package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ButtonVariant represents the The visual variant to apply to this button.
type ButtonVariant string

// ButtonVariant values
const (
    ButtonVariantAccent ButtonVariant = "accent"
    ButtonVariantPrimary ButtonVariant = "primary"
    ButtonVariantSecondary ButtonVariant = "secondary"
    ButtonVariantNegative ButtonVariant = "negative"
)
// ButtonTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type ButtonTarget string

// ButtonTarget values
const (
    ButtonTarget_blank ButtonTarget = "_blank"
    ButtonTarget_parent ButtonTarget = "_parent"
    ButtonTarget_self ButtonTarget = "_self"
    ButtonTarget_top ButtonTarget = "_top"
)
// ButtonReferrerpolicy represents the How much of the referrer to send when following the link.
type ButtonReferrerpolicy string

// ButtonReferrerpolicy values
const (
    ButtonReferrerpolicyNoReferrer ButtonReferrerpolicy = "no-referrer"
    ButtonReferrerpolicyNoReferrerWhenDowngrade ButtonReferrerpolicy = "no-referrer-when-downgrade"
    ButtonReferrerpolicyOrigin ButtonReferrerpolicy = "origin"
    ButtonReferrerpolicyOriginWhenCrossOrigin ButtonReferrerpolicy = "origin-when-cross-origin"
    ButtonReferrerpolicySameOrigin ButtonReferrerpolicy = "same-origin"
    ButtonReferrerpolicyStrictOrigin ButtonReferrerpolicy = "strict-origin"
    ButtonReferrerpolicyStrictOriginWhenCrossOrigin ButtonReferrerpolicy = "strict-origin-when-cross-origin"
    ButtonReferrerpolicyUnsafeUrl ButtonReferrerpolicy = "unsafe-url"
)
// ButtonStaticColor represents the The static color variant to use for this button.
type ButtonStaticColor string

// ButtonStaticColor values
const (
    ButtonStaticColorBlack ButtonStaticColor = "black"
    ButtonStaticColorWhite ButtonStaticColor = "white"
)
// ButtonTreatment represents the The visual treatment to apply to this button.
type ButtonTreatment string

// ButtonTreatment values
const (
    ButtonTreatmentFill ButtonTreatment = "fill"
    ButtonTreatmentOutline ButtonTreatment = "outline"
)
// ButtonType represents the The default behavior of the button.
type ButtonType string

// ButtonType values
const (
    ButtonTypeButton ButtonType = "button"
    ButtonTypeSubmit ButtonType = "submit"
    ButtonTypeReset ButtonType = "reset"
)

// spectrumButton represents an sp-button component
type spectrumButton struct {
    app.Compo
    *styler[*spectrumButton]

    // Properties
    // The visual variant to apply to this button.
    PropVariant ButtonVariant
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Style this button to be less obvious
    PropQuiet bool
    // The URL that the hyperlink points to.
    PropHref string
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
    PropTarget ButtonTarget
    // Whether the button appears in an active state
    PropActive bool
    // Causes the browser to treat the linked URL as a download.
    PropDownload string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // Disables text wrapping within the button component's label. Please note that this option is not a part of the design specification and should be used carefully, with consideration of this overflow behavior and the readability of the button's content.
    PropNoWrap bool
    // Whether the button is in a pending state, showing a progress indicator
    PropPending bool
    // Text shown while the button is in pending state
    PropPendingLabel string
    // How much of the referrer to send when following the link.
    PropReferrerpolicy ButtonReferrerpolicy
    // The relationship of the linked URL as space-separated link types.
    PropRel string
    // The static color variant to use for this button.
    PropStaticColor ButtonStaticColor
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // The visual treatment to apply to this button.
    PropTreatment ButtonTreatment
    // The default behavior of the button.
    PropType ButtonType

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnClick app.EventHandler
}

// Button creates a new sp-button component
func Button() *spectrumButton {
    element := &spectrumButton{
        PropVariant: ButtonVariantAccent,
        PropDisabled: false,
        PropQuiet: false,
        PropActive: false,
        PropNoWrap: false,
        PropPending: false,
        PropPendingLabel: "Pending",
        PropTabindex: 0,
        PropTreatment: ButtonTreatmentFill,
        PropType: ButtonTypeButton,
    }

    element.styler = newStyler(element)

    return element
}

// The visual variant to apply to this button.
func (c *spectrumButton) Variant(variant ButtonVariant) *spectrumButton {
    c.PropVariant = variant
    return c
}

func (c *spectrumButton) VariantAccent() *spectrumButton {
    return c.Variant(ButtonVariantAccent)
}
func (c *spectrumButton) VariantPrimary() *spectrumButton {
    return c.Variant(ButtonVariantPrimary)
}
func (c *spectrumButton) VariantSecondary() *spectrumButton {
    return c.Variant(ButtonVariantSecondary)
}
func (c *spectrumButton) VariantNegative() *spectrumButton {
    return c.Variant(ButtonVariantNegative)
}
// Disable this control. It will not receive focus or events
func (c *spectrumButton) Disabled(disabled bool) *spectrumButton {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumButton) SetDisabled() *spectrumButton {
    return c.Disabled(true)
}

// Style this button to be less obvious
func (c *spectrumButton) Quiet(quiet bool) *spectrumButton {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumButton) SetQuiet() *spectrumButton {
    return c.Quiet(true)
}

