package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumSidenavOverflow represents an sp-sidenav-overflow component
type spectrumSidenavOverflow struct {
    app.Compo
    *styler[*spectrumSidenavOverflow]

    // Properties
    // The text label for the overflow menu
    PropLabel string

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


}

// SidenavOverflow creates a new sp-sidenav-overflow component
func SidenavOverflow() *spectrumSidenavOverflow {
    element := &spectrumSidenavOverflow{
        PropLabel: "More",
    }

    element.styler = newStyler(element)

    return element
}

// The text label for the overflow menu
func (c *spectrumSidenavOverflow) Label(label string) *spectrumSidenavOverflow {
    c.PropLabel = label
    return c
}


// Text sets the text content for the default slot
func (c *spectrumSidenavOverflow) Text(text string) *spectrumSidenavOverflow {
    c.PropText = text
    return c
}


// Icon to display for the overflow menu
func (c *spectrumSidenavOverflow) Icon(content app.UI) *spectrumSidenavOverflow {
    c.PropIconSlot = content

    return c
}




// Render renders the sp-sidenav-overflow component
func (c *spectrumSidenavOverflow) Render() app.UI {
    element := app.Elem("sp-sidenav-overflow")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
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