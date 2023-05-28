package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        ID
	Title     Title
	Body      Body
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(title, body string, currentTime time.Time) (*Todo, error) {
	t, err := NewTitle(title)
	if err != nil {
		return nil, fmt.Errorf("failed to generate title: %w", err)
	}

	b, err := NewBody(body)
	if err != nil {
		return nil, fmt.Errorf("failed to generate body: %w", err)
	}

	return &Todo{
		ID:        NewID(),
		Title:     t,
		Body:      b,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}, nil
}
