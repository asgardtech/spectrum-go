package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSidenavItem represents an sp-sidenav-item component
type spectrumSidenavItem struct {
	app.Compo

	// Properties
	PropDisabled       bool
	PropDownload       string
	PropExpanded       bool
	PropHref           string
	PropLabel          string
	PropReferrerpolicy string
	PropRel            string
	PropSelected       bool
	PropTabIndex       int
	PropTarget         string
	PropValue          string

	// Slots
	PropIcon     app.UI
	PropChildren []*spectrumSidenavItem

	// Event handlers
	PropOnClick app.EventHandler
}

// SidenavItem creates a new sidenav item component
func SidenavItem() *spectrumSidenavItem {
	return &spectrumSidenavItem{}
}

// Disabled sets whether the item is disabled
func (si *spectrumSidenavItem) Disabled(disabled bool) *spectrumSidenavItem {
	si.PropDisabled = disabled
	return si
}

// Download sets the download attribute
func (si *spectrumSidenavItem) Download(download string) *spectrumSidenavItem {
	si.PropDownload = download
	return si
}

// Expanded sets whether the item is expanded to show child items
func (si *spectrumSidenavItem) Expanded(expanded bool) *spectrumSidenavItem {
	si.PropExpanded = expanded
	return si
}

// Href sets the URL the item points to
func (si *spectrumSidenavItem) Href(href string) *spectrumSidenavItem {
	si.PropHref = href
	return si
}

// Label sets the accessible label for the item
func (si *spectrumSidenavItem) Label(label string) *spectrumSidenavItem {
	si.PropLabel = label
	return si
}

// Referrerpolicy sets the referrer policy for the item
func (si *spectrumSidenavItem) Referrerpolicy(policy string) *spectrumSidenavItem {
	si.PropReferrerpolicy = policy
	return si
}

// Rel sets the relationship of the linked URL
func (si *spectrumSidenavItem) Rel(rel string) *spectrumSidenavItem {
	si.PropRel = rel
	return si
}

// Selected sets whether the item is selected
func (si *spectrumSidenavItem) Selected(selected bool) *spectrumSidenavItem {
	si.PropSelected = selected
	return si
}

// TabIndex sets the tab index of the item
func (si *spectrumSidenavItem) TabIndex(tabIndex int) *spectrumSidenavItem {
	si.PropTabIndex = tabIndex
	return si
}

// Target sets where to display the linked URL
func (si *spectrumSidenavItem) Target(target string) *spectrumSidenavItem {
	si.PropTarget = target
	return si
}

// Value sets the value of the item
func (si *spectrumSidenavItem) Value(value string) *spectrumSidenavItem {
	si.PropValue = value
	return si
}

// Icon sets content for the icon slot
func (si *spectrumSidenavItem) Icon(icon app.UI) *spectrumSidenavItem {
	si.PropIcon = icon
	return si
}

// Children sets the child sidenav items
func (si *spectrumSidenavItem) Children(children ...*spectrumSidenavItem) *spectrumSidenavItem {
	si.PropChildren = children
	return si
}

// OnClick sets the click event handler
func (si *spectrumSidenavItem) OnClick(handler app.EventHandler) *spectrumSidenavItem {
	si.PropOnClick = handler
	return si
}

// Render renders the sidenav item component
func (si *spectrumSidenavItem) Render() app.UI {
	sidenavItem := app.Elem("sp-sidenav-item")

	// Set attributes
	if si.PropDisabled {
		sidenavItem = sidenavItem.Attr("disabled", true)
	}
	if si.PropDownload != "" {
		sidenavItem = sidenavItem.Attr("download", si.PropDownload)
	}
	if si.PropExpanded {
		sidenavItem = sidenavItem.Attr("expanded", true)
	}
	if si.PropHref != "" {
		sidenavItem = sidenavItem.Attr("href", si.PropHref)
	}
	if si.PropLabel != "" {
		sidenavItem = sidenavItem.Attr("label", si.PropLabel)
	}
	if si.PropReferrerpolicy != "" {
		sidenavItem = sidenavItem.Attr("referrerpolicy", si.PropReferrerpolicy)
	}
	if si.PropRel != "" {
		sidenavItem = sidenavItem.Attr("rel", si.PropRel)
	}
	if si.PropSelected {
		sidenavItem = sidenavItem.Attr("selected", true)
	}
	if si.PropTabIndex != 0 {
		sidenavItem = sidenavItem.Attr("tabIndex", si.PropTabIndex)
	}
	if si.PropTarget != "" {
		sidenavItem = sidenavItem.Attr("target", si.PropTarget)
	}
	if si.PropValue != "" {
		sidenavItem = sidenavItem.Attr("value", si.PropValue)
	}

	// Add event handlers
	if si.PropOnClick != nil {
		sidenavItem = sidenavItem.On("click", si.PropOnClick)
	}

	// Collect all elements
	var elements []app.UI

	// Add icon if provided
	if si.PropIcon != nil {
		icon := si.PropIcon
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
	if len(si.PropChildren) > 0 {
		// Convert spectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range si.PropChildren {
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
