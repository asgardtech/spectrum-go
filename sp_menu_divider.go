package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumMenuDivider represents an sp-menu-divider component
type spectrumMenuDivider struct {
    app.Compo
    *styler[*spectrumMenuDivider]

    // Properties




}

// MenuDivider creates a new sp-menu-divider component
func MenuDivider() *spectrumMenuDivider {
    element := &spectrumMenuDivider{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the sp-menu-divider component
func (c *spectrumMenuDivider) Render() app.UI {
    element := app.Elem("sp-menu-divider")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 