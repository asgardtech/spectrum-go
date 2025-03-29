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
type spectrumLink struct {
	app.Compo

	// Properties
	PropText           string
	PropVariant        LinkVariant
	PropHref           string
	PropTarget         LinkTarget
	PropDisabled       bool
	PropQuiet          bool
	PropDownload       string
	PropLabel          string
	PropReferrerPolicy ReferrerPolicy
	PropRel            string
	PropStaticColor    LinkStaticColor
	PropTabIndex       int
}

// Link creates a new link component
func Link() *spectrumLink {
	return &spectrumLink{}
}

// Text sets the link text content
func (l *spectrumLink) Text(text string) *spectrumLink {
	l.PropText = text
	return l
}

// Variant sets the link variant
func (l *spectrumLink) Variant(variant LinkVariant) *spectrumLink {
	l.PropVariant = variant
	return l
}

// Href sets the URL the link points to
func (l *spectrumLink) Href(href string) *spectrumLink {
	l.PropHref = href
	return l
}

// Target sets the target for the link
func (l *spectrumLink) Target(target LinkTarget) *spectrumLink {
	l.PropTarget = target
	return l
}

// Disabled sets the disabled state
func (l *spectrumLink) Disabled(disabled bool) *spectrumLink {
	l.PropDisabled = disabled
	return l
}

// Quiet sets whether the link uses quiet styles
func (l *spectrumLink) Quiet(quiet bool) *spectrumLink {
	l.PropQuiet = quiet
	return l
}

// Download sets the download attribute
func (l *spectrumLink) Download(download string) *spectrumLink {
	l.PropDownload = download
	return l
}

// Label sets the accessible label
func (l *spectrumLink) Label(label string) *spectrumLink {
	l.PropLabel = label
	return l
}

// ReferrerPolicy sets the referrer policy
func (l *spectrumLink) ReferrerPolicy(policy ReferrerPolicy) *spectrumLink {
	l.PropReferrerPolicy = policy
	return l
}

// Rel sets the rel attribute
func (l *spectrumLink) Rel(rel string) *spectrumLink {
	l.PropRel = rel
	return l
}

// StaticColor sets the static color
func (l *spectrumLink) StaticColor(color LinkStaticColor) *spectrumLink {
	l.PropStaticColor = color
	return l
}

// TabIndex sets the tab index
func (l *spectrumLink) TabIndex(index int) *spectrumLink {
	l.PropTabIndex = index
	return l
}

// Render renders the link component
func (l *spectrumLink) Render() app.UI {
	link := app.Elem("sp-link").
		Attr("variant", string(l.PropVariant)).
		Attr("href", l.PropHref).
		Attr("target", string(l.PropTarget)).
		Attr("disabled", l.PropDisabled).
		Attr("quiet", l.PropQuiet).
		Attr("download", l.PropDownload).
		Attr("label", l.PropLabel).
		Attr("referrerpolicy", string(l.PropReferrerPolicy)).
		Attr("rel", l.PropRel).
		Attr("static-color", string(l.PropStaticColor)).
		Attr("tabindex", l.PropTabIndex)

	// Add text content
	if l.PropText != "" {
		link = link.Text(l.PropText)
	}

	return link
}
