package prism

import (
	"fmt"
	"strings"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ValidationResult holds validation state for a field
type ValidationResult struct {
	Valid   bool
	Message string
}

// Validator is a function that validates a field value
type Validator func(value string) ValidationResult

// FormField represents a single form field with validation
type FormField struct {
	Name         string
	Label        string
	Type         string
	Value        string
	Required     bool
	Placeholder  string
	Validators   []Validator
	Errors       []string
	HelpText     string
	Disabled     bool
}

// FormData holds all form fields and validation state
type FormData struct {
	Fields       map[string]*FormField
	IsSubmitting bool
	HasErrors    bool
	SubmitError  string
	OnSubmit     func(data map[string]string) error
	OnValidate   func(fieldName, value string) ValidationResult
}

// NewFormData creates a new form data structure
func NewFormData() *FormData {
	return &FormData{
		Fields:       make(map[string]*FormField),
		IsSubmitting: false,
		HasErrors:    false,
	}
}

// AddField adds a field to the form
func (fd *FormData) AddField(field *FormField) *FormData {
	fd.Fields[field.Name] = field
	return fd
}

// SetFieldValue sets a field value and validates it
func (fd *FormData) SetFieldValue(name, value string) *FormData {
	if field, exists := fd.Fields[name]; exists {
		field.Value = value
		field.Errors = fd.validateField(field)
	}
	fd.updateHasErrors()
	return fd
}

// GetFieldValue gets a field value
func (fd *FormData) GetFieldValue(name string) string {
	if field, exists := fd.Fields[name]; exists {
		return field.Value
	}
	return ""
}

// ValidateAll validates all fields in the form
func (fd *FormData) ValidateAll() bool {
	allValid := true
	for _, field := range fd.Fields {
		field.Errors = fd.validateField(field)
		if len(field.Errors) > 0 {
			allValid = false
		}
	}
	fd.HasErrors = !allValid
	return allValid
}

// validateField validates a single field
func (fd *FormData) validateField(field *FormField) []string {
	var errors []string

	// Required validation
	if field.Required && strings.TrimSpace(field.Value) == "" {
		errors = append(errors, fmt.Sprintf("%s is required", field.Label))
	}

	// Custom validators
	for _, validator := range field.Validators {
		if result := validator(field.Value); !result.Valid {
			errors = append(errors, result.Message)
		}
	}

	// Custom validation callback
	if fd.OnValidate != nil {
		if result := fd.OnValidate(field.Name, field.Value); !result.Valid {
			errors = append(errors, result.Message)
		}
	}

	return errors
}

// updateHasErrors updates the overall form error state
func (fd *FormData) updateHasErrors() {
	hasErrors := false
	for _, field := range fd.Fields {
		if len(field.Errors) > 0 {
			hasErrors = true
			break
		}
	}
	fd.HasErrors = hasErrors
}

// GetFormValues returns all field values as a map
func (fd *FormData) GetFormValues() map[string]string {
	values := make(map[string]string)
	for name, field := range fd.Fields {
		values[name] = field.Value
	}
	return values
}

// Form component that renders a complete form with validation
type FormComponent struct {
	app.Compo
	
	FormData     *FormData
	Title        string
	SubmitText   string
	CancelText   string
	OnSubmit     func(app.Context, map[string]string) error
	OnCancel     func(app.Context)
	ShowCancel   bool
	Layout       string // "vertical" or "horizontal"
}

// NewForm creates a new form component
func NewForm() *FormComponent {
	return &FormComponent{
		FormData:   NewFormData(),
		SubmitText: "Submit",
		CancelText: "Cancel",
		ShowCancel: false,
		Layout:     "vertical",
	}
}

func (f *FormComponent) WithTitle(title string) *FormComponent {
	f.Title = title
	return f
}

func (f *FormComponent) WithFormData(data *FormData) *FormComponent {
	f.FormData = data
	return f
}

func (f *FormComponent) WithSubmitText(text string) *FormComponent {
	f.SubmitText = text
	return f
}

func (f *FormComponent) WithCancelText(text string) *FormComponent {
	f.CancelText = text
	return f
}

func (f *FormComponent) WithLayout(layout string) *FormComponent {
	f.Layout = layout
	return f
}

func (f *FormComponent) WithShowCancel(show bool) *FormComponent {
	f.ShowCancel = show
	return f
}

func (f *FormComponent) WithOnSubmit(fn func(app.Context, map[string]string) error) *FormComponent {
	f.OnSubmit = fn
	return f
}

func (f *FormComponent) WithOnCancel(fn func(app.Context)) *FormComponent {
	f.OnCancel = fn
	return f
}

func (f *FormComponent) Render() app.UI {
	return app.Form().
		Class("prism-form").
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("gap", "16px").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()
			f.handleSubmit(ctx)
		}).
		Body(
			app.If(f.Title != "", func() app.UI {
				return app.H2().
					Style("margin", "0 0 16px 0").
					Text(f.Title)
			}),
			app.If(f.FormData.SubmitError != "", func() app.UI {
				return f.renderError(f.FormData.SubmitError)
			}),
			f.renderFields(),
			f.renderActions(),
		)
}

