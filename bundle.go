package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumBundle represents a bundle of Spectrum components
type SpectrumBundle struct {
	app.Compo

	// Properties
	name         string
	version      string
	components   []app.UI
	onBundleLoad app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumBundle creates a new Spectrum bundle component
func NewSpectrumBundle() *SpectrumBundle {
	return &SpectrumBundle{}
}

// Name sets the bundle name
func (b *SpectrumBundle) Name(name string) *SpectrumBundle {
	b.name = name
	return b
}

// Version sets the bundle version
func (b *SpectrumBundle) Version(version string) *SpectrumBundle {
	b.version = version
	return b
}

// AddComponent adds a component to the bundle
func (b *SpectrumBundle) AddComponent(component app.UI) *SpectrumBundle {
	b.components = append(b.components, component)
	return b
}

// AddComponents adds multiple components to the bundle
func (b *SpectrumBundle) AddComponents(components ...app.UI) *SpectrumBundle {
	b.components = append(b.components, components...)
	return b
}

// ClearComponents clears all components from the bundle
func (b *SpectrumBundle) ClearComponents() *SpectrumBundle {
	b.components = nil
	return b
}

// OnBundleLoad sets the bundle load event handler
func (b *SpectrumBundle) OnBundleLoad(handler app.EventHandler) *SpectrumBundle {
	b.onBundleLoad = handler
	return b
}

// Child adds a child element
func (b *SpectrumBundle) Child(child app.UI) *SpectrumBundle {
	b.children = append(b.children, child)
	return b
}

// Children adds multiple child elements
func (b *SpectrumBundle) Children(children ...app.UI) *SpectrumBundle {
	b.children = append(b.children, children...)
	return b
}

// Render renders the Spectrum bundle component
func (b *SpectrumBundle) Render() app.UI {
	bundle := app.Elem("div")

	// Set attributes
	if b.name != "" {
		bundle = bundle.Attr("data-bundle-name", b.name)
	}
	if b.version != "" {
		bundle = bundle.Attr("data-bundle-version", b.version)
	}

	// Add event handler
	if b.onBundleLoad != nil {
		bundle = bundle.On("bundle-load", b.onBundleLoad)
	}

	// Add components if provided
	if len(b.components) > 0 {
		bundle = bundle.Body(b.components...)
	}

	// Add children if provided
	if len(b.children) > 0 {
		bundle = bundle.Body(b.children...)
	}

	return bundle
}

// BundleManager represents a manager for Spectrum bundles
type BundleManager struct {
	app.Compo

	// Properties
	bundles        map[string]*SpectrumBundle
	onBundleChange app.EventHandler

	// Children
	children []app.UI
}

// NewBundleManager creates a new bundle manager component
func NewBundleManager() *BundleManager {
	return &BundleManager{
		bundles: make(map[string]*SpectrumBundle),
	}
}

// AddBundle adds a bundle to the manager
func (m *BundleManager) AddBundle(bundle *SpectrumBundle) *BundleManager {
	if bundle.name != "" {
		m.bundles[bundle.name] = bundle
	}
	return m
}

// RemoveBundle removes a bundle from the manager
func (m *BundleManager) RemoveBundle(name string) *BundleManager {
	delete(m.bundles, name)
	return m
}

// GetBundle gets a bundle by name
func (m *BundleManager) GetBundle(name string) *SpectrumBundle {
	return m.bundles[name]
}

// ClearBundles clears all bundles from the manager
func (m *BundleManager) ClearBundles() *BundleManager {
	m.bundles = make(map[string]*SpectrumBundle)
	return m
}

// OnBundleChange sets the bundle change event handler
func (m *BundleManager) OnBundleChange(handler app.EventHandler) *BundleManager {
	m.onBundleChange = handler
	return m
}

// Child adds a child element
func (m *BundleManager) Child(child app.UI) *BundleManager {
	m.children = append(m.children, child)
	return m
}

// Children adds multiple child elements
func (m *BundleManager) Children(children ...app.UI) *BundleManager {
	m.children = append(m.children, children...)
	return m
}

// Render renders the bundle manager component
func (m *BundleManager) Render() app.UI {
	manager := app.Elem("div")

	// Add bundles
	for _, bundle := range m.bundles {
		manager = manager.Body(bundle)
	}

	// Add event handler
	if m.onBundleChange != nil {
		manager = manager.On("bundle-change", m.onBundleChange)
	}

	// Add children if provided
	if len(m.children) > 0 {
		manager = manager.Body(m.children...)
	}

	return manager
}
