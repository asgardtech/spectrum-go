package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumBreadcrumbItem represents an sp-breadcrumb-item component
type SpectrumBreadcrumbItem struct {
	app.Compo

	// Properties
	disabled       bool
	download       string
	href           string
	label          string
	referrerpolicy string
	rel            string
	tabIndex       int
	target         string
	value          string
	slot           string

	// Text content
	text string

	// Event handlers
	onClick app.EventHandler
}

// BreadcrumbItem creates a new breadcrumb item component
func BreadcrumbItem() *SpectrumBreadcrumbItem {
	return &SpectrumBreadcrumbItem{}
}

// Disabled sets whether the breadcrumb item is disabled
func (bi *SpectrumBreadcrumbItem) Disabled(disabled bool) *SpectrumBreadcrumbItem {
	bi.disabled = disabled
	return bi
}

// Download sets the download attribute
func (bi *SpectrumBreadcrumbItem) Download(download string) *SpectrumBreadcrumbItem {
	bi.download = download
	return bi
}

// Href sets the URL the breadcrumb item points to
func (bi *SpectrumBreadcrumbItem) Href(href string) *SpectrumBreadcrumbItem {
	bi.href = href
	return bi
}

// Label sets the accessible label
func (bi *SpectrumBreadcrumbItem) Label(label string) *SpectrumBreadcrumbItem {
	bi.label = label
	return bi
}

// Referrerpolicy sets the referrer policy for the breadcrumb item
func (bi *SpectrumBreadcrumbItem) Referrerpolicy(policy string) *SpectrumBreadcrumbItem {
	bi.referrerpolicy = policy
	return bi
}

// Rel sets the relationship of the linked URL
func (bi *SpectrumBreadcrumbItem) Rel(rel string) *SpectrumBreadcrumbItem {
	bi.rel = rel
	return bi
}

// TabIndex sets the tab index of the breadcrumb item
func (bi *SpectrumBreadcrumbItem) TabIndex(tabIndex int) *SpectrumBreadcrumbItem {
	bi.tabIndex = tabIndex
	return bi
}

// Target sets where to display the linked URL
func (bi *SpectrumBreadcrumbItem) Target(target string) *SpectrumBreadcrumbItem {
	bi.target = target
	return bi
}

// Value sets the value of the breadcrumb item
func (bi *SpectrumBreadcrumbItem) Value(value string) *SpectrumBreadcrumbItem {
	bi.value = value
	return bi
}

// Slot sets the slot attribute of the breadcrumb item
func (bi *SpectrumBreadcrumbItem) Slot(slot string) *SpectrumBreadcrumbItem {
	bi.slot = slot
	return bi
}

// Root sets the slot to "root"
func (bi *SpectrumBreadcrumbItem) Root() *SpectrumBreadcrumbItem {
	return bi.Slot("root")
}

// Text sets the text content of the breadcrumb item
func (bi *SpectrumBreadcrumbItem) Text(text string) *SpectrumBreadcrumbItem {
	bi.text = text
	return bi
}

// OnClick sets the click event handler
func (bi *SpectrumBreadcrumbItem) OnClick(handler app.EventHandler) *SpectrumBreadcrumbItem {
	bi.onClick = handler
	return bi
}

// Render renders the breadcrumb item component
func (bi *SpectrumBreadcrumbItem) Render() app.UI {
	breadcrumbItem := app.Elem("sp-breadcrumb-item")

	// Set attributes
	if bi.disabled {
		breadcrumbItem = breadcrumbItem.Attr("disabled", true)
	}
	if bi.download != "" {
		breadcrumbItem = breadcrumbItem.Attr("download", bi.download)
	}
	if bi.href != "" {
		breadcrumbItem = breadcrumbItem.Attr("href", bi.href)
	}
	if bi.label != "" {
		breadcrumbItem = breadcrumbItem.Attr("label", bi.label)
	}
	if bi.referrerpolicy != "" {
		breadcrumbItem = breadcrumbItem.Attr("referrerpolicy", bi.referrerpolicy)
	}
	if bi.rel != "" {
		breadcrumbItem = breadcrumbItem.Attr("rel", bi.rel)
	}
	if bi.tabIndex != 0 {
		breadcrumbItem = breadcrumbItem.Attr("tabIndex", bi.tabIndex)
	}
	if bi.target != "" {
		breadcrumbItem = breadcrumbItem.Attr("target", bi.target)
	}
	if bi.value != "" {
		breadcrumbItem = breadcrumbItem.Attr("value", bi.value)
	}
	if bi.slot != "" {
		breadcrumbItem = breadcrumbItem.Attr("slot", bi.slot)
	}

	// Add event handlers
	if bi.onClick != nil {
		breadcrumbItem = breadcrumbItem.On("click", bi.onClick)
	}

	// Add the text content
	if bi.text != "" {
		breadcrumbItem = breadcrumbItem.Text(bi.text)
	}

	return breadcrumbItem
}
