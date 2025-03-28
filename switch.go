package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SwitchSize represents the visual size of a switch
type SwitchSize string

// Switch sizes
const (
	SwitchSizeS  SwitchSize = "s"
	SwitchSizeM  SwitchSize = "m"
	SwitchSizeL  SwitchSize = "l"
	SwitchSizeXL SwitchSize = "xl"
)

// SpectrumSwitch represents an sp-switch component
type SpectrumSwitch struct {
	app.Compo

	// Properties
	label      string
	size       SwitchSize
	checked    bool
	disabled   bool
	emphasized bool
	name       string
	readonly   bool
	tabIndex   int

	// Event handlers
	onChange app.EventHandler
}

// Switch creates a new switch component
func Switch() *SpectrumSwitch {
	return &SpectrumSwitch{
		size: SwitchSizeM, // Default size is medium
	}
}

// Label sets the switch label text content
func (s *SpectrumSwitch) Label(label string) *SpectrumSwitch {
	s.label = label
	return s
}

// Size sets the visual size of the switch
func (s *SpectrumSwitch) Size(size SwitchSize) *SpectrumSwitch {
	s.size = size
	return s
}

// Checked sets the checked state of the switch
func (s *SpectrumSwitch) Checked(checked bool) *SpectrumSwitch {
	s.checked = checked
	return s
}

// Disabled sets the disabled state of the switch
func (s *SpectrumSwitch) Disabled(disabled bool) *SpectrumSwitch {
	s.disabled = disabled
	return s
}

// Emphasized sets the emphasized style of the switch
func (s *SpectrumSwitch) Emphasized(emphasized bool) *SpectrumSwitch {
	s.emphasized = emphasized
	return s
}

// Name sets the name attribute of the switch
func (s *SpectrumSwitch) Name(name string) *SpectrumSwitch {
	s.name = name
	return s
}

// Readonly sets the readonly state of the switch
func (s *SpectrumSwitch) Readonly(readonly bool) *SpectrumSwitch {
	s.readonly = readonly
	return s
}

// TabIndex sets the tab index for the switch
func (s *SpectrumSwitch) TabIndex(index int) *SpectrumSwitch {
	s.tabIndex = index
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSwitch) OnChange(handler app.EventHandler) *SpectrumSwitch {
	s.onChange = handler
	return s
}

// Render renders the switch component
func (s *SpectrumSwitch) Render() app.UI {
	switchElem := app.Elem("sp-switch").
		Attr("size", string(s.size)).
		Attr("checked", s.checked).
		Attr("disabled", s.disabled).
		Attr("emphasized", s.emphasized).
		Attr("name", s.name).
		Attr("readonly", s.readonly).
		Attr("tabindex", s.tabIndex)

	// Add event handlers
	if s.onChange != nil {
		switchElem = switchElem.On("change", s.onChange)
	}

	// Add the label as text content
	if s.label != "" {
		switchElem = switchElem.Text(s.label)
	}

	return switchElem
}
