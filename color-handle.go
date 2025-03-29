package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorHandle represents an sp-color-handle component
type spectrumColorHandle struct {
	app.Compo

	// Properties
	PropColor    string
	PropDisabled bool
	PropFocused  bool
	PropOpen     bool
}

// ColorHandle creates a new color handle component
func ColorHandle() *spectrumColorHandle {
	return &spectrumColorHandle{
		PropColor: "rgba(255, 0, 0, 0.5)", // Default color
	}
}

// Color sets the color of the handle
func (ch *spectrumColorHandle) Color(color string) *spectrumColorHandle {
	ch.PropColor = color
	return ch
}

// Disabled sets whether the color handle is disabled
func (ch *spectrumColorHandle) Disabled(disabled bool) *spectrumColorHandle {
	ch.PropDisabled = disabled
	return ch
}

// Focused sets whether the color handle is focused
func (ch *spectrumColorHandle) Focused(focused bool) *spectrumColorHandle {
	ch.PropFocused = focused
	return ch
}

// Open sets whether the color handle is in open state
// When open, the color loupe can be shown above the handle
func (ch *spectrumColorHandle) Open(open bool) *spectrumColorHandle {
	ch.PropOpen = open
	return ch
}

// Render renders the color handle component
func (ch *spectrumColorHandle) Render() app.UI {
	colorHandle := app.Elem("sp-color-handle")

	// Set attributes
	if ch.PropColor != "rgba(255, 0, 0, 0.5)" {
		colorHandle = colorHandle.Attr("color", ch.PropColor)
	}
	if ch.PropDisabled {
		colorHandle = colorHandle.Attr("disabled", true)
	}
	if ch.PropFocused {
		colorHandle = colorHandle.Attr("focused", true)
	}
	if ch.PropOpen {
		colorHandle = colorHandle.Attr("open", true)
	}

	return colorHandle
}
