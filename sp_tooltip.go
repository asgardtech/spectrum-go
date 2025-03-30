package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TooltipPlacement represents the Position of the tooltip relative to its target
type TooltipPlacement string

// TooltipPlacement values
const (
    TooltipPlacementTop TooltipPlacement = "top"
    TooltipPlacementTopStart TooltipPlacement = "top-start"
    TooltipPlacementTopEnd TooltipPlacement = "top-end"
    TooltipPlacementRight TooltipPlacement = "right"
    TooltipPlacementRightStart TooltipPlacement = "right-start"
    TooltipPlacementRightEnd TooltipPlacement = "right-end"
    TooltipPlacementBottom TooltipPlacement = "bottom"
    TooltipPlacementBottomStart TooltipPlacement = "bottom-start"
    TooltipPlacementBottomEnd TooltipPlacement = "bottom-end"
    TooltipPlacementLeft TooltipPlacement = "left"
    TooltipPlacementLeftStart TooltipPlacement = "left-start"
    TooltipPlacementLeftEnd TooltipPlacement = "left-end"
)
// TooltipVariant represents the Visual variant of the tooltip
type TooltipVariant string

// TooltipVariant values
const (
    TooltipVariantInfo TooltipVariant = "info"
    TooltipVariantPositive TooltipVariant = "positive"
    TooltipVariantNegative TooltipVariant = "negative"
)

// spectrumTooltip represents an sp-tooltip component
type spectrumTooltip struct {
    app.Compo
    *styler[*spectrumTooltip]

    // Properties
    // Distance between the tooltip and its target element
    PropOffset float64
    // Whether the tooltip is currently visible
    PropOpen bool
    // Position of the tooltip relative to its target
    PropPlacement TooltipPlacement
    // Automatically bind to the parent element of the assigned slot or the parent element of the sp-tooltip. Without this, you must provide your own overlay-trigger.
    PropSelfManaged bool
    // Distance from the edge of the tooltip to the tip
    PropTipPadding float64
    // Visual variant of the tooltip
    PropVariant TooltipVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnTransitionend app.EventHandler
}

// Tooltip creates a new sp-tooltip component
func Tooltip() *spectrumTooltip {
    element := &spectrumTooltip{
        PropOffset: 0,
        PropOpen: false,
        PropPlacement: TooltipPlacementTop,
        PropSelfManaged: false,
        PropTipPadding: 0,
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Distance between the tooltip and its target element
func (c *spectrumTooltip) Offset(offset float64) *spectrumTooltip {
    c.PropOffset = offset
    return c
}

// Whether the tooltip is currently visible
func (c *spectrumTooltip) Open(open bool) *spectrumTooltip {
    c.PropOpen = open
    return c
}

func (c *spectrumTooltip) SetOpen() *spectrumTooltip {
    return c.Open(true)
}

// Position of the tooltip relative to its target
func (c *spectrumTooltip) Placement(placement TooltipPlacement) *spectrumTooltip {
    c.PropPlacement = placement
    return c
}

func (c *spectrumTooltip) PlacementTop() *spectrumTooltip {
    return c.Placement(TooltipPlacementTop)
}
func (c *spectrumTooltip) PlacementTopStart() *spectrumTooltip {
    return c.Placement(TooltipPlacementTopStart)
}
func (c *spectrumTooltip) PlacementTopEnd() *spectrumTooltip {
    return c.Placement(TooltipPlacementTopEnd)
}
func (c *spectrumTooltip) PlacementRight() *spectrumTooltip {
    return c.Placement(TooltipPlacementRight)
}
func (c *spectrumTooltip) PlacementRightStart() *spectrumTooltip {
    return c.Placement(TooltipPlacementRightStart)
}
func (c *spectrumTooltip) PlacementRightEnd() *spectrumTooltip {
    return c.Placement(TooltipPlacementRightEnd)
}
func (c *spectrumTooltip) PlacementBottom() *spectrumTooltip {
    return c.Placement(TooltipPlacementBottom)
}
func (c *spectrumTooltip) PlacementBottomStart() *spectrumTooltip {
    return c.Placement(TooltipPlacementBottomStart)
}
func (c *spectrumTooltip) PlacementBottomEnd() *spectrumTooltip {
    return c.Placement(TooltipPlacementBottomEnd)
}
func (c *spectrumTooltip) PlacementLeft() *spectrumTooltip {
    return c.Placement(TooltipPlacementLeft)
}
func (c *spectrumTooltip) PlacementLeftStart() *spectrumTooltip {
    return c.Placement(TooltipPlacementLeftStart)
}
func (c *spectrumTooltip) PlacementLeftEnd() *spectrumTooltip {
    return c.Placement(TooltipPlacementLeftEnd)
}
// Automatically bind to the parent element of the assigned slot or the parent element of the sp-tooltip. Without this, you must provide your own overlay-trigger.
func (c *spectrumTooltip) SelfManaged(selfManaged bool) *spectrumTooltip {
    c.PropSelfManaged = selfManaged
    return c
}

func (c *spectrumTooltip) SetSelfManaged() *spectrumTooltip {
    return c.SelfManaged(true)
}

// Distance from the edge of the tooltip to the tip
func (c *spectrumTooltip) TipPadding(tipPadding float64) *spectrumTooltip {
    c.PropTipPadding = tipPadding
    return c
}

// Visual variant of the tooltip
func (c *spectrumTooltip) Variant(variant TooltipVariant) *spectrumTooltip {
    c.PropVariant = variant
    return c
}

func (c *spectrumTooltip) VariantInfo() *spectrumTooltip {
    return c.Variant(TooltipVariantInfo)
}
func (c *spectrumTooltip) VariantPositive() *spectrumTooltip {
    return c.Variant(TooltipVariantPositive)
}
func (c *spectrumTooltip) VariantNegative() *spectrumTooltip {
    return c.Variant(TooltipVariantNegative)
}

// Text sets the text content for the default slot
func (c *spectrumTooltip) Text(text string) *spectrumTooltip {
    c.PropText = text
    return c
}


// The icon element appearing at the start of the label
func (c *spectrumTooltip) Icon(content app.UI) *spectrumTooltip {
    c.PropIconSlot = content

    return c
}



// Fired when a transition completes
func (c *spectrumTooltip) OnTransitionend(handler app.EventHandler) *spectrumTooltip {
    c.PropOnTransitionend = handler

    return c
}


// Render renders the sp-tooltip component
func (c *spectrumTooltip) Render() app.UI {
    element := app.Elem("sp-tooltip")

    // Set attributes
    if c.PropOffset != 0 {
        element = element.Attr("offset", c.PropOffset)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropSelfManaged {
        element = element.Attr("self-managed", true)
    }
    if c.PropTipPadding != 0 {
        element = element.Attr("tip-padding", c.PropTipPadding)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnTransitionend != nil {
        element = element.On("transitionend", c.PropOnTransitionend)
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