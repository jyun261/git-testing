package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "merge", "jin")
	cmd.Run()
}
