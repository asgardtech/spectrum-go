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

// spectrumSlider represents an sp-slider component
type spectrumSlider struct {
	app.Compo

	// Properties
	PropSize            SliderSize
	PropVariant         SliderVariant
	PropLabel           string
	PropLabelVisibility SliderLabelVisibility
	PropValue           float64
	PropMin             float64
	PropMax             float64
	PropStep            *float64
	PropDefaultValue    *float64
	PropDisabled        bool
	PropDragging        bool
	PropHighlight       bool
	PropEditable        bool
	PropHideStepper     bool
	PropFillStart       interface{} // Can be float64 or bool
	PropTickStep        int
	PropTickLabels      bool
	PropIndeterminate   bool
	PropQuiet           bool
	PropFormatOptions   map[string]interface{}

	// Child handles
	PropHandles []app.UI

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
}

// Slider creates a new slider component
func Slider() *spectrumSlider {
	return &spectrumSlider{
		PropSize:          SliderSizeM, // Default size is medium
		PropMin:           0,           // Default min is 0
		PropMax:           100,         // Default max is 100
		PropFormatOptions: make(map[string]interface{}),
	}
}

// Size sets the visual size of the slider
func (s *spectrumSlider) Size(size SliderSize) *spectrumSlider {
	s.PropSize = size
	return s
}

// Variant sets the visual variant of the slider
func (s *spectrumSlider) Variant(variant SliderVariant) *spectrumSlider {
	s.PropVariant = variant
	return s
}

// Label sets the accessibility label
func (s *spectrumSlider) Label(label string) *spectrumSlider {
	s.PropLabel = label
	return s
}

// LabelVisibility sets which labels should be visible
func (s *spectrumSlider) LabelVisibility(visibility SliderLabelVisibility) *spectrumSlider {
	s.PropLabelVisibility = visibility
	return s
}

// Value sets the slider value (for single handle sliders)
func (s *spectrumSlider) Value(value float64) *spectrumSlider {
	s.PropValue = value
	return s
}

// Min sets the minimum allowed value
func (s *spectrumSlider) Min(min float64) *spectrumSlider {
	s.PropMin = min
	return s
}

// Max sets the maximum allowed value
func (s *spectrumSlider) Max(max float64) *spectrumSlider {
	s.PropMax = max
	return s
}

// Step sets the step value for incrementing/decrementing
func (s *spectrumSlider) Step(step float64) *spectrumSlider {
	s.PropStep = &step
	return s
}

// DefaultValue sets the default value to reset to on double-click or escape key
func (s *spectrumSlider) DefaultValue(value float64) *spectrumSlider {
	s.PropDefaultValue = &value
	return s
}

// Disabled sets the disabled state
func (s *spectrumSlider) Disabled(disabled bool) *spectrumSlider {
	s.PropDisabled = disabled
	return s
}

// Dragging sets the dragging state
func (s *spectrumSlider) Dragging(dragging bool) *spectrumSlider {
	s.PropDragging = dragging
	return s
}

// Highlight sets the highlight state
func (s *spectrumSlider) Highlight(highlight bool) *spectrumSlider {
	s.PropHighlight = highlight
	return s
}

// Editable sets whether the slider should display a number field
func (s *spectrumSlider) Editable(editable bool) *spectrumSlider {
	s.PropEditable = editable
	return s
}

// HideStepper sets whether to hide the stepper UI in the number field
func (s *spectrumSlider) HideStepper(hide bool) *spectrumSlider {
	s.PropHideStepper = hide
	return s
}

// FillStart sets a custom fill start value when variant is "filled"
func (s *spectrumSlider) FillStart(start float64) *spectrumSlider {
	s.PropFillStart = start
	return s
}

// EnableFillStart enables fill start with the default value
func (s *spectrumSlider) EnableFillStart() *spectrumSlider {
	s.PropFillStart = true
	return s
}

// TickStep sets the step size for tick marks when variant is "tick"
func (s *spectrumSlider) TickStep(step int) *spectrumSlider {
	s.PropTickStep = step
	return s
}

