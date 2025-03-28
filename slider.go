package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SliderSize represents the visual size of a slider
type SliderSize string

// Slider sizes
const (
	SliderSizeS  SliderSize = "s"
	SliderSizeM  SliderSize = "m"
	SliderSizeL  SliderSize = "l"
	SliderSizeXL SliderSize = "xl"
)

// SliderVariant represents the visual variant of a slider
type SliderVariant string

// Slider variants
const (
	SliderVariantDefault SliderVariant = ""
	SliderVariantFilled  SliderVariant = "filled"
	SliderVariantTick    SliderVariant = "tick"
	SliderVariantRamp    SliderVariant = "ramp"
	SliderVariantRange   SliderVariant = "range"
)

// SliderLabelVisibility represents the visibility options for slider labels
type SliderLabelVisibility string

// Slider label visibility options
const (
	SliderLabelVisibilityText  SliderLabelVisibility = "text"
	SliderLabelVisibilityValue SliderLabelVisibility = "value"
	SliderLabelVisibilityNone  SliderLabelVisibility = "none"
)

// SpectrumSlider represents an sp-slider component
type SpectrumSlider struct {
	app.Compo

	// Properties
	size            SliderSize
	variant         SliderVariant
	label           string
	labelVisibility SliderLabelVisibility
	value           float64
	min             float64
	max             float64
	step            *float64
	defaultValue    *float64
	disabled        bool
	dragging        bool
	highlight       bool
	editable        bool
	hideStepper     bool
	fillStart       interface{} // Can be float64 or bool
	tickStep        int
	tickLabels      bool
	indeterminate   bool
	quiet           bool
	formatOptions   map[string]interface{}

	// Child handles
	handles []app.UI

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
}

// Slider creates a new slider component
func Slider() *SpectrumSlider {
	return &SpectrumSlider{
		size:          SliderSizeM, // Default size is medium
		min:           0,           // Default min is 0
		max:           100,         // Default max is 100
		formatOptions: make(map[string]interface{}),
	}
}

// Size sets the visual size of the slider
func (s *SpectrumSlider) Size(size SliderSize) *SpectrumSlider {
	s.size = size
	return s
}

// Variant sets the visual variant of the slider
func (s *SpectrumSlider) Variant(variant SliderVariant) *SpectrumSlider {
	s.variant = variant
	return s
}

// Label sets the accessibility label
func (s *SpectrumSlider) Label(label string) *SpectrumSlider {
	s.label = label
	return s
}

// LabelVisibility sets which labels should be visible
func (s *SpectrumSlider) LabelVisibility(visibility SliderLabelVisibility) *SpectrumSlider {
	s.labelVisibility = visibility
	return s
}

// Value sets the slider value (for single handle sliders)
func (s *SpectrumSlider) Value(value float64) *SpectrumSlider {
	s.value = value
	return s
}

// Min sets the minimum allowed value
func (s *SpectrumSlider) Min(min float64) *SpectrumSlider {
	s.min = min
	return s
}

// Max sets the maximum allowed value
func (s *SpectrumSlider) Max(max float64) *SpectrumSlider {
	s.max = max
	return s
}

// Step sets the step value for incrementing/decrementing
func (s *SpectrumSlider) Step(step float64) *SpectrumSlider {
	s.step = &step
	return s
}

// DefaultValue sets the default value to reset to on double-click or escape key
func (s *SpectrumSlider) DefaultValue(value float64) *SpectrumSlider {
	s.defaultValue = &value
	return s
}

// Disabled sets the disabled state
func (s *SpectrumSlider) Disabled(disabled bool) *SpectrumSlider {
	s.disabled = disabled
	return s
}

// Dragging sets the dragging state
func (s *SpectrumSlider) Dragging(dragging bool) *SpectrumSlider {
	s.dragging = dragging
	return s
}

// Highlight sets the highlight state
func (s *SpectrumSlider) Highlight(highlight bool) *SpectrumSlider {
	s.highlight = highlight
	return s
}

// Editable sets whether the slider should display a number field
func (s *SpectrumSlider) Editable(editable bool) *SpectrumSlider {
	s.editable = editable
	return s
}

// HideStepper sets whether to hide the stepper UI in the number field
func (s *SpectrumSlider) HideStepper(hide bool) *SpectrumSlider {
	s.hideStepper = hide
	return s
}

