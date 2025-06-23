package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// EmptyState represents an empty state display component
type EmptyState struct {
	app.Compo

	Icon        Icon
	Illustration string
	Title       string
	Description string
	Actions     []EmptyStateAction
	Size        string // "small", "medium", "large"
	Variant     string // "default", "error", "no-results", "no-data"
}

// EmptyStateAction represents an action button in the empty state
type EmptyStateAction struct {
	ID       string
	Label    string
	Icon     Icon
	Variant  string // "primary", "secondary", "link"
	Disabled bool
	OnClick  func(app.Context)
}

// NewEmptyState creates a new empty state component
func NewEmptyState() *EmptyState {
	return &EmptyState{
		Size:    "medium",
		Variant: "default",
		Actions: make([]EmptyStateAction, 0),
	}
}

func (es *EmptyState) WithIcon(icon Icon) *EmptyState {
	es.Icon = icon
	return es
}

func (es *EmptyState) WithIllustration(illustration string) *EmptyState {
	es.Illustration = illustration
	return es
}

func (es *EmptyState) WithTitle(title string) *EmptyState {
	es.Title = title
	return es
}

func (es *EmptyState) WithDescription(description string) *EmptyState {
	es.Description = description
	return es
}

func (es *EmptyState) WithActions(actions ...EmptyStateAction) *EmptyState {
	es.Actions = actions
	return es
}

func (es *EmptyState) WithSize(size string) *EmptyState {
	es.Size = size
	return es
}

func (es *EmptyState) WithVariant(variant string) *EmptyState {
	es.Variant = variant
	return es
}

func (es *EmptyState) AddAction(action EmptyStateAction) *EmptyState {
	es.Actions = append(es.Actions, action)
	return es
}

func (es *EmptyState) Render() app.UI {
	return app.Div().
		Class(es.getClasses()).
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("text-align", "center").
		Style("padding", es.getPadding()).
		Style("min-height", es.getMinHeight()).
		Body(
			es.renderIcon(),
			es.renderTitle(),
			es.renderDescription(),
			es.renderActions(),
		)
}

func (es *EmptyState) renderIcon() app.UI {
	if es.Illustration != "" {
		return app.Img().
			Src(es.Illustration).
			Alt("Empty state illustration").
			Class("prism-empty-state__illustration").
			Style("width", es.getIllustrationSize()).
			Style("height", "auto").
			Style("margin-bottom", "24px").
			Style("opacity", "0.8")
	}

	if es.Icon != "" {
		return app.Div().
			Class("prism-empty-state__icon").
			Style("font-size", es.getIconSize()).
			Style("color", es.getIconColor()).
			Style("margin-bottom", "24px").
			Style("opacity", "0.7").
			Text(string(es.Icon))
	}

	// Default icon based on variant
	defaultIcon := es.getDefaultIcon()
	if defaultIcon != "" {
		return app.Div().
			Class("prism-empty-state__icon").
			Style("font-size", es.getIconSize()).
			Style("color", es.getIconColor()).
			Style("margin-bottom", "24px").
			Style("opacity", "0.7").
			Text(defaultIcon)
	}

	return app.Div()
}

func (es *EmptyState) renderTitle() app.UI {
	if es.Title == "" {
		return app.Div()
	}

	return app.H3().
		Class("prism-empty-state__title").
		Style("margin", "0 0 16px 0").
		Style("font-size", es.getTitleFontSize()).
		Style("font-weight", "600").
		Style("color", es.getTitleColor()).
		Style("line-height", "1.3").
		Text(es.Title)
}

func (es *EmptyState) renderDescription() app.UI {
	if es.Description == "" {
		return app.Div()
	}

	return app.P().
		Class("prism-empty-state__description").
		Style("margin", "0 0 32px 0").
		Style("font-size", es.getDescriptionFontSize()).
		Style("color", "#666").
		Style("line-height", "1.5").
		Style("max-width", "400px").
		Text(es.Description)
}

func (es *EmptyState) renderActions() app.UI {
	if len(es.Actions) == 0 {
		return app.Div()
	}

	actionElements := []app.UI{}

	for _, action := range es.Actions {
		var btn app.UI

		if action.Variant == "link" {
			if !action.Disabled && action.OnClick != nil {
				btn = app.A().
					Style("color", "#0066cc").
					Style("text-decoration", "none").
					Style("font-weight", "500").
					Style("cursor", "pointer").
					Text(action.Label).
					OnClick(func(ctx app.Context, e app.Event) {
						e.PreventDefault()
						action.OnClick(ctx)
					})
			} else {
				btn = app.Span().
					Style("color", "#0066cc").
					Style("text-decoration", "none").
					Style("font-weight", "500").
					Style("cursor", IfElseString(action.Disabled, "not-allowed", "pointer")).
					Style("opacity", IfElseString(action.Disabled, "0.5", "1")).
					Text(action.Label)
			}
		} else {
			btnElement := sp.Button().
				Text(action.Label).
				Disabled(action.Disabled)

			if action.OnClick != nil {
				btnElement = btnElement.OnClick(func(ctx app.Context, e app.Event) {
					action.OnClick(ctx)
				})
			}
			btn = btnElement
		}

		actionElements = append(actionElements, btn)
	}

	return app.Div().
		Class("prism-empty-state__actions").
		Style("display", "flex").
		Style("gap", "12px").
		Style("flex-wrap", "wrap").
		Style("justify-content", "center").
		Body(actionElements...)
}

