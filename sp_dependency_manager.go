package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumDependencyManager represents an  component
type spectrumDependencyManager struct {
    app.Compo
    *styler[*spectrumDependencyManager]

    // Properties
    // Indicates whether all managed dependencies have been loaded
    PropLoaded bool




}

// DependencyManager creates a new  component
func DependencyManager() *spectrumDependencyManager {
    element := &spectrumDependencyManager{
        PropLoaded: false,
    }

    element.styler = newStyler(element)

    return element
}

// Indicates whether all managed dependencies have been loaded
func (c *spectrumDependencyManager) Loaded(loaded bool) *spectrumDependencyManager {
    c.PropLoaded = loaded
    return c
}

func (c *spectrumDependencyManager) SetLoaded() *spectrumDependencyManager {
    return c.Loaded(true)
}






// Render renders the  component
func (c *spectrumDependencyManager) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropLoaded {
        element = element.Attr("loaded", true)
    }



    element = element.Styles(c.styler.styles)

    return element
} 