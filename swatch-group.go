package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SwatchDensity represents the density options for swatch groups
type SwatchDensity string

// Swatch density values
const (
	SwatchDensityCompact  SwatchDensity = "compact"
	SwatchDensitySpacious SwatchDensity = "spacious"
)

// SwatchSelects represents the selection mode for swatch groups
type SwatchSelects string

// Swatch selection mode values
const (
	SwatchSelectsSingle   SwatchSelects = "single"
	SwatchSelectsMultiple SwatchSelects = "multiple"
)

// spectrumSwatchGroup represents an sp-swatch-group component
type spectrumSwatchGroup struct {
	app.Compo

	// Properties
	PropSize     SwatchSize
	PropShape    SwatchShape
	PropRounding SwatchRounding
	PropBorder   SwatchBorder
	PropDensity  SwatchDensity
	PropSelects  SwatchSelects
	PropSelected []string

	// Child swatches
	PropSwatches []app.UI

	// Event handlers
	PropOnChange app.EventHandler
}

// SwatchGroup creates a new swatch group component
func SwatchGroup() *spectrumSwatchGroup {
	return &spectrumSwatchGroup{
		PropSize:     SwatchSizeM, // Default size is medium
		PropSelected: []string{},
	}
}

// Size sets the visual size of all swatches in the group
func (s *spectrumSwatchGroup) Size(size SwatchSize) *spectrumSwatchGroup {
	s.PropSize = size
	return s
}

// Shape sets the shape of all swatches in the group
func (s *spectrumSwatchGroup) Shape(shape SwatchShape) *spectrumSwatchGroup {
	s.PropShape = shape
	return s
}

// Rounding sets the corner rounding of all swatches in the group
func (s *spectrumSwatchGroup) Rounding(rounding SwatchRounding) *spectrumSwatchGroup {
	s.PropRounding = rounding
	return s
}

// Border sets the border style of all swatches in the group
func (s *spectrumSwatchGroup) Border(border SwatchBorder) *spectrumSwatchGroup {
	s.PropBorder = border
	return s
}

// Density sets the spacing density of the group
func (s *spectrumSwatchGroup) Density(density SwatchDensity) *spectrumSwatchGroup {
	s.PropDensity = density
	return s
}

// Selects sets the selection mode (single or multiple)
func (s *spectrumSwatchGroup) Selects(selects SwatchSelects) *spectrumSwatchGroup {
	s.PropSelects = selects
	return s
}

// Selected sets the selected values in the group
func (s *spectrumSwatchGroup) Selected(selected []string) *spectrumSwatchGroup {
	s.PropSelected = selected
	return s
}

// SelectedJSON sets the selected values in the group from a JSON string
func (s *spectrumSwatchGroup) SelectedJSON(selectedJSON string) *spectrumSwatchGroup {
	var selected []string
	if err := json.Unmarshal([]byte(selectedJSON), &selected); err == nil {
		s.PropSelected = selected
	}
	return s
}

// OnChange sets the change event handler
func (s *spectrumSwatchGroup) OnChange(handler app.EventHandler) *spectrumSwatchGroup {
	s.PropOnChange = handler
	return s
}

// Swatches sets the child swatch components
func (s *spectrumSwatchGroup) Swatches(swatches ...app.UI) *spectrumSwatchGroup {
	s.PropSwatches = swatches
	return s
}

// Convenience methods for setting specific density variations

// Compact sets density to compact
func (s *spectrumSwatchGroup) Compact() *spectrumSwatchGroup {
	return s.Density(SwatchDensityCompact)
}

// Spacious sets density to spacious
func (s *spectrumSwatchGroup) Spacious() *spectrumSwatchGroup {
	return s.Density(SwatchDensitySpacious)
}

// Convenience methods for setting specific selection modes

// SingleSelect sets selects to single
func (s *spectrumSwatchGroup) SingleSelect() *spectrumSwatchGroup {
	return s.Selects(SwatchSelectsSingle)
}

// MultipleSelect sets selects to multiple
func (s *spectrumSwatchGroup) MultipleSelect() *spectrumSwatchGroup {
	return s.Selects(SwatchSelectsMultiple)
}

// Convenience methods for setting specific shape variations

// Rectangle sets the shape to rectangle for all swatches
func (s *spectrumSwatchGroup) Rectangle() *spectrumSwatchGroup {
	return s.Shape(SwatchShapeRectangle)
}

// Convenience methods for setting specific rounding variations

// RoundingNone sets rounding to none for all swatches
func (s *spectrumSwatchGroup) RoundingNone() *spectrumSwatchGroup {
	return s.Rounding(SwatchRoundingNone)
}

// RoundingFull sets rounding to full for all swatches
func (s *spectrumSwatchGroup) RoundingFull() *spectrumSwatchGroup {
	return s.Rounding(SwatchRoundingFull)
}

// Convenience methods for setting specific border variations

// BorderNone sets border to none for all swatches
func (s *spectrumSwatchGroup) BorderNone() *spectrumSwatchGroup {
	return s.Border(SwatchBorderNone)
}

// BorderLight sets border to light for all swatches
func (s *spectrumSwatchGroup) BorderLight() *spectrumSwatchGroup {
	return s.Border(SwatchBorderLight)
}

// Render renders the swatch group component
func (s *spectrumSwatchGroup) Render() app.UI {
	swatchGroup := app.Elem("sp-swatch-group")

	// Add size if specified
	if s.PropSize != "" {
		swatchGroup = swatchGroup.Attr("size", string(s.PropSize))
	}

	// Add shape if specified
	if s.PropShape != "" {
		swatchGroup = swatchGroup.Attr("shape", string(s.PropShape))
	}

	// Add rounding if specified
	if s.PropRounding != "" {
		swatchGroup = swatchGroup.Attr("rounding", string(s.PropRounding))
	}

	// Add border if specified
	if s.PropBorder != "" {
		swatchGroup = swatchGroup.Attr("border", string(s.PropBorder))
	}

	// Add density if specified
	if s.PropDensity != "" {
		swatchGroup = swatchGroup.Attr("density", string(s.PropDensity))
	}

	// Add selects if specified
	if s.PropSelects != "" {
		swatchGroup = swatchGroup.Attr("selects", string(s.PropSelects))
	}

	// Add selected values if specified
	if len(s.PropSelected) > 0 {
		selectedJSON, _ := json.Marshal(s.PropSelected)
		swatchGroup = swatchGroup.Attr("selected", string(selectedJSON))
	}

	// Add event handlers
	if s.PropOnChange != nil {
		swatchGroup = swatchGroup.On("change", s.PropOnChange)
	}

	// Add swatch elements if provided
	if len(s.PropSwatches) > 0 {
		swatchGroup = swatchGroup.Body(s.PropSwatches...)
	}

	return swatchGroup
}
