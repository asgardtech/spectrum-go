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

// SpectrumAvatar represents an sp-avatar component
type SpectrumAvatar struct {
	app.Compo

	// Properties
	disabled       bool
	download       string
	href           string
	label          string
	referrerpolicy AvatarReferrerPolicy
	rel            string
	size           AvatarSize
	src            string
	tabIndex       int
	target         AvatarTarget
}

// Avatar creates a new avatar component
func Avatar() *SpectrumAvatar {
	return &SpectrumAvatar{
		size: AvatarSize100, // Default size is 100
	}
}

// Disabled sets the disabled state of the avatar
func (a *SpectrumAvatar) Disabled(disabled bool) *SpectrumAvatar {
	a.disabled = disabled
	return a
}

// Download sets the download attribute
func (a *SpectrumAvatar) Download(download string) *SpectrumAvatar {
	a.download = download
	return a
}

// Href sets the URL that the avatar links to
func (a *SpectrumAvatar) Href(href string) *SpectrumAvatar {
	a.href = href
	return a
}

// Label sets the accessible label for the avatar
func (a *SpectrumAvatar) Label(label string) *SpectrumAvatar {
	a.label = label
	return a
}

// Referrerpolicy sets the referrer policy for the link
func (a *SpectrumAvatar) Referrerpolicy(referrerpolicy AvatarReferrerPolicy) *SpectrumAvatar {
	a.referrerpolicy = referrerpolicy
	return a
}

// Rel sets the relationship of the linked URL
func (a *SpectrumAvatar) Rel(rel string) *SpectrumAvatar {
	a.rel = rel
	return a
}

// Size sets the visual size of the avatar
func (a *SpectrumAvatar) Size(size AvatarSize) *SpectrumAvatar {
	a.size = size
	return a
}

// Src sets the image source URL
func (a *SpectrumAvatar) Src(src string) *SpectrumAvatar {
	a.src = src
	return a
}

// TabIndex sets the tab index
func (a *SpectrumAvatar) TabIndex(tabIndex int) *SpectrumAvatar {
	a.tabIndex = tabIndex
	return a
}

// Target sets where to display the linked URL
func (a *SpectrumAvatar) Target(target AvatarTarget) *SpectrumAvatar {
	a.target = target
	return a
}

// Render renders the avatar component
func (a *SpectrumAvatar) Render() app.UI {
	avatar := app.Elem("sp-avatar").
		Attr("disabled", a.disabled).
		Attr("src", a.src)

	if a.download != "" {
		avatar = avatar.Attr("download", a.download)
	}
	if a.href != "" {
		avatar = avatar.Attr("href", a.href)
	}
	if a.label != "" {
		avatar = avatar.Attr("label", a.label)
	}
	if a.referrerpolicy != "" {
		avatar = avatar.Attr("referrerpolicy", string(a.referrerpolicy))
	}
	if a.rel != "" {
		avatar = avatar.Attr("rel", a.rel)
	}
	if a.size != "" {
		avatar = avatar.Attr("size", string(a.size))
	}
	if a.tabIndex != 0 {
		avatar = avatar.Attr("tabindex", a.tabIndex)
	}
	if a.target != "" {
		avatar = avatar.Attr("target", string(a.target))
	}

	return avatar
}
