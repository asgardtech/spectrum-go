// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwatchGroupBorder represents the Border style for all child swatches
type SwatchGroupBorder string

// SwatchGroupBorder values
const (
	SwatchGroupBorderNone  SwatchGroupBorder = "none"
	SwatchGroupBorderLight SwatchGroupBorder = "light"
)

// SwatchGroupDensity represents the Spacing density between swatches
type SwatchGroupDensity string

// SwatchGroupDensity values
const (
	SwatchGroupDensityCompact  SwatchGroupDensity = "compact"
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
	SwatchGroupSelectsSingle   SwatchGroupSelects = "single"
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
	SwatchGroupSizeS  SwatchGroupSize = "s"
	SwatchGroupSizeM  SwatchGroupSize = "m"
	SwatchGroupSizeL  SwatchGroupSize = "l"
)

// spectrumSwatchGroup represents an sp-swatch-group component
type spectrumSwatchGroup struct {
	app.Compo
	*styler[*spectrumSwatchGroup]
	*classer[*spectrumSwatchGroup]
	*ider[*spectrumSwatchGroup]

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

	// Content for default slot
	PropBody []app.UI

	// Content slots

	// Event handlers
	PropOnChange app.EventHandler
}

// ISwatchGroup is the interface for sp-swatch-group component methods
type ISwatchGroup interface {
	app.UI
	Styler[ISwatchGroup]
	Classer[ISwatchGroup]
	Ider[ISwatchGroup]
	Border(SwatchGroupBorder) ISwatchGroup
	BorderNone() ISwatchGroup
	BorderLight() ISwatchGroup
	Density(SwatchGroupDensity) ISwatchGroup
	DensityCompact() ISwatchGroup
	DensitySpacious() ISwatchGroup
	Rounding(SwatchGroupRounding) ISwatchGroup
	RoundingNone() ISwatchGroup
	RoundingFull() ISwatchGroup
	Selects(SwatchGroupSelects) ISwatchGroup
	SelectsSingle() ISwatchGroup
	SelectsMultiple() ISwatchGroup
	Shape(SwatchGroupShape) ISwatchGroup
	ShapeRectangle() ISwatchGroup
	Size(SwatchGroupSize) ISwatchGroup
	SizeXs() ISwatchGroup
	SizeS() ISwatchGroup
	SizeM() ISwatchGroup
	SizeL() ISwatchGroup

	Body(...app.UI) ISwatchGroup
	AddToBody(app.UI) ISwatchGroup
	Text(string) ISwatchGroup

	OnChange(app.EventHandler) ISwatchGroup
}

// SwatchGroup A grouping of swatch elements that are related to each other.
func SwatchGroup() ISwatchGroup {
	element := &spectrumSwatchGroup{
		PropBorder:   "",
		PropDensity:  "",
		PropRounding: "",
		PropSelects:  "",
		PropShape:    "",
		PropSize:     SwatchGroupSizeM,
		PropBody:     []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Border Border style for all child swatches
func (c *spectrumSwatchGroup) Border(border SwatchGroupBorder) ISwatchGroup {
	c.PropBorder = border
	return c
}

func (c *spectrumSwatchGroup) BorderNone() ISwatchGroup {
	return c.Border(SwatchGroupBorderNone)
}
func (c *spectrumSwatchGroup) BorderLight() ISwatchGroup {
	return c.Border(SwatchGroupBorderLight)
}

// Density Spacing density between swatches
func (c *spectrumSwatchGroup) Density(density SwatchGroupDensity) ISwatchGroup {
	c.PropDensity = density
	return c
}

func (c *spectrumSwatchGroup) DensityCompact() ISwatchGroup {
	return c.Density(SwatchGroupDensityCompact)
}
func (c *spectrumSwatchGroup) DensitySpacious() ISwatchGroup {
	return c.Density(SwatchGroupDensitySpacious)
}

// Rounding Rounding style for all child swatches
func (c *spectrumSwatchGroup) Rounding(rounding SwatchGroupRounding) ISwatchGroup {
	c.PropRounding = rounding
	return c
}

func (c *spectrumSwatchGroup) RoundingNone() ISwatchGroup {
	return c.Rounding(SwatchGroupRoundingNone)
}
func (c *spectrumSwatchGroup) RoundingFull() ISwatchGroup {
	return c.Rounding(SwatchGroupRoundingFull)
}

// Selects Selection mode for the swatch group
func (c *spectrumSwatchGroup) Selects(selects SwatchGroupSelects) ISwatchGroup {
	c.PropSelects = selects
	return c
}

func (c *spectrumSwatchGroup) SelectsSingle() ISwatchGroup {
	return c.Selects(SwatchGroupSelectsSingle)
}
func (c *spectrumSwatchGroup) SelectsMultiple() ISwatchGroup {
	return c.Selects(SwatchGroupSelectsMultiple)
}

// Shape Shape of all child swatches
func (c *spectrumSwatchGroup) Shape(shape SwatchGroupShape) ISwatchGroup {
	c.PropShape = shape
	return c
}

func (c *spectrumSwatchGroup) ShapeRectangle() ISwatchGroup {
	return c.Shape(SwatchGroupShapeRectangle)
}

// Size Size of all child swatches
func (c *spectrumSwatchGroup) Size(size SwatchGroupSize) ISwatchGroup {
	c.PropSize = size
	return c
}

func (c *spectrumSwatchGroup) SizeXs() ISwatchGroup {
	return c.Size(SwatchGroupSizeXs)
}
func (c *spectrumSwatchGroup) SizeS() ISwatchGroup {
	return c.Size(SwatchGroupSizeS)
}
func (c *spectrumSwatchGroup) SizeM() ISwatchGroup {
	return c.Size(SwatchGroupSizeM)
}
func (c *spectrumSwatchGroup) SizeL() ISwatchGroup {
	return c.Size(SwatchGroupSizeL)
}

// Body sets the content for the default slot
func (c *spectrumSwatchGroup) Body(elements ...app.UI) ISwatchGroup {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumSwatchGroup) AddToBody(element app.UI) ISwatchGroup {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumSwatchGroup) Text(text string) ISwatchGroup {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Fired when the selection within the group changes
func (c *spectrumSwatchGroup) OnChange(handler app.EventHandler) ISwatchGroup {
	c.PropOnChange = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumSwatchGroup) Style(key, format string, values ...any) ISwatchGroup {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumSwatchGroup) Styles(styles map[string]string) ISwatchGroup {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumSwatchGroup) Class(class string) ISwatchGroup {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumSwatchGroup) Classes(classes ...string) ISwatchGroup {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumSwatchGroup) Id(id string) ISwatchGroup {
	return c.ider.Id(id)
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

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add all elements to the component
	if len(slotElements) > 0 {
		element = element.Body(slotElements...)
	}

	// Apply styles, classes, and id
	element = element.Styles(c.styler.styles)

	// Apply classes if any
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}

	// Apply id if set
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}

	return element
}
