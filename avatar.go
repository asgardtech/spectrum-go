package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AvatarSize represents the visual size of an avatar
type AvatarSize string

// Avatar sizes
const (
	AvatarSize50  AvatarSize = "50"
	AvatarSize75  AvatarSize = "75"
	AvatarSize100 AvatarSize = "100"
	AvatarSize200 AvatarSize = "200"
	AvatarSize300 AvatarSize = "300"
	AvatarSize400 AvatarSize = "400"
	AvatarSize500 AvatarSize = "500"
	AvatarSize600 AvatarSize = "600"
	AvatarSize700 AvatarSize = "700"
)

// AvatarTarget represents where a link avatar opens its URL
type AvatarTarget string

// Avatar targets
const (
	AvatarTargetBlank  AvatarTarget = "_blank"
	AvatarTargetParent AvatarTarget = "_parent"
	AvatarTargetSelf   AvatarTarget = "_self"
	AvatarTargetTop    AvatarTarget = "_top"
)

// AvatarReferrerPolicy represents how much referrer information is sent
type AvatarReferrerPolicy string

// Referrer policies
const (
	AvatarReferrerPolicyNoReferrer                  AvatarReferrerPolicy = "no-referrer"
	AvatarReferrerPolicyNoReferrerWhenDowngrade     AvatarReferrerPolicy = "no-referrer-when-downgrade"
	AvatarReferrerPolicyOrigin                      AvatarReferrerPolicy = "origin"
	AvatarReferrerPolicyOriginWhenCrossOrigin       AvatarReferrerPolicy = "origin-when-cross-origin"
	AvatarReferrerPolicySameOrigin                  AvatarReferrerPolicy = "same-origin"
	AvatarReferrerPolicyStrictOrigin                AvatarReferrerPolicy = "strict-origin"
	AvatarReferrerPolicyStrictOriginWhenCrossOrigin AvatarReferrerPolicy = "strict-origin-when-cross-origin"
	AvatarReferrerPolicyUnsafeUrl                   AvatarReferrerPolicy = "unsafe-url"
)

// spectrumAvatar represents an sp-avatar component
type spectrumAvatar struct {
	app.Compo

	// Properties
	PropDisabled       bool
	PropDownload       string
	PropHref           string
	PropLabel          string
	PropReferrerpolicy AvatarReferrerPolicy
	PropRel            string
	PropSize           AvatarSize
	PropSrc            string
	PropTabIndex       int
	PropTarget         AvatarTarget
}

// Avatar creates a new avatar component
func Avatar() *spectrumAvatar {
	return &spectrumAvatar{
		PropSize: AvatarSize100, // Default size is 100
	}
}

// Disabled sets the disabled state of the avatar
func (a *spectrumAvatar) Disabled(disabled bool) *spectrumAvatar {
	a.PropDisabled = disabled
	return a
}

// Download sets the download attribute
func (a *spectrumAvatar) Download(download string) *spectrumAvatar {
	a.PropDownload = download
	return a
}

// Href sets the URL that the avatar links to
func (a *spectrumAvatar) Href(href string) *spectrumAvatar {
	a.PropHref = href
	return a
}

// Label sets the accessible label for the avatar
func (a *spectrumAvatar) Label(label string) *spectrumAvatar {
	a.PropLabel = label
	return a
}

// Referrerpolicy sets the referrer policy for the link
func (a *spectrumAvatar) Referrerpolicy(referrerpolicy AvatarReferrerPolicy) *spectrumAvatar {
	a.PropReferrerpolicy = referrerpolicy
	return a
}

// Rel sets the relationship of the linked URL
func (a *spectrumAvatar) Rel(rel string) *spectrumAvatar {
	a.PropRel = rel
	return a
}

// Size sets the visual size of the avatar
func (a *spectrumAvatar) Size(size AvatarSize) *spectrumAvatar {
	a.PropSize = size
	return a
}

// Src sets the image source URL
func (a *spectrumAvatar) Src(src string) *spectrumAvatar {
	a.PropSrc = src
	return a
}

// TabIndex sets the tab index
func (a *spectrumAvatar) TabIndex(tabIndex int) *spectrumAvatar {
	a.PropTabIndex = tabIndex
	return a
}

// Target sets where to display the linked URL
func (a *spectrumAvatar) Target(target AvatarTarget) *spectrumAvatar {
	a.PropTarget = target
	return a
}

// Render renders the avatar component
func (a *spectrumAvatar) Render() app.UI {
	avatar := app.Elem("sp-avatar").
		Attr("disabled", a.PropDisabled).
		Attr("src", a.PropSrc)

	if a.PropDownload != "" {
		avatar = avatar.Attr("download", a.PropDownload)
	}
	if a.PropHref != "" {
		avatar = avatar.Attr("href", a.PropHref)
	}
	if a.PropLabel != "" {
		avatar = avatar.Attr("label", a.PropLabel)
	}
	if a.PropReferrerpolicy != "" {
		avatar = avatar.Attr("referrerpolicy", string(a.PropReferrerpolicy))
	}
	if a.PropRel != "" {
		avatar = avatar.Attr("rel", a.PropRel)
	}
	if a.PropSize != "" {
		avatar = avatar.Attr("size", string(a.PropSize))
	}
	if a.PropTabIndex != 0 {
		avatar = avatar.Attr("tabindex", a.PropTabIndex)
	}
	if a.PropTarget != "" {
		avatar = avatar.Attr("target", string(a.PropTarget))
	}

	return avatar
}
