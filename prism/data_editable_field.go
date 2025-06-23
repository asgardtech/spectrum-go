package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// EditableField represents a single form field that can be edited
type EditableField struct {
	app.Compo

	Name        string
	Label       string
	Value       string
	Type        string
	Placeholder string
	Required    bool
	Disabled    bool
	ReadOnly    bool
	HelpText    string
	Errors      []string
	Options     []FieldOption // For select, radio, checkbox
	OnChange    func(app.Context, string)
	OnBlur      func(app.Context, string)
	OnFocus     func(app.Context)
}

// FieldOption represents an option for select, radio, or checkbox fields
type FieldOption struct {
	Value    string
	Label    string
	Disabled bool
	Selected bool
}

// NewEditableField creates a new editable field
func NewEditableField(name, label string) *EditableField {
	return &EditableField{
		Name:     name,
		Label:    label,
		Type:     "text",
		Required: false,
		Disabled: false,
		ReadOnly: false,
	}
}

func (ef *EditableField) WithValue(value string) *EditableField {
	ef.Value = value
	return ef
}

func (ef *EditableField) WithType(fieldType string) *EditableField {
	ef.Type = fieldType
	return ef
}

func (ef *EditableField) WithPlaceholder(placeholder string) *EditableField {
	ef.Placeholder = placeholder
	return ef
}

func (ef *EditableField) WithRequired(required bool) *EditableField {
	ef.Required = required
	return ef
}

func (ef *EditableField) WithDisabled(disabled bool) *EditableField {
	ef.Disabled = disabled
	return ef
}

func (ef *EditableField) WithReadOnly(readOnly bool) *EditableField {
	ef.ReadOnly = readOnly
	return ef
}

func (ef *EditableField) WithHelpText(helpText string) *EditableField {
	ef.HelpText = helpText
	return ef
}

func (ef *EditableField) WithErrors(errors []string) *EditableField {
	ef.Errors = errors
	return ef
}

func (ef *EditableField) WithOptions(options []FieldOption) *EditableField {
	ef.Options = options
	return ef
}

func (ef *EditableField) WithOnChange(fn func(app.Context, string)) *EditableField {
	ef.OnChange = fn
	return ef
}

func (ef *EditableField) WithOnBlur(fn func(app.Context, string)) *EditableField {
	ef.OnBlur = fn
	return ef
}

func (ef *EditableField) WithOnFocus(fn func(app.Context)) *EditableField {
	ef.OnFocus = fn
	return ef
}

func (ef *EditableField) Render() app.UI {
	fieldId := "field-" + ef.Name
	hasError := len(ef.Errors) > 0

	return app.Div().
		Class("prism-editable-field").
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("gap", "4px").
		Body(
			ef.renderLabel(fieldId),
			ef.renderInput(fieldId, hasError),
			ef.renderHelpText(),
			ef.renderErrors(),
		)
}

func (ef *EditableField) renderLabel(fieldId string) app.UI {
	if ef.Label == "" {
		return app.Div()
	}

	return sp.FieldLabel().
		For(fieldId).
		Required(ef.Required).
		Text(ef.Label)
}

func (ef *EditableField) renderInput(fieldId string, hasError bool) app.UI {
	switch ef.Type {
	case "textarea":
		return ef.renderTextarea(fieldId, hasError)
	case "select":
		return ef.renderSelect(fieldId, hasError)
	case "radio":
		return ef.renderRadioGroup(fieldId)
	case "checkbox":
		return ef.renderCheckbox(fieldId)
	case "checkboxgroup":
		return ef.renderCheckboxGroup(fieldId)
	default:
		return ef.renderTextInput(fieldId, hasError)
	}
}

func (ef *EditableField) renderTextInput(fieldId string, hasError bool) app.UI {
	return sp.Textfield().
		Id(fieldId).
		Name(ef.Name).
		Value(ef.Value).
		Placeholder(ef.Placeholder).
		Required(ef.Required).
		Disabled(ef.Disabled).
		Readonly(ef.ReadOnly).
		Invalid(hasError).
		OnInput(func(ctx app.Context, e app.Event) {
			value := ctx.JSSrc().Get("value").String()
			if ef.OnChange != nil {
				ef.OnChange(ctx, value)
			}
		})
}

