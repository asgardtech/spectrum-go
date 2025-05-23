// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumReactiveControllers represents an  component
type spectrumReactiveControllers struct {
	app.Compo
	*styler[*spectrumReactiveControllers]
	*classer[*spectrumReactiveControllers]
	*ider[*spectrumReactiveControllers]

	// Properties

}

// IReactiveControllers is the interface for  component methods
type IReactiveControllers interface {
	app.UI
	Styler[IReactiveControllers]
	Classer[IReactiveControllers]
	Ider[IReactiveControllers]
}

// ReactiveControllers Reactive controllers are a tool for code reuse and composition within Lit, a core dependency of Spectrum Web Components. They can be reused across components to reduce both code complexity and size, and to deliver a consistent user experience.
func ReactiveControllers() IReactiveControllers {
	element := &spectrumReactiveControllers{}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Style sets a style property with a value
func (c *spectrumReactiveControllers) Style(key, format string, values ...any) IReactiveControllers {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumReactiveControllers) Styles(styles map[string]string) IReactiveControllers {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumReactiveControllers) Class(class string) IReactiveControllers {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumReactiveControllers) Classes(classes ...string) IReactiveControllers {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumReactiveControllers) Id(id string) IReactiveControllers {
	return c.ider.Id(id)
}

// Render renders the  component
func (c *spectrumReactiveControllers) Render() app.UI {
	element := app.Elem("")

	// Set attributes

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
