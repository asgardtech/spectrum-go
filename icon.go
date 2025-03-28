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

// SpectrumIcon represents an icon component that renders an icon through
// various methods: icon sets via name, an image via src, or custom SVG via children.
type SpectrumIcon struct {
	app.Compo

	// Properties
	label string
	name  string
	size  IconSize
	src   string
	slot  string

	// Content
	children []app.UI
}

// Icon creates a new icon component.
func Icon() *SpectrumIcon {
	return &SpectrumIcon{}
}

// Label sets the accessible label for the icon.
// If provided, removes aria-hidden and sets aria-label.
func (i *SpectrumIcon) Label(label string) *SpectrumIcon {
	i.label = label
	return i
}

// Name sets the name of the icon from a registered icon set.
func (i *SpectrumIcon) Name(name string) *SpectrumIcon {
	i.name = name
	return i
}

// Size sets the size of the icon.
func (i *SpectrumIcon) Size(size IconSize) *SpectrumIcon {
	i.size = size
	return i
}

// Src sets the URL source for an image-based icon.
func (i *SpectrumIcon) Src(src string) *SpectrumIcon {
	i.src = src
	return i
}

// Slot sets the slot attribute on the icon.
// This allows the icon to be placed in named slots of other components.
func (i *SpectrumIcon) Slot(slot string) *SpectrumIcon {
	i.slot = slot
	return i
}

// Children sets the content of the icon, typically an SVG element.
func (i *SpectrumIcon) Children(children ...app.UI) *SpectrumIcon {
	i.children = children
	return i
}

// Render renders the icon component.
func (i *SpectrumIcon) Render() app.UI {
	icon := app.Elem("sp-icon")

	// Set attributes based on properties
	if i.label != "" {
		icon.Attr("label", i.label)
	}

	if i.name != "" {
		icon.Attr("name", i.name)
	}

	if i.size != "" {
		icon.Attr("size", string(i.size))
	}

	if i.src != "" {
		icon.Attr("src", i.src)
	}

	if i.slot != "" {
		icon.Attr("slot", i.slot)
	}

	// Add children content (typically SVG)
	for _, child := range i.children {
		icon.Body(child)
	}

	return icon
}
