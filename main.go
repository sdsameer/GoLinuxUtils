package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Select a project to run:")
	fmt.Println("1. Uptime Checker")
	fmt.Println("2. File Viewer")
	fmt.Println("3. Disk Usage Checker")
	fmt.Println("4. CPU Information Display")
	fmt.Println("5. List of contents in the system")
	fmt.Print("Enter your choice (1-4): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _, _ := reader.ReadLine()

	switch string(choice) {
	case "1":
		checkUptime()
	case "2":
		viewFile()
	case "3":
		checkDiskUsage()
	case "4":
		displayCPUInfo()
	case "5":
		listOfFiles()
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}
