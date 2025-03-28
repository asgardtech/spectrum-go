package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionBarVariant represents the variant of an action bar
type ActionBarVariant string

// Action bar variants
const (
	ActionBarVariantFixed  ActionBarVariant = "fixed"
	ActionBarVariantSticky ActionBarVariant = "sticky"
)

// SpectrumActionBar represents an sp-action-bar component
type SpectrumActionBar struct {
	app.Compo

	// Properties
	emphasized bool
	flexible   bool
	open       bool
	variant    ActionBarVariant
	content    string
	innerHTML  string

	// Child elements
	buttons []app.UI
	child   app.UI
}

// ActionBar creates a new action bar component
func ActionBar() *SpectrumActionBar {
	return &SpectrumActionBar{}
}

// Emphasized sets whether the action bar has additional visual emphasis
func (a *SpectrumActionBar) Emphasized(emphasized bool) *SpectrumActionBar {
	a.emphasized = emphasized
	return a
}

// Flexible sets whether the action bar sizes itself to its content
func (a *SpectrumActionBar) Flexible(flexible bool) *SpectrumActionBar {
	a.flexible = flexible
	return a
}

// Open sets whether the action bar is visible
func (a *SpectrumActionBar) Open(open bool) *SpectrumActionBar {
	a.open = open
	return a
}

// Variant sets the variant of the action bar (fixed or sticky)
func (a *SpectrumActionBar) Variant(variant ActionBarVariant) *SpectrumActionBar {
	a.variant = variant
	return a
}

// Content sets the text content of the action bar
func (a *SpectrumActionBar) Content(content string) *SpectrumActionBar {
	a.content = content
	return a
}

// InnerHTML sets the inner HTML of the action bar
func (a *SpectrumActionBar) InnerHTML(html string) *SpectrumActionBar {
	a.innerHTML = html
	return a
}

// AddButton adds a button to the action bar
func (a *SpectrumActionBar) AddButton(button app.UI) *SpectrumActionBar {
	a.buttons = append(a.buttons, button)
	return a
}

// Child sets a UI element as the child of the action bar
func (a *SpectrumActionBar) Child(child app.UI) *SpectrumActionBar {
	a.child = child
	return a
}

// Render renders the action bar component
func (a *SpectrumActionBar) Render() app.UI {
	actionBar := app.Elem("sp-action-bar").
		Attr("emphasized", a.emphasized).
		Attr("flexible", a.flexible).
		Attr("open", a.open)

	if a.variant != "" {
		actionBar = actionBar.Attr("variant", string(a.variant))
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add content or child first
	if a.child != nil {
		elements = append(elements, a.child)
	} else if a.innerHTML != "" {
		elements = append(elements, app.Raw(a.innerHTML))
	} else if a.content != "" {
		elements = append(elements, app.Text(a.content))
	}

	// Add buttons in the buttons slot
	for _, button := range a.buttons {
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
