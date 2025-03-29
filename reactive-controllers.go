package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ReactiveController represents a controller for managing reactive state and behavior
type ReactiveController struct {
	app.Compo

	// Properties
	state         map[string]interface{}
	observers     map[string][]app.EventHandler
	onStateChange app.EventHandler

	// Children
	children []app.UI
}

// NewReactiveController creates a new reactive controller component
func NewReactiveController() *ReactiveController {
	return &ReactiveController{
		state:     make(map[string]interface{}),
		observers: make(map[string][]app.EventHandler),
	}
}

// SetState sets a state value
func (c *ReactiveController) SetState(key string, value interface{}) *ReactiveController {
	c.state[key] = value
	return c
}

// GetState gets a state value
func (c *ReactiveController) GetState(key string) interface{} {
	return c.state[key]
}

// DeleteState deletes a state value
func (c *ReactiveController) DeleteState(key string) *ReactiveController {
	delete(c.state, key)
	return c
}

// ClearState clears all state values
func (c *ReactiveController) ClearState() *ReactiveController {
	c.state = make(map[string]interface{})
	return c
}

// Observe adds an observer for a state key
func (c *ReactiveController) Observe(key string, handler app.EventHandler) *ReactiveController {
	c.observers[key] = append(c.observers[key], handler)
	return c
}

// Unobserve removes an observer for a state key
func (c *ReactiveController) Unobserve(key string, handler app.EventHandler) *ReactiveController {
	// Instead of comparing function pointers, we'll remove the last added handler
	// This is a limitation of Go's function comparison
	if handlers, exists := c.observers[key]; exists && len(handlers) > 0 {
		c.observers[key] = handlers[:len(handlers)-1]
	}
	return c
}

// ClearObservers clears all observers
func (c *ReactiveController) ClearObservers() *ReactiveController {
	c.observers = make(map[string][]app.EventHandler)
	return c
}

// OnStateChange sets the state change event handler
func (c *ReactiveController) OnStateChange(handler app.EventHandler) *ReactiveController {
	c.onStateChange = handler
	return c
}

// Child adds a child element
func (c *ReactiveController) Child(child app.UI) *ReactiveController {
	c.children = append(c.children, child)
	return c
}

// Children adds multiple child elements
func (c *ReactiveController) Children(children ...app.UI) *ReactiveController {
	c.children = append(c.children, children...)
	return c
}

// Render renders the reactive controller component
func (c *ReactiveController) Render() app.UI {
	controller := app.Elem("div")

	// Set state
	for key, value := range c.state {
		controller = controller.Attr("data-state-"+key, value)
	}

	// Add event handler
	if c.onStateChange != nil {
		controller = controller.On("state-change", c.onStateChange)
	}

	// Add children if provided
	if len(c.children) > 0 {
		controller = controller.Body(c.children...)
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
