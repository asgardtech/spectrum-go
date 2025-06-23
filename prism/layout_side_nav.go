package prism

import (
	"fmt"
	
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// SideNav represents a side navigation component
type SideNav struct {
	app.Compo

	Title        string
	Logo         string
	Items        []SideNavItem
	Groups       []SideNavGroup
	IsCollapsed  bool
	IsOpen       bool
	Width        string
	Position     string // "left" or "right"
	Variant      string // "fixed", "overlay", "push"
	ShowToggle   bool
	ShowFooter   bool
	FooterItems  []SideNavItem
	OnToggle     func(bool)
	OnItemClick  func(item SideNavItem)
}

// SideNavItem represents a navigation item
type SideNavItem struct {
	ID       string
	Label    string
	Icon     Icon
	Href     string
	Active   bool
	Disabled bool
	Badge    string
	Children []SideNavItem
	OnClick  func(app.Context)
}

// SideNavGroup represents a group of navigation items
type SideNavGroup struct {
	ID       string
	Label    string
	Items    []SideNavItem
	Expanded bool
}

// NewSideNav creates a new side navigation component
func NewSideNav() *SideNav {
	return &SideNav{
		Width:       "240px",
		Position:    "left",
		Variant:     "fixed",
		IsOpen:      true,
		ShowToggle:  true,
		ShowFooter:  false,
		Items:       make([]SideNavItem, 0),
		Groups:      make([]SideNavGroup, 0),
		FooterItems: make([]SideNavItem, 0),
	}
}

func (sn *SideNav) WithTitle(title string) *SideNav {
	sn.Title = title
	return sn
}

func (sn *SideNav) WithLogo(logo string) *SideNav {
	sn.Logo = logo
	return sn
}

func (sn *SideNav) WithItems(items ...SideNavItem) *SideNav {
	sn.Items = items
	return sn
}

func (sn *SideNav) WithGroups(groups ...SideNavGroup) *SideNav {
	sn.Groups = groups
	return sn
}

func (sn *SideNav) WithWidth(width string) *SideNav {
	sn.Width = width
	return sn
}

func (sn *SideNav) WithPosition(position string) *SideNav {
	sn.Position = position
	return sn
}

func (sn *SideNav) WithVariant(variant string) *SideNav {
	sn.Variant = variant
	return sn
}

func (sn *SideNav) WithCollapsed(collapsed bool) *SideNav {
	sn.IsCollapsed = collapsed
	return sn
}

func (sn *SideNav) WithOpen(open bool) *SideNav {
	sn.IsOpen = open
	return sn
}

func (sn *SideNav) WithShowToggle(showToggle bool) *SideNav {
	sn.ShowToggle = showToggle
	return sn
}

func (sn *SideNav) WithShowFooter(showFooter bool) *SideNav {
	sn.ShowFooter = showFooter
	return sn
}

func (sn *SideNav) WithFooterItems(items ...SideNavItem) *SideNav {
	sn.FooterItems = items
	return sn
}

func (sn *SideNav) WithOnToggle(fn func(bool)) *SideNav {
	sn.OnToggle = fn
	return sn
}

func (sn *SideNav) WithOnItemClick(fn func(item SideNavItem)) *SideNav {
	sn.OnItemClick = fn
	return sn
}

func (sn *SideNav) AddItem(item SideNavItem) *SideNav {
	sn.Items = append(sn.Items, item)
	return sn
}

func (sn *SideNav) AddGroup(group SideNavGroup) *SideNav {
	sn.Groups = append(sn.Groups, group)
	return sn
}

func (sn *SideNav) Toggle() {
	sn.IsOpen = !sn.IsOpen
	if sn.OnToggle != nil {
		sn.OnToggle(sn.IsOpen)
	}
}

func (sn *SideNav) Collapse() {
	sn.IsCollapsed = !sn.IsCollapsed
}

func (sn *SideNav) Close() {
	sn.IsOpen = false
	if sn.OnToggle != nil {
		sn.OnToggle(sn.IsOpen)
	}
}

func (sn *SideNav) Open() {
	sn.IsOpen = true
	if sn.OnToggle != nil {
		sn.OnToggle(sn.IsOpen)
	}
}

func (sn *SideNav) Render() app.UI {
	return app.Div().
		Class(sn.getClasses()).
		Style("width", IfElseString(sn.IsOpen, sn.Width, "0")).
		Style("position", IfElseString(sn.Variant == "fixed", "fixed", "relative")).
		Style("top", IfElseString(sn.Variant == "fixed", "0", "auto")).
		Style("left", IfElseString(sn.Position == "left", "0", "auto")).
		Style("right", IfElseString(sn.Position == "right", "0", "auto")).
		Style("height", IfElseString(sn.Variant == "fixed", "100vh", "auto")).
		Style("background-color", "#fff").
		Style("border-right", IfElseString(sn.Position == "left", "1px solid #e1e1e1", "none")).
		Style("border-left", IfElseString(sn.Position == "right", "1px solid #e1e1e1", "none")).
		Style("box-shadow", sn.getShadow()).
		Style("transform", sn.getTransform()).
		Style("transition", "all 0.3s ease").
		Style("overflow", "hidden").
		Style("z-index", "1000").
		Body(
			sn.renderHeader(),
			sn.renderContent(),
			sn.renderFooter(),
		)
}

func (sn *SideNav) renderHeader() app.UI {
	if sn.Title == "" && sn.Logo == "" && !sn.ShowToggle {
		return app.Div()
	}

	return app.Div().
		Class("prism-sidenav__header").
		Style("padding", "16px").
		Style("border-bottom", "1px solid #e1e1e1").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "space-between").
		Body(
			app.Div().
				Style("display", "flex").
				Style("align-items", "center").
				Style("gap", "12px").
				Body(
					sn.renderLogo(),
					sn.renderTitle(),
				),
			sn.renderToggleButton(),
		)
}

