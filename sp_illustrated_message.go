package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumIllustratedMessage represents an sp-illustrated-message component
type spectrumIllustratedMessage struct {
    app.Compo
    *styler[*spectrumIllustratedMessage]

    // Properties
    // Additional descriptive text explaining the illustrated message
    PropDescription string
    // Main heading or title for the illustrated message
    PropHeading string

    // Text content for the default slot
    PropText string

    // Content slots
    PropDescriptionSlot app.UI
    PropHeadingSlot app.UI


}

// IllustratedMessage creates a new sp-illustrated-message component
func IllustratedMessage() *spectrumIllustratedMessage {
    element := &spectrumIllustratedMessage{
        PropDescription: "",
        PropHeading: "",
    }

    element.styler = newStyler(element)

    return element
}

// Additional descriptive text explaining the illustrated message
func (c *spectrumIllustratedMessage) Description(description string) *spectrumIllustratedMessage {
    c.PropDescription = description
    return c
}

// Main heading or title for the illustrated message
func (c *spectrumIllustratedMessage) Heading(heading string) *spectrumIllustratedMessage {
    c.PropHeading = heading
    return c
}


// Text sets the text content for the default slot
func (c *spectrumIllustratedMessage) Text(text string) *spectrumIllustratedMessage {
    c.PropText = text
    return c
}


// Description text for the illustration
func (c *spectrumIllustratedMessage) DescriptionContent(content app.UI) *spectrumIllustratedMessage {
    c.PropDescriptionSlot = content

    return c
}

// Headline for the message
func (c *spectrumIllustratedMessage) HeadingContent(content app.UI) *spectrumIllustratedMessage {
    c.PropHeadingSlot = content

    return c
}




// Render renders the sp-illustrated-message component
func (c *spectrumIllustratedMessage) Render() app.UI {
    element := app.Elem("sp-illustrated-message")

    // Set attributes
    if c.PropDescription != "" {
        element = element.Attr("description", c.PropDescription)
    }
    if c.PropHeading != "" {
        element = element.Attr("heading", c.PropHeading)
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add description slot
    if c.PropDescriptionSlot != nil {
        slotElem := c.PropDescriptionSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("description")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "description").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add heading slot
    if c.PropHeadingSlot != nil {
        slotElem := c.PropHeadingSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("heading")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "heading").
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