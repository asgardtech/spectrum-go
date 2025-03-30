package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DialogWrapperHeadlinevisibility represents the Controls visibility of the headline
type DialogWrapperHeadlinevisibility string

// DialogWrapperHeadlinevisibility values
const (
    DialogWrapperHeadlinevisibilityNone DialogWrapperHeadlinevisibility = "none"
)
// DialogWrapperMode represents the The display mode of the dialog
type DialogWrapperMode string

// DialogWrapperMode values
const (
    DialogWrapperModeFullscreen DialogWrapperMode = "fullscreen"
    DialogWrapperModeFullscreentakeover DialogWrapperMode = "fullscreenTakeover"
)
// DialogWrapperSize represents the Size of the dialog
type DialogWrapperSize string

// DialogWrapperSize values
const (
    DialogWrapperSizeS DialogWrapperSize = "s"
    DialogWrapperSizeM DialogWrapperSize = "m"
    DialogWrapperSizeL DialogWrapperSize = "l"
)

// spectrumDialogWrapper represents an sp-dialog-wrapper component
type spectrumDialogWrapper struct {
    app.Compo
    *styler[*spectrumDialogWrapper]

    // Properties
    // Label for the cancel button
    PropCancellabel string
    // Label for the confirm button
    PropConfirmlabel string
    // Label for the dismiss button
    PropDismisslabel string
    // Whether the dialog has a close button
    PropDismissable bool
    // Whether the dialog represents an error state
    PropError bool
    // Text content for the footer
    PropFooter string
    // Text for the dialog heading
    PropHeadline string
    // Controls visibility of the headline
    PropHeadlinevisibility DialogWrapperHeadlinevisibility
    // URL or content for hero image
    PropHero string
    // Accessible label for the hero content
    PropHerolabel string
    // The display mode of the dialog
    PropMode DialogWrapperMode
    // Whether to hide the divider between header and content
    PropNodivider bool
    // Whether the dialog is open
    PropOpen bool
    // When set to true, fills screens smaller than 350px high and 400px wide with the full dialog
    PropResponsive bool
    // Label for the secondary button
    PropSecondarylabel string
    // Size of the dialog
    PropSize DialogWrapperSize
    // Whether to show an underlay beneath the dialog
    PropUnderlay bool

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnCancel app.EventHandler
    PropOnClose app.EventHandler
    PropOnConfirm app.EventHandler
    PropOnSecondary app.EventHandler
}

// DialogWrapper creates a new sp-dialog-wrapper component
func DialogWrapper() *spectrumDialogWrapper {
    element := &spectrumDialogWrapper{
        PropCancellabel: "",
        PropConfirmlabel: "",
        PropDismisslabel: "Close",
        PropDismissable: false,
        PropError: false,
        PropFooter: "",
        PropHeadline: "",
        PropHeadlinevisibility: "",
        PropHero: "",
        PropHerolabel: "",
        PropMode: "",
        PropNodivider: false,
        PropOpen: false,
        PropResponsive: false,
        PropSecondarylabel: "",
        PropSize: "",
        PropUnderlay: false,
    }

    element.styler = newStyler(element)

    return element
}

// Label for the cancel button
func (c *spectrumDialogWrapper) Cancellabel(cancelLabel string) *spectrumDialogWrapper {
    c.PropCancellabel = cancelLabel
    return c
}

// Label for the confirm button
func (c *spectrumDialogWrapper) Confirmlabel(confirmLabel string) *spectrumDialogWrapper {
    c.PropConfirmlabel = confirmLabel
    return c
}

// Label for the dismiss button
func (c *spectrumDialogWrapper) Dismisslabel(dismissLabel string) *spectrumDialogWrapper {
    c.PropDismisslabel = dismissLabel
    return c
}

// Whether the dialog has a close button
func (c *spectrumDialogWrapper) Dismissable(dismissable bool) *spectrumDialogWrapper {
    c.PropDismissable = dismissable
    return c
}

func (c *spectrumDialogWrapper) SetDismissable() *spectrumDialogWrapper {
    return c.Dismissable(true)
}

// Whether the dialog represents an error state
func (c *spectrumDialogWrapper) Error(error bool) *spectrumDialogWrapper {
    c.PropError = error
    return c
}

func (c *spectrumDialogWrapper) SetError() *spectrumDialogWrapper {
    return c.Error(true)
}

// Text content for the footer
func (c *spectrumDialogWrapper) Footer(footer string) *spectrumDialogWrapper {
    c.PropFooter = footer
    return c
}

// Text for the dialog heading
func (c *spectrumDialogWrapper) Headline(headline string) *spectrumDialogWrapper {
    c.PropHeadline = headline
    return c
}

// Controls visibility of the headline
func (c *spectrumDialogWrapper) Headlinevisibility(headlineVisibility DialogWrapperHeadlinevisibility) *spectrumDialogWrapper {
    c.PropHeadlinevisibility = headlineVisibility
    return c
}

func (c *spectrumDialogWrapper) HeadlinevisibilityNone() *spectrumDialogWrapper {
    return c.Headlinevisibility(DialogWrapperHeadlinevisibilityNone)
}
// URL or content for hero image
func (c *spectrumDialogWrapper) Hero(hero string) *spectrumDialogWrapper {
    c.PropHero = hero
    return c
}

// Accessible label for the hero content
func (c *spectrumDialogWrapper) Herolabel(heroLabel string) *spectrumDialogWrapper {
    c.PropHerolabel = heroLabel
    return c
}

// The display mode of the dialog
func (c *spectrumDialogWrapper) Mode(mode DialogWrapperMode) *spectrumDialogWrapper {
    c.PropMode = mode
    return c
}

