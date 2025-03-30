package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorSlider represents an sp-color-slider component
type spectrumColorSlider struct {
    app.Compo
    *styler[*spectrumColorSlider]

    // Properties
    // The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
    PropColor string
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Indicates whether the slider has focus
    PropFocused bool
    // Accessible label for the slider
    PropLabel string
    // The increment amount for keyboard navigation
    PropStep float64
    // The tab index to apply to this control
    PropTabindex float64
    // The current value of the slider
    PropValue float64
    // Whether the slider is displayed vertically
    PropVertical bool


    // Content slots
    PropGradientSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// ColorSlider creates a new sp-color-slider component
func ColorSlider() *spectrumColorSlider {
    element := &spectrumColorSlider{
        PropDisabled: false,
        PropFocused: false,
        PropLabel: "hue",
        PropStep: 1,
        PropTabindex: 0,
        PropValue: 0,
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
func (c *spectrumColorSlider) Color(color string) *spectrumColorSlider {
    c.PropColor = color
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumColorSlider) Disabled(disabled bool) *spectrumColorSlider {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumColorSlider) SetDisabled() *spectrumColorSlider {
    return c.Disabled(true)
}

// Indicates whether the slider has focus
func (c *spectrumColorSlider) Focused(focused bool) *spectrumColorSlider {
    c.PropFocused = focused
    return c
}

func (c *spectrumColorSlider) SetFocused() *spectrumColorSlider {
    return c.Focused(true)
}

// Accessible label for the slider
func (c *spectrumColorSlider) Label(label string) *spectrumColorSlider {
    c.PropLabel = label
    return c
}

// The increment amount for keyboard navigation
func (c *spectrumColorSlider) Step(step float64) *spectrumColorSlider {
    c.PropStep = step
    return c
}

// The tab index to apply to this control
func (c *spectrumColorSlider) Tabindex(tabindex float64) *spectrumColorSlider {
    c.PropTabindex = tabindex
    return c
}

// The current value of the slider
func (c *spectrumColorSlider) Value(value float64) *spectrumColorSlider {
    c.PropValue = value
    return c
}

// Whether the slider is displayed vertically
func (c *spectrumColorSlider) Vertical(vertical bool) *spectrumColorSlider {
    c.PropVertical = vertical
    return c
}

func (c *spectrumColorSlider) SetVertical() *spectrumColorSlider {
    return c.Vertical(true)
}



// A custom gradient visually outlining the available color values
func (c *spectrumColorSlider) Gradient(content app.UI) *spectrumColorSlider {
    c.PropGradientSlot = content

    return c
}



// An alteration to the value of the Color Slider has been committed by the user
func (c *spectrumColorSlider) OnChange(handler app.EventHandler) *spectrumColorSlider {
    c.PropOnChange = handler

    return c
}

// The value of the Color Slider has changed
func (c *spectrumColorSlider) OnInput(handler app.EventHandler) *spectrumColorSlider {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-color-slider component
func (c *spectrumColorSlider) Render() app.UI {
    element := app.Elem("sp-color-slider")

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
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropStep != 0 {
        element = element.Attr("step", c.PropStep)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }
    if c.PropValue != 0 {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
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