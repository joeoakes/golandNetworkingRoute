package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("route", "print") //Windows
	//cmd := exec.Command("ip", "route") // This is for Unix-like systems
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	}
	fmt.Println("Route table:\n", out.String())
}
