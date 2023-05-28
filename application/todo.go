package application

import (
	"context"
	"fmt"
	"time"

	"github.com/katsuharu/todo-app/domain/object/todo"
	"github.com/katsuharu/todo-app/domain/repository"
)

type wrapper struct {
	r repository.Todo
}

func NewTodo(r repository.Todo) Todo {
	return &wrapper{r: r}
}

type Todo interface {
	Create(ctc context.Context, title, body string) (*CreateTodoResponse, error)
	List(ctx context.Context) ([]*ListTodoResponse, error)
}

type CreateTodoResponse struct {
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListTodoResponse struct {
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w wrapper) Create(ctx context.Context, title, body string) (*CreateTodoResponse, error) {
	entity, err := todo.New(title, body, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to generate todo: %w", err)
	}

	t, err := w.r.Create(ctx, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return &CreateTodoResponse{
		Title:     t.Title.String(),
		Body:      t.Body.String(),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (w wrapper) List(ctx context.Context) ([]*ListTodoResponse, error) {
	results, err := w.r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo list: %w", err)
	}

	var response []*ListTodoResponse
	for _, v := range results {
		response = append(response, &ListTodoResponse{
			Title:     v.Title.String(),
			Body:      v.Body.String(),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return response, nil
}
