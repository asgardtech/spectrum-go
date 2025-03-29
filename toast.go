package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ToastVariant represents the visual style of a toast notification
type ToastVariant string

// Toast variants
const (
	ToastVariantNeutral  ToastVariant = ""
	ToastVariantNegative ToastVariant = "negative"
	ToastVariantPositive ToastVariant = "positive"
	ToastVariantInfo     ToastVariant = "info"
	ToastVariantError    ToastVariant = "error"
	ToastVariantWarning  ToastVariant = "warning"
)

// spectrumToast represents an sp-toast component
type spectrumToast struct {
	app.Compo

	// Properties
	PropIconLabel string
	PropOpen      bool
	PropTimeout   int
	PropVariant   ToastVariant

	// Content
	PropContent string
	PropAction  app.UI

	// Event handlers
	PropOnClose app.EventHandler
}

// Toast creates a new toast component
func Toast() *spectrumToast {
	return &spectrumToast{}
}

// IconLabel sets the accessible label for the icon element
func (t *spectrumToast) IconLabel(label string) *spectrumToast {
	t.PropIconLabel = label
	return t
}

// Open sets whether the toast is visible
func (t *spectrumToast) Open(open bool) *spectrumToast {
	t.PropOpen = open
	return t
}

// Timeout sets the number of milliseconds before the toast automatically dismisses itself
func (t *spectrumToast) Timeout(ms int) *spectrumToast {
	t.PropTimeout = ms
	return t
}

// Variant sets the visual style of the toast
func (t *spectrumToast) Variant(variant ToastVariant) *spectrumToast {
	t.PropVariant = variant
	return t
}

// Negative sets the variant to "negative"
func (t *spectrumToast) Negative() *spectrumToast {
	return t.Variant(ToastVariantNegative)
}

// Positive sets the variant to "positive"
func (t *spectrumToast) Positive() *spectrumToast {
	return t.Variant(ToastVariantPositive)
}

// Info sets the variant to "info"
func (t *spectrumToast) Info() *spectrumToast {
	return t.Variant(ToastVariantInfo)
}

// Error sets the variant to "error"
func (t *spectrumToast) Error() *spectrumToast {
	return t.Variant(ToastVariantError)
}

// Warning sets the variant to "warning"
func (t *spectrumToast) Warning() *spectrumToast {
	return t.Variant(ToastVariantWarning)
}

// Content sets the text content of the toast
func (t *spectrumToast) Content(content string) *spectrumToast {
	t.PropContent = content
	return t
}

// Action sets the action button in the toast
func (t *spectrumToast) Action(action app.UI) *spectrumToast {
	t.PropAction = action
	return t
}

// OnClose sets the close event handler
func (t *spectrumToast) OnClose(handler app.EventHandler) *spectrumToast {
	t.PropOnClose = handler
	return t
}

// Render renders the toast component
func (t *spectrumToast) Render() app.UI {
	toast := app.Elem("sp-toast")

	// Set attributes
	if t.PropIconLabel != "" {
		toast = toast.Attr("icon-label", t.PropIconLabel)
	}
	if t.PropOpen {
		toast = toast.Attr("open", true)
	}
	if t.PropTimeout > 0 {
		toast = toast.Attr("timeout", t.PropTimeout)
	}
	if t.PropVariant != ToastVariantNeutral {
		toast = toast.Attr("variant", string(t.PropVariant))
	}

	// Add event handlers
	if t.PropOnClose != nil {
		toast = toast.On("close", t.PropOnClose)
	}

	// Add content and action
	elements := []app.UI{}

	// Add text content if provided
	if t.PropContent != "" {
		elements = append(elements, app.Text(t.PropContent))
	}

	// Add action button if provided
	if t.PropAction != nil {
		// If action doesn't explicitly set its slot, set it to "action"
		if actionWithSlot, ok := t.PropAction.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, actionWithSlot.Slot("action"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "action").
				Body(t.PropAction))
		}
	}

	if len(elements) > 0 {
		toast = toast.Body(elements...)
	}

	return toast
}
