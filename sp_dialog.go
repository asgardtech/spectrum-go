package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DialogMode represents the The display mode of the dialog
type DialogMode string

// DialogMode values
const (
    DialogModeFullscreen DialogMode = "fullscreen"
    DialogModeFullscreentakeover DialogMode = "fullscreenTakeover"
)
// DialogSize represents the Size of the dialog
type DialogSize string

// DialogSize values
const (
    DialogSizeS DialogSize = "s"
    DialogSizeM DialogSize = "m"
    DialogSizeL DialogSize = "l"
)

// spectrumDialog represents an sp-dialog component
type spectrumDialog struct {
    app.Compo
    *styler[*spectrumDialog]

    // Properties
    // Label for the dismiss button
    PropDismisslabel string
    // Whether the dialog has a close button
    PropDismissable bool
    // Whether the dialog represents an error state
    PropError bool
    // The display mode of the dialog
    PropMode DialogMode
    // Whether to hide the divider between header and content
    PropNodivider bool
    // Size of the dialog
    PropSize DialogSize
    // Alert dialog variant
    PropVariant string

    // Text content for the default slot
    PropText string

    // Content slots
    PropButtonSlot app.UI
    PropFooterSlot app.UI
    PropHeadingSlot app.UI
    PropHeroSlot app.UI


    // Event handlers
    PropOnClose app.EventHandler
}

// Dialog creates a new sp-dialog component
func Dialog() *spectrumDialog {
    element := &spectrumDialog{
        PropDismisslabel: "Close",
        PropDismissable: false,
        PropError: false,
        PropMode: "",
        PropNodivider: false,
        PropSize: "",
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Label for the dismiss button
func (c *spectrumDialog) Dismisslabel(dismissLabel string) *spectrumDialog {
    c.PropDismisslabel = dismissLabel
    return c
}

// Whether the dialog has a close button
func (c *spectrumDialog) Dismissable(dismissable bool) *spectrumDialog {
    c.PropDismissable = dismissable
    return c
}

func (c *spectrumDialog) SetDismissable() *spectrumDialog {
    return c.Dismissable(true)
}

// Whether the dialog represents an error state
func (c *spectrumDialog) Error(error bool) *spectrumDialog {
    c.PropError = error
    return c
}

func (c *spectrumDialog) SetError() *spectrumDialog {
    return c.Error(true)
}

// The display mode of the dialog
func (c *spectrumDialog) Mode(mode DialogMode) *spectrumDialog {
    c.PropMode = mode
    return c
}

func (c *spectrumDialog) ModeFullscreen() *spectrumDialog {
    return c.Mode(DialogModeFullscreen)
}
func (c *spectrumDialog) ModeFullscreentakeover() *spectrumDialog {
    return c.Mode(DialogModeFullscreentakeover)
}
// Whether to hide the divider between header and content
func (c *spectrumDialog) Nodivider(noDivider bool) *spectrumDialog {
    c.PropNodivider = noDivider
    return c
}

func (c *spectrumDialog) SetNodivider() *spectrumDialog {
    return c.Nodivider(true)
}

// Size of the dialog
func (c *spectrumDialog) Size(size DialogSize) *spectrumDialog {
    c.PropSize = size
    return c
}

func (c *spectrumDialog) SizeS() *spectrumDialog {
    return c.Size(DialogSizeS)
}
func (c *spectrumDialog) SizeM() *spectrumDialog {
    return c.Size(DialogSizeM)
}
func (c *spectrumDialog) SizeL() *spectrumDialog {
    return c.Size(DialogSizeL)
}
// Alert dialog variant
func (c *spectrumDialog) Variant(variant string) *spectrumDialog {
    c.PropVariant = variant
    return c
}


// Text sets the text content for the default slot
func (c *spectrumDialog) Text(text string) *spectrumDialog {
    c.PropText = text
    return c
}


// Button elements addressed to this slot may be placed below the content when not delivered in a fullscreen mode
func (c *spectrumDialog) Button(content app.UI) *spectrumDialog {
    c.PropButtonSlot = content

    return c
}

// Content addressed to the footer will be placed below the main content and to the side of any [slot='button'] content
func (c *spectrumDialog) Footer(content app.UI) *spectrumDialog {
    c.PropFooterSlot = content

    return c
}

// Acts as the heading of the dialog. This should be an actual heading tag
func (c *spectrumDialog) Heading(content app.UI) *spectrumDialog {
    c.PropHeadingSlot = content

    return c
}

// Accepts a hero image to display at the top of the dialog
func (c *spectrumDialog) Hero(content app.UI) *spectrumDialog {
    c.PropHeroSlot = content

    return c
}



// Announces that the dialog has been closed
func (c *spectrumDialog) OnClose(handler app.EventHandler) *spectrumDialog {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-dialog component
func (c *spectrumDialog) Render() app.UI {
    element := app.Elem("sp-dialog")

    // Set attributes
    if c.PropDismisslabel != "" {
        element = element.Attr("dismissLabel", c.PropDismisslabel)
    }
    if c.PropDismissable {
        element = element.Attr("dismissable", true)
    }
    if c.PropError {
        element = element.Attr("error", true)
    }
    if c.PropMode != "" {
        element = element.Attr("mode", string(c.PropMode))
    }
    if c.PropNodivider {
        element = element.Attr("noDivider", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", c.PropVariant)
    }

    // Add event handlers
    if c.PropOnClose != nil {
        element = element.On("close", c.PropOnClose)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add button slot
    if c.PropButtonSlot != nil {
        slotElem := c.PropButtonSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("button")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "button").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add footer slot
    if c.PropFooterSlot != nil {
        slotElem := c.PropFooterSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("footer")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "footer").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add heading slot
    if c.PropHeadingSlot != nil {
        slotElem := c.PropHeadingSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("heading")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "heading").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add hero slot
    if c.PropHeroSlot != nil {
        slotElem := c.PropHeroSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("hero")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "hero").
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