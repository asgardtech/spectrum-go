package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumSplitView represents an sp-split-view component
type spectrumSplitView struct {
    app.Compo
    *styler[*spectrumSplitView]

    // Properties
    // Whether panels can be collapsed completely
    PropCollapsible bool
    // Accessible label for the resizable divider
    PropLabel string
    // The maximum size of the primary pane (in pixels)
    PropPrimaryMax float64
    // The minimum size of the primary pane (in pixels)
    PropPrimaryMin float64
    // The initial size of the primary pane (in pixels or percentage)
    PropPrimarySize string
    // Whether the split view can be resized by the user
    PropResizable bool
    // The maximum size of the secondary pane (in pixels)
    PropSecondaryMax float64
    // The minimum size of the secondary pane (in pixels)
    PropSecondaryMin float64
    // The current splitter position (in pixels)
    PropSplitterPos float64
    // Whether the split view is arranged vertically
    PropVertical bool


    // Content slots
    PropPrimaryPanelSlot app.UI
    PropSecondaryPanelSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
}

// SplitView creates a new sp-split-view component
func SplitView() *spectrumSplitView {
    element := &spectrumSplitView{
        PropCollapsible: false,
        PropLabel: "Resize the panels",
        PropPrimaryMax: 0,
        PropPrimaryMin: 0,
        PropResizable: false,
        PropSecondaryMax: 0,
        PropSecondaryMin: 0,
        PropSplitterPos: 0,
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Whether panels can be collapsed completely
func (c *spectrumSplitView) Collapsible(collapsible bool) *spectrumSplitView {
    c.PropCollapsible = collapsible
    return c
}

func (c *spectrumSplitView) SetCollapsible() *spectrumSplitView {
    return c.Collapsible(true)
}

// Accessible label for the resizable divider
func (c *spectrumSplitView) Label(label string) *spectrumSplitView {
    c.PropLabel = label
    return c
}

// The maximum size of the primary pane (in pixels)
func (c *spectrumSplitView) PrimaryMax(primaryMax float64) *spectrumSplitView {
    c.PropPrimaryMax = primaryMax
    return c
}

// The minimum size of the primary pane (in pixels)
func (c *spectrumSplitView) PrimaryMin(primaryMin float64) *spectrumSplitView {
    c.PropPrimaryMin = primaryMin
    return c
}

// The initial size of the primary pane (in pixels or percentage)
func (c *spectrumSplitView) PrimarySize(primarySize string) *spectrumSplitView {
    c.PropPrimarySize = primarySize
    return c
}

// Whether the split view can be resized by the user
func (c *spectrumSplitView) Resizable(resizable bool) *spectrumSplitView {
    c.PropResizable = resizable
    return c
}

func (c *spectrumSplitView) SetResizable() *spectrumSplitView {
    return c.Resizable(true)
}

// The maximum size of the secondary pane (in pixels)
func (c *spectrumSplitView) SecondaryMax(secondaryMax float64) *spectrumSplitView {
    c.PropSecondaryMax = secondaryMax
    return c
}

// The minimum size of the secondary pane (in pixels)
func (c *spectrumSplitView) SecondaryMin(secondaryMin float64) *spectrumSplitView {
    c.PropSecondaryMin = secondaryMin
    return c
}

// The current splitter position (in pixels)
func (c *spectrumSplitView) SplitterPos(splitterPos float64) *spectrumSplitView {
    c.PropSplitterPos = splitterPos
    return c
}

// Whether the split view is arranged vertically
func (c *spectrumSplitView) Vertical(vertical bool) *spectrumSplitView {
    c.PropVertical = vertical
    return c
}

func (c *spectrumSplitView) SetVertical() *spectrumSplitView {
    return c.Vertical(true)
}



// Two sibling elements to be sized by the element attributes
func (c *spectrumSplitView) PrimaryPanel(content app.UI) *spectrumSplitView {
    c.PropPrimaryPanelSlot = content

    return c
}

// Two sibling elements to be sized by the element attributes
func (c *spectrumSplitView) SecondaryPanel(content app.UI) *spectrumSplitView {
    c.PropSecondaryPanelSlot = content

    return c
}



// Announces the new position of the splitter
func (c *spectrumSplitView) OnChange(handler app.EventHandler) *spectrumSplitView {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-split-view component
func (c *spectrumSplitView) Render() app.UI {
    element := app.Elem("sp-split-view")

    // Set attributes
    if c.PropCollapsible {
        element = element.Attr("collapsible", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropPrimaryMax != 0 {
        element = element.Attr("primary-max", c.PropPrimaryMax)
    }
    if c.PropPrimaryMin != 0 {
        element = element.Attr("primary-min", c.PropPrimaryMin)
    }
    if c.PropPrimarySize != "" {
        element = element.Attr("primary-size", c.PropPrimarySize)
    }
    if c.PropResizable {
        element = element.Attr("resizable", true)
    }
    if c.PropSecondaryMax != 0 {
        element = element.Attr("secondary-max", c.PropSecondaryMax)
    }
    if c.PropSecondaryMin != 0 {
        element = element.Attr("secondary-min", c.PropSecondaryMin)
    }
    if c.PropSplitterPos != 0 {
        element = element.Attr("splitter-pos", c.PropSplitterPos)
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }

    // Add slots and children
    slotElements := []app.UI{}


    // Add primary-panel slot
    if c.PropPrimaryPanelSlot != nil {
        slotElem := c.PropPrimaryPanelSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("primary-panel")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "primary-panel").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add secondary-panel slot
    if c.PropSecondaryPanelSlot != nil {
        slotElem := c.PropSecondaryPanelSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("secondary-panel")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "secondary-panel").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 