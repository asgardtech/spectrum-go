package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type SectionVisibility int

const (
	SectionVisibilityDefault SectionVisibility = iota
	SectionVisibilityVisible
	SectionVisibilityHidden
)

type Page struct {
	app.Compo

	currentTheme sp.ThemeColor

	Name               string
	Description        string
	TopNavVisibility   SectionVisibility
	SidenavVisibility  SectionVisibility
	UserMenuVisibility SectionVisibility
	Icon               string
	UrlPath            string
	SidenavGroup       string
	TopNavGroup        string
	UserMenuGroup      string
	Form               Form[any]

	// Content
	content []app.UI
	app     *App

	// Context
	ctx app.Context
}

type PageConstructor func() IPage

func NewPage(options PageOptions) *Page {
	return &Page{
		Name:               options.Name,
		Description:        options.Description,
		TopNavVisibility:   options.TopNavVisibility,
		SidenavVisibility:  options.SidenavVisibility,
		UserMenuVisibility: options.UserMenuVisibility,
		Icon:               options.Icon,
		UrlPath:            options.UrlPath,
		SidenavGroup:       options.SidenavGroup,
		TopNavGroup:        options.TopNavGroup,
		UserMenuGroup:      options.UserMenuGroup,
		Form:               options.Form,
	}
}

func PageConstructorWithDefaults(options PageOptions) PageConstructor {
	return func() IPage {
		page := NewPage(options)
		return page
	}
}

type PageOptions struct {
	Name               string
	Description        string
	TopNavVisibility   SectionVisibility
	SidenavVisibility  SectionVisibility
	UserMenuVisibility SectionVisibility
	Icon               string
	UrlPath            string
	SidenavGroup       string
	TopNavGroup        string
	UserMenuGroup      string
	Form               Form[any]
}

type IPage interface {
	app.Composer

	GetApp() *App
	GetForm() Form[any]
	GetName() string
	GetDescription() string
	GetIcon() string
	GetUrlPath() string
	GetSidenavLinkVisibility() SectionVisibility
	GetTopNavLinkVisibility() SectionVisibility
	GetUserMenuLinkVisibility() SectionVisibility
	GetSidenavGroup() string
	GetTopNavGroup() string
	GetUserMenuGroup() string

	setApp(app *App)
}

func (p *Page) setApp(app *App) {
	p.app = app
}

func (p *Page) GetApp() *App {
	return p.app
}

func (p *Page) GetForm() Form[any] {
	return p.Form
}

func (p *Page) GetName() string {
	return p.Name
}

func (p *Page) GetDescription() string {
	return p.Description
}

func (p *Page) GetIcon() string {
	return p.Icon
}

func (p *Page) GetUrlPath() string {
	return p.UrlPath
}

func (p *Page) GetSidenavLinkVisibility() SectionVisibility {
	return p.SidenavVisibility
}

func (p *Page) GetTopNavLinkVisibility() SectionVisibility {
	return p.TopNavVisibility
}

func (p *Page) GetUserMenuLinkVisibility() SectionVisibility {
	return p.UserMenuVisibility
}

func (p *Page) GetSidenavGroup() string {
	return p.SidenavGroup
}

func (p *Page) GetTopNavGroup() string {
	return p.TopNavGroup
}

func (p *Page) GetUserMenuGroup() string {
	return p.UserMenuGroup
}

func (p *Page) WithSidenavGroup(group string) *Page {
	p.SidenavGroup = group

	return p
}

func (p *Page) WithTopNavGroup(group string) *Page {
	p.TopNavGroup = group
	return p
}

func (p *Page) WithUserMenu(visibility SectionVisibility) *Page {
	p.UserMenuVisibility = visibility
	return p
}

func (p *Page) WithUserMenuGroup(group string) *Page {
	p.UserMenuGroup = group
	return p
}

func (p *Page) WithName(name string) *Page {
	p.Name = name
	return p
}

func (p *Page) WithIcon(icon string) *Page {
	p.Icon = icon
	return p
}

func (p *Page) WithSidenavLinkVisibility(visibility SectionVisibility) *Page {
	p.SidenavVisibility = visibility
	return p
}

func (p *Page) WithTopNavLinkVisibility(visibility SectionVisibility) *Page {
	p.TopNavVisibility = visibility
	return p
}

func (p *Page) WithUserMenuLinkVisibility(visibility SectionVisibility) *Page {
	p.UserMenuVisibility = visibility
	return p
}

func (p *Page) WithUrlPath(urlPath string) *Page {
	p.UrlPath = urlPath
	return p
}

func (p *Page) WithForm(form Form[any]) *Page {
	p.Form = form
	return p
}

func (p *Page) WithDescription(description string) *Page {
	p.Description = description
	return p
}

