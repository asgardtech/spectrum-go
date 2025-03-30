package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTabPanel represents an sp-tab-panel component
type spectrumTabPanel struct {
    app.Compo
    *styler[*spectrumTabPanel]

    // Properties
    // Indicates whether this panel is currently selected and visible
    PropSelected bool
    // Value that associates this panel with its corresponding tab
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots


}

// TabPanel creates a new sp-tab-panel component
func TabPanel() *spectrumTabPanel {
    element := &spectrumTabPanel{
        PropSelected: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Indicates whether this panel is currently selected and visible
func (c *spectrumTabPanel) Selected(selected bool) *spectrumTabPanel {
    c.PropSelected = selected
    return c
}

func (c *spectrumTabPanel) SetSelected() *spectrumTabPanel {
    return c.Selected(true)
}

// Value that associates this panel with its corresponding tab
func (c *spectrumTabPanel) Value(value string) *spectrumTabPanel {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTabPanel) Text(text string) *spectrumTabPanel {
    c.PropText = text
    return c
}





// Render renders the sp-tab-panel component
func (c *spectrumTabPanel) Render() app.UI {
    element := app.Elem("sp-tab-panel")

    // Set attributes
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
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