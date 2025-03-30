package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwatchGroupBorder represents the Border style for all child swatches
type SwatchGroupBorder string

// SwatchGroupBorder values
const (
    SwatchGroupBorderNone SwatchGroupBorder = "none"
    SwatchGroupBorderLight SwatchGroupBorder = "light"
)
// SwatchGroupDensity represents the Spacing density between swatches
type SwatchGroupDensity string

// SwatchGroupDensity values
const (
    SwatchGroupDensityCompact SwatchGroupDensity = "compact"
    SwatchGroupDensitySpacious SwatchGroupDensity = "spacious"
)
// SwatchGroupRounding represents the Rounding style for all child swatches
type SwatchGroupRounding string

// SwatchGroupRounding values
const (
    SwatchGroupRoundingNone SwatchGroupRounding = "none"
    SwatchGroupRoundingFull SwatchGroupRounding = "full"
)
// SwatchGroupSelects represents the Selection mode for the swatch group
type SwatchGroupSelects string

// SwatchGroupSelects values
const (
    SwatchGroupSelectsSingle SwatchGroupSelects = "single"
    SwatchGroupSelectsMultiple SwatchGroupSelects = "multiple"
)
// SwatchGroupShape represents the Shape of all child swatches
type SwatchGroupShape string

// SwatchGroupShape values
const (
    SwatchGroupShapeRectangle SwatchGroupShape = "rectangle"
)
// SwatchGroupSize represents the Size of all child swatches
type SwatchGroupSize string

// SwatchGroupSize values
const (
    SwatchGroupSizeXs SwatchGroupSize = "xs"
    SwatchGroupSizeS SwatchGroupSize = "s"
    SwatchGroupSizeM SwatchGroupSize = "m"
    SwatchGroupSizeL SwatchGroupSize = "l"
)

// spectrumSwatchGroup represents an sp-swatch-group component
type spectrumSwatchGroup struct {
    app.Compo
    *styler[*spectrumSwatchGroup]

    // Properties
    // Border style for all child swatches
    PropBorder SwatchGroupBorder
    // Spacing density between swatches
    PropDensity SwatchGroupDensity
    // Rounding style for all child swatches
    PropRounding SwatchGroupRounding
    // Selection mode for the swatch group
    PropSelects SwatchGroupSelects
    // Shape of all child swatches
    PropShape SwatchGroupShape
    // Size of all child swatches
    PropSize SwatchGroupSize

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// SwatchGroup creates a new sp-swatch-group component
func SwatchGroup() *spectrumSwatchGroup {
    element := &spectrumSwatchGroup{
        PropBorder: "",
        PropDensity: "",
        PropRounding: "",
        PropSelects: "",
        PropShape: "",
        PropSize: SwatchGroupSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Border style for all child swatches
func (c *spectrumSwatchGroup) Border(border SwatchGroupBorder) *spectrumSwatchGroup {
    c.PropBorder = border
    return c
}

func (c *spectrumSwatchGroup) BorderNone() *spectrumSwatchGroup {
    return c.Border(SwatchGroupBorderNone)
}
func (c *spectrumSwatchGroup) BorderLight() *spectrumSwatchGroup {
    return c.Border(SwatchGroupBorderLight)
}
// Spacing density between swatches
func (c *spectrumSwatchGroup) Density(density SwatchGroupDensity) *spectrumSwatchGroup {
    c.PropDensity = density
    return c
}

func (c *spectrumSwatchGroup) DensityCompact() *spectrumSwatchGroup {
    return c.Density(SwatchGroupDensityCompact)
}
func (c *spectrumSwatchGroup) DensitySpacious() *spectrumSwatchGroup {
    return c.Density(SwatchGroupDensitySpacious)
}
// Rounding style for all child swatches
func (c *spectrumSwatchGroup) Rounding(rounding SwatchGroupRounding) *spectrumSwatchGroup {
    c.PropRounding = rounding
    return c
}

func (c *spectrumSwatchGroup) RoundingNone() *spectrumSwatchGroup {
    return c.Rounding(SwatchGroupRoundingNone)
}
func (c *spectrumSwatchGroup) RoundingFull() *spectrumSwatchGroup {
    return c.Rounding(SwatchGroupRoundingFull)
}
// Selection mode for the swatch group
func (c *spectrumSwatchGroup) Selects(selects SwatchGroupSelects) *spectrumSwatchGroup {
    c.PropSelects = selects
    return c
}

func (c *spectrumSwatchGroup) SelectsSingle() *spectrumSwatchGroup {
    return c.Selects(SwatchGroupSelectsSingle)
}
func (c *spectrumSwatchGroup) SelectsMultiple() *spectrumSwatchGroup {
    return c.Selects(SwatchGroupSelectsMultiple)
}
// Shape of all child swatches
func (c *spectrumSwatchGroup) Shape(shape SwatchGroupShape) *spectrumSwatchGroup {
    c.PropShape = shape
    return c
}

func (c *spectrumSwatchGroup) ShapeRectangle() *spectrumSwatchGroup {
    return c.Shape(SwatchGroupShapeRectangle)
}
// Size of all child swatches
func (c *spectrumSwatchGroup) Size(size SwatchGroupSize) *spectrumSwatchGroup {
    c.PropSize = size
    return c
}

func (c *spectrumSwatchGroup) SizeXs() *spectrumSwatchGroup {
    return c.Size(SwatchGroupSizeXs)
}
func (c *spectrumSwatchGroup) SizeS() *spectrumSwatchGroup {
    return c.Size(SwatchGroupSizeS)
}
func (c *spectrumSwatchGroup) SizeM() *spectrumSwatchGroup {
    return c.Size(SwatchGroupSizeM)
}
func (c *spectrumSwatchGroup) SizeL() *spectrumSwatchGroup {
    return c.Size(SwatchGroupSizeL)
}

// Text sets the text content for the default slot
func (c *spectrumSwatchGroup) Text(text string) *spectrumSwatchGroup {
    c.PropText = text
    return c
}




// Fired when the selection within the group changes
func (c *spectrumSwatchGroup) OnChange(handler app.EventHandler) *spectrumSwatchGroup {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-swatch-group component
func (c *spectrumSwatchGroup) Render() app.UI {
    element := app.Elem("sp-swatch-group")

    // Set attributes
    if c.PropBorder != "" {
        element = element.Attr("border", string(c.PropBorder))
    }
    if c.PropDensity != "" {
        element = element.Attr("density", string(c.PropDensity))
    }
    if c.PropRounding != "" {
        element = element.Attr("rounding", string(c.PropRounding))
    }
    if c.PropSelects != "" {
        element = element.Attr("selects", string(c.PropSelects))
    }
    if c.PropShape != "" {
        element = element.Attr("shape", string(c.PropShape))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 