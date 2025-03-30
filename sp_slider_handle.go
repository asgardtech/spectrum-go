package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumSliderHandle represents an sp-slider-handle component
type spectrumSliderHandle struct {
    app.Compo
    *styler[*spectrumSliderHandle]

    // Properties
    // Set the default value of the handle. Setting this property will cause the handle to reset to the default value on double click or pressing the escape key.
    PropDefaultValue float64
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether the handle is currently being dragged
    PropDragging bool
    // Number formatting options as a JSON string
    PropFormatOptions string
    // Whether the handle should be highlighted
    PropHighlight bool
    // Accessible text label for the handle
    PropLabel string
    // Maximum value of the handle. Can be a number or 'next' to constrain by the next handle's value
    PropMax string
    // Minimum value of the handle. Can be a number or 'previous' to constrain by the previous handle's value
    PropMin string
    // Name of the handle for form submission
    PropName string
    // Step value for increments/decrements
    PropStep float64
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // The current value of the handle. By default, the value will be halfway between min and max values, or the min value when max is less than min.
    PropValue float64




    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
    PropOnSpSliderHandleReady app.EventHandler
}

// SliderHandle creates a new sp-slider-handle component
func SliderHandle() *spectrumSliderHandle {
    element := &spectrumSliderHandle{
        PropDefaultValue: 0,
        PropDisabled: false,
        PropDragging: false,
        PropHighlight: false,
        PropLabel: "",
        PropName: "",
        PropStep: 0,
        PropTabindex: 0,
        PropValue: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Set the default value of the handle. Setting this property will cause the handle to reset to the default value on double click or pressing the escape key.
func (c *spectrumSliderHandle) DefaultValue(defaultValue float64) *spectrumSliderHandle {
    c.PropDefaultValue = defaultValue
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumSliderHandle) Disabled(disabled bool) *spectrumSliderHandle {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSliderHandle) SetDisabled() *spectrumSliderHandle {
    return c.Disabled(true)
}

// Whether the handle is currently being dragged
func (c *spectrumSliderHandle) Dragging(dragging bool) *spectrumSliderHandle {
    c.PropDragging = dragging
    return c
}

func (c *spectrumSliderHandle) SetDragging() *spectrumSliderHandle {
    return c.Dragging(true)
}

// Number formatting options as a JSON string
func (c *spectrumSliderHandle) FormatOptions(formatOptions string) *spectrumSliderHandle {
    c.PropFormatOptions = formatOptions
    return c
}

// Whether the handle should be highlighted
func (c *spectrumSliderHandle) Highlight(highlight bool) *spectrumSliderHandle {
    c.PropHighlight = highlight
    return c
}

func (c *spectrumSliderHandle) SetHighlight() *spectrumSliderHandle {
    return c.Highlight(true)
}

// Accessible text label for the handle
func (c *spectrumSliderHandle) Label(label string) *spectrumSliderHandle {
    c.PropLabel = label
    return c
}

// Maximum value of the handle. Can be a number or 'next' to constrain by the next handle's value
func (c *spectrumSliderHandle) Max(max string) *spectrumSliderHandle {
    c.PropMax = max
    return c
}

// Minimum value of the handle. Can be a number or 'previous' to constrain by the previous handle's value
func (c *spectrumSliderHandle) Min(min string) *spectrumSliderHandle {
    c.PropMin = min
    return c
}

// Name of the handle for form submission
func (c *spectrumSliderHandle) Name(name string) *spectrumSliderHandle {
    c.PropName = name
    return c
}

// Step value for increments/decrements
func (c *spectrumSliderHandle) Step(step float64) *spectrumSliderHandle {
    c.PropStep = step
    return c
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumSliderHandle) Tabindex(tabindex float64) *spectrumSliderHandle {
    c.PropTabindex = tabindex
    return c
}

// The current value of the handle. By default, the value will be halfway between min and max values, or the min value when max is less than min.
func (c *spectrumSliderHandle) Value(value float64) *spectrumSliderHandle {
    c.PropValue = value
    return c
}





// An alteration to the value of the element has been committed by the user.
func (c *spectrumSliderHandle) OnChange(handler app.EventHandler) *spectrumSliderHandle {
    c.PropOnChange = handler

    return c
}

// The value of the element has changed.
func (c *spectrumSliderHandle) OnInput(handler app.EventHandler) *spectrumSliderHandle {
    c.PropOnInput = handler

    return c
}

// Fired when a slider handle has been initialized
func (c *spectrumSliderHandle) OnSpSliderHandleReady(handler app.EventHandler) *spectrumSliderHandle {
    c.PropOnSpSliderHandleReady = handler

    return c
}


// Render renders the sp-slider-handle component
func (c *spectrumSliderHandle) Render() app.UI {
    element := app.Elem("sp-slider-handle")

    // Set attributes
    if c.PropDefaultValue != 0 {
        element = element.Attr("default-value", c.PropDefaultValue)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDragging {
        element = element.Attr("dragging", true)
    }
    if c.PropFormatOptions != "" {
        element = element.Attr("format-options", c.PropFormatOptions)
    }
    if c.PropHighlight {
        element = element.Attr("highlight", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMax != "" {
        element = element.Attr("max", c.PropMax)
    }
    if c.PropMin != "" {
        element = element.Attr("min", c.PropMin)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
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
    if c.PropOnSpSliderHandleReady != nil {
        element = element.On("sp-slider-handle-ready", c.PropOnSpSliderHandleReady)
    }


    element = element.Styles(c.styler.styles)

    return element
} 