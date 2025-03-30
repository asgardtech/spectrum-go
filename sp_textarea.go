package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TextareaSize represents the The size of the textarea
type TextareaSize string

// TextareaSize values
const (
    TextareaSizeS TextareaSize = "s"
    TextareaSizeM TextareaSize = "m"
    TextareaSizeL TextareaSize = "l"
    TextareaSizeXl TextareaSize = "xl"
)

// spectrumTextarea represents an sp-textfield component
type spectrumTextarea struct {
    app.Compo
    *styler[*spectrumTextarea]

    // Properties
    // What form of assistance should be provided when attempting to supply a value to the form control
    PropAutocomplete string
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether a form control delivered with the multiline attribute will change size vertically to accommodate longer input
    PropGrows bool
    // Whether the value held by the form control is invalid
    PropInvalid bool
    // A string applied via aria-label to the form control when a user visible label is not provided
    PropLabel string
    // Defines the maximum string length that the user can enter
    PropMaxlength float64
    // Defines the minimum string length that the user can enter
    PropMinlength float64
    // Whether the form control should accept a value longer than one line
    PropMultiline bool
    // Name of the form control
    PropName string
    // Pattern the value must match to be valid
    PropPattern string
    // Text that appears in the form control when it has no value set
    PropPlaceholder string
    // Whether to display the form control with no visible background
    PropQuiet bool
    // Whether a user can interact with the value of the form control
    PropReadonly bool
    // Whether the form control will be found to be invalid when it holds no value
    PropRequired bool
    // The specific number of rows the form control should provide in the user interface
    PropRows float64
    // The size of the textarea
    PropSize TextareaSize
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Whether the value held by the form control is valid
    PropValid bool
    // The value held by the form control
    PropValue string


    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// Textarea creates a new sp-textfield component
func Textarea() *spectrumTextarea {
    element := &spectrumTextarea{
        PropDisabled: false,
        PropGrows: false,
        PropInvalid: false,
        PropLabel: "",
        PropMaxlength: -1,
        PropMinlength: -1,
        PropMultiline: true,
        PropPlaceholder: "",
        PropQuiet: false,
        PropReadonly: false,
        PropRequired: false,
        PropRows: -1,
        PropSize: TextareaSizeM,
        PropTabindex: 0,
        PropValid: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// What form of assistance should be provided when attempting to supply a value to the form control
func (c *spectrumTextarea) Autocomplete(autocomplete string) *spectrumTextarea {
    c.PropAutocomplete = autocomplete
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumTextarea) Disabled(disabled bool) *spectrumTextarea {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumTextarea) SetDisabled() *spectrumTextarea {
    return c.Disabled(true)
}

// Whether a form control delivered with the multiline attribute will change size vertically to accommodate longer input
func (c *spectrumTextarea) Grows(grows bool) *spectrumTextarea {
    c.PropGrows = grows
    return c
}

func (c *spectrumTextarea) SetGrows() *spectrumTextarea {
    return c.Grows(true)
}

// Whether the value held by the form control is invalid
func (c *spectrumTextarea) Invalid(invalid bool) *spectrumTextarea {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumTextarea) SetInvalid() *spectrumTextarea {
    return c.Invalid(true)
}

// A string applied via aria-label to the form control when a user visible label is not provided
func (c *spectrumTextarea) Label(label string) *spectrumTextarea {
    c.PropLabel = label
    return c
}

// Defines the maximum string length that the user can enter
func (c *spectrumTextarea) Maxlength(maxlength float64) *spectrumTextarea {
    c.PropMaxlength = maxlength
    return c
}

// Defines the minimum string length that the user can enter
func (c *spectrumTextarea) Minlength(minlength float64) *spectrumTextarea {
    c.PropMinlength = minlength
    return c
}

// Whether the form control should accept a value longer than one line
func (c *spectrumTextarea) Multiline(multiline bool) *spectrumTextarea {
    c.PropMultiline = multiline
    return c
}

func (c *spectrumTextarea) SetMultiline() *spectrumTextarea {
    return c.Multiline(true)
}

// Name of the form control
func (c *spectrumTextarea) Name(name string) *spectrumTextarea {
    c.PropName = name
    return c
}

// Pattern the value must match to be valid
func (c *spectrumTextarea) Pattern(pattern string) *spectrumTextarea {
    c.PropPattern = pattern
    return c
}

// Text that appears in the form control when it has no value set
func (c *spectrumTextarea) Placeholder(placeholder string) *spectrumTextarea {
    c.PropPlaceholder = placeholder
    return c
}

// Whether to display the form control with no visible background
func (c *spectrumTextarea) Quiet(quiet bool) *spectrumTextarea {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumTextarea) SetQuiet() *spectrumTextarea {
    return c.Quiet(true)
}

// Whether a user can interact with the value of the form control
func (c *spectrumTextarea) Readonly(readonly bool) *spectrumTextarea {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumTextarea) SetReadonly() *spectrumTextarea {
    return c.Readonly(true)
}

// Whether the form control will be found to be invalid when it holds no value
func (c *spectrumTextarea) Required(required bool) *spectrumTextarea {
    c.PropRequired = required
    return c
}

func (c *spectrumTextarea) SetRequired() *spectrumTextarea {
    return c.Required(true)
}

// The specific number of rows the form control should provide in the user interface
func (c *spectrumTextarea) Rows(rows float64) *spectrumTextarea {
    c.PropRows = rows
    return c
}

// The size of the textarea
func (c *spectrumTextarea) Size(size TextareaSize) *spectrumTextarea {
    c.PropSize = size
    return c
}

func (c *spectrumTextarea) SizeS() *spectrumTextarea {
    return c.Size(TextareaSizeS)
}
func (c *spectrumTextarea) SizeM() *spectrumTextarea {
    return c.Size(TextareaSizeM)
}
func (c *spectrumTextarea) SizeL() *spectrumTextarea {
    return c.Size(TextareaSizeL)
}
func (c *spectrumTextarea) SizeXl() *spectrumTextarea {
    return c.Size(TextareaSizeXl)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumTextarea) Tabindex(tabindex float64) *spectrumTextarea {
    c.PropTabindex = tabindex
    return c
}

// Whether the value held by the form control is valid
func (c *spectrumTextarea) Valid(valid bool) *spectrumTextarea {
    c.PropValid = valid
    return c
}

func (c *spectrumTextarea) SetValid() *spectrumTextarea {
    return c.Valid(true)
}

// The value held by the form control
func (c *spectrumTextarea) Value(value string) *spectrumTextarea {
    c.PropValue = value
    return c
}



// Default or non-negative help text to associate to your form element
func (c *spectrumTextarea) HelpText(content app.UI) *spectrumTextarea {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when invalid
func (c *spectrumTextarea) NegativeHelpText(content app.UI) *spectrumTextarea {
    c.PropNegativeHelpTextSlot = content

    return c
}



// An alteration to the value of the element has been committed by the user
func (c *spectrumTextarea) OnChange(handler app.EventHandler) *spectrumTextarea {
    c.PropOnChange = handler

    return c
}

// The value of the element has changed
func (c *spectrumTextarea) OnInput(handler app.EventHandler) *spectrumTextarea {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-textfield component
func (c *spectrumTextarea) Render() app.UI {
    element := app.Elem("sp-textfield")

    // Set attributes
    if c.PropAutocomplete != "" {
        element = element.Attr("autocomplete", c.PropAutocomplete)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropGrows {
        element = element.Attr("grows", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMaxlength != 0 {
        element = element.Attr("maxlength", c.PropMaxlength)
    }
    if c.PropMinlength != 0 {
        element = element.Attr("minlength", c.PropMinlength)
    }
    if c.PropMultiline {
        element = element.Attr("multiline", true)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropPattern != "" {
        element = element.Attr("pattern", c.PropPattern)
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
    if c.PropRows != 0 {
        element = element.Attr("rows", c.PropRows)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }
    if c.PropValid {
        element = element.Attr("valid", true)
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