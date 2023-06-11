package main

func runToDo(dic *diContainer) error {
	tdh := dic.toDoHandler()
	cliCmd := dic.cliDI.CLI()
	cliCmd.Run = tdh.todo
	err := cliCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
