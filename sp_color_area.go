package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorArea represents an sp-color-area component
type spectrumColorArea struct {
    app.Compo
    *styler[*spectrumColorArea]

    // Properties
    // The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
    PropColor string
    // Disables the color area interaction
    PropDisabled bool
    // Indicates whether the color area has focus
    PropFocused bool
    // The hue value to use for the color area's gradient
    PropHue float64
    // Accessible label for the x-axis of the color area
    PropLabelX string
    // Accessible label for the y-axis of the color area
    PropLabelY string
    // The increment amount for keyboard navigation
    PropStep float64
    // The x-coordinate value (typically represents saturation)
    PropX float64
    // The y-coordinate value (typically represents luminosity)
    PropY float64


    // Content slots
    PropGradientSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// ColorArea creates a new sp-color-area component
func ColorArea() *spectrumColorArea {
    element := &spectrumColorArea{
        PropDisabled: false,
        PropFocused: false,
        PropHue: 0,
        PropLabelX: "saturation",
        PropLabelY: "luminosity",
        PropStep: 0.01,
        PropX: 0,
        PropY: 0,
    }

    element.styler = newStyler(element)

    return element
}

// The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
func (c *spectrumColorArea) Color(color string) *spectrumColorArea {
    c.PropColor = color
    return c
}

// Disables the color area interaction
func (c *spectrumColorArea) Disabled(disabled bool) *spectrumColorArea {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumColorArea) SetDisabled() *spectrumColorArea {
    return c.Disabled(true)
}

// Indicates whether the color area has focus
func (c *spectrumColorArea) Focused(focused bool) *spectrumColorArea {
    c.PropFocused = focused
    return c
}

func (c *spectrumColorArea) SetFocused() *spectrumColorArea {
    return c.Focused(true)
}

// The hue value to use for the color area's gradient
func (c *spectrumColorArea) Hue(hue float64) *spectrumColorArea {
    c.PropHue = hue
    return c
}

// Accessible label for the x-axis of the color area
func (c *spectrumColorArea) LabelX(labelX string) *spectrumColorArea {
    c.PropLabelX = labelX
    return c
}

// Accessible label for the y-axis of the color area
func (c *spectrumColorArea) LabelY(labelY string) *spectrumColorArea {
    c.PropLabelY = labelY
    return c
}

// The increment amount for keyboard navigation
func (c *spectrumColorArea) Step(step float64) *spectrumColorArea {
    c.PropStep = step
    return c
}

// The x-coordinate value (typically represents saturation)
func (c *spectrumColorArea) X(x float64) *spectrumColorArea {
    c.PropX = x
    return c
}

// The y-coordinate value (typically represents luminosity)
func (c *spectrumColorArea) Y(y float64) *spectrumColorArea {
    c.PropY = y
    return c
}



// A custom gradient visually outlining the available color values
func (c *spectrumColorArea) Gradient(content app.UI) *spectrumColorArea {
    c.PropGradientSlot = content

    return c
}



// An alteration to the value of the Color Area has been committed by the user
func (c *spectrumColorArea) OnChange(handler app.EventHandler) *spectrumColorArea {
    c.PropOnChange = handler

    return c
}

// The value of the Color Area has changed
func (c *spectrumColorArea) OnInput(handler app.EventHandler) *spectrumColorArea {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-color-area component
func (c *spectrumColorArea) Render() app.UI {
    element := app.Elem("sp-color-area")

    // Set attributes
    if c.PropColor != "" {
        element = element.Attr("color", c.PropColor)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFocused {
        element = element.Attr("focused", true)
    }
    if c.PropHue != 0 {
        element = element.Attr("hue", c.PropHue)
    }
    if c.PropLabelX != "" {
        element = element.Attr("label-x", c.PropLabelX)
    }
    if c.PropLabelY != "" {
        element = element.Attr("label-y", c.PropLabelY)
    }
    if c.PropStep != 0 {
        element = element.Attr("step", c.PropStep)
    }
    if c.PropX != 0 {
        element = element.Attr("x", c.PropX)
    }
    if c.PropY != 0 {
        element = element.Attr("y", c.PropY)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnInput != nil {
        element = element.On("input", c.PropOnInput)
    }

    // Add slots and children
    slotElements := []app.UI{}


    // Add gradient slot
    if c.PropGradientSlot != nil {
        slotElem := c.PropGradientSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("gradient")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "gradient").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 