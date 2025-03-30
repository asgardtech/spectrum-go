package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumSidenavItem represents an sp-sidenav-item component
type spectrumSidenavItem struct {
    app.Compo
    *styler[*spectrumSidenavItem]

    // Properties
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Causes the browser to treat the linked URL as a download
    PropDownload string
    // Whether the child items are visible
    PropExpanded bool
    // The URL that the hyperlink points to
    PropHref string
    // An accessible label describing the component. Applied to aria-label but not visually rendered
    PropLabel string
    // How much of the referrer to send when following the link
    PropReferrerpolicy string
    // The relationship of the linked URL as space-separated link types
    PropRel string
    // Whether this item is currently selected
    PropSelected bool
    // The tab index to apply to this control
    PropTabindex float64
    // Where to display the linked URL (e.g., '_blank', '_self')
    PropTarget string
    // The value of this item when selected
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


    // Event handlers
    PropOnClick app.EventHandler
}

// SidenavItem creates a new sp-sidenav-item component
func SidenavItem() *spectrumSidenavItem {
    element := &spectrumSidenavItem{
        PropDisabled: false,
        PropDownload: "",
        PropExpanded: false,
        PropHref: "",
        PropLabel: "",
        PropReferrerpolicy: "",
        PropRel: "",
        PropSelected: false,
        PropTabindex: 0,
        PropTarget: "",
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control. It will not receive focus or events
func (c *spectrumSidenavItem) Disabled(disabled bool) *spectrumSidenavItem {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSidenavItem) SetDisabled() *spectrumSidenavItem {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download
func (c *spectrumSidenavItem) Download(download string) *spectrumSidenavItem {
    c.PropDownload = download
    return c
}

// Whether the child items are visible
func (c *spectrumSidenavItem) Expanded(expanded bool) *spectrumSidenavItem {
    c.PropExpanded = expanded
    return c
}

func (c *spectrumSidenavItem) SetExpanded() *spectrumSidenavItem {
    return c.Expanded(true)
}

// The URL that the hyperlink points to
func (c *spectrumSidenavItem) Href(href string) *spectrumSidenavItem {
    c.PropHref = href
    return c
}

// An accessible label describing the component. Applied to aria-label but not visually rendered
func (c *spectrumSidenavItem) Label(label string) *spectrumSidenavItem {
    c.PropLabel = label
    return c
}

// How much of the referrer to send when following the link
func (c *spectrumSidenavItem) Referrerpolicy(referrerpolicy string) *spectrumSidenavItem {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

// The relationship of the linked URL as space-separated link types
func (c *spectrumSidenavItem) Rel(rel string) *spectrumSidenavItem {
    c.PropRel = rel
    return c
}

// Whether this item is currently selected
func (c *spectrumSidenavItem) Selected(selected bool) *spectrumSidenavItem {
    c.PropSelected = selected
    return c
}

func (c *spectrumSidenavItem) SetSelected() *spectrumSidenavItem {
    return c.Selected(true)
}

// The tab index to apply to this control
func (c *spectrumSidenavItem) Tabindex(tabIndex float64) *spectrumSidenavItem {
    c.PropTabindex = tabIndex
    return c
}

// Where to display the linked URL (e.g., '_blank', '_self')
func (c *spectrumSidenavItem) Target(target string) *spectrumSidenavItem {
    c.PropTarget = target
    return c
}

// The value of this item when selected
func (c *spectrumSidenavItem) Value(value string) *spectrumSidenavItem {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumSidenavItem) Text(text string) *spectrumSidenavItem {
    c.PropText = text
    return c
}


// Optional icon to display with the item
func (c *spectrumSidenavItem) Icon(content app.UI) *spectrumSidenavItem {
    c.PropIconSlot = content

    return c
}



// Fired when the item is clicked
func (c *spectrumSidenavItem) OnClick(handler app.EventHandler) *spectrumSidenavItem {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-sidenav-item component
func (c *spectrumSidenavItem) Render() app.UI {
    element := app.Elem("sp-sidenav-item")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropExpanded {
        element = element.Attr("expanded", true)
    }
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropReferrerpolicy != "" {
        element = element.Attr("referrerpolicy", c.PropReferrerpolicy)
    }
    if c.PropRel != "" {
        element = element.Attr("rel", c.PropRel)
    }
    if c.PropSelected {
        element = element.Attr("selected", true)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", c.PropTarget)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }

    // Add event handlers
    if c.PropOnClick != nil {
        element = element.On("click", c.PropOnClick)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
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

    element = element.Styles(c.styler.styles)

    return element
} 