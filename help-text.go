package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// HelpTextVariant represents the visual variant of a help text
type HelpTextVariant string

// Help text variants
const (
	HelpTextVariantNeutral  HelpTextVariant = "neutral"
	HelpTextVariantNegative HelpTextVariant = "negative"
)

// HelpTextSize represents the visual size of a help text
type HelpTextSize string

// Help text sizes
const (
	HelpTextSizeS  HelpTextSize = "s"
	HelpTextSizeM  HelpTextSize = "m"
	HelpTextSizeL  HelpTextSize = "l"
	HelpTextSizeXL HelpTextSize = "xl"
)

// SpectrumHelpText represents an sp-help-text component
type spectrumHelpText struct {
	app.Compo

	// Properties
	PropVariant   HelpTextVariant
	PropSize      HelpTextSize
	PropIcon      bool
	PropDisabled  bool
	PropSlot      string
	PropContent   string
	PropInnerHTML string
	PropChild     app.UI
}

// HelpText creates a new help text component
func HelpText() *spectrumHelpText {
	return &spectrumHelpText{
		PropVariant: HelpTextVariantNeutral,
		PropSize:    HelpTextSizeM,
	}
}

// Variant sets the visual variant of the help text
func (h *spectrumHelpText) Variant(variant HelpTextVariant) *spectrumHelpText {
	h.PropVariant = variant
	return h
}

// Size sets the visual size of the help text
func (h *spectrumHelpText) Size(size HelpTextSize) *spectrumHelpText {
	h.PropSize = size
	return h
}

// Icon enables the icon for negative variant
func (h *spectrumHelpText) Icon(icon bool) *spectrumHelpText {
	h.PropIcon = icon
	return h
}

// Disabled sets the disabled state
func (h *spectrumHelpText) Disabled(disabled bool) *spectrumHelpText {
	h.PropDisabled = disabled
	return h
}

// Slot sets the slot attribute
func (h *spectrumHelpText) Slot(slot string) *spectrumHelpText {
	h.PropSlot = slot
	return h
}

// Content sets the text content of the help text
func (h *spectrumHelpText) Content(content string) *spectrumHelpText {
	h.PropContent = content
	return h
}

// InnerHTML sets the inner HTML of the help text
func (h *spectrumHelpText) InnerHTML(html string) *spectrumHelpText {
	h.PropInnerHTML = html
	return h
}

// Child sets a UI element as the child of the help text
func (h *spectrumHelpText) Child(child app.UI) *spectrumHelpText {
	h.PropChild = child
	return h
}

// Negative is a convenience method to set the variant to negative
func (h *spectrumHelpText) Negative() *spectrumHelpText {
	h.PropVariant = HelpTextVariantNegative
	return h
}

// Render renders the help text component
func (h *spectrumHelpText) Render() app.UI {
	helpText := app.Elem("sp-help-text").
		Attr("variant", string(h.PropVariant)).
		Attr("size", string(h.PropSize)).
		Attr("icon", h.PropIcon).
		Attr("disabled", h.PropDisabled)

	if h.PropSlot != "" {
		helpText = helpText.Attr("slot", h.PropSlot)
	}

	// Handle content in the right order of precedence
	if h.PropInnerHTML != "" {
		helpText = helpText.Body(app.Raw(h.PropInnerHTML))
	} else if h.PropChild != nil {
		helpText = helpText.Body(h.PropChild)
	} else if h.PropContent != "" {
		helpText = helpText.Body(app.Text(h.PropContent))
	}

	return helpText
}
