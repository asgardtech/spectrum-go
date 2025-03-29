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

// spectrumButton represents an sp-button component
// This struct name is different from the constructor to avoid naming conflicts
type spectrumButton struct {
	app.Compo

	// Properties
	PropText           string
	PropVariant        ButtonVariant
	PropDisabled       bool
	PropQuiet          bool
	PropHref           string
	PropTarget         ButtonTarget
	PropActive         bool
	PropDownload       string
	PropLabel          string
	PropNoWrap         bool
	PropPending        bool
	PropPendingLabel   string
	PropReferrerPolicy ReferrerPolicy
	PropRel            string
	PropStaticColor    ButtonStaticColor
	PropTabIndex       int
	PropTreatment      string
	PropButtonType     ButtonType
	PropOnClick        app.EventHandler

	// Slots
	PropIcon *spectrumIcon
}

// Button creates a new button component
// Function name matches the Spectrum component name for intuitive usage
func Button() *spectrumButton {
	return &spectrumButton{
		PropButtonType: ButtonTypeButton, // Default is "button"
		PropTreatment:  "fill",           // Default treatment
	}
}

// Text sets the button text content
// This text is placed in the default slot of the component
func (b *spectrumButton) Text(text string) *spectrumButton {
	b.PropText = text
	return b
}

// Variant sets the button variant
func (b *spectrumButton) Variant(variant ButtonVariant) *spectrumButton {
	b.PropVariant = variant
	return b
}

// Disabled sets the button disabled state
func (b *spectrumButton) Disabled(disabled bool) *spectrumButton {
	b.PropDisabled = disabled
	return b
}

// Quiet sets the button quiet mode
// Makes the button appear less prominent
func (b *spectrumButton) Quiet(quiet bool) *spectrumButton {
	b.PropQuiet = quiet
	return b
}

// Href sets the button href for link buttons
// When set, the button acts as a link
func (b *spectrumButton) Href(href string) *spectrumButton {
	b.PropHref = href
	return b
}

// Target sets the target for link buttons
func (b *spectrumButton) Target(target ButtonTarget) *spectrumButton {
	b.PropTarget = target
	return b
}

// Active sets the active state
func (b *spectrumButton) Active(active bool) *spectrumButton {
	b.PropActive = active
	return b
}

// Download sets the download attribute
// Causes the browser to treat the linked URL as a download
func (b *spectrumButton) Download(download string) *spectrumButton {
	b.PropDownload = download
	return b
}

// Label sets the accessible label
// Applied to aria-label but not visually rendered
func (b *spectrumButton) Label(label string) *spectrumButton {
	b.PropLabel = label
	return b
}

// NoWrap disables text wrapping
// Not part of the design spec, use carefully
func (b *spectrumButton) NoWrap(noWrap bool) *spectrumButton {
	b.PropNoWrap = noWrap
	return b
}

// Pending sets the pending state
// Visually indicates the button is waiting for an action to complete
func (b *spectrumButton) Pending(pending bool) *spectrumButton {
	b.PropPending = pending
	return b
}

// PendingLabel sets the pending label text
// Default is 'Pending'
func (b *spectrumButton) PendingLabel(pendingLabel string) *spectrumButton {
	b.PropPendingLabel = pendingLabel
	return b
}

// ReferrerPolicy sets the referrer policy
func (b *spectrumButton) ReferrerPolicy(policy ReferrerPolicy) *spectrumButton {
	b.PropReferrerPolicy = policy
	return b
}

// Rel sets the rel attribute
// Describes relationship of linked URL
func (b *spectrumButton) Rel(rel string) *spectrumButton {
	b.PropRel = rel
	return b
}

// StaticColor sets the static color
func (b *spectrumButton) StaticColor(color ButtonStaticColor) *spectrumButton {
	b.PropStaticColor = color
	return b
}

// TabIndex sets the tab index
// Controls keyboard navigation order
func (b *spectrumButton) TabIndex(index int) *spectrumButton {
	b.PropTabIndex = index
	return b
}

// Treatment sets the treatment
// Default is "fill"
func (b *spectrumButton) Treatment(treatment string) *spectrumButton {
	b.PropTreatment = treatment
	return b
}

// Type sets the button type
func (b *spectrumButton) Type(buttonType ButtonType) *spectrumButton {
	b.PropButtonType = buttonType
	return b
}

// Icon sets the icon in the icon slot
// Takes a spectrumIcon component which will be placed in the icon slot
func (b *spectrumButton) Icon(icon *spectrumIcon) *spectrumButton {
	b.PropIcon = icon
	return b
}

// OnClick sets the click event handler
func (b *spectrumButton) OnClick(handler app.EventHandler) *spectrumButton {
	b.PropOnClick = handler
	return b
}

// Render renders the button component
// Creates the sp-button element with all configured attributes and slots
func (b *spectrumButton) Render() app.UI {
	// Create the base button element using app.Elem with the custom tag name
	button := app.Elem("sp-button").
		Attr("variant", string(b.PropVariant)).
		Attr("disabled", b.PropDisabled).
		Attr("quiet", b.PropQuiet).
		Attr("href", b.PropHref).
		Attr("target", string(b.PropTarget)).
		Attr("active", b.PropActive).
		Attr("download", b.PropDownload).
		Attr("label", b.PropLabel).
		Attr("no-wrap", b.PropNoWrap).
		Attr("pending", b.PropPending).
		Attr("pending-label", b.PropPendingLabel).
		Attr("referrerpolicy", string(b.PropReferrerPolicy)).
		Attr("rel", b.PropRel).
		Attr("static-color", string(b.PropStaticColor)).
		Attr("tabindex", b.PropTabIndex).
		Attr("treatment", b.PropTreatment).
		Attr("type", string(b.PropButtonType))

	// Add event handler if set
	if b.PropOnClick != nil {
		button = button.OnClick(b.PropOnClick)
	}

	// Handle body content - text and icon
	if b.PropIcon != nil {
		// If there's an icon, we need to explicitly set both the icon and text in the body
		b.PropIcon.Slot("icon")

		button.Body(
			// Add the icon with its slot
			b.PropIcon,

			// Add the text content
			app.Text(b.PropText),
		)
	} else {
		// If there's no icon, just set the text directly
		button.Text(b.PropText)
	}

	return button
}
