package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumBundle represents an sp-bundle component
type spectrumBundle struct {
    app.Compo
    *styler[*spectrumBundle]

    // Properties




}

// Bundle creates a new sp-bundle component
func Bundle() *spectrumBundle {
    element := &spectrumBundle{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the sp-bundle component
func (c *spectrumBundle) Render() app.UI {
    element := app.Elem("sp-bundle")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 