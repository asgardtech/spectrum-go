package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTags represents an sp-tags component
type spectrumTags struct {
    app.Compo
    *styler[*spectrumTags]

    // Properties


    // Content slots

    // Child components
    PropChildren []app.UI

}

// Tags creates a new sp-tags component
func Tags() *spectrumTags {
    element := &spectrumTags{
    }

    element.styler = newStyler(element)

    return element
}




// Children sets the child components
func (c *spectrumTags) Children(children ...app.UI) *spectrumTags {
    c.PropChildren = children

    return c
}

// AddChild adds a child component
func (c *spectrumTags) AddChild(child app.UI) *spectrumTags {
    c.PropChildren = append(c.PropChildren, child)

    return c
}



// Render renders the sp-tags component
func (c *spectrumTags) Render() app.UI {
    element := app.Elem("sp-tags")

    // Set attributes


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