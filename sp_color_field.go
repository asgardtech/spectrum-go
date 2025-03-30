package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ColorFieldAutocomplete represents the What form of assistance should be provided when attempting to supply a value to the form control
type ColorFieldAutocomplete string

// ColorFieldAutocomplete values
const (
    ColorFieldAutocompleteList ColorFieldAutocomplete = "list"
    ColorFieldAutocompleteNone ColorFieldAutocomplete = "none"
)
// ColorFieldSize represents the Size of the color field
type ColorFieldSize string

// ColorFieldSize values
const (
    ColorFieldSizeS ColorFieldSize = "s"
    ColorFieldSizeM ColorFieldSize = "m"
    ColorFieldSizeL ColorFieldSize = "l"
    ColorFieldSizeXl ColorFieldSize = "xl"
)

// spectrumColorField represents an sp-color-field component
type spectrumColorField struct {
    app.Compo
    *styler[*spectrumColorField]

    // Properties
    // What form of assistance should be provided when attempting to supply a value to the form control
    PropAutocomplete ColorFieldAutocomplete
    // Disables the color field. It will not receive focus or events
    PropDisabled bool
    // Whether a multiline form control will change size vertically to accommodate longer input
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
    // Size of the color field
    PropSize ColorFieldSize
    // The tab index to apply to this control
    PropTabindex float64
    // Whether the value held by the form control is valid
    PropValid bool
    // The color value held by the form control
    PropValue string
    // Whether to display a color handle showing the current color
    PropViewcolor bool




    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// ColorField creates a new sp-color-field component
func ColorField() *spectrumColorField {
    element := &spectrumColorField{
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
        PropSize: ColorFieldSizeM,
        PropTabindex: 0,
        PropValid: false,
        PropViewcolor: false,
    }

    element.styler = newStyler(element)

    return element
}

// What form of assistance should be provided when attempting to supply a value to the form control
func (c *spectrumColorField) Autocomplete(autocomplete ColorFieldAutocomplete) *spectrumColorField {
    c.PropAutocomplete = autocomplete
    return c
}

func (c *spectrumColorField) AutocompleteList() *spectrumColorField {
    return c.Autocomplete(ColorFieldAutocompleteList)
}
func (c *spectrumColorField) AutocompleteNone() *spectrumColorField {
    return c.Autocomplete(ColorFieldAutocompleteNone)
}
// Disables the color field. It will not receive focus or events
func (c *spectrumColorField) Disabled(disabled bool) *spectrumColorField {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumColorField) SetDisabled() *spectrumColorField {
    return c.Disabled(true)
}

// Whether a multiline form control will change size vertically to accommodate longer input
func (c *spectrumColorField) Grows(grows bool) *spectrumColorField {
    c.PropGrows = grows
    return c
}

func (c *spectrumColorField) SetGrows() *spectrumColorField {
    return c.Grows(true)
}

// Whether the value held by the form control is invalid
func (c *spectrumColorField) Invalid(invalid bool) *spectrumColorField {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumColorField) SetInvalid() *spectrumColorField {
    return c.Invalid(true)
}

// A string applied via aria-label to the form control when a user visible label is not provided
func (c *spectrumColorField) Label(label string) *spectrumColorField {
    c.PropLabel = label
    return c
}

// Defines the maximum string length that the user can enter
func (c *spectrumColorField) Maxlength(maxlength float64) *spectrumColorField {
    c.PropMaxlength = maxlength
    return c
}

// Defines the minimum string length that the user can enter
func (c *spectrumColorField) Minlength(minlength float64) *spectrumColorField {
    c.PropMinlength = minlength
    return c
}

// Whether the form control should accept a value longer than one line
func (c *spectrumColorField) Multiline(multiline bool) *spectrumColorField {
    c.PropMultiline = multiline
    return c
}

func (c *spectrumColorField) SetMultiline() *spectrumColorField {
    return c.Multiline(true)
}

// Name of the form control
func (c *spectrumColorField) Name(name string) *spectrumColorField {
    c.PropName = name
    return c
}

// Pattern the value must match to be valid
func (c *spectrumColorField) Pattern(pattern string) *spectrumColorField {
    c.PropPattern = pattern
    return c
}

// Text that appears in the form control when it has no value set
func (c *spectrumColorField) Placeholder(placeholder string) *spectrumColorField {
    c.PropPlaceholder = placeholder
    return c
}

// Whether to display the form control with no visible background
func (c *spectrumColorField) Quiet(quiet bool) *spectrumColorField {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumColorField) SetQuiet() *spectrumColorField {
    return c.Quiet(true)
}

// Whether a user can interact with the value of the form control
func (c *spectrumColorField) Readonly(readonly bool) *spectrumColorField {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumColorField) SetReadonly() *spectrumColorField {
    return c.Readonly(true)
}

// Whether the form control will be found to be invalid when it holds no value
func (c *spectrumColorField) Required(required bool) *spectrumColorField {
    c.PropRequired = required
    return c
}

func (c *spectrumColorField) SetRequired() *spectrumColorField {
    return c.Required(true)
}

// Size of the color field
func (c *spectrumColorField) Size(size ColorFieldSize) *spectrumColorField {
    c.PropSize = size
    return c
}

func (c *spectrumColorField) SizeS() *spectrumColorField {
    return c.Size(ColorFieldSizeS)
}
func (c *spectrumColorField) SizeM() *spectrumColorField {
    return c.Size(ColorFieldSizeM)
}
func (c *spectrumColorField) SizeL() *spectrumColorField {
    return c.Size(ColorFieldSizeL)
}
func (c *spectrumColorField) SizeXl() *spectrumColorField {
    return c.Size(ColorFieldSizeXl)
}
// The tab index to apply to this control
func (c *spectrumColorField) Tabindex(tabIndex float64) *spectrumColorField {
    c.PropTabindex = tabIndex
    return c
}

// Whether the value held by the form control is valid
func (c *spectrumColorField) Valid(valid bool) *spectrumColorField {
    c.PropValid = valid
    return c
}

func (c *spectrumColorField) SetValid() *spectrumColorField {
    return c.Valid(true)
}

// The color value held by the form control
func (c *spectrumColorField) Value(value string) *spectrumColorField {
    c.PropValue = value
    return c
}

// Whether to display a color handle showing the current color
func (c *spectrumColorField) Viewcolor(viewColor bool) *spectrumColorField {
    c.PropViewcolor = viewColor
    return c
}

func (c *spectrumColorField) SetViewcolor() *spectrumColorField {
    return c.Viewcolor(true)
}





// An alteration to the value of the color field has been committed by the user
func (c *spectrumColorField) OnChange(handler app.EventHandler) *spectrumColorField {
    c.PropOnChange = handler

    return c
}

// The value of the color field has changed
func (c *spectrumColorField) OnInput(handler app.EventHandler) *spectrumColorField {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-color-field component
func (c *spectrumColorField) Render() app.UI {
    element := app.Elem("sp-color-field")

    // Set attributes
    if c.PropAutocomplete != "" {
        element = element.Attr("autocomplete", string(c.PropAutocomplete))
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
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropValid {
        element = element.Attr("valid", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropViewcolor {
        element = element.Attr("viewColor", true)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnInput != nil {
        element = element.On("input", c.PropOnInput)
    }


    element = element.Styles(c.styler.styles)

    return element
} 