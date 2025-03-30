package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumMatchMedia represents an  component
type spectrumMatchMedia struct {
    app.Compo
    *styler[*spectrumMatchMedia]

    // Properties
    // Indicates whether the media query currently matches
    PropMatches bool




}

// MatchMedia creates a new  component
func MatchMedia() *spectrumMatchMedia {
    element := &spectrumMatchMedia{
        PropMatches: false,
    }

    element.styler = newStyler(element)

    return element
}

// Indicates whether the media query currently matches
func (c *spectrumMatchMedia) Matches(matches bool) *spectrumMatchMedia {
    c.PropMatches = matches
    return c
}

func (c *spectrumMatchMedia) SetMatches() *spectrumMatchMedia {
    return c.Matches(true)
}






// Render renders the  component
func (c *spectrumMatchMedia) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropMatches {
        element = element.Attr("matches", true)
    }



    element = element.Styles(c.styler.styles)

    return element
} 