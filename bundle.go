package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumBundle represents a bundle of Spectrum components
type spectrumBundle struct {
	app.Compo

	// Properties
	PropName         string
	PropVersion      string
	PropComponents   []app.UI
	PropOnBundleLoad app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumBundle creates a new Spectrum bundle component
func NewSpectrumBundle() *spectrumBundle {
	return &spectrumBundle{}
}

// Name sets the bundle name
func (b *spectrumBundle) Name(name string) *spectrumBundle {
	b.PropName = name
	return b
}

// Version sets the bundle version
func (b *spectrumBundle) Version(version string) *spectrumBundle {
	b.PropVersion = version
	return b
}

// AddComponent adds a component to the bundle
func (b *spectrumBundle) AddComponent(component app.UI) *spectrumBundle {
	b.PropComponents = append(b.PropComponents, component)
	return b
}

// AddComponents adds multiple components to the bundle
func (b *spectrumBundle) AddComponents(components ...app.UI) *spectrumBundle {
	b.PropComponents = append(b.PropComponents, components...)
	return b
}

// ClearComponents clears all components from the bundle
func (b *spectrumBundle) ClearComponents() *spectrumBundle {
	b.PropComponents = nil
	return b
}

// OnBundleLoad sets the bundle load event handler
func (b *spectrumBundle) OnBundleLoad(handler app.EventHandler) *spectrumBundle {
	b.PropOnBundleLoad = handler
	return b
}

// Child adds a child element
func (b *spectrumBundle) Child(child app.UI) *spectrumBundle {
	b.PropChildren = append(b.PropChildren, child)
	return b
}

// Children adds multiple child elements
func (b *spectrumBundle) Children(children ...app.UI) *spectrumBundle {
	b.PropChildren = append(b.PropChildren, children...)
	return b
}

// Render renders the Spectrum bundle component
func (b *spectrumBundle) Render() app.UI {
	bundle := app.Elem("div")

	// Set attributes
	if b.PropName != "" {
		bundle = bundle.Attr("data-bundle-name", b.PropName)
	}
	if b.PropVersion != "" {
		bundle = bundle.Attr("data-bundle-version", b.PropVersion)
	}

	// Add event handler
	if b.PropOnBundleLoad != nil {
		bundle = bundle.On("bundle-load", b.PropOnBundleLoad)
	}

	// Add components if provided
	if len(b.PropComponents) > 0 {
		bundle = bundle.Body(b.PropComponents...)
	}

	// Add children if provided
	if len(b.PropChildren) > 0 {
		bundle = bundle.Body(b.PropChildren...)
	}

	return bundle
}

// spectrumBundleManager represents a manager for Spectrum bundles
type spectrumBundleManager struct {
	app.Compo

	// Properties
	PropBundles        map[string]*spectrumBundle
	PropOnBundleChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewBundleManager creates a new bundle manager component
func NewBundleManager() *spectrumBundleManager {
	return &spectrumBundleManager{
		PropBundles: make(map[string]*spectrumBundle),
	}
}

// AddBundle adds a bundle to the manager
func (m *spectrumBundleManager) AddBundle(bundle *spectrumBundle) *spectrumBundleManager {
	if bundle.PropName != "" {
		m.PropBundles[bundle.PropName] = bundle
	}
	return m
}

// RemoveBundle removes a bundle from the manager
func (m *spectrumBundleManager) RemoveBundle(name string) *spectrumBundleManager {
	delete(m.PropBundles, name)
	return m
}

// GetBundle gets a bundle by name
func (m *spectrumBundleManager) GetBundle(name string) *spectrumBundle {
	return m.PropBundles[name]
}

// ClearBundles clears all bundles from the manager
func (m *spectrumBundleManager) ClearBundles() *spectrumBundleManager {
	m.PropBundles = make(map[string]*spectrumBundle)
	return m
}

// OnBundleChange sets the bundle change event handler
func (m *spectrumBundleManager) OnBundleChange(handler app.EventHandler) *spectrumBundleManager {
	m.PropOnBundleChange = handler
	return m
}

// Child adds a child element
func (m *spectrumBundleManager) Child(child app.UI) *spectrumBundleManager {
	m.PropChildren = append(m.PropChildren, child)
	return m
}

// Children adds multiple child elements
func (m *spectrumBundleManager) Children(children ...app.UI) *spectrumBundleManager {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// Render renders the bundle manager component
func (m *spectrumBundleManager) Render() app.UI {
	manager := app.Elem("div")

	// Add bundles
	for _, bundle := range m.PropBundles {
		manager = manager.Body(bundle)
	}

	// Add event handler
	if m.PropOnBundleChange != nil {
		manager = manager.On("bundle-change", m.PropOnBundleChange)
	}

	// Add children if provided
	if len(m.PropChildren) > 0 {
		manager = manager.Body(m.PropChildren...)
	}

	return manager
}
