package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumNumberField represents an sp-number-field component
type spectrumNumberField struct {
    app.Compo
    *styler[*spectrumNumberField]

    // Properties
    // Form of assistance provided when supplying a value
    PropAutocomplete string
    // Disable this control
    PropDisabled bool
    // Intl.NumberFormatOptions object for value formatting
    PropFormatoptions any
    // Whether the stepper UI is hidden
    PropHidestepper bool
    // Whether the value is invalid
    PropInvalid bool
    // Accessible label when visible label not provided
    PropLabel string
    // Maximum allowed value
    PropMax float64
    // Minimum allowed value
    PropMin float64
    // Text shown when no value is set
    PropPlaceholder string
    // Display without visible background
    PropQuiet bool
    // Whether user can interact with the value
    PropReadonly bool
    // Whether field is required
    PropRequired bool
    // Increment/decrement step value
    PropStep float64
    // Multiplier for step when using shift key
    PropStepmodifier float64
    // The value held by the control
    PropValue string


    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// NumberField creates a new sp-number-field component
func NumberField() *spectrumNumberField {
    element := &spectrumNumberField{
        PropAutocomplete: "",
        PropDisabled: false,
        PropFormatoptions: map[string]interface{}{},
        PropHidestepper: false,
        PropInvalid: false,
        PropLabel: "",
        PropMax: 0,
        PropMin: 0,
        PropPlaceholder: "",
        PropQuiet: false,
        PropReadonly: false,
        PropRequired: false,
        PropStep: 1,
        PropStepmodifier: 10,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Form of assistance provided when supplying a value
func (c *spectrumNumberField) Autocomplete(autocomplete string) *spectrumNumberField {
    c.PropAutocomplete = autocomplete
    return c
}

// Disable this control
func (c *spectrumNumberField) Disabled(disabled bool) *spectrumNumberField {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumNumberField) SetDisabled() *spectrumNumberField {
    return c.Disabled(true)
}

// Intl.NumberFormatOptions object for value formatting
func (c *spectrumNumberField) Formatoptions(formatOptions any) *spectrumNumberField {
    c.PropFormatoptions = formatOptions
    return c
}

// Whether the stepper UI is hidden
func (c *spectrumNumberField) Hidestepper(hideStepper bool) *spectrumNumberField {
    c.PropHidestepper = hideStepper
    return c
}

func (c *spectrumNumberField) SetHidestepper() *spectrumNumberField {
    return c.Hidestepper(true)
}

// Whether the value is invalid
func (c *spectrumNumberField) Invalid(invalid bool) *spectrumNumberField {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumNumberField) SetInvalid() *spectrumNumberField {
    return c.Invalid(true)
}

// Accessible label when visible label not provided
func (c *spectrumNumberField) Label(label string) *spectrumNumberField {
    c.PropLabel = label
    return c
}

// Maximum allowed value
func (c *spectrumNumberField) Max(max float64) *spectrumNumberField {
    c.PropMax = max
    return c
}

// Minimum allowed value
func (c *spectrumNumberField) Min(min float64) *spectrumNumberField {
    c.PropMin = min
    return c
}

// Text shown when no value is set
func (c *spectrumNumberField) Placeholder(placeholder string) *spectrumNumberField {
    c.PropPlaceholder = placeholder
    return c
}

// Display without visible background
func (c *spectrumNumberField) Quiet(quiet bool) *spectrumNumberField {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumNumberField) SetQuiet() *spectrumNumberField {
    return c.Quiet(true)
}

// Whether user can interact with the value
func (c *spectrumNumberField) Readonly(readonly bool) *spectrumNumberField {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumNumberField) SetReadonly() *spectrumNumberField {
    return c.Readonly(true)
}

// Whether field is required
func (c *spectrumNumberField) Required(required bool) *spectrumNumberField {
    c.PropRequired = required
    return c
}

func (c *spectrumNumberField) SetRequired() *spectrumNumberField {
    return c.Required(true)
}

// Increment/decrement step value
func (c *spectrumNumberField) Step(step float64) *spectrumNumberField {
    c.PropStep = step
    return c
}

// Multiplier for step when using shift key
func (c *spectrumNumberField) Stepmodifier(stepModifier float64) *spectrumNumberField {
    c.PropStepmodifier = stepModifier
    return c
}

// The value held by the control
func (c *spectrumNumberField) Value(value string) *spectrumNumberField {
    c.PropValue = value
    return c
}



// Default or non-negative help text
func (c *spectrumNumberField) HelpText(content app.UI) *spectrumNumberField {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text shown when invalid
func (c *spectrumNumberField) NegativeHelpText(content app.UI) *spectrumNumberField {
    c.PropNegativeHelpTextSlot = content

    return c
}



// Value has been committed by the user
func (c *spectrumNumberField) OnChange(handler app.EventHandler) *spectrumNumberField {
    c.PropOnChange = handler

    return c
}

// Value of the element has changed
func (c *spectrumNumberField) OnInput(handler app.EventHandler) *spectrumNumberField {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-number-field component
func (c *spectrumNumberField) Render() app.UI {
    element := app.Elem("sp-number-field")

    // Set attributes
    if c.PropAutocomplete != "" {
        element = element.Attr("autocomplete", c.PropAutocomplete)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFormatoptions != nil {
        element = element.Attr("formatOptions", c.PropFormatoptions)
    }
    if c.PropHidestepper {
        element = element.Attr("hideStepper", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMax != 0 {
        element = element.Attr("max", c.PropMax)
    }
    if c.PropMin != 0 {
        element = element.Attr("min", c.PropMin)
    }
    if c.PropPlaceholder != "" {
        element = element.Attr("placeholder", c.PropPlaceholder)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropRequired {
        element = element.Attr("required", true)
    }
    if c.PropStep != 0 {
        element = element.Attr("step", c.PropStep)
    }
    if c.PropStepmodifier != 0 {
        element = element.Attr("stepModifier", c.PropStepmodifier)
    }
    if c.PropValue != "" {
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


    // Add help-text slot
    if c.PropHelpTextSlot != nil {
        slotElem := c.PropHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "help-text").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add negative-help-text slot
    if c.PropNegativeHelpTextSlot != nil {
        slotElem := c.PropNegativeHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("negative-help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "negative-help-text").
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