package sp

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SpectrumColorArea represents an sp-color-area component.
type SpectrumColorArea struct {
	app.Compo

	// Properties
	color    string
	disabled bool
	focused  bool
	hue      float64
	labelX   string
	labelY   string
	step     float64
	value    string
	x        float64
	y        float64
	gradient app.UI
	style    string

	// Event handlers
	onChange app.EventHandler
	onInput  app.EventHandler
}

// ColorArea creates a new SpectrumColorArea instance.
func ColorArea() *SpectrumColorArea {
	return &SpectrumColorArea{
		labelX: "saturation",
		labelY: "luminosity",
		step:   0.01,
	}
}

// Color sets the color for the color area.
func (c *SpectrumColorArea) Color(color string) *SpectrumColorArea {
	c.color = color
	return c
}

// Disabled sets whether the color area is disabled.
func (c *SpectrumColorArea) Disabled(disabled bool) *SpectrumColorArea {
	c.disabled = disabled
	return c
}

// Focused sets whether the color area is focused.
func (c *SpectrumColorArea) Focused(focused bool) *SpectrumColorArea {
	c.focused = focused
	return c
}

// Hue sets the hue value for the color area.
func (c *SpectrumColorArea) Hue(hue float64) *SpectrumColorArea {
	c.hue = hue
	return c
}

// LabelX sets the x-axis label.
func (c *SpectrumColorArea) LabelX(label string) *SpectrumColorArea {
	c.labelX = label
	return c
}

// LabelY sets the y-axis label.
func (c *SpectrumColorArea) LabelY(label string) *SpectrumColorArea {
	c.labelY = label
	return c
}

// Step sets the step value for the color area.
func (c *SpectrumColorArea) Step(step float64) *SpectrumColorArea {
	c.step = step
	return c
}

// Value sets the value for the color area.
func (c *SpectrumColorArea) Value(value string) *SpectrumColorArea {
	c.value = value
	return c
}

// X sets the x position value.
func (c *SpectrumColorArea) X(x float64) *SpectrumColorArea {
	c.x = x
	return c
}

// Y sets the y position value.
func (c *SpectrumColorArea) Y(y float64) *SpectrumColorArea {
	c.y = y
	return c
}

// Gradient sets a custom gradient UI for the color area.
func (c *SpectrumColorArea) Gradient(gradient app.UI) *SpectrumColorArea {
	c.gradient = gradient
	return c
}

// Style sets custom CSS styles for the color area.
func (c *SpectrumColorArea) Style(style string) *SpectrumColorArea {
	c.style = style
	return c
}

// Width sets the width of the color area.
func (c *SpectrumColorArea) Width(width int) *SpectrumColorArea {
	if c.style == "" {
		c.style = "width: " + strconv.Itoa(width) + "px;"
	} else {
		c.style += " width: " + strconv.Itoa(width) + "px;"
	}
	return c
}

// Height sets the height of the color area.
func (c *SpectrumColorArea) Height(height int) *SpectrumColorArea {
	if c.style == "" {
		c.style = "height: " + strconv.Itoa(height) + "px;"
	} else {
		c.style += " height: " + strconv.Itoa(height) + "px;"
	}
	return c
}

// Size sets both width and height of the color area to the same value.
func (c *SpectrumColorArea) Size(size int) *SpectrumColorArea {
	if c.style == "" {
		c.style = "width: " + strconv.Itoa(size) + "px; height: " + strconv.Itoa(size) + "px;"
	} else {
		c.style += " width: " + strconv.Itoa(size) + "px; height: " + strconv.Itoa(size) + "px;"
	}
	return c
}

// OnChange sets the handler for when a user commits a color change.
func (c *SpectrumColorArea) OnChange(h app.EventHandler) *SpectrumColorArea {
	c.onChange = h
	return c
}

// OnInput sets the handler for when the color value changes.
func (c *SpectrumColorArea) OnInput(h app.EventHandler) *SpectrumColorArea {
	c.onInput = h
	return c
}

// Render renders the component.
func (c *SpectrumColorArea) Render() app.UI {
	colorArea := app.Elem("sp-color-area")

	// Set attributes
	if c.color != "" {
		colorArea.Attr("color", c.color)
	}

	if c.disabled {
		colorArea.Attr("disabled", "")
	}

	if c.focused {
		colorArea.Attr("focused", "")
	}

	if c.hue != 0 {
		colorArea.Attr("hue", c.hue)
	}

	if c.labelX != "saturation" {
		colorArea.Attr("label-x", c.labelX)
	}

	if c.labelY != "luminosity" {
		colorArea.Attr("label-y", c.labelY)
	}

	if c.step != 0.01 {
		colorArea.Attr("step", c.step)
	}

	if c.value != "" {
		colorArea.Attr("value", c.value)
	}

	if c.x != 0 {
		colorArea.Attr("x", c.x)
	}

	if c.y != 0 {
		colorArea.Attr("y", c.y)
	}

	if c.style != "" {
		colorArea.Attr("style", c.style)
	}

	// Add event handlers
	if c.onChange != nil {
		colorArea = colorArea.On("change", c.onChange)
	}

	if c.onInput != nil {
		colorArea = colorArea.On("input", c.onInput)
	}

	// Add the gradient slot if provided
	if c.gradient != nil {
		colorArea = colorArea.Body(
			app.Elem("div").Attr("slot", "gradient").Body(c.gradient),
		)
	}

	return colorArea
}
