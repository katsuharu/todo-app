package todo

import "github.com/google/uuid"

type ID string

func NewID() ID {
	return ID(uuid.Must(uuid.NewRandom()).String())
}

func (id ID) String() string {
	return string(id)
}
