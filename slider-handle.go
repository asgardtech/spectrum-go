package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SpectrumSliderHandle represents an sp-slider-handle component
type SpectrumSliderHandle struct {
	app.Compo

	// Properties
	label         string
	value         float64
	name          string
	min           interface{} // Can be float64 or string ("previous")
	max           interface{} // Can be float64 or string ("next")
	step          *float64
	defaultValue  *float64
	disabled      bool
	dragging      bool
	highlight     bool
	tabIndex      int
	formatOptions map[string]interface{}

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
}

// SliderHandle creates a new slider handle component
func SliderHandle() *SpectrumSliderHandle {
	return &SpectrumSliderHandle{
		formatOptions: make(map[string]interface{}),
	}
}

// Label sets the accessibility label
func (s *SpectrumSliderHandle) Label(label string) *SpectrumSliderHandle {
	s.label = label
	return s
}

// Value sets the handle value
func (s *SpectrumSliderHandle) Value(value float64) *SpectrumSliderHandle {
	s.value = value
	return s
}

// Name sets the form control name
func (s *SpectrumSliderHandle) Name(name string) *SpectrumSliderHandle {
	s.name = name
	return s
}

// Min sets the minimum allowed value or "previous" to use the previous handle's value
func (s *SpectrumSliderHandle) Min(min interface{}) *SpectrumSliderHandle {
	s.min = min
	return s
}

// MinValue sets a numeric minimum value
func (s *SpectrumSliderHandle) MinValue(min float64) *SpectrumSliderHandle {
	s.min = min
	return s
}

// MinPrevious sets the minimum to the previous handle's value
func (s *SpectrumSliderHandle) MinPrevious() *SpectrumSliderHandle {
	s.min = "previous"
	return s
}

// Max sets the maximum allowed value or "next" to use the next handle's value
func (s *SpectrumSliderHandle) Max(max interface{}) *SpectrumSliderHandle {
	s.max = max
	return s
}

// MaxValue sets a numeric maximum value
func (s *SpectrumSliderHandle) MaxValue(max float64) *SpectrumSliderHandle {
	s.max = max
	return s
}

// MaxNext sets the maximum to the next handle's value
func (s *SpectrumSliderHandle) MaxNext() *SpectrumSliderHandle {
	s.max = "next"
	return s
}

// Step sets the step value for incrementing/decrementing
func (s *SpectrumSliderHandle) Step(step float64) *SpectrumSliderHandle {
	s.step = &step
	return s
}

// DefaultValue sets the default value to reset to on double-click or escape key
func (s *SpectrumSliderHandle) DefaultValue(value float64) *SpectrumSliderHandle {
	s.defaultValue = &value
	return s
}

// Disabled sets the disabled state
func (s *SpectrumSliderHandle) Disabled(disabled bool) *SpectrumSliderHandle {
	s.disabled = disabled
	return s
}

// Dragging sets the dragging state
func (s *SpectrumSliderHandle) Dragging(dragging bool) *SpectrumSliderHandle {
	s.dragging = dragging
	return s
}

// Highlight sets the highlight state
func (s *SpectrumSliderHandle) Highlight(highlight bool) *SpectrumSliderHandle {
	s.highlight = highlight
	return s
}

// TabIndex sets the tab index
func (s *SpectrumSliderHandle) TabIndex(tabIndex int) *SpectrumSliderHandle {
	s.tabIndex = tabIndex
	return s
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (s *SpectrumSliderHandle) FormatOptions(options map[string]interface{}) *SpectrumSliderHandle {
	s.formatOptions = options
	return s
}

// OnInput sets the input event handler
func (s *SpectrumSliderHandle) OnInput(handler app.EventHandler) *SpectrumSliderHandle {
	s.onInput = handler
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSliderHandle) OnChange(handler app.EventHandler) *SpectrumSliderHandle {
	s.onChange = handler
	return s
}

// Slot sets the slot to "handle" by default since this is required for slider handles
func (s *SpectrumSliderHandle) Slot(slot string) app.UI {
	return app.Elem("sp-slider-handle").
		Attr("slot", slot).
		Attr("label", s.label).
		Attr("value", s.value).
		Attr("name", s.name).
		Attr("disabled", s.disabled).
		Attr("dragging", s.dragging).
		Attr("highlight", s.highlight).
		Attr("tabindex", s.tabIndex)
}

// Render renders the slider handle component
func (s *SpectrumSliderHandle) Render() app.UI {
	sliderHandle := app.Elem("sp-slider-handle").
		Attr("slot", "handle"). // Default slot for slider handles
		Attr("label", s.label).
		Attr("value", s.value).
		Attr("name", s.name).
		Attr("disabled", s.disabled).
		Attr("dragging", s.dragging).
		Attr("highlight", s.highlight).
		Attr("tabindex", s.tabIndex)

	// Handle min attribute which can be a number or "previous"
	if s.min != nil {
		switch min := s.min.(type) {
		case float64:
			sliderHandle = sliderHandle.Attr("min", min)
		case string:
			sliderHandle = sliderHandle.Attr("min", min)
		}
	}

	// Handle max attribute which can be a number or "next"
	if s.max != nil {
		switch max := s.max.(type) {
		case float64:
			sliderHandle = sliderHandle.Attr("max", max)
		case string:
			sliderHandle = sliderHandle.Attr("max", max)
		}
	}

	// Add step if set
	if s.step != nil {
		sliderHandle = sliderHandle.Attr("step", *s.step)
	}

	// Add default value if set
	if s.defaultValue != nil {
		sliderHandle = sliderHandle.Attr("default-value", *s.defaultValue)
	}

	// Add format options if provided
	if len(s.formatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(s.formatOptions)
		sliderHandle = sliderHandle.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if s.onInput != nil {
		sliderHandle = sliderHandle.On("input", s.onInput)
	}
	if s.onChange != nil {
		sliderHandle = sliderHandle.On("change", s.onChange)
	}

	return sliderHandle
}
