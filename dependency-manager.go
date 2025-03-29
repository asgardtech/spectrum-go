package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DependencyManager represents a manager for component dependencies
type DependencyManager struct {
	app.Compo

	// Properties
	dependencies       map[string]bool
	onDependencyChange app.EventHandler

	// Children
	children []app.UI
}

// NewDependencyManager creates a new dependency manager component
func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		dependencies: make(map[string]bool),
	}
}

// AddDependency adds a dependency
func (d *DependencyManager) AddDependency(name string) *DependencyManager {
	d.dependencies[name] = true
	return d
}

// RemoveDependency removes a dependency
func (d *DependencyManager) RemoveDependency(name string) *DependencyManager {
	delete(d.dependencies, name)
	return d
}

// HasDependency checks if a dependency exists
func (d *DependencyManager) HasDependency(name string) bool {
	return d.dependencies[name]
}

// ClearDependencies clears all dependencies
func (d *DependencyManager) ClearDependencies() *DependencyManager {
	d.dependencies = make(map[string]bool)
	return d
}

// OnDependencyChange sets the dependency change event handler
func (d *DependencyManager) OnDependencyChange(handler app.EventHandler) *DependencyManager {
	d.onDependencyChange = handler
	return d
}

// Child adds a child element
func (d *DependencyManager) Child(child app.UI) *DependencyManager {
	d.children = append(d.children, child)
	return d
}

// Children adds multiple child elements
func (d *DependencyManager) Children(children ...app.UI) *DependencyManager {
	d.children = append(d.children, children...)
	return d
}

// Render renders the dependency manager component
func (d *DependencyManager) Render() app.UI {
	manager := app.Elem("div")

	// Set dependencies
	for name := range d.dependencies {
		manager = manager.Attr("data-dependency", name)
	}

	// Add event handler
	if d.onDependencyChange != nil {
		manager = manager.On("dependency-change", d.onDependencyChange)
	}

	// Add children if provided
	if len(d.children) > 0 {
		manager = manager.Body(d.children...)
	}

	return manager
}

// DependencyController represents a controller for managing component dependencies
type DependencyController struct {
	app.Compo

	// Properties
	manager            *DependencyManager
	onDependencyChange app.EventHandler

	// Children
	children []app.UI
}

// NewDependencyController creates a new dependency controller component
func NewDependencyController() *DependencyController {
	return &DependencyController{
		manager: NewDependencyManager(),
	}
}

// Manager returns the dependency manager
func (c *DependencyController) Manager() *DependencyManager {
	return c.manager
}

// OnDependencyChange sets the dependency change event handler
func (c *DependencyController) OnDependencyChange(handler app.EventHandler) *DependencyController {
	c.onDependencyChange = handler
	return c
}

// Child adds a child element
func (c *DependencyController) Child(child app.UI) *DependencyController {
	c.children = append(c.children, child)
	return c
}

// Children adds multiple child elements
func (c *DependencyController) Children(children ...app.UI) *DependencyController {
	c.children = append(c.children, children...)
	return c
}

// Render renders the dependency controller component
func (c *DependencyController) Render() app.UI {
	controller := app.Elem("div")

	// Add manager
	controller = controller.Body(c.manager)

	// Add event handler
	if c.onDependencyChange != nil {
		controller = controller.On("dependency-change", c.onDependencyChange)
	}

	// Add children if provided
	if len(c.children) > 0 {
		controller = controller.Body(c.children...)
	}

	return controller
}
