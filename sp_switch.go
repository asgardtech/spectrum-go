package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwitchSize represents the The size of the switch
type SwitchSize string

// SwitchSize values
const (
    SwitchSizeS SwitchSize = "s"
    SwitchSizeM SwitchSize = "m"
    SwitchSizeL SwitchSize = "l"
    SwitchSizeXl SwitchSize = "xl"
)

// spectrumSwitch represents an sp-switch component
type spectrumSwitch struct {
    app.Compo
    *styler[*spectrumSwitch]

    // Properties
    // Whether the switch is checked (on)
    PropChecked bool
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether to use the emphasized (blue) visual style
    PropEmphasized bool
    // The name of the switch when used in a form
    PropName string
    // Whether the switch is readonly
    PropReadonly bool
    // The size of the switch
    PropSize SwitchSize
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// Switch creates a new sp-switch component
func Switch() *spectrumSwitch {
    element := &spectrumSwitch{
        PropChecked: false,
        PropDisabled: false,
        PropEmphasized: false,
        PropReadonly: false,
        PropSize: SwitchSizeM,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the switch is checked (on)
func (c *spectrumSwitch) Checked(checked bool) *spectrumSwitch {
    c.PropChecked = checked
    return c
}

func (c *spectrumSwitch) SetChecked() *spectrumSwitch {
    return c.Checked(true)
}

// Disable this control. It will not receive focus or events
func (c *spectrumSwitch) Disabled(disabled bool) *spectrumSwitch {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSwitch) SetDisabled() *spectrumSwitch {
    return c.Disabled(true)
}

// Whether to use the emphasized (blue) visual style
func (c *spectrumSwitch) Emphasized(emphasized bool) *spectrumSwitch {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumSwitch) SetEmphasized() *spectrumSwitch {
    return c.Emphasized(true)
}

// The name of the switch when used in a form
func (c *spectrumSwitch) Name(name string) *spectrumSwitch {
    c.PropName = name
    return c
}

// Whether the switch is readonly
func (c *spectrumSwitch) Readonly(readonly bool) *spectrumSwitch {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumSwitch) SetReadonly() *spectrumSwitch {
    return c.Readonly(true)
}

// The size of the switch
func (c *spectrumSwitch) Size(size SwitchSize) *spectrumSwitch {
    c.PropSize = size
    return c
}

func (c *spectrumSwitch) SizeS() *spectrumSwitch {
    return c.Size(SwitchSizeS)
}
func (c *spectrumSwitch) SizeM() *spectrumSwitch {
    return c.Size(SwitchSizeM)
}
func (c *spectrumSwitch) SizeL() *spectrumSwitch {
    return c.Size(SwitchSizeL)
}
func (c *spectrumSwitch) SizeXl() *spectrumSwitch {
    return c.Size(SwitchSizeXl)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumSwitch) Tabindex(tabindex float64) *spectrumSwitch {
    c.PropTabindex = tabindex
    return c
}


// Text sets the text content for the default slot
func (c *spectrumSwitch) Text(text string) *spectrumSwitch {
    c.PropText = text
    return c
}




// Announces a change in the checked property of a Switch
func (c *spectrumSwitch) OnChange(handler app.EventHandler) *spectrumSwitch {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-switch component
func (c *spectrumSwitch) Render() app.UI {
    element := app.Elem("sp-switch")

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