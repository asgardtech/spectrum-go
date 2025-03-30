package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumBreadcrumbs represents an sp-breadcrumbs component
type spectrumBreadcrumbs struct {
    app.Compo
    *styler[*spectrumBreadcrumbs]

    // Properties
    // Compact option is useful for reducing the height of the breadcrumbs
    PropCompact bool
    // Accessible name for the Breadcrumbs component
    PropLabel string
    // Override the maximum number of visible items
    PropMaxVisibleItems float64
    // Change the default label of the action menu
    PropMenuLabel string

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI
    PropRootSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
}

// Breadcrumbs creates a new sp-breadcrumbs component
func Breadcrumbs() *spectrumBreadcrumbs {
    element := &spectrumBreadcrumbs{
        PropCompact: false,
        PropLabel: "",
        PropMaxVisibleItems: 4,
        PropMenuLabel: "More items",
    }

    element.styler = newStyler(element)

    return element
}

// Compact option is useful for reducing the height of the breadcrumbs
func (c *spectrumBreadcrumbs) Compact(compact bool) *spectrumBreadcrumbs {
    c.PropCompact = compact
    return c
}

func (c *spectrumBreadcrumbs) SetCompact() *spectrumBreadcrumbs {
    return c.Compact(true)
}

// Accessible name for the Breadcrumbs component
func (c *spectrumBreadcrumbs) Label(label string) *spectrumBreadcrumbs {
    c.PropLabel = label
    return c
}

// Override the maximum number of visible items
func (c *spectrumBreadcrumbs) MaxVisibleItems(maxVisibleItems float64) *spectrumBreadcrumbs {
    c.PropMaxVisibleItems = maxVisibleItems
    return c
}

// Change the default label of the action menu
func (c *spectrumBreadcrumbs) MenuLabel(menuLabel string) *spectrumBreadcrumbs {
    c.PropMenuLabel = menuLabel
    return c
}


// Text sets the text content for the default slot
func (c *spectrumBreadcrumbs) Text(text string) *spectrumBreadcrumbs {
    c.PropText = text
    return c
}


// Change the default icon of the action menu
func (c *spectrumBreadcrumbs) Icon(content app.UI) *spectrumBreadcrumbs {
    c.PropIconSlot = content

    return c
}

// Breadcrumb item to always display
func (c *spectrumBreadcrumbs) Root(content app.UI) *spectrumBreadcrumbs {
    c.PropRootSlot = content

    return c
}



// Announces the selected breadcrumb item
func (c *spectrumBreadcrumbs) OnChange(handler app.EventHandler) *spectrumBreadcrumbs {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-breadcrumbs component
func (c *spectrumBreadcrumbs) Render() app.UI {
    element := app.Elem("sp-breadcrumbs")

    // Set attributes
    if c.PropCompact {
        element = element.Attr("compact", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMaxVisibleItems != 0 {
        element = element.Attr("max-visible-items", c.PropMaxVisibleItems)
    }
    if c.PropMenuLabel != "" {
        element = element.Attr("menu-label", c.PropMenuLabel)
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
    // Add root slot
    if c.PropRootSlot != nil {
        slotElem := c.PropRootSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("root")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "root").
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