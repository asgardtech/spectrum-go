package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionBarVariant represents the The variant applies specific positioning: 'sticky' positions relative to scrolling content, 'fixed' positions relative to the page
type ActionBarVariant string

// ActionBarVariant values
const (
    ActionBarVariantFixed ActionBarVariant = "fixed"
    ActionBarVariantSticky ActionBarVariant = "sticky"
)

// spectrumActionBar represents an sp-action-bar component
type spectrumActionBar struct {
    app.Compo
    *styler[*spectrumActionBar]

    // Properties
    // Deliver the Action Bar with additional visual emphasis
    PropEmphasized bool
    // When 'flexible' the action bar sizes itself to its content rather than a specific width
    PropFlexible bool
    // Whether the action bar is visible
    PropOpen bool
    // The variant applies specific positioning: 'sticky' positions relative to scrolling content, 'fixed' positions relative to the page
    PropVariant ActionBarVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropButtonsSlot app.UI


    // Event handlers
    PropOnClose app.EventHandler
}

// ActionBar creates a new sp-action-bar component
func ActionBar() *spectrumActionBar {
    element := &spectrumActionBar{
        PropEmphasized: false,
        PropFlexible: false,
        PropOpen: false,
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Deliver the Action Bar with additional visual emphasis
func (c *spectrumActionBar) Emphasized(emphasized bool) *spectrumActionBar {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumActionBar) SetEmphasized() *spectrumActionBar {
    return c.Emphasized(true)
}

// When 'flexible' the action bar sizes itself to its content rather than a specific width
func (c *spectrumActionBar) Flexible(flexible bool) *spectrumActionBar {
    c.PropFlexible = flexible
    return c
}

func (c *spectrumActionBar) SetFlexible() *spectrumActionBar {
    return c.Flexible(true)
}

// Whether the action bar is visible
func (c *spectrumActionBar) Open(open bool) *spectrumActionBar {
    c.PropOpen = open
    return c
}

func (c *spectrumActionBar) SetOpen() *spectrumActionBar {
    return c.Open(true)
}

// The variant applies specific positioning: 'sticky' positions relative to scrolling content, 'fixed' positions relative to the page
func (c *spectrumActionBar) Variant(variant ActionBarVariant) *spectrumActionBar {
    c.PropVariant = variant
    return c
}

func (c *spectrumActionBar) VariantFixed() *spectrumActionBar {
    return c.Variant(ActionBarVariantFixed)
}
func (c *spectrumActionBar) VariantSticky() *spectrumActionBar {
    return c.Variant(ActionBarVariantSticky)
}

// Text sets the text content for the default slot
func (c *spectrumActionBar) Text(text string) *spectrumActionBar {
    c.PropText = text
    return c
}


// Action buttons to display in the Action Bar
func (c *spectrumActionBar) Buttons(content app.UI) *spectrumActionBar {
    c.PropButtonsSlot = content

    return c
}



// Dispatched when the action bar is closed
func (c *spectrumActionBar) OnClose(handler app.EventHandler) *spectrumActionBar {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-action-bar component
func (c *spectrumActionBar) Render() app.UI {
    element := app.Elem("sp-action-bar")

    // Set attributes
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
    }
    if c.PropFlexible {
        element = element.Attr("flexible", true)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnClose != nil {
        element = element.On("close", c.PropOnClose)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add buttons slot
    if c.PropButtonsSlot != nil {
        slotElem := c.PropButtonsSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("buttons")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "buttons").
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