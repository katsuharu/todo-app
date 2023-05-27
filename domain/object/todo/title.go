package todo

import "errors"

type Title string

const titleMaxLength = 20

func NewTitle(title string) (Title, error) {
	if len(title) > titleMaxLength {
		return Title(""), errors.New("")
	}

	return Title(title), nil
}

func (t Title) String() string {
	return string(t)
}
