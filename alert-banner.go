package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AlertBannerVariant represents the visual variant of an alert banner
type AlertBannerVariant string

// Alert banner variants
const (
	AlertBannerVariantInfo     AlertBannerVariant = "info"
	AlertBannerVariantNegative AlertBannerVariant = "negative"
)

// SpectrumAlertBanner represents an sp-alert-banner component
type SpectrumAlertBanner struct {
	app.Compo

	// Properties
	dismissible bool
	open        bool
	variant     AlertBannerVariant

	// Event handlers
	onClose app.EventHandler

	// Content
	content   string
	innerHTML string
	action    app.UI
	child     app.UI
}

// AlertBanner creates a new alert banner component
func AlertBanner() *SpectrumAlertBanner {
	return &SpectrumAlertBanner{}
}

// Dismissible sets whether the alert banner should include a close button
func (a *SpectrumAlertBanner) Dismissible(dismissible bool) *SpectrumAlertBanner {
	a.dismissible = dismissible
	return a
}

// Open sets whether the alert banner is visible
func (a *SpectrumAlertBanner) Open(open bool) *SpectrumAlertBanner {
	a.open = open
	return a
}

// Variant sets the visual variant of the alert banner
func (a *SpectrumAlertBanner) Variant(variant AlertBannerVariant) *SpectrumAlertBanner {
	a.variant = variant
	return a
}

// OnClose sets the close event handler
func (a *SpectrumAlertBanner) OnClose(handler app.EventHandler) *SpectrumAlertBanner {
	a.onClose = handler
	return a
}

// Content sets the text content of the alert banner
func (a *SpectrumAlertBanner) Content(content string) *SpectrumAlertBanner {
	a.content = content
	return a
}

// InnerHTML sets the inner HTML of the alert banner
func (a *SpectrumAlertBanner) InnerHTML(html string) *SpectrumAlertBanner {
	a.innerHTML = html
	return a
}

// Action sets the action button UI element in the action slot
func (a *SpectrumAlertBanner) Action(action app.UI) *SpectrumAlertBanner {
	a.action = action
	return a
}

// Child sets a UI element as the child of the alert banner
func (a *SpectrumAlertBanner) Child(child app.UI) *SpectrumAlertBanner {
	a.child = child
	return a
}

// Render renders the alert banner component
func (a *SpectrumAlertBanner) Render() app.UI {
	alertBanner := app.Elem("sp-alert-banner").
		Attr("dismissible", a.dismissible).
		Attr("open", a.open)

	if a.variant != "" {
		alertBanner = alertBanner.Attr("variant", string(a.variant))
	}

	// Add event handlers
	if a.onClose != nil {
		alertBanner = alertBanner.On("close", a.onClose)
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child
	if a.child != nil {
		elements = append(elements, a.child)
	} else if a.innerHTML != "" {
		elements = append(elements, app.Raw(a.innerHTML))
	} else if a.content != "" {
		elements = append(elements, app.Text(a.content))
	}

	// Add action button if provided
	if a.action != nil {
		actionElem := a.action
		if actionWithSlot, ok := actionElem.(interface{ Slot(string) app.UI }); ok {
			actionElem = actionWithSlot.Slot("action")
		} else {
			actionElem = app.Elem("div").
				Attr("slot", "action").
				Body(actionElem)
		}
		elements = append(elements, actionElem)
	}

	// Add all elements to the alert banner
	if len(elements) > 0 {
		alertBanner = alertBanner.Body(elements...)
	}

	return alertBanner
}