// TickLabels sets whether to show labels with tick marks
func (s *spectrumSlider) TickLabels(show bool) *spectrumSlider {
	s.PropTickLabels = show
	return s
}

// Indeterminate sets the indeterminate state
func (s *spectrumSlider) Indeterminate(indeterminate bool) *spectrumSlider {
	s.PropIndeterminate = indeterminate
	return s
}

// Quiet sets whether the slider uses quiet styling
func (s *spectrumSlider) Quiet(quiet bool) *spectrumSlider {
	s.PropQuiet = quiet
	return s
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (s *spectrumSlider) FormatOptions(options map[string]interface{}) *spectrumSlider {
	s.PropFormatOptions = options
	return s
}

// Handles sets the slider handles when using multiple handles
func (s *spectrumSlider) Handles(handles ...app.UI) *spectrumSlider {
	s.PropHandles = make([]app.UI, len(handles))
	for i, handle := range handles {
		s.PropHandles[i] = handle
	}
	return s
}

// OnInput sets the input event handler
func (s *spectrumSlider) OnInput(handler app.EventHandler) *spectrumSlider {
	s.PropOnInput = handler
	return s
}

// OnChange sets the change event handler
func (s *spectrumSlider) OnChange(handler app.EventHandler) *spectrumSlider {
	s.PropOnChange = handler
	return s
}

// Convenience methods for setting variants

// Filled sets the variant to "filled"
func (s *spectrumSlider) Filled() *spectrumSlider {
	return s.Variant(SliderVariantFilled)
}

// Tick sets the variant to "tick"
func (s *spectrumSlider) Tick() *spectrumSlider {
	return s.Variant(SliderVariantTick)
}

// Ramp sets the variant to "ramp"
func (s *spectrumSlider) Ramp() *spectrumSlider {
	return s.Variant(SliderVariantRamp)
}

// Range sets the variant to "range"
func (s *spectrumSlider) Range() *spectrumSlider {
	return s.Variant(SliderVariantRange)
}

// Render renders the slider component
func (s *spectrumSlider) Render() app.UI {
	slider := app.Elem("sp-slider").
		Attr("size", string(s.PropSize)).
		Attr("variant", string(s.PropVariant)).
		Attr("label", s.PropLabel).
		Attr("value", s.PropValue).
		Attr("min", s.PropMin).
		Attr("max", s.PropMax).
		Attr("disabled", s.PropDisabled).
		Attr("dragging", s.PropDragging).
		Attr("highlight", s.PropHighlight).
		Attr("editable", s.PropEditable).
		Attr("hide-stepper", s.PropHideStepper).
		Attr("tick-step", s.PropTickStep).
		Attr("tick-labels", s.PropTickLabels).
		Attr("indeterminate", s.PropIndeterminate).
		Attr("quiet", s.PropQuiet)

	// Add the label visibility if specified
	if s.PropLabelVisibility != "" {
		slider = slider.Attr("label-visibility", string(s.PropLabelVisibility))
	}

	// Add step if set
	if s.PropStep != nil {
		slider = slider.Attr("step", *s.PropStep)
	}

	// Add default value if set
	if s.PropDefaultValue != nil {
		slider = slider.Attr("default-value", *s.PropDefaultValue)
	}

	// Handle fillStart which can be a number or boolean
	if s.PropFillStart != nil {
		switch fillStart := s.PropFillStart.(type) {
		case float64:
			slider = slider.Attr("fill-start", fillStart)
		case bool:
			if fillStart {
				slider = slider.Attr("fill-start", "")
			}
		}
	}

	// Add format options if provided
	if len(s.PropFormatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(s.PropFormatOptions)
		slider = slider.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if s.PropOnInput != nil {
		slider = slider.On("input", s.PropOnInput)
	}
	if s.PropOnChange != nil {
		slider = slider.On("change", s.PropOnChange)
	}

	// Add slider handles if provided
	if len(s.PropHandles) > 0 {
		slider = slider.Body(s.PropHandles...)
	}

	return slider
}
