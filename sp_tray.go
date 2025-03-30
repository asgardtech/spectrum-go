package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTray represents an sp-tray component
type spectrumTray struct {
    app.Compo
    *styler[*spectrumTray]

    // Properties
    // Whether the tray is currently open and visible
    PropOpen bool

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnClose app.EventHandler
}

// Tray creates a new sp-tray component
func Tray() *spectrumTray {
    element := &spectrumTray{
        PropOpen: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the tray is currently open and visible
func (c *spectrumTray) Open(open bool) *spectrumTray {
    c.PropOpen = open
    return c
}

func (c *spectrumTray) SetOpen() *spectrumTray {
    return c.Open(true)
}


// Text sets the text content for the default slot
func (c *spectrumTray) Text(text string) *spectrumTray {
    c.PropText = text
    return c
}




// Announces that the tray has been closed
func (c *spectrumTray) OnClose(handler app.EventHandler) *spectrumTray {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-tray component
func (c *spectrumTray) Render() app.UI {
    element := app.Elem("sp-tray")

    // Set attributes
    if c.PropOpen {
        element = element.Attr("open", true)
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



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 