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

// SpectrumCoachIndicator represents an sp-coach-indicator component
type SpectrumCoachIndicator struct {
	app.Compo

	// Properties
	quiet       bool
	staticColor StaticColorType
}

// CoachIndicator creates a new coach indicator component
func CoachIndicator() *SpectrumCoachIndicator {
	return &SpectrumCoachIndicator{}
}

// Quiet sets whether the coach indicator has a quiet appearance
func (ci *SpectrumCoachIndicator) Quiet(quiet bool) *SpectrumCoachIndicator {
	ci.quiet = quiet
	return ci
}

// StaticColor sets the static color of the coach indicator
func (ci *SpectrumCoachIndicator) StaticColor(color StaticColorType) *SpectrumCoachIndicator {
	ci.staticColor = color
	return ci
}

// StaticColorDark sets the static color to dark
func (ci *SpectrumCoachIndicator) StaticColorDark() *SpectrumCoachIndicator {
	return ci.StaticColor(StaticColorDark)
}

// StaticColorLight sets the static color to light
func (ci *SpectrumCoachIndicator) StaticColorLight() *SpectrumCoachIndicator {
	return ci.StaticColor(StaticColorLight)
}

// StaticColorWhite sets the static color to white
func (ci *SpectrumCoachIndicator) StaticColorWhite() *SpectrumCoachIndicator {
	return ci.StaticColor(StaticColorWhite)
}

// StaticColorBlack sets the static color to black
func (ci *SpectrumCoachIndicator) StaticColorBlack() *SpectrumCoachIndicator {
	return ci.StaticColor(StaticColorBlack)
}

// Render renders the coach indicator component
func (ci *SpectrumCoachIndicator) Render() app.UI {
	indicator := app.Elem("sp-coach-indicator")

	// Set attributes
	if ci.quiet {
		indicator = indicator.Attr("quiet", true)
	}
	if ci.staticColor != "" {
		indicator = indicator.Attr("static-color", string(ci.staticColor))
	}

	return indicator
}
