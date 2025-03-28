package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ButtonVariant represents the visual variant of a button
type ButtonVariant string

// Button variants
const (
	ButtonVariantPrimary   ButtonVariant = "primary"
	ButtonVariantSecondary ButtonVariant = "secondary"
	ButtonVariantCTA       ButtonVariant = "cta"
	ButtonVariantNegative  ButtonVariant = "negative"
)

// ButtonStaticColor represents the static color for a button
type ButtonStaticColor string

// Button static colors
const (
	ButtonStaticColorBlack ButtonStaticColor = "black"
	ButtonStaticColorWhite ButtonStaticColor = "white"
)

// ButtonTarget represents where a link button opens its URL
type ButtonTarget string

// Button targets
const (
	ButtonTargetBlank  ButtonTarget = "_blank"
	ButtonTargetParent ButtonTarget = "_parent"
	ButtonTargetSelf   ButtonTarget = "_self"
	ButtonTargetTop    ButtonTarget = "_top"
)

// ButtonType represents the HTML type attribute for a button
type ButtonType string

// Button types
const (
	ButtonTypeButton ButtonType = "button"
	ButtonTypeSubmit ButtonType = "submit"
	ButtonTypeReset  ButtonType = "reset"
)

// ReferrerPolicy represents how much referrer information is sent
type ReferrerPolicy string

// Referrer policies
const (
	ReferrerPolicyNoReferrer                  ReferrerPolicy = "no-referrer"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicy = "no-referrer-when-downgrade"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "origin-when-cross-origin"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "same-origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strict-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"
	ReferrerPolicyUnsafeUrl                   ReferrerPolicy = "unsafe-url"
)

// SpectrumButton represents an sp-button component
// This struct name is different from the constructor to avoid naming conflicts
type SpectrumButton struct {
	app.Compo

	// Properties
	text           string
	variant        ButtonVariant
	disabled       bool
	quiet          bool
	href           string
	target         ButtonTarget
	active         bool
	download       string
	label          string
	noWrap         bool
	pending        bool
	pendingLabel   string
	referrerPolicy ReferrerPolicy
	rel            string
	staticColor    ButtonStaticColor
	tabIndex       int
	treatment      string
	buttonType     ButtonType
	onClick        app.EventHandler

	// Slots
	icon *SpectrumIcon
}

// Button creates a new button component
// Function name matches the Spectrum component name for intuitive usage
func Button() *SpectrumButton {
	return &SpectrumButton{
		buttonType: ButtonTypeButton, // Default is "button"
		treatment:  "fill",           // Default treatment
	}
}

// Text sets the button text content
// This text is placed in the default slot of the component
func (b *SpectrumButton) Text(text string) *SpectrumButton {
	b.text = text
	return b
}

// Variant sets the button variant
func (b *SpectrumButton) Variant(variant ButtonVariant) *SpectrumButton {
	b.variant = variant
	return b
}

// Disabled sets the button disabled state
func (b *SpectrumButton) Disabled(disabled bool) *SpectrumButton {
	b.disabled = disabled
	return b
}

// Quiet sets the button quiet mode
// Makes the button appear less prominent
func (b *SpectrumButton) Quiet(quiet bool) *SpectrumButton {
	b.quiet = quiet
	return b
}

// Href sets the button href for link buttons
// When set, the button acts as a link
func (b *SpectrumButton) Href(href string) *SpectrumButton {
	b.href = href
	return b
}

// Target sets the target for link buttons
func (b *SpectrumButton) Target(target ButtonTarget) *SpectrumButton {
	b.target = target
	return b
}

// Active sets the active state
func (b *SpectrumButton) Active(active bool) *SpectrumButton {
	b.active = active
	return b
}

// Download sets the download attribute
// Causes the browser to treat the linked URL as a download
func (b *SpectrumButton) Download(download string) *SpectrumButton {
	b.download = download
	return b
}

// Label sets the accessible label
// Applied to aria-label but not visually rendered
func (b *SpectrumButton) Label(label string) *SpectrumButton {
	b.label = label
	return b
}

// NoWrap disables text wrapping
// Not part of the design spec, use carefully
func (b *SpectrumButton) NoWrap(noWrap bool) *SpectrumButton {
	b.noWrap = noWrap
	return b
}

// Pending sets the pending state
// Visually indicates the button is waiting for an action to complete
func (b *SpectrumButton) Pending(pending bool) *SpectrumButton {
	b.pending = pending
	return b
}

// PendingLabel sets the pending label text
// Default is 'Pending'
func (b *SpectrumButton) PendingLabel(pendingLabel string) *SpectrumButton {
	b.pendingLabel = pendingLabel
	return b
}

// ReferrerPolicy sets the referrer policy
func (b *SpectrumButton) ReferrerPolicy(policy ReferrerPolicy) *SpectrumButton {
	b.referrerPolicy = policy
	return b
}

// Rel sets the rel attribute
// Describes relationship of linked URL
func (b *SpectrumButton) Rel(rel string) *SpectrumButton {
	b.rel = rel
	return b
}

// StaticColor sets the static color
func (b *SpectrumButton) StaticColor(color ButtonStaticColor) *SpectrumButton {
	b.staticColor = color
	return b
}

// TabIndex sets the tab index
// Controls keyboard navigation order
func (b *SpectrumButton) TabIndex(index int) *SpectrumButton {
	b.tabIndex = index
	return b
}

// Treatment sets the treatment
// Default is "fill"
func (b *SpectrumButton) Treatment(treatment string) *SpectrumButton {
	b.treatment = treatment
	return b
}

// Type sets the button type
func (b *SpectrumButton) Type(buttonType ButtonType) *SpectrumButton {
	b.buttonType = buttonType
	return b
}

// Icon sets the icon in the icon slot
// Takes a SpectrumIcon component which will be placed in the icon slot
func (b *SpectrumButton) Icon(icon *SpectrumIcon) *SpectrumButton {
	b.icon = icon
	return b
}

// OnClick sets the click event handler
func (b *SpectrumButton) OnClick(handler app.EventHandler) *SpectrumButton {
	b.onClick = handler
	return b
}

// Render renders the button component
// Creates the sp-button element with all configured attributes and slots
func (b *SpectrumButton) Render() app.UI {
	// Create the base button element using app.Elem with the custom tag name
	button := app.Elem("sp-button").
		Attr("variant", string(b.variant)).
		Attr("disabled", b.disabled).
		Attr("quiet", b.quiet).
		Attr("href", b.href).
		Attr("target", string(b.target)).
		Attr("active", b.active).
		Attr("download", b.download).
		Attr("label", b.label).
		Attr("no-wrap", b.noWrap).
		Attr("pending", b.pending).
		Attr("pending-label", b.pendingLabel).
		Attr("referrerpolicy", string(b.referrerPolicy)).
		Attr("rel", b.rel).
		Attr("static-color", string(b.staticColor)).
		Attr("tabindex", b.tabIndex).
		Attr("treatment", b.treatment).
		Attr("type", string(b.buttonType))

	// Add event handler if set
	if b.onClick != nil {
		button = button.OnClick(b.onClick)
	}

	// Handle body content - text and icon
	if b.icon != nil {
		// If there's an icon, we need to explicitly set both the icon and text in the body
		b.icon.Slot("icon")

		button.Body(
			// Add the icon with its slot
			b.icon,

			// Add the text content
			app.Text(b.text),
		)
	} else {
		// If there's no icon, just set the text directly
		button.Text(b.text)
	}

	return button
}
