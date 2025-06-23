package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TopNav represents the top navigation bar
type TopNav struct {
	app.Compo

	Title        string
	Logo         string
	UserMenu     *UserMenu
	SearchBar    *SearchBar
	Actions      []TopNavAction
	Theme        string // "light", "dark", "auto"
	Sticky       bool
	OnThemeToggle func(theme string)
}

// TopNavAction represents an action button in the top nav
type TopNavAction struct {
	Icon     Icon
	Label    string
	Variant  string // "primary", "secondary", "negative"
	Disabled bool
	OnClick  func(app.Context)
}

// SearchBar represents a search input in the top nav
type SearchBar struct {
	Placeholder string
	Value       string
	OnSearch    func(query string)
	OnClear     func()
}

// NewTopNav creates a new top navigation bar
func NewTopNav() *TopNav {
	return &TopNav{
		Theme:  "light",
		Sticky: true,
		Actions: make([]TopNavAction, 0),
	}
}

func (tn *TopNav) WithTitle(title string) *TopNav {
	tn.Title = title
	return tn
}

func (tn *TopNav) WithLogo(logo string) *TopNav {
	tn.Logo = logo
	return tn
}

func (tn *TopNav) WithUserMenu(userMenu *UserMenu) *TopNav {
	tn.UserMenu = userMenu
	return tn
}

func (tn *TopNav) WithSearchBar(searchBar *SearchBar) *TopNav {
	tn.SearchBar = searchBar
	return tn
}

func (tn *TopNav) WithActions(actions ...TopNavAction) *TopNav {
	tn.Actions = actions
	return tn
}

func (tn *TopNav) WithTheme(theme string) *TopNav {
	tn.Theme = theme
	return tn
}

func (tn *TopNav) WithSticky(sticky bool) *TopNav {
	tn.Sticky = sticky
	return tn
}

func (tn *TopNav) WithOnThemeToggle(fn func(theme string)) *TopNav {
	tn.OnThemeToggle = fn
	return tn
}

func (tn *TopNav) AddAction(action TopNavAction) *TopNav {
	tn.Actions = append(tn.Actions, action)
	return tn
}

func (tn *TopNav) Render() app.UI {
	classes := "prism-top-nav"
	if tn.Sticky {
		classes += " prism-top-nav--sticky"
	}

	return app.Header().
		Class(classes).
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "space-between").
		Style("padding", "0 24px").
		Style("height", "64px").
		Style("background", IfElseString(tn.Theme == "dark", "#1a1a1a", "white")).
		Style("border-bottom", "1px solid #e0e0e0").
		Style("box-shadow", "0 1px 3px rgba(0,0,0,0.1)").
		Style("z-index", "1000").
		Style("position", IfElseString(tn.Sticky, "sticky", "relative")).
		Style("top", IfElseString(tn.Sticky, "0", "auto")).
		Body(
			tn.renderLeftSection(),
			tn.renderCenterSection(),
			tn.renderRightSection(),
		)
}

func (tn *TopNav) renderLeftSection() app.UI {
	leftItems := []app.UI{}

	// Logo
	if tn.Logo != "" {
		leftItems = append(leftItems,
			app.Img().
				Src(tn.Logo).
				Alt("Logo").
				Style("height", "32px").
				Style("width", "auto").
				Style("margin-right", "16px"),
		)
	}

	// Title
	if tn.Title != "" {
		leftItems = append(leftItems,
			app.H1().
				Style("margin", "0").
				Style("font-size", "18px").
				Style("font-weight", "600").
				Style("color", IfElseString(tn.Theme == "dark", "white", "#333")).
				Text(tn.Title),
		)
	}

	return app.Div().
		Class("prism-top-nav__left").
		Style("display", "flex").
		Style("align-items", "center").
		Style("flex", "0 0 auto").
		Body(leftItems...)
}

func (tn *TopNav) renderCenterSection() app.UI {
	if tn.SearchBar == nil {
		return app.Div()
	}

	return app.Div().
		Class("prism-top-nav__center").
		Style("display", "flex").
		Style("flex", "1 1 auto").
		Style("justify-content", "center").
		Style("max-width", "600px").
		Style("margin", "0 24px").
		Body(tn.renderSearchBar())
}

