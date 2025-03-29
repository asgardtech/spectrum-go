package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTray represents an sp-tray component
type spectrumTray struct {
	app.Compo

	// Properties
	PropOpen bool

	// Content
	PropContent app.UI
	PropSlot    string

	// Event handlers
	PropOnClose app.EventHandler
}

// Tray creates a new tray component
func Tray() *spectrumTray {
	return &spectrumTray{}
}

// Open sets whether the tray is visible
func (t *spectrumTray) Open(open bool) *spectrumTray {
	t.PropOpen = open
	return t
}

// Content sets the content of the tray
func (t *spectrumTray) Content(content app.UI) *spectrumTray {
	t.PropContent = content
	return t
}

// Slot sets the slot attribute of the tray
func (t *spectrumTray) Slot(slot string) *spectrumTray {
	t.PropSlot = slot
	return t
}

// OnClose sets the close event handler
func (t *spectrumTray) OnClose(handler app.EventHandler) *spectrumTray {
	t.PropOnClose = handler
	return t
}

// Render renders the tray component
func (t *spectrumTray) Render() app.UI {
	tray := app.Elem("sp-tray")

	// Set attributes
	if t.PropOpen {
		tray = tray.Attr("open", true)
	}
	if t.PropSlot != "" {
		tray = tray.Attr("slot", t.PropSlot)
	}

	// Add event handlers
	if t.PropOnClose != nil {
		tray = tray.On("close", t.PropOnClose)
	}

	// Add content if provided
	if t.PropContent != nil {
		tray = tray.Body(t.PropContent)
	}

	return tray
}
