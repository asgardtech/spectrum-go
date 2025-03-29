package sp

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// spectrumColorArea represents an sp-color-area component.
type spectrumColorArea struct {
	app.Compo

	// Properties
	PropColor    string
	PropDisabled bool
	PropFocused  bool
	PropHue      float64
	PropLabelX   string
	PropLabelY   string
	PropStep     float64
	PropValue    string
	PropX        float64
	PropY        float64
	PropGradient app.UI
	PropStyle    string

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// ColorArea creates a new spectrumColorArea instance.
func ColorArea() *spectrumColorArea {
	return &spectrumColorArea{
		PropLabelX: "saturation",
		PropLabelY: "luminosity",
		PropStep:   0.01,
	}
}

// Color sets the color for the color area.
func (c *spectrumColorArea) Color(color string) *spectrumColorArea {
	c.PropColor = color
	return c
}

// Disabled sets whether the color area is disabled.
func (c *spectrumColorArea) Disabled(disabled bool) *spectrumColorArea {
	c.PropDisabled = disabled
	return c
}

// Focused sets whether the color area is focused.
func (c *spectrumColorArea) Focused(focused bool) *spectrumColorArea {
	c.PropFocused = focused
	return c
}

// Hue sets the hue value for the color area.
func (c *spectrumColorArea) Hue(hue float64) *spectrumColorArea {
	c.PropHue = hue
	return c
}

// LabelX sets the x-axis label.
func (c *spectrumColorArea) LabelX(label string) *spectrumColorArea {
	c.PropLabelX = label
	return c
}

// LabelY sets the y-axis label.
func (c *spectrumColorArea) LabelY(label string) *spectrumColorArea {
	c.PropLabelY = label
	return c
}

// Step sets the step value for the color area.
func (c *spectrumColorArea) Step(step float64) *spectrumColorArea {
	c.PropStep = step
	return c
}

// Value sets the value for the color area.
func (c *spectrumColorArea) Value(value string) *spectrumColorArea {
	c.PropValue = value
	return c
}

// X sets the x position value.
func (c *spectrumColorArea) X(x float64) *spectrumColorArea {
	c.PropX = x
	return c
}

// Y sets the y position value.
func (c *spectrumColorArea) Y(y float64) *spectrumColorArea {
	c.PropY = y
	return c
}

// Gradient sets a custom gradient UI for the color area.
func (c *spectrumColorArea) Gradient(gradient app.UI) *spectrumColorArea {
	c.PropGradient = gradient
	return c
}

// Style sets custom CSS styles for the color area.
func (c *spectrumColorArea) Style(style string) *spectrumColorArea {
	c.PropStyle = style
	return c
}

// Width sets the width of the color area.
func (c *spectrumColorArea) Width(width int) *spectrumColorArea {
	if c.PropStyle == "" {
		c.PropStyle = "width: " + strconv.Itoa(width) + "px;"
	} else {
		c.PropStyle += " width: " + strconv.Itoa(width) + "px;"
	}
	return c
}

// Height sets the height of the color area.
func (c *spectrumColorArea) Height(height int) *spectrumColorArea {
	if c.PropStyle == "" {
		c.PropStyle = "height: " + strconv.Itoa(height) + "px;"
	} else {
		c.PropStyle += " height: " + strconv.Itoa(height) + "px;"
	}
	return c
}

// Size sets both width and height of the color area to the same value.
func (c *spectrumColorArea) Size(size int) *spectrumColorArea {
	if c.PropStyle == "" {
		c.PropStyle = "width: " + strconv.Itoa(size) + "px; height: " + strconv.Itoa(size) + "px;"
	} else {
		c.PropStyle += " width: " + strconv.Itoa(size) + "px; height: " + strconv.Itoa(size) + "px;"
	}
	return c
}

// OnChange sets the handler for when a user commits a color change.
func (c *spectrumColorArea) OnChange(h app.EventHandler) *spectrumColorArea {
	c.PropOnChange = h
	return c
}

// OnInput sets the handler for when the color value changes.
func (c *spectrumColorArea) OnInput(h app.EventHandler) *spectrumColorArea {
	c.PropOnInput = h
	return c
}

// Render renders the component.
func (c *spectrumColorArea) Render() app.UI {
	colorArea := app.Elem("sp-color-area")

	// Set attributes
	if c.PropColor != "" {
		colorArea.Attr("color", c.PropColor)
	}

	if c.PropDisabled {
		colorArea.Attr("disabled", "")
	}

	if c.PropFocused {
		colorArea.Attr("focused", "")
	}

	if c.PropHue != 0 {
		colorArea.Attr("hue", c.PropHue)
	}

	if c.PropLabelX != "saturation" {
		colorArea.Attr("label-x", c.PropLabelX)
	}

	if c.PropLabelY != "luminosity" {
		colorArea.Attr("label-y", c.PropLabelY)
	}

	if c.PropStep != 0.01 {
		colorArea.Attr("step", c.PropStep)
	}

	if c.PropValue != "" {
		colorArea.Attr("value", c.PropValue)
	}

	if c.PropX != 0 {
		colorArea.Attr("x", c.PropX)
	}

	if c.PropY != 0 {
		colorArea.Attr("y", c.PropY)
	}

	if c.PropStyle != "" {
		colorArea.Attr("style", c.PropStyle)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		colorArea = colorArea.On("change", c.PropOnChange)
	}

	if c.PropOnInput != nil {
		colorArea = colorArea.On("input", c.PropOnInput)
	}

	// Add the gradient slot if provided
	if c.PropGradient != nil {
		colorArea = colorArea.Body(
			app.Elem("div").Attr("slot", "gradient").Body(c.PropGradient),
		)
	}

	return colorArea
}
