package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumColorHandle represents an sp-color-handle component
type SpectrumColorHandle struct {
	app.Compo

	// Properties
	color    string
	disabled bool
	focused  bool
	open     bool
}

// ColorHandle creates a new color handle component
func ColorHandle() *SpectrumColorHandle {
	return &SpectrumColorHandle{
		color: "rgba(255, 0, 0, 0.5)", // Default color
	}
}

// Color sets the color of the handle
func (ch *SpectrumColorHandle) Color(color string) *SpectrumColorHandle {
	ch.color = color
	return ch
}

// Disabled sets whether the color handle is disabled
func (ch *SpectrumColorHandle) Disabled(disabled bool) *SpectrumColorHandle {
	ch.disabled = disabled
	return ch
}

// Focused sets whether the color handle is focused
func (ch *SpectrumColorHandle) Focused(focused bool) *SpectrumColorHandle {
	ch.focused = focused
	return ch
}

// Open sets whether the color handle is in open state
// When open, the color loupe can be shown above the handle
func (ch *SpectrumColorHandle) Open(open bool) *SpectrumColorHandle {
	ch.open = open
	return ch
}

// Render renders the color handle component
func (ch *SpectrumColorHandle) Render() app.UI {
	colorHandle := app.Elem("sp-color-handle")

	// Set attributes
	if ch.color != "rgba(255, 0, 0, 0.5)" {
		colorHandle = colorHandle.Attr("color", ch.color)
	}
	if ch.disabled {
		colorHandle = colorHandle.Attr("disabled", true)
	}
	if ch.focused {
		colorHandle = colorHandle.Attr("focused", true)
	}
	if ch.open {
		colorHandle = colorHandle.Attr("open", true)
	}

	return colorHandle
}