func (f *FormComponent) renderFields() app.UI {
	var fields []app.UI
	
	for _, field := range f.FormData.Fields {
		fields = append(fields, f.renderField(field))
	}
	
	return app.Div().
		Class("prism-form-fields").
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("gap", "16px").
		Body(fields...)
}

func (f *FormComponent) renderField(field *FormField) app.UI {
	fieldId := "field-" + field.Name
	hasError := len(field.Errors) > 0
	
	return app.Div().
		Class("prism-form-field").
		Body(
			sp.FieldLabel().
				For(fieldId).
				Required(field.Required).
				Text(field.Label),
			f.renderFieldInput(field, fieldId, hasError),
			app.If(field.HelpText != "", func() app.UI {
				return sp.HelpText().
					Text(field.HelpText)
			}),
			app.If(hasError, func() app.UI {
				return f.renderFieldErrors(field.Errors)
			}),
		)
}

func (f *FormComponent) renderFieldInput(field *FormField, fieldId string, hasError bool) app.UI {
	baseInput := sp.Textfield().
		Id(fieldId).
		Name(field.Name).
		Placeholder(field.Placeholder).
		Value(field.Value).
		Invalid(hasError).
		Disabled(field.Disabled).
		OnInput(func(ctx app.Context, e app.Event) {
			value := ctx.JSSrc().Get("value").String()
			f.FormData.SetFieldValue(field.Name, value)
			ctx.Reload()
		})

	switch field.Type {
	case "email":
		return baseInput.Type("email")
	case "password":
		return baseInput.Type("password")
	case "number":
		return baseInput.Type("number")
	case "tel":
		return baseInput.Type("tel")
	case "url":
		return baseInput.Type("url")
	case "textarea":
		return sp.Textarea().
			Id(fieldId).
			Name(field.Name).
			Placeholder(field.Placeholder).
			Value(field.Value).
			Invalid(hasError).
			Disabled(field.Disabled).
			OnInput(func(ctx app.Context, e app.Event) {
				value := ctx.JSSrc().Get("value").String()
				f.FormData.SetFieldValue(field.Name, value)
				ctx.Reload()
			})
	default:
		return baseInput.Type("text")
	}
}

func (f *FormComponent) renderFieldErrors(errors []string) app.UI {
	var errorElements []app.UI
	for _, error := range errors {
		errorElements = append(errorElements, 
			app.Div().
				Style("color", "#d73a49").
				Style("font-size", "14px").
				Style("margin-top", "4px").
				Text(error),
		)
	}
	
	return app.Div().
		Class("prism-form-field-errors").
		Body(errorElements...)
}

func (f *FormComponent) renderError(message string) app.UI {
	return sp.AlertBanner().
		Variant("negative").
		Body(
			app.Div().Text(message),
		)
}

func (f *FormComponent) renderActions() app.UI {
	return app.Div().
		Class("prism-form-actions").
		Style("display", "flex").
		Style("gap", "12px").
		Style("justify-content", "flex-end").
		Style("margin-top", "24px").
		Body(
			app.If(f.ShowCancel, func() app.UI {
				return sp.Button().
					Text(f.CancelText).
					Variant("secondary").
					OnClick(func(ctx app.Context, e app.Event) {
						if f.OnCancel != nil {
							f.OnCancel(ctx)
						}
					})
			}),
			sp.Button().
				Text(f.SubmitText).
				Variant("primary").
				Type("submit").
				Disabled(f.FormData.IsSubmitting || f.FormData.HasErrors),
		)
}

func (f *FormComponent) handleSubmit(ctx app.Context) {
	if !f.FormData.ValidateAll() {
		ctx.Reload()
		return
	}

	if f.OnSubmit == nil {
		return
	}

	f.FormData.IsSubmitting = true
	f.FormData.SubmitError = ""
	ctx.Reload()

	values := f.FormData.GetFormValues()
	if err := f.OnSubmit(ctx, values); err != nil {
		f.FormData.SubmitError = err.Error()
	}

	f.FormData.IsSubmitting = false
	ctx.Reload()
}

// Common validators
func EmailValidator() Validator {
	return func(value string) ValidationResult {
		if value == "" {
			return ValidationResult{Valid: true}
		}
		if !strings.Contains(value, "@") || !strings.Contains(value, ".") {
			return ValidationResult{Valid: false, Message: "Please enter a valid email address"}
		}
		return ValidationResult{Valid: true}
	}
}

func MinLengthValidator(minLength int) Validator {
	return func(value string) ValidationResult {
		if len(value) < minLength {
			return ValidationResult{
				Valid:   false,
				Message: fmt.Sprintf("Must be at least %d characters long", minLength),
			}
		}
		return ValidationResult{Valid: true}
	}
}

func MaxLengthValidator(maxLength int) Validator {
	return func(value string) ValidationResult {
		if len(value) > maxLength {
			return ValidationResult{
				Valid:   false,
				Message: fmt.Sprintf("Must be no more than %d characters long", maxLength),
			}
		}
		return ValidationResult{Valid: true}
	}
}

func PatternValidator(pattern string, message string) Validator {
	return func(value string) ValidationResult {
		// Simple pattern matching - in real implementation would use regexp
		if value != "" && !strings.Contains(value, pattern) {
			return ValidationResult{Valid: false, Message: message}
		}
		return ValidationResult{Valid: true}
	}
}