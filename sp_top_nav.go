package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumTopNav represents an sp-top-nav component
type spectrumTopNav struct {
    app.Compo
    *styler[*spectrumTopNav]

    // Properties
    // A space separated list of parts of the URL to ignore when matching for the selected Top Nav Item. Supported values: 'hash', 'search'
    PropIgnoreUrlParts string
    // Accessible label for the navigation
    PropLabel string
    // The Top Nav is displayed without a border
    PropQuiet bool
    // The currently selected navigation item
    PropSelected string
    // Style to apply to the selection indicator
    PropSelectionIndicatorStyle string

    // Text content for the default slot
    PropText string

    // Content slots


}

// TopNav creates a new sp-top-nav component
func TopNav() *spectrumTopNav {
    element := &spectrumTopNav{
        PropIgnoreUrlParts: "",
        PropLabel: "",
        PropQuiet: false,
    }

    element.styler = newStyler(element)

    return element
}

// A space separated list of parts of the URL to ignore when matching for the selected Top Nav Item. Supported values: 'hash', 'search'
func (c *spectrumTopNav) IgnoreUrlParts(ignoreUrlParts string) *spectrumTopNav {
    c.PropIgnoreUrlParts = ignoreUrlParts
    return c
}

// Accessible label for the navigation
func (c *spectrumTopNav) Label(label string) *spectrumTopNav {
    c.PropLabel = label
    return c
}

// The Top Nav is displayed without a border
func (c *spectrumTopNav) Quiet(quiet bool) *spectrumTopNav {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumTopNav) SetQuiet() *spectrumTopNav {
    return c.Quiet(true)
}

// The currently selected navigation item
func (c *spectrumTopNav) Selected(selected string) *spectrumTopNav {
    c.PropSelected = selected
    return c
}

// Style to apply to the selection indicator
func (c *spectrumTopNav) SelectionIndicatorStyle(selectionIndicatorStyle string) *spectrumTopNav {
    c.PropSelectionIndicatorStyle = selectionIndicatorStyle
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTopNav) Text(text string) *spectrumTopNav {
    c.PropText = text
    return c
}





// Render renders the sp-top-nav component
func (c *spectrumTopNav) Render() app.UI {
    element := app.Elem("sp-top-nav")

    // Set attributes
    if c.PropIgnoreUrlParts != "" {
        element = element.Attr("ignore-url-parts", c.PropIgnoreUrlParts)
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
    if c.PropSelectionIndicatorStyle != "" {
        element = element.Attr("selection-indicator-style", c.PropSelectionIndicatorStyle)
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