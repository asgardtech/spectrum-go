package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ThemeColor represents the color theme variant
type ThemeColor string

// Theme colors
const (
	ThemeColorLightest ThemeColor = "lightest"
	ThemeColorLight    ThemeColor = "light"
	ThemeColorDark     ThemeColor = "dark"
	ThemeColorDarkest  ThemeColor = "darkest"
)

// ThemeScale represents the size scale variant
type ThemeScale string

// Theme scales
const (
	ThemeScaleMedium ThemeScale = "medium"
	ThemeScaleLarge  ThemeScale = "large"
)

// ThemeSystem represents the design system variant
type ThemeSystem string

// Theme systems
const (
	ThemeSystemSpectrum ThemeSystem = "spectrum"
	ThemeSystemExpress  ThemeSystem = "express"
)

// spectrumTheme represents an sp-theme component
type spectrumTheme struct {
	app.Compo

	// Properties
	PropColor    ThemeColor
	PropScale    ThemeScale
	PropSystem   ThemeSystem
	PropDir      string
	PropLang     string
	PropOnUpdate app.EventHandler

	// Children
	children []app.UI
}

// Theme creates a new theme component
func Theme() *spectrumTheme {
	return &spectrumTheme{
		PropColor:  ThemeColorLight,     // Default is light
		PropScale:  ThemeScaleMedium,    // Default is medium
		PropSystem: ThemeSystemSpectrum, // Default is spectrum
		PropDir:    "ltr",               // Default is left-to-right
	}
}

// Color sets the theme color
func (t *spectrumTheme) Color(color ThemeColor) *spectrumTheme {
	t.PropColor = color
	return t
}

// Scale sets the theme scale
func (t *spectrumTheme) Scale(scale ThemeScale) *spectrumTheme {
	t.PropScale = scale
	return t
}

// System sets the theme system
func (t *spectrumTheme) System(system ThemeSystem) *spectrumTheme {
	t.PropSystem = system
	return t
}

// Dir sets the text direction
func (t *spectrumTheme) Dir(dir string) *spectrumTheme {
	t.PropDir = dir
	return t
}

// Lang sets the language
func (t *spectrumTheme) Lang(lang string) *spectrumTheme {
	t.PropLang = lang
	return t
}

// OnUpdate sets the update event handler
func (t *spectrumTheme) OnUpdate(handler app.EventHandler) *spectrumTheme {
	t.PropOnUpdate = handler
	return t
}

// Child adds a child element to the theme
func (t *spectrumTheme) Child(child app.UI) *spectrumTheme {
	t.children = append(t.children, child)
	return t
}

// Children adds multiple child elements to the theme
func (t *spectrumTheme) Children(children ...app.UI) *spectrumTheme {
	t.children = append(t.children, children...)
	return t
}

// Render renders the theme component
func (t *spectrumTheme) Render() app.UI {
	// Create the base theme element
	theme := app.Elem("sp-theme").
		Attr("color", string(t.PropColor)).
		Attr("scale", string(t.PropScale)).
		Attr("system", string(t.PropSystem)).
		Attr("dir", t.PropDir).
		Attr("lang", t.PropLang).
		Body(t.children...).
		On("update", t.PropOnUpdate)

	return theme
}
