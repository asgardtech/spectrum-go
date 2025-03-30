package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorController represents an  component
type spectrumColorController struct {
    app.Compo
    *styler[*spectrumColorController]

    // Properties
    // Gets or sets the current color value in various formats (strings, objects, or instances of the Color class)
    PropColor string
    // Gets the color value in various formats based on the original color input
    PropColorvalue string
    // Gets or sets the hue value of the current color
    PropHue float64




}

// ColorController creates a new  component
func ColorController() *spectrumColorController {
    element := &spectrumColorController{
        PropHue: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Gets or sets the current color value in various formats (strings, objects, or instances of the Color class)
func (c *spectrumColorController) Color(color string) *spectrumColorController {
    c.PropColor = color
    return c
}

// Gets the color value in various formats based on the original color input
func (c *spectrumColorController) Colorvalue(colorValue string) *spectrumColorController {
    c.PropColorvalue = colorValue
    return c
}

// Gets or sets the hue value of the current color
func (c *spectrumColorController) Hue(hue float64) *spectrumColorController {
    c.PropHue = hue
    return c
}






// Render renders the  component
func (c *spectrumColorController) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropColor != "" {
        element = element.Attr("color", c.PropColor)
    }
    if c.PropColorvalue != "" {
        element = element.Attr("colorValue", c.PropColorvalue)
    }
    if c.PropHue != 0 {
        element = element.Attr("hue", c.PropHue)
    }



    element = element.Styles(c.styler.styles)

    return element
} 