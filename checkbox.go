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

// SpectrumCheckbox represents an sp-checkbox component
type SpectrumCheckbox struct {
	app.Compo

	// Properties
	label         string
	size          CheckboxSize
	checked       bool
	disabled      bool
	emphasized    bool
	indeterminate bool
	invalid       bool
	name          string
	readonly      bool
	tabIndex      int

	// Event handlers
	onChange app.EventHandler
}

// Checkbox creates a new checkbox component
func Checkbox() *SpectrumCheckbox {
	return &SpectrumCheckbox{
		size: CheckboxSizeM, // Default size is medium
	}
}

// Label sets the checkbox label text content
func (c *SpectrumCheckbox) Label(label string) *SpectrumCheckbox {
	c.label = label
	return c
}

// Size sets the visual size of the checkbox
func (c *SpectrumCheckbox) Size(size CheckboxSize) *SpectrumCheckbox {
	c.size = size
	return c
}

// Checked sets the checked state of the checkbox
func (c *SpectrumCheckbox) Checked(checked bool) *SpectrumCheckbox {
	c.checked = checked
	return c
}

// Disabled sets the disabled state of the checkbox
func (c *SpectrumCheckbox) Disabled(disabled bool) *SpectrumCheckbox {
	c.disabled = disabled
	return c
}

// Emphasized sets the emphasized style of the checkbox
func (c *SpectrumCheckbox) Emphasized(emphasized bool) *SpectrumCheckbox {
	c.emphasized = emphasized
	return c
}

// Indeterminate sets the indeterminate state of the checkbox
func (c *SpectrumCheckbox) Indeterminate(indeterminate bool) *SpectrumCheckbox {
	c.indeterminate = indeterminate
	return c
}

// Invalid sets the invalid state of the checkbox
func (c *SpectrumCheckbox) Invalid(invalid bool) *SpectrumCheckbox {
	c.invalid = invalid
	return c
}

// Name sets the name attribute of the checkbox
func (c *SpectrumCheckbox) Name(name string) *SpectrumCheckbox {
	c.name = name
	return c
}

// Readonly sets the readonly state of the checkbox
func (c *SpectrumCheckbox) Readonly(readonly bool) *SpectrumCheckbox {
	c.readonly = readonly
	return c
}

// TabIndex sets the tab index for the checkbox
func (c *SpectrumCheckbox) TabIndex(index int) *SpectrumCheckbox {
	c.tabIndex = index
	return c
}

// OnChange sets the change event handler
func (c *SpectrumCheckbox) OnChange(handler app.EventHandler) *SpectrumCheckbox {
	c.onChange = handler
	return c
}

// Render renders the checkbox component
func (c *SpectrumCheckbox) Render() app.UI {
	checkbox := app.Elem("sp-checkbox").
		Attr("size", string(c.size)).
		Attr("checked", c.checked).
		Attr("disabled", c.disabled).
		Attr("emphasized", c.emphasized).
		Attr("indeterminate", c.indeterminate).
		Attr("invalid", c.invalid).
		Attr("name", c.name).
		Attr("readonly", c.readonly).
		Attr("tabindex", c.tabIndex)

	// Add event handlers
	if c.onChange != nil {
		checkbox = checkbox.On("change", c.onChange)
	}

	// Add the label as text content
	if c.label != "" {
		checkbox = checkbox.Text(c.label)
	}

	return checkbox
}
