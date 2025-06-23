package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// UserMenu represents a user menu dropdown
type UserMenu struct {
	app.Compo

	User         *User
	Avatar       string
	ShowAvatar   bool
	ShowName     bool
	MenuItems    []UserMenuItem
	IsOpen       bool
	Placement    string // "bottom-start", "bottom-end"
	OnToggle     func(bool)
}

// UserMenuItem represents an item in the user menu
type UserMenuItem struct {
	Label     string
	Icon      Icon
	Href      string
	Variant   string // "default", "negative"
	Disabled  bool
	Separator bool
	OnClick   func(app.Context)
}

// NewUserMenu creates a new user menu
func NewUserMenu() *UserMenu {
	return &UserMenu{
		ShowAvatar: true,
		ShowName:   true,
		Placement:  "bottom-end",
		MenuItems:  make([]UserMenuItem, 0),
	}
}

func (um *UserMenu) WithUser(user *User) *UserMenu {
	um.User = user
	return um
}

func (um *UserMenu) WithAvatar(avatar string) *UserMenu {
	um.Avatar = avatar
	return um
}

func (um *UserMenu) WithShowAvatar(show bool) *UserMenu {
	um.ShowAvatar = show
	return um
}

func (um *UserMenu) WithShowName(show bool) *UserMenu {
	um.ShowName = show
	return um
}

func (um *UserMenu) WithMenuItems(items ...UserMenuItem) *UserMenu {
	um.MenuItems = items
	return um
}

func (um *UserMenu) WithPlacement(placement string) *UserMenu {
	um.Placement = placement
	return um
}

func (um *UserMenu) WithOnToggle(fn func(bool)) *UserMenu {
	um.OnToggle = fn
	return um
}

func (um *UserMenu) AddMenuItem(item UserMenuItem) *UserMenu {
	um.MenuItems = append(um.MenuItems, item)
	return um
}

func (um *UserMenu) AddSeparator() *UserMenu {
	um.MenuItems = append(um.MenuItems, UserMenuItem{Separator: true})
	return um
}

func (um *UserMenu) Toggle() {
	um.IsOpen = !um.IsOpen
	if um.OnToggle != nil {
		um.OnToggle(um.IsOpen)
	}
}

func (um *UserMenu) Close() {
	um.IsOpen = false
	if um.OnToggle != nil {
		um.OnToggle(false)
	}
}

func (um *UserMenu) Render() app.UI {
	return app.Div().
		Class("prism-user-menu").
		Style("position", "relative").
		Body(
			um.renderTrigger(),
			app.If(um.IsOpen, func() app.UI {
				return um.renderMenu()
			}),
		)
}

func (um *UserMenu) renderTrigger() app.UI {
	triggerContent := []app.UI{}

	// Avatar
	if um.ShowAvatar {
		avatarSrc := um.Avatar
		if avatarSrc == "" && um.User != nil {
			avatarSrc = um.User.Avatar
		}
		
		if avatarSrc != "" {
			triggerContent = append(triggerContent,
				app.Img().
					Src(avatarSrc).
					Alt("User Avatar").
					Style("width", "32px").
					Style("height", "32px").
					Style("border-radius", "50%").
					Style("object-fit", "cover"),
			)
		} else {
			// Default avatar
			triggerContent = append(triggerContent,
				app.Div().
					Style("width", "32px").
					Style("height", "32px").
					Style("border-radius", "50%").
					Style("background", "#0066cc").
					Style("color", "white").
					Style("display", "flex").
					Style("align-items", "center").
					Style("justify-content", "center").
					Style("font-weight", "600").
					Style("font-size", "14px").
					Text(um.getUserInitials()),
			)
		}
	}

	// Name
	if um.ShowName && um.User != nil {
		triggerContent = append(triggerContent,
			app.Span().
				Style("margin-left", "8px").
				Style("font-weight", "500").
				Style("color", "#333").
				Text(um.User.GetDisplayName()),
		)
	}

	// Dropdown arrow
	triggerContent = append(triggerContent,
		app.Span().
			Style("margin-left", "8px").
			Style("color", "#666").
			Style("font-size", "12px").
			Style("transition", "transform 0.2s").
			Style("transform", IfElseString(um.IsOpen, "rotate(180deg)", "rotate(0deg)")).
			Text("▼"),
	)

	return sp.ActionButton().
		Quiet(true).
		OnClick(func(ctx app.Context, e app.Event) {
			um.Toggle()
			ctx.Reload()
		}).
		Style("display", "flex").
		Style("align-items", "center").
		Style("padding", "8px 12px").
		Style("border-radius", "4px").
		Style("cursor", "pointer").
		Style("transition", "background-color 0.2s").
		Body(triggerContent...)
}

