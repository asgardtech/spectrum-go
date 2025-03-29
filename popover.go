package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumPopover represents an sp-popover component
type spectrumPopover struct {
	app.Compo

	// Properties
	PropOpen      bool
	PropPlacement OverlayPlacement
	PropTip       bool
	PropDialog    bool
	PropVariant   string

	// Child content
	PropChildren []app.UI
}

// Popover creates a new popover component
func Popover() *spectrumPopover {
	return &spectrumPopover{}
}

// Open sets whether the popover is visible
func (p *spectrumPopover) Open(open bool) *spectrumPopover {
	p.PropOpen = open
	return p
}

// Placement sets the placement of the popover relative to its trigger
func (p *spectrumPopover) Placement(placement OverlayPlacement) *spectrumPopover {
	p.PropPlacement = placement
	return p
}

// Tip sets whether the popover has a tip pointing to the trigger
func (p *spectrumPopover) Tip(tip bool) *spectrumPopover {
	p.PropTip = tip
	return p
}

// Dialog sets whether the popover should have dialog styling
func (p *spectrumPopover) Dialog(dialog bool) *spectrumPopover {
	p.PropDialog = dialog
	return p
}

// Variant sets the visual variant of the popover
func (p *spectrumPopover) Variant(variant string) *spectrumPopover {
	p.PropVariant = variant
	return p
}

// Child adds a child element to the popover
func (p *spectrumPopover) Child(child app.UI) *spectrumPopover {
	p.PropChildren = append(p.PropChildren, child)
	return p
}

// Children sets all child elements of the popover
func (p *spectrumPopover) Children(children ...app.UI) *spectrumPopover {
	p.PropChildren = children
	return p
}

// Render renders the popover component
func (p *spectrumPopover) Render() app.UI {
	popover := app.Elem("sp-popover")

	// Add properties
	if p.PropOpen {
		popover = popover.Attr("open", true)
	}
	if p.PropPlacement != "" {
		popover = popover.Attr("placement", string(p.PropPlacement))
	}
	if p.PropTip {
		popover = popover.Attr("tip", true)
	}
	if p.PropDialog {
		popover = popover.Attr("dialog", true)
	}
	if p.PropVariant != "" {
		popover = popover.Attr("variant", p.PropVariant)
	}

	// Add child elements if provided
	if len(p.PropChildren) > 0 {
		popover = popover.Body(p.PropChildren...)
	}

	return popover
}

// Convenience methods for placement

// PlacementTop sets placement to top
func (p *spectrumPopover) PlacementTop() *spectrumPopover {
	return p.Placement(OverlayPlacementTop)
}

// PlacementTopStart sets placement to top-start
func (p *spectrumPopover) PlacementTopStart() *spectrumPopover {
	return p.Placement(OverlayPlacementTopStart)
}

// PlacementTopEnd sets placement to top-end
func (p *spectrumPopover) PlacementTopEnd() *spectrumPopover {
	return p.Placement(OverlayPlacementTopEnd)
}

// PlacementRight sets placement to right
func (p *spectrumPopover) PlacementRight() *spectrumPopover {
	return p.Placement(OverlayPlacementRight)
}

// PlacementRightStart sets placement to right-start
func (p *spectrumPopover) PlacementRightStart() *spectrumPopover {
	return p.Placement(OverlayPlacementRightStart)
}

// PlacementRightEnd sets placement to right-end
func (p *spectrumPopover) PlacementRightEnd() *spectrumPopover {
	return p.Placement(OverlayPlacementRightEnd)
}

// PlacementBottom sets placement to bottom
func (p *spectrumPopover) PlacementBottom() *spectrumPopover {
	return p.Placement(OverlayPlacementBottom)
}

// PlacementBottomStart sets placement to bottom-start
func (p *spectrumPopover) PlacementBottomStart() *spectrumPopover {
	return p.Placement(OverlayPlacementBottomStart)
}

// PlacementBottomEnd sets placement to bottom-end
func (p *spectrumPopover) PlacementBottomEnd() *spectrumPopover {
	return p.Placement(OverlayPlacementBottomEnd)
}

// PlacementLeft sets placement to left
func (p *spectrumPopover) PlacementLeft() *spectrumPopover {
	return p.Placement(OverlayPlacementLeft)
}

// PlacementLeftStart sets placement to left-start
func (p *spectrumPopover) PlacementLeftStart() *spectrumPopover {
	return p.Placement(OverlayPlacementLeftStart)
}

// PlacementLeftEnd sets placement to left-end
func (p *spectrumPopover) PlacementLeftEnd() *spectrumPopover {
	return p.Placement(OverlayPlacementLeftEnd)
}
