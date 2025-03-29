package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumBreadcrumbItem represents an sp-breadcrumb-item component
type spectrumBreadcrumbItem struct {
	app.Compo

	// Properties
	PropDisabled       bool
	PropDownload       string
	PropHref           string
	PropLabel          string
	PropReferrerpolicy string
	PropRel            string
	PropTabIndex       int
	PropTarget         string
	PropValue          string
	PropSlot           string

	// Text content
	PropText string

	// Event handlers
	PropOnClick app.EventHandler
}

// BreadcrumbItem creates a new breadcrumb item component
func BreadcrumbItem() *spectrumBreadcrumbItem {
	return &spectrumBreadcrumbItem{}
}

// Disabled sets whether the breadcrumb item is disabled
func (bi *spectrumBreadcrumbItem) Disabled(disabled bool) *spectrumBreadcrumbItem {
	bi.PropDisabled = disabled
	return bi
}

// Download sets the download attribute
func (bi *spectrumBreadcrumbItem) Download(download string) *spectrumBreadcrumbItem {
	bi.PropDownload = download
	return bi
}

// Href sets the URL the breadcrumb item points to
func (bi *spectrumBreadcrumbItem) Href(href string) *spectrumBreadcrumbItem {
	bi.PropHref = href
	return bi
}

// Label sets the accessible label
func (bi *spectrumBreadcrumbItem) Label(label string) *spectrumBreadcrumbItem {
	bi.PropLabel = label
	return bi
}

// Referrerpolicy sets the referrer policy for the breadcrumb item
func (bi *spectrumBreadcrumbItem) Referrerpolicy(policy string) *spectrumBreadcrumbItem {
	bi.PropReferrerpolicy = policy
	return bi
}

// Rel sets the relationship of the linked URL
func (bi *spectrumBreadcrumbItem) Rel(rel string) *spectrumBreadcrumbItem {
	bi.PropRel = rel
	return bi
}

// TabIndex sets the tab index of the breadcrumb item
func (bi *spectrumBreadcrumbItem) TabIndex(tabIndex int) *spectrumBreadcrumbItem {
	bi.PropTabIndex = tabIndex
	return bi
}

// Target sets where to display the linked URL
func (bi *spectrumBreadcrumbItem) Target(target string) *spectrumBreadcrumbItem {
	bi.PropTarget = target
	return bi
}

// Value sets the value of the breadcrumb item
func (bi *spectrumBreadcrumbItem) Value(value string) *spectrumBreadcrumbItem {
	bi.PropValue = value
	return bi
}

// Slot sets the slot attribute of the breadcrumb item
func (bi *spectrumBreadcrumbItem) Slot(slot string) *spectrumBreadcrumbItem {
	bi.PropSlot = slot
	return bi
}

// Root sets the slot to "root"
func (bi *spectrumBreadcrumbItem) Root() *spectrumBreadcrumbItem {
	return bi.Slot("root")
}

// Text sets the text content of the breadcrumb item
func (bi *spectrumBreadcrumbItem) Text(text string) *spectrumBreadcrumbItem {
	bi.PropText = text
	return bi
}

// OnClick sets the click event handler
func (bi *spectrumBreadcrumbItem) OnClick(handler app.EventHandler) *spectrumBreadcrumbItem {
	bi.PropOnClick = handler
	return bi
}

// Render renders the breadcrumb item component
func (bi *spectrumBreadcrumbItem) Render() app.UI {
	breadcrumbItem := app.Elem("sp-breadcrumb-item")

	// Set attributes
	if bi.PropDisabled {
		breadcrumbItem = breadcrumbItem.Attr("disabled", true)
	}
	if bi.PropDownload != "" {
		breadcrumbItem = breadcrumbItem.Attr("download", bi.PropDownload)
	}
	if bi.PropHref != "" {
		breadcrumbItem = breadcrumbItem.Attr("href", bi.PropHref)
	}
	if bi.PropLabel != "" {
		breadcrumbItem = breadcrumbItem.Attr("label", bi.PropLabel)
	}
	if bi.PropReferrerpolicy != "" {
		breadcrumbItem = breadcrumbItem.Attr("referrerpolicy", bi.PropReferrerpolicy)
	}
	if bi.PropRel != "" {
		breadcrumbItem = breadcrumbItem.Attr("rel", bi.PropRel)
	}
	if bi.PropTabIndex != 0 {
		breadcrumbItem = breadcrumbItem.Attr("tabIndex", bi.PropTabIndex)
	}
	if bi.PropTarget != "" {
		breadcrumbItem = breadcrumbItem.Attr("target", bi.PropTarget)
	}
	if bi.PropValue != "" {
		breadcrumbItem = breadcrumbItem.Attr("value", bi.PropValue)
	}
	if bi.PropSlot != "" {
		breadcrumbItem = breadcrumbItem.Attr("slot", bi.PropSlot)
	}

	// Add event handlers
	if bi.PropOnClick != nil {
		breadcrumbItem = breadcrumbItem.On("click", bi.PropOnClick)
	}

	// Add the text content
	if bi.PropText != "" {
		breadcrumbItem = breadcrumbItem.Text(bi.PropText)
	}

	return breadcrumbItem
}
