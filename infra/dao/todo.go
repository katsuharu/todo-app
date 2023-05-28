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

var datas []TodoDTO

func (w wrapper) Create(ctx context.Context, entity *todo.Todo) (*todo.Todo, error) {

	datas = append(datas, TodoDTO{
		ID:        entity.ID.String(),
		Title:     entity.Title.String(),
		Body:      entity.Body.String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	})

	return entity, nil
}
