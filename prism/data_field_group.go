package prism

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// FieldGroup represents a group of related form fields
type FieldGroup struct {
	app.Compo

	Title       string
	Description string
	Fields      []*EditableField
	Layout      string // "vertical", "horizontal", "grid"
	Columns     int    // For grid layout
	Required    bool
	Collapsible bool
	Collapsed   bool
}

// NewFieldGroup creates a new field group
func NewFieldGroup() *FieldGroup {
	return &FieldGroup{
		Layout:      "vertical",
		Columns:     2,
		Required:    false,
		Collapsible: false,
		Collapsed:   false,
	}
}

func (fg *FieldGroup) WithTitle(title string) *FieldGroup {
	fg.Title = title
	return fg
}

func (fg *FieldGroup) WithDescription(description string) *FieldGroup {
	fg.Description = description
	return fg
}

func (fg *FieldGroup) WithFields(fields ...*EditableField) *FieldGroup {
	fg.Fields = fields
	return fg
}

func (fg *FieldGroup) AddField(field *EditableField) *FieldGroup {
	fg.Fields = append(fg.Fields, field)
	return fg
}

func (fg *FieldGroup) WithLayout(layout string) *FieldGroup {
	fg.Layout = layout
	return fg
}

func (fg *FieldGroup) WithColumns(columns int) *FieldGroup {
	fg.Columns = columns
	return fg
}

func (fg *FieldGroup) WithRequired(required bool) *FieldGroup {
	fg.Required = required
	return fg
}

func (fg *FieldGroup) WithCollapsible(collapsible bool) *FieldGroup {
	fg.Collapsible = collapsible
	return fg
}

func (fg *FieldGroup) WithCollapsed(collapsed bool) *FieldGroup {
	fg.Collapsed = collapsed
	return fg
}

func (fg *FieldGroup) Render() app.UI {
	return app.Div().
		Class("prism-field-group").
		Style("margin-bottom", "24px").
		Body(
			fg.renderHeader(),
			fg.renderContent(),
		)
}

func (fg *FieldGroup) renderHeader() app.UI {
	if fg.Title == "" && fg.Description == "" {
		return app.Div()
	}

	headerContent := []app.UI{}

	if fg.Title != "" {
		var titleElement app.UI
		
		titleText := fg.Title
		if fg.Required {
			titleText += " *"
		}

		if fg.Collapsible {
			titleElement = app.Button().
				Style("background", "none").
				Style("border", "none").
				Style("padding", "0").
				Style("font-size", "16px").
				Style("font-weight", "600").
				Style("cursor", "pointer").
				Style("display", "flex").
				Style("align-items", "center").
				Style("gap", "8px").
				OnClick(func(ctx app.Context, e app.Event) {
					fg.Collapsed = !fg.Collapsed
					ctx.Reload()
				}).
				Body(
					app.Span().Text(IfElseString(fg.Collapsed, "▶", "▼")),
					app.Span().
						Style("font-size", "16px").
						Style("font-weight", "600").
						Style("color", IfElseString(fg.Required, "#d73a49", "inherit")).
						Text(titleText),
				)
		} else {
			titleElement = app.H3().
				Style("margin", "0 0 8px 0").
				Style("font-size", "16px").
				Style("font-weight", "600").
				Style("color", IfElseString(fg.Required, "#d73a49", "inherit")).
				Text(titleText)
		}

		headerContent = append(headerContent, titleElement)
	}

	if fg.Description != "" {
		headerContent = append(headerContent,
			app.P().
				Style("margin", "0 0 16px 0").
				Style("color", "#666").
				Style("font-size", "14px").
				Text(fg.Description),
		)
	}

	if len(headerContent) == 0 {
		return app.Div()
	}

	return app.Div().
		Class("prism-field-group-header").
		Style("margin-bottom", "16px").
		Body(headerContent...)
}

