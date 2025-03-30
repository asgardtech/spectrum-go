package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumOpacityCheckerboard represents an  component
type spectrumOpacityCheckerboard struct {
    app.Compo
    *styler[*spectrumOpacityCheckerboard]

    // Properties




}

// OpacityCheckerboard creates a new  component
func OpacityCheckerboard() *spectrumOpacityCheckerboard {
    element := &spectrumOpacityCheckerboard{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the  component
func (c *spectrumOpacityCheckerboard) Render() app.UI {
    element := app.Elem("")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 