package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"sync"
)

type toDoHandler struct {
	task []string
}

func newToDoHandlerProvider() func() *toDoHandler {
	var toDo *toDoHandler
	var mu sync.Mutex
	return func() *toDoHandler {
		mu.Lock()
		defer mu.Unlock()
		toDo = newTodo()
		return toDo
	}
}

func newTodo() *toDoHandler {
	task := make([]string, 0)
	return &toDoHandler{task: task}
}

func (tdh *toDoHandler) todo(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to the To-Do List CLI!")
	fmt.Println("Enter 'help' to see available commands.")

	for {
		var command string
		fmt.Print(">> ")
		fmt.Scanln(&command)
		var arg string
		switch strings.ToLower(command) {
		case "add":
			addCmd := &cobra.Command{
				Use:   "add [task]",
				Short: "Add a new task",
				Args:  cobra.MinimumNArgs(1),
				Run:   tdh.add,
			}
			addCmd.Run(cmd, append([]string{arg}))

		case "list":
			listCmd := &cobra.Command{
				Use:   "list",
				Short: "List all tasks",
				Run:   tdh.show,
			}
			listCmd.Run(cmd, args)

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid command. Please try again or enter 'help' for assistance.")
		}
	}
}

func (tdh *toDoHandler) add(cmd *cobra.Command, args []string) {
	var arg string
	fmt.Print("task description :")
	fmt.Scanln(&arg)
	tdh.task = append(tdh.task, arg)
	fmt.Println("task added", arg)

	fmt.Println("all task")
	fmt.Println(tdh.task)
}

func (tdh *toDoHandler) show(cmd *cobra.Command, args []string) {
	for i, v := range tdh.task {
		fmt.Println(i+1, ": ", v)
	}
}
