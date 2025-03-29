package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorLoupe represents the sp-color-loupe component
// It shows the output color that would otherwise be covered by a cursor, stylus, or finger during color selection.
type spectrumColorLoupe struct {
	app.Compo

	PropColor string
	PropOpen  bool
}

// ColorLoupe creates a new spectrumColorLoupe instance.
func ColorLoupe() *spectrumColorLoupe {
	return &spectrumColorLoupe{
		PropColor: "rgba(255, 0, 0, 0.5)",
	}
}

// Color sets the color property of the spectrumColorLoupe
func (c *spectrumColorLoupe) Color(color string) *spectrumColorLoupe {
	c.PropColor = color
	return c
}

// Open sets the open property of the spectrumColorLoupe
func (c *spectrumColorLoupe) Open(open bool) *spectrumColorLoupe {
	c.PropOpen = open
	return c
}

// Render renders the spectrumColorLoupe component.
func (c *spectrumColorLoupe) Render() app.UI {
	colorLoupe := app.Elem("sp-color-loupe").
		Attr("color", c.PropColor)

	if c.PropOpen {
		colorLoupe = colorLoupe.Attr("open", true)
	}

	return colorLoupe
}
