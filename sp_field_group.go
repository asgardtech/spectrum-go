package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumFieldGroup represents an sp-field-group component
type spectrumFieldGroup struct {
    app.Compo
    *styler[*spectrumFieldGroup]

    // Properties
    // Whether the fields should be arranged horizontally
    PropHorizontal bool
    // Whether the field group is in an invalid state
    PropInvalid bool
    // An accessible label for the field group
    PropLabel string
    // Whether the fields should be arranged vertically
    PropVertical bool

    // Text content for the default slot
    PropText string

    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI


}

// FieldGroup creates a new sp-field-group component
func FieldGroup() *spectrumFieldGroup {
    element := &spectrumFieldGroup{
        PropHorizontal: false,
        PropInvalid: false,
        PropLabel: "",
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the fields should be arranged horizontally
func (c *spectrumFieldGroup) Horizontal(horizontal bool) *spectrumFieldGroup {
    c.PropHorizontal = horizontal
    return c
}

func (c *spectrumFieldGroup) SetHorizontal() *spectrumFieldGroup {
    return c.Horizontal(true)
}

// Whether the field group is in an invalid state
func (c *spectrumFieldGroup) Invalid(invalid bool) *spectrumFieldGroup {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumFieldGroup) SetInvalid() *spectrumFieldGroup {
    return c.Invalid(true)
}

// An accessible label for the field group
func (c *spectrumFieldGroup) Label(label string) *spectrumFieldGroup {
    c.PropLabel = label
    return c
}

// Whether the fields should be arranged vertically
func (c *spectrumFieldGroup) Vertical(vertical bool) *spectrumFieldGroup {
    c.PropVertical = vertical
    return c
}

func (c *spectrumFieldGroup) SetVertical() *spectrumFieldGroup {
    return c.Vertical(true)
}


// Text sets the text content for the default slot
func (c *spectrumFieldGroup) Text(text string) *spectrumFieldGroup {
    c.PropText = text
    return c
}


// Default or non-negative help text to associate to your form element
func (c *spectrumFieldGroup) HelpText(content app.UI) *spectrumFieldGroup {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when invalid
func (c *spectrumFieldGroup) NegativeHelpText(content app.UI) *spectrumFieldGroup {
    c.PropNegativeHelpTextSlot = content

    return c
}




// Render renders the sp-field-group component
func (c *spectrumFieldGroup) Render() app.UI {
    element := app.Elem("sp-field-group")

    // Set attributes
    if c.PropHorizontal {
        element = element.Attr("horizontal", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

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