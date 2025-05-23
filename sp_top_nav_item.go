// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTopNavItem represents an sp-top-nav-item component
type spectrumTopNavItem struct {
	app.Compo
	*styler[*spectrumTopNavItem]
	*classer[*spectrumTopNavItem]
	*ider[*spectrumTopNavItem]

	// Properties
	// URL that the navigation item points to
	PropHref string
	// Accessible label for the navigation item
	PropLabel string
	// Whether the navigation item is currently selected
	PropSelected bool
	// Specifies where to open the linked document
	PropTarget string

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropIconSlot app.UI

	// Event handlers
	PropOnClick app.EventHandler
}

// ITopNavItem is the interface for sp-top-nav-item component methods
type ITopNavItem interface {
	app.UI
	Styler[ITopNavItem]
	Classer[ITopNavItem]
	Ider[ITopNavItem]
	Href(string) ITopNavItem
	Label(string) ITopNavItem
	Selected(bool) ITopNavItem
	SetSelected() ITopNavItem
	Target(string) ITopNavItem

	Body(...app.UI) ITopNavItem
	AddToBody(app.UI) ITopNavItem
	Text(string) ITopNavItem

	Icon(app.UI) ITopNavItem

	OnClick(app.EventHandler) ITopNavItem
}

// TopNavItem A navigation item used within sp-top-nav to represent a destination in site navigation.
func TopNavItem() ITopNavItem {
	element := &spectrumTopNavItem{
		PropLabel:    "",
		PropSelected: false,
		PropBody:     []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Href URL that the navigation item points to
func (c *spectrumTopNavItem) Href(href string) ITopNavItem {
	c.PropHref = href
	return c
}

// Label Accessible label for the navigation item
func (c *spectrumTopNavItem) Label(label string) ITopNavItem {
	c.PropLabel = label
	return c
}

// Selected Whether the navigation item is currently selected
func (c *spectrumTopNavItem) Selected(selected bool) ITopNavItem {
	c.PropSelected = selected
	return c
}

func (c *spectrumTopNavItem) SetSelected() ITopNavItem {
	return c.Selected(true)
}

// Target Specifies where to open the linked document
func (c *spectrumTopNavItem) Target(target string) ITopNavItem {
	c.PropTarget = target
	return c
}

// Body sets the content for the default slot
func (c *spectrumTopNavItem) Body(elements ...app.UI) ITopNavItem {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumTopNavItem) AddToBody(element app.UI) ITopNavItem {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumTopNavItem) Text(text string) ITopNavItem {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Icon to display with the navigation item
func (c *spectrumTopNavItem) Icon(content app.UI) ITopNavItem {
	c.PropIconSlot = content

	return c
}

// Triggered when the navigation item is clicked
func (c *spectrumTopNavItem) OnClick(handler app.EventHandler) ITopNavItem {
	c.PropOnClick = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumTopNavItem) Style(key, format string, values ...any) ITopNavItem {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumTopNavItem) Styles(styles map[string]string) ITopNavItem {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumTopNavItem) Class(class string) ITopNavItem {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumTopNavItem) Classes(classes ...string) ITopNavItem {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumTopNavItem) Id(id string) ITopNavItem {
	return c.ider.Id(id)
}

// Render renders the sp-top-nav-item component
func (c *spectrumTopNavItem) Render() app.UI {
	element := app.Elem("sp-top-nav-item")

	// Set attributes
	if c.PropHref != "" {
		element = element.Attr("href", c.PropHref)
	}
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}
	if c.PropSelected {
		element = element.Attr("selected", true)
	}
	if c.PropTarget != "" {
		element = element.Attr("target", c.PropTarget)
	}

	// Add event handlers
	if c.PropOnClick != nil {
		element = element.On("click", c.PropOnClick)
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
