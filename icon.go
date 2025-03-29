package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconSize represents the size variants available for the Icon.
type IconSize string

// Icon size constants.
const (
	IconXXS IconSize = "xxs"
	IconXS  IconSize = "xs"
	IconS   IconSize = "s"
	IconM   IconSize = "m"
	IconL   IconSize = "l"
	IconXL  IconSize = "xl"
	IconXXL IconSize = "xxl"
)

// spectrumIcon represents an icon component that renders an icon through
// various methods: icon sets via name, an image via src, or custom SVG via children.
type spectrumIcon struct {
	app.Compo

	// Properties
	PropLabel string
	PropName  string
	PropSize  IconSize
	PropSrc   string
	PropSlot  string

	// Content
	PropChildren []app.UI
}

// Icon creates a new icon component.
func Icon() *spectrumIcon {
	return &spectrumIcon{}
}

// Label sets the accessible label for the icon.
// If provided, removes aria-hidden and sets aria-label.
func (i *spectrumIcon) Label(label string) *spectrumIcon {
	i.PropLabel = label
	return i
}

// Name sets the name of the icon from a registered icon set.
func (i *spectrumIcon) Name(name string) *spectrumIcon {
	i.PropName = name
	return i
}

// Size sets the size of the icon.
func (i *spectrumIcon) Size(size IconSize) *spectrumIcon {
	i.PropSize = size
	return i
}

// Src sets the URL source for an image-based icon.
func (i *spectrumIcon) Src(src string) *spectrumIcon {
	i.PropSrc = src
	return i
}

// Slot sets the slot attribute on the icon.
// This allows the icon to be placed in named slots of other components.
func (i *spectrumIcon) Slot(slot string) *spectrumIcon {
	i.PropSlot = slot
	return i
}

// Children sets the content of the icon, typically an SVG element.
func (i *spectrumIcon) Children(children ...app.UI) *spectrumIcon {
	i.PropChildren = children
	return i
}

// Render renders the icon component.
func (i *spectrumIcon) Render() app.UI {
	icon := app.Elem("sp-icon")

	// Set attributes based on properties
	if i.PropLabel != "" {
		icon.Attr("label", i.PropLabel)
	}

	if i.PropName != "" {
		icon.Attr("name", i.PropName)
	}

	if i.PropSize != "" {
		icon.Attr("size", string(i.PropSize))
	}

	if i.PropSrc != "" {
		icon.Attr("src", i.PropSrc)
	}

	if i.PropSlot != "" {
		icon.Attr("slot", i.PropSlot)
	}

	// Add children content (typically SVG)
	for _, child := range i.PropChildren {
		icon.Body(child)
	}

	return icon
}
