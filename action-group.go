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

// SpectrumActionGroup represents an sp-action-group component
type SpectrumActionGroup struct {
	app.Compo

	// Properties
	compact     bool
	emphasized  bool
	justified   bool
	label       string
	quiet       bool
	selects     ActionGroupSelects
	staticColor ActionGroupStaticColor
	vertical    bool
	size        ActionGroupSize
	selected    []string

	// Event handlers
	onChange app.EventHandler

	// Child elements
	actions []app.UI
}

// ActionGroup creates a new action group component
func ActionGroup() *SpectrumActionGroup {
	return &SpectrumActionGroup{
		size:     ActionGroupSizeM, // Default size is medium
		selected: []string{},
	}
}

// Compact sets whether the buttons are displayed more compactly
func (a *SpectrumActionGroup) Compact(compact bool) *SpectrumActionGroup {
	a.compact = compact
	return a
}

// Emphasized sets whether the buttons have an emphasized appearance
func (a *SpectrumActionGroup) Emphasized(emphasized bool) *SpectrumActionGroup {
	a.emphasized = emphasized
	return a
}

// Justified sets whether the buttons fill the available horizontal space
func (a *SpectrumActionGroup) Justified(justified bool) *SpectrumActionGroup {
	a.justified = justified
	return a
}

// Label sets the accessible label for the action group
func (a *SpectrumActionGroup) Label(label string) *SpectrumActionGroup {
	a.label = label
	return a
}

// Quiet sets whether the buttons have a quiet appearance
func (a *SpectrumActionGroup) Quiet(quiet bool) *SpectrumActionGroup {
	a.quiet = quiet
	return a
}

// Selects sets the selection mode (single or multiple)
func (a *SpectrumActionGroup) Selects(selects ActionGroupSelects) *SpectrumActionGroup {
	a.selects = selects
	return a
}

// StaticColor sets the static color variant
func (a *SpectrumActionGroup) StaticColor(staticColor ActionGroupStaticColor) *SpectrumActionGroup {
	a.staticColor = staticColor
	return a
}

// Vertical sets whether the buttons are arranged vertically
func (a *SpectrumActionGroup) Vertical(vertical bool) *SpectrumActionGroup {
	a.vertical = vertical
	return a
}

// Size sets the visual size of the action group
func (a *SpectrumActionGroup) Size(size ActionGroupSize) *SpectrumActionGroup {
	a.size = size
	return a
}

// Selected sets the selected values
func (a *SpectrumActionGroup) Selected(selected []string) *SpectrumActionGroup {
	a.selected = selected
	return a
}

// OnChange sets the change event handler
func (a *SpectrumActionGroup) OnChange(handler app.EventHandler) *SpectrumActionGroup {
	a.onChange = handler
	return a
}

// AddAction adds an action button to the group
func (a *SpectrumActionGroup) AddAction(action app.UI) *SpectrumActionGroup {
	a.actions = append(a.actions, action)
	return a
}

// Actions sets multiple action buttons at once
func (a *SpectrumActionGroup) Actions(actions ...app.UI) *SpectrumActionGroup {
	a.actions = actions
	return a
}

// Render renders the action group component
func (a *SpectrumActionGroup) Render() app.UI {
	// Convert the selected array to JSON
	selectedJSON, _ := json.Marshal(a.selected)

	actionGroup := app.Elem("sp-action-group").
		Attr("compact", a.compact).
		Attr("emphasized", a.emphasized).
		Attr("justified", a.justified).
		Attr("quiet", a.quiet).
		Attr("vertical", a.vertical)

	if a.label != "" {
		actionGroup = actionGroup.Attr("label", a.label)
	}
	if a.selects != "" {
		actionGroup = actionGroup.Attr("selects", string(a.selects))
	}
	if a.staticColor != "" {
		actionGroup = actionGroup.Attr("static-color", string(a.staticColor))
	}
	if a.size != "" {
		actionGroup = actionGroup.Attr("size", string(a.size))
	}
	if len(a.selected) > 0 {
		actionGroup = actionGroup.Attr("selected", string(selectedJSON))
	}

	// Add event handlers
	if a.onChange != nil {
		actionGroup = actionGroup.On("change", a.onChange)
	}

	// Add action buttons
	if len(a.actions) > 0 {
		actionGroup = actionGroup.Body(a.actions...)
	}

	return actionGroup
}
