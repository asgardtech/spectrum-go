package sp

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// GridItem represents an item in the grid
type GridItem struct {
	// Data to be rendered
	Data interface{}
	// Width of the item in pixels
	Width int
	// Height of the item in pixels
	Height int
}

// SpectrumGrid represents a component that displays a virtualized grid
type SpectrumGrid struct {
	app.Compo

	// Properties
	items             []GridItem
	itemSize          int
	gap               int
	focusableSelector string
	onFocus           app.EventHandler
	onBlur            app.EventHandler
	onKeyDown         app.EventHandler
	onScroll          app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumGrid creates a new grid component
func NewSpectrumGrid() *SpectrumGrid {
	return &SpectrumGrid{}
}

// Items sets the grid items
func (g *SpectrumGrid) Items(items []GridItem) *SpectrumGrid {
	g.items = items
	return g
}

// ItemSize sets the size of each grid item
func (g *SpectrumGrid) ItemSize(size int) *SpectrumGrid {
	g.itemSize = size
	return g
}

// Gap sets the gap between grid items
func (g *SpectrumGrid) Gap(gap int) *SpectrumGrid {
	g.gap = gap
	return g
}

// FocusableSelector sets the selector for focusable elements
func (g *SpectrumGrid) FocusableSelector(selector string) *SpectrumGrid {
	g.focusableSelector = selector
	return g
}

// OnFocus sets the focus event handler
func (g *SpectrumGrid) OnFocus(handler app.EventHandler) *SpectrumGrid {
	g.onFocus = handler
	return g
}

// OnBlur sets the blur event handler
func (g *SpectrumGrid) OnBlur(handler app.EventHandler) *SpectrumGrid {
	g.onBlur = handler
	return g
}

// OnKeyDown sets the key down event handler
func (g *SpectrumGrid) OnKeyDown(handler app.EventHandler) *SpectrumGrid {
	g.onKeyDown = handler
	return g
}

// OnScroll sets the scroll event handler
func (g *SpectrumGrid) OnScroll(handler app.EventHandler) *SpectrumGrid {
	g.onScroll = handler
	return g
}

// Child adds a child element
func (g *SpectrumGrid) Child(child app.UI) *SpectrumGrid {
	g.children = append(g.children, child)
	return g
}

// Children adds multiple child elements
func (g *SpectrumGrid) Children(children ...app.UI) *SpectrumGrid {
	g.children = append(g.children, children...)
	return g
}

// Render renders the grid component
func (g *SpectrumGrid) Render() app.UI {
	grid := app.Elem("div").
		Class("spectrum-Grid")

	// Set attributes
	if g.itemSize > 0 {
		grid = grid.Attr("data-item-size", g.itemSize)
	}
	if g.gap > 0 {
		grid = grid.Attr("data-gap", g.gap)
	}
	if g.focusableSelector != "" {
		grid = grid.Attr("data-focusable-selector", g.focusableSelector)
	}

	// Add event handlers
	if g.onFocus != nil {
		grid = grid.On("focus", g.onFocus)
	}
	if g.onBlur != nil {
		grid = grid.On("blur", g.onBlur)
	}
	if g.onKeyDown != nil {
		grid = grid.On("keydown", g.onKeyDown)
	}
	if g.onScroll != nil {
		grid = grid.On("scroll", g.onScroll)
	}

	// Add items if provided
	if len(g.items) > 0 {
		itemsContainer := app.Elem("div").
			Class("spectrum-Grid-items")
		for _, item := range g.items {
			itemElem := app.Elem("div").
				Class("spectrum-Grid-item")
			if item.Width > 0 {
				itemElem = itemElem.Style("width", fmt.Sprintf("%dpx", item.Width))
			}
			if item.Height > 0 {
				itemElem = itemElem.Style("height", fmt.Sprintf("%dpx", item.Height))
			}
			if item.Data != nil {
				switch data := item.Data.(type) {
				case string:
					itemElem = itemElem.Body(app.Raw(data))
				case app.UI:
					itemElem = itemElem.Body(data)
				default:
					itemElem = itemElem.Body(app.Text(fmt.Sprintf("%v", data)))
				}
			}
			itemsContainer = itemsContainer.Body(itemElem)
		}
		grid = grid.Body(itemsContainer)
	}

	// Add children if provided
	if len(g.children) > 0 {
		grid = grid.Body(g.children...)
	}

	return grid
}

// SpectrumGridManager represents a manager for grid components
type SpectrumGridManager struct {
	app.Compo

	// Properties
	components        map[string]*SpectrumGrid
	onComponentChange app.EventHandler

	// Children
	children []app.UI
}

// NewSpectrumGridManager creates a new grid manager component
func NewSpectrumGridManager() *SpectrumGridManager {
	return &SpectrumGridManager{
		components: make(map[string]*SpectrumGrid),
	}
}

// AddComponent adds a grid component
func (m *SpectrumGridManager) AddComponent(name string, component *SpectrumGrid) *SpectrumGridManager {
	m.components[name] = component
	return m
}

// RemoveComponent removes a grid component
func (m *SpectrumGridManager) RemoveComponent(name string) *SpectrumGridManager {
	delete(m.components, name)
	return m
}

// GetComponent gets a grid component by name
func (m *SpectrumGridManager) GetComponent(name string) *SpectrumGrid {
	return m.components[name]
}

// ClearComponents clears all grid components
func (m *SpectrumGridManager) ClearComponents() *SpectrumGridManager {
	m.components = make(map[string]*SpectrumGrid)
	return m
}

// OnComponentChange sets the component change event handler
func (m *SpectrumGridManager) OnComponentChange(handler app.EventHandler) *SpectrumGridManager {
	m.onComponentChange = handler
	return m
}

// Child adds a child element
func (m *SpectrumGridManager) Child(child app.UI) *SpectrumGridManager {
	m.children = append(m.children, child)
	return m
}

// Children adds multiple child elements
func (m *SpectrumGridManager) Children(children ...app.UI) *SpectrumGridManager {
	m.children = append(m.children, children...)
	return m
}

// Render renders the grid manager component
func (m *SpectrumGridManager) Render() app.UI {
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
