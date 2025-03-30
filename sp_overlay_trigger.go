package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// OverlayTriggerPlacement represents the The placement of the overlay relative to the trigger
type OverlayTriggerPlacement string

// OverlayTriggerPlacement values
const (
    OverlayTriggerPlacementTop OverlayTriggerPlacement = "top"
    OverlayTriggerPlacementTopStart OverlayTriggerPlacement = "top-start"
    OverlayTriggerPlacementTopEnd OverlayTriggerPlacement = "top-end"
    OverlayTriggerPlacementRight OverlayTriggerPlacement = "right"
    OverlayTriggerPlacementRightStart OverlayTriggerPlacement = "right-start"
    OverlayTriggerPlacementRightEnd OverlayTriggerPlacement = "right-end"
    OverlayTriggerPlacementBottom OverlayTriggerPlacement = "bottom"
    OverlayTriggerPlacementBottomStart OverlayTriggerPlacement = "bottom-start"
    OverlayTriggerPlacementBottomEnd OverlayTriggerPlacement = "bottom-end"
    OverlayTriggerPlacementLeft OverlayTriggerPlacement = "left"
    OverlayTriggerPlacementLeftStart OverlayTriggerPlacement = "left-start"
    OverlayTriggerPlacementLeftEnd OverlayTriggerPlacement = "left-end"
)
// OverlayTriggerReceivesfocus represents the How focus should be handled when the overlay opens
type OverlayTriggerReceivesfocus string

// OverlayTriggerReceivesfocus values
const (
    OverlayTriggerReceivesfocusTrue OverlayTriggerReceivesfocus = "true"
    OverlayTriggerReceivesfocusFalse OverlayTriggerReceivesfocus = "false"
    OverlayTriggerReceivesfocusAuto OverlayTriggerReceivesfocus = "auto"
)

// spectrumOverlayTrigger represents an sp-overlay-trigger component
type spectrumOverlayTrigger struct {
    app.Compo
    *styler[*spectrumOverlayTrigger]

    // Properties
    // Whether the overlay trigger is disabled
    PropDisabled bool
    // The distance between the overlay and the trigger
    PropOffset float64
    // The type of overlay that is currently open
    PropOpen string
    // The placement of the overlay relative to the trigger
    PropPlacement OverlayTriggerPlacement
    // How focus should be handled when the overlay opens
    PropReceivesfocus OverlayTriggerReceivesfocus
    // Explicitly declares which types of overlays will be used (space-separated: 'click hover longpress')
    PropTriggeredby string
    // How the overlay appears in the tab order
    PropType string


    // Content slots
    PropTriggerSlot app.UI
    PropClickContentSlot app.UI
    PropHoverContentSlot app.UI
    PropLongpressContentSlot app.UI
    PropLongpressDescribedbyDescriptorSlot app.UI


    // Event handlers
    PropOnSpOpened app.EventHandler
    PropOnSpClosed app.EventHandler
}

// OverlayTrigger creates a new sp-overlay-trigger component
func OverlayTrigger() *spectrumOverlayTrigger {
    element := &spectrumOverlayTrigger{
        PropDisabled: false,
        PropOffset: 6,
        PropOpen: "",
        PropPlacement: "",
        PropReceivesfocus: OverlayTriggerReceivesfocusAuto,
        PropTriggeredby: "",
        PropType: "",
    }

    element.styler = newStyler(element)

    return element
}

// Whether the overlay trigger is disabled
func (c *spectrumOverlayTrigger) Disabled(disabled bool) *spectrumOverlayTrigger {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumOverlayTrigger) SetDisabled() *spectrumOverlayTrigger {
    return c.Disabled(true)
}

// The distance between the overlay and the trigger
func (c *spectrumOverlayTrigger) Offset(offset float64) *spectrumOverlayTrigger {
    c.PropOffset = offset
    return c
}

// The type of overlay that is currently open
func (c *spectrumOverlayTrigger) Open(open string) *spectrumOverlayTrigger {
    c.PropOpen = open
    return c
}

// The placement of the overlay relative to the trigger
func (c *spectrumOverlayTrigger) Placement(placement OverlayTriggerPlacement) *spectrumOverlayTrigger {
    c.PropPlacement = placement
    return c
}

func (c *spectrumOverlayTrigger) PlacementTop() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementTop)
}
func (c *spectrumOverlayTrigger) PlacementTopStart() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementTopStart)
}
func (c *spectrumOverlayTrigger) PlacementTopEnd() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementTopEnd)
}
func (c *spectrumOverlayTrigger) PlacementRight() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementRight)
}
func (c *spectrumOverlayTrigger) PlacementRightStart() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementRightStart)
}
func (c *spectrumOverlayTrigger) PlacementRightEnd() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementRightEnd)
}
func (c *spectrumOverlayTrigger) PlacementBottom() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementBottom)
}
func (c *spectrumOverlayTrigger) PlacementBottomStart() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementBottomStart)
}
func (c *spectrumOverlayTrigger) PlacementBottomEnd() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementBottomEnd)
}
func (c *spectrumOverlayTrigger) PlacementLeft() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementLeft)
}
func (c *spectrumOverlayTrigger) PlacementLeftStart() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementLeftStart)
}
func (c *spectrumOverlayTrigger) PlacementLeftEnd() *spectrumOverlayTrigger {
    return c.Placement(OverlayTriggerPlacementLeftEnd)
}
// How focus should be handled when the overlay opens
func (c *spectrumOverlayTrigger) Receivesfocus(receivesFocus OverlayTriggerReceivesfocus) *spectrumOverlayTrigger {
    c.PropReceivesfocus = receivesFocus
    return c
}

