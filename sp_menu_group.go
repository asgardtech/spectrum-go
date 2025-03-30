package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MenuGroupSelects represents the How the menu group allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
type MenuGroupSelects string

// MenuGroupSelects values
const (
    MenuGroupSelectsInherit MenuGroupSelects = "inherit"
    MenuGroupSelectsSingle MenuGroupSelects = "single"
    MenuGroupSelectsMultiple MenuGroupSelects = "multiple"
)

// spectrumMenuGroup represents an sp-menu-group component
type spectrumMenuGroup struct {
    app.Compo
    *styler[*spectrumMenuGroup]

    // Properties
    // Whether menu group should be ignored by roving tabindex controller
    PropIgnore bool
    // Label of the menu group for accessibility
    PropLabel string
    // How the menu group allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
    PropSelects MenuGroupSelects
    // Value of the selected item(s)
    PropValue string
    // Separator for multiple values when selects is multiple
    PropValueSeparator string


    // Content slots
    PropHeaderSlot app.UI

    // Child components
    PropChildren []app.UI

    // Event handlers
    PropOnChange app.EventHandler
    PropOnClose app.EventHandler
}

// MenuGroup creates a new sp-menu-group component
func MenuGroup() *spectrumMenuGroup {
    element := &spectrumMenuGroup{
        PropIgnore: false,
        PropLabel: "",
        PropValue: "",
        PropValueSeparator: ",",
    }

    element.styler = newStyler(element)

    return element
}

// Whether menu group should be ignored by roving tabindex controller
func (c *spectrumMenuGroup) Ignore(ignore bool) *spectrumMenuGroup {
    c.PropIgnore = ignore
    return c
}

func (c *spectrumMenuGroup) SetIgnore() *spectrumMenuGroup {
    return c.Ignore(true)
}

// Label of the menu group for accessibility
func (c *spectrumMenuGroup) Label(label string) *spectrumMenuGroup {
    c.PropLabel = label
    return c
}

// How the menu group allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
func (c *spectrumMenuGroup) Selects(selects MenuGroupSelects) *spectrumMenuGroup {
    c.PropSelects = selects
    return c
}

func (c *spectrumMenuGroup) SelectsInherit() *spectrumMenuGroup {
    return c.Selects(MenuGroupSelectsInherit)
}
func (c *spectrumMenuGroup) SelectsSingle() *spectrumMenuGroup {
    return c.Selects(MenuGroupSelectsSingle)
}
func (c *spectrumMenuGroup) SelectsMultiple() *spectrumMenuGroup {
    return c.Selects(MenuGroupSelectsMultiple)
}
// Value of the selected item(s)
func (c *spectrumMenuGroup) Value(value string) *spectrumMenuGroup {
    c.PropValue = value
    return c
}

// Separator for multiple values when selects is multiple
func (c *spectrumMenuGroup) ValueSeparator(valueSeparator string) *spectrumMenuGroup {
    c.PropValueSeparator = valueSeparator
    return c
}



// Headline of the menu group
func (c *spectrumMenuGroup) Header(content app.UI) *spectrumMenuGroup {
    c.PropHeaderSlot = content

    return c
}


// Children sets the child components
func (c *spectrumMenuGroup) Children(children ...app.UI) *spectrumMenuGroup {
    c.PropChildren = children

    return c
}

// AddChild adds a child component
func (c *spectrumMenuGroup) AddChild(child app.UI) *spectrumMenuGroup {
    c.PropChildren = append(c.PropChildren, child)

    return c
}


// Announces that the value of the element has changed
func (c *spectrumMenuGroup) OnChange(handler app.EventHandler) *spectrumMenuGroup {
    c.PropOnChange = handler

    return c
}

// Fired when the menu group should be closed
func (c *spectrumMenuGroup) OnClose(handler app.EventHandler) *spectrumMenuGroup {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-menu-group component
func (c *spectrumMenuGroup) Render() app.UI {
    element := app.Elem("sp-menu-group")

    // Set attributes
    if c.PropIgnore {
        element = element.Attr("ignore", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropSelects != "" {
        element = element.Attr("selects", string(c.PropSelects))
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropValueSeparator != "" {
        element = element.Attr("value-separator", c.PropValueSeparator)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnClose != nil {
        element = element.On("close", c.PropOnClose)
    }

    // Add slots and children
    slotElements := []app.UI{}


    // Add header slot
    if c.PropHeaderSlot != nil {
        slotElem := c.PropHeaderSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("header")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "header").
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