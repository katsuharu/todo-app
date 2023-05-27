package dao

import (
	"context"
	"fmt"

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

func (w wrapper) Create(ctx context.Context, entity *todo.Todo) (*todo.Todo, error) {
	q := `
	INSERT INTO todos (
		id,
		title,
		body,
		created_at,
		updated_at
	) VALUES(?, ?, ?, ?, ?)`

	if _, err := w.db.Write.ExecContext(
		ctx,
		q,
		entity.ID.String(),
		entity.Title.String(),
		entity.Body,
		entity.CreatedAt,
		entity.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to ExecContext: %w", err)
	}

	return entity, nil
}
