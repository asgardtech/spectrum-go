package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AlertDialogVariant represents the The variant defines the style and purpose of the alert dialog
type AlertDialogVariant string

// AlertDialogVariant values
const (
    AlertDialogVariantConfirmation AlertDialogVariant = "confirmation"
    AlertDialogVariantInformation AlertDialogVariant = "information"
    AlertDialogVariantWarning AlertDialogVariant = "warning"
    AlertDialogVariantError AlertDialogVariant = "error"
    AlertDialogVariantDestructive AlertDialogVariant = "destructive"
    AlertDialogVariantSecondary AlertDialogVariant = "secondary"
)

// spectrumAlertDialog represents an sp-alert-dialog component
type spectrumAlertDialog struct {
    app.Compo
    *styler[*spectrumAlertDialog]

    // Properties
    // The variant defines the style and purpose of the alert dialog
    PropVariant AlertDialogVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropHeadingSlot app.UI
    PropButtonSlot app.UI


    // Event handlers
    PropOnClose app.EventHandler
}

// AlertDialog creates a new sp-alert-dialog component
func AlertDialog() *spectrumAlertDialog {
    element := &spectrumAlertDialog{
        PropVariant: AlertDialogVariantConfirmation,
    }

    element.styler = newStyler(element)

    return element
}

// The variant defines the style and purpose of the alert dialog
func (c *spectrumAlertDialog) Variant(variant AlertDialogVariant) *spectrumAlertDialog {
    c.PropVariant = variant
    return c
}

func (c *spectrumAlertDialog) VariantConfirmation() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantConfirmation)
}
func (c *spectrumAlertDialog) VariantInformation() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantInformation)
}
func (c *spectrumAlertDialog) VariantWarning() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantWarning)
}
func (c *spectrumAlertDialog) VariantError() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantError)
}
func (c *spectrumAlertDialog) VariantDestructive() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantDestructive)
}
func (c *spectrumAlertDialog) VariantSecondary() *spectrumAlertDialog {
    return c.Variant(AlertDialogVariantSecondary)
}

// Text sets the text content for the default slot
func (c *spectrumAlertDialog) Text(text string) *spectrumAlertDialog {
    c.PropText = text
    return c
}


// Dialog heading or title
func (c *spectrumAlertDialog) Heading(content app.UI) *spectrumAlertDialog {
    c.PropHeadingSlot = content

    return c
}

// Buttons for user actions (multiple buttons can be slotted here)
func (c *spectrumAlertDialog) Button(content app.UI) *spectrumAlertDialog {
    c.PropButtonSlot = content

    return c
}



// Dispatched when the dialog should be closed
func (c *spectrumAlertDialog) OnClose(handler app.EventHandler) *spectrumAlertDialog {
    c.PropOnClose = handler

    return c
}


// Render renders the sp-alert-dialog component
func (c *spectrumAlertDialog) Render() app.UI {
    element := app.Elem("sp-alert-dialog")

    // Set attributes
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


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 