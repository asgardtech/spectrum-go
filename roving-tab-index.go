package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// RovingTabIndexDirection represents the direction for roving tab index navigation
type RovingTabIndexDirection string

// Roving tab index directions
const (
	RovingTabIndexDirectionBoth       RovingTabIndexDirection = "both"
	RovingTabIndexDirectionVertical   RovingTabIndexDirection = "vertical"
	RovingTabIndexDirectionHorizontal RovingTabIndexDirection = "horizontal"
	RovingTabIndexDirectionGrid       RovingTabIndexDirection = "grid"
)

// SpectrumRovingTabIndex represents a component that manages keyboard navigation
type SpectrumRovingTabIndex struct {
	app.Compo

	// Properties
	direction         string
	focusableSelector string
	onFocus           app.EventHandler
	onBlur            app.EventHandler
	onKeyDown         app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumRovingTabIndex creates a new roving tab index component
func NewSpectrumRovingTabIndex() *SpectrumRovingTabIndex {
	return &SpectrumRovingTabIndex{}
}

// Direction sets the navigation direction
func (r *SpectrumRovingTabIndex) Direction(direction string) *SpectrumRovingTabIndex {
	r.direction = direction
	return r
}

// FocusableSelector sets the selector for focusable elements
func (r *SpectrumRovingTabIndex) FocusableSelector(selector string) *SpectrumRovingTabIndex {
	r.focusableSelector = selector
	return r
}

// OnFocus sets the focus event handler
func (r *SpectrumRovingTabIndex) OnFocus(handler app.EventHandler) *SpectrumRovingTabIndex {
	r.onFocus = handler
	return r
}

// OnBlur sets the blur event handler
func (r *SpectrumRovingTabIndex) OnBlur(handler app.EventHandler) *SpectrumRovingTabIndex {
	r.onBlur = handler
	return r
}

// OnKeyDown sets the key down event handler
func (r *SpectrumRovingTabIndex) OnKeyDown(handler app.EventHandler) *SpectrumRovingTabIndex {
	r.onKeyDown = handler
	return r
}

// Child adds a child element
func (r *SpectrumRovingTabIndex) Child(child app.UI) *SpectrumRovingTabIndex {
	r.children = append(r.children, child)
	return r
}

// Children adds multiple child elements
func (r *SpectrumRovingTabIndex) Children(children ...app.UI) *SpectrumRovingTabIndex {
	r.children = append(r.children, children...)
	return r
}

// Render renders the roving tab index component
func (r *SpectrumRovingTabIndex) Render() app.UI {
	roving := app.Elem("div")

	// Set attributes
	if r.direction != "" {
		roving = roving.Attr("data-direction", r.direction)
	}
	if r.focusableSelector != "" {
		roving = roving.Attr("data-focusable-selector", r.focusableSelector)
	}

	// Add event handlers
	if r.onFocus != nil {
		roving = roving.On("focus", r.onFocus)
	}
	if r.onBlur != nil {
		roving = roving.On("blur", r.onBlur)
	}
	if r.onKeyDown != nil {
		roving = roving.On("keydown", r.onKeyDown)
	}

	// Add children if provided
	if len(r.children) > 0 {
		roving = roving.Body(r.children...)
	}

	return roving
}

// SpectrumRovingTabIndexManager represents a manager for roving tab index components
type SpectrumRovingTabIndexManager struct {
	app.Compo

	// Properties
	components        map[string]*SpectrumRovingTabIndex
	onComponentChange app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumRovingTabIndexManager creates a new roving tab index manager component
func NewSpectrumRovingTabIndexManager() *SpectrumRovingTabIndexManager {
	return &SpectrumRovingTabIndexManager{
		components: make(map[string]*SpectrumRovingTabIndex),
	}
}

// AddComponent adds a roving tab index component
func (m *SpectrumRovingTabIndexManager) AddComponent(name string, component *SpectrumRovingTabIndex) *SpectrumRovingTabIndexManager {
	m.components[name] = component
	return m
}

// RemoveComponent removes a roving tab index component
func (m *SpectrumRovingTabIndexManager) RemoveComponent(name string) *SpectrumRovingTabIndexManager {
	delete(m.components, name)
	return m
}

// GetComponent gets a roving tab index component by name
func (m *SpectrumRovingTabIndexManager) GetComponent(name string) *SpectrumRovingTabIndex {
	return m.components[name]
}

// ClearComponents clears all roving tab index components
func (m *SpectrumRovingTabIndexManager) ClearComponents() *SpectrumRovingTabIndexManager {
	m.components = make(map[string]*SpectrumRovingTabIndex)
	return m
}

// OnComponentChange sets the component change event handler
func (m *SpectrumRovingTabIndexManager) OnComponentChange(handler app.EventHandler) *SpectrumRovingTabIndexManager {
	m.onComponentChange = handler
	return m
}

// Child adds a child element
func (m *SpectrumRovingTabIndexManager) Child(child app.UI) *SpectrumRovingTabIndexManager {
	m.children = append(m.children, child)
	return m
}

// Children adds multiple child elements
func (m *SpectrumRovingTabIndexManager) Children(children ...app.UI) *SpectrumRovingTabIndexManager {
	m.children = append(m.children, children...)
	return m
}

// Render renders the roving tab index manager component
func (m *SpectrumRovingTabIndexManager) Render() app.UI {
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
