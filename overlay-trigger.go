package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// OverlayTriggerType represents the interaction model of an overlay trigger
type OverlayTriggerType string

// Overlay trigger type values
const (
	OverlayTriggerTypeInline  OverlayTriggerType = "inline"
	OverlayTriggerTypeReplace OverlayTriggerType = "replace"
	OverlayTriggerTypeModal   OverlayTriggerType = "modal"
)

// TriggeredByType represents which content types the overlay trigger should use
type TriggeredByType string

// Triggered by values (can be space-separated combinations)
const (
	TriggeredByClick     = "click"
	TriggeredByHover     = "hover"
	TriggeredByLongpress = "longpress"
)

// SpectrumOverlayTrigger represents an overlay-trigger component
type SpectrumOverlayTrigger struct {
	app.Compo

	// Properties
	disabled      bool
	offset        int
	open          string // Can be one of the OverlayContentTypes
	placement     OverlayPlacement
	receivesFocus string // 'true', 'false', or 'auto'
	triggeredBy   string // Space-separated list of trigger types
	triggerType   OverlayTriggerType

	// Slots
	trigger              app.UI
	clickContent         app.UI
	hoverContent         app.UI
	longpressContent     app.UI
	longpressDescription app.UI

	// Event handlers
	onOpened app.EventHandler
	onClosed app.EventHandler
}

// OverlayTrigger creates a new overlay trigger component
func OverlayTrigger() *SpectrumOverlayTrigger {
	return &SpectrumOverlayTrigger{
		offset:        6,      // Default offset is 6
		receivesFocus: "auto", // Default receives-focus is "auto"
	}
}

// Disabled sets whether the overlay trigger is disabled
func (ot *SpectrumOverlayTrigger) Disabled(disabled bool) *SpectrumOverlayTrigger {
	ot.disabled = disabled
	return ot
}

// Offset sets the distance between the overlay and the trigger
func (ot *SpectrumOverlayTrigger) Offset(offset int) *SpectrumOverlayTrigger {
	ot.offset = offset
	return ot
}

// Open sets which overlay content type should be open
func (ot *SpectrumOverlayTrigger) Open(open string) *SpectrumOverlayTrigger {
	ot.open = open
	return ot
}

// Placement sets the placement of the overlay relative to the trigger
func (ot *SpectrumOverlayTrigger) Placement(placement OverlayPlacement) *SpectrumOverlayTrigger {
	ot.placement = placement
	return ot
}

// ReceivesFocus sets how focus should be handled
func (ot *SpectrumOverlayTrigger) ReceivesFocus(receivesFocus string) *SpectrumOverlayTrigger {
	ot.receivesFocus = receivesFocus
	return ot
}

// TriggeredBy sets which overlay content types should be used
func (ot *SpectrumOverlayTrigger) TriggeredBy(triggeredBy string) *SpectrumOverlayTrigger {
	ot.triggeredBy = triggeredBy
	return ot
}

// Type sets the overlay trigger interaction model
func (ot *SpectrumOverlayTrigger) Type(triggerType OverlayTriggerType) *SpectrumOverlayTrigger {
	ot.triggerType = triggerType
	return ot
}

// Trigger sets the trigger element
func (ot *SpectrumOverlayTrigger) Trigger(trigger app.UI) *SpectrumOverlayTrigger {
	ot.trigger = trigger
	return ot
}

// ClickContent sets the content displayed on click
func (ot *SpectrumOverlayTrigger) ClickContent(content app.UI) *SpectrumOverlayTrigger {
	ot.clickContent = content
	return ot
}

// HoverContent sets the content displayed on hover
func (ot *SpectrumOverlayTrigger) HoverContent(content app.UI) *SpectrumOverlayTrigger {
	ot.hoverContent = content
	return ot
}

// LongpressContent sets the content displayed on longpress
func (ot *SpectrumOverlayTrigger) LongpressContent(content app.UI) *SpectrumOverlayTrigger {
	ot.longpressContent = content
	return ot
}

// LongpressDescription sets the description for longpress content
func (ot *SpectrumOverlayTrigger) LongpressDescription(description app.UI) *SpectrumOverlayTrigger {
	ot.longpressDescription = description
	return ot
}

// OnOpened sets the opened event handler
func (ot *SpectrumOverlayTrigger) OnOpened(handler app.EventHandler) *SpectrumOverlayTrigger {
	ot.onOpened = handler
	return ot
}

// OnClosed sets the closed event handler
func (ot *SpectrumOverlayTrigger) OnClosed(handler app.EventHandler) *SpectrumOverlayTrigger {
	ot.onClosed = handler
	return ot
}

// Convenience methods for setting overlay trigger types

// Inline sets type to inline
func (ot *SpectrumOverlayTrigger) Inline() *SpectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeInline)
}

// Replace sets type to replace
func (ot *SpectrumOverlayTrigger) Replace() *SpectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeReplace)
}

// Modal sets type to modal
func (ot *SpectrumOverlayTrigger) Modal() *SpectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeModal)
}

