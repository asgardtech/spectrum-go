package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type IPage interface {
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
	GetForm() Form[any]

	WithName(name string) *Page
	WithDescription(description string) *Page
	WithIcon(icon string) *Page
	WithUrlPath(urlPath string) *Page
	WithSidenav(visibility SectionVisibility) *Page
	WithTopNav(visibility SectionVisibility) *Page
	WithUser(user User) *Page
	WithUserMenu(visibility SectionVisibility) *Page
	WithSidenavGroup(group string) *Page
}

type SectionVisibility string

const (
	SectionVisibilityHidden  SectionVisibility = "hidden"
	SectionVisibilityVisible SectionVisibility = "visible"
	SectionVisibilityDefault SectionVisibility = "default"
)

type Page struct {
	app.Compo

	currentTheme sp.ThemeColor

	Name               string
	Description        string
	TopNavVisibility   SectionVisibility
	SidenavVisibility  SectionVisibility
	UserMenuVisibility SectionVisibility
	User               User
	Icon               string
	UrlPath            string
	SidenavGroup       string
	TopNavGroup        string
	UserMenuGroup      string
	Form               Form[any]

	// Content
	content []app.UI

	// Context
	ctx app.Context
}

var _ IPage = (*Page)(nil)

func NewPage() *Page {
	page := &Page{}

	return page.
		WithTopNav(SectionVisibilityDefault).
		WithSidenav(SectionVisibilityDefault)
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

func (p Page) WithForm(form Form[any]) Page {
	p.Form = form
	return p
}

func (l *Page) WithDescription(description string) *Page {
	l.Description = description
	return l
}

func (l *Page) WithTopNav(visibility SectionVisibility) *Page {
	l.TopNavVisibility = visibility
	return l
}

func (l *Page) WithSidenav(visibility SectionVisibility) *Page {
	l.SidenavVisibility = visibility
	return l
}

func (l *Page) WithUser(user User) *Page {
	l.User = user
	return l
}

func (p *Page) Content(v ...app.UI) *Page {
	p.content = app.FilterUIElems(v...)
	return p
}

func (p *Page) OnMount(ctx app.Context) {
	p.ctx = ctx
	ctx.ObserveState("current-theme", &p.currentTheme)
}

func (p *Page) OnNav(ctx app.Context) {
	ctx.Page().SetTitle(p.Name)
	ctx.Page().SetDescription(p.Description)
}

func (p *Page) Render() app.UI {
	// Create a main container with theme-aware styling
	return sp.Theme().
		Color(p.currentTheme).
		Scale(sp.ThemeScaleMedium).
		Body(
			p.renderTopNav(),
			app.If(
				p.SidenavVisibility == SectionVisibilityVisible ||
					(p.SidenavVisibility == SectionVisibilityDefault && p.User.LoggedIn),
				func() app.UI {
					return sp.SplitView().
						Collapsible(true).
						PrimarySize("250px").
						PrimaryMin(200).
						Resizable(false).
						AddToBody(p.renderSidenav()).
						AddToBody(p.renderContent())
				}).Else(
				func() app.UI {
					return p.renderContent()
				},
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
			(p.TopNavVisibility == SectionVisibilityDefault && p.User.LoggedIn),
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
	return app.If(
		p.SidenavVisibility == SectionVisibilityVisible ||
			(p.SidenavVisibility == SectionVisibilityDefault && p.User.LoggedIn),
		func() app.UI {
			return app.Div().
				Class("spectrum-sidenav").
				Body(
					app.Div().
						Class("spectrum-sidenav-heading").
						Text("Components"),
					app.Div().
						Class("spectrum-sidenav-body").
						Body(
							sp.Sidenav().
								Label("Component Navigation").
								AddToBody(sp.SidenavItem().
									Label("Home").
									Value("/").
									Href("/"),
								).
								// Implemented components (mark as active)
								AddToBody(sp.SidenavItem().
									Label("Button").
									Value("button").
									Href("/button"),
								).
								AddToBody(sp.SidenavItem().
									Label("Checkbox").
									Value("checkbox").
									Href("/checkbox"),
								).
								AddToBody(sp.SidenavItem().
									Label("Badge").
									Value("badge").
									Href("/badge"),
								).
								AddToBody(sp.SidenavItem().
									Label("Radio").
									Value("radio").
									Href("/radio"),
								).
								AddToBody(sp.SidenavItem().
									Label("Switch").
									Value("switch").
									Href("/switch"),
								).
								// New components added
								AddToBody(sp.SidenavItem().
									Label("Progress Bar").
									Value("progress-bar").
									Href("/progress-bar"),
								).
								AddToBody(sp.SidenavItem().
									Label("Progress Circle").
									Value("progress-circle").
									Href("/progress-circle"),
								).
								AddToBody(sp.SidenavItem().
									Label("Link").
									Value("link").
									Href("/link"),
								).
								AddToBody(sp.SidenavItem().
									Label("Icon").
									Value("icon").
									Href("/icon"),
								).
								AddToBody(sp.SidenavItem().
									Label("Avatar").
									Value("avatar").
									Href("/avatar"),
								).
								AddToBody(sp.SidenavItem().
									Label("Status Light").
									Value("status-light").
									Href("/status-light"),
								).
								AddToBody(sp.SidenavItem().
									Label("Accordion").
									Value("accordion").
									Href("/accordion"),
								).
								AddToBody(sp.SidenavItem().
									Label("Coachmark").
									Value("coachmark").
									Href("/coachmark"),
								).
								AddToBody(sp.SidenavItem().
									Label("Alert Banner").
									Value("alert-banner").
									Href("/alert-banner"),
								).
								AddToBody(sp.SidenavItem().
									Label("Alert Dialog").
									Value("alert-dialog").
									Href("/alert-dialog"),
								).
								AddToBody(sp.SidenavItem().
									Label("Asset").
									Value("asset").
									Href("/asset"),
								).
								AddToBody(sp.SidenavItem().
									Label("Button Group").
									Value("button-group").
									Href("/button-group"),
								).
								AddToBody(sp.SidenavItem().
									Label("Card").
									Value("card").
									Href("/card"),
								).
								AddToBody(sp.SidenavItem().
									Label("Toast").
									Value("toast").
									Href("/toast"),
								).
								AddToBody(sp.SidenavItem().
									Label("Tooltip").
									Value("tooltip").
									Href("/tooltip"),
								).
								AddToBody(sp.SidenavItem().
									Label("Breadcrumbs").
									Value("breadcrumbs").
									Href("/breadcrumbs"),
								).
								AddToBody(sp.SidenavItem().
									Label("Tabs").
									Value("tabs").
									Href("/tabs"),
								).
								AddToBody(sp.SidenavItem().
									Label("Tags").
									Value("tags").
									Href("/tags"),
								).
								AddToBody(sp.SidenavItem().
									Label("Divider").
									Value("divider").
									Href("/divider"),
								).
								AddToBody(sp.SidenavItem().
									Label("Textfield").
									Value("textfield").
									Href("/textfield"),
								).
								AddToBody(sp.SidenavItem().
									Label("Textarea").
									Value("textarea").
									Href("/textarea"),
								).
								AddToBody(sp.SidenavItem().
									Label("Popover").
									Value("popover").
									Href("/popover"),
								).
								AddToBody(sp.SidenavItem().
									Label("Menu").
									Value("menu").
									Href("/menu"),
								).
								AddToBody(sp.SidenavItem().
									Label("Meter").
									Value("meter").
									Href("/meter"),
								).
								AddToBody(sp.SidenavItem().
									Label("Search").
									Value("search").
									Href("/search"),
								).
								AddToBody(sp.SidenavItem().
									Label("Field Label").
									Value("field-label").
									Href("/field-label"),
								).
								AddToBody(sp.SidenavItem().
									Label("Help Text").
									Value("help-text").
									Href("/help-text"),
								).
								AddToBody(sp.SidenavItem().
									Label("Slider").
									Value("slider").
									Href("/slider"),
								).
								AddToBody(sp.SidenavItem().
									Label("Sidenav").
									Value("sidenav").
									Href("/sidenav"),
								).
								AddToBody(sp.SidenavItem().
									Label("Number Field").
									Value("number-field").
									Href("/number-field"),
								).
								AddToBody(sp.SidenavItem().
									Label("Combobox").
									Value("combobox").
									Href("/combobox"),
								).
								AddToBody(sp.SidenavItem().
									Label("Split View").
									Value("split-view").
									Href("/split-view"),
								).
								AddToBody(sp.SidenavItem().
									Label("Dialog").
									Value("dialog").
									Href("/dialog"),
								).
								AddToBody(sp.SidenavItem().
									Label("Contextual Help").
									Value("contextual-help").
									Href("/contextual-help"),
								).
								AddToBody(sp.SidenavItem().
									Label("Field Group").
									Value("field-group").
									Href("/field-group"),
								).
								AddToBody(sp.SidenavItem().
									Label("Illustrated Message").
									Value("illustrated-message").
									Href("/illustrated-message"),
								).
								AddToBody(sp.SidenavItem().
									Label("Thumbnail").
									Value("thumbnail").
									Href("/thumbnail"),
								).
								AddToBody(sp.SidenavItem().
									Label("Picker").
									Value("picker").
									Href("/picker"),
								).
								AddToBody(sp.SidenavItem().
									Label("Picker Button").
									Value("picker-button").
									Href("/picker-button"),
								).
								AddToBody(sp.SidenavItem().
									Label("Swatch").
									Value("swatch").
									Href("/swatch"),
								).
								AddToBody(sp.SidenavItem().
									Label("Swatch Group").
									Value("swatch-group").
									Href("/swatch-group"),
								).
								AddToBody(sp.SidenavItem().
									Label("Tray").
									Value("tray").
									Href("/tray"),
								).
								AddToBody(sp.SidenavItem().
									Label("Overlay").
									Value("overlay").
									Href("/overlay"),
								).
								AddToBody(sp.SidenavItem().
									Label("Overlay Trigger").
									Value("overlay-trigger").
									Href("/overlay-trigger"),
								).
								// Not yet implemented components
								AddToBody(sp.SidenavItem().Label("Accordion Item").Value("accordion-item").Disabled(true)).
								AddToBody(sp.SidenavItem().
									Label("Action Bar").
									Value("action-bar").
									Href("/action-bar"),
								).
								AddToBody(sp.SidenavItem().
									Label("Action Button").
									Value("action-button").
									Href("/action-button"),
								).
								AddToBody(sp.SidenavItem().Label("Action Group").Value("action-group").Disabled(true)).
								AddToBody(sp.SidenavItem().
									Label("Action Menu").
									Value("action-menu").
									Href("/action-menu"),
								).
								AddToBody(sp.SidenavItem().
									Label("Top Nav").
									Value("top-nav").
									Href("/top-nav"),
								).
								AddToBody(sp.SidenavItem().Label("Action Group").Value("action-group").Href("/action-group")).
								AddToBody(sp.SidenavItem().Label("Action Menu").Value("action-menu").Href("/action-menu")).
								AddToBody(sp.SidenavItem().Label("Color").Value("color").Href("/color")).
								AddToBody(sp.SidenavItem().Label("Dropzone").Value("dropzone").Href("/dropzone")).
								AddToBody(sp.SidenavItem().Label("Infield Button").Value("infield-button").Href("/infield-button")).
								AddToBody(sp.SidenavItem().Label("Opacity Checkerboard").Value("opacity-checkerboard").Href("/opacity-checkerboard")).
								AddToBody(sp.SidenavItem().Label("Table").Value("table").Href("/table")).
								AddToBody(sp.SidenavItem().Label("Theme").Value("theme").Href("/theme")).
								AddToBody(sp.SidenavItem().Label("Tray").Value("tray").Href("/tray")).
								AddToBody(sp.SidenavItem().Label("Truncated").Value("truncated").Href("/truncated")).
								AddToBody(sp.SidenavItem().Label("Underlay").Value("underlay").Href("/underlay")),
						),
				)
		}).Else(
		func() app.UI {
			return app.Div()
		},
	)
}

func (p *Page) renderContent() app.UI {
	// Content area - will be replaced with component showcase
	return app.Div().
		Body(
			p.content...,
		)
}
