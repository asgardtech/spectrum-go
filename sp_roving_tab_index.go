package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumRovingTabIndex represents an  component
type spectrumRovingTabIndex struct {
    app.Compo
    *styler[*spectrumRovingTabIndex]

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

// RovingTabIndex creates a new  component
func RovingTabIndex() *spectrumRovingTabIndex {
    element := &spectrumRovingTabIndex{
        PropDirection: "both",
        PropListenerscope: map[string]interface{}{},
    }

    element.styler = newStyler(element)

    return element
}

// Customizes which arrow keys manage element focus
func (c *spectrumRovingTabIndex) Direction(direction string) *spectrumRovingTabIndex {
    c.PropDirection = direction
    return c
}

// Controls which element receives tabindex=0 while focus is outside the container
func (c *spectrumRovingTabIndex) Focusinindex(focusInIndex string) *spectrumRovingTabIndex {
    c.PropFocusinindex = focusInIndex
    return c
}

// Provides the elements that will have their tabindex managed
func (c *spectrumRovingTabIndex) Elements(elements string) *spectrumRovingTabIndex {
    c.PropElements = elements
    return c
}

// Describes the state an element must be in to receive focus
func (c *spectrumRovingTabIndex) Isfocusableelement(isFocusableElement string) *spectrumRovingTabIndex {
    c.PropIsfocusableelement = isFocusableElement
    return c
}

// Outlines which parts of a container's DOM to listen for arrow key presses
func (c *spectrumRovingTabIndex) Listenerscope(listenerScope any) *spectrumRovingTabIndex {
    c.PropListenerscope = listenerScope
    return c
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



    element = element.Styles(c.styler.styles)

    return element
} 