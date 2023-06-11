package main

const (
	use         = "To Do"
	description = "To do list"
)

func main() {
	dic := neDIContainer(use, description)
	err := runToDo(dic)
	if err != nil {
		panic(err)
	}
}
