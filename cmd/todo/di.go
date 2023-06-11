package main

import (
	"github.com/spf13/cobra"
	"task-manager/cli"
)

type diContainer struct {
	cliDI interface {
		CLI() *cobra.Command
	}
	toDoHandler func() *toDoHandler
}

func neDIContainer(use, description string) *diContainer {
	dic := &diContainer{}
	dic.cliDI = cli.NewDIContainer(use, description)
	dic.toDoHandler = newToDoHandlerProvider()
	return dic
}