func (sn *SideNav) renderLogo() app.UI {
	if sn.Logo == "" {
		return app.Div()
	}

	return app.Img().
		Src(sn.Logo).
		Alt("Logo").
		Class("prism-sidenav__logo").
		Style("width", "32px").
		Style("height", "32px").
		Style("object-fit", "contain")
}

func (sn *SideNav) renderTitle() app.UI {
	if sn.Title == "" {
		return app.Div()
	}

	return app.H2().
		Class("prism-sidenav__title").
		Style("margin", "0").
		Style("font-size", "18px").
		Style("font-weight", "600").
		Style("color", "#333").
		Style("white-space", "nowrap").
		Style("overflow", "hidden").
		Style("text-overflow", "ellipsis").
		Text(sn.Title)
}

func (sn *SideNav) renderToggleButton() app.UI {
	if !sn.ShowToggle {
		return app.Div()
	}

	return sp.Button().
		Text(IfElseString(sn.IsCollapsed, "▶", "◀")).
		OnClick(func(ctx app.Context, e app.Event) {
			sn.Collapse()
		}).
		Style("min-width", "auto").
		Style("padding", "4px 8px").
		Style("font-size", "12px")
}

func (sn *SideNav) renderContent() app.UI {
	return app.Div().
		Class("prism-sidenav__content").
		Style("flex", "1").
		Style("overflow-y", "auto").
		Style("padding", "8px 0").
		Body(
			sn.renderItems(),
			sn.renderGroups(),
		)
}

func (sn *SideNav) renderItems() app.UI {
	if len(sn.Items) == 0 {
		return app.Div()
	}

	items := []app.UI{}
	for _, item := range sn.Items {
		items = append(items, sn.renderNavItem(item, 0))
	}

	return app.Div().
		Class("prism-sidenav__items").
		Body(items...)
}

func (sn *SideNav) renderGroups() app.UI {
	if len(sn.Groups) == 0 {
		return app.Div()
	}

	groups := []app.UI{}
	for _, group := range sn.Groups {
		groups = append(groups, sn.renderNavGroup(group))
	}

	return app.Div().
		Class("prism-sidenav__groups").
		Body(groups...)
}

func (sn *SideNav) renderNavGroup(group SideNavGroup) app.UI {
	return app.Div().
		Class("prism-sidenav__group").
		Style("margin-bottom", "16px").
		Body(
			app.Div().
				Class("prism-sidenav__group-header").
				Style("padding", "8px 16px").
				Style("font-size", "12px").
				Style("font-weight", "600").
				Style("color", "#666").
				Style("text-transform", "uppercase").
				Style("letter-spacing", "0.5px").
				Style("cursor", "pointer").
				Text(group.Label).
				OnClick(func(ctx app.Context, e app.Event) {
					// Toggle group expansion
					for i := range sn.Groups {
						if sn.Groups[i].ID == group.ID {
							sn.Groups[i].Expanded = !sn.Groups[i].Expanded
							break
						}
					}
				}),
			app.If(group.Expanded, func() app.UI {
				return app.Div().
					Class("prism-sidenav__group-items").
					Body(sn.renderGroupItems(group.Items)...)
			}),
		)
}

