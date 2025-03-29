package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ElementResolverUpdatedSymbol is a sentinel value to indicate when an element reference has been updated
var ElementResolverUpdatedSymbol = "element-resolver-updated"

// spectrumElementResolutionController keeps an active reference to another element in the same DOM tree
// It manages periodic checking of the DOM tree to ensure that the reference it holds is always the first matched element or nil
type spectrumElementResolutionController struct {
	PropHost       app.Composer
	PropElement    interface{}
	PropSelector   string
	observing      bool
	pollingTimerID interface{}
	pollingEnabled bool
}

// ElementResolutionController creates a new spectrumElementResolutionController
func ElementResolutionController(host app.Composer) *spectrumElementResolutionController {
	controller := &spectrumElementResolutionController{
		PropHost: host,
	}

	// Initialize the controller
	return controller
}

// SetSelector sets the CSS selector to use for element resolution and updates the reference
func (c *spectrumElementResolutionController) SetSelector(selector string) *spectrumElementResolutionController {
	c.PropSelector = selector
	c.updateReference()
	return c
}

// Connect starts observing DOM changes if a selector is provided
func (c *spectrumElementResolutionController) Connect() {
	if c.PropSelector != "" && !c.observing {
		c.startObserving()
	}
}

// Disconnect stops observing DOM changes
func (c *spectrumElementResolutionController) Disconnect() {
	if c.observing {
		c.stopObserving()
	}
}

// startObserving sets up a polling mechanism to check for DOM changes
func (c *spectrumElementResolutionController) startObserving() {
	if c.pollingEnabled {
		return
	}

	// Set up polling using JavaScript setInterval
	c.pollingEnabled = true
	c.setupPolling()

	// Make initial reference update
	c.updateReference()
	c.observing = true
}

// setupPolling creates a recurring timer to check for DOM changes
func (c *spectrumElementResolutionController) setupPolling() {
	// Clear any existing timer
	if c.pollingTimerID != nil {
		app.Window().Call("clearInterval", c.pollingTimerID)
		c.pollingTimerID = nil
	}

	// Only set up polling if enabled
	if !c.pollingEnabled {
		return
	}

	// Create a function to check for updates periodically
	checkFunc := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		if c.pollingEnabled {
			c.updateReference()
		}
		return nil
	})

	// Set up interval to check every 500ms
	// This is a compromise between responsiveness and performance
	c.pollingTimerID = app.Window().Call("setInterval", checkFunc, 500)
}

// stopObserving stops the polling mechanism
func (c *spectrumElementResolutionController) stopObserving() {
	c.pollingEnabled = false

	// Clear the polling timer if it exists
	if c.pollingTimerID != nil {
		app.Window().Call("clearInterval", c.pollingTimerID)
		c.pollingTimerID = nil
	}

	c.observing = false
}

// updateReference queries the DOM and updates the element reference
func (c *spectrumElementResolutionController) updateReference() {
	if c.PropSelector == "" {
		c.PropElement = nil
		// Notify the host component that the reference has been updated
		notifyHostOfUpdate(c.PropHost)
		return
	}

	// Query the DOM for the first element matching the selector
	doc := app.Window().Get("document")
	element := doc.Call("querySelector", c.PropSelector)

	// Check if the element has changed
	oldElement := c.PropElement

	if !element.IsNull() {
		// Store the JavaScript DOM element reference
		c.PropElement = element
	} else {
		c.PropElement = nil
	}

	// If the element reference has changed, notify the host
	if oldElement != c.PropElement {
		// Notify the host component that the reference has been updated
		notifyHostOfUpdate(c.PropHost)
	}
}

// notifyHostOfUpdate triggers an update on the host component
func notifyHostOfUpdate(host app.Composer) {
	// In go-app v10, we need to use Update method if the component implements app.Updater
	if updater, ok := host.(interface{ Update() }); ok {
		updater.Update()
	}
}

// GetElementProperty retrieves a property value from the current element reference
func (c *spectrumElementResolutionController) GetElementProperty(property string) interface{} {
	if c.PropElement == nil {
		return nil
	}

	// Access element property via the JavaScript value
	if jsValue, ok := c.PropElement.(app.Value); ok && !jsValue.IsNull() {
		return jsValue.Get(property).Get("value").String()
	}

	return nil
}

// CallElementMethod calls a method on the current element reference
func (c *spectrumElementResolutionController) CallElementMethod(method string, args ...interface{}) interface{} {
	if c.PropElement == nil {
		return nil
	}

	// Call element method via the JavaScript value
	if jsValue, ok := c.PropElement.(app.Value); ok && !jsValue.IsNull() {
		return jsValue.Call(method, args...)
	}

	return nil
}
