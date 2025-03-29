package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTruncated represents a component that truncates text with tooltips
type SpectrumTruncated struct {
	app.Compo

	// Properties
	placement string
	content   string
	innerHTML string
	child     app.UI
	overflow  app.UI

	// Children
	children []app.UI
}

// NewSpectrumTruncated creates a new truncated component
func NewSpectrumTruncated() *SpectrumTruncated {
	return &SpectrumTruncated{}
}

// Placement sets the tooltip placement
func (t *SpectrumTruncated) Placement(placement string) *SpectrumTruncated {
	t.placement = placement
	return t
}

// Content sets the content to be truncated
func (t *SpectrumTruncated) Content(content string) *SpectrumTruncated {
	t.content = content
	return t
}

// InnerHTML sets the inner HTML content
func (t *SpectrumTruncated) InnerHTML(innerHTML string) *SpectrumTruncated {
	t.innerHTML = innerHTML
	return t
}

// Child sets the child element
func (t *SpectrumTruncated) Child(child app.UI) *SpectrumTruncated {
	t.child = child
	return t
}

// Overflow sets the overflow content
func (t *SpectrumTruncated) Overflow(overflow app.UI) *SpectrumTruncated {
	t.overflow = overflow
	return t
}

// Children adds multiple child elements
func (t *SpectrumTruncated) Children(children ...app.UI) *SpectrumTruncated {
	t.children = append(t.children, children...)
	return t
}

// Render renders the truncated component
func (t *SpectrumTruncated) Render() app.UI {
	truncated := app.Elem("div").
		Class("spectrum-Truncated")

	// Set attributes
	if t.placement != "" {
		truncated = truncated.Attr("data-placement", t.placement)
	}

	// Set content
	if t.content != "" {
		truncated = truncated.Text(t.content)
	}
	if t.innerHTML != "" {
		truncated = truncated.Body(app.Raw(t.innerHTML))
	}

	// Add child if provided
	if t.child != nil {
		truncated = truncated.Body(t.child)
	}

	// Add overflow if provided
	if t.overflow != nil {
		truncated = truncated.Body(app.Elem("div").
			Class("spectrum-Truncated-overflow").
			Body(t.overflow))
	}

	// Add children if provided
	if len(t.children) > 0 {
		truncated = truncated.Body(t.children...)
	}

	return truncated
}

// SpectrumTruncatedManager represents a manager for truncated components
type SpectrumTruncatedManager struct {
	app.Compo

	// Properties
	components        map[string]*SpectrumTruncated
	onComponentChange app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumTruncatedManager creates a new truncated manager component
func NewSpectrumTruncatedManager() *SpectrumTruncatedManager {
	return &SpectrumTruncatedManager{
		components: make(map[string]*SpectrumTruncated),
	}
}

// AddComponent adds a truncated component
func (m *SpectrumTruncatedManager) AddComponent(name string, component *SpectrumTruncated) *SpectrumTruncatedManager {
	m.components[name] = component
	return m
}

// RemoveComponent removes a truncated component
func (m *SpectrumTruncatedManager) RemoveComponent(name string) *SpectrumTruncatedManager {
	delete(m.components, name)
	return m
}

// GetComponent gets a truncated component by name
func (m *SpectrumTruncatedManager) GetComponent(name string) *SpectrumTruncated {
	return m.components[name]
}

// ClearComponents clears all truncated components
func (m *SpectrumTruncatedManager) ClearComponents() *SpectrumTruncatedManager {
	m.components = make(map[string]*SpectrumTruncated)
	return m
}

// OnComponentChange sets the component change event handler
func (m *SpectrumTruncatedManager) OnComponentChange(handler app.EventHandler) *SpectrumTruncatedManager {
	m.onComponentChange = handler
	return m
}

// Child adds a child element
func (m *SpectrumTruncatedManager) Child(child app.UI) *SpectrumTruncatedManager {
	m.children = append(m.children, child)
	return m
}

// Children adds multiple child elements
func (m *SpectrumTruncatedManager) Children(children ...app.UI) *SpectrumTruncatedManager {
	m.children = append(m.children, children...)
	return m
}

// Render renders the truncated manager component
func (m *SpectrumTruncatedManager) Render() app.UI {
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