func (sn *SideNav) renderGroupItems(items []SideNavItem) []app.UI {
	elements := []app.UI{}
	for _, item := range items {
		elements = append(elements, sn.renderNavItem(item, 1))
	}
	return elements
}

func (sn *SideNav) renderNavItem(item SideNavItem, level int) app.UI {
	paddingLeft := 16 + (level * 24)

	element := app.Div().
		Class(sn.getItemClasses(item)).
		Style("padding", fmt.Sprintf("%dpx %dpx %dpx %dpx", 8, 16, 8, paddingLeft)).
		Style("display", "flex").
		Style("align-items", "center").
		Style("gap", "12px").
		Style("cursor", IfElseString(item.Disabled, "not-allowed", "pointer")).
		Style("opacity", IfElseString(item.Disabled, "0.5", "1")).
		Style("transition", "all 0.2s ease").
		Body(
			sn.renderItemIcon(item),
			sn.renderItemLabel(item),
			sn.renderItemBadge(item),
		)

	if !item.Disabled {
		element = element.OnClick(func(ctx app.Context, e app.Event) {
			if item.OnClick != nil {
				item.OnClick(ctx)
			}
			if sn.OnItemClick != nil {
				sn.OnItemClick(item)
			}
			if item.Href != "" {
				ctx.Navigate(item.Href)
			}
		})
	}

	// Add hover effects
	if !item.Disabled {
		element = element.
			Style("&:hover", "background-color: #f5f5f5")
	}

	return element
}

func (sn *SideNav) renderItemIcon(item SideNavItem) app.UI {
	if item.Icon == "" {
		return app.Span().
			Style("width", "20px").
			Style("height", "20px").
			Style("flex-shrink", "0")
	}

	return app.Span().
		Class("prism-sidenav__item-icon").
		Style("width", "20px").
		Style("height", "20px").
		Style("flex-shrink", "0").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("font-size", "16px").
		Style("color", IfElseString(item.Active, "#0066cc", "#666")).
		Text(string(item.Icon))
}

func (sn *SideNav) renderItemLabel(item SideNavItem) app.UI {
	return app.Span().
		Class("prism-sidenav__item-label").
		Style("flex", "1").
		Style("font-size", "14px").
		Style("font-weight", IfElseString(item.Active, "600", "400")).
		Style("color", IfElseString(item.Active, "#0066cc", "#333")).
		Style("white-space", "nowrap").
		Style("overflow", "hidden").
		Style("text-overflow", "ellipsis").
		Text(item.Label)
}

func (sn *SideNav) renderItemBadge(item SideNavItem) app.UI {
	if item.Badge == "" {
		return app.Div()
	}

	return app.Span().
		Class("prism-sidenav__item-badge").
		Style("background-color", "#0066cc").
		Style("color", "#fff").
		Style("font-size", "10px").
		Style("font-weight", "600").
		Style("padding", "2px 6px").
		Style("border-radius", "10px").
		Style("min-width", "16px").
		Style("text-align", "center").
		Text(item.Badge)
}

func (sn *SideNav) renderFooter() app.UI {
	if !sn.ShowFooter || len(sn.FooterItems) == 0 {
		return app.Div()
	}

	footerItems := []app.UI{}
	for _, item := range sn.FooterItems {
		footerItems = append(footerItems, sn.renderNavItem(item, 0))
	}

	return app.Div().
		Class("prism-sidenav__footer").
		Style("border-top", "1px solid #e1e1e1").
		Style("padding", "8px 0").
		Body(footerItems...)
}

func (sn *SideNav) getClasses() string {
	classes := "prism-sidenav"

	if sn.IsCollapsed {
		classes += " prism-sidenav--collapsed"
	}

	if !sn.IsOpen {
		classes += " prism-sidenav--closed"
	}

	if sn.Variant != "fixed" {
		classes += " prism-sidenav--" + sn.Variant
	}

	if sn.Position != "left" {
		classes += " prism-sidenav--" + sn.Position
	}

	return classes
}

