package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ToastVariant represents the Applies specific styling to the toast
type ToastVariant string

// ToastVariant values
const (
    ToastVariantNegative ToastVariant = "negative"
    ToastVariantPositive ToastVariant = "positive"
    ToastVariantInfo ToastVariant = "info"
    ToastVariantError ToastVariant = "error"
    ToastVariantWarning ToastVariant = "warning"
)

// spectrumToast represents an sp-toast component
type spectrumToast struct {
    app.Compo
    *styler[*spectrumToast]

    // Properties
    // Sets the label attribute on the icon element, providing a text alternative for accessibility
    PropIconLabel string
    // Indicates whether the toast is visible or hidden
    PropOpen bool
    // The number of milliseconds before the toast automatically dismisses itself
    PropTimeout float64
    // Applies specific styling to the toast
    PropVariant ToastVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropActionSlot app.UI


    // Event handlers
    PropOnClose app.EventHandler
}

// Toast creates a new sp-toast component
func Toast() *spectrumToast {
    element := &spectrumToast{
        PropOpen: false,
        PropTimeout: 0,
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Sets the label attribute on the icon element, providing a text alternative for accessibility
func (c *spectrumToast) IconLabel(iconLabel string) *spectrumToast {
    c.PropIconLabel = iconLabel
    return c
}

// Indicates whether the toast is visible or hidden
func (c *spectrumToast) Open(open bool) *spectrumToast {
    c.PropOpen = open
    return c
}

func (c *spectrumToast) SetOpen() *spectrumToast {
    return c.Open(true)
}

// The number of milliseconds before the toast automatically dismisses itself
func (c *spectrumToast) Timeout(timeout float64) *spectrumToast {
    c.PropTimeout = timeout
    return c
}

// Applies specific styling to the toast
func (c *spectrumToast) Variant(variant ToastVariant) *spectrumToast {
    c.PropVariant = variant
    return c
}

func (c *spectrumToast) VariantNegative() *spectrumToast {
    return c.Variant(ToastVariantNegative)
}
func (c *spectrumToast) VariantPositive() *spectrumToast {
    return c.Variant(ToastVariantPositive)
}
func (c *spectrumToast) VariantInfo() *spectrumToast {
    return c.Variant(ToastVariantInfo)
}
func (c *spectrumToast) VariantError() *spectrumToast {
    return c.Variant(ToastVariantError)
}
func (c *spectrumToast) VariantWarning() *spectrumToast {
    return c.Variant(ToastVariantWarning)
}

// Text sets the text content for the default slot
func (c *spectrumToast) Text(text string) *spectrumToast {
    c.PropText = text
    return c
}


// Button element surfacing an action in the toast
func (c *spectrumToast) Action(content app.UI) *spectrumToast {
    c.PropActionSlot = content

    return c
}



// Announces that the toast has been closed by the user or by its timeout
func (c *spectrumToast) OnClose(handler app.EventHandler) *spectrumToast {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-toast component
func (c *spectrumToast) Render() app.UI {
    element := app.Elem("sp-toast")

    // Set attributes
    if c.PropIconLabel != "" {
        element = element.Attr("icon-label", c.PropIconLabel)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropTimeout != 0 {
        element = element.Attr("timeout", c.PropTimeout)
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