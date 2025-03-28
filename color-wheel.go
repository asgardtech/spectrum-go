package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumColorWheel represents the sp-color-wheel component
// It lets users visually change an individual channel of a color on a circular track.
type SpectrumColorWheel struct {
	app.Compo

	// Properties
	color    string
	disabled bool
	focused  bool
	label    string
	step     float64
	tabIndex int
	value    float64

	// Event handlers
	onChange app.EventHandler
	onInput  app.EventHandler
}

// ColorWheel creates a new SpectrumColorWheel instance.
func ColorWheel() *SpectrumColorWheel {
	return &SpectrumColorWheel{
		label: "hue",
		step:  1,
	}
}

// Color sets the color property of the SpectrumColorWheel
func (c *SpectrumColorWheel) Color(color string) *SpectrumColorWheel {
	c.color = color
	return c
}

// Disabled sets whether the color wheel is disabled
func (c *SpectrumColorWheel) Disabled(disabled bool) *SpectrumColorWheel {
	c.disabled = disabled
	return c
}

// Focused sets whether the color wheel is focused
func (c *SpectrumColorWheel) Focused(focused bool) *SpectrumColorWheel {
	c.focused = focused
	return c
}

// Label sets the label property of the SpectrumColorWheel
func (c *SpectrumColorWheel) Label(label string) *SpectrumColorWheel {
	c.label = label
	return c
}

// Step sets the step property of the SpectrumColorWheel
func (c *SpectrumColorWheel) Step(step float64) *SpectrumColorWheel {
	c.step = step
	return c
}

// TabIndex sets the tabIndex property of the SpectrumColorWheel
func (c *SpectrumColorWheel) TabIndex(tabIndex int) *SpectrumColorWheel {
	c.tabIndex = tabIndex
	return c
}

// Value sets the value property of the SpectrumColorWheel
func (c *SpectrumColorWheel) Value(value float64) *SpectrumColorWheel {
	c.value = value
	return c
}

// OnChange sets the change event handler
func (c *SpectrumColorWheel) OnChange(handler app.EventHandler) *SpectrumColorWheel {
	c.onChange = handler
	return c
}

// OnInput sets the input event handler
func (c *SpectrumColorWheel) OnInput(handler app.EventHandler) *SpectrumColorWheel {
	c.onInput = handler
	return c
}

// Render renders the SpectrumColorWheel component.
func (c *SpectrumColorWheel) Render() app.UI {
	colorWheel := app.Elem("sp-color-wheel").
		Attr("label", c.label).
		Attr("step", c.step)

	if c.color != "" {
		colorWheel = colorWheel.Attr("color", c.color)
	}
	if c.disabled {
		colorWheel = colorWheel.Attr("disabled", true)
	}
	if c.focused {
		colorWheel = colorWheel.Attr("focused", true)
	}
	if c.tabIndex != 0 {
		colorWheel = colorWheel.Attr("tabindex", c.tabIndex)
	}
	if c.value != 0 {
		colorWheel = colorWheel.Attr("value", c.value)
	}

	// Add event handlers
	if c.onChange != nil {
		colorWheel = colorWheel.On("change", c.onChange)
	}
	if c.onInput != nil {
		colorWheel = colorWheel.On("input", c.onInput)
	}

	return colorWheel
}
