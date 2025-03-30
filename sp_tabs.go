package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TabsDirection represents the Direction of the tabs layout
type TabsDirection string

// TabsDirection values
const (
    TabsDirectionHorizontal TabsDirection = "horizontal"
    TabsDirectionVertical TabsDirection = "vertical"
    TabsDirectionVerticalRight TabsDirection = "vertical-right"
)
// TabsSize represents the Size of the tabs
type TabsSize string

// TabsSize values
const (
    TabsSizeS TabsSize = "s"
    TabsSizeM TabsSize = "m"
    TabsSizeL TabsSize = "l"
    TabsSizeXl TabsSize = "xl"
)

// spectrumTabs represents an sp-tabs component
type spectrumTabs struct {
    app.Compo
    *styler[*spectrumTabs]

    // Properties
    // Whether to activate a tab on keyboard focus or not
    PropAuto bool
    // The tab items are displayed closer together
    PropCompact bool
    // Direction of the tabs layout
    PropDirection TabsDirection
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Emphasize the selected tab
    PropEmphasized bool
    // Enable scrolling when tabs exceed available space
    PropEnabletabsscroll bool
    // Accessible label for the tabs component
    PropLabel string
    // The tab list is displayed without a border
    PropQuiet bool
    // Value of the currently selected tab
    PropSelected string
    // Size of the tabs
    PropSize TabsSize
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64

    // Text content for the default slot
    PropText string

    // Content slots
    PropTabPanelSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
}

// Tabs creates a new sp-tabs component
func Tabs() *spectrumTabs {
    element := &spectrumTabs{
        PropAuto: false,
        PropCompact: false,
        PropDirection: TabsDirectionHorizontal,
        PropDisabled: false,
        PropEmphasized: false,
        PropEnabletabsscroll: false,
        PropLabel: "",
        PropQuiet: false,
        PropSelected: "",
        PropSize: TabsSizeM,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Whether to activate a tab on keyboard focus or not
func (c *spectrumTabs) Auto(auto bool) *spectrumTabs {
    c.PropAuto = auto
    return c
}

func (c *spectrumTabs) SetAuto() *spectrumTabs {
    return c.Auto(true)
}

// The tab items are displayed closer together
func (c *spectrumTabs) Compact(compact bool) *spectrumTabs {
    c.PropCompact = compact
    return c
}

func (c *spectrumTabs) SetCompact() *spectrumTabs {
    return c.Compact(true)
}

// Direction of the tabs layout
func (c *spectrumTabs) Direction(direction TabsDirection) *spectrumTabs {
    c.PropDirection = direction
    return c
}

func (c *spectrumTabs) DirectionHorizontal() *spectrumTabs {
    return c.Direction(TabsDirectionHorizontal)
}
func (c *spectrumTabs) DirectionVertical() *spectrumTabs {
    return c.Direction(TabsDirectionVertical)
}
func (c *spectrumTabs) DirectionVerticalRight() *spectrumTabs {
    return c.Direction(TabsDirectionVerticalRight)
}
// Disable this control. It will not receive focus or events
func (c *spectrumTabs) Disabled(disabled bool) *spectrumTabs {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumTabs) SetDisabled() *spectrumTabs {
    return c.Disabled(true)
}

// Emphasize the selected tab
func (c *spectrumTabs) Emphasized(emphasized bool) *spectrumTabs {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumTabs) SetEmphasized() *spectrumTabs {
    return c.Emphasized(true)
}

// Enable scrolling when tabs exceed available space
func (c *spectrumTabs) Enabletabsscroll(enableTabsScroll bool) *spectrumTabs {
    c.PropEnabletabsscroll = enableTabsScroll
    return c
}

func (c *spectrumTabs) SetEnabletabsscroll() *spectrumTabs {
    return c.Enabletabsscroll(true)
}

// Accessible label for the tabs component
func (c *spectrumTabs) Label(label string) *spectrumTabs {
    c.PropLabel = label
    return c
}

// The tab list is displayed without a border
func (c *spectrumTabs) Quiet(quiet bool) *spectrumTabs {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumTabs) SetQuiet() *spectrumTabs {
    return c.Quiet(true)
}

// Value of the currently selected tab
func (c *spectrumTabs) Selected(selected string) *spectrumTabs {
    c.PropSelected = selected
    return c
}

// Size of the tabs
func (c *spectrumTabs) Size(size TabsSize) *spectrumTabs {
    c.PropSize = size
    return c
}

func (c *spectrumTabs) SizeS() *spectrumTabs {
    return c.Size(TabsSizeS)
}
func (c *spectrumTabs) SizeM() *spectrumTabs {
    return c.Size(TabsSizeM)
}
func (c *spectrumTabs) SizeL() *spectrumTabs {
    return c.Size(TabsSizeL)
}
func (c *spectrumTabs) SizeXl() *spectrumTabs {
    return c.Size(TabsSizeXl)
}
// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumTabs) Tabindex(tabIndex float64) *spectrumTabs {
    c.PropTabindex = tabIndex
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTabs) Text(text string) *spectrumTabs {
    c.PropText = text
    return c
}


// Tab Panel elements related to the listed Tab elements
func (c *spectrumTabs) TabPanel(content app.UI) *spectrumTabs {
    c.PropTabPanelSlot = content

    return c
}



// The selected Tab child has changed
func (c *spectrumTabs) OnChange(handler app.EventHandler) *spectrumTabs {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-tabs component
func (c *spectrumTabs) Render() app.UI {
    element := app.Elem("sp-tabs")

    // Set attributes
    if c.PropAuto {
        element = element.Attr("auto", true)
    }
    if c.PropCompact {
        element = element.Attr("compact", true)
    }
    if c.PropDirection != "" {
        element = element.Attr("direction", string(c.PropDirection))
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
    }
    if c.PropEnabletabsscroll {
        element = element.Attr("enableTabsScroll", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropSelected != "" {
        element = element.Attr("selected", c.PropSelected)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add tab-panel slot
    if c.PropTabPanelSlot != nil {
        slotElem := c.PropTabPanelSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("tab-panel")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "tab-panel").
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