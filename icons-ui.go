package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumIconsUI represents a component that provides access to the Spectrum UI icons.
// This is a wrapper around the Adobe Spectrum Web Components icons-ui package.
type SpectrumIconsUI struct {
	app.Compo

	// Properties
	label string
	size  IconSize
}

// IconsUI creates a new icons-ui component.
func IconsUI() *SpectrumIconsUI {
	return &SpectrumIconsUI{}
}

// Label sets the accessible label for the icon.
func (i *SpectrumIconsUI) Label(label string) *SpectrumIconsUI {
	i.label = label
	return i
}

// Size sets the size of the icons.
func (i *SpectrumIconsUI) Size(size IconSize) *SpectrumIconsUI {
	i.size = size
	return i
}

// Render renders the icons-ui component.
func (i *SpectrumIconsUI) Render() app.UI {
	iconsUI := app.Elem("sp-icons-ui")

	// Set attributes based on properties
	if i.label != "" {
		iconsUI.Attr("label", i.label)
	}

	if i.size != "" {
		iconsUI.Attr("size", string(i.size))
	}

	return iconsUI
}

// Icon utility functions for specific UI icons
// These are convenience methods to create specific icon elements

// ArrowIcon creates an Arrow icon
func ArrowIcon() app.UI {
	return app.Elem("sp-icon-arrow100")
}

// CheckmarkIcon creates a Checkmark icon
func CheckmarkIcon() app.UI {
	return app.Elem("sp-icon-checkmark100")
}

// ChevronIcon creates a Chevron icon
func ChevronIcon() app.UI {
	return app.Elem("sp-icon-chevron100")
}

// CrossIcon creates a Cross/X icon
func CrossIcon() app.UI {
	return app.Elem("sp-icon-cross100")
}

// DashIcon creates a Dash/Minus icon
func DashIcon() app.UI {
	return app.Elem("sp-icon-dash100")
}

// InfoIcon creates an Info icon
func InfoIcon() app.UI {
	return app.Elem("sp-icon-info")
}

// AlertIcon creates an Alert icon
func AlertIcon() app.UI {
	return app.Elem("sp-icon-alert")
}

// CornerTriangleIcon creates a Corner Triangle icon
func CornerTriangleIcon() app.UI {
	return app.Elem("sp-icon-corner-triangle100")
}

// The above are just examples. In a real implementation, all UI icons would be added here.
