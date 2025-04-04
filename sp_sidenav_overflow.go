// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSidenavOverflow represents an sp-sidenav-overflow component
type spectrumSidenavOverflow struct {
	app.Compo
	*styler[*spectrumSidenavOverflow]
	*classer[*spectrumSidenavOverflow]
	*ider[*spectrumSidenavOverflow]

	// Properties
	// The text label for the overflow menu
	PropLabel string

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropIconSlot app.UI
}

// ISidenavOverflow is the interface for sp-sidenav-overflow component methods
type ISidenavOverflow interface {
	app.UI
	Styler[ISidenavOverflow]
	Classer[ISidenavOverflow]
	Ider[ISidenavOverflow]
	Label(string) ISidenavOverflow

	Body(...app.UI) ISidenavOverflow
	AddToBody(app.UI) ISidenavOverflow
	Text(string) ISidenavOverflow

	Icon(app.UI) ISidenavOverflow
}

// SidenavOverflow A component used to handle overflow of navigation items in a sidenav component. It displays additional items in a submenu when there are too many to display in the main navigation.
func SidenavOverflow() ISidenavOverflow {
	element := &spectrumSidenavOverflow{
		PropLabel: "More",
		PropBody:  []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Label The text label for the overflow menu
func (c *spectrumSidenavOverflow) Label(label string) ISidenavOverflow {
	c.PropLabel = label
	return c
}

// Body sets the content for the default slot
func (c *spectrumSidenavOverflow) Body(elements ...app.UI) ISidenavOverflow {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumSidenavOverflow) AddToBody(element app.UI) ISidenavOverflow {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumSidenavOverflow) Text(text string) ISidenavOverflow {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Icon to display for the overflow menu
func (c *spectrumSidenavOverflow) Icon(content app.UI) ISidenavOverflow {
	c.PropIconSlot = content

	return c
}

// Style sets a style property with a value
func (c *spectrumSidenavOverflow) Style(key, format string, values ...any) ISidenavOverflow {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumSidenavOverflow) Styles(styles map[string]string) ISidenavOverflow {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumSidenavOverflow) Class(class string) ISidenavOverflow {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumSidenavOverflow) Classes(classes ...string) ISidenavOverflow {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumSidenavOverflow) Id(id string) ISidenavOverflow {
	return c.ider.Id(id)
}

// Render renders the sp-sidenav-overflow component
func (c *spectrumSidenavOverflow) Render() app.UI {
	element := app.Elem("sp-sidenav-overflow")

	// Set attributes
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add icon slot
	if c.PropIconSlot != nil {
		slotElem := c.PropIconSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("icon")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "icon").
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
