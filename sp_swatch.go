package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwatchBorder represents the Controls the border style of the swatch
type SwatchBorder string

// SwatchBorder values
const (
    SwatchBorderNone SwatchBorder = "none"
    SwatchBorderLight SwatchBorder = "light"
)
// SwatchRounding represents the Controls the corner rounding of the swatch
type SwatchRounding string

// SwatchRounding values
const (
    SwatchRoundingNone SwatchRounding = "none"
    SwatchRoundingFull SwatchRounding = "full"
)
// SwatchShape represents the The shape of the swatch (default is square)
type SwatchShape string

// SwatchShape values
const (
    SwatchShapeRectangle SwatchShape = "rectangle"
)
// SwatchSize represents the The size of the swatch
type SwatchSize string

// SwatchSize values
const (
    SwatchSizeXs SwatchSize = "xs"
    SwatchSizeS SwatchSize = "s"
    SwatchSizeM SwatchSize = "m"
    SwatchSizeL SwatchSize = "l"
)

// spectrumSwatch represents an sp-swatch component
type spectrumSwatch struct {
    app.Compo
    *styler[*spectrumSwatch]

    // Properties
    // Controls the border style of the swatch
    PropBorder SwatchBorder
    // The color value that the swatch will display
    PropColor string
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // An accessible label for the swatch
    PropLabel string
    // Indicates the swatch represents more than one color
    PropMixedValue bool
    // Indicates the swatch represents no color or transparent
    PropNothing bool
    // Controls the corner rounding of the swatch
    PropRounding SwatchRounding
    // Whether the swatch is selected
    PropSelected bool
    // The shape of the swatch (default is square)
    PropShape SwatchShape
    // The size of the swatch
    PropSize SwatchSize




    // Event handlers
    PropOnChange app.EventHandler
}

// Swatch creates a new sp-swatch component
func Swatch() *spectrumSwatch {
    element := &spectrumSwatch{
        PropColor: "",
        PropDisabled: false,
        PropLabel: "",
        PropMixedValue: false,
        PropNothing: false,
        PropSelected: false,
        PropSize: SwatchSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Controls the border style of the swatch
func (c *spectrumSwatch) Border(border SwatchBorder) *spectrumSwatch {
    c.PropBorder = border
    return c
}

func (c *spectrumSwatch) BorderNone() *spectrumSwatch {
    return c.Border(SwatchBorderNone)
}
func (c *spectrumSwatch) BorderLight() *spectrumSwatch {
    return c.Border(SwatchBorderLight)
}
// The color value that the swatch will display
func (c *spectrumSwatch) Color(color string) *spectrumSwatch {
    c.PropColor = color
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumSwatch) Disabled(disabled bool) *spectrumSwatch {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSwatch) SetDisabled() *spectrumSwatch {
    return c.Disabled(true)
}

// An accessible label for the swatch
func (c *spectrumSwatch) Label(label string) *spectrumSwatch {
    c.PropLabel = label
    return c
}

// Indicates the swatch represents more than one color
func (c *spectrumSwatch) MixedValue(mixedValue bool) *spectrumSwatch {
    c.PropMixedValue = mixedValue
    return c
}

func (c *spectrumSwatch) SetMixedValue() *spectrumSwatch {
    return c.MixedValue(true)
}

// Indicates the swatch represents no color or transparent
func (c *spectrumSwatch) Nothing(nothing bool) *spectrumSwatch {
    c.PropNothing = nothing
    return c
}

func (c *spectrumSwatch) SetNothing() *spectrumSwatch {
    return c.Nothing(true)
}

// Controls the corner rounding of the swatch
func (c *spectrumSwatch) Rounding(rounding SwatchRounding) *spectrumSwatch {
    c.PropRounding = rounding
    return c
}

func (c *spectrumSwatch) RoundingNone() *spectrumSwatch {
    return c.Rounding(SwatchRoundingNone)
}
func (c *spectrumSwatch) RoundingFull() *spectrumSwatch {
    return c.Rounding(SwatchRoundingFull)
}
// Whether the swatch is selected
func (c *spectrumSwatch) Selected(selected bool) *spectrumSwatch {
    c.PropSelected = selected
    return c
}

func (c *spectrumSwatch) SetSelected() *spectrumSwatch {
    return c.Selected(true)
}

// The shape of the swatch (default is square)
func (c *spectrumSwatch) Shape(shape SwatchShape) *spectrumSwatch {
    c.PropShape = shape
    return c
}

func (c *spectrumSwatch) ShapeRectangle() *spectrumSwatch {
    return c.Shape(SwatchShapeRectangle)
}
// The size of the swatch
func (c *spectrumSwatch) Size(size SwatchSize) *spectrumSwatch {
    c.PropSize = size
    return c
}

func (c *spectrumSwatch) SizeXs() *spectrumSwatch {
    return c.Size(SwatchSizeXs)
}
func (c *spectrumSwatch) SizeS() *spectrumSwatch {
    return c.Size(SwatchSizeS)
}
func (c *spectrumSwatch) SizeM() *spectrumSwatch {
    return c.Size(SwatchSizeM)
}
func (c *spectrumSwatch) SizeL() *spectrumSwatch {
    return c.Size(SwatchSizeL)
}




// Fired when the swatch is selected
func (c *spectrumSwatch) OnChange(handler app.EventHandler) *spectrumSwatch {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-swatch component
func (c *spectrumSwatch) Render() app.UI {
    element := app.Elem("sp-swatch")

    // Set attributes
    if c.PropBorder != "" {
        element = element.Attr("border", string(c.PropBorder))
    }
    if c.PropColor != "" {
        element = element.Attr("color", c.PropColor)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMixedValue {
        element = element.Attr("mixed-value", true)
    }
    if c.PropNothing {
        element = element.Attr("nothing", true)
    }
    if c.PropRounding != "" {
        element = element.Attr("rounding", string(c.PropRounding))
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropShape != "" {
        element = element.Attr("shape", string(c.PropShape))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }


    element = element.Styles(c.styler.styles)

    return element
} 