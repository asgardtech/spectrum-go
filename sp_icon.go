package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconSize represents the Size of the icon
type IconSize string

// IconSize values
const (
    IconSizeXxs IconSize = "xxs"
    IconSizeXs IconSize = "xs"
    IconSizeS IconSize = "s"
    IconSizeM IconSize = "m"
    IconSizeL IconSize = "l"
    IconSizeXl IconSize = "xl"
    IconSizeXxl IconSize = "xxl"
)

// spectrumIcon represents an sp-icon component
type spectrumIcon struct {
    app.Compo
    *styler[*spectrumIcon]

    // Properties
    // Accessible label for the icon. When provided, removes aria-hidden and sets aria-label instead
    PropLabel string
    // Name of the icon in the registered icon set to display
    PropName string
    // Size of the icon
    PropSize IconSize
    // URL to an image to use as the icon
    PropSrc string

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnError app.EventHandler
}

// Icon creates a new sp-icon component
func Icon() *spectrumIcon {
    element := &spectrumIcon{
        PropLabel: "",
        PropSize: IconSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Accessible label for the icon. When provided, removes aria-hidden and sets aria-label instead
func (c *spectrumIcon) Label(label string) *spectrumIcon {
    c.PropLabel = label
    return c
}

// Name of the icon in the registered icon set to display
func (c *spectrumIcon) Name(name string) *spectrumIcon {
    c.PropName = name
    return c
}

// Size of the icon
func (c *spectrumIcon) Size(size IconSize) *spectrumIcon {
    c.PropSize = size
    return c
}

func (c *spectrumIcon) SizeXxs() *spectrumIcon {
    return c.Size(IconSizeXxs)
}
func (c *spectrumIcon) SizeXs() *spectrumIcon {
    return c.Size(IconSizeXs)
}
func (c *spectrumIcon) SizeS() *spectrumIcon {
    return c.Size(IconSizeS)
}
func (c *spectrumIcon) SizeM() *spectrumIcon {
    return c.Size(IconSizeM)
}
func (c *spectrumIcon) SizeL() *spectrumIcon {
    return c.Size(IconSizeL)
}
func (c *spectrumIcon) SizeXl() *spectrumIcon {
    return c.Size(IconSizeXl)
}
func (c *spectrumIcon) SizeXxl() *spectrumIcon {
    return c.Size(IconSizeXxl)
}
// URL to an image to use as the icon
func (c *spectrumIcon) Src(src string) *spectrumIcon {
    c.PropSrc = src
    return c
}


// Text sets the text content for the default slot
func (c *spectrumIcon) Text(text string) *spectrumIcon {
    c.PropText = text
    return c
}




// Fired when there is an error loading the icon
func (c *spectrumIcon) OnError(handler app.EventHandler) *spectrumIcon {
    c.PropOnError = handler

    return c
}


// Render renders the sp-icon component
func (c *spectrumIcon) Render() app.UI {
    element := app.Elem("sp-icon")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropSrc != "" {
        element = element.Attr("src", c.PropSrc)
    }

    // Add event handlers
    if c.PropOnError != nil {
        element = element.On("error", c.PropOnError)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 