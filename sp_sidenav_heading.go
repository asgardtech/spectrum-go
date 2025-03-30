package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumSidenavHeading represents an sp-sidenav-heading component
type spectrumSidenavHeading struct {
    app.Compo
    *styler[*spectrumSidenavHeading]

    // Properties
    // The text label for the heading
    PropLabel string

    // Text content for the default slot
    PropText string

    // Content slots


}

// SidenavHeading creates a new sp-sidenav-heading component
func SidenavHeading() *spectrumSidenavHeading {
    element := &spectrumSidenavHeading{
        PropLabel: "",
    }

    element.styler = newStyler(element)

    return element
}

// The text label for the heading
func (c *spectrumSidenavHeading) Label(label string) *spectrumSidenavHeading {
    c.PropLabel = label
    return c
}


// Text sets the text content for the default slot
func (c *spectrumSidenavHeading) Text(text string) *spectrumSidenavHeading {
    c.PropText = text
    return c
}





// Render renders the sp-sidenav-heading component
func (c *spectrumSidenavHeading) Render() app.UI {
    element := app.Elem("sp-sidenav-heading")

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



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 