func (es *EmptyState) getClasses() string {
	classes := "prism-empty-state"
	
	if es.Variant != "default" {
		classes += " prism-empty-state--" + es.Variant
	}
	
	if es.Size != "medium" {
		classes += " prism-empty-state--" + es.Size
	}

	return classes
}

func (es *EmptyState) getPadding() string {
	switch es.Size {
	case "small":
		return "24px"
	case "large":
		return "64px 32px"
	default: // medium
		return "48px 24px"
	}
}

func (es *EmptyState) getMinHeight() string {
	switch es.Size {
	case "small":
		return "200px"
	case "large":
		return "400px"
	default: // medium
		return "300px"
	}
}

func (es *EmptyState) getIconSize() string {
	switch es.Size {
	case "small":
		return "48px"
	case "large":
		return "96px"
	default: // medium
		return "64px"
	}
}

func (es *EmptyState) getIllustrationSize() string {
	switch es.Size {
	case "small":
		return "120px"
	case "large":
		return "240px"
	default: // medium
		return "180px"
	}
}

func (es *EmptyState) getTitleFontSize() string {
	switch es.Size {
	case "small":
		return "16px"
	case "large":
		return "24px"
	default: // medium
		return "20px"
	}
}

func (es *EmptyState) getDescriptionFontSize() string {
	switch es.Size {
	case "small":
		return "14px"
	case "large":
		return "16px"
	default: // medium
		return "16px"
	}
}

func (es *EmptyState) getIconColor() string {
	switch es.Variant {
	case "error":
		return "#d73a49"
	case "no-results":
		return "#0066cc"
	default:
		return "#666"
	}
}

func (es *EmptyState) getTitleColor() string {
	switch es.Variant {
	case "error":
		return "#d73a49"
	default:
		return "#333"
	}
}

func (es *EmptyState) getDefaultIcon() string {
	switch es.Variant {
	case "error":
		return "⚠️"
	case "no-results":
		return "🔍"
	case "no-data":
		return "📄"
	default:
		return "📭"
	}
}

// Helper function to create an empty state action
func NewEmptyStateAction(id, label string) EmptyStateAction {
	return EmptyStateAction{
		ID:      id,
		Label:   label,
		Variant: "secondary",
	}
}

func (esa EmptyStateAction) WithIcon(icon Icon) EmptyStateAction {
	esa.Icon = icon
	return esa
}

func (esa EmptyStateAction) WithVariant(variant string) EmptyStateAction {
	esa.Variant = variant
	return esa
}

func (esa EmptyStateAction) WithDisabled(disabled bool) EmptyStateAction {
	esa.Disabled = disabled
	return esa
}

func (esa EmptyStateAction) WithOnClick(fn func(app.Context)) EmptyStateAction {
	esa.OnClick = fn
	return esa
}

// Helper functions for common empty state patterns

// NoDataEmptyState creates an empty state for when there's no data
func NoDataEmptyState(title, description string) *EmptyState {
	return NewEmptyState().
		WithVariant("no-data").
		WithTitle(title).
		WithDescription(description)
}

// NoResultsEmptyState creates an empty state for when search/filter returns no results
func NoResultsEmptyState(query string) *EmptyState {
	title := "No results found"
	description := "Try adjusting your search or filter criteria."
	
	if query != "" {
		description = "No results found for \"" + query + "\". " + description
	}

	return NewEmptyState().
		WithVariant("no-results").
		WithTitle(title).
		WithDescription(description).
		WithActions(
			NewEmptyStateAction("clear", "Clear filters").
				WithVariant("link"),
		)
}

// ErrorEmptyState creates an empty state for error scenarios
func ErrorEmptyState(title, description string, onRetry func(app.Context)) *EmptyState {
	emptyState := NewEmptyState().
		WithVariant("error").
		WithTitle(title).
		WithDescription(description)

	if onRetry != nil {
		emptyState = emptyState.WithActions(
			NewEmptyStateAction("retry", "Try again").
				WithVariant("primary").
				WithOnClick(onRetry),
		)
	}

	return emptyState
}

// FirstTimeEmptyState creates an empty state for first-time users
func FirstTimeEmptyState(title, description, actionLabel string, onAction func(app.Context)) *EmptyState {
	return NewEmptyState().
		WithTitle(title).
		WithDescription(description).
		WithActions(
			NewEmptyStateAction("get-started", actionLabel).
				WithVariant("primary").
				WithOnClick(onAction),
		)
}

// LoadingEmptyState creates an empty state with loading indicator
func LoadingEmptyState(message string) *EmptyState {
	if message == "" {
		message = "Loading..."
	}

	return NewEmptyState().
		WithIcon("").
		WithTitle(message).
		WithDescription("")
		// Custom render for loading spinner would go here
}

// ComingSoonEmptyState creates an empty state for coming soon features
func ComingSoonEmptyState(feature string) *EmptyState {
	return NewEmptyState().
		WithTitle("Coming Soon").
		WithDescription(feature + " will be available in a future update.").
		WithActions(
			NewEmptyStateAction("notify", "Notify me").
				WithVariant("link"),
		)
}