// Convenience methods for setting triggered-by attribute

// TriggeredByClickOnly sets triggered-by to "click"
func (ot *SpectrumOverlayTrigger) TriggeredByClickOnly() *SpectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick)
}

// TriggeredByHoverOnly sets triggered-by to "hover"
func (ot *SpectrumOverlayTrigger) TriggeredByHoverOnly() *SpectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByHover)
}

// TriggeredByLongpressOnly sets triggered-by to "longpress"
func (ot *SpectrumOverlayTrigger) TriggeredByLongpressOnly() *SpectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByLongpress)
}

// TriggeredByClickAndHover sets triggered-by to "click hover"
func (ot *SpectrumOverlayTrigger) TriggeredByClickAndHover() *SpectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick + " " + TriggeredByHover)
}

// TriggeredByAll sets triggered-by to include all interaction types
func (ot *SpectrumOverlayTrigger) TriggeredByAll() *SpectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick + " " + TriggeredByHover + " " + TriggeredByLongpress)
}

// Render renders the overlay trigger component
func (ot *SpectrumOverlayTrigger) Render() app.UI {
	overlayTrigger := app.Elem("overlay-trigger")

	// Add properties
	if ot.disabled {
		overlayTrigger = overlayTrigger.Attr("disabled", true)
	}
	if ot.offset != 0 {
		overlayTrigger = overlayTrigger.Attr("offset", ot.offset)
	}
	if ot.open != "" {
		overlayTrigger = overlayTrigger.Attr("open", ot.open)
	}
	if ot.placement != "" {
		overlayTrigger = overlayTrigger.Attr("placement", string(ot.placement))
	}
	if ot.receivesFocus != "" {
		overlayTrigger = overlayTrigger.Attr("receives-focus", ot.receivesFocus)
	}
	if ot.triggeredBy != "" {
		overlayTrigger = overlayTrigger.Attr("triggered-by", ot.triggeredBy)
	}
	if ot.triggerType != "" {
		overlayTrigger = overlayTrigger.Attr("type", string(ot.triggerType))
	}

	// Add event handlers
	if ot.onOpened != nil {
		overlayTrigger = overlayTrigger.On("sp-opened", ot.onOpened)
	}
	if ot.onClosed != nil {
		overlayTrigger = overlayTrigger.On("sp-closed", ot.onClosed)
	}

	// Collect slot elements
	elements := []app.UI{}

	// Add trigger slot if provided
	if ot.trigger != nil {
		triggerElem := ot.trigger
		if triggerWithSlot, ok := triggerElem.(interface{ Slot(string) app.UI }); ok {
			triggerElem = triggerWithSlot.Slot("trigger")
		} else {
			triggerElem = app.Elem("div").
				Attr("slot", "trigger").
				Body(triggerElem)
		}
		elements = append(elements, triggerElem)
	}

	// Add click-content slot if provided
	if ot.clickContent != nil {
		clickContentElem := ot.clickContent
		if clickContentWithSlot, ok := clickContentElem.(interface{ Slot(string) app.UI }); ok {
			clickContentElem = clickContentWithSlot.Slot("click-content")
		} else {
			clickContentElem = app.Elem("div").
				Attr("slot", "click-content").
				Body(clickContentElem)
		}
		elements = append(elements, clickContentElem)
	}

	// Add hover-content slot if provided
	if ot.hoverContent != nil {
		hoverContentElem := ot.hoverContent
		if hoverContentWithSlot, ok := hoverContentElem.(interface{ Slot(string) app.UI }); ok {
			hoverContentElem = hoverContentWithSlot.Slot("hover-content")
		} else {
			hoverContentElem = app.Elem("div").
				Attr("slot", "hover-content").
				Body(hoverContentElem)
		}
		elements = append(elements, hoverContentElem)
	}

	// Add longpress-content slot if provided
	if ot.longpressContent != nil {
		longpressContentElem := ot.longpressContent
		if longpressContentWithSlot, ok := longpressContentElem.(interface{ Slot(string) app.UI }); ok {
			longpressContentElem = longpressContentWithSlot.Slot("longpress-content")
		} else {
			longpressContentElem = app.Elem("div").
				Attr("slot", "longpress-content").
				Body(longpressContentElem)
		}
		elements = append(elements, longpressContentElem)
	}

	// Add longpress-describedby-descriptor slot if provided
	if ot.longpressDescription != nil {
		longpressDescElem := ot.longpressDescription
		if longpressDescWithSlot, ok := longpressDescElem.(interface{ Slot(string) app.UI }); ok {
			longpressDescElem = longpressDescWithSlot.Slot("longpress-describedby-descriptor")
		} else {
			longpressDescElem = app.Elem("div").
				Attr("slot", "longpress-describedby-descriptor").
				Body(longpressDescElem)
		}
		elements = append(elements, longpressDescElem)
	}

	// Add all elements to the overlay trigger
	if len(elements) > 0 {
		overlayTrigger = overlayTrigger.Body(elements...)
	}

	return overlayTrigger
}
