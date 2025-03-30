package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumUnderlay represents an sp-underlay component
type spectrumUnderlay struct {
    app.Compo
    *styler[*spectrumUnderlay]

    // Properties
    // Whether the underlay is currently visible
    PropOpen bool




    // Event handlers
    PropOnClose app.EventHandler
}

// Underlay creates a new sp-underlay component
func Underlay() *spectrumUnderlay {
    element := &spectrumUnderlay{
        PropOpen: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the underlay is currently visible
func (c *spectrumUnderlay) Open(open bool) *spectrumUnderlay {
    c.PropOpen = open
    return c
}

func (c *spectrumUnderlay) SetOpen() *spectrumUnderlay {
    return c.Open(true)
}





// When the underlay is clicked and the consuming pattern should choose whether to close based on that interaction
func (c *spectrumUnderlay) OnClose(handler app.EventHandler) *spectrumUnderlay {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-underlay component
func (c *spectrumUnderlay) Render() app.UI {
    element := app.Elem("sp-underlay")

    // Set attributes
    if c.PropOpen {
        element = element.Attr("open", true)
    }

    // Add event handlers
    if c.PropOnClose != nil {
        element = element.On("close", c.PropOnClose)
    }


    element = element.Styles(c.styler.styles)

    return element
} 