package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"path/filepath"

	"time"

	"github.com/fsnotify/fsnotify"
)

const version = "0.0.4"

var vagrantFile = `
Vagrant.configure("2") do |config|
  config.vm.box = "%s"
  config.vm.network :private_network, ip: "%s"
  config.vm.synced_folder ".", "/vagrant", type: "nfs"
  config.vm.provider "virtualbox" do |v|
	v.memory = %d
	v.cpus = %d
  end

  config.vm.provision "shell",
	inline: "curl -SsL %s | sh",
	privileged: false
end
`
var ip = flag.String("ip", "", "the internal IP address of this agent")
var name = flag.String("name", "", "the name of this agent")
var baseImage = flag.String("image", "", "the vagrant base box of this agent")
var showVersion = flag.Bool("version", false, "prints the version")
var memory = flag.Int("memory", 2048, "sets the memory of the VM")
var cpu = flag.Int("cpu", 3, "sets the number of cpus for the VM")
var provisionScript = flag.String("provision", "", "Script to provision the VM, runs using the shell provisioner")
var help = flag.Bool("help", false, "prints help message")

func watchForRollback() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	rollbackFile, err := filepath.Abs("rollback.txt")
	if err != nil {
		log.Fatal(err)
	}
	sigchan := make(chan os.Signal, 10)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					if event.Name == rollbackFile {
						execute("vagrant", "sandbox", "rollback")
					}
				}
			case err := <-watcher.Errors:
				if err != nil {
					log.Println("error:", err)
				}
			case <-sigchan:
				log.Println("Received Kill Signal. Cleaning up before exit..")
				log.Printf("Halting VM %s\n", *name)
				execute("vagrant", "halt")
				log.Printf("Destroying VM %s\n", *name)
				execute("vagrant", "destroy")
				log.Println("Cleaning files")
				os.Remove("Vagrantfile")
				os.Remove("rollback.txt")
				log.Println("Done, exiting..")
				os.Exit(0)
			}
		}
	}()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(wd)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *ip == "" || *name == "" || *baseImage == "" || *provisionScript == "" {
		fmt.Println("Usage: cider --ip <agent_ip> --name <agent_name> --image <base_vegrant_box_name> --provision <provision_sh_script>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	checkEnvSet("GAUGE_DOWNLOADS_IP")
	checkEnvSet("GO_SERVER_URL")
	checkEnvSet("AGENT_AUTO_REGISTER_KEY")

	execute("vagrant", "plugin", "install", "sahara")

	writeVagrantFile(*baseImage, *ip, *memory, *cpu, *provisionScript)

	execute("vagrant", "up", "--provision")

	execute("vagrant", "ssh", "-c", fmt.Sprintf("echo -e \"agent.auto.register.key=%s\nagent.auto.register.resources=FT,UT,darwin\nagent.auto.register.environments=all\nagent.auto.register.hostname=%s\" > go-agent-17.4.0/config/autoregister.properties", os.Getenv("AGENT_AUTO_REGISTER_KEY"), *name))

	execute("vagrant", "ssh", "-c", fmt.Sprintf("sudo /bin/sh -c \"echo '%s downloads.gauge.org' >> /etc/hosts\"", os.Getenv("GAUGE_DOWNLOADS_IP")))

	execute("vagrant", "ssh", "-c", fmt.Sprintf("GO_SERVER_URL=%s nohup /bin/sh -c \"sh /Users/vagrant/go-agent-17.4.0/agent.sh & disown\"", os.Getenv("GO_SERVER_URL")))

	fmt.Println("Waiting for Go Agent to start.")
	time.Sleep(20000)

	execute("vagrant", "sandbox", "on")

	watchForRollback()
}

func checkEnvSet(envName string) {
	if _, ok := os.LookupEnv(envName); !ok {
		fmt.Printf("%s env is not set.\n", envName)
		os.Exit(1)
	}
}

func writeVagrantFile(baseBox, ip string, memory, cpu int, provisionScript string) {
	err := ioutil.WriteFile("Vagrantfile", []byte(fmt.Sprintf(vagrantFile, baseBox, ip, memory, cpu, provisionScript)), 0644)
	if err != nil {
		log.Fatalf("Error writing Vagrantfile `%s`", err.Error())
	}
}

func execute(command string, args ...string) {
	cmd := exec.Command(command, args...)
	log.Println(fmt.Sprintf("Running %s %s", command, strings.Join(args, " ")))
	wd, err := os.Getwd()
	if err == nil {
		cmd.Dir = wd
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Command `%s %s` failed. Error: %s", command, strings.Join(args, " "), err.Error())
	}
}
