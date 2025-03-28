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

// SpectrumToast represents an sp-toast component
type SpectrumToast struct {
	app.Compo

	// Properties
	iconLabel string
	open      bool
	timeout   int
	variant   ToastVariant

	// Content
	content string
	action  app.UI

	// Event handlers
	onClose app.EventHandler
}

// Toast creates a new toast component
func Toast() *SpectrumToast {
	return &SpectrumToast{}
}

// IconLabel sets the accessible label for the icon element
func (t *SpectrumToast) IconLabel(label string) *SpectrumToast {
	t.iconLabel = label
	return t
}

// Open sets whether the toast is visible
func (t *SpectrumToast) Open(open bool) *SpectrumToast {
	t.open = open
	return t
}

// Timeout sets the number of milliseconds before the toast automatically dismisses itself
func (t *SpectrumToast) Timeout(ms int) *SpectrumToast {
	t.timeout = ms
	return t
}

// Variant sets the visual style of the toast
func (t *SpectrumToast) Variant(variant ToastVariant) *SpectrumToast {
	t.variant = variant
	return t
}

// Negative sets the variant to "negative"
func (t *SpectrumToast) Negative() *SpectrumToast {
	return t.Variant(ToastVariantNegative)
}

// Positive sets the variant to "positive"
func (t *SpectrumToast) Positive() *SpectrumToast {
	return t.Variant(ToastVariantPositive)
}

// Info sets the variant to "info"
func (t *SpectrumToast) Info() *SpectrumToast {
	return t.Variant(ToastVariantInfo)
}

// Error sets the variant to "error"
func (t *SpectrumToast) Error() *SpectrumToast {
	return t.Variant(ToastVariantError)
}

// Warning sets the variant to "warning"
func (t *SpectrumToast) Warning() *SpectrumToast {
	return t.Variant(ToastVariantWarning)
}

// Content sets the text content of the toast
func (t *SpectrumToast) Content(content string) *SpectrumToast {
	t.content = content
	return t
}

// Action sets the action button in the toast
func (t *SpectrumToast) Action(action app.UI) *SpectrumToast {
	t.action = action
	return t
}

// OnClose sets the close event handler
func (t *SpectrumToast) OnClose(handler app.EventHandler) *SpectrumToast {
	t.onClose = handler
	return t
}

// Render renders the toast component
func (t *SpectrumToast) Render() app.UI {
	toast := app.Elem("sp-toast")

	// Set attributes
	if t.iconLabel != "" {
		toast = toast.Attr("icon-label", t.iconLabel)
	}
	if t.open {
		toast = toast.Attr("open", true)
	}
	if t.timeout > 0 {
		toast = toast.Attr("timeout", t.timeout)
	}
	if t.variant != ToastVariantNeutral {
		toast = toast.Attr("variant", string(t.variant))
	}

	// Add event handlers
	if t.onClose != nil {
		toast = toast.On("close", t.onClose)
	}

	// Add content and action
	elements := []app.UI{}

	// Add text content if provided
	if t.content != "" {
		elements = append(elements, app.Text(t.content))
	}

	// Add action button if provided
	if t.action != nil {
		// If action doesn't explicitly set its slot, set it to "action"
		if actionWithSlot, ok := t.action.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, actionWithSlot.Slot("action"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "action").
				Body(t.action))
		}
	}

	if len(elements) > 0 {
		toast = toast.Body(elements...)
	}

	return toast
}
