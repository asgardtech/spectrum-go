package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// OverlayType represents the The type of overlay which determines its behavior
type OverlayType string

// OverlayType values
const (
    OverlayTypeAuto OverlayType = "auto"
    OverlayTypeHint OverlayType = "hint"
    OverlayTypeManual OverlayType = "manual"
    OverlayTypeModal OverlayType = "modal"
    OverlayTypePage OverlayType = "page"
)

// spectrumOverlay represents an sp-overlay component
type spectrumOverlay struct {
    app.Compo
    *styler[*spectrumOverlay]

    // Properties
    // Whether the overlay is currently open
    PropOpen bool
    // Whether the opening of the overlay should be delayed
    PropDelayed bool
    // Distance from the trigger element, can be a single number or an array of two numbers [x, y]
    PropOffset float64
    // Where the overlay should be positioned relative to its trigger element (e.g., 'top', 'bottom', 'right-start')
    PropPlacement string
    // How the overlay should receive focus
    PropReceivesFocus string
    // ID of the element that triggers the overlay, can include interaction type (e.g., 'trigger-id@click')
    PropTrigger string
    // The type of overlay which determines its behavior
    PropType OverlayType

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnSlottableRequest app.EventHandler
    PropOnSpClosed app.EventHandler
    PropOnSpOpened app.EventHandler
}

// Overlay creates a new sp-overlay component
func Overlay() *spectrumOverlay {
    element := &spectrumOverlay{
        PropOpen: false,
        PropDelayed: false,
        PropOffset: 0,
        PropReceivesFocus: "auto",
        PropType: OverlayTypeAuto,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the overlay is currently open
func (c *spectrumOverlay) Open(open bool) *spectrumOverlay {
    c.PropOpen = open
    return c
}

func (c *spectrumOverlay) SetOpen() *spectrumOverlay {
    return c.Open(true)
}

// Whether the opening of the overlay should be delayed
func (c *spectrumOverlay) Delayed(delayed bool) *spectrumOverlay {
    c.PropDelayed = delayed
    return c
}

func (c *spectrumOverlay) SetDelayed() *spectrumOverlay {
    return c.Delayed(true)
}

// Distance from the trigger element, can be a single number or an array of two numbers [x, y]
func (c *spectrumOverlay) Offset(offset float64) *spectrumOverlay {
    c.PropOffset = offset
    return c
}

// Where the overlay should be positioned relative to its trigger element (e.g., 'top', 'bottom', 'right-start')
func (c *spectrumOverlay) Placement(placement string) *spectrumOverlay {
    c.PropPlacement = placement
    return c
}

// How the overlay should receive focus
func (c *spectrumOverlay) ReceivesFocus(receivesFocus string) *spectrumOverlay {
    c.PropReceivesFocus = receivesFocus
    return c
}

// ID of the element that triggers the overlay, can include interaction type (e.g., 'trigger-id@click')
func (c *spectrumOverlay) Trigger(trigger string) *spectrumOverlay {
    c.PropTrigger = trigger
    return c
}

// The type of overlay which determines its behavior
func (c *spectrumOverlay) Type(typeValue OverlayType) *spectrumOverlay {
    c.PropType = typeValue
    return c
}

func (c *spectrumOverlay) TypeAuto() *spectrumOverlay {
    return c.Type(OverlayTypeAuto)
}
func (c *spectrumOverlay) TypeHint() *spectrumOverlay {
    return c.Type(OverlayTypeHint)
}
func (c *spectrumOverlay) TypeManual() *spectrumOverlay {
    return c.Type(OverlayTypeManual)
}
func (c *spectrumOverlay) TypeModal() *spectrumOverlay {
    return c.Type(OverlayTypeModal)
}
func (c *spectrumOverlay) TypePage() *spectrumOverlay {
    return c.Type(OverlayTypePage)
}

// Text sets the text content for the default slot
func (c *spectrumOverlay) Text(text string) *spectrumOverlay {
    c.PropText = text
    return c
}




// Requests to add or remove slottable content
func (c *spectrumOverlay) OnSlottableRequest(handler app.EventHandler) *spectrumOverlay {
    c.PropOnSlottableRequest = handler

    return c
}

// Announces that an overlay has completed any exit animations
func (c *spectrumOverlay) OnSpClosed(handler app.EventHandler) *spectrumOverlay {
    c.PropOnSpClosed = handler

    return c
}

// Announces that an overlay has completed any entry animations
func (c *spectrumOverlay) OnSpOpened(handler app.EventHandler) *spectrumOverlay {
    c.PropOnSpOpened = handler

    return c
}


// Render renders the sp-overlay component
func (c *spectrumOverlay) Render() app.UI {
    element := app.Elem("sp-overlay")

    // Set attributes
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropDelayed {
        element = element.Attr("delayed", true)
    }
    if c.PropOffset != 0 {
        element = element.Attr("offset", c.PropOffset)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", c.PropPlacement)
    }
    if c.PropReceivesFocus != "" {
        element = element.Attr("receives-focus", c.PropReceivesFocus)
    }
    if c.PropTrigger != "" {
        element = element.Attr("trigger", c.PropTrigger)
    }
    if c.PropType != "" {
        element = element.Attr("type", string(c.PropType))
    }

    // Add event handlers
    if c.PropOnSlottableRequest != nil {
        element = element.On("slottable-request", c.PropOnSlottableRequest)
    }
    if c.PropOnSpClosed != nil {
        element = element.On("sp-closed", c.PropOnSpClosed)
    }
    if c.PropOnSpOpened != nil {
        element = element.On("sp-opened", c.PropOnSpOpened)
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