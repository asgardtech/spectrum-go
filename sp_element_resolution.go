package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumElementResolution represents an  component
type spectrumElementResolution struct {
    app.Compo
    *styler[*spectrumElementResolution]

    // Properties
    // The CSS selector to query for matching elements
    PropSelector string
    // Reference to the first element matching the selector
    PropElement any




}

// ElementResolution creates a new  component
func ElementResolution() *spectrumElementResolution {
    element := &spectrumElementResolution{
        PropElement: map[string]interface{}{},
    }

    element.styler = newStyler(element)

    return element
}

// The CSS selector to query for matching elements
func (c *spectrumElementResolution) Selector(selector string) *spectrumElementResolution {
    c.PropSelector = selector
    return c
}

// Reference to the first element matching the selector
func (c *spectrumElementResolution) Element(element any) *spectrumElementResolution {
    c.PropElement = element
    return c
}






// Render renders the  component
func (c *spectrumElementResolution) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropSelector != "" {
        element = element.Attr("selector", c.PropSelector)
    }
    if c.PropElement != nil {
        element = element.Attr("element", c.PropElement)
    }



    element = element.Styles(c.styler.styles)

    return element
} 