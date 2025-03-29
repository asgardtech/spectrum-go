package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumReactiveController represents a controller for managing reactive state and behavior
type spectrumReactiveController struct {
	app.Compo

	// Properties
	PropState         map[string]interface{}
	PropObservers     map[string][]app.EventHandler
	PropOnStateChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewReactiveController creates a new reactive controller component
func NewReactiveController() *spectrumReactiveController {
	return &spectrumReactiveController{
		PropState:     make(map[string]interface{}),
		PropObservers: make(map[string][]app.EventHandler),
	}
}

// SetState sets a state value
func (c *spectrumReactiveController) SetState(key string, value interface{}) *spectrumReactiveController {
	c.PropState[key] = value
	return c
}

// GetState gets a state value
func (c *spectrumReactiveController) GetState(key string) interface{} {
	return c.PropState[key]
}

// DeleteState deletes a state value
func (c *spectrumReactiveController) DeleteState(key string) *spectrumReactiveController {
	delete(c.PropState, key)
	return c
}

// ClearState clears all state values
func (c *spectrumReactiveController) ClearState() *spectrumReactiveController {
	c.PropState = make(map[string]interface{})
	return c
}

// Observe adds an observer for a state key
func (c *spectrumReactiveController) Observe(key string, handler app.EventHandler) *spectrumReactiveController {
	c.PropObservers[key] = append(c.PropObservers[key], handler)
	return c
}

// Unobserve removes an observer for a state key
func (c *spectrumReactiveController) Unobserve(key string, handler app.EventHandler) *spectrumReactiveController {
	// Instead of comparing function pointers, we'll remove the last added handler
	// This is a limitation of Go's function comparison
	if handlers, exists := c.PropObservers[key]; exists && len(handlers) > 0 {
		c.PropObservers[key] = handlers[:len(handlers)-1]
	}
	return c
}

// ClearObservers clears all observers
func (c *spectrumReactiveController) ClearObservers() *spectrumReactiveController {
	c.PropObservers = make(map[string][]app.EventHandler)
	return c
}

// OnStateChange sets the state change event handler
func (c *spectrumReactiveController) OnStateChange(handler app.EventHandler) *spectrumReactiveController {
	c.PropOnStateChange = handler
	return c
}

// Child adds a child element
func (c *spectrumReactiveController) Child(child app.UI) *spectrumReactiveController {
	c.PropChildren = append(c.PropChildren, child)
	return c
}

// Children adds multiple child elements
func (c *spectrumReactiveController) Children(children ...app.UI) *spectrumReactiveController {
	c.PropChildren = append(c.PropChildren, children...)
	return c
}

// Render renders the reactive controller component
func (c *spectrumReactiveController) Render() app.UI {
	controller := app.Elem("div").
		Class("spectrum-ReactiveController")

	// Add event handlers
	if c.PropOnStateChange != nil {
		controller = controller.On("state-change", c.PropOnStateChange)
	}

	// Add children if provided
	if len(c.PropChildren) > 0 {
		controller = controller.Body(c.PropChildren...)
	}

	return controller
}

// ReactiveState represents a reactive state value
type ReactiveState struct {
	app.Compo

	// Properties
	value    interface{}
	onChange app.EventHandler
	onUpdate app.EventHandler

	// Children
	children []app.UI
}

// NewReactiveState creates a new reactive state component
func NewReactiveState() *ReactiveState {
	return &ReactiveState{}
}

// Value sets the state value
func (s *ReactiveState) Value(value interface{}) *ReactiveState {
	s.value = value
	return s
}

// GetValue gets the state value
func (s *ReactiveState) GetValue() interface{} {
	return s.value
}

// OnChange sets the change event handler
func (s *ReactiveState) OnChange(handler app.EventHandler) *ReactiveState {
	s.onChange = handler
	return s
}

// OnUpdate sets the update event handler
func (s *ReactiveState) OnUpdate(handler app.EventHandler) *ReactiveState {
	s.onUpdate = handler
	return s
}

// Child adds a child element
func (s *ReactiveState) Child(child app.UI) *ReactiveState {
	s.children = append(s.children, child)
	return s
}

// Children adds multiple child elements
func (s *ReactiveState) Children(children ...app.UI) *ReactiveState {
	s.children = append(s.children, children...)
	return s
}

// Render renders the reactive state component
func (s *ReactiveState) Render() app.UI {
	state := app.Elem("div")

	// Set value
	if s.value != nil {
		state = state.Attr("data-value", s.value)
	}

	// Add event handlers
	if s.onChange != nil {
		state = state.On("change", s.onChange)
	}
	if s.onUpdate != nil {
		state = state.On("update", s.onUpdate)
	}

	// Add children if provided
	if len(s.children) > 0 {
		state = state.Body(s.children...)
	}

	return state
}
