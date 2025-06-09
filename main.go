package main

import (
	"fmt"
	"os"

	"github.com/lllllan02/task-cli/command"
	"github.com/lllllan02/task-cli/task"
)

func printUsage() {
	fmt.Println("Usage: task-cli <command> [args...]")
	fmt.Println("Commands:")
	fmt.Println("  add \"<task name>\" - Add a new task")
	fmt.Println("  list [/done/todo/in-progress] - List all tasks")
	fmt.Println("  update <task id> \"<new task name>\" - Update a task")
	fmt.Println("  delete <task id> - Delete a task")
	fmt.Println("  mark-in-progress <task id> - Mark a task as in progress")
	fmt.Println("  mark-done <task id> - Mark a task as done")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := command.ParseCommand(os.Args[1:])
	if cmd == nil {
		printUsage()
		return
	}

	switch cmd.Name {
	case "add":
		if len(cmd.Args) == 0 {
			fmt.Println("Error: Task name is required")
			return
		}
		if err := task.AddTask(cmd.Args[0]); err != nil {
			fmt.Println("Error adding task:", err)
		}

	case "list":
		var status string
		if len(cmd.Args) > 0 {
			status = cmd.Args[0]
		}
		if err := task.ListTasks(status); err != nil {
			fmt.Println("Error listing tasks:", err)
		}

	case "update":
		if len(cmd.Args) < 2 {
			fmt.Println("Error: Task ID and new name are required")
			return
		}
		if err := task.UpdateTask(cmd.Args[0], cmd.Args[1]); err != nil {
			fmt.Println("Error updating task:", err)
		}

	case "delete":
		if len(cmd.Args) == 0 {
			fmt.Println("Error: Task ID is required")
			return
		}
		if err := task.DeleteTask(cmd.Args[0]); err != nil {
			fmt.Println("Error deleting task:", err)
		}

	case "mark-in-progress":
		if len(cmd.Args) == 0 {
			fmt.Println("Error: Task ID is required")
			return
		}
		if err := task.MarkTaskInProgress(cmd.Args[0]); err != nil {
			fmt.Println("Error marking task as in progress:", err)
		}

	case "mark-done":
		if len(cmd.Args) == 0 {
			fmt.Println("Error: Task ID is required")
			return
		}
		if err := task.MarkTaskDone(cmd.Args[0]); err != nil {
			fmt.Println("Error marking task as done:", err)
		}

	default:
		printUsage()
		return
	}
}
