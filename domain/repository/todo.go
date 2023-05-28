package repository

import (
	"context"

	"github.com/katsuharu/todo-app/domain/object/todo"
)

type Todo interface {
	Create(ctx context.Context, todo *todo.Todo) (*todo.Todo, error)
}
