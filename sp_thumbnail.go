package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ThumbnailSize represents the The size of the thumbnail in pixels
type ThumbnailSize string

// ThumbnailSize values
const (
    ThumbnailSize50 ThumbnailSize = "50"
    ThumbnailSize75 ThumbnailSize = "75"
    ThumbnailSize100 ThumbnailSize = "100"
    ThumbnailSize200 ThumbnailSize = "200"
    ThumbnailSize300 ThumbnailSize = "300"
    ThumbnailSize400 ThumbnailSize = "400"
    ThumbnailSize500 ThumbnailSize = "500"
    ThumbnailSize600 ThumbnailSize = "600"
    ThumbnailSize700 ThumbnailSize = "700"
    ThumbnailSize800 ThumbnailSize = "800"
    ThumbnailSize900 ThumbnailSize = "900"
    ThumbnailSize1000 ThumbnailSize = "1000"
)

// spectrumThumbnail represents an sp-thumbnail component
type spectrumThumbnail struct {
    app.Compo
    *styler[*spectrumThumbnail]

    // Properties
    // CSS background property value to customize the letterboxing background
    PropBackground string
    // Causes the content to fill the space provided by the thumbnail element
    PropCover bool
    // Disables the thumbnail
    PropDisabled bool
    // Displays the thumbnail with a focus indicator
    PropFocused bool
    // Shows the thumbnail with a layer style, for use in layer management interfaces
    PropLayer bool
    // Shows the thumbnail in a selected state
    PropSelected bool
    // The size of the thumbnail in pixels
    PropSize ThumbnailSize

    // Text content for the default slot
    PropText string

    // Content slots
    PropImageSlot app.UI


}

// Thumbnail creates a new sp-thumbnail component
func Thumbnail() *spectrumThumbnail {
    element := &spectrumThumbnail{
        PropCover: false,
        PropDisabled: false,
        PropFocused: false,
        PropLayer: false,
        PropSelected: false,
        PropSize: ThumbnailSize100,
    }

    element.styler = newStyler(element)

    return element
}

// CSS background property value to customize the letterboxing background
func (c *spectrumThumbnail) Background(background string) *spectrumThumbnail {
    c.PropBackground = background
    return c
}

// Causes the content to fill the space provided by the thumbnail element
func (c *spectrumThumbnail) Cover(cover bool) *spectrumThumbnail {
    c.PropCover = cover
    return c
}

func (c *spectrumThumbnail) SetCover() *spectrumThumbnail {
    return c.Cover(true)
}

// Disables the thumbnail
func (c *spectrumThumbnail) Disabled(disabled bool) *spectrumThumbnail {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumThumbnail) SetDisabled() *spectrumThumbnail {
    return c.Disabled(true)
}

// Displays the thumbnail with a focus indicator
func (c *spectrumThumbnail) Focused(focused bool) *spectrumThumbnail {
    c.PropFocused = focused
    return c
}

func (c *spectrumThumbnail) SetFocused() *spectrumThumbnail {
    return c.Focused(true)
}

// Shows the thumbnail with a layer style, for use in layer management interfaces
func (c *spectrumThumbnail) Layer(layer bool) *spectrumThumbnail {
    c.PropLayer = layer
    return c
}

func (c *spectrumThumbnail) SetLayer() *spectrumThumbnail {
    return c.Layer(true)
}

// Shows the thumbnail in a selected state
func (c *spectrumThumbnail) Selected(selected bool) *spectrumThumbnail {
    c.PropSelected = selected
    return c
}

func (c *spectrumThumbnail) SetSelected() *spectrumThumbnail {
    return c.Selected(true)
}

// The size of the thumbnail in pixels
func (c *spectrumThumbnail) Size(size ThumbnailSize) *spectrumThumbnail {
    c.PropSize = size
    return c
}

func (c *spectrumThumbnail) Size50() *spectrumThumbnail {
    return c.Size(ThumbnailSize50)
}
func (c *spectrumThumbnail) Size75() *spectrumThumbnail {
    return c.Size(ThumbnailSize75)
}
func (c *spectrumThumbnail) Size100() *spectrumThumbnail {
    return c.Size(ThumbnailSize100)
}
func (c *spectrumThumbnail) Size200() *spectrumThumbnail {
    return c.Size(ThumbnailSize200)
}
func (c *spectrumThumbnail) Size300() *spectrumThumbnail {
    return c.Size(ThumbnailSize300)
}
func (c *spectrumThumbnail) Size400() *spectrumThumbnail {
    return c.Size(ThumbnailSize400)
}
func (c *spectrumThumbnail) Size500() *spectrumThumbnail {
    return c.Size(ThumbnailSize500)
}
func (c *spectrumThumbnail) Size600() *spectrumThumbnail {
    return c.Size(ThumbnailSize600)
}
func (c *spectrumThumbnail) Size700() *spectrumThumbnail {
    return c.Size(ThumbnailSize700)
}
func (c *spectrumThumbnail) Size800() *spectrumThumbnail {
    return c.Size(ThumbnailSize800)
}
func (c *spectrumThumbnail) Size900() *spectrumThumbnail {
    return c.Size(ThumbnailSize900)
}
func (c *spectrumThumbnail) Size1000() *spectrumThumbnail {
    return c.Size(ThumbnailSize1000)
}

// Text sets the text content for the default slot
func (c *spectrumThumbnail) Text(text string) *spectrumThumbnail {
    c.PropText = text
    return c
}


// Image element to present in the thumbnail
func (c *spectrumThumbnail) Image(content app.UI) *spectrumThumbnail {
    c.PropImageSlot = content

    return c
}




// Render renders the sp-thumbnail component
func (c *spectrumThumbnail) Render() app.UI {
    element := app.Elem("sp-thumbnail")

    // Set attributes
    if c.PropBackground != "" {
        element = element.Attr("background", c.PropBackground)
    }
    if c.PropCover {
        element = element.Attr("cover", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFocused {
        element = element.Attr("focused", true)
    }
    if c.PropLayer {
        element = element.Attr("layer", true)
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add image slot
    if c.PropImageSlot != nil {
        slotElem := c.PropImageSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("image")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "image").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 