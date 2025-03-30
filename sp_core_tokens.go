package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumCoreTokens represents an  component
type spectrumCoreTokens struct {
    app.Compo
    *styler[*spectrumCoreTokens]

    // Properties
    // Custom property used to support calculating the difference between delivery in medium and large scales
    PropSwcScaleFactor string




}

// CoreTokens creates a new  component
func CoreTokens() *spectrumCoreTokens {
    element := &spectrumCoreTokens{
        PropSwcScaleFactor: "1",
    }

    element.styler = newStyler(element)

    return element
}

// Custom property used to support calculating the difference between delivery in medium and large scales
func (c *spectrumCoreTokens) SwcScaleFactor(swcScaleFactor string) *spectrumCoreTokens {
    c.PropSwcScaleFactor = swcScaleFactor
    return c
}






// Render renders the  component
func (c *spectrumCoreTokens) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropSwcScaleFactor != "" {
        element = element.Attr("swc-scale-factor", c.PropSwcScaleFactor)
    }



    element = element.Styles(c.styler.styles)

    return element
} 