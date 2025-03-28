package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTabPanel represents a tab panel with content that is shown when its associated tab is selected
type SpectrumTabPanel struct {
	app.Compo

	// Properties
	value     string
	selected  bool
	content   string
	innerHTML string
	child     app.UI
}

// TabPanel creates a new tab panel
func TabPanel() *SpectrumTabPanel {
	return &SpectrumTabPanel{}
}

// Value sets the value of the tab panel, used to associate it with a tab
func (t *SpectrumTabPanel) Value(value string) *SpectrumTabPanel {
	t.value = value
	return t
}

// Selected sets whether the tab panel is currently selected/visible
func (t *SpectrumTabPanel) Selected(selected bool) *SpectrumTabPanel {
	t.selected = selected
	return t
}

// Content sets the text content of the tab panel
func (t *SpectrumTabPanel) Content(content string) *SpectrumTabPanel {
	t.content = content
	return t
}

// InnerHTML sets the inner HTML of the tab panel
func (t *SpectrumTabPanel) InnerHTML(html string) *SpectrumTabPanel {
	t.innerHTML = html
	return t
}

// Child sets a UI element as the child of the tab panel
func (t *SpectrumTabPanel) Child(child app.UI) *SpectrumTabPanel {
	t.child = child
	return t
}

// Render renders the tab panel component
func (t *SpectrumTabPanel) Render() app.UI {
	tabPanel := app.Elem("sp-tab-panel")

	if t.value != "" {
		tabPanel = tabPanel.Attr("value", t.value)
	}

	if t.selected {
		tabPanel = tabPanel.Attr("selected", "")
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child
	if t.content != "" {
		elements = append(elements, app.Text(t.content))
	}
	if t.innerHTML != "" {
		elements = append(elements, app.Raw(t.innerHTML))
	}
	if t.child != nil {
		elements = append(elements, t.child)
	}

	// Add all elements to the tab panel
	if len(elements) > 0 {
		tabPanel = tabPanel.Body(elements...)
	}

	return tabPanel
}
