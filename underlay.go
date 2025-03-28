package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumUnderlay represents an sp-underlay component
type SpectrumUnderlay struct {
	app.Compo

	// Properties
	open bool

	// Event handlers
	onClose app.EventHandler
}

// Underlay creates a new underlay component
func Underlay() *SpectrumUnderlay {
	return &SpectrumUnderlay{}
}

// Open sets whether the underlay is visible
func (u *SpectrumUnderlay) Open(open bool) *SpectrumUnderlay {
	u.open = open
	return u
}

// OnClose sets the close event handler
func (u *SpectrumUnderlay) OnClose(handler app.EventHandler) *SpectrumUnderlay {
	u.onClose = handler
	return u
}

// Render renders the underlay component
func (u *SpectrumUnderlay) Render() app.UI {
	underlay := app.Elem("sp-underlay")

	// Set attributes
	if u.open {
		underlay = underlay.Attr("open", true)
	}

	// Add event handlers
	if u.onClose != nil {
		underlay = underlay.On("close", u.onClose)
	}

	return underlay
}
