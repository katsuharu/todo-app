package dao

import (
	"context"
	"time"

	"github.com/katsuharu/todo-app/domain/object/todo"
	"github.com/katsuharu/todo-app/domain/repository"
	"github.com/katsuharu/todo-app/infra/db"
)

type wrapper struct {
	db *db.DB
}

func NewTodo(db *db.DB) repository.Todo {
	return &wrapper{
		db: db,
	}
}

type TodoDTO struct {
	ID        string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var datas []*todo.Todo

func (w wrapper) Create(ctx context.Context, entity *todo.Todo) (*todo.Todo, error) {

	datas = append(datas, &todo.Todo{
		ID:        entity.ID,
		Title:     entity.Title,
		Body:      entity.Body,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	})

	return entity, nil
}

func (w wrapper) List(ctx context.Context) ([]*todo.Todo, error) {
	return datas, nil
}
