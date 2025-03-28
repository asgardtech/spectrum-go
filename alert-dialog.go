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

// SpectrumAlertDialog represents an sp-alert-dialog component
type SpectrumAlertDialog struct {
	app.Compo

	// Properties
	variant AlertDialogVariant

	// Content slots
	heading         app.UI
	content         app.UI
	buttonSlots     []app.UI
	confirmButton   app.UI
	cancelButton    app.UI
	secondaryButton app.UI

	// Event handlers
	onClose app.EventHandler
}

// AlertDialog creates a new alert dialog component
func AlertDialog() *SpectrumAlertDialog {
	return &SpectrumAlertDialog{
		variant: AlertDialogVariantConfirmation, // Default variant
	}
}

// Variant sets the visual style of the alert dialog
func (ad *SpectrumAlertDialog) Variant(variant AlertDialogVariant) *SpectrumAlertDialog {
	ad.variant = variant
	return ad
}

// Confirmation sets the variant to "confirmation"
func (ad *SpectrumAlertDialog) Confirmation() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantConfirmation)
}

// Information sets the variant to "information"
func (ad *SpectrumAlertDialog) Information() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantInformation)
}

// Warning sets the variant to "warning"
func (ad *SpectrumAlertDialog) Warning() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantWarning)
}

// Error sets the variant to "error"
func (ad *SpectrumAlertDialog) Error() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantError)
}

// Destructive sets the variant to "destructive"
func (ad *SpectrumAlertDialog) Destructive() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantDestructive)
}

// Secondary sets the variant to "secondary"
func (ad *SpectrumAlertDialog) Secondary() *SpectrumAlertDialog {
	return ad.Variant(AlertDialogVariantSecondary)
}

// Heading sets the heading content in the heading slot
func (ad *SpectrumAlertDialog) Heading(heading app.UI) *SpectrumAlertDialog {
	ad.heading = heading
	return ad
}

// HeadingText sets a text heading in the heading slot
func (ad *SpectrumAlertDialog) HeadingText(text string) *SpectrumAlertDialog {
	ad.heading = app.H2().Text(text)
	return ad
}

// Content sets the main content of the alert dialog
func (ad *SpectrumAlertDialog) Content(content app.UI) *SpectrumAlertDialog {
	ad.content = content
	return ad
}

// ContentText sets text content in the alert dialog
func (ad *SpectrumAlertDialog) ContentText(text string) *SpectrumAlertDialog {
	ad.content = app.Text(text)
	return ad
}

// AddButton adds a button to the alert dialog's button slot
func (ad *SpectrumAlertDialog) AddButton(button app.UI) *SpectrumAlertDialog {
	// If the button implements a Slot method, use it, otherwise wrap the button
	if buttonWithSlot, ok := button.(interface{ Slot(string) app.UI }); ok {
		ad.buttonSlots = append(ad.buttonSlots, buttonWithSlot.Slot("button"))
	} else {
		ad.buttonSlots = append(ad.buttonSlots, app.Elem("div").
			Attr("slot", "button").
			Body(button))
	}
	return ad
}

// ConfirmButton sets the confirm button
func (ad *SpectrumAlertDialog) ConfirmButton(button app.UI) *SpectrumAlertDialog {
	ad.confirmButton = button
	return ad
}

// CancelButton sets the cancel button
func (ad *SpectrumAlertDialog) CancelButton(button app.UI) *SpectrumAlertDialog {
	ad.cancelButton = button
	return ad
}

// SecondaryButton sets the secondary button (for secondary variant)
func (ad *SpectrumAlertDialog) SecondaryButton(button app.UI) *SpectrumAlertDialog {
	ad.secondaryButton = button
	return ad
}

// OnClose sets the close event handler
func (ad *SpectrumAlertDialog) OnClose(handler app.EventHandler) *SpectrumAlertDialog {
	ad.onClose = handler
	return ad
}

// Render renders the alert dialog component
func (ad *SpectrumAlertDialog) Render() app.UI {
	alertDialog := app.Elem("sp-alert-dialog")

	// Set attributes
	if ad.variant != "" {
		alertDialog = alertDialog.Attr("variant", string(ad.variant))
	}

	// Add event handlers
	if ad.onClose != nil {
		alertDialog = alertDialog.On("close", ad.onClose)
	}

	// Add content elements
	elements := []app.UI{}

	// Add heading if provided
	if ad.heading != nil {
		if headingWithSlot, ok := ad.heading.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, headingWithSlot.Slot("heading"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "heading").
				Body(ad.heading))
		}
	}

	// Add main content if provided
	if ad.content != nil {
		elements = append(elements, ad.content)
	}

	// Add buttons from buttonSlots array
	if len(ad.buttonSlots) > 0 {
		elements = append(elements, ad.buttonSlots...)
	}

	// Add specific buttons if provided and not already added via buttonSlots
	if ad.confirmButton != nil {
		if buttonWithSlot, ok := ad.confirmButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "confirmButton").
				Body(ad.confirmButton)
			elements = append(elements, button)
		}
	}

	if ad.cancelButton != nil {
		if buttonWithSlot, ok := ad.cancelButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "cancelButton").
				Body(ad.cancelButton)
			elements = append(elements, button)
		}
	}

	if ad.secondaryButton != nil {
		if buttonWithSlot, ok := ad.secondaryButton.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, buttonWithSlot.Slot("button"))
		} else {
			button := app.Elem("div").
				Attr("slot", "button").
				Attr("id", "secondaryButton").
				Body(ad.secondaryButton)
			elements = append(elements, button)
		}
	}

	if len(elements) > 0 {
		alertDialog = alertDialog.Body(elements...)
	}

	return alertDialog
}
