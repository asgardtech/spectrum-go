package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ProgressCircleSize represents the The size of the progress circle
type ProgressCircleSize string

// ProgressCircleSize values
const (
    ProgressCircleSizeS ProgressCircleSize = "s"
    ProgressCircleSizeM ProgressCircleSize = "m"
    ProgressCircleSizeL ProgressCircleSize = "l"
)
// ProgressCircleStaticColor represents the A static color variant for when the progress circle is displayed over content
type ProgressCircleStaticColor string

// ProgressCircleStaticColor values
const (
    ProgressCircleStaticColorWhite ProgressCircleStaticColor = "white"
)

// spectrumProgressCircle represents an sp-progress-circle component
type spectrumProgressCircle struct {
    app.Compo
    *styler[*spectrumProgressCircle]

    // Properties
    // Whether progress is indeterminate (unknown duration)
    PropIndeterminate bool
    // Accessible label for the progress circle (not visually displayed)
    PropLabel string
    // The progress value from 0-100
    PropProgress float64
    // The size of the progress circle
    PropSize ProgressCircleSize
    // A static color variant for when the progress circle is displayed over content
    PropStaticColor ProgressCircleStaticColor




}

// ProgressCircle creates a new sp-progress-circle component
func ProgressCircle() *spectrumProgressCircle {
    element := &spectrumProgressCircle{
        PropIndeterminate: false,
        PropLabel: "",
        PropProgress: 0,
        PropSize: ProgressCircleSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Whether progress is indeterminate (unknown duration)
func (c *spectrumProgressCircle) Indeterminate(indeterminate bool) *spectrumProgressCircle {
    c.PropIndeterminate = indeterminate
    return c
}

func (c *spectrumProgressCircle) SetIndeterminate() *spectrumProgressCircle {
    return c.Indeterminate(true)
}

// Accessible label for the progress circle (not visually displayed)
func (c *spectrumProgressCircle) Label(label string) *spectrumProgressCircle {
    c.PropLabel = label
    return c
}

// The progress value from 0-100
func (c *spectrumProgressCircle) Progress(progress float64) *spectrumProgressCircle {
    c.PropProgress = progress
    return c
}

// The size of the progress circle
func (c *spectrumProgressCircle) Size(size ProgressCircleSize) *spectrumProgressCircle {
    c.PropSize = size
    return c
}

func (c *spectrumProgressCircle) SizeS() *spectrumProgressCircle {
    return c.Size(ProgressCircleSizeS)
}
func (c *spectrumProgressCircle) SizeM() *spectrumProgressCircle {
    return c.Size(ProgressCircleSizeM)
}
func (c *spectrumProgressCircle) SizeL() *spectrumProgressCircle {
    return c.Size(ProgressCircleSizeL)
}
// A static color variant for when the progress circle is displayed over content
func (c *spectrumProgressCircle) StaticColor(staticColor ProgressCircleStaticColor) *spectrumProgressCircle {
    c.PropStaticColor = staticColor
    return c
}

func (c *spectrumProgressCircle) StaticColorWhite() *spectrumProgressCircle {
    return c.StaticColor(ProgressCircleStaticColorWhite)
}





// Render renders the sp-progress-circle component
func (c *spectrumProgressCircle) Render() app.UI {
    element := app.Elem("sp-progress-circle")

    // Set attributes
    if c.PropIndeterminate {
        element = element.Attr("indeterminate", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropProgress != 0 {
        element = element.Attr("progress", c.PropProgress)
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