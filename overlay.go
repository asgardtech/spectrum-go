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

// spectrumOverlay represents an sp-overlay component
type spectrumOverlay struct {
	app.Compo

	// Properties
	PropOpen               bool
	PropDelayed            bool
	PropOffset             interface{} // Can be a number or [number, number]
	PropPlacement          OverlayPlacement
	PropReceivesFocus      OverlayReceivesFocus
	PropTrigger            string
	PropTriggerElement     app.UI
	PropTriggerInteraction OverlayTriggerInteraction
	PropOverlayType        OverlayType

	// Child element to display in the overlay
	PropChild app.UI

	// Event handlers
	PropOnOpened app.EventHandler
	PropOnClosed app.EventHandler
}

// Overlay creates a new overlay component
func Overlay() *spectrumOverlay {
	return &spectrumOverlay{
		PropReceivesFocus: OverlayReceivesFocusAuto, // Default receives-focus is "auto"
	}
}

// Open sets whether the overlay is open
func (o *spectrumOverlay) Open(open bool) *spectrumOverlay {
	o.PropOpen = open
	return o
}

// Delayed sets whether the overlay should be delayed
func (o *spectrumOverlay) Delayed(delayed bool) *spectrumOverlay {
	o.PropDelayed = delayed
	return o
}

// Offset sets the distance between the overlay and the trigger
func (o *spectrumOverlay) Offset(offset interface{}) *spectrumOverlay {
	o.PropOffset = offset
	return o
}

// Placement sets the placement of the overlay relative to the trigger
func (o *spectrumOverlay) Placement(placement OverlayPlacement) *spectrumOverlay {
	o.PropPlacement = placement
	return o
}

// ReceivesFocus sets how focus should be handled
func (o *spectrumOverlay) ReceivesFocus(receivesFocus OverlayReceivesFocus) *spectrumOverlay {
	o.PropReceivesFocus = receivesFocus
	return o
}

// Trigger sets the trigger element ID or ID@interaction string
func (o *spectrumOverlay) Trigger(trigger string) *spectrumOverlay {
	o.PropTrigger = trigger
	return o
}

// TriggerElement sets the trigger element directly
func (o *spectrumOverlay) TriggerElement(triggerElement app.UI) *spectrumOverlay {
	o.PropTriggerElement = triggerElement
	return o
}

// TriggerInteraction sets the interaction type for triggering the overlay
func (o *spectrumOverlay) TriggerInteraction(triggerInteraction OverlayTriggerInteraction) *spectrumOverlay {
	o.PropTriggerInteraction = triggerInteraction
	return o
}

// Type sets the overlay type
func (o *spectrumOverlay) Type(overlayType OverlayType) *spectrumOverlay {
	o.PropOverlayType = overlayType
	return o
}

// Child sets the child element to display in the overlay
func (o *spectrumOverlay) Child(child app.UI) *spectrumOverlay {
	o.PropChild = child
	return o
}

// OnOpened sets the opened event handler
func (o *spectrumOverlay) OnOpened(handler app.EventHandler) *spectrumOverlay {
	o.PropOnOpened = handler
	return o
}

// OnClosed sets the closed event handler
func (o *spectrumOverlay) OnClosed(handler app.EventHandler) *spectrumOverlay {
	o.PropOnClosed = handler
	return o
}

// Convenience methods for setting overlay types

// Auto sets type to auto
func (o *spectrumOverlay) Auto() *spectrumOverlay {
	return o.Type(OverlayTypeAuto)
}

// Hint sets type to hint
func (o *spectrumOverlay) Hint() *spectrumOverlay {
	return o.Type(OverlayTypeHint)
}

// Manual sets type to manual
func (o *spectrumOverlay) Manual() *spectrumOverlay {
	return o.Type(OverlayTypeManual)
}

// Modal sets type to modal
func (o *spectrumOverlay) Modal() *spectrumOverlay {
	return o.Type(OverlayTypeModal)
}

// Page sets type to page
func (o *spectrumOverlay) Page() *spectrumOverlay {
	return o.Type(OverlayTypePage)
}

// Convenience methods for setting trigger interactions

// ClickTrigger sets trigger interaction to click
func (o *spectrumOverlay) ClickTrigger() *spectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionClick)
}

// HoverTrigger sets trigger interaction to hover
func (o *spectrumOverlay) HoverTrigger() *spectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionHover)
}

// LongpressTrigger sets trigger interaction to longpress
func (o *spectrumOverlay) LongpressTrigger() *spectrumOverlay {
	return o.TriggerInteraction(OverlayTriggerInteractionLongpress)
}

// Render renders the overlay component
func (o *spectrumOverlay) Render() app.UI {
	overlay := app.Elem("sp-overlay")

	// Add properties
	if o.PropOpen {
		overlay = overlay.Attr("open", true)
	}
	if o.PropDelayed {
		overlay = overlay.Attr("delayed", true)
	}
	if o.PropOffset != nil {
		overlay = overlay.Attr("offset", o.PropOffset)
	}
	if o.PropPlacement != "" {
		overlay = overlay.Attr("placement", string(o.PropPlacement))
	}
	if o.PropReceivesFocus != "" {
		overlay = overlay.Attr("receives-focus", string(o.PropReceivesFocus))
	}
	if o.PropTrigger != "" {
		overlay = overlay.Attr("trigger", o.PropTrigger)
	}
	if o.PropOverlayType != "" {
		overlay = overlay.Attr("type", string(o.PropOverlayType))
	}

	// Add event handlers
	if o.PropOnOpened != nil {
		overlay = overlay.On("sp-opened", o.PropOnOpened)
	}
	if o.PropOnClosed != nil {
		overlay = overlay.On("sp-closed", o.PropOnClosed)
	}

	// Add the child element if provided
	if o.PropChild != nil {
		overlay = overlay.Body(o.PropChild)
	}

	return overlay
}
