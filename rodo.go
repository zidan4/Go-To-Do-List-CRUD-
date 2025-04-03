package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			tasks = append(tasks, scanner.Text())
			fmt.Println("Task added!")
		case "2":
			fmt.Println("\nTo-Do List:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