func (c *spectrumOverlayTrigger) ReceivesfocusTrue() *spectrumOverlayTrigger {
    return c.Receivesfocus(OverlayTriggerReceivesfocusTrue)
}
func (c *spectrumOverlayTrigger) ReceivesfocusFalse() *spectrumOverlayTrigger {
    return c.Receivesfocus(OverlayTriggerReceivesfocusFalse)
}
func (c *spectrumOverlayTrigger) ReceivesfocusAuto() *spectrumOverlayTrigger {
    return c.Receivesfocus(OverlayTriggerReceivesfocusAuto)
}
// Explicitly declares which types of overlays will be used (space-separated: 'click hover longpress')
func (c *spectrumOverlayTrigger) Triggeredby(triggeredBy string) *spectrumOverlayTrigger {
    c.PropTriggeredby = triggeredBy
    return c
}

// How the overlay appears in the tab order
func (c *spectrumOverlayTrigger) Type(typeValue string) *spectrumOverlayTrigger {
    c.PropType = typeValue
    return c
}



// The content that will trigger the various overlays
func (c *spectrumOverlayTrigger) Trigger(content app.UI) *spectrumOverlayTrigger {
    c.PropTriggerSlot = content

    return c
}

// The content that will be displayed on click
func (c *spectrumOverlayTrigger) ClickContent(content app.UI) *spectrumOverlayTrigger {
    c.PropClickContentSlot = content

    return c
}

// The content that will be displayed on hover
func (c *spectrumOverlayTrigger) HoverContent(content app.UI) *spectrumOverlayTrigger {
    c.PropHoverContentSlot = content

    return c
}

// The content that will be displayed on longpress
func (c *spectrumOverlayTrigger) LongpressContent(content app.UI) *spectrumOverlayTrigger {
    c.PropLongpressContentSlot = content

    return c
}

// Description for longpress content
func (c *spectrumOverlayTrigger) LongpressDescribedbyDescriptor(content app.UI) *spectrumOverlayTrigger {
    c.PropLongpressDescribedbyDescriptorSlot = content

    return c
}



// Announces that the overlay has been opened
func (c *spectrumOverlayTrigger) OnSpOpened(handler app.EventHandler) *spectrumOverlayTrigger {
    c.PropOnSpOpened = handler

    return c
}

// Announces that the overlay has been closed
func (c *spectrumOverlayTrigger) OnSpClosed(handler app.EventHandler) *spectrumOverlayTrigger {
    c.PropOnSpClosed = handler

    return c
}


// Render renders the sp-overlay-trigger component
func (c *spectrumOverlayTrigger) Render() app.UI {
    element := app.Elem("sp-overlay-trigger")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropOffset != 0 {
        element = element.Attr("offset", c.PropOffset)
    }
    if c.PropOpen != "" {
        element = element.Attr("open", c.PropOpen)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropReceivesfocus != "" {
        element = element.Attr("receivesFocus", string(c.PropReceivesfocus))
    }
    if c.PropTriggeredby != "" {
        element = element.Attr("triggeredBy", c.PropTriggeredby)
    }
    if c.PropType != "" {
        element = element.Attr("type", c.PropType)
    }

    // Add event handlers
    if c.PropOnSpOpened != nil {
        element = element.On("sp-opened", c.PropOnSpOpened)
    }
    if c.PropOnSpClosed != nil {
        element = element.On("sp-closed", c.PropOnSpClosed)
    }

    // Add slots and children
    slotElements := []app.UI{}


    // Add trigger slot
    if c.PropTriggerSlot != nil {
        slotElem := c.PropTriggerSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("trigger")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "trigger").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add click-content slot
    if c.PropClickContentSlot != nil {
        slotElem := c.PropClickContentSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("click-content")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "click-content").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add hover-content slot
    if c.PropHoverContentSlot != nil {
        slotElem := c.PropHoverContentSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("hover-content")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "hover-content").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add longpress-content slot
    if c.PropLongpressContentSlot != nil {
        slotElem := c.PropLongpressContentSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("longpress-content")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "longpress-content").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add longpress-describedby-descriptor slot
    if c.PropLongpressDescribedbyDescriptorSlot != nil {
        slotElem := c.PropLongpressDescribedbyDescriptorSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("longpress-describedby-descriptor")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "longpress-describedby-descriptor").
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