// FillStart sets a custom fill start value when variant is "filled"
func (s *SpectrumSlider) FillStart(start float64) *SpectrumSlider {
	s.fillStart = start
	return s
}

// EnableFillStart enables fill start with the default value
func (s *SpectrumSlider) EnableFillStart() *SpectrumSlider {
	s.fillStart = true
	return s
}

// TickStep sets the step size for tick marks when variant is "tick"
func (s *SpectrumSlider) TickStep(step int) *SpectrumSlider {
	s.tickStep = step
	return s
}

// TickLabels sets whether to show labels with tick marks
func (s *SpectrumSlider) TickLabels(show bool) *SpectrumSlider {
	s.tickLabels = show
	return s
}

// Indeterminate sets the indeterminate state
func (s *SpectrumSlider) Indeterminate(indeterminate bool) *SpectrumSlider {
	s.indeterminate = indeterminate
	return s
}

// Quiet sets whether the slider uses quiet styling
func (s *SpectrumSlider) Quiet(quiet bool) *SpectrumSlider {
	s.quiet = quiet
	return s
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (s *SpectrumSlider) FormatOptions(options map[string]interface{}) *SpectrumSlider {
	s.formatOptions = options
	return s
}

// Handles sets the slider handles when using multiple handles
func (s *SpectrumSlider) Handles(handles ...*SpectrumSliderHandle) *SpectrumSlider {
	s.handles = make([]app.UI, len(handles))
	for i, handle := range handles {
		s.handles[i] = handle
	}
	return s
}

// OnInput sets the input event handler
func (s *SpectrumSlider) OnInput(handler app.EventHandler) *SpectrumSlider {
	s.onInput = handler
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSlider) OnChange(handler app.EventHandler) *SpectrumSlider {
	s.onChange = handler
	return s
}

// Convenience methods for setting variants

// Filled sets the variant to "filled"
func (s *SpectrumSlider) Filled() *SpectrumSlider {
	return s.Variant(SliderVariantFilled)
}

// Tick sets the variant to "tick"
func (s *SpectrumSlider) Tick() *SpectrumSlider {
	return s.Variant(SliderVariantTick)
}

// Ramp sets the variant to "ramp"
func (s *SpectrumSlider) Ramp() *SpectrumSlider {
	return s.Variant(SliderVariantRamp)
}

// Range sets the variant to "range"
func (s *SpectrumSlider) Range() *SpectrumSlider {
	return s.Variant(SliderVariantRange)
}

// Render renders the slider component
func (s *SpectrumSlider) Render() app.UI {
	slider := app.Elem("sp-slider").
		Attr("size", string(s.size)).
		Attr("variant", string(s.variant)).
		Attr("label", s.label).
		Attr("value", s.value).
		Attr("min", s.min).
		Attr("max", s.max).
		Attr("disabled", s.disabled).
		Attr("dragging", s.dragging).
		Attr("highlight", s.highlight).
		Attr("editable", s.editable).
		Attr("hide-stepper", s.hideStepper).
		Attr("tick-step", s.tickStep).
		Attr("tick-labels", s.tickLabels).
		Attr("indeterminate", s.indeterminate).
		Attr("quiet", s.quiet)

	// Add the label visibility if specified
	if s.labelVisibility != "" {
		slider = slider.Attr("label-visibility", string(s.labelVisibility))
	}

	// Add step if set
	if s.step != nil {
		slider = slider.Attr("step", *s.step)
	}

	// Add default value if set
	if s.defaultValue != nil {
		slider = slider.Attr("default-value", *s.defaultValue)
	}

	// Handle fillStart which can be a number or boolean
	if s.fillStart != nil {
		switch fillStart := s.fillStart.(type) {
		case float64:
			slider = slider.Attr("fill-start", fillStart)
		case bool:
			if fillStart {
				slider = slider.Attr("fill-start", "")
			}
		}
	}

	// Add format options if provided
	if len(s.formatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(s.formatOptions)
		slider = slider.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if s.onInput != nil {
		slider = slider.On("input", s.onInput)
	}
	if s.onChange != nil {
		slider = slider.On("change", s.onChange)
	}

	// Add slider handles if provided
	if len(s.handles) > 0 {
		slider = slider.Body(s.handles...)
	}

	return slider
}
