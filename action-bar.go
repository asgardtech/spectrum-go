package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionBarVariant represents the variant of an action bar
type ActionBarVariant string

// Action bar variants
const (
	ActionBarVariantFixed  ActionBarVariant = "fixed"
	ActionBarVariantSticky ActionBarVariant = "sticky"
)

// spectrumActionBar represents an sp-action-bar component
type spectrumActionBar struct {
	app.Compo

	// Properties
	PropEmphasized bool
	PropFlexible   bool
	PropOpen       bool
	PropVariant    ActionBarVariant
	PropContent    string
	PropInnerHTML  string

	// Child elements
	PropButtons []app.UI
	PropChild   app.UI
}

// ActionBar creates a new action bar component
func ActionBar() *spectrumActionBar {
	return &spectrumActionBar{}
}

// Emphasized sets whether the action bar has additional visual emphasis
func (a *spectrumActionBar) Emphasized(emphasized bool) *spectrumActionBar {
	a.PropEmphasized = emphasized
	return a
}

// Flexible sets whether the action bar sizes itself to its content
func (a *spectrumActionBar) Flexible(flexible bool) *spectrumActionBar {
	a.PropFlexible = flexible
	return a
}

// Open sets whether the action bar is visible
func (a *spectrumActionBar) Open(open bool) *spectrumActionBar {
	a.PropOpen = open
	return a
}

// Variant sets the variant of the action bar (fixed or sticky)
func (a *spectrumActionBar) Variant(variant ActionBarVariant) *spectrumActionBar {
	a.PropVariant = variant
	return a
}

// Content sets the text content of the action bar
func (a *spectrumActionBar) Content(content string) *spectrumActionBar {
	a.PropContent = content
	return a
}

// InnerHTML sets the inner HTML of the action bar
func (a *spectrumActionBar) InnerHTML(html string) *spectrumActionBar {
	a.PropInnerHTML = html
	return a
}

// AddButton adds a button to the action bar
func (a *spectrumActionBar) AddButton(button app.UI) *spectrumActionBar {
	a.PropButtons = append(a.PropButtons, button)
	return a
}

// Child sets a UI element as the child of the action bar
func (a *spectrumActionBar) Child(child app.UI) *spectrumActionBar {
	a.PropChild = child
	return a
}

// Render renders the action bar component
func (a *spectrumActionBar) Render() app.UI {
	actionBar := app.Elem("sp-action-bar").
		Attr("emphasized", a.PropEmphasized).
		Attr("flexible", a.PropFlexible).
		Attr("open", a.PropOpen)

	if a.PropVariant != "" {
		actionBar = actionBar.Attr("variant", string(a.PropVariant))
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child first
	if a.PropChild != nil {
		elements = append(elements, a.PropChild)
	} else if a.PropInnerHTML != "" {
		elements = append(elements, app.Raw(a.PropInnerHTML))
	} else if a.PropContent != "" {
		elements = append(elements, app.Text(a.PropContent))
	}

	// Add buttons in the buttons slot
	for _, button := range a.PropButtons {
		buttonElem := button
		if buttonWithSlot, ok := buttonElem.(interface{ Slot(string) app.UI }); ok {
			buttonElem = buttonWithSlot.Slot("buttons")
		} else {
			buttonElem = app.Elem("div").
				Attr("slot", "buttons").
				Body(buttonElem)
		}
		elements = append(elements, buttonElem)
	}

	// Add all elements to the action bar
	if len(elements) > 0 {
		actionBar = actionBar.Body(elements...)
	}

	return actionBar
}
