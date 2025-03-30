package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// RadioSize represents the The size of the radio button
type RadioSize string

// RadioSize values
const (
    RadioSizeS RadioSize = "s"
    RadioSizeM RadioSize = "m"
    RadioSizeL RadioSize = "l"
    RadioSizeXl RadioSize = "xl"
)

// spectrumRadio represents an sp-radio component
type spectrumRadio struct {
    app.Compo
    *styler[*spectrumRadio]

    // Properties
    // Represents when the input is checked
    PropChecked bool
    // Uses the disabled style
    PropDisabled bool
    // Whether to use the emphasized (blue) visual style
    PropEmphasized bool
    // Uses the invalid style
    PropInvalid bool
    // Whether the radio is readonly
    PropReadonly bool
    // Identifies this radio button within its radio group
    PropValue string
    // The size of the radio button
    PropSize RadioSize
    // The name of the radio button when used in a form
    PropName string

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// Radio creates a new sp-radio component
func Radio() *spectrumRadio {
    element := &spectrumRadio{
        PropChecked: false,
        PropDisabled: false,
        PropEmphasized: false,
        PropInvalid: false,
        PropReadonly: false,
        PropValue: "",
        PropSize: RadioSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Represents when the input is checked
func (c *spectrumRadio) Checked(checked bool) *spectrumRadio {
    c.PropChecked = checked
    return c
}

func (c *spectrumRadio) SetChecked() *spectrumRadio {
    return c.Checked(true)
}

// Uses the disabled style
func (c *spectrumRadio) Disabled(disabled bool) *spectrumRadio {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumRadio) SetDisabled() *spectrumRadio {
    return c.Disabled(true)
}

// Whether to use the emphasized (blue) visual style
func (c *spectrumRadio) Emphasized(emphasized bool) *spectrumRadio {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumRadio) SetEmphasized() *spectrumRadio {
    return c.Emphasized(true)
}

// Uses the invalid style
func (c *spectrumRadio) Invalid(invalid bool) *spectrumRadio {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumRadio) SetInvalid() *spectrumRadio {
    return c.Invalid(true)
}

// Whether the radio is readonly
func (c *spectrumRadio) Readonly(readonly bool) *spectrumRadio {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumRadio) SetReadonly() *spectrumRadio {
    return c.Readonly(true)
}

// Identifies this radio button within its radio group
func (c *spectrumRadio) Value(value string) *spectrumRadio {
    c.PropValue = value
    return c
}

// The size of the radio button
func (c *spectrumRadio) Size(size RadioSize) *spectrumRadio {
    c.PropSize = size
    return c
}

func (c *spectrumRadio) SizeS() *spectrumRadio {
    return c.Size(RadioSizeS)
}
func (c *spectrumRadio) SizeM() *spectrumRadio {
    return c.Size(RadioSizeM)
}
func (c *spectrumRadio) SizeL() *spectrumRadio {
    return c.Size(RadioSizeL)
}
func (c *spectrumRadio) SizeXl() *spectrumRadio {
    return c.Size(RadioSizeXl)
}
// The name of the radio button when used in a form
func (c *spectrumRadio) Name(name string) *spectrumRadio {
    c.PropName = name
    return c
}


// Text sets the text content for the default slot
func (c *spectrumRadio) Text(text string) *spectrumRadio {
    c.PropText = text
    return c
}




// When the input is interacted with and its state is changed
func (c *spectrumRadio) OnChange(handler app.EventHandler) *spectrumRadio {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-radio component
func (c *spectrumRadio) Render() app.UI {
    element := app.Elem("sp-radio")

    // Set attributes
    if c.PropChecked {
        element = element.Attr("checked", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 