package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// spectrumSliderHandle represents an sp-slider-handle component
type spectrumSliderHandle struct {
	app.Compo

	// Properties
	PropLabel         string
	PropValue         float64
	PropName          string
	PropMin           interface{} // Can be float64 or string ("previous")
	PropMax           interface{} // Can be float64 or string ("next")
	PropStep          *float64
	PropDefaultValue  *float64
	PropDisabled      bool
	PropDragging      bool
	PropHighlight     bool
	PropTabIndex      int
	PropFormatOptions map[string]interface{}

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
}

// SliderHandle creates a new slider handle component
func SliderHandle() *spectrumSliderHandle {
	return &spectrumSliderHandle{
		PropFormatOptions: make(map[string]interface{}),
	}
}

// Label sets the accessibility label
func (s *spectrumSliderHandle) Label(label string) *spectrumSliderHandle {
	s.PropLabel = label
	return s
}

// Value sets the handle value
func (s *spectrumSliderHandle) Value(value float64) *spectrumSliderHandle {
	s.PropValue = value
	return s
}

// Name sets the form control name
func (s *spectrumSliderHandle) Name(name string) *spectrumSliderHandle {
	s.PropName = name
	return s
}

// Min sets the minimum allowed value or "previous" to use the previous handle's value
func (s *spectrumSliderHandle) Min(min interface{}) *spectrumSliderHandle {
	s.PropMin = min
	return s
}

// MinValue sets a numeric minimum value
func (s *spectrumSliderHandle) MinValue(min float64) *spectrumSliderHandle {
	s.PropMin = min
	return s
}

// MinPrevious sets the minimum to the previous handle's value
func (s *spectrumSliderHandle) MinPrevious() *spectrumSliderHandle {
	s.PropMin = "previous"
	return s
}

// Max sets the maximum allowed value or "next" to use the next handle's value
func (s *spectrumSliderHandle) Max(max interface{}) *spectrumSliderHandle {
	s.PropMax = max
	return s
}

// MaxValue sets a numeric maximum value
func (s *spectrumSliderHandle) MaxValue(max float64) *spectrumSliderHandle {
	s.PropMax = max
	return s
}

// MaxNext sets the maximum to the next handle's value
func (s *spectrumSliderHandle) MaxNext() *spectrumSliderHandle {
	s.PropMax = "next"
	return s
}

// Step sets the step value for incrementing/decrementing
func (s *spectrumSliderHandle) Step(step float64) *spectrumSliderHandle {
	s.PropStep = &step
	return s
}

// DefaultValue sets the default value to reset to on double-click or escape key
func (s *spectrumSliderHandle) DefaultValue(value float64) *spectrumSliderHandle {
	s.PropDefaultValue = &value
	return s
}

// Disabled sets the disabled state
func (s *spectrumSliderHandle) Disabled(disabled bool) *spectrumSliderHandle {
	s.PropDisabled = disabled
	return s
}

// Dragging sets the dragging state
func (s *spectrumSliderHandle) Dragging(dragging bool) *spectrumSliderHandle {
	s.PropDragging = dragging
	return s
}

// Highlight sets the highlight state
func (s *spectrumSliderHandle) Highlight(highlight bool) *spectrumSliderHandle {
	s.PropHighlight = highlight
	return s
}

// TabIndex sets the tab index
func (s *spectrumSliderHandle) TabIndex(tabIndex int) *spectrumSliderHandle {
	s.PropTabIndex = tabIndex
	return s
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (s *spectrumSliderHandle) FormatOptions(options map[string]interface{}) *spectrumSliderHandle {
	s.PropFormatOptions = options
	return s
}

// OnInput sets the input event handler
func (s *spectrumSliderHandle) OnInput(handler app.EventHandler) *spectrumSliderHandle {
	s.PropOnInput = handler
	return s
}

// OnChange sets the change event handler
func (s *spectrumSliderHandle) OnChange(handler app.EventHandler) *spectrumSliderHandle {
	s.PropOnChange = handler
	return s
}

// Slot sets the slot to "handle" by default since this is required for slider handles
func (s *spectrumSliderHandle) Slot(slot string) app.UI {
	return app.Elem("sp-slider-handle").
		Attr("slot", slot).
		Attr("label", s.PropLabel).
		Attr("value", s.PropValue).
		Attr("name", s.PropName).
		Attr("disabled", s.PropDisabled).
		Attr("dragging", s.PropDragging).
		Attr("highlight", s.PropHighlight).
		Attr("tabindex", s.PropTabIndex)
}

// Render renders the slider handle component
func (s *spectrumSliderHandle) Render() app.UI {
	sliderHandle := app.Elem("sp-slider-handle").
		Attr("slot", "handle"). // Default slot for slider handles
		Attr("label", s.PropLabel).
		Attr("value", s.PropValue).
		Attr("name", s.PropName).
		Attr("disabled", s.PropDisabled).
		Attr("dragging", s.PropDragging).
		Attr("highlight", s.PropHighlight).
		Attr("tabindex", s.PropTabIndex)

	// Handle min attribute which can be a number or "previous"
	if s.PropMin != nil {
		switch min := s.PropMin.(type) {
		case float64:
			sliderHandle = sliderHandle.Attr("min", min)
		case string:
			sliderHandle = sliderHandle.Attr("min", min)
		}
	}

	// Handle max attribute which can be a number or "next"
	if s.PropMax != nil {
		switch max := s.PropMax.(type) {
		case float64:
			sliderHandle = sliderHandle.Attr("max", max)
		case string:
			sliderHandle = sliderHandle.Attr("max", max)
		}
	}

	// Add step if set
	if s.PropStep != nil {
		sliderHandle = sliderHandle.Attr("step", *s.PropStep)
	}

	// Add default value if set
	if s.PropDefaultValue != nil {
		sliderHandle = sliderHandle.Attr("default-value", *s.PropDefaultValue)
	}

	// Add format options if provided
	if len(s.PropFormatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(s.PropFormatOptions)
		sliderHandle = sliderHandle.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if s.PropOnInput != nil {
		sliderHandle = sliderHandle.On("input", s.PropOnInput)
	}
	if s.PropOnChange != nil {
		sliderHandle = sliderHandle.On("change", s.PropOnChange)
	}

	return sliderHandle
}
