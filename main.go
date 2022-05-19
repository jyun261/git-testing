package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "commit", "-m", `"testing commits in go"`)
	cmd.Run()
}
