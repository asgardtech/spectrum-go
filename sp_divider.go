package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DividerSize represents the The thickness of the divider
type DividerSize string

// DividerSize values
const (
    DividerSizeS DividerSize = "s"
    DividerSizeM DividerSize = "m"
    DividerSizeL DividerSize = "l"
)

// spectrumDivider represents an sp-divider component
type spectrumDivider struct {
    app.Compo
    *styler[*spectrumDivider]

    // Properties
    // Whether the divider is vertical instead of horizontal
    PropVertical bool
    // The thickness of the divider
    PropSize DividerSize




}

// Divider creates a new sp-divider component
func Divider() *spectrumDivider {
    element := &spectrumDivider{
        PropVertical: false,
        PropSize: DividerSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the divider is vertical instead of horizontal
func (c *spectrumDivider) Vertical(vertical bool) *spectrumDivider {
    c.PropVertical = vertical
    return c
}

func (c *spectrumDivider) SetVertical() *spectrumDivider {
    return c.Vertical(true)
}

// The thickness of the divider
func (c *spectrumDivider) Size(size DividerSize) *spectrumDivider {
    c.PropSize = size
    return c
}

func (c *spectrumDivider) SizeS() *spectrumDivider {
    return c.Size(DividerSizeS)
}
func (c *spectrumDivider) SizeM() *spectrumDivider {
    return c.Size(DividerSizeM)
}
func (c *spectrumDivider) SizeL() *spectrumDivider {
    return c.Size(DividerSizeL)
}





// Render renders the sp-divider component
func (c *spectrumDivider) Render() app.UI {
    element := app.Elem("sp-divider")

    // Set attributes
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }



    element = element.Styles(c.styler.styles)

    return element
} 