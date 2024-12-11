package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkUptime() {
	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Error reading /proc/uptime:", err)
		return
	}

	parts := strings.Fields(string(data))
	if len(parts) > 0 {
		fmt.Printf("System Uptime: %s seconds\n", parts[0])
	} else {
		fmt.Println("Could not parse uptime data.")
	}
}
