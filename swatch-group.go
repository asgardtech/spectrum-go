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

// SpectrumSwatchGroup represents an sp-swatch-group component
type SpectrumSwatchGroup struct {
	app.Compo

	// Properties
	size     SwatchSize
	shape    SwatchShape
	rounding SwatchRounding
	border   SwatchBorder
	density  SwatchDensity
	selects  SwatchSelects
	selected []string

	// Child swatches
	swatches []*SpectrumSwatch

	// Event handlers
	onChange app.EventHandler
}

// SwatchGroup creates a new swatch group component
func SwatchGroup() *SpectrumSwatchGroup {
	return &SpectrumSwatchGroup{
		size:     SwatchSizeM, // Default size is medium
		selected: []string{},
	}
}

// Size sets the visual size of all swatches in the group
func (s *SpectrumSwatchGroup) Size(size SwatchSize) *SpectrumSwatchGroup {
	s.size = size
	return s
}

// Shape sets the shape of all swatches in the group
func (s *SpectrumSwatchGroup) Shape(shape SwatchShape) *SpectrumSwatchGroup {
	s.shape = shape
	return s
}

// Rounding sets the corner rounding of all swatches in the group
func (s *SpectrumSwatchGroup) Rounding(rounding SwatchRounding) *SpectrumSwatchGroup {
	s.rounding = rounding
	return s
}

// Border sets the border style of all swatches in the group
func (s *SpectrumSwatchGroup) Border(border SwatchBorder) *SpectrumSwatchGroup {
	s.border = border
	return s
}

// Density sets the spacing density of the group
func (s *SpectrumSwatchGroup) Density(density SwatchDensity) *SpectrumSwatchGroup {
	s.density = density
	return s
}

// Selects sets the selection mode (single or multiple)
func (s *SpectrumSwatchGroup) Selects(selects SwatchSelects) *SpectrumSwatchGroup {
	s.selects = selects
	return s
}

// Selected sets the selected values in the group
func (s *SpectrumSwatchGroup) Selected(selected []string) *SpectrumSwatchGroup {
	s.selected = selected
	return s
}

// SelectedJSON sets the selected values in the group from a JSON string
func (s *SpectrumSwatchGroup) SelectedJSON(selectedJSON string) *SpectrumSwatchGroup {
	var selected []string
	if err := json.Unmarshal([]byte(selectedJSON), &selected); err == nil {
		s.selected = selected
	}
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSwatchGroup) OnChange(handler app.EventHandler) *SpectrumSwatchGroup {
	s.onChange = handler
	return s
}

// Swatches sets the child swatch components
func (s *SpectrumSwatchGroup) Swatches(swatches ...*SpectrumSwatch) *SpectrumSwatchGroup {
	s.swatches = swatches
	return s
}

// Convenience methods for setting specific density variations

// Compact sets density to compact
func (s *SpectrumSwatchGroup) Compact() *SpectrumSwatchGroup {
	return s.Density(SwatchDensityCompact)
}

// Spacious sets density to spacious
func (s *SpectrumSwatchGroup) Spacious() *SpectrumSwatchGroup {
	return s.Density(SwatchDensitySpacious)
}

// Convenience methods for setting specific selection modes

// SingleSelect sets selects to single
func (s *SpectrumSwatchGroup) SingleSelect() *SpectrumSwatchGroup {
	return s.Selects(SwatchSelectsSingle)
}

// MultipleSelect sets selects to multiple
func (s *SpectrumSwatchGroup) MultipleSelect() *SpectrumSwatchGroup {
	return s.Selects(SwatchSelectsMultiple)
}

// Convenience methods for setting specific shape variations

// Rectangle sets the shape to rectangle for all swatches
func (s *SpectrumSwatchGroup) Rectangle() *SpectrumSwatchGroup {
	return s.Shape(SwatchShapeRectangle)
}

// Convenience methods for setting specific rounding variations

// RoundingNone sets rounding to none for all swatches
func (s *SpectrumSwatchGroup) RoundingNone() *SpectrumSwatchGroup {
	return s.Rounding(SwatchRoundingNone)
}

// RoundingFull sets rounding to full for all swatches
func (s *SpectrumSwatchGroup) RoundingFull() *SpectrumSwatchGroup {
	return s.Rounding(SwatchRoundingFull)
}

// Convenience methods for setting specific border variations

// BorderNone sets border to none for all swatches
func (s *SpectrumSwatchGroup) BorderNone() *SpectrumSwatchGroup {
	return s.Border(SwatchBorderNone)
}

// BorderLight sets border to light for all swatches
func (s *SpectrumSwatchGroup) BorderLight() *SpectrumSwatchGroup {
	return s.Border(SwatchBorderLight)
}

// Render renders the swatch group component
func (s *SpectrumSwatchGroup) Render() app.UI {
	swatchGroup := app.Elem("sp-swatch-group")

	// Add size if specified
	if s.size != "" {
		swatchGroup = swatchGroup.Attr("size", string(s.size))
	}

	// Add shape if specified
	if s.shape != "" {
		swatchGroup = swatchGroup.Attr("shape", string(s.shape))
	}

	// Add rounding if specified
	if s.rounding != "" {
		swatchGroup = swatchGroup.Attr("rounding", string(s.rounding))
	}

	// Add border if specified
	if s.border != "" {
		swatchGroup = swatchGroup.Attr("border", string(s.border))
	}

	// Add density if specified
	if s.density != "" {
		swatchGroup = swatchGroup.Attr("density", string(s.density))
	}

	// Add selects if specified
	if s.selects != "" {
		swatchGroup = swatchGroup.Attr("selects", string(s.selects))
	}

	// Add selected values if specified
	if len(s.selected) > 0 {
		selectedJSON, _ := json.Marshal(s.selected)
		swatchGroup = swatchGroup.Attr("selected", string(selectedJSON))
	}

	// Add event handlers
	if s.onChange != nil {
		swatchGroup = swatchGroup.On("change", s.onChange)
	}

	// Add swatch elements if provided
	if len(s.swatches) > 0 {
		swatchUIs := make([]app.UI, len(s.swatches))
		for i, swatch := range s.swatches {
			swatchUIs[i] = swatch
		}
		swatchGroup = swatchGroup.Body(swatchUIs...)
	}

	return swatchGroup
}
