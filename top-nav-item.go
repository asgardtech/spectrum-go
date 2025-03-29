package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTopNavItem represents an sp-top-nav-item component
type spectrumTopNavItem struct {
	app.Compo

	// Properties
	PropDisabled       bool
	PropDownload       string
	PropHref           string
	PropLabel          string
	PropReferrerpolicy string
	PropRel            string
	PropSelected       bool
	PropTabIndex       int
	PropTarget         string
	PropValue          string

	// Content
	PropText string

	// Event handlers
	PropOnClick app.EventHandler
}

// TopNavItem creates a new top nav item component
func TopNavItem() *spectrumTopNavItem {
	return &spectrumTopNavItem{}
}

// Disabled sets whether the nav item is disabled
func (ni *spectrumTopNavItem) Disabled(disabled bool) *spectrumTopNavItem {
	ni.PropDisabled = disabled
	return ni
}

// Download sets the download attribute for the nav item
func (ni *spectrumTopNavItem) Download(download string) *spectrumTopNavItem {
	ni.PropDownload = download
	return ni
}

// Href sets the URL for the nav item
func (ni *spectrumTopNavItem) Href(href string) *spectrumTopNavItem {
	ni.PropHref = href
	return ni
}

// Label sets the accessible label for the nav item
func (ni *spectrumTopNavItem) Label(label string) *spectrumTopNavItem {
	ni.PropLabel = label
	return ni
}

// Referrerpolicy sets the referrer policy for the nav item
func (ni *spectrumTopNavItem) Referrerpolicy(policy string) *spectrumTopNavItem {
	ni.PropReferrerpolicy = policy
	return ni
}

// Rel sets the relationship of the linked URL
func (ni *spectrumTopNavItem) Rel(rel string) *spectrumTopNavItem {
	ni.PropRel = rel
	return ni
}

// Selected sets whether the nav item is selected
func (ni *spectrumTopNavItem) Selected(selected bool) *spectrumTopNavItem {
	ni.PropSelected = selected
	return ni
}

// TabIndex sets the tab index of the nav item
func (ni *spectrumTopNavItem) TabIndex(tabIndex int) *spectrumTopNavItem {
	ni.PropTabIndex = tabIndex
	return ni
}

// Target sets where to display the linked URL
func (ni *spectrumTopNavItem) Target(target string) *spectrumTopNavItem {
	ni.PropTarget = target
	return ni
}

// Value sets the value of the nav item
func (ni *spectrumTopNavItem) Value(value string) *spectrumTopNavItem {
	ni.PropValue = value
	return ni
}

// Text sets the text content of the nav item
func (ni *spectrumTopNavItem) Text(text string) *spectrumTopNavItem {
	ni.PropText = text
	return ni
}

// OnClick sets the click event handler
func (ni *spectrumTopNavItem) OnClick(handler app.EventHandler) *spectrumTopNavItem {
	ni.PropOnClick = handler
	return ni
}

// Render renders the top nav item component
func (ni *spectrumTopNavItem) Render() app.UI {
	navItem := app.Elem("sp-top-nav-item")

	// Set attributes
	if ni.PropDisabled {
		navItem = navItem.Attr("disabled", true)
	}
	if ni.PropDownload != "" {
		navItem = navItem.Attr("download", ni.PropDownload)
	}
	if ni.PropHref != "" {
		navItem = navItem.Attr("href", ni.PropHref)
	}
	if ni.PropLabel != "" {
		navItem = navItem.Attr("label", ni.PropLabel)
	}
	if ni.PropReferrerpolicy != "" {
		navItem = navItem.Attr("referrerpolicy", ni.PropReferrerpolicy)
	}
	if ni.PropRel != "" {
		navItem = navItem.Attr("rel", ni.PropRel)
	}
	if ni.PropSelected {
		navItem = navItem.Attr("selected", true)
	}
	if ni.PropTabIndex != 0 {
		navItem = navItem.Attr("tabindex", ni.PropTabIndex)
	}
	if ni.PropTarget != "" {
		navItem = navItem.Attr("target", ni.PropTarget)
	}
	if ni.PropValue != "" {
		navItem = navItem.Attr("value", ni.PropValue)
	}

	// Add event handlers
	if ni.PropOnClick != nil {
		navItem = navItem.On("click", ni.PropOnClick)
	}

	// Add text content if provided
	if ni.PropText != "" {
		navItem = navItem.Text(ni.PropText)
	}

	return navItem
}
