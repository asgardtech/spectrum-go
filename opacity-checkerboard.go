package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SpectrumOpacityCheckerboard represents a component that displays an opacity checkerboard pattern
type SpectrumOpacityCheckerboard struct {
	app.Compo

	// Properties
	width  string
	height string
	onLoad app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumOpacityCheckerboard creates a new opacity checkerboard component
func NewSpectrumOpacityCheckerboard() *SpectrumOpacityCheckerboard {
	return &SpectrumOpacityCheckerboard{}
}

// Width sets the width of the checkerboard
func (c *SpectrumOpacityCheckerboard) Width(width string) *SpectrumOpacityCheckerboard {
	c.width = width
	return c
}

// Height sets the height of the checkerboard
func (c *SpectrumOpacityCheckerboard) Height(height string) *SpectrumOpacityCheckerboard {
	c.height = height
	return c
}

// Size sets both width and height of the checkerboard
func (c *SpectrumOpacityCheckerboard) Size(width, height string) *SpectrumOpacityCheckerboard {
	c.width = width
	c.height = height
	return c
}

// OnLoad sets the load event handler
func (c *SpectrumOpacityCheckerboard) OnLoad(handler app.EventHandler) *SpectrumOpacityCheckerboard {
	c.onLoad = handler
	return c
}

// Child adds a child element
func (c *SpectrumOpacityCheckerboard) Child(child app.UI) *SpectrumOpacityCheckerboard {
	c.children = append(c.children, child)
	return c
}

// Children adds multiple child elements
func (c *SpectrumOpacityCheckerboard) Children(children ...app.UI) *SpectrumOpacityCheckerboard {
	c.children = append(c.children, children...)
	return c
}

// Render renders the opacity checkerboard component
func (c *SpectrumOpacityCheckerboard) Render() app.UI {
	checkerboard := app.Elem("div").
		Class("spectrum-OpacityCheckerboard")

	// Set dimensions
	if c.width != "" {
		checkerboard = checkerboard.Style("--spectrum-opacity-checkerboard-width", c.width)
	}
	if c.height != "" {
		checkerboard = checkerboard.Style("--spectrum-opacity-checkerboard-height", c.height)
	}

	// Add event handler
	if c.onLoad != nil {
		checkerboard = checkerboard.On("load", c.onLoad)
	}

	// Add children if provided
	if len(c.children) > 0 {
		checkerboard = checkerboard.Body(c.children...)
	}

	return checkerboard
}

// SpectrumOpacityCheckerboardManager represents a manager for opacity checkerboard components
type SpectrumOpacityCheckerboardManager struct {
	app.Compo

	// Properties
	components        map[string]*SpectrumOpacityCheckerboard
	onComponentChange app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumOpacityCheckerboardManager creates a new opacity checkerboard manager component
func NewSpectrumOpacityCheckerboardManager() *SpectrumOpacityCheckerboardManager {
	return &SpectrumOpacityCheckerboardManager{
		components: make(map[string]*SpectrumOpacityCheckerboard),
	}
}

// AddComponent adds an opacity checkerboard component
func (m *SpectrumOpacityCheckerboardManager) AddComponent(name string, component *SpectrumOpacityCheckerboard) *SpectrumOpacityCheckerboardManager {
	m.components[name] = component
	return m
}

// RemoveComponent removes an opacity checkerboard component
func (m *SpectrumOpacityCheckerboardManager) RemoveComponent(name string) *SpectrumOpacityCheckerboardManager {
	delete(m.components, name)
	return m
}

// GetComponent gets an opacity checkerboard component by name
func (m *SpectrumOpacityCheckerboardManager) GetComponent(name string) *SpectrumOpacityCheckerboard {
	return m.components[name]
}

// ClearComponents clears all opacity checkerboard components
func (m *SpectrumOpacityCheckerboardManager) ClearComponents() *SpectrumOpacityCheckerboardManager {
	m.components = make(map[string]*SpectrumOpacityCheckerboard)
	return m
}

// OnComponentChange sets the component change event handler
func (m *SpectrumOpacityCheckerboardManager) OnComponentChange(handler app.EventHandler) *SpectrumOpacityCheckerboardManager {
	m.onComponentChange = handler
	return m
}

// Child adds a child element
func (m *SpectrumOpacityCheckerboardManager) Child(child app.UI) *SpectrumOpacityCheckerboardManager {
	m.children = append(m.children, child)
	return m
}

// Children adds multiple child elements
func (m *SpectrumOpacityCheckerboardManager) Children(children ...app.UI) *SpectrumOpacityCheckerboardManager {
	m.children = append(m.children, children...)
	return m
}

// Render renders the opacity checkerboard manager component
func (m *SpectrumOpacityCheckerboardManager) Render() app.UI {
	manager := app.Elem("div")

	// Add components
	for _, component := range m.components {
		manager = manager.Body(component)
	}

	// Add event handler
	if m.onComponentChange != nil {
		manager = manager.On("component-change", m.onComponentChange)
	}

	// Add children if provided
	if len(m.children) > 0 {
		manager = manager.Body(m.children...)
	}

	return manager
}
