package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumButtonGroup represents an sp-button-group component
type spectrumButtonGroup struct {
	app.Compo

	// Properties
	PropVertical bool

	// Child elements
	PropButtons []app.UI
}

// ButtonGroup creates a new button group component
func ButtonGroup() *spectrumButtonGroup {
	return &spectrumButtonGroup{}
}

// Vertical sets whether the buttons are arranged vertically
func (b *spectrumButtonGroup) Vertical(vertical bool) *spectrumButtonGroup {
	b.PropVertical = vertical
	return b
}

// AddButton adds a button to the group
func (b *spectrumButtonGroup) AddButton(button app.UI) *spectrumButtonGroup {
	b.PropButtons = append(b.PropButtons, button)
	return b
}

// Buttons sets multiple buttons at once
func (b *spectrumButtonGroup) Buttons(buttons ...app.UI) *spectrumButtonGroup {
	b.PropButtons = buttons
	return b
}

// Render renders the button group component
func (b *spectrumButtonGroup) Render() app.UI {
	buttonGroup := app.Elem("sp-button-group")

	if b.PropVertical {
		buttonGroup = buttonGroup.Attr("vertical", "")
	}

	// Add buttons
	if len(b.PropButtons) > 0 {
		buttonGroup = buttonGroup.Body(b.PropButtons...)
	}

	return buttonGroup
}
