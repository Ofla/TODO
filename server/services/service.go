package services

import (
	"context"
	svc "github.com/Ofla/TODO/proto/v1"
)

func (srv *SvcRunner) AddTodo(ctx context.Context, todo *svc.Todo) (*svc.Reply, error) {
	panic("implement me")
}

func (srv *SvcRunner) FindAll(message *svc.FindAllMessage, server svc.TodoService_FindAllServer) error {
	panic("implement me")
}

func (srv *SvcRunner) RemoveTodo(ctx context.Context, message *svc.RemoveTodoMessage) (*svc.Reply, error) {
	panic("implement me")
}

func (srv *SvcRunner) UpdateTodo(ctx context.Context, message *svc.UpdateTodoMessage) (*svc.Reply, error) {
	panic("implement me")
}

func (srv *SvcRunner) UpdateStatusOfTodo(ctx context.Context, message *svc.UpdateStatusOfTodoMessage) (*svc.Reply, error) {
	panic("implement me")
}
