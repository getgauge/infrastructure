package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var vagrantFile = `
Vagrant.configure("2") do |config|
  config.vm.box = "%s"
  config.vm.network :private_network, ip: "192.168.33.10"
  config.vm.synced_folder ".", "/vagrant", type: "nfs"
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
	if len(os.Args) < 2 {
		fmt.Println("Missing box name as program argument.\nUsage: gvm <box name>")
		os.Exit(1)
	}

	name := os.Args[1]

	execute("vagrant", "plugin", "install", "sahara")

	writeVagrantFile(name)

	execute("vagrant", "up")

	execute("vagrant", "sandbox", "on")

	watchForRollback()
}

func writeVagrantFile(name string) {
	err := ioutil.WriteFile("Vagrantfile", []byte(fmt.Sprintf(vagrantFile, name)), 0644)
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