func (fg *FieldGroup) renderContent() app.UI {
	if fg.Collapsible && fg.Collapsed {
		return app.Div()
	}

	if len(fg.Fields) == 0 {
		return app.Div()
	}

	container := app.Div().Class("prism-field-group-content")

	switch fg.Layout {
	case "horizontal":
		container = container.
			Style("display", "flex").
			Style("gap", "16px").
			Style("flex-wrap", "wrap")
	case "grid":
		container = container.
			Style("display", "grid").
			Style("grid-template-columns", fmt.Sprintf("repeat(%d, 1fr)", fg.Columns)).
			Style("gap", "16px")
	default: // vertical
		container = container.
			Style("display", "flex").
			Style("flex-direction", "column").
			Style("gap", "16px")
	}

	fieldElements := make([]app.UI, len(fg.Fields))
	for i, field := range fg.Fields {
		fieldElements[i] = field
	}

	return container.Body(fieldElements...)
}

// FormBuilder helps build complex forms with field groups
type FormBuilder struct {
	Title       string
	Description string
	Groups      []*FieldGroup
	SubmitText  string
	CancelText  string
	ShowCancel  bool
	OnSubmit    func(app.Context, map[string]string) error
	OnCancel    func(app.Context)
}

// NewFormBuilder creates a new form builder
func NewFormBuilder() *FormBuilder {
	return &FormBuilder{
		SubmitText: "Submit",
		CancelText: "Cancel",
		ShowCancel: false,
	}
}

func (fb *FormBuilder) WithTitle(title string) *FormBuilder {
	fb.Title = title
	return fb
}

func (fb *FormBuilder) WithDescription(description string) *FormBuilder {
	fb.Description = description
	return fb
}

func (fb *FormBuilder) AddGroup(group *FieldGroup) *FormBuilder {
	fb.Groups = append(fb.Groups, group)
	return fb
}

func (fb *FormBuilder) WithSubmitText(text string) *FormBuilder {
	fb.SubmitText = text
	return fb
}

func (fb *FormBuilder) WithCancelText(text string) *FormBuilder {
	fb.CancelText = text
	return fb
}

func (fb *FormBuilder) WithShowCancel(show bool) *FormBuilder {
	fb.ShowCancel = show
	return fb
}

func (fb *FormBuilder) WithOnSubmit(fn func(app.Context, map[string]string) error) *FormBuilder {
	fb.OnSubmit = fn
	return fb
}

func (fb *FormBuilder) WithOnCancel(fn func(app.Context)) *FormBuilder {
	fb.OnCancel = fn
	return fb
}

func (fb *FormBuilder) Build() app.UI {
	return app.Form().
		Class("prism-form-builder").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()
			if fb.OnSubmit != nil {
				values := fb.collectValues()
				fb.OnSubmit(ctx, values)
			}
		}).
		Body(
			app.If(fb.Title != "", func() app.UI {
				return app.H1().
					Style("margin", "0 0 8px 0").
					Text(fb.Title)
			}),
			app.If(fb.Description != "", func() app.UI {
				return app.P().
					Style("margin", "0 0 24px 0").
					Style("color", "#666").
					Text(fb.Description)
			}),
			app.Div().
				Class("prism-form-groups").
				Body(
					app.Range(fb.Groups).Slice(func(i int) app.UI {
						return fb.Groups[i]
					}),
				),
			fb.renderActions(),
		)
}

func (fb *FormBuilder) renderActions() app.UI {
	return app.Div().
		Class("prism-form-actions").
		Style("display", "flex").
		Style("gap", "12px").
		Style("justify-content", "flex-end").
		Style("margin-top", "32px").
		Style("padding-top", "24px").
		Style("border-top", "1px solid #e0e0e0").
		Body(
			app.If(fb.ShowCancel, func() app.UI {
				return app.Button().
					Type("button").
					Style("padding", "8px 16px").
					Style("border", "1px solid #ccc").
					Style("background", "white").
					Style("cursor", "pointer").
					Text(fb.CancelText).
					OnClick(func(ctx app.Context, e app.Event) {
						if fb.OnCancel != nil {
							fb.OnCancel(ctx)
						}
					})
			}),
			app.Button().
				Type("submit").
				Style("padding", "8px 16px").
				Style("background", "#0066cc").
				Style("color", "white").
				Style("border", "none").
				Style("border-radius", "4px").
				Style("cursor", "pointer").
				Text(fb.SubmitText),
		)
}

func (fb *FormBuilder) collectValues() map[string]string {
	values := make(map[string]string)
	for _, group := range fb.Groups {
		for _, field := range group.Fields {
			values[field.Name] = field.Value
		}
	}
	return values
}

