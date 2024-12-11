package main

import (
	"bufio"
	"fmt"
	"os"
)

func viewFile() {
	fmt.Print("Enter the file path to view: ")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Printf("Contents of %s:\n", fileName)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
