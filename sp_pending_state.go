package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumPendingState represents an  component
type spectrumPendingState struct {
    app.Compo
    *styler[*spectrumPendingState]

    // Properties
    // Whether the items are currently loading
    PropPending bool
    // Whether the host is disabled
    PropDisabled bool
    // Defines a string value that labels the element while it is in pending state
    PropPendingLabel string




}

// PendingState creates a new  component
func PendingState() *spectrumPendingState {
    element := &spectrumPendingState{
        PropPending: false,
        PropDisabled: false,
        PropPendingLabel: "Pending",
    }

    element.styler = newStyler(element)

    return element
}

// Whether the items are currently loading
func (c *spectrumPendingState) Pending(pending bool) *spectrumPendingState {
    c.PropPending = pending
    return c
}

func (c *spectrumPendingState) SetPending() *spectrumPendingState {
    return c.Pending(true)
}

// Whether the host is disabled
func (c *spectrumPendingState) Disabled(disabled bool) *spectrumPendingState {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumPendingState) SetDisabled() *spectrumPendingState {
    return c.Disabled(true)
}

// Defines a string value that labels the element while it is in pending state
func (c *spectrumPendingState) PendingLabel(pendingLabel string) *spectrumPendingState {
    c.PropPendingLabel = pendingLabel
    return c
}






// Render renders the  component
func (c *spectrumPendingState) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropPending {
        element = element.Attr("pending", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropPendingLabel != "" {
        element = element.Attr("pending-label", c.PropPendingLabel)
    }



    element = element.Styles(c.styler.styles)

    return element
} 