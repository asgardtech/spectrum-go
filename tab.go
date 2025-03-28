package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTab represents a tab in a tab list
type SpectrumTab struct {
	app.Compo

	value     string
	label     string
	disabled  bool
	selected  bool
	vertical  bool
	icon      app.UI
	content   string
	innerHTML string
	child     app.UI
}

// Tab creates a new tab
func Tab() *SpectrumTab {
	return &SpectrumTab{}
}

// Value sets the value of the tab
func (t *SpectrumTab) Value(value string) *SpectrumTab {
	t.value = value
	return t
}

// Label sets the label of the tab
func (t *SpectrumTab) Label(label string) *SpectrumTab {
	t.label = label
	return t
}

// Disabled sets the disabled state of the tab
func (t *SpectrumTab) Disabled(disabled bool) *SpectrumTab {
	t.disabled = disabled
	return t
}

// Selected sets the selected state of the tab
func (t *SpectrumTab) Selected(selected bool) *SpectrumTab {
	t.selected = selected
	return t
}

// Vertical sets the vertical orientation of the tab
func (t *SpectrumTab) Vertical(vertical bool) *SpectrumTab {
	t.vertical = vertical
	return t
}

// Icon sets the icon of the tab
func (t *SpectrumTab) Icon(icon app.UI) *SpectrumTab {
	t.icon = icon
	return t
}

// Content sets the content of the tab
func (t *SpectrumTab) Content(content string) *SpectrumTab {
	t.content = content
	return t
}

// InnerHTML sets the inner HTML of the tab
func (t *SpectrumTab) InnerHTML(innerHTML string) *SpectrumTab {
	t.innerHTML = innerHTML
	return t
}

// Child sets the child component of the tab
func (t *SpectrumTab) Child(child app.UI) *SpectrumTab {
	t.child = child
	return t
}

// Render renders the tab component
func (t *SpectrumTab) Render() app.UI {
	tab := app.Elem("sp-tab")

	if t.value != "" {
		tab = tab.Attr("value", t.value)
	}
	if t.label != "" {
		tab = tab.Attr("label", t.label)
	}
	if t.disabled {
		tab = tab.Attr("disabled", "")
	}
	if t.selected {
		tab = tab.Attr("selected", "")
	}
	if t.vertical {
		tab = tab.Attr("vertical", "")
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add icon if provided
	if t.icon != nil {
		iconElem := app.Elem("div").
			Attr("slot", "icon").
			Body(t.icon)
		elements = append(elements, iconElem)
	}

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

	// Add all elements to the tab
	if len(elements) > 0 {
		tab = tab.Body(elements...)
	}

	return tab
}
