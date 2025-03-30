package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconsUISize represents the Size of the icon
type IconsUISize string

// IconsUISize values
const (
    IconsUISizeXxs IconsUISize = "xxs"
    IconsUISizeXs IconsUISize = "xs"
    IconsUISizeS IconsUISize = "s"
    IconsUISizeM IconsUISize = "m"
    IconsUISizeL IconsUISize = "l"
    IconsUISizeXl IconsUISize = "xl"
    IconsUISizeXxl IconsUISize = "xxl"
)

// spectrumIconsUI represents an sp-icon-ui component
type spectrumIconsUI struct {
    app.Compo
    *styler[*spectrumIconsUI]

    // Properties
    // Accessibility label for the icon
    PropLabel string
    // Size of the icon
    PropSize IconsUISize




}

// IconsUI creates a new sp-icon-ui component
func IconsUI() *spectrumIconsUI {
    element := &spectrumIconsUI{
        PropLabel: "",
        PropSize: "",
    }

    element.styler = newStyler(element)

    return element
}

// Accessibility label for the icon
func (c *spectrumIconsUI) Label(label string) *spectrumIconsUI {
    c.PropLabel = label
    return c
}

// Size of the icon
func (c *spectrumIconsUI) Size(size IconsUISize) *spectrumIconsUI {
    c.PropSize = size
    return c
}

func (c *spectrumIconsUI) SizeXxs() *spectrumIconsUI {
    return c.Size(IconsUISizeXxs)
}
func (c *spectrumIconsUI) SizeXs() *spectrumIconsUI {
    return c.Size(IconsUISizeXs)
}
func (c *spectrumIconsUI) SizeS() *spectrumIconsUI {
    return c.Size(IconsUISizeS)
}
func (c *spectrumIconsUI) SizeM() *spectrumIconsUI {
    return c.Size(IconsUISizeM)
}
func (c *spectrumIconsUI) SizeL() *spectrumIconsUI {
    return c.Size(IconsUISizeL)
}
func (c *spectrumIconsUI) SizeXl() *spectrumIconsUI {
    return c.Size(IconsUISizeXl)
}
func (c *spectrumIconsUI) SizeXxl() *spectrumIconsUI {
    return c.Size(IconsUISizeXxl)
}





// Render renders the sp-icon-ui component
func (c *spectrumIconsUI) Render() app.UI {
    element := app.Elem("sp-icon-ui")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }



    element = element.Styles(c.styler.styles)

    return element
} 