package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumPopover represents an sp-popover component
type SpectrumPopover struct {
	app.Compo

	// Properties
	open      bool
	placement OverlayPlacement
	tip       bool
	dialog    bool
	variant   string

	// Child content
	children []app.UI
}

// Popover creates a new popover component
func Popover() *SpectrumPopover {
	return &SpectrumPopover{}
}

// Open sets whether the popover is visible
func (p *SpectrumPopover) Open(open bool) *SpectrumPopover {
	p.open = open
	return p
}

// Placement sets the placement of the popover relative to its trigger
func (p *SpectrumPopover) Placement(placement OverlayPlacement) *SpectrumPopover {
	p.placement = placement
	return p
}

// Tip sets whether the popover has a tip pointing to the trigger
func (p *SpectrumPopover) Tip(tip bool) *SpectrumPopover {
	p.tip = tip
	return p
}

// Dialog sets whether the popover should have dialog styling
func (p *SpectrumPopover) Dialog(dialog bool) *SpectrumPopover {
	p.dialog = dialog
	return p
}

// Variant sets the visual variant of the popover
func (p *SpectrumPopover) Variant(variant string) *SpectrumPopover {
	p.variant = variant
	return p
}

// Child adds a child element to the popover
func (p *SpectrumPopover) Child(child app.UI) *SpectrumPopover {
	p.children = append(p.children, child)
	return p
}

// Children sets all child elements of the popover
func (p *SpectrumPopover) Children(children ...app.UI) *SpectrumPopover {
	p.children = children
	return p
}

// Render renders the popover component
func (p *SpectrumPopover) Render() app.UI {
	popover := app.Elem("sp-popover")

	// Add properties
	if p.open {
		popover = popover.Attr("open", true)
	}
	if p.placement != "" {
		popover = popover.Attr("placement", string(p.placement))
	}
	if p.tip {
		popover = popover.Attr("tip", true)
	}
	if p.dialog {
		popover = popover.Attr("dialog", true)
	}
	if p.variant != "" {
		popover = popover.Attr("variant", p.variant)
	}

	// Add child elements if provided
	if len(p.children) > 0 {
		popover = popover.Body(p.children...)
	}

	return popover
}

// Convenience methods for placement

// PlacementTop sets placement to top
func (p *SpectrumPopover) PlacementTop() *SpectrumPopover {
	return p.Placement(OverlayPlacementTop)
}

// PlacementTopStart sets placement to top-start
func (p *SpectrumPopover) PlacementTopStart() *SpectrumPopover {
	return p.Placement(OverlayPlacementTopStart)
}

// PlacementTopEnd sets placement to top-end
func (p *SpectrumPopover) PlacementTopEnd() *SpectrumPopover {
	return p.Placement(OverlayPlacementTopEnd)
}

// PlacementRight sets placement to right
func (p *SpectrumPopover) PlacementRight() *SpectrumPopover {
	return p.Placement(OverlayPlacementRight)
}

// PlacementRightStart sets placement to right-start
func (p *SpectrumPopover) PlacementRightStart() *SpectrumPopover {
	return p.Placement(OverlayPlacementRightStart)
}

// PlacementRightEnd sets placement to right-end
func (p *SpectrumPopover) PlacementRightEnd() *SpectrumPopover {
	return p.Placement(OverlayPlacementRightEnd)
}

// PlacementBottom sets placement to bottom
func (p *SpectrumPopover) PlacementBottom() *SpectrumPopover {
	return p.Placement(OverlayPlacementBottom)
}

// PlacementBottomStart sets placement to bottom-start
func (p *SpectrumPopover) PlacementBottomStart() *SpectrumPopover {
	return p.Placement(OverlayPlacementBottomStart)
}

// PlacementBottomEnd sets placement to bottom-end
func (p *SpectrumPopover) PlacementBottomEnd() *SpectrumPopover {
	return p.Placement(OverlayPlacementBottomEnd)
}

// PlacementLeft sets placement to left
func (p *SpectrumPopover) PlacementLeft() *SpectrumPopover {
	return p.Placement(OverlayPlacementLeft)
}

// PlacementLeftStart sets placement to left-start
func (p *SpectrumPopover) PlacementLeftStart() *SpectrumPopover {
	return p.Placement(OverlayPlacementLeftStart)
}

// PlacementLeftEnd sets placement to left-end
func (p *SpectrumPopover) PlacementLeftEnd() *SpectrumPopover {
	return p.Placement(OverlayPlacementLeftEnd)
}
