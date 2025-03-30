package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumShared represents an  component
type spectrumShared struct {
    app.Compo
    *styler[*spectrumShared]

    // Properties




}

// Shared creates a new  component
func Shared() *spectrumShared {
    element := &spectrumShared{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the  component
func (c *spectrumShared) Render() app.UI {
    element := app.Elem("")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 