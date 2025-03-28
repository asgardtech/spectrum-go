package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumColorSlider represents the sp-color-slider component
// It lets users visually change an individual channel of a color, representing color properties
// such as hues, color channel values, or opacity.
type SpectrumColorSlider struct {
	app.Compo

	// Properties
	color    string
	disabled bool
	focused  bool
	label    string
	step     float64
	tabIndex int
	value    float64
	vertical bool

	// Event handlers
	onChange app.EventHandler
	onInput  app.EventHandler
}

// ColorSlider creates a new SpectrumColorSlider instance.
func ColorSlider() *SpectrumColorSlider {
	return &SpectrumColorSlider{
		label: "hue",
		step:  1,
	}
}

// Color sets the color property of the SpectrumColorSlider
func (c *SpectrumColorSlider) Color(color string) *SpectrumColorSlider {
	c.color = color
	return c
}

// Disabled sets whether the color slider is disabled
func (c *SpectrumColorSlider) Disabled(disabled bool) *SpectrumColorSlider {
	c.disabled = disabled
	return c
}

// Focused sets whether the color slider is focused
func (c *SpectrumColorSlider) Focused(focused bool) *SpectrumColorSlider {
	c.focused = focused
	return c
}

// Label sets the label property of the SpectrumColorSlider
func (c *SpectrumColorSlider) Label(label string) *SpectrumColorSlider {
	c.label = label
	return c
}

// Step sets the step property of the SpectrumColorSlider
func (c *SpectrumColorSlider) Step(step float64) *SpectrumColorSlider {
	c.step = step
	return c
}

// TabIndex sets the tabIndex property of the SpectrumColorSlider
func (c *SpectrumColorSlider) TabIndex(tabIndex int) *SpectrumColorSlider {
	c.tabIndex = tabIndex
	return c
}

// Value sets the value property of the SpectrumColorSlider
func (c *SpectrumColorSlider) Value(value float64) *SpectrumColorSlider {
	c.value = value
	return c
}

// Vertical sets whether the color slider is displayed vertically
func (c *SpectrumColorSlider) Vertical(vertical bool) *SpectrumColorSlider {
	c.vertical = vertical
	return c
}

// OnChange sets the change event handler
func (c *SpectrumColorSlider) OnChange(handler app.EventHandler) *SpectrumColorSlider {
	c.onChange = handler
	return c
}

// OnInput sets the input event handler
func (c *SpectrumColorSlider) OnInput(handler app.EventHandler) *SpectrumColorSlider {
	c.onInput = handler
	return c
}

// Render renders the SpectrumColorSlider component.
func (c *SpectrumColorSlider) Render() app.UI {
	colorSlider := app.Elem("sp-color-slider").
		Attr("label", c.label).
		Attr("step", c.step)

	if c.color != "" {
		colorSlider = colorSlider.Attr("color", c.color)
	}
	if c.disabled {
		colorSlider = colorSlider.Attr("disabled", true)
	}
	if c.focused {
		colorSlider = colorSlider.Attr("focused", true)
	}
	if c.tabIndex != 0 {
		colorSlider = colorSlider.Attr("tabindex", c.tabIndex)
	}
	if c.value != 0 {
		colorSlider = colorSlider.Attr("value", c.value)
	}
	if c.vertical {
		colorSlider = colorSlider.Attr("vertical", true)
	}

	// Add event handlers
	if c.onChange != nil {
		colorSlider = colorSlider.On("change", c.onChange)
	}
	if c.onInput != nil {
		colorSlider = colorSlider.On("input", c.onInput)
	}

	return colorSlider
}
