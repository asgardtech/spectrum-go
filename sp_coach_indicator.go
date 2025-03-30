package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// CoachIndicatorStaticcolor represents the Applies a specific color treatment to the indicator
type CoachIndicatorStaticcolor string

// CoachIndicatorStaticcolor values
const (
    CoachIndicatorStaticcolorWhite CoachIndicatorStaticcolor = "white"
    CoachIndicatorStaticcolorBlack CoachIndicatorStaticcolor = "black"
    CoachIndicatorStaticcolorLight CoachIndicatorStaticcolor = "light"
    CoachIndicatorStaticcolorDark CoachIndicatorStaticcolor = "dark"
)

// spectrumCoachIndicator represents an sp-coach-indicator component
type spectrumCoachIndicator struct {
    app.Compo
    *styler[*spectrumCoachIndicator]

    // Properties
    // Applies a less visually prominent style
    PropQuiet bool
    // Applies a specific color treatment to the indicator
    PropStaticcolor CoachIndicatorStaticcolor




}

// CoachIndicator creates a new sp-coach-indicator component
func CoachIndicator() *spectrumCoachIndicator {
    element := &spectrumCoachIndicator{
        PropQuiet: false,
        PropStaticcolor: "",
    }

    element.styler = newStyler(element)

    return element
}

// Applies a less visually prominent style
func (c *spectrumCoachIndicator) Quiet(quiet bool) *spectrumCoachIndicator {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumCoachIndicator) SetQuiet() *spectrumCoachIndicator {
    return c.Quiet(true)
}

// Applies a specific color treatment to the indicator
func (c *spectrumCoachIndicator) Staticcolor(staticColor CoachIndicatorStaticcolor) *spectrumCoachIndicator {
    c.PropStaticcolor = staticColor
    return c
}

func (c *spectrumCoachIndicator) StaticcolorWhite() *spectrumCoachIndicator {
    return c.Staticcolor(CoachIndicatorStaticcolorWhite)
}
func (c *spectrumCoachIndicator) StaticcolorBlack() *spectrumCoachIndicator {
    return c.Staticcolor(CoachIndicatorStaticcolorBlack)
}
func (c *spectrumCoachIndicator) StaticcolorLight() *spectrumCoachIndicator {
    return c.Staticcolor(CoachIndicatorStaticcolorLight)
}
func (c *spectrumCoachIndicator) StaticcolorDark() *spectrumCoachIndicator {
    return c.Staticcolor(CoachIndicatorStaticcolorDark)
}





// Render renders the sp-coach-indicator component
func (c *spectrumCoachIndicator) Render() app.UI {
    element := app.Elem("sp-coach-indicator")

    // Set attributes
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropStaticcolor != "" {
        element = element.Attr("staticColor", string(c.PropStaticcolor))
    }



    element = element.Styles(c.styler.styles)

    return element
} 