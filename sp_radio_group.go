package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumRadioGroup represents an sp-radio-group component
type spectrumRadioGroup struct {
    app.Compo
    *styler[*spectrumRadioGroup]

    // Properties
    // Whether the radio buttons should be arranged horizontally
    PropHorizontal bool
    // Whether the radio group is in an invalid state
    PropInvalid bool
    // The accessible label for the radio group
    PropLabel string
    // The name of the radio group when used in a form
    PropName string
    // The value of the currently selected radio button
    PropSelected string
    // Whether the radio buttons should be arranged vertically
    PropVertical bool


    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI

    // Child components
    PropChildren []app.UI

    // Event handlers
    PropOnChange app.EventHandler
}

// RadioGroup creates a new sp-radio-group component
func RadioGroup() *spectrumRadioGroup {
    element := &spectrumRadioGroup{
        PropHorizontal: false,
        PropInvalid: false,
        PropLabel: "",
        PropName: "",
        PropSelected: "",
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the radio buttons should be arranged horizontally
func (c *spectrumRadioGroup) Horizontal(horizontal bool) *spectrumRadioGroup {
    c.PropHorizontal = horizontal
    return c
}

func (c *spectrumRadioGroup) SetHorizontal() *spectrumRadioGroup {
    return c.Horizontal(true)
}

// Whether the radio group is in an invalid state
func (c *spectrumRadioGroup) Invalid(invalid bool) *spectrumRadioGroup {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumRadioGroup) SetInvalid() *spectrumRadioGroup {
    return c.Invalid(true)
}

// The accessible label for the radio group
func (c *spectrumRadioGroup) Label(label string) *spectrumRadioGroup {
    c.PropLabel = label
    return c
}

// The name of the radio group when used in a form
func (c *spectrumRadioGroup) Name(name string) *spectrumRadioGroup {
    c.PropName = name
    return c
}

// The value of the currently selected radio button
func (c *spectrumRadioGroup) Selected(selected string) *spectrumRadioGroup {
    c.PropSelected = selected
    return c
}

// Whether the radio buttons should be arranged vertically
func (c *spectrumRadioGroup) Vertical(vertical bool) *spectrumRadioGroup {
    c.PropVertical = vertical
    return c
}

func (c *spectrumRadioGroup) SetVertical() *spectrumRadioGroup {
    return c.Vertical(true)
}



// Default or non-negative help text to associate to your form element
func (c *spectrumRadioGroup) HelpText(content app.UI) *spectrumRadioGroup {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when invalid
func (c *spectrumRadioGroup) NegativeHelpText(content app.UI) *spectrumRadioGroup {
    c.PropNegativeHelpTextSlot = content

    return c
}


// Children sets the child components
func (c *spectrumRadioGroup) Children(children ...app.UI) *spectrumRadioGroup {
    c.PropChildren = children

    return c
}

// AddChild adds a child component
func (c *spectrumRadioGroup) AddChild(child app.UI) *spectrumRadioGroup {
    c.PropChildren = append(c.PropChildren, child)

    return c
}


// An alteration to the value of the element has been committed by the user.
func (c *spectrumRadioGroup) OnChange(handler app.EventHandler) *spectrumRadioGroup {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-radio-group component
func (c *spectrumRadioGroup) Render() app.UI {
    element := app.Elem("sp-radio-group")

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
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropSelected != "" {
        element = element.Attr("selected", c.PropSelected)
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
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

    // Add children
    if len(c.PropChildren) > 0 {
        for _, child := range c.PropChildren {
            slotElements = append(slotElements, child)
        }
    }

    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 