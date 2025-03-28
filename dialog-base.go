package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumDialogBase represents an sp-dialog-base component
type SpectrumDialogBase struct {
	app.Compo

	// Properties
	open        bool
	dismissable bool
	mode        DialogMode
	responsive  bool
	underlay    bool

	// Child content (usually an sp-dialog)
	child app.UI

	// Event handlers
	onClose app.EventHandler
}

// DialogBase creates a new dialog base component
func DialogBase() *SpectrumDialogBase {
	return &SpectrumDialogBase{}
}

// Open sets whether the dialog base is open
func (db *SpectrumDialogBase) Open(open bool) *SpectrumDialogBase {
	db.open = open
	return db
}

// Dismissable sets whether the dialog is dismissable
func (db *SpectrumDialogBase) Dismissable(dismissable bool) *SpectrumDialogBase {
	db.dismissable = dismissable
	return db
}

// Mode sets the display mode of the dialog
func (db *SpectrumDialogBase) Mode(mode DialogMode) *SpectrumDialogBase {
	db.mode = mode
	return db
}

// Responsive sets whether the dialog should be responsive
func (db *SpectrumDialogBase) Responsive(responsive bool) *SpectrumDialogBase {
	db.responsive = responsive
	return db
}

// Underlay sets whether the dialog should have an underlay
func (db *SpectrumDialogBase) Underlay(underlay bool) *SpectrumDialogBase {
	db.underlay = underlay
	return db
}

// Child sets the child element (usually an sp-dialog)
func (db *SpectrumDialogBase) Child(child app.UI) *SpectrumDialogBase {
	db.child = child
	return db
}

// OnClose sets the close event handler
func (db *SpectrumDialogBase) OnClose(handler app.EventHandler) *SpectrumDialogBase {
	db.onClose = handler
	return db
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (db *SpectrumDialogBase) Fullscreen() *SpectrumDialogBase {
	return db.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (db *SpectrumDialogBase) FullscreenTakeover() *SpectrumDialogBase {
	return db.Mode(DialogModeFullscreenTakeover)
}

// Render renders the dialog base component
func (db *SpectrumDialogBase) Render() app.UI {
	dialogBase := app.Elem("sp-dialog-base")

	// Add properties
	if db.open {
		dialogBase = dialogBase.Attr("open", true)
	}
	if db.dismissable {
		dialogBase = dialogBase.Attr("dismissable", true)
	}
	if db.mode != "" {
		dialogBase = dialogBase.Attr("mode", string(db.mode))
	}
	if db.responsive {
		dialogBase = dialogBase.Attr("responsive", true)
	}
	if db.underlay {
		dialogBase = dialogBase.Attr("underlay", true)
	}

	// Add event handlers
	if db.onClose != nil {
		dialogBase = dialogBase.On("close", db.onClose)
	}

	// Add child element if provided
	if db.child != nil {
		dialogBase = dialogBase.Body(db.child)
	}

	return dialogBase
}
