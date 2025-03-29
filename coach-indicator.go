package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// StaticColorType represents the static color of a coach indicator
type StaticColorType string

const (
	StaticColorWhite StaticColorType = "white"
	StaticColorBlack StaticColorType = "black"
	StaticColorDark  StaticColorType = "dark"
	StaticColorLight StaticColorType = "light"
)

// spectrumCoachIndicator represents an sp-coach-indicator component
type spectrumCoachIndicator struct {
	app.Compo

	// Properties
	PropQuiet       bool
	PropStaticColor StaticColorType
}

// CoachIndicator creates a new coach indicator component
func CoachIndicator() *spectrumCoachIndicator {
	return &spectrumCoachIndicator{}
}

// Quiet sets whether the coach indicator has a quiet appearance
func (ci *spectrumCoachIndicator) Quiet(quiet bool) *spectrumCoachIndicator {
	ci.PropQuiet = quiet
	return ci
}

// StaticColor sets the static color of the coach indicator
func (ci *spectrumCoachIndicator) StaticColor(color StaticColorType) *spectrumCoachIndicator {
	ci.PropStaticColor = color
	return ci
}

// StaticColorDark sets the static color to dark
func (ci *spectrumCoachIndicator) StaticColorDark() *spectrumCoachIndicator {
	return ci.StaticColor(StaticColorDark)
}

// StaticColorLight sets the static color to light
func (ci *spectrumCoachIndicator) StaticColorLight() *spectrumCoachIndicator {
	return ci.StaticColor(StaticColorLight)
}

// StaticColorWhite sets the static color to white
func (ci *spectrumCoachIndicator) StaticColorWhite() *spectrumCoachIndicator {
	return ci.StaticColor(StaticColorWhite)
}

// StaticColorBlack sets the static color to black
func (ci *spectrumCoachIndicator) StaticColorBlack() *spectrumCoachIndicator {
	return ci.StaticColor(StaticColorBlack)
}

// Render renders the coach indicator component
func (ci *spectrumCoachIndicator) Render() app.UI {
	indicator := app.Elem("sp-coach-indicator")

	// Set attributes
	if ci.PropQuiet {
		indicator = indicator.Attr("quiet", true)
	}
	if ci.PropStaticColor != "" {
		indicator = indicator.Attr("static-color", string(ci.PropStaticColor))
	}

	return indicator
}
