package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type ider[T app.UI] struct {
	id      string
	element T
}

func newIder[T app.UI](element T) *ider[T] {
	return &ider[T]{
		id:      "",
		element: element,
	}
}

func (s *ider[T]) Id(id string) T {
	s.id = id
	return s.element
}

type Ider[T app.UI] interface {
	Id(id string) T
}
