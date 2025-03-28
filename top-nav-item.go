package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTopNavItem represents an sp-top-nav-item component
type SpectrumTopNavItem struct {
	app.Compo

	// Properties
	disabled       bool
	download       string
	href           string
	label          string
	referrerpolicy string
	rel            string
	selected       bool
	tabIndex       int
	target         string
	value          string

	// Content
	text string

	// Event handlers
	onClick app.EventHandler
}

// TopNavItem creates a new top nav item component
func TopNavItem() *SpectrumTopNavItem {
	return &SpectrumTopNavItem{}
}

// Disabled sets whether the nav item is disabled
func (ni *SpectrumTopNavItem) Disabled(disabled bool) *SpectrumTopNavItem {
	ni.disabled = disabled
	return ni
}

// Download sets the download attribute for the nav item
func (ni *SpectrumTopNavItem) Download(download string) *SpectrumTopNavItem {
	ni.download = download
	return ni
}

// Href sets the URL for the nav item
func (ni *SpectrumTopNavItem) Href(href string) *SpectrumTopNavItem {
	ni.href = href
	return ni
}

// Label sets the accessible label for the nav item
func (ni *SpectrumTopNavItem) Label(label string) *SpectrumTopNavItem {
	ni.label = label
	return ni
}

// Referrerpolicy sets the referrer policy for the nav item
func (ni *SpectrumTopNavItem) Referrerpolicy(policy string) *SpectrumTopNavItem {
	ni.referrerpolicy = policy
	return ni
}

// Rel sets the relationship of the linked URL
func (ni *SpectrumTopNavItem) Rel(rel string) *SpectrumTopNavItem {
	ni.rel = rel
	return ni
}

// Selected sets whether the nav item is selected
func (ni *SpectrumTopNavItem) Selected(selected bool) *SpectrumTopNavItem {
	ni.selected = selected
	return ni
}

// TabIndex sets the tab index of the nav item
func (ni *SpectrumTopNavItem) TabIndex(tabIndex int) *SpectrumTopNavItem {
	ni.tabIndex = tabIndex
	return ni
}

// Target sets where to display the linked URL
func (ni *SpectrumTopNavItem) Target(target string) *SpectrumTopNavItem {
	ni.target = target
	return ni
}

// Value sets the value of the nav item
func (ni *SpectrumTopNavItem) Value(value string) *SpectrumTopNavItem {
	ni.value = value
	return ni
}

// Text sets the text content of the nav item
func (ni *SpectrumTopNavItem) Text(text string) *SpectrumTopNavItem {
	ni.text = text
	return ni
}

// OnClick sets the click event handler
func (ni *SpectrumTopNavItem) OnClick(handler app.EventHandler) *SpectrumTopNavItem {
	ni.onClick = handler
	return ni
}

// Render renders the top nav item component
func (ni *SpectrumTopNavItem) Render() app.UI {
	navItem := app.Elem("sp-top-nav-item")

	// Set attributes
	if ni.disabled {
		navItem = navItem.Attr("disabled", true)
	}
	if ni.download != "" {
		navItem = navItem.Attr("download", ni.download)
	}
	if ni.href != "" {
		navItem = navItem.Attr("href", ni.href)
	}
	if ni.label != "" {
		navItem = navItem.Attr("label", ni.label)
	}
	if ni.referrerpolicy != "" {
		navItem = navItem.Attr("referrerpolicy", ni.referrerpolicy)
	}
	if ni.rel != "" {
		navItem = navItem.Attr("rel", ni.rel)
	}
	if ni.selected {
		navItem = navItem.Attr("selected", true)
	}
	if ni.tabIndex != 0 {
		navItem = navItem.Attr("tabindex", ni.tabIndex)
	}
	if ni.target != "" {
		navItem = navItem.Attr("target", ni.target)
	}
	if ni.value != "" {
		navItem = navItem.Attr("value", ni.value)
	}

	// Add event handlers
	if ni.onClick != nil {
		navItem = navItem.On("click", ni.onClick)
	}

	// Add text content if provided
	if ni.text != "" {
		navItem = navItem.Text(ni.text)
	}

	return navItem
}
