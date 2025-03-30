package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ProgressBarSize represents the The size of the progress bar
type ProgressBarSize string

// ProgressBarSize values
const (
    ProgressBarSizeS ProgressBarSize = "s"
    ProgressBarSizeM ProgressBarSize = "m"
    ProgressBarSizeL ProgressBarSize = "l"
    ProgressBarSizeXl ProgressBarSize = "xl"
)
// ProgressBarStaticColor represents the A static color variant for the progress bar
type ProgressBarStaticColor string

// ProgressBarStaticColor values
const (
    ProgressBarStaticColorWhite ProgressBarStaticColor = "white"
)

// spectrumProgressBar represents an sp-progress-bar component
type spectrumProgressBar struct {
    app.Compo
    *styler[*spectrumProgressBar]

    // Properties
    // Whether progress is indeterminate (unknown duration)
    PropIndeterminate bool
    // Text label to display with the progress bar
    PropLabel string
    // Whether the progress bar is displayed over a colored background
    PropOverBackground bool
    // The progress value from 0-100
    PropProgress float64
    // Whether to display the label beside the progress bar instead of above it
    PropSideLabel bool
    // The size of the progress bar
    PropSize ProgressBarSize
    // A static color variant for the progress bar
    PropStaticColor ProgressBarStaticColor




}

// ProgressBar creates a new sp-progress-bar component
func ProgressBar() *spectrumProgressBar {
    element := &spectrumProgressBar{
        PropIndeterminate: false,
        PropLabel: "",
        PropOverBackground: false,
        PropProgress: 0,
        PropSideLabel: false,
        PropSize: ProgressBarSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Whether progress is indeterminate (unknown duration)
func (c *spectrumProgressBar) Indeterminate(indeterminate bool) *spectrumProgressBar {
    c.PropIndeterminate = indeterminate
    return c
}

func (c *spectrumProgressBar) SetIndeterminate() *spectrumProgressBar {
    return c.Indeterminate(true)
}

// Text label to display with the progress bar
func (c *spectrumProgressBar) Label(label string) *spectrumProgressBar {
    c.PropLabel = label
    return c
}

// Whether the progress bar is displayed over a colored background
func (c *spectrumProgressBar) OverBackground(overBackground bool) *spectrumProgressBar {
    c.PropOverBackground = overBackground
    return c
}

func (c *spectrumProgressBar) SetOverBackground() *spectrumProgressBar {
    return c.OverBackground(true)
}

// The progress value from 0-100
func (c *spectrumProgressBar) Progress(progress float64) *spectrumProgressBar {
    c.PropProgress = progress
    return c
}

// Whether to display the label beside the progress bar instead of above it
func (c *spectrumProgressBar) SideLabel(sideLabel bool) *spectrumProgressBar {
    c.PropSideLabel = sideLabel
    return c
}

func (c *spectrumProgressBar) SetSideLabel() *spectrumProgressBar {
    return c.SideLabel(true)
}

// The size of the progress bar
func (c *spectrumProgressBar) Size(size ProgressBarSize) *spectrumProgressBar {
    c.PropSize = size
    return c
}

func (c *spectrumProgressBar) SizeS() *spectrumProgressBar {
    return c.Size(ProgressBarSizeS)
}
func (c *spectrumProgressBar) SizeM() *spectrumProgressBar {
    return c.Size(ProgressBarSizeM)
}
func (c *spectrumProgressBar) SizeL() *spectrumProgressBar {
    return c.Size(ProgressBarSizeL)
}
func (c *spectrumProgressBar) SizeXl() *spectrumProgressBar {
    return c.Size(ProgressBarSizeXl)
}
// A static color variant for the progress bar
func (c *spectrumProgressBar) StaticColor(staticColor ProgressBarStaticColor) *spectrumProgressBar {
    c.PropStaticColor = staticColor
    return c
}

func (c *spectrumProgressBar) StaticColorWhite() *spectrumProgressBar {
    return c.StaticColor(ProgressBarStaticColorWhite)
}





// Render renders the sp-progress-bar component
func (c *spectrumProgressBar) Render() app.UI {
    element := app.Elem("sp-progress-bar")

    // Set attributes
    if c.PropIndeterminate {
        element = element.Attr("indeterminate", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropOverBackground {
        element = element.Attr("over-background", true)
    }
    if c.PropProgress != 0 {
        element = element.Attr("progress", c.PropProgress)
    }
    if c.PropSideLabel {
        element = element.Attr("side-label", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropStaticColor != "" {
        element = element.Attr("static-color", string(c.PropStaticColor))
    }



    element = element.Styles(c.styler.styles)

    return element
} 