package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumMenuDivider represents an sp-menu-divider component
type spectrumMenuDivider struct {
	app.Compo

	// Properties
	PropSize string
}

// MenuDivider creates a new menu divider component
func MenuDivider() *spectrumMenuDivider {
	return &spectrumMenuDivider{}
}

// Size sets the size of the menu divider
func (m *spectrumMenuDivider) Size(size string) *spectrumMenuDivider {
	m.PropSize = size
	return m
}

// Render renders the menu divider component
func (m *spectrumMenuDivider) Render() app.UI {
	menuDivider := app.Elem("sp-menu-divider")

	// Set attributes
	if m.PropSize != "" {
		menuDivider = menuDivider.Attr("size", m.PropSize)
	}

	return menuDivider
}