func (p *Page) WithTopNav(visibility SectionVisibility) *Page {
	p.TopNavVisibility = visibility
	return p
}

func (p *Page) WithSidenav(visibility SectionVisibility) *Page {
	p.SidenavVisibility = visibility
	return p
}

func (p *Page) Content(v ...app.UI) *Page {
	p.content = app.FilterUIElems(v...)
	return p
}

func (p *Page) OnMount(ctx app.Context) {
	p.app.setContext(ctx)

	p.ctx = ctx
	ctx.ObserveState("current-theme", &p.currentTheme)
}

func (p *Page) OnNav(ctx app.Context) {
	p.app.setContext(ctx)

	ctx.Page().SetTitle(p.Name)
	ctx.Page().SetDescription(p.Description)
}

func (p *Page) Render() app.UI {
	// Create a main container with theme-aware styling
	return app.Div().
		Class("spectrum-Typography").
		Body(
			sp.Theme().
				Dir(sp.ThemeDirLtr).
				Color(p.currentTheme).
				Scale(sp.ThemeScaleLarge).
				// Style("background-color", "var(--spectrum-gray-100)").
				Body(
					p.renderTopNav(),
					app.If(
						p.SidenavVisibility == SectionVisibilityVisible ||
							(p.SidenavVisibility == SectionVisibilityDefault && p.app.user.LoggedIn),
						func() app.UI {
							return sp.SplitView().
								Collapsible(true).
								PrimarySize("250px").
								PrimaryMin(200).
								Resizable(true).
								AddToBody(p.renderSidenav()).
								AddToBody(p.renderContent())
						}).Else(
						func() app.UI {
							return p.renderContent()
						},
					),
				),
		)
}

// Toggle theme between light and dark
func (p *Page) toggleTheme() {
	currentTheme := p.currentTheme
	if currentTheme == "" {
		currentTheme = sp.ThemeColorLight
	}

	if currentTheme == sp.ThemeColorLight {
		currentTheme = sp.ThemeColorDark
	} else {
		currentTheme = sp.ThemeColorLight
	}

	p.ctx.SetState("current-theme", currentTheme).
		Persist().
		Broadcast()
}

func (p *Page) renderTopNav() app.UI {
	return app.If(
		p.TopNavVisibility == SectionVisibilityVisible ||
			(p.TopNavVisibility == SectionVisibilityDefault && p.app.user.LoggedIn),
		func() app.UI {
			return sp.TopNav().
				Body(
					sp.TopNavItem().
						Label("Home").
						Text("Home"),
					sp.ActionMenu().
						Placement("bottom-end").
						Style("margin-inline-start", "auto").
						Quiet(true).
						AddToBody(
							sp.MenuItem().
								Label("Toggle Theme").
								Text("Toggle Theme").
								//	Icon(p.getThemeIcon()).
								OnClick(func(ctx app.Context, e app.Event) {
									app.Log("Current Theme", p.currentTheme)
									p.toggleTheme()
								}),
						),
				)
		}).Else(
		func() app.UI {
			return app.Div()
		},
	)
}

func (p *Page) renderSidenav() app.UI {
	linksByGroup := make(map[string][]IPage)

	for _, page := range p.app.pageConstructors {
		page := page()
		if page.GetSidenavLinkVisibility() == SectionVisibilityHidden {
			continue
		}

		linksByGroup[page.GetSidenavGroup()] = append(linksByGroup[page.GetSidenavGroup()], page)
	}

	groupsWithLinks := []struct {
		Group string
		Links []IPage
	}{}

	for group, links := range linksByGroup {
		groupsWithLinks = append(groupsWithLinks, struct {
			Group string
			Links []IPage
		}{
			Group: group,
			Links: links,
		})
	}

	return app.If(
		p.SidenavVisibility == SectionVisibilityVisible ||
			(p.SidenavVisibility == SectionVisibilityDefault && p.app.user.LoggedIn),
		func() app.UI {
			return app.Div().
				Class("spectrum-sidenav").
				Body(
					app.Range(groupsWithLinks).Slice(func(i int) app.UI {
						group := groupsWithLinks[i]

						return sp.SidenavItem().
							Expanded(true).
							Label(group.Group).
							Body(
								app.Range(group.Links).Slice(func(i int) app.UI {
									page := group.Links[i]

									return sp.SidenavItem().
										Label(page.GetName()).
										Value(page.GetUrlPath()).
										Href(page.GetUrlPath())
								}),
							)
					}),
				)
		}).Else(
		func() app.UI {
			return app.Div()
		},
	)
}

func (p *Page) renderContent() app.UI {
	return app.Div().
		Class("spectrum-Body").
		Body(
			p.content...,
		)
}
