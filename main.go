package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "branch", "James")
	cmd.Run()
}