func (c *spectrumDialogWrapper) ModeFullscreen() *spectrumDialogWrapper {
    return c.Mode(DialogWrapperModeFullscreen)
}
func (c *spectrumDialogWrapper) ModeFullscreentakeover() *spectrumDialogWrapper {
    return c.Mode(DialogWrapperModeFullscreentakeover)
}
// Whether to hide the divider between header and content
func (c *spectrumDialogWrapper) Nodivider(noDivider bool) *spectrumDialogWrapper {
    c.PropNodivider = noDivider
    return c
}

func (c *spectrumDialogWrapper) SetNodivider() *spectrumDialogWrapper {
    return c.Nodivider(true)
}

// Whether the dialog is open
func (c *spectrumDialogWrapper) Open(open bool) *spectrumDialogWrapper {
    c.PropOpen = open
    return c
}

func (c *spectrumDialogWrapper) SetOpen() *spectrumDialogWrapper {
    return c.Open(true)
}

// When set to true, fills screens smaller than 350px high and 400px wide with the full dialog
func (c *spectrumDialogWrapper) Responsive(responsive bool) *spectrumDialogWrapper {
    c.PropResponsive = responsive
    return c
}

func (c *spectrumDialogWrapper) SetResponsive() *spectrumDialogWrapper {
    return c.Responsive(true)
}

// Label for the secondary button
func (c *spectrumDialogWrapper) Secondarylabel(secondaryLabel string) *spectrumDialogWrapper {
    c.PropSecondarylabel = secondaryLabel
    return c
}

// Size of the dialog
func (c *spectrumDialogWrapper) Size(size DialogWrapperSize) *spectrumDialogWrapper {
    c.PropSize = size
    return c
}

func (c *spectrumDialogWrapper) SizeS() *spectrumDialogWrapper {
    return c.Size(DialogWrapperSizeS)
}
func (c *spectrumDialogWrapper) SizeM() *spectrumDialogWrapper {
    return c.Size(DialogWrapperSizeM)
}
func (c *spectrumDialogWrapper) SizeL() *spectrumDialogWrapper {
    return c.Size(DialogWrapperSizeL)
}
// Whether to show an underlay beneath the dialog
func (c *spectrumDialogWrapper) Underlay(underlay bool) *spectrumDialogWrapper {
    c.PropUnderlay = underlay
    return c
}

func (c *spectrumDialogWrapper) SetUnderlay() *spectrumDialogWrapper {
    return c.Underlay(true)
}


// Text sets the text content for the default slot
func (c *spectrumDialogWrapper) Text(text string) *spectrumDialogWrapper {
    c.PropText = text
    return c
}




// Announces that the cancel button has been clicked
func (c *spectrumDialogWrapper) OnCancel(handler app.EventHandler) *spectrumDialogWrapper {
    c.PropOnCancel = handler

    return c
}

// Announces that the dialog has been closed
func (c *spectrumDialogWrapper) OnClose(handler app.EventHandler) *spectrumDialogWrapper {
    c.PropOnClose = handler

    return c
}

// Announces that the confirm button has been clicked
func (c *spectrumDialogWrapper) OnConfirm(handler app.EventHandler) *spectrumDialogWrapper {
    c.PropOnConfirm = handler

    return c
}

// Announces that the secondary button has been clicked
func (c *spectrumDialogWrapper) OnSecondary(handler app.EventHandler) *spectrumDialogWrapper {
    c.PropOnSecondary = handler

    return c
}


// Render renders the sp-dialog-wrapper component
func (c *spectrumDialogWrapper) Render() app.UI {
    element := app.Elem("sp-dialog-wrapper")

    // Set attributes
    if c.PropCancellabel != "" {
        element = element.Attr("cancelLabel", c.PropCancellabel)
    }
    if c.PropConfirmlabel != "" {
        element = element.Attr("confirmLabel", c.PropConfirmlabel)
    }
    if c.PropDismisslabel != "" {
        element = element.Attr("dismissLabel", c.PropDismisslabel)
    }
    if c.PropDismissable {
        element = element.Attr("dismissable", true)
    }
    if c.PropError {
        element = element.Attr("error", true)
    }
    if c.PropFooter != "" {
        element = element.Attr("footer", c.PropFooter)
    }
    if c.PropHeadline != "" {
        element = element.Attr("headline", c.PropHeadline)
    }
    if c.PropHeadlinevisibility != "" {
        element = element.Attr("headlineVisibility", string(c.PropHeadlinevisibility))
    }
    if c.PropHero != "" {
        element = element.Attr("hero", c.PropHero)
    }
    if c.PropHerolabel != "" {
        element = element.Attr("heroLabel", c.PropHerolabel)
    }
    if c.PropMode != "" {
        element = element.Attr("mode", string(c.PropMode))
    }
    if c.PropNodivider {
        element = element.Attr("noDivider", true)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropResponsive {
        element = element.Attr("responsive", true)
    }
    if c.PropSecondarylabel != "" {
        element = element.Attr("secondaryLabel", c.PropSecondarylabel)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropUnderlay {
        element = element.Attr("underlay", true)
    }

    // Add event handlers
    if c.PropOnCancel != nil {
        element = element.On("cancel", c.PropOnCancel)
    }
    if c.PropOnClose != nil {
        element = element.On("close", c.PropOnClose)
    }
    if c.PropOnConfirm != nil {
        element = element.On("confirm", c.PropOnConfirm)
    }
    if c.PropOnSecondary != nil {
        element = element.On("secondary", c.PropOnSecondary)
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