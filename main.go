package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "merge", "jin")
	cmd.Run()
	fmt.Print("It worked!")
}
