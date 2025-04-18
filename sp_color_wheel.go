// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorWheel represents an sp-color-wheel component
type spectrumColorWheel struct {
	app.Compo
	*styler[*spectrumColorWheel]
	*classer[*spectrumColorWheel]
	*ider[*spectrumColorWheel]

	// Properties
	// The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
	PropColor string
	// Disable this control. It will not receive focus or events
	PropDisabled bool
	// Indicates whether the wheel has focus
	PropFocused bool
	// Accessible label for the wheel
	PropLabel string
	// The increment amount for keyboard navigation
	PropStep float64
	// The tab index to apply to this control
	PropTabindex float64
	// The current value of the wheel (hue)
	PropValue float64

	// Content slots
	PropGradientSlot app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// IColorWheel is the interface for sp-color-wheel component methods
type IColorWheel interface {
	app.UI
	Styler[IColorWheel]
	Classer[IColorWheel]
	Ider[IColorWheel]
	Color(string) IColorWheel
	Disabled(bool) IColorWheel
	SetDisabled() IColorWheel
	Focused(bool) IColorWheel
	SetFocused() IColorWheel
	Label(string) IColorWheel
	Step(float64) IColorWheel
	Tabindex(float64) IColorWheel
	Value(float64) IColorWheel

	Gradient(app.UI) IColorWheel

	OnChange(app.EventHandler) IColorWheel
	OnInput(app.EventHandler) IColorWheel
}

// ColorWheel An sp-color-wheel lets users visually change an individual channel of a color on a circular track.
func ColorWheel() IColorWheel {
	element := &spectrumColorWheel{
		PropDisabled: false,
		PropFocused:  false,
		PropLabel:    "hue",
		PropStep:     1,
		PropTabindex: 0,
		PropValue:    0,
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Color The current color value in various supported formats (Hex, HSV, HSL, RGB, color strings)
func (c *spectrumColorWheel) Color(color string) IColorWheel {
	c.PropColor = color
	return c
}

// Disabled Disable this control. It will not receive focus or events
func (c *spectrumColorWheel) Disabled(disabled bool) IColorWheel {
	c.PropDisabled = disabled
	return c
}

func (c *spectrumColorWheel) SetDisabled() IColorWheel {
	return c.Disabled(true)
}

// Focused Indicates whether the wheel has focus
func (c *spectrumColorWheel) Focused(focused bool) IColorWheel {
	c.PropFocused = focused
	return c
}

func (c *spectrumColorWheel) SetFocused() IColorWheel {
	return c.Focused(true)
}

// Label Accessible label for the wheel
func (c *spectrumColorWheel) Label(label string) IColorWheel {
	c.PropLabel = label
	return c
}

// Step The increment amount for keyboard navigation
func (c *spectrumColorWheel) Step(step float64) IColorWheel {
	c.PropStep = step
	return c
}

// Tabindex The tab index to apply to this control
func (c *spectrumColorWheel) Tabindex(tabindex float64) IColorWheel {
	c.PropTabindex = tabindex
	return c
}

// Value The current value of the wheel (hue)
func (c *spectrumColorWheel) Value(value float64) IColorWheel {
	c.PropValue = value
	return c
}

// A custom gradient visually outlining the available color values
func (c *spectrumColorWheel) Gradient(content app.UI) IColorWheel {
	c.PropGradientSlot = content

	return c
}

// An alteration to the value of the Color Wheel has been committed by the user
func (c *spectrumColorWheel) OnChange(handler app.EventHandler) IColorWheel {
	c.PropOnChange = handler

	return c
}

// The value of the Color Wheel has changed
func (c *spectrumColorWheel) OnInput(handler app.EventHandler) IColorWheel {
	c.PropOnInput = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumColorWheel) Style(key, format string, values ...any) IColorWheel {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumColorWheel) Styles(styles map[string]string) IColorWheel {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumColorWheel) Class(class string) IColorWheel {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumColorWheel) Classes(classes ...string) IColorWheel {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumColorWheel) Id(id string) IColorWheel {
	return c.ider.Id(id)
}

// Render renders the sp-color-wheel component
func (c *spectrumColorWheel) Render() app.UI {
	element := app.Elem("sp-color-wheel")

	// Set attributes
	if c.PropColor != "" {
		element = element.Attr("color", c.PropColor)
	}
	if c.PropDisabled {
		element = element.Attr("disabled", true)
	}
	if c.PropFocused {
		element = element.Attr("focused", true)
	}
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}
	if c.PropStep != 0 {
		element = element.Attr("step", c.PropStep)
	}
	if c.PropTabindex != 0 {
		element = element.Attr("tabindex", c.PropTabindex)
	}
	if c.PropValue != 0 {
		element = element.Attr("value", c.PropValue)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		element = element.On("change", c.PropOnChange)
	}
	if c.PropOnInput != nil {
		element = element.On("input", c.PropOnInput)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add gradient slot
	if c.PropGradientSlot != nil {
		slotElem := c.PropGradientSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("gradient")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "gradient").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
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
