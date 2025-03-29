package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// CheckboxSize represents the visual size of a checkbox
type CheckboxSize string

// Checkbox sizes
const (
	CheckboxSizeS  CheckboxSize = "s"
	CheckboxSizeM  CheckboxSize = "m"
	CheckboxSizeL  CheckboxSize = "l"
	CheckboxSizeXL CheckboxSize = "xl"
)

// spectrumCheckbox represents an sp-checkbox component
type spectrumCheckbox struct {
	app.Compo

	// Properties
	PropLabel         string
	PropSize          CheckboxSize
	PropChecked       bool
	PropDisabled      bool
	PropEmphasized    bool
	PropIndeterminate bool
	PropInvalid       bool
	PropName          string
	PropReadonly      bool
	PropTabIndex      int

	// Event handlers
	PropOnChange app.EventHandler
}

// Checkbox creates a new checkbox component
func Checkbox() *spectrumCheckbox {
	return &spectrumCheckbox{
		PropSize: CheckboxSizeM, // Default size is medium
	}
}

// Label sets the checkbox label text content
func (c *spectrumCheckbox) Label(label string) *spectrumCheckbox {
	c.PropLabel = label
	return c
}

// Size sets the visual size of the checkbox
func (c *spectrumCheckbox) Size(size CheckboxSize) *spectrumCheckbox {
	c.PropSize = size
	return c
}

// Checked sets the checked state of the checkbox
func (c *spectrumCheckbox) Checked(checked bool) *spectrumCheckbox {
	c.PropChecked = checked
	return c
}

// Disabled sets the disabled state of the checkbox
func (c *spectrumCheckbox) Disabled(disabled bool) *spectrumCheckbox {
	c.PropDisabled = disabled
	return c
}

// Emphasized sets the emphasized style of the checkbox
func (c *spectrumCheckbox) Emphasized(emphasized bool) *spectrumCheckbox {
	c.PropEmphasized = emphasized
	return c
}

// Indeterminate sets the indeterminate state of the checkbox
func (c *spectrumCheckbox) Indeterminate(indeterminate bool) *spectrumCheckbox {
	c.PropIndeterminate = indeterminate
	return c
}

// Invalid sets the invalid state of the checkbox
func (c *spectrumCheckbox) Invalid(invalid bool) *spectrumCheckbox {
	c.PropInvalid = invalid
	return c
}

// Name sets the name attribute of the checkbox
func (c *spectrumCheckbox) Name(name string) *spectrumCheckbox {
	c.PropName = name
	return c
}

// Readonly sets the readonly state of the checkbox
func (c *spectrumCheckbox) Readonly(readonly bool) *spectrumCheckbox {
	c.PropReadonly = readonly
	return c
}

// TabIndex sets the tab index for the checkbox
func (c *spectrumCheckbox) TabIndex(index int) *spectrumCheckbox {
	c.PropTabIndex = index
	return c
}

// OnChange sets the change event handler
func (c *spectrumCheckbox) OnChange(handler app.EventHandler) *spectrumCheckbox {
	c.PropOnChange = handler
	return c
}

// Render renders the checkbox component
func (c *spectrumCheckbox) Render() app.UI {
	checkbox := app.Elem("sp-checkbox").
		Attr("size", string(c.PropSize)).
		Attr("checked", c.PropChecked).
		Attr("disabled", c.PropDisabled).
		Attr("emphasized", c.PropEmphasized).
		Attr("indeterminate", c.PropIndeterminate).
		Attr("invalid", c.PropInvalid).
		Attr("name", c.PropName).
		Attr("readonly", c.PropReadonly).
		Attr("tabindex", c.PropTabIndex)

	// Add event handlers
	if c.PropOnChange != nil {
		checkbox = checkbox.On("change", c.PropOnChange)
	}

	// Add the label as text content
	if c.PropLabel != "" {
		checkbox = checkbox.Text(c.PropLabel)
	}

	return checkbox
}
