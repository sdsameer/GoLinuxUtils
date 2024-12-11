package main

import (
	"fmt"
	"os/exec"
)

func listOfFiles() {
	cmd := exec.Command("ls", "/")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing ls command:", err)
		return
	}

	fmt.Println("The following are the list of contents available in the system /:")
	fmt.Println(string(output))
}
