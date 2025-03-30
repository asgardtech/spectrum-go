package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AvatarReferrerpolicy represents the How much of the referrer to send when following the link.
type AvatarReferrerpolicy string

// AvatarReferrerpolicy values
const (
    AvatarReferrerpolicyNoReferrer AvatarReferrerpolicy = "no-referrer"
    AvatarReferrerpolicyNoReferrerWhenDowngrade AvatarReferrerpolicy = "no-referrer-when-downgrade"
    AvatarReferrerpolicyOrigin AvatarReferrerpolicy = "origin"
    AvatarReferrerpolicyOriginWhenCrossOrigin AvatarReferrerpolicy = "origin-when-cross-origin"
    AvatarReferrerpolicySameOrigin AvatarReferrerpolicy = "same-origin"
    AvatarReferrerpolicyStrictOrigin AvatarReferrerpolicy = "strict-origin"
    AvatarReferrerpolicyStrictOriginWhenCrossOrigin AvatarReferrerpolicy = "strict-origin-when-cross-origin"
    AvatarReferrerpolicyUnsafeUrl AvatarReferrerpolicy = "unsafe-url"
)
// AvatarSize represents the The size of the avatar
type AvatarSize string

// AvatarSize values
const (
    AvatarSize50 AvatarSize = "50"
    AvatarSize75 AvatarSize = "75"
    AvatarSize100 AvatarSize = "100"
    AvatarSize200 AvatarSize = "200"
    AvatarSize300 AvatarSize = "300"
    AvatarSize400 AvatarSize = "400"
    AvatarSize500 AvatarSize = "500"
    AvatarSize600 AvatarSize = "600"
    AvatarSize700 AvatarSize = "700"
)
// AvatarTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type AvatarTarget string

// AvatarTarget values
const (
    AvatarTarget_blank AvatarTarget = "_blank"
    AvatarTarget_parent AvatarTarget = "_parent"
    AvatarTarget_self AvatarTarget = "_self"
    AvatarTarget_top AvatarTarget = "_top"
)

// spectrumAvatar represents an sp-avatar component
type spectrumAvatar struct {
    app.Compo
    *styler[*spectrumAvatar]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download.
    PropDownload string
    // The URL that the hyperlink points to.
    PropHref string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // How much of the referrer to send when following the link.
    PropReferrerpolicy AvatarReferrerpolicy
    // The relationship of the linked URL as space-separated link types.
    PropRel string
    // The size of the avatar
    PropSize AvatarSize
    // The URL of the image to display for the avatar
    PropSrc string
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
    PropTarget AvatarTarget




    // Event handlers
    PropOnClick app.EventHandler
}

// Avatar creates a new sp-avatar component
func Avatar() *spectrumAvatar {
    element := &spectrumAvatar{
        PropDisabled: false,
        PropSize: AvatarSize100,
        PropSrc: "",
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumAvatar) Disabled(disabled bool) *spectrumAvatar {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumAvatar) SetDisabled() *spectrumAvatar {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumAvatar) Download(download string) *spectrumAvatar {
    c.PropDownload = download
    return c
}

// The URL that the hyperlink points to.
func (c *spectrumAvatar) Href(href string) *spectrumAvatar {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumAvatar) Label(label string) *spectrumAvatar {
    c.PropLabel = label
    return c
}

// How much of the referrer to send when following the link.
func (c *spectrumAvatar) Referrerpolicy(referrerpolicy AvatarReferrerpolicy) *spectrumAvatar {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumAvatar) ReferrerpolicyNoReferrer() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyNoReferrer)
}
func (c *spectrumAvatar) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumAvatar) ReferrerpolicyOrigin() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyOrigin)
}
func (c *spectrumAvatar) ReferrerpolicyOriginWhenCrossOrigin() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumAvatar) ReferrerpolicySameOrigin() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicySameOrigin)
}
func (c *spectrumAvatar) ReferrerpolicyStrictOrigin() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyStrictOrigin)
}
func (c *spectrumAvatar) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumAvatar) ReferrerpolicyUnsafeUrl() *spectrumAvatar {
    return c.Referrerpolicy(AvatarReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types.
func (c *spectrumAvatar) Rel(rel string) *spectrumAvatar {
    c.PropRel = rel
    return c
}

// The size of the avatar
func (c *spectrumAvatar) Size(size AvatarSize) *spectrumAvatar {
    c.PropSize = size
    return c
}

func (c *spectrumAvatar) Size50() *spectrumAvatar {
    return c.Size(AvatarSize50)
}
func (c *spectrumAvatar) Size75() *spectrumAvatar {
    return c.Size(AvatarSize75)
}
func (c *spectrumAvatar) Size100() *spectrumAvatar {
    return c.Size(AvatarSize100)
}
func (c *spectrumAvatar) Size200() *spectrumAvatar {
    return c.Size(AvatarSize200)
}
func (c *spectrumAvatar) Size300() *spectrumAvatar {
    return c.Size(AvatarSize300)
}
func (c *spectrumAvatar) Size400() *spectrumAvatar {
    return c.Size(AvatarSize400)
}
func (c *spectrumAvatar) Size500() *spectrumAvatar {
    return c.Size(AvatarSize500)
}
func (c *spectrumAvatar) Size600() *spectrumAvatar {
    return c.Size(AvatarSize600)
}
func (c *spectrumAvatar) Size700() *spectrumAvatar {
    return c.Size(AvatarSize700)
}
// The URL of the image to display for the avatar
func (c *spectrumAvatar) Src(src string) *spectrumAvatar {
    c.PropSrc = src
    return c
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumAvatar) Tabindex(tabindex float64) *spectrumAvatar {
    c.PropTabindex = tabindex
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumAvatar) Target(target AvatarTarget) *spectrumAvatar {
    c.PropTarget = target
    return c
}

func (c *spectrumAvatar) Target_blank() *spectrumAvatar {
    return c.Target(AvatarTarget_blank)
}
func (c *spectrumAvatar) Target_parent() *spectrumAvatar {
    return c.Target(AvatarTarget_parent)
}
func (c *spectrumAvatar) Target_self() *spectrumAvatar {
    return c.Target(AvatarTarget_self)
}
func (c *spectrumAvatar) Target_top() *spectrumAvatar {
    return c.Target(AvatarTarget_top)
}




// Fired when the avatar is clicked when it has an href attribute
func (c *spectrumAvatar) OnClick(handler app.EventHandler) *spectrumAvatar {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-avatar component
func (c *spectrumAvatar) Render() app.UI {
    element := app.Elem("sp-avatar")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropReferrerpolicy != "" {
        element = element.Attr("referrerpolicy", string(c.PropReferrerpolicy))
    }
    if c.PropRel != "" {
        element = element.Attr("rel", c.PropRel)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropSrc != "" {
        element = element.Attr("src", c.PropSrc)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", string(c.PropTarget))
    }

    // Add event handlers
    if c.PropOnClick != nil {
        element = element.On("click", c.PropOnClick)
    }


    element = element.Styles(c.styler.styles)

    return element
} 