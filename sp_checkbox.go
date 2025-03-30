package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// CheckboxSize represents the The size of the checkbox
type CheckboxSize string

// CheckboxSize values
const (
    CheckboxSizeS CheckboxSize = "s"
    CheckboxSizeM CheckboxSize = "m"
    CheckboxSizeL CheckboxSize = "l"
    CheckboxSizeXl CheckboxSize = "xl"
)

// spectrumCheckbox represents an sp-checkbox component
type spectrumCheckbox struct {
    app.Compo
    *styler[*spectrumCheckbox]

    // Properties
    // Whether the checkbox is checked
    PropChecked bool
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether to use the emphasized (blue) visual style for the checkbox
    PropEmphasized bool
    // Whether the checkbox is in an indeterminate state
    PropIndeterminate bool
    // Whether the checkbox is in an invalid state
    PropInvalid bool
    // The name of the checkbox when used in a form
    PropName string
    // Whether the checkbox is readonly
    PropReadonly bool
    // The size of the checkbox
    PropSize CheckboxSize
    // The tab index of the checkbox
    PropTabindex float64

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// Checkbox creates a new sp-checkbox component
func Checkbox() *spectrumCheckbox {
    element := &spectrumCheckbox{
        PropChecked: false,
        PropDisabled: false,
        PropEmphasized: false,
        PropIndeterminate: false,
        PropInvalid: false,
        PropReadonly: false,
        PropSize: CheckboxSizeM,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the checkbox is checked
func (c *spectrumCheckbox) Checked(checked bool) *spectrumCheckbox {
    c.PropChecked = checked
    return c
}

func (c *spectrumCheckbox) SetChecked() *spectrumCheckbox {
    return c.Checked(true)
}

// Disable this control. It will not receive focus or events
func (c *spectrumCheckbox) Disabled(disabled bool) *spectrumCheckbox {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumCheckbox) SetDisabled() *spectrumCheckbox {
    return c.Disabled(true)
}

// Whether to use the emphasized (blue) visual style for the checkbox
func (c *spectrumCheckbox) Emphasized(emphasized bool) *spectrumCheckbox {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumCheckbox) SetEmphasized() *spectrumCheckbox {
    return c.Emphasized(true)
}

// Whether the checkbox is in an indeterminate state
func (c *spectrumCheckbox) Indeterminate(indeterminate bool) *spectrumCheckbox {
    c.PropIndeterminate = indeterminate
    return c
}

func (c *spectrumCheckbox) SetIndeterminate() *spectrumCheckbox {
    return c.Indeterminate(true)
}

// Whether the checkbox is in an invalid state
func (c *spectrumCheckbox) Invalid(invalid bool) *spectrumCheckbox {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumCheckbox) SetInvalid() *spectrumCheckbox {
    return c.Invalid(true)
}

// The name of the checkbox when used in a form
func (c *spectrumCheckbox) Name(name string) *spectrumCheckbox {
    c.PropName = name
    return c
}

// Whether the checkbox is readonly
func (c *spectrumCheckbox) Readonly(readonly bool) *spectrumCheckbox {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumCheckbox) SetReadonly() *spectrumCheckbox {
    return c.Readonly(true)
}

// The size of the checkbox
func (c *spectrumCheckbox) Size(size CheckboxSize) *spectrumCheckbox {
    c.PropSize = size
    return c
}

func (c *spectrumCheckbox) SizeS() *spectrumCheckbox {
    return c.Size(CheckboxSizeS)
}
func (c *spectrumCheckbox) SizeM() *spectrumCheckbox {
    return c.Size(CheckboxSizeM)
}
func (c *spectrumCheckbox) SizeL() *spectrumCheckbox {
    return c.Size(CheckboxSizeL)
}
func (c *spectrumCheckbox) SizeXl() *spectrumCheckbox {
    return c.Size(CheckboxSizeXl)
}
// The tab index of the checkbox
func (c *spectrumCheckbox) Tabindex(tabindex float64) *spectrumCheckbox {
    c.PropTabindex = tabindex
    return c
}


// Text sets the text content for the default slot
func (c *spectrumCheckbox) Text(text string) *spectrumCheckbox {
    c.PropText = text
    return c
}




// Announces a change in the checked property of a Checkbox
func (c *spectrumCheckbox) OnChange(handler app.EventHandler) *spectrumCheckbox {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-checkbox component
func (c *spectrumCheckbox) Render() app.UI {
    element := app.Elem("sp-checkbox")

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
    if c.PropIndeterminate {
        element = element.Attr("indeterminate", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
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