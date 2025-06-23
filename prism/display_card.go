package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Card represents a content card component
type Card struct {
	app.Compo

	Title         string
	Subtitle      string
	Content       app.UI
	Image         string
	ImageAlt      string
	Actions       []CardAction
	Footer        app.UI
	Header        app.UI
	Variant       string // "default", "quiet", "outlined"
	Size          string // "small", "medium", "large"
	Horizontal    bool
	Clickable     bool
	Selected      bool
	Loading       bool
	OnClick       func(app.Context)
	OnActionClick func(string, app.Context)
}

// CardAction represents an action button in the card
type CardAction struct {
	ID       string
	Label    string
	Icon     Icon
	Variant  string // "primary", "secondary", "negative"
	Disabled bool
	OnClick  func(app.Context)
}

// NewCard creates a new card component
func NewCard() *Card {
	return &Card{
		Variant: "default",
		Size:    "medium",
		Actions: make([]CardAction, 0),
	}
}

func (c *Card) WithTitle(title string) *Card {
	c.Title = title
	return c
}

func (c *Card) WithSubtitle(subtitle string) *Card {
	c.Subtitle = subtitle
	return c
}

func (c *Card) WithContent(content app.UI) *Card {
	c.Content = content
	return c
}

func (c *Card) WithImage(src, alt string) *Card {
	c.Image = src
	c.ImageAlt = alt
	return c
}

func (c *Card) WithActions(actions ...CardAction) *Card {
	c.Actions = actions
	return c
}

func (c *Card) WithFooter(footer app.UI) *Card {
	c.Footer = footer
	return c
}

func (c *Card) WithHeader(header app.UI) *Card {
	c.Header = header
	return c
}

func (c *Card) WithVariant(variant string) *Card {
	c.Variant = variant
	return c
}

func (c *Card) WithSize(size string) *Card {
	c.Size = size
	return c
}

func (c *Card) WithHorizontal(horizontal bool) *Card {
	c.Horizontal = horizontal
	return c
}

func (c *Card) WithClickable(clickable bool) *Card {
	c.Clickable = clickable
	return c
}

func (c *Card) WithSelected(selected bool) *Card {
	c.Selected = selected
	return c
}

func (c *Card) WithLoading(loading bool) *Card {
	c.Loading = loading
	return c
}

func (c *Card) WithOnClick(fn func(app.Context)) *Card {
	c.OnClick = fn
	return c
}

func (c *Card) WithOnActionClick(fn func(string, app.Context)) *Card {
	c.OnActionClick = fn
	return c
}

func (c *Card) AddAction(action CardAction) *Card {
	c.Actions = append(c.Actions, action)
	return c
}

func (c *Card) Render() app.UI {
	classes := c.getCardClasses()
	styles := c.getCardStyles()

	cardElement := app.Div().
		Class(classes).
		Style("display", "flex").
		Style("flex-direction", IfElseString(c.Horizontal, "row", "column")).
		Style("background", "white").
		Style("border-radius", "4px").
		Style("box-shadow", c.getBoxShadow()).
		Style("border", c.getBorder()).
		Style("overflow", "hidden").
		Style("transition", "all 0.2s ease").
		Style("position", "relative")

	// Apply custom styles
	for key, value := range styles {
		cardElement = cardElement.Style(key, value)
	}

	// Add click handler if clickable
	if c.Clickable && c.OnClick != nil {
		cardElement = cardElement.
			Style("cursor", "pointer").
			OnClick(func(ctx app.Context, e app.Event) {
				c.OnClick(ctx)
			}).
			OnMouseEnter(func(ctx app.Context, e app.Event) {
				// Hover effect handled by CSS
			})
	}

	content := []app.UI{}

	// Image section
	if c.Image != "" {
		content = append(content, c.renderImage())
	}

	// Main content container
	mainContent := []app.UI{}

	// Header
	if c.Header != nil {
		mainContent = append(mainContent, c.Header)
	}

	// Title/Subtitle section
	if c.Title != "" || c.Subtitle != "" {
		mainContent = append(mainContent, c.renderTitleSection())
	}

	// Content
	if c.Content != nil {
		mainContent = append(mainContent, c.renderContentSection())
	}

	// Actions
	if len(c.Actions) > 0 {
		mainContent = append(mainContent, c.renderActions())
	}

	// Footer
	if c.Footer != nil {
		mainContent = append(mainContent, c.Footer)
	}

	// Add main content to card
	if len(mainContent) > 0 {
		content = append(content,
			app.Div().
				Class("prism-card__body").
				Style("flex", "1").
				Style("padding", c.getContentPadding()).
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "16px").
				Body(mainContent...),
		)
	}

	// Loading overlay
	if c.Loading {
		content = append(content, c.renderLoadingOverlay())
	}

	return cardElement.Body(content...)
}

func (c *Card) renderImage() app.UI {
	imageStyle := map[string]string{
		"width":       "100%",
		"object-fit":  "cover",
		"display":     "block",
	}

	if c.Horizontal {
		imageStyle["width"] = "200px"
		imageStyle["height"] = "100%"
		imageStyle["flex-shrink"] = "0"
	} else {
		imageStyle["height"] = "200px"
	}

	img := app.Img().
		Src(c.Image).
		Alt(c.ImageAlt).
		Class("prism-card__image")

	for key, value := range imageStyle {
		img = img.Style(key, value)
	}

	return img
}

