package memdatabase

import (
	"fmt"
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/errors"
	vld "github.com/Ofla/TODO/helpers/validators"
	mdl "github.com/Ofla/TODO/models"
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

// Repo is a struct which represents
type Repo struct {
	*memdb.MemDB
	log  *logrus.Logger
	conf *config.Database
}

func NewTodoRepo(conf *config.Database, log *logrus.Logger) *Repo {
	db := ConnectToDB(conf)
	repo := Repo{
		MemDB: db,
		log:   log,
		conf:  conf,
	}
	return &repo
}

// AddTodo is a func which adds a new todo item
func (repo *Repo) AddTodo(item *mdl.Todo) error {
	if ok := vld.IsNilOrEmpty(item); ok {
		return errors.InvalidInputError
	}
	// Create a write transaction
	txn := repo.MemDB.Txn(true)
	if err := txn.Insert(repo.conf.Tablename, item); err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", item.Hash)).Error())
		return errors.DBCreateError
	}
	// Commit the transaction
	txn.Commit()
	return nil
}

// RemoveTodo is a unc which deletes a specific todo item
func (repo *Repo) RemoveTodo(id string) (bool, error) {

	txn := repo.MemDB.Txn(false)
	oldItem, err := txn.First(repo.conf.Tablename, "id", id)
	if err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBGetError
	}
	if oldItem == nil {
		repo.log.Errorf(errors.CreateInternalError(errors.NotFoundError, fmt.Sprintf("%v", id)).Error())
		return false, nil
	}
	oldTodo := oldItem.(mdl.Todo)
	itemToDelete := mdl.Todo{
		Name:        oldTodo.Name,
		Status:      oldTodo.Status,
		Description: oldTodo.Description,
		ItemType:    oldTodo.ItemType,
		Hash:        id,
	}

	// Create a write transaction
	txn = repo.MemDB.Txn(true)
	if err := txn.Delete(repo.conf.Tablename, itemToDelete); err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBDeleteError
	}
	txn.Commit()
	return true, nil
}

// UpdateTodo is a unc which updates a specific todo item
func (repo *Repo) UpdateTodo(id string, item mdl.Todo) (bool, error) {

	txn := repo.MemDB.Txn(false)
	defer txn.Abort()
	oldItem, err := txn.First(repo.conf.Tablename, "id", id)
	if err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBGetError
	}

	if oldItem == nil {
		repo.log.Errorf(errors.CreateInternalError(errors.NotFoundError, fmt.Sprintf("%v", id)).Error())
		return false, nil
	}
	oldTodo := oldItem.(mdl.Todo)
	itemToDelete := mdl.Todo{
		Name:        oldTodo.Name,
		Status:      oldTodo.Status,
		Description: oldTodo.Description,
		ItemType:    oldTodo.ItemType,
		Hash:        id,
	}
	txn = repo.MemDB.Txn(true)
	if err := txn.Insert(repo.conf.Tablename, item); err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", item.Hash)).Error())
		return false, errors.DBCreateError
	}
	// having much more data is better than having none
	if err := txn.Delete(repo.conf.Tablename, itemToDelete); err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBDeleteError
	}

	txn.Commit()
	return true, nil
}

// UpdateStatus is a func which updates the status of a  specific todo item
func (repo *Repo) UpdateStatus(id string, newStatus mdl.TodoStatus) (bool, error) {

	txn := repo.MemDB.Txn(false)
	oldItem, err := txn.First(repo.conf.Tablename, "id", id)
	if err != nil {
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBGetError

	}

	if oldItem == nil {
		repo.log.Errorf(errors.CreateInternalError(errors.NotFoundError, fmt.Sprintf("%v", id)).Error())
		return false, nil
	}
	oldTodo := oldItem.(mdl.Todo)
	newItem := mdl.Todo{
		Name:        oldTodo.Name,
		Status:      newStatus,
		Description: oldTodo.Description,
		ItemType:    oldTodo.ItemType,
		Hash:        id,
	}
	// Create a write transaction
	txn = repo.MemDB.Txn(true)
	if err := txn.Insert(repo.conf.Tablename, newItem); err != nil {
		txn.Commit()
		repo.log.Errorf(errors.CreateInternalError(err, fmt.Sprintf("%v", id)).Error())
		return false, errors.DBUpdateStatusError
	}

	// Commit the transaction
	txn.Commit()
	return true, nil
}

// FindAllTodos is a unc which feths all todo items
func (repo *Repo) FindAllTodos() (mdl.ListOfIds, mdl.ListOfTodo, error) {

	// Create a write transaction
	txn := repo.MemDB.Txn(false)
	defer txn.Abort()

	var (
		todosL    []mdl.Todo
		IdsOfTodo []string
		loT       = mdl.ListOfTodo{}
		IDs       = mdl.ListOfIds{}
	)
	// List all the todos
	allItems, err := txn.Get(repo.conf.Tablename, "id")
	if err != nil {
		repo.log.Errorf(errors.CreateInternalError(err).Error())
		return IDs, loT, errors.DBGetAllError
	}
	for item := allItems.Next(); item != nil; item = allItems.Next() {
		todo := item.(mdl.Todo)
		t := mdl.Todo{
			Name:        todo.Name,
			Description: todo.Description,
			Status:      todo.Status,
			ItemType:    todo.ItemType,
		}
		todosL = append(todosL, t)
		IdsOfTodo = append(IdsOfTodo, todo.Hash)
	}

	loT = mdl.ListOfTodo{Todos: todosL}
	IDs = mdl.ListOfIds{IDs: IdsOfTodo}
	return IDs, loT, nil

}
