package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AlertDialogVariant represents the visual style of an alert dialog
type AlertDialogVariant string

// Alert dialog variants
const (
	AlertDialogVariantConfirmation AlertDialogVariant = "confirmation"
	AlertDialogVariantInformation  AlertDialogVariant = "information"
	AlertDialogVariantWarning      AlertDialogVariant = "warning"
	AlertDialogVariantError        AlertDialogVariant = "error"
	AlertDialogVariantDestructive  AlertDialogVariant = "destructive"
	AlertDialogVariantSecondary    AlertDialogVariant = "secondary"
)

// spectrumAlertDialog represents an sp-alert-dialog component
type spectrumAlertDialog struct {
	app.Compo

	// Properties
	PropVariant AlertDialogVariant

	// Content slots
	PropHeading         app.UI
	PropContent         app.UI
	PropButtonSlots     []app.UI
	PropConfirmButton   app.UI
	PropCancelButton    app.UI
	PropSecondaryButton app.UI

	// Event handlers
	PropOnClose app.EventHandler
}

// AlertDialog creates a new alert dialog component
func AlertDialog() *spectrumAlertDialog {
	return &spectrumAlertDialog{
		PropVariant: AlertDialogVariantConfirmation, // Default variant
	}
}

// Variant sets the visual style of the alert dialog
func (ad *spectrumAlertDialog) Variant(variant AlertDialogVariant) *spectrumAlertDialog {
	ad.PropVariant = variant
	return ad
}

// Confirmation sets the variant to "confirmation"
func (ad *spectrumAlertDialog) Confirmation() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantConfirmation)
}

// Information sets the variant to "information"
func (ad *spectrumAlertDialog) Information() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantInformation)
}

// Warning sets the variant to "warning"
func (ad *spectrumAlertDialog) Warning() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantWarning)
}

// Error sets the variant to "error"
func (ad *spectrumAlertDialog) Error() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantError)
}

// Destructive sets the variant to "destructive"
func (ad *spectrumAlertDialog) Destructive() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantDestructive)
}

// Secondary sets the variant to "secondary"
func (ad *spectrumAlertDialog) Secondary() *spectrumAlertDialog {
	return ad.Variant(AlertDialogVariantSecondary)
}

// Heading sets the heading content in the heading slot
func (ad *spectrumAlertDialog) Heading(heading app.UI) *spectrumAlertDialog {
	ad.PropHeading = heading
	return ad
}

// HeadingText sets a text heading in the heading slot
func (ad *spectrumAlertDialog) HeadingText(text string) *spectrumAlertDialog {
	ad.PropHeading = app.H2().Text(text)
	return ad
}

// Content sets the main content of the alert dialog
func (ad *spectrumAlertDialog) Content(content app.UI) *spectrumAlertDialog {
	ad.PropContent = content
	return ad
}

// ContentText sets text content in the alert dialog
func (ad *spectrumAlertDialog) ContentText(text string) *spectrumAlertDialog {
	ad.PropContent = app.Text(text)
	return ad
}

// AddButton adds a button to the alert dialog's button slot
func (ad *spectrumAlertDialog) AddButton(button app.UI) *spectrumAlertDialog {
	// If the button implements a Slot method, use it, otherwise wrap the button
	if buttonWithSlot, ok := button.(interface{ Slot(string) app.UI }); ok {
		ad.PropButtonSlots = append(ad.PropButtonSlots, buttonWithSlot.Slot("button"))
	} else {
		ad.PropButtonSlots = append(ad.PropButtonSlots, app.Elem("div").
			Attr("slot", "button").
			Body(button))
	}
	return ad
}

// ConfirmButton sets the confirm button
func (ad *spectrumAlertDialog) ConfirmButton(button app.UI) *spectrumAlertDialog {
	ad.PropConfirmButton = button
	return ad
}

// CancelButton sets the cancel button
func (ad *spectrumAlertDialog) CancelButton(button app.UI) *spectrumAlertDialog {
	ad.PropCancelButton = button
	return ad
}

// SecondaryButton sets the secondary button (for secondary variant)
func (ad *spectrumAlertDialog) SecondaryButton(button app.UI) *spectrumAlertDialog {
	ad.PropSecondaryButton = button
	return ad
}

// OnClose sets the close event handler
func (ad *spectrumAlertDialog) OnClose(handler app.EventHandler) *spectrumAlertDialog {
	ad.PropOnClose = handler
	return ad
}

// Render renders the alert dialog component
func (ad *spectrumAlertDialog) Render() app.UI {
	alertDialog := app.Elem("sp-alert-dialog")

	// Set attributes
	if ad.PropVariant != "" {
		alertDialog = alertDialog.Attr("variant", string(ad.PropVariant))
	}

	// Add event handlers
	if ad.PropOnClose != nil {
		alertDialog = alertDialog.On("close", ad.PropOnClose)
	}

	// Add content elements
	elements := []app.UI{}

	// Add heading if provided
	if ad.PropHeading != nil {
		if headingWithSlot, ok := ad.PropHeading.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, headingWithSlot.Slot("heading"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "heading").
				Body(ad.PropHeading))
		}
	}

	// Add main content if provided
	if ad.PropContent != nil {
		elements = append(elements, ad.PropContent)
	}

	// Add buttons from buttonSlots array
	if len(ad.PropButtonSlots) > 0 {
		elements = append(elements, ad.PropButtonSlots...)
	}

	// Add specific buttons if provided and not already added via buttonSlots
	if ad.PropConfirmButton != nil {
		if buttonWithSlot, ok := ad.PropConfirmButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "confirmButton").
				Body(ad.PropConfirmButton)
			elements = append(elements, button)
		}
	}

	if ad.PropCancelButton != nil {
		if buttonWithSlot, ok := ad.PropCancelButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "cancelButton").
				Body(ad.PropCancelButton)
			elements = append(elements, button)
		}
	}

	if ad.PropSecondaryButton != nil {
		if buttonWithSlot, ok := ad.PropSecondaryButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "secondaryButton").
				Body(ad.PropSecondaryButton)
			elements = append(elements, button)
		}
	}

	if len(elements) > 0 {
		alertDialog = alertDialog.Body(elements...)
	}

	return alertDialog
}
