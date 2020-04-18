package TestData

import (
	mdl "github.com/Ofla/TODO/models"
)

// TTAddTodo represents table test structure of AddTodo test
type TTAddTodo struct {
	Name     string
	Todo     mdl.Todo
	HasError bool
}

// CreateTTAddTodo creates table test for AddTodo test
func CreateTTAddTodo() []TTAddTodo {
	return []TTAddTodo{
		{
			Name: "valid request",
			Todo: mdl.Todo{
				Name:        "My todo 1",
				Description: "some description for my todo 1",
				Status:      mdl.INPROG,
				ItemType:    "Task",
				Hash:        "1234",
			},
			HasError: false,
		},
		{
			Name: "invalid request: empty id",
			Todo: mdl.Todo{
				Name:        "My todo 1",
				Description: "some description for my todo 1",
				Status:      mdl.INPROG,
				ItemType:    "Task",
				Hash:        "",
			},
			HasError: true,
		},
	}
}
