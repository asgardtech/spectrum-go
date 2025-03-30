package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// InfieldButtonType represents the The default behavior of the button. Possible values are: 'button' (default), 'submit', and 'reset'.
type InfieldButtonType string

// InfieldButtonType values
const (
    InfieldButtonTypeButton InfieldButtonType = "button"
    InfieldButtonTypeSubmit InfieldButtonType = "submit"
    InfieldButtonTypeReset InfieldButtonType = "reset"
)
// InfieldButtonSize represents the Size of the button: 's', 'm', 'l', 'xl'
type InfieldButtonSize string

// InfieldButtonSize values
const (
    InfieldButtonSizeS InfieldButtonSize = "s"
    InfieldButtonSizeM InfieldButtonSize = "m"
    InfieldButtonSizeL InfieldButtonSize = "l"
    InfieldButtonSizeXl InfieldButtonSize = "xl"
)
// InfieldButtonInline represents the Position of the button relative to the field: 'start' or 'end'
type InfieldButtonInline string

// InfieldButtonInline values
const (
    InfieldButtonInlineStart InfieldButtonInline = "start"
    InfieldButtonInlineEnd InfieldButtonInline = "end"
)
// InfieldButtonBlock represents the Position of the button in a vertical stack: 'start' or 'end'
type InfieldButtonBlock string

// InfieldButtonBlock values
const (
    InfieldButtonBlockStart InfieldButtonBlock = "start"
    InfieldButtonBlockEnd InfieldButtonBlock = "end"
)

// spectrumInfieldButton represents an sp-infield-button component
type spectrumInfieldButton struct {
    app.Compo
    *styler[*spectrumInfieldButton]

    // Properties
    // Whether the button appears in an active state
    PropActive bool
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download.
    PropDownload string
    // The URL that the hyperlink points to.
    PropHref string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // How much of the referrer to send when following the link.
    PropReferrerpolicy string
    // The relationship of the linked URL as space-separated link types.
    PropRel string
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
    PropTarget string
    // The default behavior of the button. Possible values are: 'button' (default), 'submit', and 'reset'.
    PropType InfieldButtonType
    // Size of the button: 's', 'm', 'l', 'xl'
    PropSize InfieldButtonSize
    // Position of the button relative to the field: 'start' or 'end'
    PropInline InfieldButtonInline
    // Position of the button in a vertical stack: 'start' or 'end'
    PropBlock InfieldButtonBlock
    // Whether the button has a diminished visual presence
    PropQuiet bool

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnClick app.EventHandler
}

