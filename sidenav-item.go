package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumSidenavItem represents an sp-sidenav-item component
type SpectrumSidenavItem struct {
	app.Compo

	// Properties
	disabled       bool
	download       string
	expanded       bool
	href           string
	label          string
	referrerpolicy string
	rel            string
	selected       bool
	tabIndex       int
	target         string
	value          string

	// Slots
	icon     app.UI
	children []*SpectrumSidenavItem

	// Event handlers
	onClick app.EventHandler
}

// SidenavItem creates a new sidenav item component
func SidenavItem() *SpectrumSidenavItem {
	return &SpectrumSidenavItem{}
}

// Disabled sets whether the item is disabled
func (si *SpectrumSidenavItem) Disabled(disabled bool) *SpectrumSidenavItem {
	si.disabled = disabled
	return si
}

// Download sets the download attribute
func (si *SpectrumSidenavItem) Download(download string) *SpectrumSidenavItem {
	si.download = download
	return si
}

// Expanded sets whether the item is expanded to show child items
func (si *SpectrumSidenavItem) Expanded(expanded bool) *SpectrumSidenavItem {
	si.expanded = expanded
	return si
}

// Href sets the URL the item points to
func (si *SpectrumSidenavItem) Href(href string) *SpectrumSidenavItem {
	si.href = href
	return si
}

// Label sets the accessible label for the item
func (si *SpectrumSidenavItem) Label(label string) *SpectrumSidenavItem {
	si.label = label
	return si
}

// Referrerpolicy sets the referrer policy for the item
func (si *SpectrumSidenavItem) Referrerpolicy(policy string) *SpectrumSidenavItem {
	si.referrerpolicy = policy
	return si
}

// Rel sets the relationship of the linked URL
func (si *SpectrumSidenavItem) Rel(rel string) *SpectrumSidenavItem {
	si.rel = rel
	return si
}

// Selected sets whether the item is selected
func (si *SpectrumSidenavItem) Selected(selected bool) *SpectrumSidenavItem {
	si.selected = selected
	return si
}

// TabIndex sets the tab index of the item
func (si *SpectrumSidenavItem) TabIndex(tabIndex int) *SpectrumSidenavItem {
	si.tabIndex = tabIndex
	return si
}

// Target sets where to display the linked URL
func (si *SpectrumSidenavItem) Target(target string) *SpectrumSidenavItem {
	si.target = target
	return si
}

// Value sets the value of the item
func (si *SpectrumSidenavItem) Value(value string) *SpectrumSidenavItem {
	si.value = value
	return si
}

// Icon sets content for the icon slot
func (si *SpectrumSidenavItem) Icon(icon app.UI) *SpectrumSidenavItem {
	si.icon = icon
	return si
}

// Children sets the child sidenav items
func (si *SpectrumSidenavItem) Children(children ...*SpectrumSidenavItem) *SpectrumSidenavItem {
	si.children = children
	return si
}

// OnClick sets the click event handler
func (si *SpectrumSidenavItem) OnClick(handler app.EventHandler) *SpectrumSidenavItem {
	si.onClick = handler
	return si
}

// Render renders the sidenav item component
func (si *SpectrumSidenavItem) Render() app.UI {
	sidenavItem := app.Elem("sp-sidenav-item")

	// Set attributes
	if si.disabled {
		sidenavItem = sidenavItem.Attr("disabled", true)
	}
	if si.download != "" {
		sidenavItem = sidenavItem.Attr("download", si.download)
	}
	if si.expanded {
		sidenavItem = sidenavItem.Attr("expanded", true)
	}
	if si.href != "" {
		sidenavItem = sidenavItem.Attr("href", si.href)
	}
	if si.label != "" {
		sidenavItem = sidenavItem.Attr("label", si.label)
	}
	if si.referrerpolicy != "" {
		sidenavItem = sidenavItem.Attr("referrerpolicy", si.referrerpolicy)
	}
	if si.rel != "" {
		sidenavItem = sidenavItem.Attr("rel", si.rel)
	}
	if si.selected {
		sidenavItem = sidenavItem.Attr("selected", true)
	}
	if si.tabIndex != 0 {
		sidenavItem = sidenavItem.Attr("tabIndex", si.tabIndex)
	}
	if si.target != "" {
		sidenavItem = sidenavItem.Attr("target", si.target)
	}
	if si.value != "" {
		sidenavItem = sidenavItem.Attr("value", si.value)
	}

	// Add event handlers
	if si.onClick != nil {
		sidenavItem = sidenavItem.OnClick(si.onClick)
	}

	// Collect all elements
	var elements []app.UI

	// Add icon if provided
	if si.icon != nil {
		icon := si.icon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		elements = append(elements, icon)
	}

	// Add child items if provided
	if len(si.children) > 0 {
		// Convert SpectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range si.children {
			items = append(items, child)
		}

		// Append children to elements list
		elements = append(elements, items...)
	}

	// Add all elements to the sidenav item
	if len(elements) > 0 {
		sidenavItem = sidenavItem.Body(elements...)
	}

	return sidenavItem
}
