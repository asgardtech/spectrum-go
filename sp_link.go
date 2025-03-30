package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// LinkReferrerpolicy represents the How much of the referrer to send when following the link.
type LinkReferrerpolicy string

// LinkReferrerpolicy values
const (
    LinkReferrerpolicyNoReferrer LinkReferrerpolicy = "no-referrer"
    LinkReferrerpolicyNoReferrerWhenDowngrade LinkReferrerpolicy = "no-referrer-when-downgrade"
    LinkReferrerpolicyOrigin LinkReferrerpolicy = "origin"
    LinkReferrerpolicyOriginWhenCrossOrigin LinkReferrerpolicy = "origin-when-cross-origin"
    LinkReferrerpolicySameOrigin LinkReferrerpolicy = "same-origin"
    LinkReferrerpolicyStrictOrigin LinkReferrerpolicy = "strict-origin"
    LinkReferrerpolicyStrictOriginWhenCrossOrigin LinkReferrerpolicy = "strict-origin-when-cross-origin"
    LinkReferrerpolicyUnsafeUrl LinkReferrerpolicy = "unsafe-url"
)
// LinkStaticColor represents the A static color variant for the link, useful when placing on colored backgrounds
type LinkStaticColor string

// LinkStaticColor values
const (
    LinkStaticColorBlack LinkStaticColor = "black"
    LinkStaticColorWhite LinkStaticColor = "white"
)
// LinkTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type LinkTarget string

// LinkTarget values
const (
    LinkTarget_blank LinkTarget = "_blank"
    LinkTarget_parent LinkTarget = "_parent"
    LinkTarget_self LinkTarget = "_self"
    LinkTarget_top LinkTarget = "_top"
)
// LinkVariant represents the The visual variant of the link
type LinkVariant string

// LinkVariant values
const (
    LinkVariantSecondary LinkVariant = "secondary"
)

// spectrumLink represents an sp-link component
type spectrumLink struct {
    app.Compo
    *styler[*spectrumLink]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download.
    PropDownload string
    // The URL that the hyperlink points to.
    PropHref string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // Whether to display the link without an underline
    PropQuiet bool
    // How much of the referrer to send when following the link.
    PropReferrerpolicy LinkReferrerpolicy
    // The relationship of the linked URL as space-separated link types.
    PropRel string
    // A static color variant for the link, useful when placing on colored backgrounds
    PropStaticColor LinkStaticColor
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
    PropTarget LinkTarget
    // The visual variant of the link
    PropVariant LinkVariant

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnClick app.EventHandler
}

// Link creates a new sp-link component
func Link() *spectrumLink {
    element := &spectrumLink{
        PropDisabled: false,
        PropQuiet: false,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumLink) Disabled(disabled bool) *spectrumLink {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumLink) SetDisabled() *spectrumLink {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumLink) Download(download string) *spectrumLink {
    c.PropDownload = download
    return c
}

// The URL that the hyperlink points to.
func (c *spectrumLink) Href(href string) *spectrumLink {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumLink) Label(label string) *spectrumLink {
    c.PropLabel = label
    return c
}

// Whether to display the link without an underline
func (c *spectrumLink) Quiet(quiet bool) *spectrumLink {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumLink) SetQuiet() *spectrumLink {
    return c.Quiet(true)
}

// How much of the referrer to send when following the link.
func (c *spectrumLink) Referrerpolicy(referrerpolicy LinkReferrerpolicy) *spectrumLink {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumLink) ReferrerpolicyNoReferrer() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyNoReferrer)
}
func (c *spectrumLink) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumLink) ReferrerpolicyOrigin() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyOrigin)
}
func (c *spectrumLink) ReferrerpolicyOriginWhenCrossOrigin() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumLink) ReferrerpolicySameOrigin() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicySameOrigin)
}
func (c *spectrumLink) ReferrerpolicyStrictOrigin() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyStrictOrigin)
}
func (c *spectrumLink) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumLink) ReferrerpolicyUnsafeUrl() *spectrumLink {
    return c.Referrerpolicy(LinkReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types.
func (c *spectrumLink) Rel(rel string) *spectrumLink {
    c.PropRel = rel
    return c
}

// A static color variant for the link, useful when placing on colored backgrounds
func (c *spectrumLink) StaticColor(staticColor LinkStaticColor) *spectrumLink {
    c.PropStaticColor = staticColor
    return c
}

func (c *spectrumLink) StaticColorBlack() *spectrumLink {
    return c.StaticColor(LinkStaticColorBlack)
}
func (c *spectrumLink) StaticColorWhite() *spectrumLink {
    return c.StaticColor(LinkStaticColorWhite)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumLink) Tabindex(tabindex float64) *spectrumLink {
    c.PropTabindex = tabindex
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumLink) Target(target LinkTarget) *spectrumLink {
    c.PropTarget = target
    return c
}

func (c *spectrumLink) Target_blank() *spectrumLink {
    return c.Target(LinkTarget_blank)
}
func (c *spectrumLink) Target_parent() *spectrumLink {
    return c.Target(LinkTarget_parent)
}
func (c *spectrumLink) Target_self() *spectrumLink {
    return c.Target(LinkTarget_self)
}
func (c *spectrumLink) Target_top() *spectrumLink {
    return c.Target(LinkTarget_top)
}
// The visual variant of the link
func (c *spectrumLink) Variant(variant LinkVariant) *spectrumLink {
    c.PropVariant = variant
    return c
}

func (c *spectrumLink) VariantSecondary() *spectrumLink {
    return c.Variant(LinkVariantSecondary)
}

// Text sets the text content for the default slot
func (c *spectrumLink) Text(text string) *spectrumLink {
    c.PropText = text
    return c
}




// Fired when the link is clicked
func (c *spectrumLink) OnClick(handler app.EventHandler) *spectrumLink {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-link component
func (c *spectrumLink) Render() app.UI {
    element := app.Elem("sp-link")

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
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropReferrerpolicy != "" {
        element = element.Attr("referrerpolicy", string(c.PropReferrerpolicy))
    }
    if c.PropRel != "" {
        element = element.Attr("rel", c.PropRel)
    }
    if c.PropStaticColor != "" {
        element = element.Attr("static-color", string(c.PropStaticColor))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", string(c.PropTarget))
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnClick != nil {
        element = element.On("click", c.PropOnClick)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 