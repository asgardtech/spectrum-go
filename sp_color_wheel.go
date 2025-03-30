package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorWheel represents an sp-color-wheel component
type spectrumColorWheel struct {
    app.Compo
    *styler[*spectrumColorWheel]

    // Properties
    // The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
    PropColor string
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Indicates whether the wheel has focus
    PropFocused bool
    // Accessible label for the wheel
    PropLabel string
    // The increment amount for keyboard navigation
    PropStep float64
    // The tab index to apply to this control
    PropTabindex float64
    // The current value of the wheel (hue)
    PropValue float64


    // Content slots
    PropGradientSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// ColorWheel creates a new sp-color-wheel component
func ColorWheel() *spectrumColorWheel {
    element := &spectrumColorWheel{
        PropDisabled: false,
        PropFocused: false,
        PropLabel: "hue",
        PropStep: 1,
        PropTabindex: 0,
        PropValue: 0,
    }

    element.styler = newStyler(element)

    return element
}

// The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
func (c *spectrumColorWheel) Color(color string) *spectrumColorWheel {
    c.PropColor = color
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumColorWheel) Disabled(disabled bool) *spectrumColorWheel {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumColorWheel) SetDisabled() *spectrumColorWheel {
    return c.Disabled(true)
}

// Indicates whether the wheel has focus
func (c *spectrumColorWheel) Focused(focused bool) *spectrumColorWheel {
    c.PropFocused = focused
    return c
}

func (c *spectrumColorWheel) SetFocused() *spectrumColorWheel {
    return c.Focused(true)
}

// Accessible label for the wheel
func (c *spectrumColorWheel) Label(label string) *spectrumColorWheel {
    c.PropLabel = label
    return c
}

// The increment amount for keyboard navigation
func (c *spectrumColorWheel) Step(step float64) *spectrumColorWheel {
    c.PropStep = step
    return c
}

// The tab index to apply to this control
func (c *spectrumColorWheel) Tabindex(tabindex float64) *spectrumColorWheel {
    c.PropTabindex = tabindex
    return c
}

// The current value of the wheel (hue)
func (c *spectrumColorWheel) Value(value float64) *spectrumColorWheel {
    c.PropValue = value
    return c
}



// A custom gradient visually outlining the available color values
func (c *spectrumColorWheel) Gradient(content app.UI) *spectrumColorWheel {
    c.PropGradientSlot = content

    return c
}



// An alteration to the value of the Color Wheel has been committed by the user
func (c *spectrumColorWheel) OnChange(handler app.EventHandler) *spectrumColorWheel {
    c.PropOnChange = handler

    return c
}

// The value of the Color Wheel has changed
func (c *spectrumColorWheel) OnInput(handler app.EventHandler) *spectrumColorWheel {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-color-wheel component
func (c *spectrumColorWheel) Render() app.UI {
    element := app.Elem("sp-color-wheel")

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