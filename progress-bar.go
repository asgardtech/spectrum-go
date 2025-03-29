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

// spectrumProgressBar represents an sp-progress-bar component
type spectrumProgressBar struct {
	app.Compo

	// Properties
	PropSize           ProgressBarSize
	PropLabel          string
	PropAriaLabel      string // Used when no visible label is desired
	PropProgress       float64
	PropIndeterminate  bool
	PropOverBackground bool
	PropSideLabel      bool
	PropStaticColor    ProgressBarStaticColor
}

// ProgressBar creates a new progress bar component
func ProgressBar() *spectrumProgressBar {
	return &spectrumProgressBar{
		PropSize: ProgressBarSizeM, // Default size is medium
	}
}

// Size sets the visual size of the progress bar
func (p *spectrumProgressBar) Size(size ProgressBarSize) *spectrumProgressBar {
	p.PropSize = size
	return p
}

// Label sets the text label content
// This label will be visible and also used for accessibility
func (p *spectrumProgressBar) Label(label string) *spectrumProgressBar {
	p.PropLabel = label
	return p
}

// AriaLabel sets the aria-label attribute
// Use this when you don't want a visible label but still need accessibility
func (p *spectrumProgressBar) AriaLabel(label string) *spectrumProgressBar {
	p.PropAriaLabel = label
	return p
}

// Progress sets the progress value (0-100)
func (p *spectrumProgressBar) Progress(progress float64) *spectrumProgressBar {
	p.PropProgress = progress
	return p
}

// Indeterminate sets whether the progress bar shows indeterminate progress
func (p *spectrumProgressBar) Indeterminate(indeterminate bool) *spectrumProgressBar {
	p.PropIndeterminate = indeterminate
	return p
}

// OverBackground sets whether the progress bar should be displayed over a background
func (p *spectrumProgressBar) OverBackground(overBackground bool) *spectrumProgressBar {
	p.PropOverBackground = overBackground
	return p
}

// SideLabel sets whether the label should be displayed beside the progress bar
func (p *spectrumProgressBar) SideLabel(sideLabel bool) *spectrumProgressBar {
	p.PropSideLabel = sideLabel
	return p
}

// StaticColor sets the static color
func (p *spectrumProgressBar) StaticColor(color ProgressBarStaticColor) *spectrumProgressBar {
	p.PropStaticColor = color
	return p
}

// Render renders the progress bar component
func (p *spectrumProgressBar) Render() app.UI {
	progressBar := app.Elem("sp-progress-bar").
		Attr("size", string(p.PropSize)).
		Attr("progress", p.PropProgress).
		Attr("indeterminate", p.PropIndeterminate).
		Attr("over-background", p.PropOverBackground).
		Attr("side-label", p.PropSideLabel).
		Attr("static-color", string(p.PropStaticColor))

	// Add label or aria-label
	if p.PropLabel != "" {
		progressBar = progressBar.Attr("label", p.PropLabel)
	}

	if p.PropAriaLabel != "" {
		progressBar = progressBar.Attr("aria-label", p.PropAriaLabel)
	}

	return progressBar
}
