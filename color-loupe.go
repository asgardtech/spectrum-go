package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumColorLoupe represents the sp-color-loupe component
// It shows the output color that would otherwise be covered by a cursor, stylus, or finger during color selection.
type SpectrumColorLoupe struct {
	app.Compo

	color string
	open  bool
}

// ColorLoupe creates a new SpectrumColorLoupe instance.
func ColorLoupe() *SpectrumColorLoupe {
	return &SpectrumColorLoupe{
		color: "rgba(255, 0, 0, 0.5)",
	}
}

// Color sets the color property of the SpectrumColorLoupe
func (c *SpectrumColorLoupe) Color(color string) *SpectrumColorLoupe {
	c.color = color
	return c
}

// Open sets the open property of the SpectrumColorLoupe
func (c *SpectrumColorLoupe) Open(open bool) *SpectrumColorLoupe {
	c.open = open
	return c
}

// Render renders the SpectrumColorLoupe component.
func (c *SpectrumColorLoupe) Render() app.UI {
	colorLoupe := app.Elem("sp-color-loupe").
		Attr("color", c.color)

	if c.open {
		colorLoupe = colorLoupe.Attr("open", true)
	}

	return colorLoupe
}
