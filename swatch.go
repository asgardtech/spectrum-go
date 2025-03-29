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

// spectrumSwatch represents an sp-swatch component
type spectrumSwatch struct {
	app.Compo

	// Properties
	PropSize       SwatchSize
	PropShape      SwatchShape
	PropRounding   SwatchRounding
	PropBorder     SwatchBorder
	PropColor      string
	PropDisabled   bool
	PropMixedValue bool
	PropNothing    bool
	PropSelected   bool
	PropLabel      string
	PropValue      string
	PropTabIndex   int
	PropRole       string

	// Event handlers
	PropOnChange app.EventHandler
}

// Swatch creates a new swatch component
func Swatch() *spectrumSwatch {
	return &spectrumSwatch{
		PropSize: SwatchSizeM, // Default size is medium
		PropRole: "button",    // Default role is button
	}
}

// Size sets the visual size of the swatch
func (s *spectrumSwatch) Size(size SwatchSize) *spectrumSwatch {
	s.PropSize = size
	return s
}

// Shape sets the shape of the swatch
func (s *spectrumSwatch) Shape(shape SwatchShape) *spectrumSwatch {
	s.PropShape = shape
	return s
}

// Rounding sets the corner rounding of the swatch
func (s *spectrumSwatch) Rounding(rounding SwatchRounding) *spectrumSwatch {
	s.PropRounding = rounding
	return s
}

// Border sets the border style of the swatch
func (s *spectrumSwatch) Border(border SwatchBorder) *spectrumSwatch {
	s.PropBorder = border
	return s
}

// Color sets the color of the swatch
func (s *spectrumSwatch) Color(color string) *spectrumSwatch {
	s.PropColor = color
	return s
}

// Disabled sets the disabled state
func (s *spectrumSwatch) Disabled(disabled bool) *spectrumSwatch {
	s.PropDisabled = disabled
	return s
}

// MixedValue sets the mixed value state (multiple colors represented)
func (s *spectrumSwatch) MixedValue(mixed bool) *spectrumSwatch {
	s.PropMixedValue = mixed
	return s
}

// Nothing sets the nothing state (represents transparent or no color)
func (s *spectrumSwatch) Nothing(nothing bool) *spectrumSwatch {
	s.PropNothing = nothing
	return s
}

// Selected sets the selected state
func (s *spectrumSwatch) Selected(selected bool) *spectrumSwatch {
	s.PropSelected = selected
	return s
}

// Label sets the accessibility label
func (s *spectrumSwatch) Label(label string) *spectrumSwatch {
	s.PropLabel = label
	return s
}

// Value sets the value of the swatch (for selection)
func (s *spectrumSwatch) Value(value string) *spectrumSwatch {
	s.PropValue = value
	return s
}

// TabIndex sets the tab index
func (s *spectrumSwatch) TabIndex(tabIndex int) *spectrumSwatch {
	s.PropTabIndex = tabIndex
	return s
}

// Role sets the ARIA role
func (s *spectrumSwatch) Role(role string) *spectrumSwatch {
	s.PropRole = role
	return s
}

// OnChange sets the change event handler
func (s *spectrumSwatch) OnChange(handler app.EventHandler) *spectrumSwatch {
	s.PropOnChange = handler
	return s
}

// Convenience methods for setting specific shape variations

// Rectangle sets the shape to rectangle
func (s *spectrumSwatch) Rectangle() *spectrumSwatch {
	return s.Shape(SwatchShapeRectangle)
}

// Convenience methods for setting specific rounding variations

// RoundingNone sets rounding to none
func (s *spectrumSwatch) RoundingNone() *spectrumSwatch {
	return s.Rounding(SwatchRoundingNone)
}

// RoundingFull sets rounding to full
func (s *spectrumSwatch) RoundingFull() *spectrumSwatch {
	return s.Rounding(SwatchRoundingFull)
}

// Convenience methods for setting specific border variations

// BorderNone sets border to none
func (s *spectrumSwatch) BorderNone() *spectrumSwatch {
	return s.Border(SwatchBorderNone)
}

// BorderLight sets border to light
func (s *spectrumSwatch) BorderLight() *spectrumSwatch {
	return s.Border(SwatchBorderLight)
}

// Render renders the swatch component
func (s *spectrumSwatch) Render() app.UI {
	swatch := app.Elem("sp-swatch").
		Attr("size", string(s.PropSize)).
		Attr("color", s.PropColor).
		Attr("disabled", s.PropDisabled).
		Attr("mixed-value", s.PropMixedValue).
		Attr("nothing", s.PropNothing).
		Attr("selected", s.PropSelected).
		Attr("label", s.PropLabel).
		Attr("value", s.PropValue).
		Attr("tabindex", s.PropTabIndex).
		Attr("role", s.PropRole)

	// Add shape if specified
	if s.PropShape != "" {
		swatch = swatch.Attr("shape", string(s.PropShape))
	}

	// Add rounding if specified
	if s.PropRounding != "" {
		swatch = swatch.Attr("rounding", string(s.PropRounding))
	}

	// Add border if specified
	if s.PropBorder != "" {
		swatch = swatch.Attr("border", string(s.PropBorder))
	}

	// Add event handlers
	if s.PropOnChange != nil {
		swatch = swatch.On("change", s.PropOnChange)
	}

	return swatch
}
