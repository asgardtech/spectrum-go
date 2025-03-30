package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTab represents an sp-tab component
type spectrumTab struct {
    app.Compo
    *styler[*spectrumTab]

    // Properties
    // Prevents visitors from selecting this tab and its contents
    PropDisabled bool
    // Text label for the tab, used when no content is provided in the default slot
    PropLabel string
    // Indicates whether this tab is currently selected
    PropSelected bool
    // Value that associates this tab with its corresponding tab panel
    PropValue string
    // Displays the tab's icon and label in a vertical alignment
    PropVertical bool

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


}

// Tab creates a new sp-tab component
func Tab() *spectrumTab {
    element := &spectrumTab{
        PropDisabled: false,
        PropLabel: "",
        PropSelected: false,
        PropValue: "",
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Prevents visitors from selecting this tab and its contents
func (c *spectrumTab) Disabled(disabled bool) *spectrumTab {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumTab) SetDisabled() *spectrumTab {
    return c.Disabled(true)
}

// Text label for the tab, used when no content is provided in the default slot
func (c *spectrumTab) Label(label string) *spectrumTab {
    c.PropLabel = label
    return c
}

// Indicates whether this tab is currently selected
func (c *spectrumTab) Selected(selected bool) *spectrumTab {
    c.PropSelected = selected
    return c
}

func (c *spectrumTab) SetSelected() *spectrumTab {
    return c.Selected(true)
}

// Value that associates this tab with its corresponding tab panel
func (c *spectrumTab) Value(value string) *spectrumTab {
    c.PropValue = value
    return c
}

// Displays the tab's icon and label in a vertical alignment
func (c *spectrumTab) Vertical(vertical bool) *spectrumTab {
    c.PropVertical = vertical
    return c
}

func (c *spectrumTab) SetVertical() *spectrumTab {
    return c.Vertical(true)
}


// Text sets the text content for the default slot
func (c *spectrumTab) Text(text string) *spectrumTab {
    c.PropText = text
    return c
}


// The icon that appears on the left of the label
func (c *spectrumTab) Icon(content app.UI) *spectrumTab {
    c.PropIconSlot = content

    return c
}




// Render renders the sp-tab component
func (c *spectrumTab) Render() app.UI {
    element := app.Elem("sp-tab")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
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