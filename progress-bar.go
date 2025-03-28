package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ProgressBarSize represents the visual size of a progress bar
type ProgressBarSize string

// Progress bar sizes
const (
	ProgressBarSizeS  ProgressBarSize = "s"
	ProgressBarSizeM  ProgressBarSize = "m"
	ProgressBarSizeL  ProgressBarSize = "l"
	ProgressBarSizeXL ProgressBarSize = "xl"
)

// ProgressBarStaticColor represents the static color for a progress bar
type ProgressBarStaticColor string

// Progress bar static colors
const (
	ProgressBarStaticColorWhite ProgressBarStaticColor = "white"
)

// SpectrumProgressBar represents an sp-progress-bar component
type SpectrumProgressBar struct {
	app.Compo

	// Properties
	size           ProgressBarSize
	label          string
	ariaLabel      string // Used when no visible label is desired
	progress       float64
	indeterminate  bool
	overBackground bool
	sideLabel      bool
	staticColor    ProgressBarStaticColor
}

// ProgressBar creates a new progress bar component
func ProgressBar() *SpectrumProgressBar {
	return &SpectrumProgressBar{
		size: ProgressBarSizeM, // Default size is medium
	}
}

// Size sets the visual size of the progress bar
func (p *SpectrumProgressBar) Size(size ProgressBarSize) *SpectrumProgressBar {
	p.size = size
	return p
}

// Label sets the text label content
// This label will be visible and also used for accessibility
func (p *SpectrumProgressBar) Label(label string) *SpectrumProgressBar {
	p.label = label
	return p
}

// AriaLabel sets the aria-label attribute
// Use this when you don't want a visible label but still need accessibility
func (p *SpectrumProgressBar) AriaLabel(label string) *SpectrumProgressBar {
	p.ariaLabel = label
	return p
}

// Progress sets the progress value (0-100)
func (p *SpectrumProgressBar) Progress(progress float64) *SpectrumProgressBar {
	p.progress = progress
	return p
}

// Indeterminate sets whether the progress bar shows indeterminate progress
func (p *SpectrumProgressBar) Indeterminate(indeterminate bool) *SpectrumProgressBar {
	p.indeterminate = indeterminate
	return p
}

// OverBackground sets whether the progress bar should be displayed over a background
func (p *SpectrumProgressBar) OverBackground(overBackground bool) *SpectrumProgressBar {
	p.overBackground = overBackground
	return p
}

// SideLabel sets whether the label should be displayed beside the progress bar
func (p *SpectrumProgressBar) SideLabel(sideLabel bool) *SpectrumProgressBar {
	p.sideLabel = sideLabel
	return p
}

// StaticColor sets the static color
func (p *SpectrumProgressBar) StaticColor(color ProgressBarStaticColor) *SpectrumProgressBar {
	p.staticColor = color
	return p
}

// Render renders the progress bar component
func (p *SpectrumProgressBar) Render() app.UI {
	progressBar := app.Elem("sp-progress-bar").
		Attr("size", string(p.size)).
		Attr("progress", p.progress).
		Attr("indeterminate", p.indeterminate).
		Attr("over-background", p.overBackground).
		Attr("side-label", p.sideLabel).
		Attr("static-color", string(p.staticColor))

	// Add label or aria-label
	if p.label != "" {
		progressBar = progressBar.Attr("label", p.label)
	}

	if p.ariaLabel != "" {
		progressBar = progressBar.Attr("aria-label", p.ariaLabel)
	}

	return progressBar
}
