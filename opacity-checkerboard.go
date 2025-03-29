package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// spectrumOpacityCheckerboard represents a component that displays an opacity checkerboard pattern
type spectrumOpacityCheckerboard struct {
	app.Compo

	// Properties
	PropWidth  string
	PropHeight string
	PropOnLoad app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumOpacityCheckerboard creates a new opacity checkerboard component
func NewSpectrumOpacityCheckerboard() *spectrumOpacityCheckerboard {
	return &spectrumOpacityCheckerboard{}
}

// Width sets the width of the checkerboard
func (c *spectrumOpacityCheckerboard) Width(width string) *spectrumOpacityCheckerboard {
	c.PropWidth = width
	return c
}

// Height sets the height of the checkerboard
func (c *spectrumOpacityCheckerboard) Height(height string) *spectrumOpacityCheckerboard {
	c.PropHeight = height
	return c
}

// Size sets both width and height of the checkerboard
func (c *spectrumOpacityCheckerboard) Size(width, height string) *spectrumOpacityCheckerboard {
	c.PropWidth = width
	c.PropHeight = height
	return c
}

// OnLoad sets the load event handler
func (c *spectrumOpacityCheckerboard) OnLoad(handler app.EventHandler) *spectrumOpacityCheckerboard {
	c.PropOnLoad = handler
	return c
}

// Child adds a child element
func (c *spectrumOpacityCheckerboard) Child(child app.UI) *spectrumOpacityCheckerboard {
	c.PropChildren = append(c.PropChildren, child)
	return c
}

// Children adds multiple child elements
func (c *spectrumOpacityCheckerboard) Children(children ...app.UI) *spectrumOpacityCheckerboard {
	c.PropChildren = append(c.PropChildren, children...)
	return c
}

// Render renders the opacity checkerboard component
func (c *spectrumOpacityCheckerboard) Render() app.UI {
	checkerboard := app.Elem("div").
		Class("spectrum-OpacityCheckerboard")

	// Set dimensions
	if c.PropWidth != "" {
		checkerboard = checkerboard.Style("--spectrum-opacity-checkerboard-width", c.PropWidth)
	}
	if c.PropHeight != "" {
		checkerboard = checkerboard.Style("--spectrum-opacity-checkerboard-height", c.PropHeight)
	}

	// Add event handler
	if c.PropOnLoad != nil {
		checkerboard = checkerboard.On("load", c.PropOnLoad)
	}

	// Add children if provided
	if len(c.PropChildren) > 0 {
		checkerboard = checkerboard.Body(c.PropChildren...)
	}

	return checkerboard
}

// spectrumOpacityCheckerboardManager represents a manager for opacity checkerboard components
type spectrumOpacityCheckerboardManager struct {
	app.Compo

	// Properties
	PropComponents        map[string]*spectrumOpacityCheckerboard
	PropOnComponentChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumOpacityCheckerboardManager creates a new opacity checkerboard manager component
func NewSpectrumOpacityCheckerboardManager() *spectrumOpacityCheckerboardManager {
	return &spectrumOpacityCheckerboardManager{
		PropComponents: make(map[string]*spectrumOpacityCheckerboard),
	}
}

// AddComponent adds an opacity checkerboard component
func (m *spectrumOpacityCheckerboardManager) AddComponent(name string, component *spectrumOpacityCheckerboard) *spectrumOpacityCheckerboardManager {
	m.PropComponents[name] = component
	return m
}

// RemoveComponent removes an opacity checkerboard component
func (m *spectrumOpacityCheckerboardManager) RemoveComponent(name string) *spectrumOpacityCheckerboardManager {
	delete(m.PropComponents, name)
	return m
}

// GetComponent gets an opacity checkerboard component by name
func (m *spectrumOpacityCheckerboardManager) GetComponent(name string) *spectrumOpacityCheckerboard {
	return m.PropComponents[name]
}

// ClearComponents clears all opacity checkerboard components
func (m *spectrumOpacityCheckerboardManager) ClearComponents() *spectrumOpacityCheckerboardManager {
	m.PropComponents = make(map[string]*spectrumOpacityCheckerboard)
	return m
}

// OnComponentChange sets the component change event handler
func (m *spectrumOpacityCheckerboardManager) OnComponentChange(handler app.EventHandler) *spectrumOpacityCheckerboardManager {
	m.PropOnComponentChange = handler
	return m
}

// Child adds a child element
func (m *spectrumOpacityCheckerboardManager) Child(child app.UI) *spectrumOpacityCheckerboardManager {
	m.PropChildren = append(m.PropChildren, child)
	return m
}

// Children adds multiple child elements
func (m *spectrumOpacityCheckerboardManager) Children(children ...app.UI) *spectrumOpacityCheckerboardManager {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// Render renders the opacity checkerboard manager component
func (m *spectrumOpacityCheckerboardManager) Render() app.UI {
	manager := app.Elem("div")

	// Add components
	for _, component := range m.PropComponents {
		manager = manager.Body(component)
	}

	// Add event handler
	if m.PropOnComponentChange != nil {
		manager = manager.On("component-change", m.PropOnComponentChange)
	}

	// Add children if provided
	if len(m.PropChildren) > 0 {
		manager = manager.Body(m.PropChildren...)
	}

	return manager
}
