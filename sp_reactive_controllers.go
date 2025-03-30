package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumReactiveControllers represents an  component
type spectrumReactiveControllers struct {
    app.Compo
    *styler[*spectrumReactiveControllers]

    // Properties




}

// ReactiveControllers creates a new  component
func ReactiveControllers() *spectrumReactiveControllers {
    element := &spectrumReactiveControllers{
    }

    element.styler = newStyler(element)

    return element
}






// Render renders the  component
func (c *spectrumReactiveControllers) Render() app.UI {
    element := app.Elem("")

    // Set attributes



    element = element.Styles(c.styler.styles)

    return element
} 