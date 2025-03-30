package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumAccordionItem represents an sp-accordion-item component
type spectrumAccordionItem struct {
    app.Compo
    *styler[*spectrumAccordionItem]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // The text label for the accordion item header
    PropLabel string
    // Whether the accordion item is open/expanded
    PropOpen bool
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnSpAccordionItemToggle app.EventHandler
}

// AccordionItem creates a new sp-accordion-item component
func AccordionItem() *spectrumAccordionItem {
    element := &spectrumAccordionItem{
        PropDisabled: false,
        PropLabel: "",
        PropOpen: false,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumAccordionItem) Disabled(disabled bool) *spectrumAccordionItem {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumAccordionItem) SetDisabled() *spectrumAccordionItem {
    return c.Disabled(true)
}

// The text label for the accordion item header
func (c *spectrumAccordionItem) Label(label string) *spectrumAccordionItem {
    c.PropLabel = label
    return c
}

// Whether the accordion item is open/expanded
func (c *spectrumAccordionItem) Open(open bool) *spectrumAccordionItem {
    c.PropOpen = open
    return c
}

func (c *spectrumAccordionItem) SetOpen() *spectrumAccordionItem {
    return c.Open(true)
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumAccordionItem) Tabindex(tabindex float64) *spectrumAccordionItem {
    c.PropTabindex = tabindex
    return c
}


// Text sets the text content for the default slot
func (c *spectrumAccordionItem) Text(text string) *spectrumAccordionItem {
    c.PropText = text
    return c
}




// Fired when an accordion item is toggled open or closed
func (c *spectrumAccordionItem) OnSpAccordionItemToggle(handler app.EventHandler) *spectrumAccordionItem {
    c.PropOnSpAccordionItemToggle = handler

    return c
}


// Render renders the sp-accordion-item component
func (c *spectrumAccordionItem) Render() app.UI {
    element := app.Elem("sp-accordion-item")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }

    // Add event handlers
    if c.PropOnSpAccordionItemToggle != nil {
        element = element.On("sp-accordion-item-toggle", c.PropOnSpAccordionItemToggle)
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