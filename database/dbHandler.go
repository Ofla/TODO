package database

import (
	"errors"
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/database/memdatabase"
	mdl "github.com/Ofla/TODO/models"
	"github.com/sirupsen/logrus"
)

// CategoryHandler holds functions related to category database
type DbHandler interface {
	AddTodo(item *mdl.Todo) error
	RemoveTodo(id string) (bool, error)
	UpdateTodo(id string, item mdl.Todo) (bool, error)
	UpdateStatus(id string, newStatus mdl.TodoStatus) (bool, error)
	FindAllTodos() (mdl.ListOfIds, mdl.ListOfTodo, error)
}

func CreateDbHandler(conf *config.Database, log *logrus.Logger) (DbHandler, error) {
	var dbHandler DbHandler
	switch conf.Type {
	case "memdb":
		dbHandler = memdatabase.NewTodoRepo(conf, log)
	default:
		return nil, errors.New("invalid DB type")
	}
	return dbHandler, nil
}
