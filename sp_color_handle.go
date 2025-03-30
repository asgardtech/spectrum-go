package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumColorHandle represents an sp-color-handle component
type spectrumColorHandle struct {
    app.Compo
    *styler[*spectrumColorHandle]

    // Properties
    // The color to display in the handle
    PropColor string
    // Disables interaction with the handle
    PropDisabled bool
    // Indicates whether the handle has focus
    PropFocused bool
    // When true, displays a color loupe above the handle to show the selected color
    PropOpen bool




}

// ColorHandle creates a new sp-color-handle component
func ColorHandle() *spectrumColorHandle {
    element := &spectrumColorHandle{
        PropColor: "rgba(255, 0, 0, 0.5)",
        PropDisabled: false,
        PropFocused: false,
        PropOpen: false,
    }

    element.styler = newStyler(element)

    return element
}

// The color to display in the handle
func (c *spectrumColorHandle) Color(color string) *spectrumColorHandle {
    c.PropColor = color
    return c
}

// Disables interaction with the handle
func (c *spectrumColorHandle) Disabled(disabled bool) *spectrumColorHandle {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumColorHandle) SetDisabled() *spectrumColorHandle {
    return c.Disabled(true)
}

// Indicates whether the handle has focus
func (c *spectrumColorHandle) Focused(focused bool) *spectrumColorHandle {
    c.PropFocused = focused
    return c
}

func (c *spectrumColorHandle) SetFocused() *spectrumColorHandle {
    return c.Focused(true)
}

// When true, displays a color loupe above the handle to show the selected color
func (c *spectrumColorHandle) Open(open bool) *spectrumColorHandle {
    c.PropOpen = open
    return c
}

func (c *spectrumColorHandle) SetOpen() *spectrumColorHandle {
    return c.Open(true)
}






// Render renders the sp-color-handle component
func (c *spectrumColorHandle) Render() app.UI {
    element := app.Elem("sp-color-handle")

    // Set attributes
    if c.PropColor != "" {
        element = element.Attr("color", c.PropColor)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFocused {
        element = element.Attr("focused", true)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }



    element = element.Styles(c.styler.styles)

    return element
} 