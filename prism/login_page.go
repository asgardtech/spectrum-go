package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type LoginPage struct {
	app.Compo

	Providers                []string
	UsePasswordLogin         bool
	WelcomeMessage           string
	LoginButtonText          string
	RegisterButtonText       string
	ForgotPasswordButtonText string
	Logo                     string

	Email    string
	Password string
}

func (l *LoginPage) WithProviders(providers ...string) *LoginPage {
	l.Providers = providers
	return l
}

func (l *LoginPage) WithPasswordLogin(usePasswordLogin bool) *LoginPage {
	l.UsePasswordLogin = usePasswordLogin
	return l
}

func (l *LoginPage) WithWelcomeMessage(welcomeMessage string) *LoginPage {
	l.WelcomeMessage = welcomeMessage
	return l
}

func (l *LoginPage) WithLoginButtonText(loginButtonText string) *LoginPage {
	l.LoginButtonText = loginButtonText
	return l
}

func (l *LoginPage) WithRegisterButtonText(registerButtonText string) *LoginPage {
	l.RegisterButtonText = registerButtonText
	return l
}

func (l *LoginPage) WithLogo(logo string) *LoginPage {
	l.Logo = logo
	return l
}

func (l *LoginPage) WithForgotPasswordButtonText(forgotPasswordButtonText string) *LoginPage {
	l.ForgotPasswordButtonText = forgotPasswordButtonText
	return l
}

func NewLoginPage() *LoginPage {
	page := &LoginPage{
		Providers:        []string{},
		UsePasswordLogin: true,
	}

	return page
}

func (p *LoginPage) Render() app.UI {
	return NewLayout().
		WithTitle("Login").
		WithDescription("Login to your account").
		WithTopNav(SectionVisibilityHidden).
		WithSidenav(SectionVisibilityHidden).
		Content(
			app.Div().
				Style("display", "flex").
				Style("align-items", "center").
				Style("justify-content", "center").
				Style("min-height", "100vh").
				Body(
					app.Div().
						Style("width", "320px").
						Style("display", "flex").
						Style("flex-direction", "column").
						Style("gap", "16px").
						Body(
							app.H1().Text(p.WelcomeMessage),
							sp.FieldLabel().Text("Email"),
							sp.Textfield().Value(p.Email),
							sp.FieldLabel().Text("Password"),
							sp.Textfield().Value(p.Password),

							app.Range(p.Providers).Slice(func(i int) app.UI {
								return app.P().Text(p.Providers[i])
							}),

							app.P().Text(p.RegisterButtonText),
							app.P().Text(p.ForgotPasswordButtonText),
							app.Img().Src(p.Logo),
						),
				),
		)
}
