package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumButtonGroup represents an sp-button-group component
type SpectrumButtonGroup struct {
	app.Compo

	// Properties
	vertical bool

	// Child elements
	buttons []app.UI
}

// ButtonGroup creates a new button group component
func ButtonGroup() *SpectrumButtonGroup {
	return &SpectrumButtonGroup{}
}

// Vertical sets whether the buttons are arranged vertically
func (b *SpectrumButtonGroup) Vertical(vertical bool) *SpectrumButtonGroup {
	b.vertical = vertical
	return b
}

// AddButton adds a button to the group
func (b *SpectrumButtonGroup) AddButton(button app.UI) *SpectrumButtonGroup {
	b.buttons = append(b.buttons, button)
	return b
}

// Buttons sets multiple buttons at once
func (b *SpectrumButtonGroup) Buttons(buttons ...app.UI) *SpectrumButtonGroup {
	b.buttons = buttons
	return b
}

// Render renders the button group component
func (b *SpectrumButtonGroup) Render() app.UI {
	buttonGroup := app.Elem("sp-button-group")

	if b.vertical {
		buttonGroup = buttonGroup.Attr("vertical", "")
	}

	// Add buttons
	if len(b.buttons) > 0 {
		buttonGroup = buttonGroup.Body(b.buttons...)
	}

	return buttonGroup
}
