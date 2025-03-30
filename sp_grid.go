package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumGrid represents an sp-grid component
type spectrumGrid struct {
    app.Compo
    *styler[*spectrumGrid]

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

// Grid creates a new sp-grid component
func Grid() *spectrumGrid {
    element := &spectrumGrid{
        PropGap: "0",
        PropItemsize: map[string]interface{}{},
        PropItems: []string{},
        PropSelected: []string{},
    }

    element.styler = newStyler(element)

    return element
}

// CSS selector that identifies which part of each item can receive focus
func (c *spectrumGrid) Focusableselector(focusableSelector string) *spectrumGrid {
    c.PropFocusableselector = focusableSelector
    return c
}

// Space between grid items (e.g. '10px')
func (c *spectrumGrid) Gap(gap string) *spectrumGrid {
    c.PropGap = gap
    return c
}

// Size of each grid item as {width: number, height: number}
func (c *spectrumGrid) Itemsize(itemSize any) *spectrumGrid {
    c.PropItemsize = itemSize
    return c
}

// Array of data objects to render in the grid
func (c *spectrumGrid) Items(items []string) *spectrumGrid {
    c.PropItems = items
    return c
}

// Padding around the grid (e.g. '10px')
func (c *spectrumGrid) Padding(padding string) *spectrumGrid {
    c.PropPadding = padding
    return c
}

// Array of currently selected items
func (c *spectrumGrid) Selected(selected []string) *spectrumGrid {
    c.PropSelected = selected
    return c
}





// Announces that the value of selected has changed
func (c *spectrumGrid) OnChange(handler app.EventHandler) *spectrumGrid {
    c.PropOnChange = handler

    return c
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


    element = element.Styles(c.styler.styles)

    return element
} 