package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// LinkVariant represents the visual variant of a link
type LinkVariant string

// Link variants
const (
	LinkVariantSecondary LinkVariant = "secondary"
)

// LinkStaticColor represents the static color for a link
type LinkStaticColor string

// Link static colors
const (
	LinkStaticColorBlack LinkStaticColor = "black"
	LinkStaticColorWhite LinkStaticColor = "white"
)

// LinkTarget represents where a link opens its URL
type LinkTarget string

// Link targets
const (
	LinkTargetBlank  LinkTarget = "_blank"
	LinkTargetParent LinkTarget = "_parent"
	LinkTargetSelf   LinkTarget = "_self"
	LinkTargetTop    LinkTarget = "_top"
)

// Note: ReferrerPolicy is already defined in button.go, so we reuse it here

// SpectrumLink represents an sp-link component
type SpectrumLink struct {
	app.Compo

	// Properties
	text           string
	variant        LinkVariant
	href           string
	target         LinkTarget
	disabled       bool
	quiet          bool
	download       string
	label          string
	referrerPolicy ReferrerPolicy
	rel            string
	staticColor    LinkStaticColor
	tabIndex       int
}

// Link creates a new link component
func Link() *SpectrumLink {
	return &SpectrumLink{}
}

// Text sets the link text content
func (l *SpectrumLink) Text(text string) *SpectrumLink {
	l.text = text
	return l
}

// Variant sets the link variant
func (l *SpectrumLink) Variant(variant LinkVariant) *SpectrumLink {
	l.variant = variant
	return l
}

// Href sets the URL the link points to
func (l *SpectrumLink) Href(href string) *SpectrumLink {
	l.href = href
	return l
}

// Target sets the target for the link
func (l *SpectrumLink) Target(target LinkTarget) *SpectrumLink {
	l.target = target
	return l
}

// Disabled sets the disabled state
func (l *SpectrumLink) Disabled(disabled bool) *SpectrumLink {
	l.disabled = disabled
	return l
}

// Quiet sets whether the link uses quiet styles
func (l *SpectrumLink) Quiet(quiet bool) *SpectrumLink {
	l.quiet = quiet
	return l
}

// Download sets the download attribute
func (l *SpectrumLink) Download(download string) *SpectrumLink {
	l.download = download
	return l
}

// Label sets the accessible label
func (l *SpectrumLink) Label(label string) *SpectrumLink {
	l.label = label
	return l
}

// ReferrerPolicy sets the referrer policy
func (l *SpectrumLink) ReferrerPolicy(policy ReferrerPolicy) *SpectrumLink {
	l.referrerPolicy = policy
	return l
}

// Rel sets the rel attribute
func (l *SpectrumLink) Rel(rel string) *SpectrumLink {
	l.rel = rel
	return l
}

// StaticColor sets the static color
func (l *SpectrumLink) StaticColor(color LinkStaticColor) *SpectrumLink {
	l.staticColor = color
	return l
}

// TabIndex sets the tab index
func (l *SpectrumLink) TabIndex(index int) *SpectrumLink {
	l.tabIndex = index
	return l
}

// Render renders the link component
func (l *SpectrumLink) Render() app.UI {
	link := app.Elem("sp-link").
		Attr("variant", string(l.variant)).
		Attr("href", l.href).
		Attr("target", string(l.target)).
		Attr("disabled", l.disabled).
		Attr("quiet", l.quiet).
		Attr("download", l.download).
		Attr("label", l.label).
		Attr("referrerpolicy", string(l.referrerPolicy)).
		Attr("rel", l.rel).
		Attr("static-color", string(l.staticColor)).
		Attr("tabindex", l.tabIndex)

	// Add text content
	if l.text != "" {
		link = link.Text(l.text)
	}

	return link
}
