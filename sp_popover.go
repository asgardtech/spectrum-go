package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PopoverPlacement represents the The placement of the popover and its tip relative to its trigger element
type PopoverPlacement string

// PopoverPlacement values
const (
    PopoverPlacementTop PopoverPlacement = "top"
    PopoverPlacementTopStart PopoverPlacement = "top-start"
    PopoverPlacementTopEnd PopoverPlacement = "top-end"
    PopoverPlacementRight PopoverPlacement = "right"
    PopoverPlacementRightStart PopoverPlacement = "right-start"
    PopoverPlacementRightEnd PopoverPlacement = "right-end"
    PopoverPlacementBottom PopoverPlacement = "bottom"
    PopoverPlacementBottomStart PopoverPlacement = "bottom-start"
    PopoverPlacementBottomEnd PopoverPlacement = "bottom-end"
    PopoverPlacementLeft PopoverPlacement = "left"
    PopoverPlacementLeftStart PopoverPlacement = "left-start"
    PopoverPlacementLeftEnd PopoverPlacement = "left-end"
)

// spectrumPopover represents an sp-popover component
type spectrumPopover struct {
    app.Compo
    *styler[*spectrumPopover]

    // Properties
    // Whether the popover is visible or not
    PropOpen bool
    // The placement of the popover and its tip relative to its trigger element
    PropPlacement PopoverPlacement
    // Whether to display the popover with a directional tip
    PropTip bool

    // Text content for the default slot
    PropText string

    // Content slots


}

// Popover creates a new sp-popover component
func Popover() *spectrumPopover {
    element := &spectrumPopover{
        PropOpen: false,
        PropPlacement: "",
        PropTip: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the popover is visible or not
func (c *spectrumPopover) Open(open bool) *spectrumPopover {
    c.PropOpen = open
    return c
}

func (c *spectrumPopover) SetOpen() *spectrumPopover {
    return c.Open(true)
}

// The placement of the popover and its tip relative to its trigger element
func (c *spectrumPopover) Placement(placement PopoverPlacement) *spectrumPopover {
    c.PropPlacement = placement
    return c
}

func (c *spectrumPopover) PlacementTop() *spectrumPopover {
    return c.Placement(PopoverPlacementTop)
}
func (c *spectrumPopover) PlacementTopStart() *spectrumPopover {
    return c.Placement(PopoverPlacementTopStart)
}
func (c *spectrumPopover) PlacementTopEnd() *spectrumPopover {
    return c.Placement(PopoverPlacementTopEnd)
}
func (c *spectrumPopover) PlacementRight() *spectrumPopover {
    return c.Placement(PopoverPlacementRight)
}
func (c *spectrumPopover) PlacementRightStart() *spectrumPopover {
    return c.Placement(PopoverPlacementRightStart)
}
func (c *spectrumPopover) PlacementRightEnd() *spectrumPopover {
    return c.Placement(PopoverPlacementRightEnd)
}
func (c *spectrumPopover) PlacementBottom() *spectrumPopover {
    return c.Placement(PopoverPlacementBottom)
}
func (c *spectrumPopover) PlacementBottomStart() *spectrumPopover {
    return c.Placement(PopoverPlacementBottomStart)
}
func (c *spectrumPopover) PlacementBottomEnd() *spectrumPopover {
    return c.Placement(PopoverPlacementBottomEnd)
}
func (c *spectrumPopover) PlacementLeft() *spectrumPopover {
    return c.Placement(PopoverPlacementLeft)
}
func (c *spectrumPopover) PlacementLeftStart() *spectrumPopover {
    return c.Placement(PopoverPlacementLeftStart)
}
func (c *spectrumPopover) PlacementLeftEnd() *spectrumPopover {
    return c.Placement(PopoverPlacementLeftEnd)
}
// Whether to display the popover with a directional tip
func (c *spectrumPopover) Tip(tip bool) *spectrumPopover {
    c.PropTip = tip
    return c
}

func (c *spectrumPopover) SetTip() *spectrumPopover {
    return c.Tip(true)
}


// Text sets the text content for the default slot
func (c *spectrumPopover) Text(text string) *spectrumPopover {
    c.PropText = text
    return c
}





// Render renders the sp-popover component
func (c *spectrumPopover) Render() app.UI {
    element := app.Elem("sp-popover")

    // Set attributes
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropTip {
        element = element.Attr("tip", true)
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