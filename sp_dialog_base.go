package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DialogBaseMode represents the The display mode of the dialog
type DialogBaseMode string

// DialogBaseMode values
const (
    DialogBaseModeFullscreen DialogBaseMode = "fullscreen"
    DialogBaseModeFullscreentakeover DialogBaseMode = "fullscreenTakeover"
)

// spectrumDialogBase represents an sp-dialog-base component
type spectrumDialogBase struct {
    app.Compo
    *styler[*spectrumDialogBase]

    // Properties
    // Whether the dialog is dismissable
    PropDismissable bool
    // The display mode of the dialog
    PropMode DialogBaseMode
    // Whether the dialog is open
    PropOpen bool
    // When set to true, fills screens smaller than 350px high and 400px wide with the full dialog
    PropResponsive bool
    // Whether to show an underlay beneath the dialog
    PropUnderlay bool

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnClose app.EventHandler
}

// DialogBase creates a new sp-dialog-base component
func DialogBase() *spectrumDialogBase {
    element := &spectrumDialogBase{
        PropDismissable: false,
        PropMode: "",
        PropOpen: false,
        PropResponsive: false,
        PropUnderlay: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the dialog is dismissable
func (c *spectrumDialogBase) Dismissable(dismissable bool) *spectrumDialogBase {
    c.PropDismissable = dismissable
    return c
}

func (c *spectrumDialogBase) SetDismissable() *spectrumDialogBase {
    return c.Dismissable(true)
}

// The display mode of the dialog
func (c *spectrumDialogBase) Mode(mode DialogBaseMode) *spectrumDialogBase {
    c.PropMode = mode
    return c
}

func (c *spectrumDialogBase) ModeFullscreen() *spectrumDialogBase {
    return c.Mode(DialogBaseModeFullscreen)
}
func (c *spectrumDialogBase) ModeFullscreentakeover() *spectrumDialogBase {
    return c.Mode(DialogBaseModeFullscreentakeover)
}
// Whether the dialog is open
func (c *spectrumDialogBase) Open(open bool) *spectrumDialogBase {
    c.PropOpen = open
    return c
}

func (c *spectrumDialogBase) SetOpen() *spectrumDialogBase {
    return c.Open(true)
}

// When set to true, fills screens smaller than 350px high and 400px wide with the full dialog
func (c *spectrumDialogBase) Responsive(responsive bool) *spectrumDialogBase {
    c.PropResponsive = responsive
    return c
}

func (c *spectrumDialogBase) SetResponsive() *spectrumDialogBase {
    return c.Responsive(true)
}

// Whether to show an underlay beneath the dialog
func (c *spectrumDialogBase) Underlay(underlay bool) *spectrumDialogBase {
    c.PropUnderlay = underlay
    return c
}

func (c *spectrumDialogBase) SetUnderlay() *spectrumDialogBase {
    return c.Underlay(true)
}


// Text sets the text content for the default slot
func (c *spectrumDialogBase) Text(text string) *spectrumDialogBase {
    c.PropText = text
    return c
}




// Announces that the dialog has been closed
func (c *spectrumDialogBase) OnClose(handler app.EventHandler) *spectrumDialogBase {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-dialog-base component
func (c *spectrumDialogBase) Render() app.UI {
    element := app.Elem("sp-dialog-base")

    // Set attributes
    if c.PropDismissable {
        element = element.Attr("dismissable", true)
    }
    if c.PropMode != "" {
        element = element.Attr("mode", string(c.PropMode))
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropResponsive {
        element = element.Attr("responsive", true)
    }
    if c.PropUnderlay {
        element = element.Attr("underlay", true)
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



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 