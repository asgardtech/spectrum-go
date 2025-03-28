package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwatchSize represents the visual size of a swatch
type SwatchSize string

// Swatch sizes
const (
	SwatchSizeXS SwatchSize = "xs"
	SwatchSizeS  SwatchSize = "s"
	SwatchSizeM  SwatchSize = "m"
	SwatchSizeL  SwatchSize = "l"
)

// SwatchShape represents the shape of a swatch
type SwatchShape string

// Swatch shapes
const (
	SwatchShapeRectangle SwatchShape = "rectangle"
)

// SwatchRounding represents the corner rounding of a swatch
type SwatchRounding string

// Swatch rounding values
const (
	SwatchRoundingNone SwatchRounding = "none"
	SwatchRoundingFull SwatchRounding = "full"
)

// SwatchBorder represents the border style of a swatch
type SwatchBorder string

// Swatch border values
const (
	SwatchBorderNone  SwatchBorder = "none"
	SwatchBorderLight SwatchBorder = "light"
)

// SpectrumSwatch represents an sp-swatch component
type SpectrumSwatch struct {
	app.Compo

	// Properties
	size       SwatchSize
	shape      SwatchShape
	rounding   SwatchRounding
	border     SwatchBorder
	color      string
	disabled   bool
	mixedValue bool
	nothing    bool
	selected   bool
	label      string
	value      string
	tabIndex   int
	role       string

	// Event handlers
	onChange app.EventHandler
}

// Swatch creates a new swatch component
func Swatch() *SpectrumSwatch {
	return &SpectrumSwatch{
		size: SwatchSizeM, // Default size is medium
		role: "button",    // Default role is button
	}
}

// Size sets the visual size of the swatch
func (s *SpectrumSwatch) Size(size SwatchSize) *SpectrumSwatch {
	s.size = size
	return s
}

// Shape sets the shape of the swatch
func (s *SpectrumSwatch) Shape(shape SwatchShape) *SpectrumSwatch {
	s.shape = shape
	return s
}

// Rounding sets the corner rounding of the swatch
func (s *SpectrumSwatch) Rounding(rounding SwatchRounding) *SpectrumSwatch {
	s.rounding = rounding
	return s
}

// Border sets the border style of the swatch
func (s *SpectrumSwatch) Border(border SwatchBorder) *SpectrumSwatch {
	s.border = border
	return s
}

// Color sets the color of the swatch
func (s *SpectrumSwatch) Color(color string) *SpectrumSwatch {
	s.color = color
	return s
}

// Disabled sets the disabled state
func (s *SpectrumSwatch) Disabled(disabled bool) *SpectrumSwatch {
	s.disabled = disabled
	return s
}

// MixedValue sets the mixed value state (multiple colors represented)
func (s *SpectrumSwatch) MixedValue(mixed bool) *SpectrumSwatch {
	s.mixedValue = mixed
	return s
}

// Nothing sets the nothing state (represents transparent or no color)
func (s *SpectrumSwatch) Nothing(nothing bool) *SpectrumSwatch {
	s.nothing = nothing
	return s
}

// Selected sets the selected state
func (s *SpectrumSwatch) Selected(selected bool) *SpectrumSwatch {
	s.selected = selected
	return s
}

// Label sets the accessibility label
func (s *SpectrumSwatch) Label(label string) *SpectrumSwatch {
	s.label = label
	return s
}

// Value sets the value of the swatch (for selection)
func (s *SpectrumSwatch) Value(value string) *SpectrumSwatch {
	s.value = value
	return s
}

// TabIndex sets the tab index
func (s *SpectrumSwatch) TabIndex(tabIndex int) *SpectrumSwatch {
	s.tabIndex = tabIndex
	return s
}

// Role sets the ARIA role
func (s *SpectrumSwatch) Role(role string) *SpectrumSwatch {
	s.role = role
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSwatch) OnChange(handler app.EventHandler) *SpectrumSwatch {
	s.onChange = handler
	return s
}

// Convenience methods for setting specific shape variations

// Rectangle sets the shape to rectangle
func (s *SpectrumSwatch) Rectangle() *SpectrumSwatch {
	return s.Shape(SwatchShapeRectangle)
}

// Convenience methods for setting specific rounding variations

// RoundingNone sets rounding to none
func (s *SpectrumSwatch) RoundingNone() *SpectrumSwatch {
	return s.Rounding(SwatchRoundingNone)
}

// RoundingFull sets rounding to full
func (s *SpectrumSwatch) RoundingFull() *SpectrumSwatch {
	return s.Rounding(SwatchRoundingFull)
}

// Convenience methods for setting specific border variations

// BorderNone sets border to none
func (s *SpectrumSwatch) BorderNone() *SpectrumSwatch {
	return s.Border(SwatchBorderNone)
}

// BorderLight sets border to light
func (s *SpectrumSwatch) BorderLight() *SpectrumSwatch {
	return s.Border(SwatchBorderLight)
}

// Render renders the swatch component
func (s *SpectrumSwatch) Render() app.UI {
	swatch := app.Elem("sp-swatch").
		Attr("size", string(s.size)).
		Attr("color", s.color).
		Attr("disabled", s.disabled).
		Attr("mixed-value", s.mixedValue).
		Attr("nothing", s.nothing).
		Attr("selected", s.selected).
		Attr("label", s.label).
		Attr("value", s.value).
		Attr("tabindex", s.tabIndex).
		Attr("role", s.role)

	// Add shape if specified
	if s.shape != "" {
		swatch = swatch.Attr("shape", string(s.shape))
	}

	// Add rounding if specified
	if s.rounding != "" {
		swatch = swatch.Attr("rounding", string(s.rounding))
	}

	// Add border if specified
	if s.border != "" {
		swatch = swatch.Attr("border", string(s.border))
	}

	// Add event handlers
	if s.onChange != nil {
		swatch = swatch.On("change", s.onChange)
	}

	return swatch
}
