package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTab represents a tab in a tab list
type spectrumTab struct {
	app.Compo

	PropValue     string
	PropLabel     string
	PropDisabled  bool
	PropSelected  bool
	PropVertical  bool
	PropIcon      app.UI
	PropContent   string
	PropInnerHTML string
	PropChild     app.UI
}

// Tab creates a new tab
func Tab() *spectrumTab {
	return &spectrumTab{}
}

// Value sets the value of the tab
func (t *spectrumTab) Value(value string) *spectrumTab {
	t.PropValue = value
	return t
}

// Label sets the label of the tab
func (t *spectrumTab) Label(label string) *spectrumTab {
	t.PropLabel = label
	return t
}

// Disabled sets the disabled state of the tab
func (t *spectrumTab) Disabled(disabled bool) *spectrumTab {
	t.PropDisabled = disabled
	return t
}

// Selected sets the selected state of the tab
func (t *spectrumTab) Selected(selected bool) *spectrumTab {
	t.PropSelected = selected
	return t
}

// Vertical sets the vertical orientation of the tab
func (t *spectrumTab) Vertical(vertical bool) *spectrumTab {
	t.PropVertical = vertical
	return t
}

// Icon sets the icon of the tab
func (t *spectrumTab) Icon(icon app.UI) *spectrumTab {
	t.PropIcon = icon
	return t
}

// Content sets the content of the tab
func (t *spectrumTab) Content(content string) *spectrumTab {
	t.PropContent = content
	return t
}

// InnerHTML sets the inner HTML of the tab
func (t *spectrumTab) InnerHTML(innerHTML string) *spectrumTab {
	t.PropInnerHTML = innerHTML
	return t
}

// Child sets the child component of the tab
func (t *spectrumTab) Child(child app.UI) *spectrumTab {
	t.PropChild = child
	return t
}

// Render renders the tab component
func (t *spectrumTab) Render() app.UI {
	tab := app.Elem("sp-tab")

	if t.PropValue != "" {
		tab = tab.Attr("value", t.PropValue)
	}
	if t.PropLabel != "" {
		tab = tab.Attr("label", t.PropLabel)
	}
	if t.PropDisabled {
		tab = tab.Attr("disabled", "")
	}
	if t.PropSelected {
		tab = tab.Attr("selected", "")
	}
	if t.PropVertical {
		tab = tab.Attr("vertical", "")
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add icon if provided
	if t.PropIcon != nil {
		iconElem := app.Elem("div").
			Attr("slot", "icon").
			Body(t.PropIcon)
		elements = append(elements, iconElem)
	}

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

	// Add all elements to the tab
	if len(elements) > 0 {
		tab = tab.Body(elements...)
	}

	return tab
}
