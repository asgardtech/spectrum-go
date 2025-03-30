package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TextfieldSize represents the The size of the textfield
type TextfieldSize string

// TextfieldSize values
const (
    TextfieldSizeS TextfieldSize = "s"
    TextfieldSizeM TextfieldSize = "m"
    TextfieldSizeL TextfieldSize = "l"
    TextfieldSizeXl TextfieldSize = "xl"
)
// TextfieldType represents the The type of input that the textfield will accept
type TextfieldType string

// TextfieldType values
const (
    TextfieldTypeText TextfieldType = "text"
    TextfieldTypePassword TextfieldType = "password"
    TextfieldTypeEmail TextfieldType = "email"
    TextfieldTypeTel TextfieldType = "tel"
    TextfieldTypeUrl TextfieldType = "url"
)

// spectrumTextfield represents an sp-textfield component
type spectrumTextfield struct {
    app.Compo
    *styler[*spectrumTextfield]

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
    // The size of the textfield
    PropSize TextfieldSize
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // The type of input that the textfield will accept
    PropType TextfieldType
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

// Textfield creates a new sp-textfield component
func Textfield() *spectrumTextfield {
    element := &spectrumTextfield{
        PropDisabled: false,
        PropGrows: false,
        PropInvalid: false,
        PropLabel: "",
        PropMaxlength: -1,
        PropMinlength: -1,
        PropMultiline: false,
        PropPlaceholder: "",
        PropQuiet: false,
        PropReadonly: false,
        PropRequired: false,
        PropRows: -1,
        PropSize: TextfieldSizeM,
        PropTabindex: 0,
        PropType: TextfieldTypeText,
        PropValid: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// What form of assistance should be provided when attempting to supply a value to the form control
func (c *spectrumTextfield) Autocomplete(autocomplete string) *spectrumTextfield {
    c.PropAutocomplete = autocomplete
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumTextfield) Disabled(disabled bool) *spectrumTextfield {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumTextfield) SetDisabled() *spectrumTextfield {
    return c.Disabled(true)
}

// Whether a form control delivered with the multiline attribute will change size vertically to accommodate longer input
func (c *spectrumTextfield) Grows(grows bool) *spectrumTextfield {
    c.PropGrows = grows
    return c
}

func (c *spectrumTextfield) SetGrows() *spectrumTextfield {
    return c.Grows(true)
}

// Whether the value held by the form control is invalid
func (c *spectrumTextfield) Invalid(invalid bool) *spectrumTextfield {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumTextfield) SetInvalid() *spectrumTextfield {
    return c.Invalid(true)
}

// A string applied via aria-label to the form control when a user visible label is not provided
func (c *spectrumTextfield) Label(label string) *spectrumTextfield {
    c.PropLabel = label
    return c
}

// Defines the maximum string length that the user can enter
func (c *spectrumTextfield) Maxlength(maxlength float64) *spectrumTextfield {
    c.PropMaxlength = maxlength
    return c
}

// Defines the minimum string length that the user can enter
func (c *spectrumTextfield) Minlength(minlength float64) *spectrumTextfield {
    c.PropMinlength = minlength
    return c
}

// Whether the form control should accept a value longer than one line
func (c *spectrumTextfield) Multiline(multiline bool) *spectrumTextfield {
    c.PropMultiline = multiline
    return c
}

func (c *spectrumTextfield) SetMultiline() *spectrumTextfield {
    return c.Multiline(true)
}

// Name of the form control
func (c *spectrumTextfield) Name(name string) *spectrumTextfield {
    c.PropName = name
    return c
}

// Pattern the value must match to be valid
func (c *spectrumTextfield) Pattern(pattern string) *spectrumTextfield {
    c.PropPattern = pattern
    return c
}

// Text that appears in the form control when it has no value set
func (c *spectrumTextfield) Placeholder(placeholder string) *spectrumTextfield {
    c.PropPlaceholder = placeholder
    return c
}

// Whether to display the form control with no visible background
func (c *spectrumTextfield) Quiet(quiet bool) *spectrumTextfield {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumTextfield) SetQuiet() *spectrumTextfield {
    return c.Quiet(true)
}

// Whether a user can interact with the value of the form control
func (c *spectrumTextfield) Readonly(readonly bool) *spectrumTextfield {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumTextfield) SetReadonly() *spectrumTextfield {
    return c.Readonly(true)
}

// Whether the form control will be found to be invalid when it holds no value
func (c *spectrumTextfield) Required(required bool) *spectrumTextfield {
    c.PropRequired = required
    return c
}

func (c *spectrumTextfield) SetRequired() *spectrumTextfield {
    return c.Required(true)
}

// The specific number of rows the form control should provide in the user interface
func (c *spectrumTextfield) Rows(rows float64) *spectrumTextfield {
    c.PropRows = rows
    return c
}

// The size of the textfield
func (c *spectrumTextfield) Size(size TextfieldSize) *spectrumTextfield {
    c.PropSize = size
    return c
}

func (c *spectrumTextfield) SizeS() *spectrumTextfield {
    return c.Size(TextfieldSizeS)
}
func (c *spectrumTextfield) SizeM() *spectrumTextfield {
    return c.Size(TextfieldSizeM)
}
func (c *spectrumTextfield) SizeL() *spectrumTextfield {
    return c.Size(TextfieldSizeL)
}
func (c *spectrumTextfield) SizeXl() *spectrumTextfield {
    return c.Size(TextfieldSizeXl)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumTextfield) Tabindex(tabindex float64) *spectrumTextfield {
    c.PropTabindex = tabindex
    return c
}

// The type of input that the textfield will accept
func (c *spectrumTextfield) Type(typeValue TextfieldType) *spectrumTextfield {
    c.PropType = typeValue
    return c
}

func (c *spectrumTextfield) TypeText() *spectrumTextfield {
    return c.Type(TextfieldTypeText)
}
func (c *spectrumTextfield) TypePassword() *spectrumTextfield {
    return c.Type(TextfieldTypePassword)
}
func (c *spectrumTextfield) TypeEmail() *spectrumTextfield {
    return c.Type(TextfieldTypeEmail)
}
func (c *spectrumTextfield) TypeTel() *spectrumTextfield {
    return c.Type(TextfieldTypeTel)
}
func (c *spectrumTextfield) TypeUrl() *spectrumTextfield {
    return c.Type(TextfieldTypeUrl)
}
// Whether the value held by the form control is valid
func (c *spectrumTextfield) Valid(valid bool) *spectrumTextfield {
    c.PropValid = valid
    return c
}

func (c *spectrumTextfield) SetValid() *spectrumTextfield {
    return c.Valid(true)
}

// The value held by the form control
func (c *spectrumTextfield) Value(value string) *spectrumTextfield {
    c.PropValue = value
    return c
}



// Default or non-negative help text to associate to your form element
func (c *spectrumTextfield) HelpText(content app.UI) *spectrumTextfield {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when invalid
func (c *spectrumTextfield) NegativeHelpText(content app.UI) *spectrumTextfield {
    c.PropNegativeHelpTextSlot = content

    return c
}



// An alteration to the value of the element has been committed by the user
func (c *spectrumTextfield) OnChange(handler app.EventHandler) *spectrumTextfield {
    c.PropOnChange = handler

    return c
}

// The value of the element has changed
func (c *spectrumTextfield) OnInput(handler app.EventHandler) *spectrumTextfield {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-textfield component
func (c *spectrumTextfield) Render() app.UI {
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
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
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