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

// spectrumGrid represents a component that displays a virtualized grid
type spectrumGrid struct {
	app.Compo

	// Properties
	PropItems             []GridItem
	PropItemSize          int
	PropGap               int
	PropFocusableSelector string
	PropOnFocus           app.EventHandler
	PropOnBlur            app.EventHandler
	PropOnKeyDown         app.EventHandler
	PropOnScroll          app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumGrid creates a new grid component
func NewSpectrumGrid() *spectrumGrid {
	return &spectrumGrid{}
}

// Items sets the grid items
func (g *spectrumGrid) Items(items []GridItem) *spectrumGrid {
	g.PropItems = items
	return g
}

// ItemSize sets the size of each grid item
func (g *spectrumGrid) ItemSize(size int) *spectrumGrid {
	g.PropItemSize = size
	return g
}

// Gap sets the gap between grid items
func (g *spectrumGrid) Gap(gap int) *spectrumGrid {
	g.PropGap = gap
	return g
}

// FocusableSelector sets the selector for focusable elements
func (g *spectrumGrid) FocusableSelector(selector string) *spectrumGrid {
	g.PropFocusableSelector = selector
	return g
}

// OnFocus sets the focus event handler
func (g *spectrumGrid) OnFocus(handler app.EventHandler) *spectrumGrid {
	g.PropOnFocus = handler
	return g
}

// OnBlur sets the blur event handler
func (g *spectrumGrid) OnBlur(handler app.EventHandler) *spectrumGrid {
	g.PropOnBlur = handler
	return g
}

// OnKeyDown sets the key down event handler
func (g *spectrumGrid) OnKeyDown(handler app.EventHandler) *spectrumGrid {
	g.PropOnKeyDown = handler
	return g
}

// OnScroll sets the scroll event handler
func (g *spectrumGrid) OnScroll(handler app.EventHandler) *spectrumGrid {
	g.PropOnScroll = handler
	return g
}

// Child adds a child element
func (g *spectrumGrid) Child(child app.UI) *spectrumGrid {
	g.PropChildren = append(g.PropChildren, child)
	return g
}

// Children adds multiple child elements
func (g *spectrumGrid) Children(children ...app.UI) *spectrumGrid {
	g.PropChildren = append(g.PropChildren, children...)
	return g
}

// Render renders the grid component
func (g *spectrumGrid) Render() app.UI {
	grid := app.Elem("div").
		Class("spectrum-Grid")

	// Set attributes
	if g.PropItemSize > 0 {
		grid = grid.Attr("data-item-size", g.PropItemSize)
	}
	if g.PropGap > 0 {
		grid = grid.Attr("data-gap", g.PropGap)
	}
	if g.PropFocusableSelector != "" {
		grid = grid.Attr("data-focusable-selector", g.PropFocusableSelector)
	}

	// Add event handlers
	if g.PropOnFocus != nil {
		grid = grid.On("focus", g.PropOnFocus)
	}
	if g.PropOnBlur != nil {
		grid = grid.On("blur", g.PropOnBlur)
	}
	if g.PropOnKeyDown != nil {
		grid = grid.On("keydown", g.PropOnKeyDown)
	}
	if g.PropOnScroll != nil {
		grid = grid.On("scroll", g.PropOnScroll)
	}

	// Add items if provided
	if len(g.PropItems) > 0 {
		itemsContainer := app.Elem("div").
			Class("spectrum-Grid-items")
		for _, item := range g.PropItems {
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
	if len(g.PropChildren) > 0 {
		grid = grid.Body(g.PropChildren...)
	}

	return grid
}

// spectrumGridManager represents a manager for grid components
type spectrumGridManager struct {
	app.Compo

	// Properties
	PropComponents        map[string]*spectrumGrid
	PropOnComponentChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewSpectrumGridManager creates a new grid manager component
func NewSpectrumGridManager() *spectrumGridManager {
	return &spectrumGridManager{
		PropComponents: make(map[string]*spectrumGrid),
	}
}

// AddComponent adds a grid component
func (m *spectrumGridManager) AddComponent(name string, component *spectrumGrid) *spectrumGridManager {
	m.PropComponents[name] = component
	return m
}

// RemoveComponent removes a grid component
func (m *spectrumGridManager) RemoveComponent(name string) *spectrumGridManager {
	delete(m.PropComponents, name)
	return m
}

// GetComponent gets a grid component by name
func (m *spectrumGridManager) GetComponent(name string) *spectrumGrid {
	return m.PropComponents[name]
}

// ClearComponents clears all grid components
func (m *spectrumGridManager) ClearComponents() *spectrumGridManager {
	m.PropComponents = make(map[string]*spectrumGrid)
	return m
}

// OnComponentChange sets the component change event handler
func (m *spectrumGridManager) OnComponentChange(handler app.EventHandler) *spectrumGridManager {
	m.PropOnComponentChange = handler
	return m
}

// Child adds a child element
func (m *spectrumGridManager) Child(child app.UI) *spectrumGridManager {
	m.PropChildren = append(m.PropChildren, child)
	return m
}

// Children adds multiple child elements
func (m *spectrumGridManager) Children(children ...app.UI) *spectrumGridManager {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// Render renders the grid manager component
func (m *spectrumGridManager) Render() app.UI {
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
