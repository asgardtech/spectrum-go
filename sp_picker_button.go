package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerButtonPosition represents the Position relative to textfield
type PickerButtonPosition string

// PickerButtonPosition values
const (
    PickerButtonPositionLeft PickerButtonPosition = "left"
    PickerButtonPositionRight PickerButtonPosition = "right"
)
// PickerButtonType represents the Button type
type PickerButtonType string

// PickerButtonType values
const (
    PickerButtonTypeButton PickerButtonType = "button"
    PickerButtonTypeSubmit PickerButtonType = "submit"
    PickerButtonTypeReset PickerButtonType = "reset"
)

// spectrumPickerButton represents an sp-picker-button component
type spectrumPickerButton struct {
    app.Compo
    *styler[*spectrumPickerButton]

    // Properties
    // Whether the button is in active state
    PropActive bool
    // Disable this control
    PropDisabled bool
    // Causes browser to treat linked URL as download
    PropDownload string
    // URL the button links to
    PropHref string
    // Accessible label for the component
    PropLabel string
    // Position relative to textfield
    PropPosition PickerButtonPosition
    // Less pronounced visual delivery
    PropQuiet bool
    // Deeply rounded corners in express system
    PropRounded bool
    // Tab index for the control
    PropTabindex float64
    // Button type
    PropType PickerButtonType

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI
    PropLabelSlot app.UI


    // Event handlers
    PropOnClick app.EventHandler
}

// PickerButton creates a new sp-picker-button component
func PickerButton() *spectrumPickerButton {
    element := &spectrumPickerButton{
        PropActive: false,
        PropDisabled: false,
        PropDownload: "",
        PropHref: "",
        PropLabel: "",
        PropPosition: PickerButtonPositionRight,
        PropQuiet: false,
        PropRounded: false,
        PropTabindex: 0,
        PropType: PickerButtonTypeButton,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the button is in active state
func (c *spectrumPickerButton) Active(active bool) *spectrumPickerButton {
    c.PropActive = active
    return c
}

func (c *spectrumPickerButton) SetActive() *spectrumPickerButton {
    return c.Active(true)
}

// Disable this control
func (c *spectrumPickerButton) Disabled(disabled bool) *spectrumPickerButton {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumPickerButton) SetDisabled() *spectrumPickerButton {
    return c.Disabled(true)
}

// Causes browser to treat linked URL as download
func (c *spectrumPickerButton) Download(download string) *spectrumPickerButton {
    c.PropDownload = download
    return c
}

// URL the button links to
func (c *spectrumPickerButton) Href(href string) *spectrumPickerButton {
    c.PropHref = href
    return c
}

// Accessible label for the component
func (c *spectrumPickerButton) Label(label string) *spectrumPickerButton {
    c.PropLabel = label
    return c
}

// Position relative to textfield
func (c *spectrumPickerButton) Position(position PickerButtonPosition) *spectrumPickerButton {
    c.PropPosition = position
    return c
}

func (c *spectrumPickerButton) PositionLeft() *spectrumPickerButton {
    return c.Position(PickerButtonPositionLeft)
}
func (c *spectrumPickerButton) PositionRight() *spectrumPickerButton {
    return c.Position(PickerButtonPositionRight)
}
// Less pronounced visual delivery
func (c *spectrumPickerButton) Quiet(quiet bool) *spectrumPickerButton {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumPickerButton) SetQuiet() *spectrumPickerButton {
    return c.Quiet(true)
}

// Deeply rounded corners in express system
func (c *spectrumPickerButton) Rounded(rounded bool) *spectrumPickerButton {
    c.PropRounded = rounded
    return c
}

func (c *spectrumPickerButton) SetRounded() *spectrumPickerButton {
    return c.Rounded(true)
}

// Tab index for the control
func (c *spectrumPickerButton) Tabindex(tabIndex float64) *spectrumPickerButton {
    c.PropTabindex = tabIndex
    return c
}

// Button type
func (c *spectrumPickerButton) Type(typeValue PickerButtonType) *spectrumPickerButton {
    c.PropType = typeValue
    return c
}

func (c *spectrumPickerButton) TypeButton() *spectrumPickerButton {
    return c.Type(PickerButtonTypeButton)
}
func (c *spectrumPickerButton) TypeSubmit() *spectrumPickerButton {
    return c.Type(PickerButtonTypeSubmit)
}
func (c *spectrumPickerButton) TypeReset() *spectrumPickerButton {
    return c.Type(PickerButtonTypeReset)
}

// Text sets the text content for the default slot
func (c *spectrumPickerButton) Text(text string) *spectrumPickerButton {
    c.PropText = text
    return c
}


// Icon element to display
func (c *spectrumPickerButton) Icon(content app.UI) *spectrumPickerButton {
    c.PropIconSlot = content

    return c
}

// Content for the button label
func (c *spectrumPickerButton) LabelContent(content app.UI) *spectrumPickerButton {
    c.PropLabelSlot = content

    return c
}



// Fired when the button is clicked
func (c *spectrumPickerButton) OnClick(handler app.EventHandler) *spectrumPickerButton {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-picker-button component
func (c *spectrumPickerButton) Render() app.UI {
    element := app.Elem("sp-picker-button")

    // Set attributes
    if c.PropActive {
        element = element.Attr("active", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropPosition != "" {
        element = element.Attr("position", string(c.PropPosition))
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropRounded {
        element = element.Attr("rounded", true)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
    }

    // Add event handlers
    if c.PropOnClick != nil {
        element = element.On("click", c.PropOnClick)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add icon slot
    if c.PropIconSlot != nil {
        slotElem := c.PropIconSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("icon")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "icon").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add label slot
    if c.PropLabelSlot != nil {
        slotElem := c.PropLabelSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("label")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "label").
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