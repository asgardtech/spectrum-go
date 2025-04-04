// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumRovingTabIndex represents an  component
type spectrumRovingTabIndex struct {
	app.Compo
	*styler[*spectrumRovingTabIndex]
	*classer[*spectrumRovingTabIndex]
	*ider[*spectrumRovingTabIndex]

	// Properties
	// Customizes which arrow keys manage element focus
	PropDirection string
	// Controls which element receives tabindex=0 while focus is outside the container
	PropFocusinindex string
	// Provides the elements that will have their tabindex managed
	PropElements string
	// Describes the state an element must be in to receive focus
	PropIsfocusableelement string
	// Outlines which parts of a container's DOM to listen for arrow key presses
	PropListenerscope any
}

// IRovingTabIndex is the interface for  component methods
type IRovingTabIndex interface {
	app.UI
	Styler[IRovingTabIndex]
	Classer[IRovingTabIndex]
	Ider[IRovingTabIndex]
	Direction(string) IRovingTabIndex
	Focusinindex(string) IRovingTabIndex
	Elements(string) IRovingTabIndex
	Isfocusableelement(string) IRovingTabIndex
	Listenerscope(any) IRovingTabIndex
}

// RovingTabIndex Roving tabindex is a pattern whereby multiple focusable elements are represented by a single tabindex=0 element, while the individual elements maintain tabindex=-1 and are accessed via arrow keys. This helps keyboard users tab through a page without stopping on every element in a large collection.
func RovingTabIndex() IRovingTabIndex {
	element := &spectrumRovingTabIndex{
		PropDirection:     "both",
		PropListenerscope: map[string]interface{}{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Direction Customizes which arrow keys manage element focus
func (c *spectrumRovingTabIndex) Direction(direction string) IRovingTabIndex {
	c.PropDirection = direction
	return c
}

// Focusinindex Controls which element receives tabindex=0 while focus is outside the container
func (c *spectrumRovingTabIndex) Focusinindex(focusInIndex string) IRovingTabIndex {
	c.PropFocusinindex = focusInIndex
	return c
}

// Elements Provides the elements that will have their tabindex managed
func (c *spectrumRovingTabIndex) Elements(elements string) IRovingTabIndex {
	c.PropElements = elements
	return c
}

// Isfocusableelement Describes the state an element must be in to receive focus
func (c *spectrumRovingTabIndex) Isfocusableelement(isFocusableElement string) IRovingTabIndex {
	c.PropIsfocusableelement = isFocusableElement
	return c
}

// Listenerscope Outlines which parts of a container's DOM to listen for arrow key presses
func (c *spectrumRovingTabIndex) Listenerscope(listenerScope any) IRovingTabIndex {
	c.PropListenerscope = listenerScope
	return c
}

// Style sets a style property with a value
func (c *spectrumRovingTabIndex) Style(key, format string, values ...any) IRovingTabIndex {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumRovingTabIndex) Styles(styles map[string]string) IRovingTabIndex {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumRovingTabIndex) Class(class string) IRovingTabIndex {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumRovingTabIndex) Classes(classes ...string) IRovingTabIndex {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumRovingTabIndex) Id(id string) IRovingTabIndex {
	return c.ider.Id(id)
}

// Render renders the  component
func (c *spectrumRovingTabIndex) Render() app.UI {
	element := app.Elem("")

	// Set attributes
	if c.PropDirection != "" {
		element = element.Attr("direction", c.PropDirection)
	}
	if c.PropFocusinindex != "" {
		element = element.Attr("focusInIndex", c.PropFocusinindex)
	}
	if c.PropElements != "" {
		element = element.Attr("elements", c.PropElements)
	}
	if c.PropIsfocusableelement != "" {
		element = element.Attr("isFocusableElement", c.PropIsfocusableelement)
	}
	if c.PropListenerscope != nil {
		element = element.Attr("listenerScope", c.PropListenerscope)
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
