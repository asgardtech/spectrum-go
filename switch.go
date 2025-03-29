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

// spectrumSwitch represents an sp-switch component
type spectrumSwitch struct {
	app.Compo

	// Properties
	PropLabel      string
	PropSize       SwitchSize
	PropChecked    bool
	PropDisabled   bool
	PropEmphasized bool
	PropName       string
	PropReadonly   bool
	PropTabIndex   int

	// Event handlers
	PropOnChange app.EventHandler
}

// Switch creates a new switch component
func Switch() *spectrumSwitch {
	return &spectrumSwitch{
		PropSize: SwitchSizeM, // Default size is medium
	}
}

// Label sets the switch label text content
func (s *spectrumSwitch) Label(label string) *spectrumSwitch {
	s.PropLabel = label
	return s
}

// Size sets the visual size of the switch
func (s *spectrumSwitch) Size(size SwitchSize) *spectrumSwitch {
	s.PropSize = size
	return s
}

// Checked sets the checked state of the switch
func (s *spectrumSwitch) Checked(checked bool) *spectrumSwitch {
	s.PropChecked = checked
	return s
}

// Disabled sets the disabled state of the switch
func (s *spectrumSwitch) Disabled(disabled bool) *spectrumSwitch {
	s.PropDisabled = disabled
	return s
}

// Emphasized sets the emphasized style of the switch
func (s *spectrumSwitch) Emphasized(emphasized bool) *spectrumSwitch {
	s.PropEmphasized = emphasized
	return s
}

// Name sets the name attribute of the switch
func (s *spectrumSwitch) Name(name string) *spectrumSwitch {
	s.PropName = name
	return s
}

// Readonly sets the readonly state of the switch
func (s *spectrumSwitch) Readonly(readonly bool) *spectrumSwitch {
	s.PropReadonly = readonly
	return s
}

// TabIndex sets the tab index for the switch
func (s *spectrumSwitch) TabIndex(index int) *spectrumSwitch {
	s.PropTabIndex = index
	return s
}

// OnChange sets the change event handler
func (s *spectrumSwitch) OnChange(handler app.EventHandler) *spectrumSwitch {
	s.PropOnChange = handler
	return s
}

// Render renders the switch component
func (s *spectrumSwitch) Render() app.UI {
	switchElem := app.Elem("sp-switch").
		Attr("size", string(s.PropSize)).
		Attr("checked", s.PropChecked).
		Attr("disabled", s.PropDisabled).
		Attr("emphasized", s.PropEmphasized).
		Attr("name", s.PropName).
		Attr("readonly", s.PropReadonly).
		Attr("tabindex", s.PropTabIndex)

	// Add event handlers
	if s.PropOnChange != nil {
		switchElem = switchElem.On("change", s.PropOnChange)
	}

	// Add the label as text content
	if s.PropLabel != "" {
		switchElem = switchElem.Text(s.PropLabel)
	}

	return switchElem
}
