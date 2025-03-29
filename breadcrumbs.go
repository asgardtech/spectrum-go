package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumBreadcrumbs represents an sp-breadcrumbs component
type spectrumBreadcrumbs struct {
	app.Compo

	// Properties
	PropCompact         bool
	PropLabel           string
	PropMaxVisibleItems int
	PropMenuLabel       string

	// Slots
	PropIcon            app.UI
	PropRootItem        *spectrumBreadcrumbItem
	PropBreadcrumbItems []*spectrumBreadcrumbItem

	// Event handlers
	PropOnChange app.EventHandler
}

// Breadcrumbs creates a new breadcrumbs component
func Breadcrumbs() *spectrumBreadcrumbs {
	return &spectrumBreadcrumbs{
		PropMaxVisibleItems: 4,            // Default maximum visible items
		PropMenuLabel:       "More items", // Default menu label
	}
}

// Compact sets whether the breadcrumbs are displayed compactly
func (b *spectrumBreadcrumbs) Compact(compact bool) *spectrumBreadcrumbs {
	b.PropCompact = compact
	return b
}

// Label sets the accessible label for the breadcrumbs
func (b *spectrumBreadcrumbs) Label(label string) *spectrumBreadcrumbs {
	b.PropLabel = label
	return b
}

// MaxVisibleItems sets the maximum number of visible breadcrumb items
func (b *spectrumBreadcrumbs) MaxVisibleItems(max int) *spectrumBreadcrumbs {
	b.PropMaxVisibleItems = max
	return b
}

// MenuLabel sets the label for the action menu
func (b *spectrumBreadcrumbs) MenuLabel(label string) *spectrumBreadcrumbs {
	b.PropMenuLabel = label
	return b
}

// Icon sets the content for the icon slot (replaces the menu icon)
func (b *spectrumBreadcrumbs) Icon(icon app.UI) *spectrumBreadcrumbs {
	b.PropIcon = icon
	return b
}

// RootItem sets a breadcrumb item to always display (in the root slot)
func (b *spectrumBreadcrumbs) RootItem(item *spectrumBreadcrumbItem) *spectrumBreadcrumbs {
	b.PropRootItem = item.Root()
	return b
}

// Items sets the breadcrumb items
func (b *spectrumBreadcrumbs) Items(items ...*spectrumBreadcrumbItem) *spectrumBreadcrumbs {
	b.PropBreadcrumbItems = items
	return b
}

// AddItem adds a single breadcrumb item
func (b *spectrumBreadcrumbs) AddItem(item *spectrumBreadcrumbItem) *spectrumBreadcrumbs {
	b.PropBreadcrumbItems = append(b.PropBreadcrumbItems, item)
	return b
}

// OnChange sets the change event handler
func (b *spectrumBreadcrumbs) OnChange(handler app.EventHandler) *spectrumBreadcrumbs {
	b.PropOnChange = handler
	return b
}

// Render renders the breadcrumbs component
func (b *spectrumBreadcrumbs) Render() app.UI {
	breadcrumbs := app.Elem("sp-breadcrumbs")

	// Set attributes
	if b.PropCompact {
		breadcrumbs = breadcrumbs.Attr("compact", true)
	}
	if b.PropLabel != "" {
		breadcrumbs = breadcrumbs.Attr("label", b.PropLabel)
	}
	if b.PropMaxVisibleItems != 4 { // Only set if not the default value
		breadcrumbs = breadcrumbs.Attr("max-visible-items", b.PropMaxVisibleItems)
	}
	if b.PropMenuLabel != "More items" { // Only set if not the default value
		breadcrumbs = breadcrumbs.Attr("menu-label", b.PropMenuLabel)
	}

	// Add event handlers
	if b.PropOnChange != nil {
		breadcrumbs = breadcrumbs.On("change", b.PropOnChange)
	}

	// Collect all elements
	var elements []app.UI

	// Add icon if provided
	if b.PropIcon != nil {
		icon := b.PropIcon
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
	if b.PropRootItem != nil {
		elements = append(elements, b.PropRootItem)
	}

	// Add breadcrumb items
	for _, item := range b.PropBreadcrumbItems {
		elements = append(elements, item)
	}

	// Add all elements to the breadcrumbs
	if len(elements) > 0 {
		breadcrumbs = breadcrumbs.Body(elements...)
	}

	return breadcrumbs
}
