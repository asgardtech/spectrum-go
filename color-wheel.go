package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorWheel represents the sp-color-wheel component
// It lets users visually change an individual channel of a color on a circular track.
type spectrumColorWheel struct {
	app.Compo

	// Properties
	PropColor    string
	PropDisabled bool
	PropFocused  bool
	PropLabel    string
	PropStep     float64
	PropTabIndex int
	PropValue    float64

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// ColorWheel creates a new spectrumColorWheel instance.
func ColorWheel() *spectrumColorWheel {
	return &spectrumColorWheel{
		PropLabel: "hue",
		PropStep:  1,
	}
}

// Color sets the color property of the spectrumColorWheel
func (c *spectrumColorWheel) Color(color string) *spectrumColorWheel {
	c.PropColor = color
	return c
}

// Disabled sets whether the color wheel is disabled
func (c *spectrumColorWheel) Disabled(disabled bool) *spectrumColorWheel {
	c.PropDisabled = disabled
	return c
}

// Focused sets whether the color wheel is focused
func (c *spectrumColorWheel) Focused(focused bool) *spectrumColorWheel {
	c.PropFocused = focused
	return c
}

// Label sets the label property of the spectrumColorWheel
func (c *spectrumColorWheel) Label(label string) *spectrumColorWheel {
	c.PropLabel = label
	return c
}

// Step sets the step property of the spectrumColorWheel
func (c *spectrumColorWheel) Step(step float64) *spectrumColorWheel {
	c.PropStep = step
	return c
}

// TabIndex sets the tabIndex property of the spectrumColorWheel
func (c *spectrumColorWheel) TabIndex(tabIndex int) *spectrumColorWheel {
	c.PropTabIndex = tabIndex
	return c
}

// Value sets the value property of the spectrumColorWheel
func (c *spectrumColorWheel) Value(value float64) *spectrumColorWheel {
	c.PropValue = value
	return c
}

// OnChange sets the change event handler
func (c *spectrumColorWheel) OnChange(handler app.EventHandler) *spectrumColorWheel {
	c.PropOnChange = handler
	return c
}

// OnInput sets the input event handler
func (c *spectrumColorWheel) OnInput(handler app.EventHandler) *spectrumColorWheel {
	c.PropOnInput = handler
	return c
}

// Render renders the spectrumColorWheel component.
func (c *spectrumColorWheel) Render() app.UI {
	colorWheel := app.Elem("sp-color-wheel").
		Attr("label", c.PropLabel).
		Attr("step", c.PropStep)

	if c.PropColor != "" {
		colorWheel = colorWheel.Attr("color", c.PropColor)
	}
	if c.PropDisabled {
		colorWheel = colorWheel.Attr("disabled", true)
	}
	if c.PropFocused {
		colorWheel = colorWheel.Attr("focused", true)
	}
	if c.PropTabIndex != 0 {
		colorWheel = colorWheel.Attr("tabindex", c.PropTabIndex)
	}
	if c.PropValue != 0 {
		colorWheel = colorWheel.Attr("value", c.PropValue)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		colorWheel = colorWheel.On("change", c.PropOnChange)
	}
	if c.PropOnInput != nil {
		colorWheel = colorWheel.On("input", c.PropOnInput)
	}

	return colorWheel
}
