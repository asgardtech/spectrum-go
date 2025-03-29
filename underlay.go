package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumUnderlay represents an sp-underlay component
type spectrumUnderlay struct {
	app.Compo

	// Properties
	PropOpen bool

	// Event handlers
	PropOnClose app.EventHandler
}

// Underlay creates a new underlay component
func Underlay() *spectrumUnderlay {
	return &spectrumUnderlay{}
}

// Open sets whether the underlay is visible
func (u *spectrumUnderlay) Open(open bool) *spectrumUnderlay {
	u.PropOpen = open
	return u
}

// OnClose sets the close event handler
func (u *spectrumUnderlay) OnClose(handler app.EventHandler) *spectrumUnderlay {
	u.PropOnClose = handler
	return u
}

// Render renders the underlay component
func (u *spectrumUnderlay) Render() app.UI {
	underlay := app.Elem("sp-underlay")

	// Set attributes
	if u.PropOpen {
		underlay = underlay.Attr("open", true)
	}

	// Add event handlers
	if u.PropOnClose != nil {
		underlay = underlay.On("close", u.PropOnClose)
	}

	return underlay
}
