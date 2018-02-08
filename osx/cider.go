package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"path/filepath"

	"time"

	"github.com/fsnotify/fsnotify"
)

var vagrantFile = `
Vagrant.configure("2") do |config|
  config.vm.box = "%s"
  config.vm.network :private_network, ip: "%s"
  config.vm.synced_folder ".", "/vagrant", type: "nfs"
  config.vm.provision "shell",
    inline: "curl -SsL https://raw.githubusercontent.com/getgauge/infrastructure/master/osx/osx.sh | sh"
end
`

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
	if len(os.Args) < 3 {
		fmt.Println("Missing box name/IP as program argument.\nUsage: cider <box name> <IP>")
		os.Exit(1)
	}

	name := os.Args[1]

	execute("vagrant", "plugin", "install", "sahara")

	writeVagrantFile(name, os.Args[2])

	execute("vagrant", "up", "--provision")

	execute("vagrant", "ssh", "-c", fmt.Sprintf("echo -e \"agent.auto.register.key=%s\nagent.auto.register.resources=FT,UT,darwin,installers\nagent.auto.register.hostname=%s\" > go-agent-17.4.0/config/autoregister.properties", os.Getenv("AGENT_AUTO_REGISTER_KEY"), os.Getenv("AGENT_NAME")))

	execute("vagrant", "ssh", "-c", fmt.Sprintf("sudo /bin/sh -c \"echo '%s downloads.getgauge.io' >> /etc/hosts\"", os.Getenv("GAUGE_DOWNLOADS_IP")))

	execute("vagrant", "ssh", "-c", fmt.Sprintf("GO_SERVER_URL=%s nohup /bin/sh -c \"sh /Users/vagrant/go-agent-17.4.0/agent.sh & disown\"", os.Getenv("GO_SERVER_URL")))

	fmt.Println("Waiting for Go Agent to start.")
	time.Sleep(20000)

	execute("vagrant", "sandbox", "on")

	watchForRollback()
}

func writeVagrantFile(name, ip string) {
	err := ioutil.WriteFile("Vagrantfile", []byte(fmt.Sprintf(vagrantFile, name, ip)), 0644)
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
