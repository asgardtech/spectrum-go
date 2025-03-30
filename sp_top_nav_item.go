package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTopNavItem represents an sp-top-nav-item component
type spectrumTopNavItem struct {
    app.Compo
    *styler[*spectrumTopNavItem]

    // Properties
    // URL that the navigation item points to
    PropHref string
    // Accessible label for the navigation item
    PropLabel string
    // Whether the navigation item is currently selected
    PropSelected bool
    // Specifies where to open the linked document
    PropTarget string

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnClick app.EventHandler
}

// TopNavItem creates a new sp-top-nav-item component
func TopNavItem() *spectrumTopNavItem {
    element := &spectrumTopNavItem{
        PropLabel: "",
        PropSelected: false,
    }

    element.styler = newStyler(element)

    return element
}

// URL that the navigation item points to
func (c *spectrumTopNavItem) Href(href string) *spectrumTopNavItem {
    c.PropHref = href
    return c
}

// Accessible label for the navigation item
func (c *spectrumTopNavItem) Label(label string) *spectrumTopNavItem {
    c.PropLabel = label
    return c
}

// Whether the navigation item is currently selected
func (c *spectrumTopNavItem) Selected(selected bool) *spectrumTopNavItem {
    c.PropSelected = selected
    return c
}

func (c *spectrumTopNavItem) SetSelected() *spectrumTopNavItem {
    return c.Selected(true)
}

// Specifies where to open the linked document
func (c *spectrumTopNavItem) Target(target string) *spectrumTopNavItem {
    c.PropTarget = target
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTopNavItem) Text(text string) *spectrumTopNavItem {
    c.PropText = text
    return c
}


// Icon to display with the navigation item
func (c *spectrumTopNavItem) Icon(content app.UI) *spectrumTopNavItem {
    c.PropIconSlot = content

    return c
}



// Triggered when the navigation item is clicked
func (c *spectrumTopNavItem) OnClick(handler app.EventHandler) *spectrumTopNavItem {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-top-nav-item component
func (c *spectrumTopNavItem) Render() app.UI {
    element := app.Elem("sp-top-nav-item")

    // Set attributes
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", c.PropTarget)
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


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 