package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// HostWithPendingState defines the interface that a host element must implement
// to work with the PendingStateController
type HostWithPendingState interface {
	app.Composer
	IsPending() bool
	IsDisabled() bool
	GetPendingLabel() string
}

// spectrumPendingStateController helps manage the pending state of a component.
// It provides a standardized way to indicate when an element is in a pending state,
// such as during an asynchronous operation.
type spectrumPendingStateController struct {
	PropHost HostWithPendingState
}

// PendingStateController creates a new spectrumPendingStateController
func PendingStateController(host HostWithPendingState) *spectrumPendingStateController {
	return &spectrumPendingStateController{
		PropHost: host,
	}
}

// RenderPendingState returns a UI element that represents the pending state
// This would typically be a progress circle or similar loading indicator
func (c *spectrumPendingStateController) RenderPendingState() app.UI {
	if !c.PropHost.IsPending() {
		return nil
	}

	// Create a container for the pending state UI
	container := app.Div().
		Class("pending-state-container").
		Style("position", "absolute").
		Style("top", "0").
		Style("left", "0").
		Style("width", "100%").
		Style("height", "100%").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("z-index", "1")

	// Create a progress circle element
	progressCircle := app.Elem("sp-progress-circle").
		Attr("indeterminate", true).
		Attr("size", "m").
		Attr("label", c.PropHost.GetPendingLabel())

	return container.Body(progressCircle)
}

// UpdateAriaLabel updates the aria-label of the host element based on its pending state
func (c *spectrumPendingStateController) UpdateAriaLabel(currentLabel string) string {
	if c.PropHost.IsPending() {
		return c.PropHost.GetPendingLabel()
	}
	return currentLabel
}

// ShouldDisableControls returns whether the controls should be disabled
// based on the pending or disabled state
func (c *spectrumPendingStateController) ShouldDisableControls() bool {
	return c.PropHost.IsPending() || c.PropHost.IsDisabled()
}
