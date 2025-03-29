package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTabPanel represents a tab panel with content that is shown when its associated tab is selected
type spectrumTabPanel struct {
	app.Compo

	// Properties
	PropValue     string
	PropSelected  bool
	PropContent   string
	PropInnerHTML string
	PropChild     app.UI
}

// TabPanel creates a new tab panel
func TabPanel() *spectrumTabPanel {
	return &spectrumTabPanel{}
}

// Value sets the value of the tab panel, used to associate it with a tab
func (t *spectrumTabPanel) Value(value string) *spectrumTabPanel {
	t.PropValue = value
	return t
}

// Selected sets whether the tab panel is currently selected/visible
func (t *spectrumTabPanel) Selected(selected bool) *spectrumTabPanel {
	t.PropSelected = selected
	return t
}

// Content sets the text content of the tab panel
func (t *spectrumTabPanel) Content(content string) *spectrumTabPanel {
	t.PropContent = content
	return t
}

// InnerHTML sets the inner HTML of the tab panel
func (t *spectrumTabPanel) InnerHTML(html string) *spectrumTabPanel {
	t.PropInnerHTML = html
	return t
}

// Child sets a UI element as the child of the tab panel
func (t *spectrumTabPanel) Child(child app.UI) *spectrumTabPanel {
	t.PropChild = child
	return t
}

// Render renders the tab panel component
func (t *spectrumTabPanel) Render() app.UI {
	tabPanel := app.Elem("sp-tab-panel")

	if t.PropValue != "" {
		tabPanel = tabPanel.Attr("value", t.PropValue)
	}

	if t.PropSelected {
		tabPanel = tabPanel.Attr("selected", "")
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child
	if t.PropContent != "" {
		elements = append(elements, app.Text(t.PropContent))
	}
	if t.PropInnerHTML != "" {
		elements = append(elements, app.Raw(t.PropInnerHTML))
	}
	if t.PropChild != nil {
		elements = append(elements, t.PropChild)
	}

	// Add all elements to the tab panel
	if len(elements) > 0 {
		tabPanel = tabPanel.Body(elements...)
	}

	return tabPanel
}
