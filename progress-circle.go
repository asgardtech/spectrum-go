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

// SpectrumProgressCircle represents an sp-progress-circle component
type SpectrumProgressCircle struct {
	app.Compo

	// Properties
	size          ProgressCircleSize
	label         string
	progress      float64
	indeterminate bool
	staticColor   ProgressCircleStaticColor
}

// ProgressCircle creates a new progress circle component
func ProgressCircle() *SpectrumProgressCircle {
	return &SpectrumProgressCircle{
		size: ProgressCircleSizeM, // Default size is medium
	}
}

// Size sets the visual size of the progress circle
func (p *SpectrumProgressCircle) Size(size ProgressCircleSize) *SpectrumProgressCircle {
	p.size = size
	return p
}

// Label sets the accessibility label
func (p *SpectrumProgressCircle) Label(label string) *SpectrumProgressCircle {
	p.label = label
	return p
}

// Progress sets the progress value (0-100)
func (p *SpectrumProgressCircle) Progress(progress float64) *SpectrumProgressCircle {
	p.progress = progress
	return p
}

// Indeterminate sets whether the progress circle shows indeterminate progress
func (p *SpectrumProgressCircle) Indeterminate(indeterminate bool) *SpectrumProgressCircle {
	p.indeterminate = indeterminate
	return p
}

// StaticColor sets the static color for use over backgrounds
func (p *SpectrumProgressCircle) StaticColor(color ProgressCircleStaticColor) *SpectrumProgressCircle {
	p.staticColor = color
	return p
}

// Render renders the progress circle component
func (p *SpectrumProgressCircle) Render() app.UI {
	progressCircle := app.Elem("sp-progress-circle").
		Attr("size", string(p.size)).
		Attr("progress", p.progress).
		Attr("indeterminate", p.indeterminate).
		Attr("static-color", string(p.staticColor))

	// Add label for accessibility
	if p.label != "" {
		progressCircle = progressCircle.Attr("label", p.label)
	}

	return progressCircle
}
