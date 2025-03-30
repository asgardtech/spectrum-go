package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTabsOverflow represents an sp-tabs-overflow component
type spectrumTabsOverflow struct {
    app.Compo
    *styler[*spectrumTabsOverflow]

    // Properties
    // Whether the tabs overflow should use a compact style
    PropCompact bool
    // Accessible label for the next button
    PropLabelnext string
    // Accessible label for the previous button
    PropLabelprevious string

    // Text content for the default slot
    PropText string

    // Content slots


}

// TabsOverflow creates a new sp-tabs-overflow component
func TabsOverflow() *spectrumTabsOverflow {
    element := &spectrumTabsOverflow{
        PropCompact: false,
        PropLabelnext: "Scroll to next tabs",
        PropLabelprevious: "Scroll to previous tabs",
    }

    element.styler = newStyler(element)

    return element
}

// Whether the tabs overflow should use a compact style
func (c *spectrumTabsOverflow) Compact(compact bool) *spectrumTabsOverflow {
    c.PropCompact = compact
    return c
}

func (c *spectrumTabsOverflow) SetCompact() *spectrumTabsOverflow {
    return c.Compact(true)
}

// Accessible label for the next button
func (c *spectrumTabsOverflow) Labelnext(labelNext string) *spectrumTabsOverflow {
    c.PropLabelnext = labelNext
    return c
}

// Accessible label for the previous button
func (c *spectrumTabsOverflow) Labelprevious(labelPrevious string) *spectrumTabsOverflow {
    c.PropLabelprevious = labelPrevious
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTabsOverflow) Text(text string) *spectrumTabsOverflow {
    c.PropText = text
    return c
}





// Render renders the sp-tabs-overflow component
func (c *spectrumTabsOverflow) Render() app.UI {
    element := app.Elem("sp-tabs-overflow")

    // Set attributes
    if c.PropCompact {
        element = element.Attr("compact", true)
    }
    if c.PropLabelnext != "" {
        element = element.Attr("labelNext", c.PropLabelnext)
    }
    if c.PropLabelprevious != "" {
        element = element.Attr("labelPrevious", c.PropLabelprevious)
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