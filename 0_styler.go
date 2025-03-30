package sp

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type styler[T app.UI] struct {
	styles  map[string]string
	element T
}

func newStyler[T app.UI](element T) *styler[T] {
	return &styler[T]{
		styles:  make(map[string]string),
		element: element,
	}
}

func (s *styler[T]) Style(k, format string, v ...any) T {
	s.styles[k] = fmt.Sprintf(format, v...)
	return s.element
}

func (s *styler[T]) Styles(styles map[string]string) T {
	for k, v := range styles {
		s.styles[k] = v
	}
	return s.element
}

type Styler[T app.UI] interface {
	Style(k, format string, v ...any) T
	Styles(styles map[string]string) T
}
