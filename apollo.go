package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var H_HOME string = os.Getenv("HOME")
var MOUNT string = fmt.Sprintf("%s/apollo", H_HOME)
var G_HOME = "/home/vagrant"

var s = fmt.Sprintf

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
	g_cmd := prefixCmd("/bin/zsh")
	execRemoteCommand(g_cmd)
}

func nonInteractive(usercmd string) {
	usercmd = prefixCmd(usercmd)
	fmt.Printf("\x1B[42m\x1B[30m\x1B[1m Running in spaaace! \x1B[0m\n")
	execRemoteCommand(usercmd)
	fmt.Printf("\x1B[42m\x1B[30m\x1B[1m ~ Fin. ~ \x1B[0m\n")
}

func up() {
	// TODO: Create the ~/apollo directory if not existing.
	// TODO: Also add down, and reload methods
	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Preparing Apollo. \x1B[0m\n")
	execCommand([]string{"vagrant", "up"})
	execCommand([]string{"sshfs", "vagrant@localhost:",
		s("%s/apollo", H_HOME),
		"-p", "2222",
		"-ofollow_symlinks,auto_cache,reconnect,volname=apollo"})
	fmt.Printf("\x1B[45m\x1B[30m\x1B[1m Ready for liftoff! \x1B[0m\n")
}

func main() {
	flag.Parse()
	if flag.Arg(0) == "i" {
		interactive()
	} else if flag.Arg(0) == "up" {
		up()
	} else {
		nonInteractive(strings.Join(flag.Args(), " "))
	}
}
