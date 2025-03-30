package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTruncated represents an sp-truncated component
type spectrumTruncated struct {
    app.Compo
    *styler[*spectrumTruncated]

    // Properties
    // Where the tooltip should be positioned relative to the truncated content
    PropPlacement string

    // Text content for the default slot
    PropText string

    // Content slots
    PropOverflowSlot app.UI


}

// Truncated creates a new sp-truncated component
func Truncated() *spectrumTruncated {
    element := &spectrumTruncated{
        PropPlacement: "bottom",
    }

    element.styler = newStyler(element)

    return element
}

// Where the tooltip should be positioned relative to the truncated content
func (c *spectrumTruncated) Placement(placement string) *spectrumTruncated {
    c.PropPlacement = placement
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTruncated) Text(text string) *spectrumTruncated {
    c.PropText = text
    return c
}


// Custom content to display in the tooltip when truncated
func (c *spectrumTruncated) Overflow(content app.UI) *spectrumTruncated {
    c.PropOverflowSlot = content

    return c
}




// Render renders the sp-truncated component
func (c *spectrumTruncated) Render() app.UI {
    element := app.Elem("sp-truncated")

    // Set attributes
    if c.PropPlacement != "" {
        element = element.Attr("placement", c.PropPlacement)
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add overflow slot
    if c.PropOverflowSlot != nil {
        slotElem := c.PropOverflowSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("overflow")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "overflow").
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