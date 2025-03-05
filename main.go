package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	id       int
	taskName string
	status   bool
}

var taskList []task

func addTask() {
	fmt.Print("Enter your task name: ")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	newTask := task{
		id:       len(taskList) + 1,
		taskName: taskName,
		status:   false,
	}

	taskList = append(taskList, newTask)
	fmt.Println("Task added successfully!")
	fmt.Println("")

}

func showTask() {
	if len(taskList) == 0 {
		fmt.Println("Tasklist is empty")
		fmt.Println("")
		return
	}

	fmt.Println("Here is your tasklist:")
	for _, task := range taskList {
		fmt.Println("----------------------------")
		fmt.Println("ID: ", task.id)
		fmt.Println("Task Name: ", task.taskName)
		fmt.Print("Status: ")
		if task.status {
			fmt.Println("Completed")
		} else {
			fmt.Println("Not completed yet")
		}
		fmt.Println("----------------------------")
	}
	fmt.Println("")
}

func updateTask() {
	if len(taskList) == 0 {
		fmt.Println("Tasklist is empty")
		fmt.Println("")
		return
	}

	fmt.Print("Enter the task ID you want to update: ")
	var taskID int
	fmt.Scanln(&taskID)

	for i := range taskList {
		if taskList[i].id == taskID {
			taskList[i].status = !taskList[i].status
			fmt.Println("Task updated successfully!")
			fmt.Println("")
			return
		}
	}

}

func deleteTask() {
	if len(taskList) == 0 {
		fmt.Println("Tasklist is empty")
		fmt.Println("")
		return
	}

	fmt.Print("Enter the task ID you want to delete: ")
	var taskID int
	fmt.Scanln(&taskID)

	for i := range taskList {
		if taskList[i].id == taskID {
			taskList = append(taskList[:i], taskList[i+1:]...)
			fmt.Println("Task deleted successfully!")
			fmt.Println("")
			break
		}
	}

	for i := range taskList {
		taskList[i].id = i + 1
	}
}

func main() {
	var option int

	for option != 5 {
		fmt.Println("Welcome to Your Tasklist!")
		fmt.Println("1. Add task")
		fmt.Println("2. Show tasklist")
		fmt.Println("3. Update task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Choose your action : ")
		fmt.Scanln(&option)

		fmt.Println("")

		switch option {
		case 1:
			addTask()
		case 2:
			showTask()
		case 3:
			updateTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Goodbye! Have a nice day!")
		default:
			fmt.Println("Invalid option")
		}
	}
}
