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

// SpectrumDialog represents an sp-dialog component
type SpectrumDialog struct {
	app.Compo

	// Properties
	dismissLabel string
	dismissable  bool
	error        bool
	mode         DialogMode
	noDivider    bool
	size         DialogSize
	variant      string

	// Slots
	buttons []app.UI
	footer  app.UI
	heading app.UI
	hero    app.UI
	content []app.UI

	// Event handlers
	onClose app.EventHandler
}

// Dialog creates a new dialog component
func Dialog() *SpectrumDialog {
	return &SpectrumDialog{
		dismissLabel: "Close", // Default dismiss label
	}
}

// DismissLabel sets the label for the dismiss button
func (d *SpectrumDialog) DismissLabel(label string) *SpectrumDialog {
	d.dismissLabel = label
	return d
}

// Dismissable sets whether the dialog has a dismiss button
func (d *SpectrumDialog) Dismissable(dismissable bool) *SpectrumDialog {
	d.dismissable = dismissable
	return d
}

// Error sets whether the dialog displays an error state
func (d *SpectrumDialog) Error(error bool) *SpectrumDialog {
	d.error = error
	return d
}

// Mode sets the display mode of the dialog
func (d *SpectrumDialog) Mode(mode DialogMode) *SpectrumDialog {
	d.mode = mode
	return d
}

// NoDivider sets whether the dialog should have a divider between header and content
func (d *SpectrumDialog) NoDivider(noDivider bool) *SpectrumDialog {
	d.noDivider = noDivider
	return d
}

// Size sets the visual size of the dialog
func (d *SpectrumDialog) Size(size DialogSize) *SpectrumDialog {
	d.size = size
	return d
}

// Variant sets the variant style of the dialog
func (d *SpectrumDialog) Variant(variant string) *SpectrumDialog {
	d.variant = variant
	return d
}

// Heading sets the heading element for the dialog
func (d *SpectrumDialog) Heading(heading app.UI) *SpectrumDialog {
	d.heading = heading
	return d
}

// HeadingText sets a text heading for the dialog
func (d *SpectrumDialog) HeadingText(text string) *SpectrumDialog {
	d.heading = app.H2().Text(text)
	return d
}

// Hero sets the hero image element for the dialog
func (d *SpectrumDialog) Hero(hero app.UI) *SpectrumDialog {
	d.hero = hero
	return d
}

// Button adds a button to the dialog's button area
func (d *SpectrumDialog) Button(button app.UI) *SpectrumDialog {
	d.buttons = append(d.buttons, button)
	return d
}

// Buttons sets all buttons for the dialog
func (d *SpectrumDialog) Buttons(buttons ...app.UI) *SpectrumDialog {
	d.buttons = buttons
	return d
}

// Footer sets the footer content for the dialog
func (d *SpectrumDialog) Footer(footer app.UI) *SpectrumDialog {
	d.footer = footer
	return d
}

// Content adds content to the dialog's main content area
func (d *SpectrumDialog) Content(content app.UI) *SpectrumDialog {
	d.content = append(d.content, content)
	return d
}

// ContentList sets all content elements for the dialog
func (d *SpectrumDialog) ContentList(content ...app.UI) *SpectrumDialog {
	d.content = content
	return d
}

// ContentText adds text content to the dialog
func (d *SpectrumDialog) ContentText(text string) *SpectrumDialog {
	d.content = append(d.content, app.P().Text(text))
	return d
}

// OnClose sets the close event handler
func (d *SpectrumDialog) OnClose(handler app.EventHandler) *SpectrumDialog {
	d.onClose = handler
	return d
}

// Convenience methods for setting dialog sizes

// Small sets size to small (s)
func (d *SpectrumDialog) Small() *SpectrumDialog {
	return d.Size(DialogSizeS)
}

// Medium sets size to medium (m)
func (d *SpectrumDialog) Medium() *SpectrumDialog {
	return d.Size(DialogSizeM)
}

// Large sets size to large (l)
func (d *SpectrumDialog) Large() *SpectrumDialog {
	return d.Size(DialogSizeL)
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (d *SpectrumDialog) Fullscreen() *SpectrumDialog {
	return d.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (d *SpectrumDialog) FullscreenTakeover() *SpectrumDialog {
	return d.Mode(DialogModeFullscreenTakeover)
}

// Render renders the dialog component
func (d *SpectrumDialog) Render() app.UI {
	dialog := app.Elem("sp-dialog")

	// Add properties
	if d.dismissLabel != "" {
		dialog = dialog.Attr("dismiss-label", d.dismissLabel)
	}
	if d.dismissable {
		dialog = dialog.Attr("dismissable", true)
	}
	if d.error {
		dialog = dialog.Attr("error", true)
	}
	if d.mode != "" {
		dialog = dialog.Attr("mode", string(d.mode))
	}
	if d.noDivider {
		dialog = dialog.Attr("no-divider", true)
	}
	if d.size != "" {
		dialog = dialog.Attr("size", string(d.size))
	}
	if d.variant != "" {
		dialog = dialog.Attr("variant", d.variant)
	}

	// Add event handlers
	if d.onClose != nil {
		dialog = dialog.On("close", d.onClose)
	}

	// Collect slot elements
	elements := []app.UI{}

	// Add heading slot if provided
	if d.heading != nil {
		headingElem := d.heading
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
	if d.hero != nil {
		heroElem := d.hero
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
	if d.footer != nil {
		footerElem := d.footer
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
	for _, button := range d.buttons {
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
	elements = append(elements, d.content...)

	// Add all elements to the dialog
	dialog = dialog.Body(elements...)

	return dialog
}
