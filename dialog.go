package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DialogSize represents the visual size of a dialog
type DialogSize string

// Dialog sizes
const (
	DialogSizeS DialogSize = "s"
	DialogSizeM DialogSize = "m"
	DialogSizeL DialogSize = "l"
)

// DialogMode represents the display mode of a dialog
type DialogMode string

// Dialog modes
const (
	DialogModeFullscreen         DialogMode = "fullscreen"
	DialogModeFullscreenTakeover DialogMode = "fullscreenTakeover"
)

// spectrumDialog represents an sp-dialog component
type spectrumDialog struct {
	app.Compo

	// Properties
	PropDismissLabel string
	PropDismissable  bool
	PropError        bool
	PropMode         DialogMode
	PropNoDivider    bool
	PropSize         DialogSize
	PropVariant      string

	// Slots
	PropButtons []app.UI
	PropFooter  app.UI
	PropHeading app.UI
	PropHero    app.UI
	PropContent []app.UI

	// Event handlers
	PropOnClose app.EventHandler
}

// Dialog creates a new dialog component
func Dialog() *spectrumDialog {
	return &spectrumDialog{
		PropDismissLabel: "Close", // Default dismiss label
	}
}

// DismissLabel sets the label for the dismiss button
func (d *spectrumDialog) DismissLabel(label string) *spectrumDialog {
	d.PropDismissLabel = label
	return d
}

// Dismissable sets whether the dialog has a dismiss button
func (d *spectrumDialog) Dismissable(dismissable bool) *spectrumDialog {
	d.PropDismissable = dismissable
	return d
}

// Error sets whether the dialog displays an error state
func (d *spectrumDialog) Error(error bool) *spectrumDialog {
	d.PropError = error
	return d
}

// Mode sets the display mode of the dialog
func (d *spectrumDialog) Mode(mode DialogMode) *spectrumDialog {
	d.PropMode = mode
	return d
}

// NoDivider sets whether the dialog should have a divider between header and content
func (d *spectrumDialog) NoDivider(noDivider bool) *spectrumDialog {
	d.PropNoDivider = noDivider
	return d
}

// Size sets the visual size of the dialog
func (d *spectrumDialog) Size(size DialogSize) *spectrumDialog {
	d.PropSize = size
	return d
}

// Variant sets the variant style of the dialog
func (d *spectrumDialog) Variant(variant string) *spectrumDialog {
	d.PropVariant = variant
	return d
}

// Heading sets the heading element for the dialog
func (d *spectrumDialog) Heading(heading app.UI) *spectrumDialog {
	d.PropHeading = heading
	return d
}

// HeadingText sets a text heading for the dialog
func (d *spectrumDialog) HeadingText(text string) *spectrumDialog {
	d.PropHeading = app.H2().Text(text)
	return d
}

// Hero sets the hero image element for the dialog
func (d *spectrumDialog) Hero(hero app.UI) *spectrumDialog {
	d.PropHero = hero
	return d
}

// Button adds a button to the dialog's button area
func (d *spectrumDialog) Button(button app.UI) *spectrumDialog {
	d.PropButtons = append(d.PropButtons, button)
	return d
}

// Buttons sets all buttons for the dialog
func (d *spectrumDialog) Buttons(buttons ...app.UI) *spectrumDialog {
	d.PropButtons = buttons
	return d
}

// Footer sets the footer content for the dialog
func (d *spectrumDialog) Footer(footer app.UI) *spectrumDialog {
	d.PropFooter = footer
	return d
}

// Content adds content to the dialog's main content area
func (d *spectrumDialog) Content(content app.UI) *spectrumDialog {
	d.PropContent = append(d.PropContent, content)
	return d
}

// ContentList sets all content elements for the dialog
func (d *spectrumDialog) ContentList(content ...app.UI) *spectrumDialog {
	d.PropContent = content
	return d
}

