package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AlertBannerVariant represents the visual variant of an alert banner
type AlertBannerVariant string

// Alert banner variants
const (
	AlertBannerVariantInfo     AlertBannerVariant = "info"
	AlertBannerVariantNegative AlertBannerVariant = "negative"
)

// spectrumAlertBanner represents an sp-alert-banner component
type spectrumAlertBanner struct {
	app.Compo

	// Properties
	PropDismissible bool
	PropOpen        bool
	PropVariant     AlertBannerVariant

	// Event handlers
	PropOnClose app.EventHandler

	// Content
	PropContent   string
	PropInnerHTML string
	PropAction    app.UI
	PropChild     app.UI
}

// AlertBanner creates a new alert banner component
func AlertBanner() *spectrumAlertBanner {
	return &spectrumAlertBanner{}
}

// Dismissible sets whether the alert banner should include a close button
func (a *spectrumAlertBanner) Dismissible(dismissible bool) *spectrumAlertBanner {
	a.PropDismissible = dismissible
	return a
}

// Open sets whether the alert banner is visible
func (a *spectrumAlertBanner) Open(open bool) *spectrumAlertBanner {
	a.PropOpen = open
	return a
}

// Variant sets the visual variant of the alert banner
func (a *spectrumAlertBanner) Variant(variant AlertBannerVariant) *spectrumAlertBanner {
	a.PropVariant = variant
	return a
}

// OnClose sets the close event handler
func (a *spectrumAlertBanner) OnClose(handler app.EventHandler) *spectrumAlertBanner {
	a.PropOnClose = handler
	return a
}

// Content sets the text content of the alert banner
func (a *spectrumAlertBanner) Content(content string) *spectrumAlertBanner {
	a.PropContent = content
	return a
}

// InnerHTML sets the inner HTML of the alert banner
func (a *spectrumAlertBanner) InnerHTML(html string) *spectrumAlertBanner {
	a.PropInnerHTML = html
	return a
}

// Action sets the action button UI element in the action slot
func (a *spectrumAlertBanner) Action(action app.UI) *spectrumAlertBanner {
	a.PropAction = action
	return a
}

// Child sets a UI element as the child of the alert banner
func (a *spectrumAlertBanner) Child(child app.UI) *spectrumAlertBanner {
	a.PropChild = child
	return a
}

// Render renders the alert banner component
func (a *spectrumAlertBanner) Render() app.UI {
	alertBanner := app.Elem("sp-alert-banner").
		Attr("dismissible", a.PropDismissible).
		Attr("open", a.PropOpen)

	if a.PropVariant != "" {
		alertBanner = alertBanner.Attr("variant", string(a.PropVariant))
	}

	// Add event handlers
	if a.PropOnClose != nil {
		alertBanner = alertBanner.On("close", a.PropOnClose)
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child
	if a.PropChild != nil {
		elements = append(elements, a.PropChild)
	} else if a.PropInnerHTML != "" {
		elements = append(elements, app.Raw(a.PropInnerHTML))
	} else if a.PropContent != "" {
		elements = append(elements, app.Text(a.PropContent))
	}

	// Add action button if provided
	if a.PropAction != nil {
		actionElem := a.PropAction
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