// The URL that the hyperlink points to.
func (c *spectrumButton) Href(href string) *spectrumButton {
    c.PropHref = href
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumButton) Target(target ButtonTarget) *spectrumButton {
    c.PropTarget = target
    return c
}

func (c *spectrumButton) Target_blank() *spectrumButton {
    return c.Target(ButtonTarget_blank)
}
func (c *spectrumButton) Target_parent() *spectrumButton {
    return c.Target(ButtonTarget_parent)
}
func (c *spectrumButton) Target_self() *spectrumButton {
    return c.Target(ButtonTarget_self)
}
func (c *spectrumButton) Target_top() *spectrumButton {
    return c.Target(ButtonTarget_top)
}
// Whether the button appears in an active state
func (c *spectrumButton) Active(active bool) *spectrumButton {
    c.PropActive = active
    return c
}

func (c *spectrumButton) SetActive() *spectrumButton {
    return c.Active(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumButton) Download(download string) *spectrumButton {
    c.PropDownload = download
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumButton) Label(label string) *spectrumButton {
    c.PropLabel = label
    return c
}

// Disables text wrapping within the button component's label. Please note that this option is not a part of the design specification and should be used carefully, with consideration of this overflow behavior and the readability of the button's content.
func (c *spectrumButton) NoWrap(noWrap bool) *spectrumButton {
    c.PropNoWrap = noWrap
    return c
}

func (c *spectrumButton) SetNoWrap() *spectrumButton {
    return c.NoWrap(true)
}

// Whether the button is in a pending state, showing a progress indicator
func (c *spectrumButton) Pending(pending bool) *spectrumButton {
    c.PropPending = pending
    return c
}

func (c *spectrumButton) SetPending() *spectrumButton {
    return c.Pending(true)
}

// Text shown while the button is in pending state
func (c *spectrumButton) PendingLabel(pendingLabel string) *spectrumButton {
    c.PropPendingLabel = pendingLabel
    return c
}

// How much of the referrer to send when following the link.
func (c *spectrumButton) Referrerpolicy(referrerpolicy ButtonReferrerpolicy) *spectrumButton {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumButton) ReferrerpolicyNoReferrer() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyNoReferrer)
}
func (c *spectrumButton) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumButton) ReferrerpolicyOrigin() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyOrigin)
}
func (c *spectrumButton) ReferrerpolicyOriginWhenCrossOrigin() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumButton) ReferrerpolicySameOrigin() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicySameOrigin)
}
func (c *spectrumButton) ReferrerpolicyStrictOrigin() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyStrictOrigin)
}
func (c *spectrumButton) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumButton) ReferrerpolicyUnsafeUrl() *spectrumButton {
    return c.Referrerpolicy(ButtonReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types.
func (c *spectrumButton) Rel(rel string) *spectrumButton {
    c.PropRel = rel
    return c
}

// The static color variant to use for this button.
func (c *spectrumButton) StaticColor(staticColor ButtonStaticColor) *spectrumButton {
    c.PropStaticColor = staticColor
    return c
}

func (c *spectrumButton) StaticColorBlack() *spectrumButton {
    return c.StaticColor(ButtonStaticColorBlack)
}
func (c *spectrumButton) StaticColorWhite() *spectrumButton {
    return c.StaticColor(ButtonStaticColorWhite)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumButton) Tabindex(tabindex float64) *spectrumButton {
    c.PropTabindex = tabindex
    return c
}

// The visual treatment to apply to this button.
func (c *spectrumButton) Treatment(treatment ButtonTreatment) *spectrumButton {
    c.PropTreatment = treatment
    return c
}

func (c *spectrumButton) TreatmentFill() *spectrumButton {
    return c.Treatment(ButtonTreatmentFill)
}
func (c *spectrumButton) TreatmentOutline() *spectrumButton {
    return c.Treatment(ButtonTreatmentOutline)
}
// The default behavior of the button.
func (c *spectrumButton) Type(typeValue ButtonType) *spectrumButton {
    c.PropType = typeValue
    return c
}

func (c *spectrumButton) TypeButton() *spectrumButton {
    return c.Type(ButtonTypeButton)
}
func (c *spectrumButton) TypeSubmit() *spectrumButton {
    return c.Type(ButtonTypeSubmit)
}
func (c *spectrumButton) TypeReset() *spectrumButton {
    return c.Type(ButtonTypeReset)
}

// Text sets the text content for the default slot
func (c *spectrumButton) Text(text string) *spectrumButton {
    c.PropText = text
    return c
}


// The icon to use for Button
func (c *spectrumButton) Icon(content app.UI) *spectrumButton {
    c.PropIconSlot = content

    return c
}



// Fired when the button is clicked
func (c *spectrumButton) OnClick(handler app.EventHandler) *spectrumButton {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-button component
func (c *spectrumButton) Render() app.UI {
    element := app.Elem("sp-button")

    // Set attributes
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", string(c.PropTarget))
    }
    if c.PropActive {
        element = element.Attr("active", true)
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropNoWrap {
        element = element.Attr("no-wrap", true)
    }
    if c.PropPending {
        element = element.Attr("pending", true)
    }
    if c.PropPendingLabel != "" {
        element = element.Attr("pending-label", c.PropPendingLabel)
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
    if c.PropTreatment != "" {
        element = element.Attr("treatment", string(c.PropTreatment))
    }
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
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

    // Add icon slot
    if c.PropIconSlot != nil {
        slotElem := c.PropIconSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("icon")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "icon").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 