package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumDialogBase represents an sp-dialog-base component
type spectrumDialogBase struct {
	app.Compo

	// Properties
	PropOpen        bool
	PropDismissable bool
	PropMode        DialogMode
	PropResponsive  bool
	PropUnderlay    bool

	// Child content (usually an sp-dialog)
	PropChild app.UI

	// Event handlers
	PropOnClose app.EventHandler
}

// DialogBase creates a new dialog base component
func DialogBase() *spectrumDialogBase {
	return &spectrumDialogBase{}
}

// Open sets whether the dialog base is open
func (db *spectrumDialogBase) Open(open bool) *spectrumDialogBase {
	db.PropOpen = open
	return db
}

// Dismissable sets whether the dialog is dismissable
func (db *spectrumDialogBase) Dismissable(dismissable bool) *spectrumDialogBase {
	db.PropDismissable = dismissable
	return db
}

// Mode sets the display mode of the dialog
func (db *spectrumDialogBase) Mode(mode DialogMode) *spectrumDialogBase {
	db.PropMode = mode
	return db
}

// Responsive sets whether the dialog should be responsive
func (db *spectrumDialogBase) Responsive(responsive bool) *spectrumDialogBase {
	db.PropResponsive = responsive
	return db
}

// Underlay sets whether the dialog should have an underlay
func (db *spectrumDialogBase) Underlay(underlay bool) *spectrumDialogBase {
	db.PropUnderlay = underlay
	return db
}

// Child sets the child element (usually an sp-dialog)
func (db *spectrumDialogBase) Child(child app.UI) *spectrumDialogBase {
	db.PropChild = child
	return db
}

// OnClose sets the close event handler
func (db *spectrumDialogBase) OnClose(handler app.EventHandler) *spectrumDialogBase {
	db.PropOnClose = handler
	return db
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (db *spectrumDialogBase) Fullscreen() *spectrumDialogBase {
	return db.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (db *spectrumDialogBase) FullscreenTakeover() *spectrumDialogBase {
	return db.Mode(DialogModeFullscreenTakeover)
}

// Render renders the dialog base component
func (db *spectrumDialogBase) Render() app.UI {
	dialogBase := app.Elem("sp-dialog-base")

	// Add properties
	if db.PropOpen {
		dialogBase = dialogBase.Attr("open", true)
	}
	if db.PropDismissable {
		dialogBase = dialogBase.Attr("dismissable", true)
	}
	if db.PropMode != "" {
		dialogBase = dialogBase.Attr("mode", string(db.PropMode))
	}
	if db.PropResponsive {
		dialogBase = dialogBase.Attr("responsive", true)
	}
	if db.PropUnderlay {
		dialogBase = dialogBase.Attr("underlay", true)
	}

	// Add event handlers
	if db.PropOnClose != nil {
		dialogBase = dialogBase.On("close", db.PropOnClose)
	}

	// Add child element if provided
	if db.PropChild != nil {
		dialogBase = dialogBase.Body(db.PropChild)
	}

	return dialogBase
}
