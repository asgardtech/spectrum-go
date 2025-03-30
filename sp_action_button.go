package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionButtonReferrerpolicy represents the How much of the referrer to send when following the link.
type ActionButtonReferrerpolicy string

// ActionButtonReferrerpolicy values
const (
    ActionButtonReferrerpolicyNoReferrer ActionButtonReferrerpolicy = "no-referrer"
    ActionButtonReferrerpolicyNoReferrerWhenDowngrade ActionButtonReferrerpolicy = "no-referrer-when-downgrade"
    ActionButtonReferrerpolicyOrigin ActionButtonReferrerpolicy = "origin"
    ActionButtonReferrerpolicyOriginWhenCrossOrigin ActionButtonReferrerpolicy = "origin-when-cross-origin"
    ActionButtonReferrerpolicySameOrigin ActionButtonReferrerpolicy = "same-origin"
    ActionButtonReferrerpolicyStrictOrigin ActionButtonReferrerpolicy = "strict-origin"
    ActionButtonReferrerpolicyStrictOriginWhenCrossOrigin ActionButtonReferrerpolicy = "strict-origin-when-cross-origin"
    ActionButtonReferrerpolicyUnsafeUrl ActionButtonReferrerpolicy = "unsafe-url"
)
// ActionButtonStaticColor represents the The static color variant to use for the action button.
type ActionButtonStaticColor string

// ActionButtonStaticColor values
const (
    ActionButtonStaticColorWhite ActionButtonStaticColor = "white"
    ActionButtonStaticColorBlack ActionButtonStaticColor = "black"
)
// ActionButtonTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type ActionButtonTarget string

// ActionButtonTarget values
const (
    ActionButtonTarget_blank ActionButtonTarget = "_blank"
    ActionButtonTarget_parent ActionButtonTarget = "_parent"
    ActionButtonTarget_self ActionButtonTarget = "_self"
    ActionButtonTarget_top ActionButtonTarget = "_top"
)
// ActionButtonType represents the The default behavior of the button. Possible values are: button (default), submit, and reset.
type ActionButtonType string

// ActionButtonType values
const (
    ActionButtonTypeButton ActionButtonType = "button"
    ActionButtonTypeSubmit ActionButtonType = "submit"
    ActionButtonTypeReset ActionButtonType = "reset"
)

// spectrumActionButton represents an sp-action-button component
type spectrumActionButton struct {
    app.Compo
    *styler[*spectrumActionButton]

    // Properties
    // Whether the button appears in an active state
    PropActive bool
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download.
    PropDownload string
    // Whether to apply the emphasized visual style
    PropEmphasized bool
    // Whether to display a visual affordance indicating longpress functionality
    PropHoldAffordance bool
    // The URL that the hyperlink points to.
    PropHref string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // Whether to apply the quiet visual style
    PropQuiet bool
    // How much of the referrer to send when following the link.
    PropReferrerpolicy ActionButtonReferrerpolicy
    // The relationship of the linked URL as space-separated link types.
    PropRel string
    // The ARIA role for the action button
    PropRole string
    // Whether an Action Button with role='button' should also be aria-pressed='true'
    PropSelected bool
    // The static color variant to use for the action button.
    PropStaticColor ActionButtonStaticColor
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
    PropTarget ActionButtonTarget
    // Whether to automatically manage the selected attribute on interaction and whether aria-pressed="false" should be used when selected === false
    PropToggles bool
    // The default behavior of the button. Possible values are: button (default), submit, and reset.
    PropType ActionButtonType
    // The value associated with the button
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnLongpress app.EventHandler
    PropOnClick app.EventHandler
}

