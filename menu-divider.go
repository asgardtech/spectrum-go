package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumMenuDivider represents an sp-menu-divider component
type SpectrumMenuDivider struct {
	app.Compo

	// Properties
	size string
}

// MenuDivider creates a new menu divider component
func MenuDivider() *SpectrumMenuDivider {
	return &SpectrumMenuDivider{}
}

// Size sets the size of the divider
func (m *SpectrumMenuDivider) Size(size string) *SpectrumMenuDivider {
	m.size = size
	return m
}

// Render renders the menu divider component
func (m *SpectrumMenuDivider) Render() app.UI {
	menuDivider := app.Elem("sp-menu-divider")

	// Set size if provided
	if m.size != "" {
		menuDivider = menuDivider.Attr("size", m.size)
	}

	return menuDivider
}
