package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AlertBannerVariant represents the The variant applies specific styling for different alert types
type AlertBannerVariant string

// AlertBannerVariant values
const (
    AlertBannerVariantInfo AlertBannerVariant = "info"
    AlertBannerVariantNegative AlertBannerVariant = "negative"
)

// spectrumAlertBanner represents an sp-alert-banner component
type spectrumAlertBanner struct {
    app.Compo
    *styler[*spectrumAlertBanner]

    // Properties
    // Whether to include an icon-only close button to dismiss the alert banner
    PropDismissible bool
    // Controls the display of the alert banner
    PropOpen bool
    // The variant applies specific styling for different alert types
    PropVariant AlertBannerVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropActionSlot app.UI


    // Event handlers
    PropOnClose app.EventHandler
}

// AlertBanner creates a new sp-alert-banner component
func AlertBanner() *spectrumAlertBanner {
    element := &spectrumAlertBanner{
        PropDismissible: false,
        PropOpen: false,
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Whether to include an icon-only close button to dismiss the alert banner
func (c *spectrumAlertBanner) Dismissible(dismissible bool) *spectrumAlertBanner {
    c.PropDismissible = dismissible
    return c
}

func (c *spectrumAlertBanner) SetDismissible() *spectrumAlertBanner {
    return c.Dismissible(true)
}

// Controls the display of the alert banner
func (c *spectrumAlertBanner) Open(open bool) *spectrumAlertBanner {
    c.PropOpen = open
    return c
}

func (c *spectrumAlertBanner) SetOpen() *spectrumAlertBanner {
    return c.Open(true)
}

// The variant applies specific styling for different alert types
func (c *spectrumAlertBanner) Variant(variant AlertBannerVariant) *spectrumAlertBanner {
    c.PropVariant = variant
    return c
}

func (c *spectrumAlertBanner) VariantInfo() *spectrumAlertBanner {
    return c.Variant(AlertBannerVariantInfo)
}
func (c *spectrumAlertBanner) VariantNegative() *spectrumAlertBanner {
    return c.Variant(AlertBannerVariantNegative)
}

// Text sets the text content for the default slot
func (c *spectrumAlertBanner) Text(text string) *spectrumAlertBanner {
    c.PropText = text
    return c
}


// Slot for the button element that surfaces the contextual action a user can take
func (c *spectrumAlertBanner) Action(content app.UI) *spectrumAlertBanner {
    c.PropActionSlot = content

    return c
}



// Announces the alert banner has been closed
func (c *spectrumAlertBanner) OnClose(handler app.EventHandler) *spectrumAlertBanner {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-alert-banner component
func (c *spectrumAlertBanner) Render() app.UI {
    element := app.Elem("sp-alert-banner")

    // Set attributes
    if c.PropDismissible {
        element = element.Attr("dismissible", true)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
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

    // Add action slot
    if c.PropActionSlot != nil {
        slotElem := c.PropActionSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("action")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "action").
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