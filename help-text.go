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
type SpectrumHelpText struct {
	app.Compo

	// Properties
	variant   HelpTextVariant
	size      HelpTextSize
	icon      bool
	disabled  bool
	slot      string
	content   string
	innerHTML string
	child     app.UI
}

// HelpText creates a new help text component
func HelpText() *SpectrumHelpText {
	return &SpectrumHelpText{
		variant: HelpTextVariantNeutral,
		size:    HelpTextSizeM,
	}
}

// Variant sets the visual variant of the help text
func (h *SpectrumHelpText) Variant(variant HelpTextVariant) *SpectrumHelpText {
	h.variant = variant
	return h
}

// Size sets the visual size of the help text
func (h *SpectrumHelpText) Size(size HelpTextSize) *SpectrumHelpText {
	h.size = size
	return h
}

// Icon enables the icon for negative variant
func (h *SpectrumHelpText) Icon(icon bool) *SpectrumHelpText {
	h.icon = icon
	return h
}

// Disabled sets the disabled state
func (h *SpectrumHelpText) Disabled(disabled bool) *SpectrumHelpText {
	h.disabled = disabled
	return h
}

// Slot sets the slot attribute
func (h *SpectrumHelpText) Slot(slot string) *SpectrumHelpText {
	h.slot = slot
	return h
}

// Content sets the text content of the help text
func (h *SpectrumHelpText) Content(content string) *SpectrumHelpText {
	h.content = content
	return h
}

// InnerHTML sets the inner HTML of the help text
func (h *SpectrumHelpText) InnerHTML(html string) *SpectrumHelpText {
	h.innerHTML = html
	return h
}

// Child sets a UI element as the child of the help text
func (h *SpectrumHelpText) Child(child app.UI) *SpectrumHelpText {
	h.child = child
	return h
}

// Negative is a convenience method to set the variant to negative
func (h *SpectrumHelpText) Negative() *SpectrumHelpText {
	h.variant = HelpTextVariantNegative
	return h
}

// Render renders the help text component
func (h *SpectrumHelpText) Render() app.UI {
	helpText := app.Elem("sp-help-text").
		Attr("variant", string(h.variant)).
		Attr("size", string(h.size)).
		Attr("icon", h.icon).
		Attr("disabled", h.disabled)

	if h.slot != "" {
		helpText = helpText.Attr("slot", h.slot)
	}

	// Handle content in the right order of precedence
	if h.innerHTML != "" {
		helpText = helpText.Body(app.Raw(h.innerHTML))
	} else if h.child != nil {
		helpText = helpText.Body(h.child)
	} else if h.content != "" {
		helpText = helpText.Body(app.Text(h.content))
	}

	return helpText
}
