package models

// Todo is a struct
type Todo struct {
	Name        string
	Description string
	Status      TodoStatus
	ItemType    string
	Hash        string
}

// TodoStatus is a type defining the status of a todo
type TodoStatus string

const (
	// INPROG is a status of the the todo item which means the todo is in Progress
	INPROG TodoStatus = "IN PROGRESS"
	// DONE is a status of the the todo item which means te todo is Done/finished
	DONE TodoStatus = "DONE"
	// NOTYET is a status of the the todo item which means the todo is not started yet
	NOTYET TodoStatus = "TODO"
)

// Reply is a struct which contains the response of the methods
type Reply struct {
	Message string
	err     bool
}

// ListOfTodo is a struct which refers to an array of Todo items
type ListOfTodo struct {
	Todos []Todo
}

// ListOfIds is a struct which refers to an array of the IDs of the Todo items
type ListOfIds struct {
	IDs []string
}