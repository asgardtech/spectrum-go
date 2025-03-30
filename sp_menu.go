package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MenuSelects represents the How the menu allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
type MenuSelects string

// MenuSelects values
const (
    MenuSelectsInherit MenuSelects = "inherit"
    MenuSelectsSingle MenuSelects = "single"
    MenuSelectsMultiple MenuSelects = "multiple"
)
// MenuSize represents the The size of the menu
type MenuSize string

// MenuSize values
const (
    MenuSizeS MenuSize = "s"
    MenuSizeM MenuSize = "m"
    MenuSizeL MenuSize = "l"
    MenuSizeXl MenuSize = "xl"
)

// spectrumMenu represents an sp-menu component
type spectrumMenu struct {
    app.Compo
    *styler[*spectrumMenu]

    // Properties
    // Whether menu should be ignored by roving tabindex controller
    PropIgnore bool
    // Label of the menu for accessibility
    PropLabel string
    // How the menu allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
    PropSelects MenuSelects
    // The size of the menu
    PropSize MenuSize
    // Value of the selected item(s)
    PropValue string
    // Separator for multiple values when selects is multiple
    PropValueSeparator string


    // Content slots

    // Child components
    PropChildren []app.UI

    // Event handlers
    PropOnChange app.EventHandler
    PropOnClose app.EventHandler
}

// Menu creates a new sp-menu component
func Menu() *spectrumMenu {
    element := &spectrumMenu{
        PropIgnore: false,
        PropLabel: "",
        PropSize: MenuSizeM,
        PropValue: "",
        PropValueSeparator: ",",
    }

    element.styler = newStyler(element)

    return element
}

// Whether menu should be ignored by roving tabindex controller
func (c *spectrumMenu) Ignore(ignore bool) *spectrumMenu {
    c.PropIgnore = ignore
    return c
}

func (c *spectrumMenu) SetIgnore() *spectrumMenu {
    return c.Ignore(true)
}

// Label of the menu for accessibility
func (c *spectrumMenu) Label(label string) *spectrumMenu {
    c.PropLabel = label
    return c
}

// How the menu allows selection of its items: - undefined (default): no selection is allowed - inherit: the selection behavior is managed from an ancestor - single: only one item can be selected at a time - multiple: multiple items can be selected
func (c *spectrumMenu) Selects(selects MenuSelects) *spectrumMenu {
    c.PropSelects = selects
    return c
}

func (c *spectrumMenu) SelectsInherit() *spectrumMenu {
    return c.Selects(MenuSelectsInherit)
}
func (c *spectrumMenu) SelectsSingle() *spectrumMenu {
    return c.Selects(MenuSelectsSingle)
}
func (c *spectrumMenu) SelectsMultiple() *spectrumMenu {
    return c.Selects(MenuSelectsMultiple)
}
// The size of the menu
func (c *spectrumMenu) Size(size MenuSize) *spectrumMenu {
    c.PropSize = size
    return c
}

func (c *spectrumMenu) SizeS() *spectrumMenu {
    return c.Size(MenuSizeS)
}
func (c *spectrumMenu) SizeM() *spectrumMenu {
    return c.Size(MenuSizeM)
}
func (c *spectrumMenu) SizeL() *spectrumMenu {
    return c.Size(MenuSizeL)
}
func (c *spectrumMenu) SizeXl() *spectrumMenu {
    return c.Size(MenuSizeXl)
}
// Value of the selected item(s)
func (c *spectrumMenu) Value(value string) *spectrumMenu {
    c.PropValue = value
    return c
}

// Separator for multiple values when selects is multiple
func (c *spectrumMenu) ValueSeparator(valueSeparator string) *spectrumMenu {
    c.PropValueSeparator = valueSeparator
    return c
}




// Children sets the child components
func (c *spectrumMenu) Children(children ...app.UI) *spectrumMenu {
    c.PropChildren = children

    return c
}

// AddChild adds a child component
func (c *spectrumMenu) AddChild(child app.UI) *spectrumMenu {
    c.PropChildren = append(c.PropChildren, child)

    return c
}


// Announces that the value of the element has changed
func (c *spectrumMenu) OnChange(handler app.EventHandler) *spectrumMenu {
    c.PropOnChange = handler

    return c
}

// Fired when the menu should be closed, typically after item selection
func (c *spectrumMenu) OnClose(handler app.EventHandler) *spectrumMenu {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-menu component
func (c *spectrumMenu) Render() app.UI {
    element := app.Elem("sp-menu")

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
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
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