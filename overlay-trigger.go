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

// spectrumOverlayTrigger represents an overlay-trigger component
type spectrumOverlayTrigger struct {
	app.Compo

	// Properties
	PropDisabled      bool
	PropOffset        int
	PropOpen          string // Can be one of the OverlayContentTypes
	PropPlacement     OverlayPlacement
	PropReceivesFocus string // 'true', 'false', or 'auto'
	PropTriggeredBy   string // Space-separated list of trigger types
	PropTriggerType   OverlayTriggerType

	// Slots
	PropTrigger              app.UI
	PropClickContent         app.UI
	PropHoverContent         app.UI
	PropLongpressContent     app.UI
	PropLongpressDescription app.UI

	// Event handlers
	PropOnOpened app.EventHandler
	PropOnClosed app.EventHandler
}

// OverlayTrigger creates a new overlay trigger component
func OverlayTrigger() *spectrumOverlayTrigger {
	return &spectrumOverlayTrigger{
		PropOffset:        6,      // Default offset is 6
		PropReceivesFocus: "auto", // Default receives-focus is "auto"
	}
}

// Disabled sets whether the overlay trigger is disabled
func (ot *spectrumOverlayTrigger) Disabled(disabled bool) *spectrumOverlayTrigger {
	ot.PropDisabled = disabled
	return ot
}

// Offset sets the distance between the overlay and the trigger
func (ot *spectrumOverlayTrigger) Offset(offset int) *spectrumOverlayTrigger {
	ot.PropOffset = offset
	return ot
}

// Open sets which overlay content type should be open
func (ot *spectrumOverlayTrigger) Open(open string) *spectrumOverlayTrigger {
	ot.PropOpen = open
	return ot
}

// Placement sets the placement of the overlay relative to the trigger
func (ot *spectrumOverlayTrigger) Placement(placement OverlayPlacement) *spectrumOverlayTrigger {
	ot.PropPlacement = placement
	return ot
}

// ReceivesFocus sets how focus should be handled
func (ot *spectrumOverlayTrigger) ReceivesFocus(receivesFocus string) *spectrumOverlayTrigger {
	ot.PropReceivesFocus = receivesFocus
	return ot
}

// TriggeredBy sets which overlay content types should be used
func (ot *spectrumOverlayTrigger) TriggeredBy(triggeredBy string) *spectrumOverlayTrigger {
	ot.PropTriggeredBy = triggeredBy
	return ot
}

// Type sets the overlay trigger interaction model
func (ot *spectrumOverlayTrigger) Type(triggerType OverlayTriggerType) *spectrumOverlayTrigger {
	ot.PropTriggerType = triggerType
	return ot
}

// Trigger sets the trigger element
func (ot *spectrumOverlayTrigger) Trigger(trigger app.UI) *spectrumOverlayTrigger {
	ot.PropTrigger = trigger
	return ot
}

// ClickContent sets the content displayed on click
func (ot *spectrumOverlayTrigger) ClickContent(content app.UI) *spectrumOverlayTrigger {
	ot.PropClickContent = content
	return ot
}

// HoverContent sets the content displayed on hover
func (ot *spectrumOverlayTrigger) HoverContent(content app.UI) *spectrumOverlayTrigger {
	ot.PropHoverContent = content
	return ot
}

// LongpressContent sets the content displayed on longpress
func (ot *spectrumOverlayTrigger) LongpressContent(content app.UI) *spectrumOverlayTrigger {
	ot.PropLongpressContent = content
	return ot
}

// LongpressDescription sets the description for longpress content
func (ot *spectrumOverlayTrigger) LongpressDescription(description app.UI) *spectrumOverlayTrigger {
	ot.PropLongpressDescription = description
	return ot
}

// OnOpened sets the opened event handler
func (ot *spectrumOverlayTrigger) OnOpened(handler app.EventHandler) *spectrumOverlayTrigger {
	ot.PropOnOpened = handler
	return ot
}

// OnClosed sets the closed event handler
func (ot *spectrumOverlayTrigger) OnClosed(handler app.EventHandler) *spectrumOverlayTrigger {
	ot.PropOnClosed = handler
	return ot
}

// Convenience methods for setting overlay trigger types

// Inline sets type to inline
func (ot *spectrumOverlayTrigger) Inline() *spectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeInline)
}

// Replace sets type to replace
func (ot *spectrumOverlayTrigger) Replace() *spectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeReplace)
}

// Modal sets type to modal
func (ot *spectrumOverlayTrigger) Modal() *spectrumOverlayTrigger {
	return ot.Type(OverlayTriggerTypeModal)
}

// Convenience methods for setting triggered-by attribute

// TriggeredByClickOnly sets triggered-by to "click"
func (ot *spectrumOverlayTrigger) TriggeredByClickOnly() *spectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick)
}

// TriggeredByHoverOnly sets triggered-by to "hover"
func (ot *spectrumOverlayTrigger) TriggeredByHoverOnly() *spectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByHover)
}

// TriggeredByLongpressOnly sets triggered-by to "longpress"
func (ot *spectrumOverlayTrigger) TriggeredByLongpressOnly() *spectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByLongpress)
}

// TriggeredByClickAndHover sets triggered-by to "click hover"
func (ot *spectrumOverlayTrigger) TriggeredByClickAndHover() *spectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick + " " + TriggeredByHover)
}

// TriggeredByAll sets triggered-by to include all interaction types
func (ot *spectrumOverlayTrigger) TriggeredByAll() *spectrumOverlayTrigger {
	return ot.TriggeredBy(TriggeredByClick + " " + TriggeredByHover + " " + TriggeredByLongpress)
}

// Render renders the overlay trigger component
func (ot *spectrumOverlayTrigger) Render() app.UI {
	overlayTrigger := app.Elem("overlay-trigger")

	// Add properties
	if ot.PropDisabled {
		overlayTrigger = overlayTrigger.Attr("disabled", true)
	}
	if ot.PropOffset != 0 {
		overlayTrigger = overlayTrigger.Attr("offset", ot.PropOffset)
	}
	if ot.PropOpen != "" {
		overlayTrigger = overlayTrigger.Attr("open", ot.PropOpen)
	}
	if ot.PropPlacement != "" {
		overlayTrigger = overlayTrigger.Attr("placement", string(ot.PropPlacement))
	}
	if ot.PropReceivesFocus != "" {
		overlayTrigger = overlayTrigger.Attr("receives-focus", ot.PropReceivesFocus)
	}
	if ot.PropTriggeredBy != "" {
		overlayTrigger = overlayTrigger.Attr("triggered-by", ot.PropTriggeredBy)
	}
	if ot.PropTriggerType != "" {
		overlayTrigger = overlayTrigger.Attr("type", string(ot.PropTriggerType))
	}

	// Add event handlers
	if ot.PropOnOpened != nil {
		overlayTrigger = overlayTrigger.On("sp-opened", ot.PropOnOpened)
	}
	if ot.PropOnClosed != nil {
		overlayTrigger = overlayTrigger.On("sp-closed", ot.PropOnClosed)
	}

	// Collect slot elements
	elements := []app.UI{}

	// Add trigger slot if provided
	if ot.PropTrigger != nil {
		triggerElem := ot.PropTrigger
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
	if ot.PropClickContent != nil {
		clickContentElem := ot.PropClickContent
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
	if ot.PropHoverContent != nil {
		hoverContentElem := ot.PropHoverContent
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
	if ot.PropLongpressContent != nil {
		longpressContentElem := ot.PropLongpressContent
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
	if ot.PropLongpressDescription != nil {
		longpressDescElem := ot.PropLongpressDescription
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
