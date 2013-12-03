package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var VERSION = "0.0"
var H_HOME string = os.Getenv("HOME")
var MOUNT string = fmt.Sprintf("%s/apollo", H_HOME)
var G_HOME = "/home/vagrant"

var s = fmt.Sprintf

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func execCommand(cmdslice []string) {
	cmd := exec.Command(cmdslice[0], cmdslice[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func auth() {
	cmd := []string{"ssh-copy-id", "-p", "2222", "vagrant@localhost"}
	execCommand(cmd)
}

func execRemoteCommand(cmdstring string) {
	// In order for this to work like a login shell, the .zshenv on the vm
	// must source the .zshrc file. That way we can get access to the aliases
	// and other path modifications that happen in the .zshrc.
	cmd := []string{"ssh", "-t", "-p", "2222", "-q",
		"vagrant@localhost", cmdstring}
	execCommand(cmd)
}

func prefixCmd(cmdstr string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cd_cmd := s("cd %s; ", pwd)
	cd_cmd = strings.Replace(cd_cmd, MOUNT, G_HOME, -1)
	return (cd_cmd + cmdstr)
}

func interactive() {
	// TODO: Make the shell configurable
	g_cmd := prefixCmd("/bin/bash")
	execRemoteCommand(g_cmd)
}

func nonInteractive(usercmd string) {
	usercmd = prefixCmd(usercmd)
	usercmd = strings.Replace(usercmd, MOUNT, G_HOME, -1)
	fmt.Printf("\x1B[42m\x1B[30m\x1B[1m Running in spaaace! \x1B[0m\n")
	execRemoteCommand(usercmd)
	fmt.Printf("\x1B[42m\x1B[30m\x1B[1m ~ Fin. ~ \x1B[0m\n")
}

func up() {
	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Preparing Apollo. \x1B[0m\n")

	if !dirExists(MOUNT) {
		fmt.Printf(s("\x1B[45m\x1B[30m\x1B[1m Creating directory for mount at %s. \x1B[0m\n", MOUNT))
		err := os.MkdirAll(MOUNT, 0777)
		if err != nil {
			log.Fatal("The directory could not be created.", err)
		}
	}

	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Starting Vagrant. \x1B[0m\n")
	execCommand([]string{"vagrant", "up"})
	fmt.Printf(s("\x1B[45m\x1B[30m\x1B[1m Mounting apollo at %s. \x1B[0m\n", MOUNT))
	auth()
	execCommand([]string{"sshfs", "vagrant@localhost:",
		MOUNT,
		"-p", "2222",
		"-ofollow_symlinks,auto_cache,reconnect,volname=apollo"})

	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Ready for liftoff! \x1B[0m\n")
}

func down() {
	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Apollo is making re-entry. \x1B[0m\n")
	execCommand([]string{"vagrant", "halt"})
	execCommand([]string{"umount", "-f", MOUNT})
	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m The Eagle has landed. \x1B[0m\n")
}

func version() {
	fmt.Printf(s("Apollo version: %s\n", VERSION))
}

func main() {
	flag.Parse()
	if flag.Arg(0) == "i" {
		interactive()
	} else if flag.Arg(0) == "up" {
		up()
	} else if flag.Arg(0) == "down" {
		down()
	} else if flag.Arg(0) == "auth" {
		auth()
	} else if flag.Arg(0) == "" {
		version()
	} else {
		nonInteractive(strings.Join(flag.Args(), " "))
	}
}
