package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTruncated represents a component that truncates text with tooltips
type spectrumTruncated struct {
	app.Compo

	// Properties
	PropPlacement string
	PropContent   string
	PropInnerHTML string
	PropChild     app.UI
	PropOverflow  app.UI

	// Children
	PropChildren []app.UI
}

// NewSpectrumTruncated creates a new truncated component
func NewSpectrumTruncated() *spectrumTruncated {
	return &spectrumTruncated{}
}

// Placement sets the tooltip placement
func (t *spectrumTruncated) Placement(placement string) *spectrumTruncated {
	t.PropPlacement = placement
	return t
}

// Content sets the content to be truncated
func (t *spectrumTruncated) Content(content string) *spectrumTruncated {
	t.PropContent = content
	return t
}

// InnerHTML sets the inner HTML content
func (t *spectrumTruncated) InnerHTML(innerHTML string) *spectrumTruncated {
	t.PropInnerHTML = innerHTML
	return t
}

// Child sets the child element
func (t *spectrumTruncated) Child(child app.UI) *spectrumTruncated {
	t.PropChild = child
	return t
}

// Overflow sets the overflow content
func (t *spectrumTruncated) Overflow(overflow app.UI) *spectrumTruncated {
	t.PropOverflow = overflow
	return t
}

// Children adds multiple child elements
func (t *spectrumTruncated) Children(children ...app.UI) *spectrumTruncated {
	t.PropChildren = append(t.PropChildren, children...)
	return t
}

// Render renders the truncated component
func (t *spectrumTruncated) Render() app.UI {
	truncated := app.Elem("div").
		Class("spectrum-Truncated")

	// Set attributes
	if t.PropPlacement != "" {
		truncated = truncated.Attr("data-placement", t.PropPlacement)
	}

	// Set content
	if t.PropContent != "" {
		truncated = truncated.Text(t.PropContent)
	}
	if t.PropInnerHTML != "" {
		truncated = truncated.Body(app.Raw(t.PropInnerHTML))
	}

	// Add child if provided
	if t.PropChild != nil {
		truncated = truncated.Body(t.PropChild)
	}

	// Add overflow if provided
	if t.PropOverflow != nil {
		truncated = truncated.Body(app.Elem("div").
			Class("spectrum-Truncated-overflow").
			Body(t.PropOverflow))
	}

	// Add children if provided
	if len(t.PropChildren) > 0 {
		truncated = truncated.Body(t.PropChildren...)
	}

	return truncated
}

// spectrumTruncatedManager represents a manager for truncated components
type spectrumTruncatedManager struct {
	app.Compo

	// Properties
	PropComponents        map[string]*spectrumTruncated
	PropOnComponentChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumTruncatedManager creates a new truncated manager component
func NewSpectrumTruncatedManager() *spectrumTruncatedManager {
	return &spectrumTruncatedManager{
		PropComponents: make(map[string]*spectrumTruncated),
	}
}

// AddComponent adds a truncated component
func (m *spectrumTruncatedManager) AddComponent(name string, component *spectrumTruncated) *spectrumTruncatedManager {
	m.PropComponents[name] = component
	return m
}

// RemoveComponent removes a truncated component
func (m *spectrumTruncatedManager) RemoveComponent(name string) *spectrumTruncatedManager {
	delete(m.PropComponents, name)
	return m
}

// GetComponent gets a truncated component by name
func (m *spectrumTruncatedManager) GetComponent(name string) *spectrumTruncated {
	return m.PropComponents[name]
}

// ClearComponents clears all truncated components
func (m *spectrumTruncatedManager) ClearComponents() *spectrumTruncatedManager {
	m.PropComponents = make(map[string]*spectrumTruncated)
	return m
}

// OnComponentChange sets the component change event handler
func (m *spectrumTruncatedManager) OnComponentChange(handler app.EventHandler) *spectrumTruncatedManager {
	m.PropOnComponentChange = handler
	return m
}

// Child adds a child element
func (m *spectrumTruncatedManager) Child(child app.UI) *spectrumTruncatedManager {
	m.PropChildren = append(m.PropChildren, child)
	return m
}

// Children adds multiple child elements
func (m *spectrumTruncatedManager) Children(children ...app.UI) *spectrumTruncatedManager {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// Render renders the truncated manager component
func (m *spectrumTruncatedManager) Render() app.UI {
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
