package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "checkout", "jin")
	cmd.Run()
}