// InfieldButton creates a new sp-infield-button component
func InfieldButton() *spectrumInfieldButton {
    element := &spectrumInfieldButton{
        PropActive: false,
        PropDisabled: false,
        PropDownload: "",
        PropHref: "",
        PropLabel: "",
        PropReferrerpolicy: "",
        PropRel: "",
        PropTabindex: 0,
        PropTarget: "",
        PropType: InfieldButtonTypeButton,
        PropSize: InfieldButtonSizeM,
        PropInline: "",
        PropBlock: "",
        PropQuiet: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the button appears in an active state
func (c *spectrumInfieldButton) Active(active bool) *spectrumInfieldButton {
    c.PropActive = active
    return c
}

func (c *spectrumInfieldButton) SetActive() *spectrumInfieldButton {
    return c.Active(true)
}

// Disable this control. It will not receive focus or events
func (c *spectrumInfieldButton) Disabled(disabled bool) *spectrumInfieldButton {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumInfieldButton) SetDisabled() *spectrumInfieldButton {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumInfieldButton) Download(download string) *spectrumInfieldButton {
    c.PropDownload = download
    return c
}

// The URL that the hyperlink points to.
func (c *spectrumInfieldButton) Href(href string) *spectrumInfieldButton {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumInfieldButton) Label(label string) *spectrumInfieldButton {
    c.PropLabel = label
    return c
}

// How much of the referrer to send when following the link.
func (c *spectrumInfieldButton) Referrerpolicy(referrerpolicy string) *spectrumInfieldButton {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

// The relationship of the linked URL as space-separated link types.
func (c *spectrumInfieldButton) Rel(rel string) *spectrumInfieldButton {
    c.PropRel = rel
    return c
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumInfieldButton) Tabindex(tabIndex float64) *spectrumInfieldButton {
    c.PropTabindex = tabIndex
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumInfieldButton) Target(target string) *spectrumInfieldButton {
    c.PropTarget = target
    return c
}

// The default behavior of the button. Possible values are: 'button' (default), 'submit', and 'reset'.
func (c *spectrumInfieldButton) Type(typeValue InfieldButtonType) *spectrumInfieldButton {
    c.PropType = typeValue
    return c
}

func (c *spectrumInfieldButton) TypeButton() *spectrumInfieldButton {
    return c.Type(InfieldButtonTypeButton)
}
func (c *spectrumInfieldButton) TypeSubmit() *spectrumInfieldButton {
    return c.Type(InfieldButtonTypeSubmit)
}
func (c *spectrumInfieldButton) TypeReset() *spectrumInfieldButton {
    return c.Type(InfieldButtonTypeReset)
}
// Size of the button: 's', 'm', 'l', 'xl'
func (c *spectrumInfieldButton) Size(size InfieldButtonSize) *spectrumInfieldButton {
    c.PropSize = size
    return c
}

func (c *spectrumInfieldButton) SizeS() *spectrumInfieldButton {
    return c.Size(InfieldButtonSizeS)
}
func (c *spectrumInfieldButton) SizeM() *spectrumInfieldButton {
    return c.Size(InfieldButtonSizeM)
}
func (c *spectrumInfieldButton) SizeL() *spectrumInfieldButton {
    return c.Size(InfieldButtonSizeL)
}
func (c *spectrumInfieldButton) SizeXl() *spectrumInfieldButton {
    return c.Size(InfieldButtonSizeXl)
}
// Position of the button relative to the field: 'start' or 'end'
func (c *spectrumInfieldButton) Inline(inline InfieldButtonInline) *spectrumInfieldButton {
    c.PropInline = inline
    return c
}

func (c *spectrumInfieldButton) InlineStart() *spectrumInfieldButton {
    return c.Inline(InfieldButtonInlineStart)
}
func (c *spectrumInfieldButton) InlineEnd() *spectrumInfieldButton {
    return c.Inline(InfieldButtonInlineEnd)
}
// Position of the button in a vertical stack: 'start' or 'end'
func (c *spectrumInfieldButton) Block(block InfieldButtonBlock) *spectrumInfieldButton {
    c.PropBlock = block
    return c
}

func (c *spectrumInfieldButton) BlockStart() *spectrumInfieldButton {
    return c.Block(InfieldButtonBlockStart)
}
func (c *spectrumInfieldButton) BlockEnd() *spectrumInfieldButton {
    return c.Block(InfieldButtonBlockEnd)
}
// Whether the button has a diminished visual presence
func (c *spectrumInfieldButton) Quiet(quiet bool) *spectrumInfieldButton {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumInfieldButton) SetQuiet() *spectrumInfieldButton {
    return c.Quiet(true)
}


// Text sets the text content for the default slot
func (c *spectrumInfieldButton) Text(text string) *spectrumInfieldButton {
    c.PropText = text
    return c
}


// icon element(s) to display at the start of the button
func (c *spectrumInfieldButton) Icon(content app.UI) *spectrumInfieldButton {
    c.PropIconSlot = content

    return c
}



// Fired when the button is clicked
func (c *spectrumInfieldButton) OnClick(handler app.EventHandler) *spectrumInfieldButton {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-infield-button component
func (c *spectrumInfieldButton) Render() app.UI {
    element := app.Elem("sp-infield-button")

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
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropReferrerpolicy != "" {
        element = element.Attr("referrerpolicy", c.PropReferrerpolicy)
    }
    if c.PropRel != "" {
        element = element.Attr("rel", c.PropRel)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", c.PropTarget)
    }
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropInline != "" {
        element = element.Attr("inline", string(c.PropInline))
    }
    if c.PropBlock != "" {
        element = element.Attr("block", string(c.PropBlock))
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
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