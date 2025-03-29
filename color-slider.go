package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorSlider represents the sp-color-slider component
// It lets users visually change an individual channel of a color, representing color properties
// such as hues, color channel values, or opacity.
type spectrumColorSlider struct {
	app.Compo

	// Properties
	PropColor    string
	PropDisabled bool
	PropFocused  bool
	PropLabel    string
	PropStep     float64
	PropTabIndex int
	PropValue    float64
	PropVertical bool

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// ColorSlider creates a new spectrumColorSlider instance.
func ColorSlider() *spectrumColorSlider {
	return &spectrumColorSlider{
		PropLabel: "hue",
		PropStep:  1,
	}
}

// Color sets the color property of the spectrumColorSlider
func (c *spectrumColorSlider) Color(color string) *spectrumColorSlider {
	c.PropColor = color
	return c
}

// Disabled sets whether the color slider is disabled
func (c *spectrumColorSlider) Disabled(disabled bool) *spectrumColorSlider {
	c.PropDisabled = disabled
	return c
}

// Focused sets whether the color slider is focused
func (c *spectrumColorSlider) Focused(focused bool) *spectrumColorSlider {
	c.PropFocused = focused
	return c
}

// Label sets the label property of the spectrumColorSlider
func (c *spectrumColorSlider) Label(label string) *spectrumColorSlider {
	c.PropLabel = label
	return c
}

// Step sets the step property of the spectrumColorSlider
func (c *spectrumColorSlider) Step(step float64) *spectrumColorSlider {
	c.PropStep = step
	return c
}

// TabIndex sets the tabIndex property of the spectrumColorSlider
func (c *spectrumColorSlider) TabIndex(tabIndex int) *spectrumColorSlider {
	c.PropTabIndex = tabIndex
	return c
}

// Value sets the value property of the spectrumColorSlider
func (c *spectrumColorSlider) Value(value float64) *spectrumColorSlider {
	c.PropValue = value
	return c
}

// Vertical sets whether the color slider is displayed vertically
func (c *spectrumColorSlider) Vertical(vertical bool) *spectrumColorSlider {
	c.PropVertical = vertical
	return c
}

// OnChange sets the change event handler
func (c *spectrumColorSlider) OnChange(handler app.EventHandler) *spectrumColorSlider {
	c.PropOnChange = handler
	return c
}

// OnInput sets the input event handler
func (c *spectrumColorSlider) OnInput(handler app.EventHandler) *spectrumColorSlider {
	c.PropOnInput = handler
	return c
}

// Render renders the spectrumColorSlider component.
func (c *spectrumColorSlider) Render() app.UI {
	colorSlider := app.Elem("sp-color-slider").
		Attr("label", c.PropLabel).
		Attr("step", c.PropStep)

	if c.PropColor != "" {
		colorSlider = colorSlider.Attr("color", c.PropColor)
	}
	if c.PropDisabled {
		colorSlider = colorSlider.Attr("disabled", true)
	}
	if c.PropFocused {
		colorSlider = colorSlider.Attr("focused", true)
	}
	if c.PropTabIndex != 0 {
		colorSlider = colorSlider.Attr("tabindex", c.PropTabIndex)
	}
	if c.PropValue != 0 {
		colorSlider = colorSlider.Attr("value", c.PropValue)
	}
	if c.PropVertical {
		colorSlider = colorSlider.Attr("vertical", true)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		colorSlider = colorSlider.On("change", c.PropOnChange)
	}
	if c.PropOnInput != nil {
		colorSlider = colorSlider.On("input", c.PropOnInput)
	}

	return colorSlider
}
