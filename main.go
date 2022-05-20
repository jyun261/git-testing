package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "merge", "jin")
	cmd.Run()
	fmt.Print("It worked!")
	fmt.Print("Showing how to do shit.")
}
