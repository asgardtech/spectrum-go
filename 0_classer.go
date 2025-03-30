package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type classer[T app.UI] struct {
	classes []string
	element T
}

func newClasser[T app.UI](element T) *classer[T] {
	return &classer[T]{
		classes: make([]string, 0),
		element: element,
	}
}

func (s *classer[T]) Class(class string) T {
	s.classes = append(s.classes, class)
	return s.element
}

func (s *classer[T]) Classes(classes ...string) T {
	s.classes = append(s.classes, classes...)
	return s.element
}

type Classer[T app.UI] interface {
	Class(class string) T
	Classes(classes ...string) T
}
