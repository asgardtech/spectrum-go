package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ActionGroupSize represents the visual size of an action group
type ActionGroupSize string

// Action group sizes
const (
	ActionGroupSizeXS ActionGroupSize = "xs" // Extra Small
	ActionGroupSizeS  ActionGroupSize = "s"  // Small
	ActionGroupSizeM  ActionGroupSize = "m"  // Medium
	ActionGroupSizeL  ActionGroupSize = "l"  // Large
	ActionGroupSizeXL ActionGroupSize = "xl" // Extra Large
)

// ActionGroupStaticColor represents the static color variant for an action group
type ActionGroupStaticColor string

// Action group static colors
const (
	ActionGroupStaticColorWhite ActionGroupStaticColor = "white"
	ActionGroupStaticColorBlack ActionGroupStaticColor = "black"
)

// ActionGroupSelects represents the selection mode of an action group
type ActionGroupSelects string

// Action group selection modes
const (
	ActionGroupSelectsSingle   ActionGroupSelects = "single"
	ActionGroupSelectsMultiple ActionGroupSelects = "multiple"
)

// spectrumActionGroup represents an sp-action-group component
type spectrumActionGroup struct {
	app.Compo

	// Properties
	PropCompact     bool
	PropEmphasized  bool
	PropJustified   bool
	PropLabel       string
	PropQuiet       bool
	PropSelects     ActionGroupSelects
	PropStaticColor ActionGroupStaticColor
	PropVertical    bool
	PropSize        ActionGroupSize
	PropSelected    []string

	// Event handlers
	PropOnChange app.EventHandler

	// Child elements
	PropActions []app.UI
}

// ActionGroup creates a new action group component
func ActionGroup() *spectrumActionGroup {
	return &spectrumActionGroup{
		PropSize:     ActionGroupSizeM, // Default size is medium
		PropSelected: []string{},
	}
}

// Compact sets whether the buttons are displayed more compactly
func (a *spectrumActionGroup) Compact(compact bool) *spectrumActionGroup {
	a.PropCompact = compact
	return a
}

// Emphasized sets whether the buttons have an emphasized appearance
func (a *spectrumActionGroup) Emphasized(emphasized bool) *spectrumActionGroup {
	a.PropEmphasized = emphasized
	return a
}

// Justified sets whether the buttons fill the available horizontal space
func (a *spectrumActionGroup) Justified(justified bool) *spectrumActionGroup {
	a.PropJustified = justified
	return a
}

// Label sets the accessible label for the action group
func (a *spectrumActionGroup) Label(label string) *spectrumActionGroup {
	a.PropLabel = label
	return a
}

// Quiet sets whether the buttons have a quiet appearance
func (a *spectrumActionGroup) Quiet(quiet bool) *spectrumActionGroup {
	a.PropQuiet = quiet
	return a
}

// Selects sets the selection mode (single or multiple)
func (a *spectrumActionGroup) Selects(selects ActionGroupSelects) *spectrumActionGroup {
	a.PropSelects = selects
	return a
}

// StaticColor sets the static color variant
func (a *spectrumActionGroup) StaticColor(staticColor ActionGroupStaticColor) *spectrumActionGroup {
	a.PropStaticColor = staticColor
	return a
}

// Vertical sets whether the buttons are arranged vertically
func (a *spectrumActionGroup) Vertical(vertical bool) *spectrumActionGroup {
	a.PropVertical = vertical
	return a
}

// Size sets the visual size of the action group
func (a *spectrumActionGroup) Size(size ActionGroupSize) *spectrumActionGroup {
	a.PropSize = size
	return a
}

// Selected sets the selected values
func (a *spectrumActionGroup) Selected(selected []string) *spectrumActionGroup {
	a.PropSelected = selected
	return a
}

// OnChange sets the change event handler
func (a *spectrumActionGroup) OnChange(handler app.EventHandler) *spectrumActionGroup {
	a.PropOnChange = handler
	return a
}

// AddAction adds an action button to the group
func (a *spectrumActionGroup) AddAction(action app.UI) *spectrumActionGroup {
	a.PropActions = append(a.PropActions, action)
	return a
}

// Actions sets multiple action buttons at once
func (a *spectrumActionGroup) Actions(actions ...app.UI) *spectrumActionGroup {
	a.PropActions = actions
	return a
}

// Render renders the action group component
func (a *spectrumActionGroup) Render() app.UI {
	// Convert the selected array to JSON
	selectedJSON, _ := json.Marshal(a.PropSelected)

	actionGroup := app.Elem("sp-action-group").
		Attr("compact", a.PropCompact).
		Attr("emphasized", a.PropEmphasized).
		Attr("justified", a.PropJustified).
		Attr("quiet", a.PropQuiet).
		Attr("vertical", a.PropVertical)

	if a.PropLabel != "" {
		actionGroup = actionGroup.Attr("label", a.PropLabel)
	}
	if a.PropSelects != "" {
		actionGroup = actionGroup.Attr("selects", string(a.PropSelects))
	}
	if a.PropStaticColor != "" {
		actionGroup = actionGroup.Attr("static-color", string(a.PropStaticColor))
	}
	if a.PropSize != "" {
		actionGroup = actionGroup.Attr("size", string(a.PropSize))
	}
	if len(a.PropSelected) > 0 {
		actionGroup = actionGroup.Attr("selected", string(selectedJSON))
	}

	// Add event handlers
	if a.PropOnChange != nil {
		actionGroup = actionGroup.On("change", a.PropOnChange)
	}

	// Add action buttons
	if len(a.PropActions) > 0 {
		actionGroup = actionGroup.Body(a.PropActions...)
	}

	return actionGroup
}
