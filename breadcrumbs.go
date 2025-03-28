package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumBreadcrumbs represents an sp-breadcrumbs component
type SpectrumBreadcrumbs struct {
	app.Compo

	// Properties
	compact         bool
	label           string
	maxVisibleItems int
	menuLabel       string

	// Slots
	icon            app.UI
	rootItem        *SpectrumBreadcrumbItem
	breadcrumbItems []*SpectrumBreadcrumbItem

	// Event handlers
	onChange app.EventHandler
}

// Breadcrumbs creates a new breadcrumbs component
func Breadcrumbs() *SpectrumBreadcrumbs {
	return &SpectrumBreadcrumbs{
		maxVisibleItems: 4,            // Default maximum visible items
		menuLabel:       "More items", // Default menu label
	}
}

// Compact sets whether the breadcrumbs are displayed compactly
func (b *SpectrumBreadcrumbs) Compact(compact bool) *SpectrumBreadcrumbs {
	b.compact = compact
	return b
}

// Label sets the accessible label for the breadcrumbs
func (b *SpectrumBreadcrumbs) Label(label string) *SpectrumBreadcrumbs {
	b.label = label
	return b
}

// MaxVisibleItems sets the maximum number of visible breadcrumb items
func (b *SpectrumBreadcrumbs) MaxVisibleItems(max int) *SpectrumBreadcrumbs {
	b.maxVisibleItems = max
	return b
}

// MenuLabel sets the label for the action menu
func (b *SpectrumBreadcrumbs) MenuLabel(label string) *SpectrumBreadcrumbs {
	b.menuLabel = label
	return b
}

// Icon sets the content for the icon slot (replaces the menu icon)
func (b *SpectrumBreadcrumbs) Icon(icon app.UI) *SpectrumBreadcrumbs {
	b.icon = icon
	return b
}

// RootItem sets a breadcrumb item to always display (in the root slot)
func (b *SpectrumBreadcrumbs) RootItem(item *SpectrumBreadcrumbItem) *SpectrumBreadcrumbs {
	b.rootItem = item.Root()
	return b
}

// Items sets the breadcrumb items
func (b *SpectrumBreadcrumbs) Items(items ...*SpectrumBreadcrumbItem) *SpectrumBreadcrumbs {
	b.breadcrumbItems = items
	return b
}

// AddItem adds a single breadcrumb item
func (b *SpectrumBreadcrumbs) AddItem(item *SpectrumBreadcrumbItem) *SpectrumBreadcrumbs {
	b.breadcrumbItems = append(b.breadcrumbItems, item)
	return b
}

// OnChange sets the change event handler
func (b *SpectrumBreadcrumbs) OnChange(handler app.EventHandler) *SpectrumBreadcrumbs {
	b.onChange = handler
	return b
}

// Render renders the breadcrumbs component
func (b *SpectrumBreadcrumbs) Render() app.UI {
	breadcrumbs := app.Elem("sp-breadcrumbs")

	// Set attributes
	if b.compact {
		breadcrumbs = breadcrumbs.Attr("compact", true)
	}
	if b.label != "" {
		breadcrumbs = breadcrumbs.Attr("label", b.label)
	}
	if b.maxVisibleItems != 4 { // Only set if not the default value
		breadcrumbs = breadcrumbs.Attr("max-visible-items", b.maxVisibleItems)
	}
	if b.menuLabel != "More items" { // Only set if not the default value
		breadcrumbs = breadcrumbs.Attr("menu-label", b.menuLabel)
	}

	// Add event handlers
	if b.onChange != nil {
		breadcrumbs = breadcrumbs.On("change", b.onChange)
	}

	// Collect all elements
	var elements []app.UI

	// Add icon if provided
	if b.icon != nil {
		icon := b.icon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		elements = append(elements, icon)
	}

	// Add root item if provided
	if b.rootItem != nil {
		elements = append(elements, b.rootItem)
	}

	// Add breadcrumb items
	for _, item := range b.breadcrumbItems {
		elements = append(elements, item)
	}

	// Add all elements to the breadcrumbs
	if len(elements) > 0 {
		breadcrumbs = breadcrumbs.Body(elements...)
	}

	return breadcrumbs
}
