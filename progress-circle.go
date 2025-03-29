package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ProgressCircleSize represents the visual size of a progress circle
type ProgressCircleSize string

// Progress circle sizes
const (
	ProgressCircleSizeS ProgressCircleSize = "s"
	ProgressCircleSizeM ProgressCircleSize = "m"
	ProgressCircleSizeL ProgressCircleSize = "l"
)

// ProgressCircleStaticColor represents the static color for a progress circle
type ProgressCircleStaticColor string

// Progress circle static colors
const (
	ProgressCircleStaticColorWhite ProgressCircleStaticColor = "white"
)

// spectrumProgressCircle represents an sp-progress-circle component
type spectrumProgressCircle struct {
	app.Compo

	// Properties
	PropSize          ProgressCircleSize
	PropLabel         string
	PropProgress      float64
	PropIndeterminate bool
	PropStaticColor   ProgressCircleStaticColor
}

// ProgressCircle creates a new progress circle component
func ProgressCircle() *spectrumProgressCircle {
	return &spectrumProgressCircle{
		PropSize: ProgressCircleSizeM, // Default size is medium
	}
}

// Size sets the visual size of the progress circle
func (p *spectrumProgressCircle) Size(size ProgressCircleSize) *spectrumProgressCircle {
	p.PropSize = size
	return p
}

// Label sets the accessibility label
func (p *spectrumProgressCircle) Label(label string) *spectrumProgressCircle {
	p.PropLabel = label
	return p
}

// Progress sets the progress value (0-100)
func (p *spectrumProgressCircle) Progress(progress float64) *spectrumProgressCircle {
	p.PropProgress = progress
	return p
}

// Indeterminate sets whether the progress circle shows indeterminate progress
func (p *spectrumProgressCircle) Indeterminate(indeterminate bool) *spectrumProgressCircle {
	p.PropIndeterminate = indeterminate
	return p
}

// StaticColor sets the static color for use over backgrounds
func (p *spectrumProgressCircle) StaticColor(color ProgressCircleStaticColor) *spectrumProgressCircle {
	p.PropStaticColor = color
	return p
}

// Render renders the progress circle component
func (p *spectrumProgressCircle) Render() app.UI {
	progressCircle := app.Elem("sp-progress-circle").
		Attr("size", string(p.PropSize)).
		Attr("progress", p.PropProgress).
		Attr("indeterminate", p.PropIndeterminate).
		Attr("static-color", string(p.PropStaticColor))

	// Add label for accessibility
	if p.PropLabel != "" {
		progressCircle = progressCircle.Attr("label", p.PropLabel)
	}

	return progressCircle
}
