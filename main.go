package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/soupaulodev/togo/task"
)

// Initialize the application, start the task manager and read user input
func main() {
	manager := task.NewTaskManager()
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- TODO CLI ---")
		fmt.Println("1. Create task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Update task status")
		fmt.Println("4. Remove task")
		fmt.Println("5. Quit")
		fmt.Print("Choose an option: ")

		reader.Scan()
		choice := reader.Text()

		fmt.Print("\033[H\033[2J")
		
		switch choice {
		case "1":
			fmt.Print("Title: ")
			reader.Scan()
			title := reader.Text()
			
			fmt.Print("Description: ")
			reader.Scan()
			description := reader.Text()
			
			fmt.Print("Priority (1-5): ")
			reader.Scan()
			priorityInt, _ := strconv.Atoi(reader.Text())
			priority := uint8(priorityInt)
			
			fmt.Print("Due Date (DD/MM/YYYY HH:MM): ")
			reader.Scan()
			dueDate, err := time.Parse("02/01/2006 15:04", reader.Text())
			if err != nil {
				fmt.Println("Invalid date format. Please use DD/MM/YYYY HH:MM")
				return
			}
			
			manager.CreateTask(title, description, priority, dueDate)

		case "2":
			manager.ShowTask()

		case "3":
			fmt.Print("Task ID: ")
			reader.Scan()
			id, _ := strconv.Atoi(reader.Text())

			fmt.Print("New Status (pending/done): ")
			reader.Scan()
			status := task.Status(reader.Text())

			manager.UpdateTask(id, status)
		
		case "4":
			fmt.Print("Task ID: ")
			reader.Scan()
			id, _ := strconv.Atoi(reader.Text())

			manager.RemoveTask(id)

		case "5":
			fmt.Println("Quitting...")
			fmt.Print("\033[H\033[2J")
			os.Exit(0)
		
		default:
			fmt.Println("Invalid option!")
		}
	}
}