package main

import (
	"fmt"
	"io/ioutil"
)

func displayCPUInfo() {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error reading /proc/cpuinfo:", err)
		return
	}

	fmt.Println("CPU Information:")
	fmt.Println(string(data))
}
