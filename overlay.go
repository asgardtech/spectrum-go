package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// OverlayPlacement represents the placement options for an overlay
type OverlayPlacement string

// Overlay placement values
const (
	OverlayPlacementTop         OverlayPlacement = "top"
	OverlayPlacementTopStart    OverlayPlacement = "top-start"
	OverlayPlacementTopEnd      OverlayPlacement = "top-end"
	OverlayPlacementBottom      OverlayPlacement = "bottom"
	OverlayPlacementBottomStart OverlayPlacement = "bottom-start"
	OverlayPlacementBottomEnd   OverlayPlacement = "bottom-end"
	OverlayPlacementLeft        OverlayPlacement = "left"
	OverlayPlacementLeftStart   OverlayPlacement = "left-start"
	OverlayPlacementLeftEnd     OverlayPlacement = "left-end"
	OverlayPlacementRight       OverlayPlacement = "right"
	OverlayPlacementRightStart  OverlayPlacement = "right-start"
	OverlayPlacementRightEnd    OverlayPlacement = "right-end"
)

// OverlayType represents the interaction type of an overlay
type OverlayType string

// Overlay type values
const (
	OverlayTypeAuto   OverlayType = "auto"
	OverlayTypeHint   OverlayType = "hint"
	OverlayTypeManual OverlayType = "manual"
	OverlayTypeModal  OverlayType = "modal"
	OverlayTypePage   OverlayType = "page"
)

// OverlayReceivesFocus represents how focus should be handled in an overlay
type OverlayReceivesFocus string

// Overlay receives focus values
const (
	OverlayReceivesFocusTrue  OverlayReceivesFocus = "true"
	OverlayReceivesFocusFalse OverlayReceivesFocus = "false"
	OverlayReceivesFocusAuto  OverlayReceivesFocus = "auto"
)

// OverlayTriggerInteraction represents the interaction type for triggering an overlay
type OverlayTriggerInteraction string

// Overlay trigger interaction values
const (
	OverlayTriggerInteractionClick     OverlayTriggerInteraction = "click"
	OverlayTriggerInteractionHover     OverlayTriggerInteraction = "hover"
	OverlayTriggerInteractionLongpress OverlayTriggerInteraction = "longpress"
)

// SpectrumOverlay represents an sp-overlay component
type SpectrumOverlay struct {
	app.Compo

	// Properties
	open               bool
	delayed            bool
	offset             interface{} // Can be a number or [number, number]
	placement          OverlayPlacement
	receivesFocus      OverlayReceivesFocus
	trigger            string
	triggerElement     app.UI
	triggerInteraction OverlayTriggerInteraction
	overlayType        OverlayType

	// Child element to display in the overlay
	child app.UI

	// Event handlers
	onOpened app.EventHandler
	onClosed app.EventHandler
}

// Overlay creates a new overlay component
func Overlay() *SpectrumOverlay {
	return &SpectrumOverlay{
		receivesFocus: OverlayReceivesFocusAuto, // Default receives-focus is "auto"
	}
}

// Open sets whether the overlay is open
func (o *SpectrumOverlay) Open(open bool) *SpectrumOverlay {
	o.open = open
	return o
}

// Delayed sets whether the overlay should be delayed
func (o *SpectrumOverlay) Delayed(delayed bool) *SpectrumOverlay {
	o.delayed = delayed
	return o
}

// Offset sets the distance between the overlay and the trigger
func (o *SpectrumOverlay) Offset(offset interface{}) *SpectrumOverlay {
	o.offset = offset
	return o
}

// Placement sets the placement of the overlay relative to the trigger
func (o *SpectrumOverlay) Placement(placement OverlayPlacement) *SpectrumOverlay {
	o.placement = placement
	return o
}

// ReceivesFocus sets how focus should be handled
func (o *SpectrumOverlay) ReceivesFocus(receivesFocus OverlayReceivesFocus) *SpectrumOverlay {
	o.receivesFocus = receivesFocus
	return o
}

