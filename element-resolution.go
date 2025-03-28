package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ElementResolverUpdatedSymbol is a sentinel value to indicate when an element reference has been updated
var ElementResolverUpdatedSymbol = "element-resolver-updated"

// SpectrumElementResolutionController keeps an active reference to another element in the same DOM tree
// It manages observing the DOM tree to ensure that the reference it holds is always the first matched element or nil
type SpectrumElementResolutionController struct {
	host        app.Composer
	Element     app.UI
	Selector    string
	mutationObs app.MutationObserver
	observing   bool
}

// ElementResolutionController creates a new SpectrumElementResolutionController
func ElementResolutionController(host app.Composer) *SpectrumElementResolutionController {
	controller := &SpectrumElementResolutionController{
		host: host,
	}

	// Initialize the controller
	return controller
}

// SetSelector sets the CSS selector to use for element resolution and updates the reference
func (c *SpectrumElementResolutionController) SetSelector(selector string) *SpectrumElementResolutionController {
	c.Selector = selector
	c.updateReference()
	return c
}

// Connect starts observing DOM mutations if a selector is provided
func (c *SpectrumElementResolutionController) Connect() {
	if c.Selector != "" && !c.observing {
		c.startObserving()
	}
}

// Disconnect stops observing DOM mutations
func (c *SpectrumElementResolutionController) Disconnect() {
	if c.observing {
		c.stopObserving()
	}
}

// startObserving sets up a MutationObserver to track DOM changes that might affect the resolved element
func (c *SpectrumElementResolutionController) startObserving() {
	if c.mutationObs != nil {
		return
	}

	// Set up mutation observer options
	options := app.MutationObserverInit{
		ChildList:       true,
		Attributes:      true,
		CharacterData:   true,
		Subtree:         true,
		AttributeFilter: []string{"class", "id", "hidden"},
	}

	// Create the mutation observer
	c.mutationObs = app.NewMutationObserver(func(mutations []app.Mutation) {
		c.handleMutations(mutations)
	})

	// Start observing the document body
	c.mutationObs.Observe(app.Window().Get("document").Get("body"), options)
	c.observing = true

	// Make initial reference update
	c.updateReference()
}

// stopObserving disconnects the mutation observer
func (c *SpectrumElementResolutionController) stopObserving() {
	if c.mutationObs != nil {
		c.mutationObs.Disconnect()
		c.mutationObs = nil
	}
	c.observing = false
}

// handleMutations processes DOM mutations and updates the element reference if needed
func (c *SpectrumElementResolutionController) handleMutations(mutations []app.Mutation) {
	// Check if any of the mutations are relevant to our selector
	shouldUpdate := false

	for _, mutation := range mutations {
		// If nodes were added or removed, we need to update
		if mutation.Type == "childList" {
			shouldUpdate = true
			break
		}

		// If attributes changed on potential target elements, we need to update
		if mutation.Type == "attributes" {
			attrName := mutation.AttributeName
			if attrName == "class" || attrName == "id" || attrName == "hidden" {
				shouldUpdate = true
				break
			}
		}
	}

	if shouldUpdate {
		c.updateReference()
	}
}

// updateReference queries the DOM and updates the element reference
func (c *SpectrumElementResolutionController) updateReference() {
	if c.Selector == "" {
		c.Element = nil
		// Notify the host component that the reference has been updated
		if _, ok := c.host.(app.Updater); ok {
			c.host.(app.Updater).RequestUpdate(ElementResolverUpdatedSymbol)
		}
		return
	}

	// Query the DOM for the first element matching the selector
	doc := app.Window().Get("document")
	element := doc.QuerySelector(c.Selector)

	// Check if the element has changed
	oldElement := c.Element

	if !element.IsNull() {
		// We found an element, convert it to app.UI type
		// This is a simplified representation - in a real implementation
		// you would need to convert the JS element to an appropriate app.UI type
		c.Element = app.Elem("div") // Placeholder - actual conversion would depend on the element type
	} else {
		c.Element = nil
	}

	// If the element reference has changed, notify the host
	if oldElement != c.Element {
		// Notify the host component that the reference has been updated
		if _, ok := c.host.(app.Updater); ok {
			c.host.(app.Updater).RequestUpdate(ElementResolverUpdatedSymbol)
		}
	}
}
