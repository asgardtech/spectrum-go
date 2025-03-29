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

// spectrumRovingTabIndex represents a component that manages keyboard navigation
type spectrumRovingTabIndex struct {
	app.Compo

	// Properties
	PropDirection         string
	PropFocusableSelector string
	PropOnFocus           app.EventHandler
	PropOnBlur            app.EventHandler
	PropOnKeyDown         app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumRovingTabIndex creates a new roving tab index component
func NewSpectrumRovingTabIndex() *spectrumRovingTabIndex {
	return &spectrumRovingTabIndex{}
}

// Direction sets the navigation direction
func (r *spectrumRovingTabIndex) Direction(direction string) *spectrumRovingTabIndex {
	r.PropDirection = direction
	return r
}

// FocusableSelector sets the selector for focusable elements
func (r *spectrumRovingTabIndex) FocusableSelector(selector string) *spectrumRovingTabIndex {
	r.PropFocusableSelector = selector
	return r
}

// OnFocus sets the focus event handler
func (r *spectrumRovingTabIndex) OnFocus(handler app.EventHandler) *spectrumRovingTabIndex {
	r.PropOnFocus = handler
	return r
}

// OnBlur sets the blur event handler
func (r *spectrumRovingTabIndex) OnBlur(handler app.EventHandler) *spectrumRovingTabIndex {
	r.PropOnBlur = handler
	return r
}

// OnKeyDown sets the key down event handler
func (r *spectrumRovingTabIndex) OnKeyDown(handler app.EventHandler) *spectrumRovingTabIndex {
	r.PropOnKeyDown = handler
	return r
}

// Child adds a child element
func (r *spectrumRovingTabIndex) Child(child app.UI) *spectrumRovingTabIndex {
	r.PropChildren = append(r.PropChildren, child)
	return r
}

// Children adds multiple child elements
func (r *spectrumRovingTabIndex) Children(children ...app.UI) *spectrumRovingTabIndex {
	r.PropChildren = append(r.PropChildren, children...)
	return r
}

// Render renders the roving tab index component
func (r *spectrumRovingTabIndex) Render() app.UI {
	roving := app.Elem("div")

	// Set attributes
	if r.PropDirection != "" {
		roving = roving.Attr("data-direction", r.PropDirection)
	}
	if r.PropFocusableSelector != "" {
		roving = roving.Attr("data-focusable-selector", r.PropFocusableSelector)
	}

	// Add event handlers
	if r.PropOnFocus != nil {
		roving = roving.On("focus", r.PropOnFocus)
	}
	if r.PropOnBlur != nil {
		roving = roving.On("blur", r.PropOnBlur)
	}
	if r.PropOnKeyDown != nil {
		roving = roving.On("keydown", r.PropOnKeyDown)
	}

	// Add children if provided
	if len(r.PropChildren) > 0 {
		roving = roving.Body(r.PropChildren...)
	}

	return roving
}

// spectrumRovingTabIndexManager represents a manager for roving tab index components
type spectrumRovingTabIndexManager struct {
	app.Compo

	// Properties
	PropComponents        map[string]*spectrumRovingTabIndex
	PropOnComponentChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumRovingTabIndexManager creates a new roving tab index manager component
func NewSpectrumRovingTabIndexManager() *spectrumRovingTabIndexManager {
	return &spectrumRovingTabIndexManager{
		PropComponents: make(map[string]*spectrumRovingTabIndex),
	}
}

// AddComponent adds a roving tab index component
func (m *spectrumRovingTabIndexManager) AddComponent(name string, component *spectrumRovingTabIndex) *spectrumRovingTabIndexManager {
	m.PropComponents[name] = component
	return m
}

// RemoveComponent removes a roving tab index component
func (m *spectrumRovingTabIndexManager) RemoveComponent(name string) *spectrumRovingTabIndexManager {
	delete(m.PropComponents, name)
	return m
}

// GetComponent gets a roving tab index component by name
func (m *spectrumRovingTabIndexManager) GetComponent(name string) *spectrumRovingTabIndex {
	return m.PropComponents[name]
}

// ClearComponents clears all roving tab index components
func (m *spectrumRovingTabIndexManager) ClearComponents() *spectrumRovingTabIndexManager {
	m.PropComponents = make(map[string]*spectrumRovingTabIndex)
	return m
}

// OnComponentChange sets the component change event handler
func (m *spectrumRovingTabIndexManager) OnComponentChange(handler app.EventHandler) *spectrumRovingTabIndexManager {
	m.PropOnComponentChange = handler
	return m
}

// Child adds a child element
func (m *spectrumRovingTabIndexManager) Child(child app.UI) *spectrumRovingTabIndexManager {
	m.PropChildren = append(m.PropChildren, child)
	return m
}

// Children adds multiple child elements
func (m *spectrumRovingTabIndexManager) Children(children ...app.UI) *spectrumRovingTabIndexManager {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// Render renders the roving tab index manager component
func (m *spectrumRovingTabIndexManager) Render() app.UI {
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
