package main

import (
	"fmt"
	"os"
)

func main() {
	cliArguments := os.Args
	argsLength := len(cliArguments)
	command := make([]string, 0, argsLength-1)
	var action string
	actionArguments := make([]string, 0, argsLength-2)

	if argsLength > 1 {
		command = cliArguments[1:]
		action = command[0]
		if len(command) > 1 {
			actionArguments = command[1:]
		}
	}

	switch action {
	case "add":
		fmt.Println("add action")
		fmt.Println("action: " + action)
		fmt.Println("action arguments : ")
		for _, arg := range actionArguments {
			fmt.Println(arg)
		}
	case "compute":
		fmt.Println("compute action")
	}
}
