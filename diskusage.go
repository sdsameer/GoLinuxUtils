package main

import (
	"fmt"
	"os/exec"
)

func checkDiskUsage() {
	cmd := exec.Command("df", "-h", "/")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing df command:", err)
		return
	}

	fmt.Println("Disk usage for /:")
	fmt.Println(string(output))
}
