package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SidenavVariant represents the When 'multilevel', supports multiple layers of hierarchical navigation items
type SidenavVariant string

// SidenavVariant values
const (
    SidenavVariantMultilevel SidenavVariant = "multilevel"
)

// spectrumSidenav represents an sp-sidenav component
type spectrumSidenav struct {
    app.Compo
    *styler[*spectrumSidenav]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // An accessible label describing the component to distinguish it from other navigation for screen reader users
    PropLabel string
    // When true, presents child items with a single tab-stop navigable with arrow keys
    PropManagetabindex bool
    // The tab index to apply to this control
    PropTabindex float64
    // The value of the currently selected item
    PropValue string
    // When 'multilevel', supports multiple layers of hierarchical navigation items
    PropVariant SidenavVariant

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// Sidenav creates a new sp-sidenav component
func Sidenav() *spectrumSidenav {
    element := &spectrumSidenav{
        PropDisabled: false,
        PropLabel: "",
        PropManagetabindex: false,
        PropTabindex: 0,
        PropValue: "",
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumSidenav) Disabled(disabled bool) *spectrumSidenav {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSidenav) SetDisabled() *spectrumSidenav {
    return c.Disabled(true)
}

// An accessible label describing the component to distinguish it from other navigation for screen reader users
func (c *spectrumSidenav) Label(label string) *spectrumSidenav {
    c.PropLabel = label
    return c
}

// When true, presents child items with a single tab-stop navigable with arrow keys
func (c *spectrumSidenav) Managetabindex(manageTabIndex bool) *spectrumSidenav {
    c.PropManagetabindex = manageTabIndex
    return c
}

func (c *spectrumSidenav) SetManagetabindex() *spectrumSidenav {
    return c.Managetabindex(true)
}

// The tab index to apply to this control
func (c *spectrumSidenav) Tabindex(tabIndex float64) *spectrumSidenav {
    c.PropTabindex = tabIndex
    return c
}

// The value of the currently selected item
func (c *spectrumSidenav) Value(value string) *spectrumSidenav {
    c.PropValue = value
    return c
}

// When 'multilevel', supports multiple layers of hierarchical navigation items
func (c *spectrumSidenav) Variant(variant SidenavVariant) *spectrumSidenav {
    c.PropVariant = variant
    return c
}

func (c *spectrumSidenav) VariantMultilevel() *spectrumSidenav {
    return c.Variant(SidenavVariantMultilevel)
}

// Text sets the text content for the default slot
func (c *spectrumSidenav) Text(text string) *spectrumSidenav {
    c.PropText = text
    return c
}




// Announces a change in the value property of the navigation element. Can be canceled via event.preventDefault()
func (c *spectrumSidenav) OnChange(handler app.EventHandler) *spectrumSidenav {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-sidenav component
func (c *spectrumSidenav) Render() app.UI {
    element := app.Elem("sp-sidenav")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropManagetabindex {
        element = element.Attr("manageTabIndex", true)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
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



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 