// Trigger sets the trigger element ID or ID@interaction string
func (o *SpectrumOverlay) Trigger(trigger string) *SpectrumOverlay {
	o.trigger = trigger
	return o
}

// TriggerElement sets the trigger element directly
func (o *SpectrumOverlay) TriggerElement(triggerElement app.UI) *SpectrumOverlay {
	o.triggerElement = triggerElement
	return o
}

// TriggerInteraction sets the interaction type for triggering the overlay
func (o *SpectrumOverlay) TriggerInteraction(triggerInteraction OverlayTriggerInteraction) *SpectrumOverlay {
	o.triggerInteraction = triggerInteraction
	return o
}

// Type sets the overlay type
func (o *SpectrumOverlay) Type(overlayType OverlayType) *SpectrumOverlay {
	o.overlayType = overlayType
	return o
}

// Child sets the child element to display in the overlay
func (o *SpectrumOverlay) Child(child app.UI) *SpectrumOverlay {
	o.child = child
	return o
}

// OnOpened sets the opened event handler
func (o *SpectrumOverlay) OnOpened(handler app.EventHandler) *SpectrumOverlay {
	o.onOpened = handler
	return o
}

// OnClosed sets the closed event handler
func (o *SpectrumOverlay) OnClosed(handler app.EventHandler) *SpectrumOverlay {
	o.onClosed = handler
	return o
}

// Convenience methods for setting overlay types

// Auto sets type to auto
func (o *SpectrumOverlay) Auto() *SpectrumOverlay {
	return o.Type(OverlayTypeAuto)
}

// Hint sets type to hint
func (o *SpectrumOverlay) Hint() *SpectrumOverlay {
	return o.Type(OverlayTypeHint)
}

// Manual sets type to manual
func (o *SpectrumOverlay) Manual() *SpectrumOverlay {
	return o.Type(OverlayTypeManual)
}

// Modal sets type to modal
func (o *SpectrumOverlay) Modal() *SpectrumOverlay {
	return o.Type(OverlayTypeModal)
}

// Page sets type to page
func (o *SpectrumOverlay) Page() *SpectrumOverlay {
	return o.Type(OverlayTypePage)
}

// Convenience methods for setting trigger interactions

// ClickTrigger sets trigger interaction to click
func (o *SpectrumOverlay) ClickTrigger() *SpectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionClick)
}

// HoverTrigger sets trigger interaction to hover
func (o *SpectrumOverlay) HoverTrigger() *SpectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionHover)
}

// LongpressTrigger sets trigger interaction to longpress
func (o *SpectrumOverlay) LongpressTrigger() *SpectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionLongpress)
}

// Render renders the overlay component
func (o *SpectrumOverlay) Render() app.UI {
	overlay := app.Elem("sp-overlay")

	// Add properties
	if o.open {
		overlay = overlay.Attr("open", true)
	}
	if o.delayed {
		overlay = overlay.Attr("delayed", true)
	}
	if o.offset != nil {
		overlay = overlay.Attr("offset", o.offset)
	}
	if o.placement != "" {
		overlay = overlay.Attr("placement", string(o.placement))
	}
	if o.receivesFocus != "" {
		overlay = overlay.Attr("receives-focus", string(o.receivesFocus))
	}
	if o.trigger != "" {
		overlay = overlay.Attr("trigger", o.trigger)
	}
	if o.overlayType != "" {
		overlay = overlay.Attr("type", string(o.overlayType))
	}

	// Add event handlers
	if o.onOpened != nil {
		overlay = overlay.On("sp-opened", o.onOpened)
	}
	if o.onClosed != nil {
		overlay = overlay.On("sp-closed", o.onClosed)
	}

	// Add the child element if provided
	if o.child != nil {
		overlay = overlay.Body(o.child)
	}

	return overlay
}
