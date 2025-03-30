package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumButtonGroup represents an sp-button-group component
type spectrumButtonGroup struct {
    app.Compo
    *styler[*spectrumButtonGroup]

    // Properties
    // Use vertical orientation when horizontal space is limited
    PropVertical bool

    // Text content for the default slot
    PropText string

    // Content slots


}

// ButtonGroup creates a new sp-button-group component
func ButtonGroup() *spectrumButtonGroup {
    element := &spectrumButtonGroup{
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Use vertical orientation when horizontal space is limited
func (c *spectrumButtonGroup) Vertical(vertical bool) *spectrumButtonGroup {
    c.PropVertical = vertical
    return c
}

func (c *spectrumButtonGroup) SetVertical() *spectrumButtonGroup {
    return c.Vertical(true)
}


// Text sets the text content for the default slot
func (c *spectrumButtonGroup) Text(text string) *spectrumButtonGroup {
    c.PropText = text
    return c
}





// Render renders the sp-button-group component
func (c *spectrumButtonGroup) Render() app.UI {
    element := app.Elem("sp-button-group")

    // Set attributes
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 