func (um *UserMenu) renderMenu() app.UI {
	menuItems := []app.UI{}

	// User info header
	if um.User != nil {
		menuItems = append(menuItems,
			app.Div().
				Class("prism-user-menu__header").
				Style("padding", "12px 16px").
				Style("border-bottom", "1px solid #e0e0e0").
				Style("background", "#f8f9fa").
				Body(
					app.Div().
						Style("font-weight", "600").
						Style("font-size", "14px").
						Style("color", "#333").
						Text(um.User.GetDisplayName()),
					app.Div().
						Style("font-size", "12px").
						Style("color", "#666").
						Style("margin-top", "2px").
						Text(um.User.Email),
				),
		)
	}

	// Menu items
	for _, item := range um.MenuItems {
		if item.Separator {
			menuItems = append(menuItems,
				app.Hr().
					Style("margin", "4px 0").
					Style("border", "none").
					Style("border-top", "1px solid #e0e0e0"),
			)
			continue
		}

		menuItem := app.Div().
			Class("prism-user-menu__item").
			Style("padding", "8px 16px").
			Style("cursor", IfElseString(item.Disabled, "not-allowed", "pointer")).
			Style("color", um.getItemColor(item)).
			Style("opacity", IfElseString(item.Disabled, "0.5", "1")).
			Style("transition", "background-color 0.2s").
			Style("font-size", "14px").
			Body(
				app.If(item.Icon != "", func() app.UI {
					return app.Span().
						Style("margin-right", "8px").
						Style("font-size", "16px").
						Text(string(item.Icon))
				}),
				app.Span().Text(item.Label),
			)

		if !item.Disabled {
			menuItem = menuItem.
				OnMouseEnter(func(ctx app.Context, e app.Event) {
					// Add hover effect via CSS or JavaScript
				}).
				OnClick(func(ctx app.Context, e app.Event) {
					if item.OnClick != nil {
						item.OnClick(ctx)
					}
					if item.Href != "" {
						ctx.Navigate(item.Href)
					}
					um.Close()
					ctx.Reload()
				})
		}

		menuItems = append(menuItems, menuItem)
	}

	return app.Div().
		Class("prism-user-menu__dropdown").
		Style("position", "absolute").
		Style("top", "100%").
		Style("right", "0").
		Style("min-width", "200px").
		Style("background", "white").
		Style("border", "1px solid #e0e0e0").
		Style("border-radius", "4px").
		Style("box-shadow", "0 4px 12px rgba(0,0,0,0.15)").
		Style("z-index", "1000").
		Style("margin-top", "4px").
		OnClick(func(ctx app.Context, e app.Event) {
			// Prevent event propagation to avoid closing when clicking inside menu
			e.PreventDefault()
		}).
		Body(menuItems...)
}

func (um *UserMenu) getUserInitials() string {
	if um.User == nil {
		return "?"
	}
	
	name := um.User.GetDisplayName()
	if len(name) == 0 {
		return "?"
	}
	
	// Simple initials extraction
	initials := string(name[0])
	for i, char := range name {
		if i > 0 && name[i-1] == ' ' {
			initials += string(char)
			break
		}
	}
	
	return initials
}

func (um *UserMenu) getItemColor(item UserMenuItem) string {
	switch item.Variant {
	case "negative":
		return "#d73a49"
	default:
		return "#333"
	}
}

// Helper function to create a menu item
func NewUserMenuItem(label string) UserMenuItem {
	return UserMenuItem{
		Label:   label,
		Variant: "default",
	}
}

func (umi UserMenuItem) WithIcon(icon Icon) UserMenuItem {
	umi.Icon = icon
	return umi
}

func (umi UserMenuItem) WithHref(href string) UserMenuItem {
	umi.Href = href
	return umi
}

func (umi UserMenuItem) WithVariant(variant string) UserMenuItem {
	umi.Variant = variant
	return umi
}

func (umi UserMenuItem) WithDisabled(disabled bool) UserMenuItem {
	umi.Disabled = disabled
	return umi
}

func (umi UserMenuItem) WithOnClick(fn func(app.Context)) UserMenuItem {
	umi.OnClick = fn
	return umi
}

// Helper function to create default user menu items
func DefaultUserMenuItems() []UserMenuItem {
	return []UserMenuItem{
		NewUserMenuItem("Profile").WithIcon("👤").WithHref("/profile"),
		NewUserMenuItem("Settings").WithIcon("⚙️").WithHref("/settings"),
		{Separator: true},
		NewUserMenuItem("Help").WithIcon("❓").WithHref("/help"),
		NewUserMenuItem("Logout").WithIcon("🚪").WithVariant("negative").WithOnClick(func(ctx app.Context) {
			// Handle logout
			ctx.Navigate("/logout")
		}),
	}
}