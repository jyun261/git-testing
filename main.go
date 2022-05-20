package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "commit", "-m", `"testing commits in go"`)
	cmd.Run()
	fmt.Print("Showing how to do shit.")
}
