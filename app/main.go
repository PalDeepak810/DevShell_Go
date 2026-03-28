package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Print

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("$ ")
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		parts := strings.Fields(strings.TrimSpace(text))
		command := parts[0]
		args := parts[1:]

		if command == "exit" {
			break
		} else if command == "echo" {
			var res string
			for _, v := range args {
				res += v + " "
			}
			fmt.Print(strings.TrimSpace(res))
		} else if command == "type" {
			if args[0] == "echo" || args[0] == "exit" || args[0] == "type" {
				fmt.Printf("%s is a shell builtin", args[0])
			} else {
				fmt.Printf("%s: not found", args[0])
			}
		} else {
			fmt.Printf("%s: command not found", command)
		}

		fmt.Print("\n")
	}
}