// ContentText adds text content to the dialog
func (d *spectrumDialog) ContentText(text string) *spectrumDialog {
	d.PropContent = append(d.PropContent, app.P().Text(text))
	return d
}

// OnClose sets the close event handler
func (d *spectrumDialog) OnClose(handler app.EventHandler) *spectrumDialog {
	d.PropOnClose = handler
	return d
}

// Convenience methods for setting dialog sizes

// Small sets size to small (s)
func (d *spectrumDialog) Small() *spectrumDialog {
	return d.Size(DialogSizeS)
}

// Medium sets size to medium (m)
func (d *spectrumDialog) Medium() *spectrumDialog {
	return d.Size(DialogSizeM)
}

// Large sets size to large (l)
func (d *spectrumDialog) Large() *spectrumDialog {
	return d.Size(DialogSizeL)
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (d *spectrumDialog) Fullscreen() *spectrumDialog {
	return d.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (d *spectrumDialog) FullscreenTakeover() *spectrumDialog {
	return d.Mode(DialogModeFullscreenTakeover)
}

// Render renders the dialog component
func (d *spectrumDialog) Render() app.UI {
	dialog := app.Elem("sp-dialog")

	// Add properties
	if d.PropDismissLabel != "" {
		dialog = dialog.Attr("dismiss-label", d.PropDismissLabel)
	}
	if d.PropDismissable {
		dialog = dialog.Attr("dismissable", true)
	}
	if d.PropError {
		dialog = dialog.Attr("error", true)
	}
	if d.PropMode != "" {
		dialog = dialog.Attr("mode", string(d.PropMode))
	}
	if d.PropNoDivider {
		dialog = dialog.Attr("no-divider", true)
	}
	if d.PropSize != "" {
		dialog = dialog.Attr("size", string(d.PropSize))
	}
	if d.PropVariant != "" {
		dialog = dialog.Attr("variant", d.PropVariant)
	}

	// Add event handlers
	if d.PropOnClose != nil {
		dialog = dialog.On("close", d.PropOnClose)
	}

	// Collect slot elements
	elements := []app.UI{}

	// Add heading slot if provided
	if d.PropHeading != nil {
		headingElem := d.PropHeading
		if headingWithSlot, ok := headingElem.(interface{ Slot(string) app.UI }); ok {
			headingElem = headingWithSlot.Slot("heading")
		} else {
			headingElem = app.Elem("div").
				Attr("slot", "heading").
				Body(headingElem)
		}
		elements = append(elements, headingElem)
	}

	// Add hero slot if provided
	if d.PropHero != nil {
		heroElem := d.PropHero
		if heroWithSlot, ok := heroElem.(interface{ Slot(string) app.UI }); ok {
			heroElem = heroWithSlot.Slot("hero")
		} else {
			heroElem = app.Elem("div").
				Attr("slot", "hero").
				Body(heroElem)
		}
		elements = append(elements, heroElem)
	}

	// Add footer slot if provided
	if d.PropFooter != nil {
		footerElem := d.PropFooter
		if footerWithSlot, ok := footerElem.(interface{ Slot(string) app.UI }); ok {
			footerElem = footerWithSlot.Slot("footer")
		} else {
			footerElem = app.Elem("div").
				Attr("slot", "footer").
				Body(footerElem)
		}
		elements = append(elements, footerElem)
	}

	// Add button slots if provided
	for _, button := range d.PropButtons {
		buttonElem := button
		if buttonWithSlot, ok := buttonElem.(interface{ Slot(string) app.UI }); ok {
			buttonElem = buttonWithSlot.Slot("button")
		} else {
			buttonElem = app.Elem("div").
				Attr("slot", "button").
				Body(buttonElem)
		}
		elements = append(elements, buttonElem)
	}

	// Add content (default slot) if provided
	elements = append(elements, d.PropContent...)

	// Add all elements to the dialog
	dialog = dialog.Body(elements...)

	return dialog
}
