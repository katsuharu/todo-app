package todo

import "errors"

type Body string

const bodyMaxLength = 255

func (b Body) String() string {
	return string(b)
}

func NewBody(body string) (Body, error) {
	if len(body) > bodyMaxLength {
		return Body(""), errors.New("body character limit exceeded")
	}

	return Body(body), nil
}