func (c *Card) renderTitleSection() app.UI {
	titleElements := []app.UI{}

	if c.Title != "" {
		titleElements = append(titleElements,
			app.H3().
				Class("prism-card__title").
				Style("margin", "0").
				Style("font-size", c.getTitleFontSize()).
				Style("font-weight", "600").
				Style("color", "#333").
				Style("line-height", "1.3").
				Text(c.Title),
		)
	}

	if c.Subtitle != "" {
		titleElements = append(titleElements,
			app.P().
				Class("prism-card__subtitle").
				Style("margin", "0").
				Style("font-size", "14px").
				Style("color", "#666").
				Style("line-height", "1.4").
				Style("margin-top", "4px").
				Text(c.Subtitle),
		)
	}

	return app.Div().
		Class("prism-card__header").
		Body(titleElements...)
}

func (c *Card) renderContentSection() app.UI {
	return app.Div().
		Class("prism-card__content").
		Style("flex", "1").
		Body(c.Content)
}

func (c *Card) renderActions() app.UI {
	actionElements := []app.UI{}

	for _, action := range c.Actions {
		btn := sp.Button().
			Text(action.Label).
			Disabled(action.Disabled)

		if action.OnClick != nil {
			btn = btn.OnClick(func(ctx app.Context, e app.Event) {
				action.OnClick(ctx)
				if c.OnActionClick != nil {
					c.OnActionClick(action.ID, ctx)
				}
			})
		}

		actionElements = append(actionElements, btn)
	}

	return app.Div().
		Class("prism-card__actions").
		Style("display", "flex").
		Style("gap", "8px").
		Style("flex-wrap", "wrap").
		Style("margin-top", "auto").
		Body(actionElements...)
}

func (c *Card) renderLoadingOverlay() app.UI {
	return app.Div().
		Class("prism-card__loading").
		Style("position", "absolute").
		Style("top", "0").
		Style("left", "0").
		Style("right", "0").
		Style("bottom", "0").
		Style("background", "rgba(255, 255, 255, 0.8)").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("z-index", "1").
		Body(
			sp.ProgressCircle().
				Indeterminate(true).
				Size("medium"),
		)
}

func (c *Card) getCardClasses() string {
	classes := "prism-card"
	
	if c.Variant != "default" {
		classes += " prism-card--" + c.Variant
	}
	
	if c.Size != "medium" {
		classes += " prism-card--" + c.Size
	}
	
	if c.Horizontal {
		classes += " prism-card--horizontal"
	}
	
	if c.Clickable {
		classes += " prism-card--clickable"
	}
	
	if c.Selected {
		classes += " prism-card--selected"
	}

	return classes
}

func (c *Card) getCardStyles() map[string]string {
	styles := make(map[string]string)

	// Size-based styles
	switch c.Size {
	case "small":
		styles["max-width"] = "300px"
	case "large":
		styles["max-width"] = "600px"
	default: // medium
		styles["max-width"] = "400px"
	}

	// Selected state
	if c.Selected {
		styles["border"] = "2px solid #0066cc"
		styles["box-shadow"] = "0 0 0 1px #0066cc"
	}

	// Clickable hover effect
	if c.Clickable {
		styles["cursor"] = "pointer"
	}

	return styles
}

func (c *Card) getBoxShadow() string {
	switch c.Variant {
	case "quiet":
		return "none"
	case "outlined":
		return "none"
	default:
		return "0 2px 8px rgba(0, 0, 0, 0.1)"
	}
}

func (c *Card) getBorder() string {
	switch c.Variant {
	case "outlined":
		return "1px solid #e0e0e0"
	case "quiet":
		return "none"
	default:
		return "1px solid #f0f0f0"
	}
}

func (c *Card) getContentPadding() string {
	switch c.Size {
	case "small":
		return "12px"
	case "large":
		return "24px"
	default: // medium
		return "16px"
	}
}

func (c *Card) getTitleFontSize() string {
	switch c.Size {
	case "small":
		return "14px"
	case "large":
		return "20px"
	default: // medium
		return "16px"
	}
}

// Helper function to create a card action
func NewCardAction(id, label string) CardAction {
	return CardAction{
		ID:      id,
		Label:   label,
		Variant: "secondary",
	}
}

func (ca CardAction) WithIcon(icon Icon) CardAction {
	ca.Icon = icon
	return ca
}

func (ca CardAction) WithVariant(variant string) CardAction {
	ca.Variant = variant
	return ca
}

func (ca CardAction) WithDisabled(disabled bool) CardAction {
	ca.Disabled = disabled
	return ca
}

func (ca CardAction) WithOnClick(fn func(app.Context)) CardAction {
	ca.OnClick = fn
	return ca
}

// Helper functions for common card patterns

// InfoCard creates a simple info card with title and content
func InfoCard(title, content string) *Card {
	return NewCard().
		WithTitle(title).
		WithContent(app.P().
			Style("margin", "0").
			Style("line-height", "1.5").
			Text(content))
}

// ActionCard creates a card with a primary action
func ActionCard(title, content, actionLabel string, onAction func(app.Context)) *Card {
	return NewCard().
		WithTitle(title).
		WithContent(app.P().
			Style("margin", "0").
			Style("line-height", "1.5").
			Text(content)).
		WithActions(
			NewCardAction("primary", actionLabel).
				WithVariant("primary").
				WithOnClick(onAction),
		)
}

// ImageCard creates a card with an image
func ImageCard(title, subtitle, imageSrc, imageAlt string) *Card {
	return NewCard().
		WithTitle(title).
		WithSubtitle(subtitle).
		WithImage(imageSrc, imageAlt)
}

// StatsCard creates a card for displaying statistics
func StatsCard(title, value, subtitle string) *Card {
	return NewCard().
		WithTitle(title).
		WithContent(
			app.Div().
				Style("text-align", "center").
				Body(
					app.Div().
						Style("font-size", "32px").
						Style("font-weight", "700").
						Style("color", "#0066cc").
						Style("margin-bottom", "8px").
						Text(value),
					app.Div().
						Style("font-size", "14px").
						Style("color", "#666").
						Text(subtitle),
				),
		)
}