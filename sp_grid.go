// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumGrid represents an sp-grid component
type spectrumGrid struct {
	app.Compo
	*styler[*spectrumGrid]
	*classer[*spectrumGrid]
	*ider[*spectrumGrid]

	// Properties
	// CSS selector that identifies which part of each item can receive focus
	PropFocusableselector string
	// Space between grid items (e.g. '10px')
	PropGap string
	// Size of each grid item as {width: number, height: number}
	PropItemsize any
	// Array of data objects to render in the grid
	PropItems []string
	// Padding around the grid (e.g. '10px')
	PropPadding string
	// Array of currently selected items
	PropSelected []string

	// Event handlers
	PropOnChange app.EventHandler
}

// IGrid is the interface for sp-grid component methods
type IGrid interface {
	app.UI
	Styler[IGrid]
	Classer[IGrid]
	Ider[IGrid]
	Focusableselector(string) IGrid
	Gap(string) IGrid
	Itemsize(any) IGrid
	Items([]string) IGrid
	Padding(string) IGrid
	Selected([]string) IGrid

	OnChange(app.EventHandler) IGrid
}

// Grid An sp-grid element displays a virtualized grid of elements built from its items array, applied to a supplied renderItem method. It extends lit-virtualizer and provides roving tabindex for keyboard navigation.
func Grid() IGrid {
	element := &spectrumGrid{
		PropGap:      "0",
		PropItemsize: map[string]interface{}{},
		PropItems:    []string{},
		PropSelected: []string{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Focusableselector CSS selector that identifies which part of each item can receive focus
func (c *spectrumGrid) Focusableselector(focusableSelector string) IGrid {
	c.PropFocusableselector = focusableSelector
	return c
}

// Gap Space between grid items (e.g. '10px')
func (c *spectrumGrid) Gap(gap string) IGrid {
	c.PropGap = gap
	return c
}

// Itemsize Size of each grid item as {width: number, height: number}
func (c *spectrumGrid) Itemsize(itemSize any) IGrid {
	c.PropItemsize = itemSize
	return c
}

// Items Array of data objects to render in the grid
func (c *spectrumGrid) Items(items []string) IGrid {
	c.PropItems = items
	return c
}

// Padding Padding around the grid (e.g. '10px')
func (c *spectrumGrid) Padding(padding string) IGrid {
	c.PropPadding = padding
	return c
}

// Selected Array of currently selected items
func (c *spectrumGrid) Selected(selected []string) IGrid {
	c.PropSelected = selected
	return c
}

// Announces that the value of selected has changed
func (c *spectrumGrid) OnChange(handler app.EventHandler) IGrid {
	c.PropOnChange = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumGrid) Style(key, format string, values ...any) IGrid {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumGrid) Styles(styles map[string]string) IGrid {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumGrid) Class(class string) IGrid {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumGrid) Classes(classes ...string) IGrid {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumGrid) Id(id string) IGrid {
	return c.ider.Id(id)
}

// Render renders the sp-grid component
func (c *spectrumGrid) Render() app.UI {
	element := app.Elem("sp-grid")

	// Set attributes
	if c.PropFocusableselector != "" {
		element = element.Attr("focusableSelector", c.PropFocusableselector)
	}
	if c.PropGap != "" {
		element = element.Attr("gap", c.PropGap)
	}
	if c.PropItemsize != nil {
		element = element.Attr("itemSize", c.PropItemsize)
	}
	if len(c.PropItems) > 0 {
		element = element.Attr("items", c.PropItems)
	}
	if c.PropPadding != "" {
		element = element.Attr("padding", c.PropPadding)
	}
	if len(c.PropSelected) > 0 {
		element = element.Attr("selected", c.PropSelected)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		element = element.On("change", c.PropOnChange)
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