func (tn *TopNav) renderSearchBar() app.UI {
	if tn.SearchBar == nil {
		return app.Div()
	}

	return app.Div().
		Class("prism-search-bar").
		Style("position", "relative").
		Style("width", "100%").
		Body(
			sp.Textfield().
				Placeholder(tn.SearchBar.Placeholder).
				Value(tn.SearchBar.Value).
				Style("width", "100%").
				OnInput(func(ctx app.Context, e app.Event) {
					value := ctx.JSSrc().Get("value").String()
					if tn.SearchBar.OnSearch != nil {
						tn.SearchBar.OnSearch(value)
					}
				}),
			
			// Search icon
			app.If(tn.SearchBar.Value == "", func() app.UI {
				return app.Div().
					Style("position", "absolute").
					Style("right", "12px").
					Style("top", "50%").
					Style("transform", "translateY(-50%)").
					Style("color", "#666").
					Style("pointer-events", "none").
					Text("🔍")
			}),

			// Clear button
			app.If(tn.SearchBar.Value != "", func() app.UI {
				return app.Button().
					Style("position", "absolute").
					Style("right", "8px").
					Style("top", "50%").
					Style("transform", "translateY(-50%)").
					Style("background", "none").
					Style("border", "none").
					Style("cursor", "pointer").
					Style("padding", "4px").
					Style("color", "#666").
					Text("✕").
					OnClick(func(ctx app.Context, e app.Event) {
						if tn.SearchBar.OnClear != nil {
							tn.SearchBar.OnClear()
						}
					})
			}),
		)
}

func (tn *TopNav) renderRightSection() app.UI {
	rightItems := []app.UI{}

	// Theme toggle button
	rightItems = append(rightItems,
		sp.ActionButton().
			Quiet(true).
			OnClick(func(ctx app.Context, e app.Event) {
				newTheme := "light"
				if tn.Theme == "light" {
					newTheme = "dark"
				}
				if tn.OnThemeToggle != nil {
					tn.OnThemeToggle(newTheme)
				}
			}).
			Body(
				app.Span().
					Style("font-size", "18px").
					Text(IfElseString(tn.Theme == "dark", "☀️", "🌙")),
			),
	)

	// Action buttons
	for _, action := range tn.Actions {
		actionBtn := sp.ActionButton().
			Quiet(action.Variant != "primary").
			Disabled(action.Disabled)

		if action.OnClick != nil {
			actionBtn = actionBtn.OnClick(func(ctx app.Context, e app.Event) {
				action.OnClick(ctx)
			})
		}

		if action.Label != "" {
			actionBtn = actionBtn.Body(app.Text(action.Label))
		}

		rightItems = append(rightItems, actionBtn)
	}

	// User menu
	if tn.UserMenu != nil {
		rightItems = append(rightItems, tn.UserMenu)
	}

	return app.Div().
		Class("prism-top-nav__right").
		Style("display", "flex").
		Style("align-items", "center").
		Style("gap", "8px").
		Style("flex", "0 0 auto").
		Body(rightItems...)
}

// Helper function to create a search bar
func NewSearchBar() *SearchBar {
	return &SearchBar{
		Placeholder: "Search...",
	}
}

func (sb *SearchBar) WithPlaceholder(placeholder string) *SearchBar {
	sb.Placeholder = placeholder
	return sb
}

func (sb *SearchBar) WithValue(value string) *SearchBar {
	sb.Value = value
	return sb
}

func (sb *SearchBar) WithOnSearch(fn func(query string)) *SearchBar {
	sb.OnSearch = fn
	return sb
}

func (sb *SearchBar) WithOnClear(fn func()) *SearchBar {
	sb.OnClear = fn
	return sb
}

// Helper function to create a top nav action
func NewTopNavAction(icon Icon, label string) TopNavAction {
	return TopNavAction{
		Icon:    icon,
		Label:   label,
		Variant: "secondary",
	}
}

func (tna TopNavAction) WithVariant(variant string) TopNavAction {
	tna.Variant = variant
	return tna
}

func (tna TopNavAction) WithDisabled(disabled bool) TopNavAction {
	tna.Disabled = disabled
	return tna
}

func (tna TopNavAction) WithOnClick(fn func(app.Context)) TopNavAction {
	tna.OnClick = fn
	return tna
}