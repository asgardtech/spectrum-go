// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorLoupe represents an sp-color-loupe component
type spectrumColorLoupe struct {
	app.Compo
	*styler[*spectrumColorLoupe]
	*classer[*spectrumColorLoupe]
	*ider[*spectrumColorLoupe]

	// Properties
	// The color to display in the loupe
	PropColor string
	// Controls whether the loupe is visible
	PropOpen bool
}

// IColorLoupe is the interface for sp-color-loupe component methods
type IColorLoupe interface {
	app.UI
	Styler[IColorLoupe]
	Classer[IColorLoupe]
	Ider[IColorLoupe]
	Color(string) IColorLoupe
	Open(bool) IColorLoupe
	SetOpen() IColorLoupe
}

// ColorLoupe An sp-color-loupe shows the output color that would otherwise be covered by a cursor, stylus, or finger during color selection.
func ColorLoupe() IColorLoupe {
	element := &spectrumColorLoupe{
		PropColor: "rgba(255, 0, 0, 0.5)",
		PropOpen:  false,
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Color The color to display in the loupe
func (c *spectrumColorLoupe) Color(color string) IColorLoupe {
	c.PropColor = color
	return c
}

// Open Controls whether the loupe is visible
func (c *spectrumColorLoupe) Open(open bool) IColorLoupe {
	c.PropOpen = open
	return c
}

func (c *spectrumColorLoupe) SetOpen() IColorLoupe {
	return c.Open(true)
}

// Style sets a style property with a value
func (c *spectrumColorLoupe) Style(key, format string, values ...any) IColorLoupe {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumColorLoupe) Styles(styles map[string]string) IColorLoupe {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumColorLoupe) Class(class string) IColorLoupe {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumColorLoupe) Classes(classes ...string) IColorLoupe {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumColorLoupe) Id(id string) IColorLoupe {
	return c.ider.Id(id)
}

// Render renders the sp-color-loupe component
func (c *spectrumColorLoupe) Render() app.UI {
	element := app.Elem("sp-color-loupe")

	// Set attributes
	if c.PropColor != "" {
		element = element.Attr("color", c.PropColor)
	}
	if c.PropOpen {
		element = element.Attr("open", true)
	}

	// Apply styles, classes, and id
	element = element.Styles(c.styler.styles)

	// Apply classes if any
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}

	// Apply id if set
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}

	return element
}
