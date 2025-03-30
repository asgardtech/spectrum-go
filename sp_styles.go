package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumStyles represents an  component
type spectrumStyles struct {
    app.Compo
    *styler[*spectrumStyles]

    // Properties




}

// Styles creates a new  component
func Styles() *spectrumStyles {
    element := &spectrumStyles{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the  component
func (c *spectrumStyles) Render() app.UI {
    element := app.Elem("")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 