// ActionButton creates a new sp-action-button component
func ActionButton() *spectrumActionButton {
    element := &spectrumActionButton{
        PropActive: false,
        PropDisabled: false,
        PropEmphasized: false,
        PropHoldAffordance: false,
        PropQuiet: false,
        PropRole: "button",
        PropSelected: false,
        PropTabindex: 0,
        PropToggles: false,
        PropType: ActionButtonTypeButton,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the button appears in an active state
func (c *spectrumActionButton) Active(active bool) *spectrumActionButton {
    c.PropActive = active
    return c
}

func (c *spectrumActionButton) SetActive() *spectrumActionButton {
    return c.Active(true)
}

// Disable this control. It will not receive focus or events
func (c *spectrumActionButton) Disabled(disabled bool) *spectrumActionButton {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumActionButton) SetDisabled() *spectrumActionButton {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumActionButton) Download(download string) *spectrumActionButton {
    c.PropDownload = download
    return c
}

// Whether to apply the emphasized visual style
func (c *spectrumActionButton) Emphasized(emphasized bool) *spectrumActionButton {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumActionButton) SetEmphasized() *spectrumActionButton {
    return c.Emphasized(true)
}

// Whether to display a visual affordance indicating longpress functionality
func (c *spectrumActionButton) HoldAffordance(holdAffordance bool) *spectrumActionButton {
    c.PropHoldAffordance = holdAffordance
    return c
}

func (c *spectrumActionButton) SetHoldAffordance() *spectrumActionButton {
    return c.HoldAffordance(true)
}

// The URL that the hyperlink points to.
func (c *spectrumActionButton) Href(href string) *spectrumActionButton {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumActionButton) Label(label string) *spectrumActionButton {
    c.PropLabel = label
    return c
}

// Whether to apply the quiet visual style
func (c *spectrumActionButton) Quiet(quiet bool) *spectrumActionButton {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumActionButton) SetQuiet() *spectrumActionButton {
    return c.Quiet(true)
}

// How much of the referrer to send when following the link.
func (c *spectrumActionButton) Referrerpolicy(referrerpolicy ActionButtonReferrerpolicy) *spectrumActionButton {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumActionButton) ReferrerpolicyNoReferrer() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyNoReferrer)
}
func (c *spectrumActionButton) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumActionButton) ReferrerpolicyOrigin() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyOrigin)
}
func (c *spectrumActionButton) ReferrerpolicyOriginWhenCrossOrigin() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumActionButton) ReferrerpolicySameOrigin() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicySameOrigin)
}
func (c *spectrumActionButton) ReferrerpolicyStrictOrigin() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyStrictOrigin)
}
func (c *spectrumActionButton) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumActionButton) ReferrerpolicyUnsafeUrl() *spectrumActionButton {
    return c.Referrerpolicy(ActionButtonReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types.
func (c *spectrumActionButton) Rel(rel string) *spectrumActionButton {
    c.PropRel = rel
    return c
}

// The ARIA role for the action button
func (c *spectrumActionButton) Role(role string) *spectrumActionButton {
    c.PropRole = role
    return c
}

// Whether an Action Button with role='button' should also be aria-pressed='true'
func (c *spectrumActionButton) Selected(selected bool) *spectrumActionButton {
    c.PropSelected = selected
    return c
}

func (c *spectrumActionButton) SetSelected() *spectrumActionButton {
    return c.Selected(true)
}

// The static color variant to use for the action button.
func (c *spectrumActionButton) StaticColor(staticColor ActionButtonStaticColor) *spectrumActionButton {
    c.PropStaticColor = staticColor
    return c
}

func (c *spectrumActionButton) StaticColorWhite() *spectrumActionButton {
    return c.StaticColor(ActionButtonStaticColorWhite)
}
func (c *spectrumActionButton) StaticColorBlack() *spectrumActionButton {
    return c.StaticColor(ActionButtonStaticColorBlack)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumActionButton) Tabindex(tabindex float64) *spectrumActionButton {
    c.PropTabindex = tabindex
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumActionButton) Target(target ActionButtonTarget) *spectrumActionButton {
    c.PropTarget = target
    return c
}

func (c *spectrumActionButton) Target_blank() *spectrumActionButton {
    return c.Target(ActionButtonTarget_blank)
}
func (c *spectrumActionButton) Target_parent() *spectrumActionButton {
    return c.Target(ActionButtonTarget_parent)
}
func (c *spectrumActionButton) Target_self() *spectrumActionButton {
    return c.Target(ActionButtonTarget_self)
}
func (c *spectrumActionButton) Target_top() *spectrumActionButton {
    return c.Target(ActionButtonTarget_top)
}
// Whether to automatically manage the selected attribute on interaction and whether aria-pressed="false" should be used when selected === false
func (c *spectrumActionButton) Toggles(toggles bool) *spectrumActionButton {
    c.PropToggles = toggles
    return c
}

func (c *spectrumActionButton) SetToggles() *spectrumActionButton {
    return c.Toggles(true)
}

// The default behavior of the button. Possible values are: button (default), submit, and reset.
func (c *spectrumActionButton) Type(typeValue ActionButtonType) *spectrumActionButton {
    c.PropType = typeValue
    return c
}

func (c *spectrumActionButton) TypeButton() *spectrumActionButton {
    return c.Type(ActionButtonTypeButton)
}
func (c *spectrumActionButton) TypeSubmit() *spectrumActionButton {
    return c.Type(ActionButtonTypeSubmit)
}
func (c *spectrumActionButton) TypeReset() *spectrumActionButton {
    return c.Type(ActionButtonTypeReset)
}
// The value associated with the button
func (c *spectrumActionButton) Value(value string) *spectrumActionButton {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumActionButton) Text(text string) *spectrumActionButton {
    c.PropText = text
    return c
}


// The icon to use for Action Button
func (c *spectrumActionButton) Icon(content app.UI) *spectrumActionButton {
    c.PropIconSlot = content

    return c
}



// Announces a change in the selected property of an action button
func (c *spectrumActionButton) OnChange(handler app.EventHandler) *spectrumActionButton {
    c.PropOnChange = handler

    return c
}

// Synthesizes a longpress interaction that signifies a pointerdown event that is >=300ms or a keyboard event where code is Space or code is ArrowDown while altKey===true.
func (c *spectrumActionButton) OnLongpress(handler app.EventHandler) *spectrumActionButton {
    c.PropOnLongpress = handler

    return c
}

// Fired when the button is clicked
func (c *spectrumActionButton) OnClick(handler app.EventHandler) *spectrumActionButton {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-action-button component
func (c *spectrumActionButton) Render() app.UI {
    element := app.Elem("sp-action-button")

    // Set attributes
    if c.PropActive {
        element = element.Attr("active", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
    }
    if c.PropHoldAffordance {
        element = element.Attr("hold-affordance", true)
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
    if c.PropRole != "" {
        element = element.Attr("role", c.PropRole)
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
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
    if c.PropToggles {
        element = element.Attr("toggles", true)
    }
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnLongpress != nil {
        element = element.On("longpress", c.PropOnLongpress)
    }
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