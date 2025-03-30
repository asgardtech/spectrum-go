package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ContextualHelpPlacement represents the The placement of the popover relative to the trigger.
type ContextualHelpPlacement string

// ContextualHelpPlacement values
const (
    ContextualHelpPlacementTop ContextualHelpPlacement = "top"
    ContextualHelpPlacementTopStart ContextualHelpPlacement = "top-start"
    ContextualHelpPlacementTopEnd ContextualHelpPlacement = "top-end"
    ContextualHelpPlacementRight ContextualHelpPlacement = "right"
    ContextualHelpPlacementRightStart ContextualHelpPlacement = "right-start"
    ContextualHelpPlacementRightEnd ContextualHelpPlacement = "right-end"
    ContextualHelpPlacementBottom ContextualHelpPlacement = "bottom"
    ContextualHelpPlacementBottomStart ContextualHelpPlacement = "bottom-start"
    ContextualHelpPlacementBottomEnd ContextualHelpPlacement = "bottom-end"
    ContextualHelpPlacementLeft ContextualHelpPlacement = "left"
    ContextualHelpPlacementLeftStart ContextualHelpPlacement = "left-start"
    ContextualHelpPlacementLeftEnd ContextualHelpPlacement = "left-end"
)
// ContextualHelpVariant represents the The variant property applies specific styling on the action button trigger. Values: 'info', 'help'
type ContextualHelpVariant string

// ContextualHelpVariant values
const (
    ContextualHelpVariantInfo ContextualHelpVariant = "info"
    ContextualHelpVariantHelp ContextualHelpVariant = "help"
)

// spectrumContextualHelp represents an sp-contextual-help component
type spectrumContextualHelp struct {
    app.Compo
    *styler[*spectrumContextualHelp]

    // Properties
    // Provides an accessible name for the action button trigger.
    PropLabel string
    // The offset property accepts either a single number, to define the offset of the Popover along the main axis from the action button, or 2-tuple, to define the offset along the main axis and the cross axis.
    PropOffset string
    // The placement of the popover relative to the trigger.
    PropPlacement ContextualHelpPlacement
    // The variant property applies specific styling on the action button trigger. Values: 'info', 'help'
    PropVariant ContextualHelpVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropHeadingSlot app.UI
    PropLinkSlot app.UI


    // Event handlers
    PropOnSpOpen app.EventHandler
    PropOnSpClosed app.EventHandler
}

// ContextualHelp creates a new sp-contextual-help component
func ContextualHelp() *spectrumContextualHelp {
    element := &spectrumContextualHelp{
        PropLabel: "",
        PropOffset: "0",
        PropPlacement: ContextualHelpPlacementBottomStart,
        PropVariant: ContextualHelpVariantInfo,
    }

    element.styler = newStyler(element)

    return element
}

// Provides an accessible name for the action button trigger.
func (c *spectrumContextualHelp) Label(label string) *spectrumContextualHelp {
    c.PropLabel = label
    return c
}

// The offset property accepts either a single number, to define the offset of the Popover along the main axis from the action button, or 2-tuple, to define the offset along the main axis and the cross axis.
func (c *spectrumContextualHelp) Offset(offset string) *spectrumContextualHelp {
    c.PropOffset = offset
    return c
}

// The placement of the popover relative to the trigger.
func (c *spectrumContextualHelp) Placement(placement ContextualHelpPlacement) *spectrumContextualHelp {
    c.PropPlacement = placement
    return c
}

func (c *spectrumContextualHelp) PlacementTop() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementTop)
}
func (c *spectrumContextualHelp) PlacementTopStart() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementTopStart)
}
func (c *spectrumContextualHelp) PlacementTopEnd() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementTopEnd)
}
func (c *spectrumContextualHelp) PlacementRight() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementRight)
}
func (c *spectrumContextualHelp) PlacementRightStart() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementRightStart)
}
func (c *spectrumContextualHelp) PlacementRightEnd() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementRightEnd)
}
func (c *spectrumContextualHelp) PlacementBottom() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementBottom)
}
func (c *spectrumContextualHelp) PlacementBottomStart() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementBottomStart)
}
func (c *spectrumContextualHelp) PlacementBottomEnd() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementBottomEnd)
}
func (c *spectrumContextualHelp) PlacementLeft() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementLeft)
}
func (c *spectrumContextualHelp) PlacementLeftStart() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementLeftStart)
}
func (c *spectrumContextualHelp) PlacementLeftEnd() *spectrumContextualHelp {
    return c.Placement(ContextualHelpPlacementLeftEnd)
}
// The variant property applies specific styling on the action button trigger. Values: 'info', 'help'
func (c *spectrumContextualHelp) Variant(variant ContextualHelpVariant) *spectrumContextualHelp {
    c.PropVariant = variant
    return c
}

func (c *spectrumContextualHelp) VariantInfo() *spectrumContextualHelp {
    return c.Variant(ContextualHelpVariantInfo)
}
func (c *spectrumContextualHelp) VariantHelp() *spectrumContextualHelp {
    return c.Variant(ContextualHelpVariantHelp)
}

// Text sets the text content for the default slot
func (c *spectrumContextualHelp) Text(text string) *spectrumContextualHelp {
    c.PropText = text
    return c
}


// content to display as the heading of the popover
func (c *spectrumContextualHelp) Heading(content app.UI) *spectrumContextualHelp {
    c.PropHeadingSlot = content

    return c
}

// link to additional informations
func (c *spectrumContextualHelp) Link(content app.UI) *spectrumContextualHelp {
    c.PropLinkSlot = content

    return c
}



// Fired when the popover opens
func (c *spectrumContextualHelp) OnSpOpen(handler app.EventHandler) *spectrumContextualHelp {
    c.PropOnSpOpen = handler

    return c
}

// Fired when the popover closes
func (c *spectrumContextualHelp) OnSpClosed(handler app.EventHandler) *spectrumContextualHelp {
    c.PropOnSpClosed = handler

    return c
}


// Render renders the sp-contextual-help component
func (c *spectrumContextualHelp) Render() app.UI {
    element := app.Elem("sp-contextual-help")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropOffset != "" {
        element = element.Attr("offset", c.PropOffset)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnSpOpen != nil {
        element = element.On("sp-open", c.PropOnSpOpen)
    }
    if c.PropOnSpClosed != nil {
        element = element.On("sp-closed", c.PropOnSpClosed)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add heading slot
    if c.PropHeadingSlot != nil {
        slotElem := c.PropHeadingSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("heading")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "heading").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add link slot
    if c.PropLinkSlot != nil {
        slotElem := c.PropLinkSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("link")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "link").
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