func (ef *EditableField) renderTextarea(fieldId string, hasError bool) app.UI {
	return sp.Textarea().
		Id(fieldId).
		Name(ef.Name).
		Value(ef.Value).
		Placeholder(ef.Placeholder).
		Required(ef.Required).
		Disabled(ef.Disabled).
		Readonly(ef.ReadOnly).
		Invalid(hasError).
		OnInput(func(ctx app.Context, e app.Event) {
			value := ctx.JSSrc().Get("value").String()
			if ef.OnChange != nil {
				ef.OnChange(ctx, value)
			}
		})
}

func (ef *EditableField) renderSelect(fieldId string, hasError bool) app.UI {
	return sp.Picker().
		Id(fieldId).
		Disabled(ef.Disabled).
		Invalid(hasError).
		Body(
			app.Range(ef.Options).Slice(func(i int) app.UI {
				option := ef.Options[i]
				return sp.MenuItem().
					Value(option.Value).
					Selected(option.Value == ef.Value).
					Disabled(option.Disabled).
					Text(option.Label).
					OnClick(func(ctx app.Context, e app.Event) {
						if ef.OnChange != nil {
							ef.OnChange(ctx, option.Value)
						}
					})
			}),
		)
}

func (ef *EditableField) renderRadioGroup(fieldId string) app.UI {
	return sp.RadioGroup().
		Body(
			app.Range(ef.Options).Slice(func(i int) app.UI {
				option := ef.Options[i]
				return sp.Radio().
					Value(option.Value).
					Checked(option.Value == ef.Value).
					Disabled(option.Disabled).
					Text(option.Label).
					OnChange(func(ctx app.Context, e app.Event) {
						if ef.OnChange != nil {
							ef.OnChange(ctx, option.Value)
						}
					})
			}),
		)
}

func (ef *EditableField) renderCheckbox(fieldId string) app.UI {
	return sp.Checkbox().
		Id(fieldId).
		Name(ef.Name).
		Checked(ef.Value == "true").
		Disabled(ef.Disabled).
		Text(ef.Label).
		OnChange(func(ctx app.Context, e app.Event) {
			checked := ctx.JSSrc().Get("checked").Bool()
			value := "false"
			if checked {
				value = "true"
			}
			if ef.OnChange != nil {
				ef.OnChange(ctx, value)
			}
		})
}

func (ef *EditableField) renderCheckboxGroup(fieldId string) app.UI {
	// Use a regular div container since CheckboxGroup doesn't exist in Spectrum
	return app.Div().
		Class("spectrum-checkboxgroup").
		Body(
			app.Range(ef.Options).Slice(func(i int) app.UI {
				option := ef.Options[i]
				return sp.Checkbox().
					Checked(option.Selected).
					Disabled(option.Disabled).
					Text(option.Label).
					OnChange(func(ctx app.Context, e app.Event) {
						// For checkbox groups, we'd need to handle multiple values
						if ef.OnChange != nil {
							ef.OnChange(ctx, option.Value)
						}
					})
			}),
		)
}

func (ef *EditableField) renderHelpText() app.UI {
	if ef.HelpText == "" {
		return app.Div()
	}

	return sp.HelpText().
		Text(ef.HelpText)
}

func (ef *EditableField) renderErrors() app.UI {
	if len(ef.Errors) == 0 {
		return app.Div()
	}

	var errorElements []app.UI
	for _, errorMsg := range ef.Errors {
		errorElements = append(errorElements,
			app.Div().
				Style("color", "#d73a49").
				Style("font-size", "14px").
				Style("margin-top", "2px").
				Text(errorMsg),
		)
	}

	return app.Div().
		Class("prism-field-errors").
		Body(errorElements...)
}