func (sn *SideNav) getItemClasses(item SideNavItem) string {
	classes := "prism-sidenav__item"

	if item.Active {
		classes += " prism-sidenav__item--active"
	}

	if item.Disabled {
		classes += " prism-sidenav__item--disabled"
	}

	return classes
}

func (sn *SideNav) getShadow() string {
	if sn.Variant == "overlay" {
		return "2px 0 8px rgba(0, 0, 0, 0.15)"
	}
	return "none"
}

func (sn *SideNav) getTransform() string {
	if !sn.IsOpen && sn.Variant == "overlay" {
		if sn.Position == "left" {
			return "translateX(-100%)"
		}
		return "translateX(100%)"
	}
	return "translateX(0)"
}

// Helper function to create a navigation item
func NewSideNavItem(id, label string) SideNavItem {
	return SideNavItem{
		ID:       id,
		Label:    label,
		Active:   false,
		Disabled: false,
		Children: make([]SideNavItem, 0),
	}
}

func (sni SideNavItem) WithIcon(icon Icon) SideNavItem {
	sni.Icon = icon
	return sni
}

func (sni SideNavItem) WithHref(href string) SideNavItem {
	sni.Href = href
	return sni
}

func (sni SideNavItem) WithActive(active bool) SideNavItem {
	sni.Active = active
	return sni
}

func (sni SideNavItem) WithDisabled(disabled bool) SideNavItem {
	sni.Disabled = disabled
	return sni
}

func (sni SideNavItem) WithBadge(badge string) SideNavItem {
	sni.Badge = badge
	return sni
}

func (sni SideNavItem) WithChildren(children ...SideNavItem) SideNavItem {
	sni.Children = children
	return sni
}

func (sni SideNavItem) WithOnClick(fn func(app.Context)) SideNavItem {
	sni.OnClick = fn
	return sni
}

// Helper function to create a navigation group
func NewSideNavGroup(id, label string) SideNavGroup {
	return SideNavGroup{
		ID:       id,
		Label:    label,
		Expanded: true,
		Items:    make([]SideNavItem, 0),
	}
}

func (sng SideNavGroup) WithItems(items ...SideNavItem) SideNavGroup {
	sng.Items = items
	return sng
}

func (sng SideNavGroup) WithExpanded(expanded bool) SideNavGroup {
	sng.Expanded = expanded
	return sng
}

// Helper functions for common navigation patterns

// MainSideNav creates a typical main navigation
func MainSideNav() *SideNav {
	return NewSideNav().
		WithTitle("Your App").
		WithItems(
			NewSideNavItem("dashboard", "Dashboard").
				WithIcon("🏠").
				WithHref("/").
				WithActive(true),
			NewSideNavItem("users", "Users").
				WithIcon("👥").
				WithHref("/users"),
			NewSideNavItem("settings", "Settings").
				WithIcon("⚙️").
				WithHref("/settings"),
		)
}

// AdminSideNav creates an admin navigation with groups
func AdminSideNav() *SideNav {
	return NewSideNav().
		WithTitle("Admin Panel").
		WithGroups(
			NewSideNavGroup("content", "Content").
				WithItems(
					NewSideNavItem("posts", "Posts").WithIcon("📝"),
					NewSideNavItem("pages", "Pages").WithIcon("📄"),
					NewSideNavItem("media", "Media").WithIcon("🖼️"),
				),
			NewSideNavGroup("users", "Users").
				WithItems(
					NewSideNavItem("all-users", "All Users").WithIcon("👥"),
					NewSideNavItem("roles", "Roles").WithIcon("🔑"),
					NewSideNavItem("permissions", "Permissions").WithIcon("🛡️"),
				),
			NewSideNavGroup("system", "System").
				WithItems(
					NewSideNavItem("settings", "Settings").WithIcon("⚙️"),
					NewSideNavItem("logs", "Logs").WithIcon("📋"),
					NewSideNavItem("backup", "Backup").WithIcon("💾"),
				),
		)
}

// MobileSideNav creates a mobile-friendly overlay navigation
func MobileSideNav() *SideNav {
	return NewSideNav().
		WithVariant("overlay").
		WithOpen(false).
		WithWidth("280px")
}