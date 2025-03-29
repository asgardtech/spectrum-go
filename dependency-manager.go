package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumDependencyManager represents a manager for component dependencies
type spectrumDependencyManager struct {
	app.Compo

	// Properties
	PropDependencies       map[string]bool
	PropOnDependencyChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewDependencyManager creates a new dependency manager component
func NewDependencyManager() *spectrumDependencyManager {
	return &spectrumDependencyManager{
		PropDependencies: make(map[string]bool),
	}
}

// AddDependency adds a dependency
func (d *spectrumDependencyManager) AddDependency(name string) *spectrumDependencyManager {
	d.PropDependencies[name] = true
	return d
}

// RemoveDependency removes a dependency
func (d *spectrumDependencyManager) RemoveDependency(name string) *spectrumDependencyManager {
	delete(d.PropDependencies, name)
	return d
}

// HasDependency checks if a dependency exists
func (d *spectrumDependencyManager) HasDependency(name string) bool {
	return d.PropDependencies[name]
}

// ClearDependencies clears all dependencies
func (d *spectrumDependencyManager) ClearDependencies() *spectrumDependencyManager {
	d.PropDependencies = make(map[string]bool)
	return d
}

// OnDependencyChange sets the dependency change event handler
func (d *spectrumDependencyManager) OnDependencyChange(handler app.EventHandler) *spectrumDependencyManager {
	d.PropOnDependencyChange = handler
	return d
}

// Child adds a child element
func (d *spectrumDependencyManager) Child(child app.UI) *spectrumDependencyManager {
	d.PropChildren = append(d.PropChildren, child)
	return d
}

// Children adds multiple child elements
func (d *spectrumDependencyManager) Children(children ...app.UI) *spectrumDependencyManager {
	d.PropChildren = append(d.PropChildren, children...)
	return d
}

// Render renders the dependency manager component
func (d *spectrumDependencyManager) Render() app.UI {
	manager := app.Elem("div")

	// Set dependencies
	for name := range d.PropDependencies {
		manager = manager.Attr("data-dependency", name)
	}

	// Add event handler
	if d.PropOnDependencyChange != nil {
		manager = manager.On("dependency-change", d.PropOnDependencyChange)
	}

	// Add children if provided
	if len(d.PropChildren) > 0 {
		manager = manager.Body(d.PropChildren...)
	}

	return manager
}

// spectrumDependencyController represents a controller for managing component dependencies
type spectrumDependencyController struct {
	app.Compo

	// Properties
	PropManager            *spectrumDependencyManager
	PropOnDependencyChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewDependencyController creates a new dependency controller component
func NewDependencyController() *spectrumDependencyController {
	return &spectrumDependencyController{
		PropManager: NewDependencyManager(),
	}
}

// Manager returns the dependency manager
func (c *spectrumDependencyController) Manager() *spectrumDependencyManager {
	return c.PropManager
}

// OnDependencyChange sets the dependency change event handler
func (c *spectrumDependencyController) OnDependencyChange(handler app.EventHandler) *spectrumDependencyController {
	c.PropOnDependencyChange = handler
	return c
}

// Child adds a child element
func (c *spectrumDependencyController) Child(child app.UI) *spectrumDependencyController {
	c.PropChildren = append(c.PropChildren, child)
	return c
}

// Children adds multiple child elements
func (c *spectrumDependencyController) Children(children ...app.UI) *spectrumDependencyController {
	c.PropChildren = append(c.PropChildren, children...)
	return c
}

// Render renders the dependency controller component
func (c *spectrumDependencyController) Render() app.UI {
	controller := app.Elem("div")

	// Add manager
	controller = controller.Body(c.PropManager)

	// Add event handler
	if c.PropOnDependencyChange != nil {
		controller = controller.On("dependency-change", c.PropOnDependencyChange)
	}

	// Add children if provided
	if len(c.PropChildren) > 0 {
		controller = controller.Body(c.PropChildren...)
	}

	return controller
}
