package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorLoupe represents an sp-color-loupe component
type spectrumColorLoupe struct {
    app.Compo
    *styler[*spectrumColorLoupe]

    // Properties
    // The color to display in the loupe
    PropColor string
    // Controls whether the loupe is visible
    PropOpen bool




}

// ColorLoupe creates a new sp-color-loupe component
func ColorLoupe() *spectrumColorLoupe {
    element := &spectrumColorLoupe{
        PropColor: "rgba(255, 0, 0, 0.5)",
        PropOpen: false,
    }

    element.styler = newStyler(element)

    return element
}

// The color to display in the loupe
func (c *spectrumColorLoupe) Color(color string) *spectrumColorLoupe {
    c.PropColor = color
    return c
}

// Controls whether the loupe is visible
func (c *spectrumColorLoupe) Open(open bool) *spectrumColorLoupe {
    c.PropOpen = open
    return c
}

func (c *spectrumColorLoupe) SetOpen() *spectrumColorLoupe {
    return c.Open(true)
}






// Render renders the sp-color-loupe component
func (c *spectrumColorLoupe) Render() app.UI {
    element := app.Elem("sp-color-loupe")

    // Set attributes
    if c.PropColor != "" {
        element = element.Attr("color", c.PropColor)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }



    element = element.Styles(c.styler.styles)

    return element
} 