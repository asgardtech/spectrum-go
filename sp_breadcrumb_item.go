package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// BreadcrumbItemReferrerpolicy represents the How much of the referrer to send when following the link
type BreadcrumbItemReferrerpolicy string

// BreadcrumbItemReferrerpolicy values
const (
    BreadcrumbItemReferrerpolicyNoReferrer BreadcrumbItemReferrerpolicy = "no-referrer"
    BreadcrumbItemReferrerpolicyNoReferrerWhenDowngrade BreadcrumbItemReferrerpolicy = "no-referrer-when-downgrade"
    BreadcrumbItemReferrerpolicyOrigin BreadcrumbItemReferrerpolicy = "origin"
    BreadcrumbItemReferrerpolicyOriginWhenCrossOrigin BreadcrumbItemReferrerpolicy = "origin-when-cross-origin"
    BreadcrumbItemReferrerpolicySameOrigin BreadcrumbItemReferrerpolicy = "same-origin"
    BreadcrumbItemReferrerpolicyStrictOrigin BreadcrumbItemReferrerpolicy = "strict-origin"
    BreadcrumbItemReferrerpolicyStrictOriginWhenCrossOrigin BreadcrumbItemReferrerpolicy = "strict-origin-when-cross-origin"
    BreadcrumbItemReferrerpolicyUnsafeUrl BreadcrumbItemReferrerpolicy = "unsafe-url"
)
// BreadcrumbItemTarget represents the Where to display the linked URL, as the name for a browsing context
type BreadcrumbItemTarget string

// BreadcrumbItemTarget values
const (
    BreadcrumbItemTarget_blank BreadcrumbItemTarget = "_blank"
    BreadcrumbItemTarget_parent BreadcrumbItemTarget = "_parent"
    BreadcrumbItemTarget_self BreadcrumbItemTarget = "_self"
    BreadcrumbItemTarget_top BreadcrumbItemTarget = "_top"
)

// spectrumBreadcrumbItem represents an sp-breadcrumb-item component
type spectrumBreadcrumbItem struct {
    app.Compo
    *styler[*spectrumBreadcrumbItem]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download
    PropDownload string
    // The URL that the hyperlink points to. When present, the item functions as a link.
    PropHref string
    // An accessible label that describes the component. Applied to aria-label but not visually rendered.
    PropLabel string
    // How much of the referrer to send when following the link
    PropReferrerpolicy BreadcrumbItemReferrerpolicy
    // The relationship of the linked URL as space-separated link types
    PropRel string
    // The tab index to apply to this control
    PropTabindex float64
    // Where to display the linked URL, as the name for a browsing context
    PropTarget BreadcrumbItemTarget
    // The value associated with this breadcrumb item, used for selection events
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots


}

// BreadcrumbItem creates a new sp-breadcrumb-item component
func BreadcrumbItem() *spectrumBreadcrumbItem {
    element := &spectrumBreadcrumbItem{
        PropDisabled: false,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumBreadcrumbItem) Disabled(disabled bool) *spectrumBreadcrumbItem {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumBreadcrumbItem) SetDisabled() *spectrumBreadcrumbItem {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download
func (c *spectrumBreadcrumbItem) Download(download string) *spectrumBreadcrumbItem {
    c.PropDownload = download
    return c
}

// The URL that the hyperlink points to. When present, the item functions as a link.
func (c *spectrumBreadcrumbItem) Href(href string) *spectrumBreadcrumbItem {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. Applied to aria-label but not visually rendered.
func (c *spectrumBreadcrumbItem) Label(label string) *spectrumBreadcrumbItem {
    c.PropLabel = label
    return c
}

// How much of the referrer to send when following the link
func (c *spectrumBreadcrumbItem) Referrerpolicy(referrerpolicy BreadcrumbItemReferrerpolicy) *spectrumBreadcrumbItem {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumBreadcrumbItem) ReferrerpolicyNoReferrer() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyNoReferrer)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyOrigin() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyOrigin)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyOriginWhenCrossOrigin() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicySameOrigin() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicySameOrigin)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyStrictOrigin() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyStrictOrigin)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumBreadcrumbItem) ReferrerpolicyUnsafeUrl() *spectrumBreadcrumbItem {
    return c.Referrerpolicy(BreadcrumbItemReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types
func (c *spectrumBreadcrumbItem) Rel(rel string) *spectrumBreadcrumbItem {
    c.PropRel = rel
    return c
}

// The tab index to apply to this control
func (c *spectrumBreadcrumbItem) Tabindex(tabIndex float64) *spectrumBreadcrumbItem {
    c.PropTabindex = tabIndex
    return c
}

// Where to display the linked URL, as the name for a browsing context
func (c *spectrumBreadcrumbItem) Target(target BreadcrumbItemTarget) *spectrumBreadcrumbItem {
    c.PropTarget = target
    return c
}

func (c *spectrumBreadcrumbItem) Target_blank() *spectrumBreadcrumbItem {
    return c.Target(BreadcrumbItemTarget_blank)
}
func (c *spectrumBreadcrumbItem) Target_parent() *spectrumBreadcrumbItem {
    return c.Target(BreadcrumbItemTarget_parent)
}
func (c *spectrumBreadcrumbItem) Target_self() *spectrumBreadcrumbItem {
    return c.Target(BreadcrumbItemTarget_self)
}
func (c *spectrumBreadcrumbItem) Target_top() *spectrumBreadcrumbItem {
    return c.Target(BreadcrumbItemTarget_top)
}
// The value associated with this breadcrumb item, used for selection events
func (c *spectrumBreadcrumbItem) Value(value string) *spectrumBreadcrumbItem {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumBreadcrumbItem) Text(text string) *spectrumBreadcrumbItem {
    c.PropText = text
    return c
}





// Render renders the sp-breadcrumb-item component
func (c *spectrumBreadcrumbItem) Render() app.UI {
    element := app.Elem("sp-breadcrumb-item")

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
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", string(c.PropTarget))
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
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