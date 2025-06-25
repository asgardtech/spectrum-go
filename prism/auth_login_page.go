package prism

import (
	"log"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/logs"
)

type LoginPage struct {
	app.Compo
	Page

	Providers                []string
	WelcomeMessage           string
	LoginButtonText          string
	RegisterButtonText       string
	ForgotPasswordButtonText string
	Logo                     string

	Email    string
	Password string
}

type LoginPageOptions struct {
	PageOptions

	Providers                []string
	WelcomeMessage           string
	LoginButtonText          string
	RegisterButtonText       string
	ForgotPasswordButtonText string
	Logo                     string
}

func NewLoginPage(options LoginPageOptions) *LoginPage {
	return &LoginPage{
		Page:                     *NewPage(options.PageOptions),
		Providers:                options.Providers,
		WelcomeMessage:           options.WelcomeMessage,
		LoginButtonText:          options.LoginButtonText,
		RegisterButtonText:       options.RegisterButtonText,
		ForgotPasswordButtonText: options.ForgotPasswordButtonText,
		Logo:                     options.Logo,
	}
}

func LoginPageConstructor(options LoginPageOptions) PageConstructor {
	return func() IPage {
		return NewLoginPage(options)
	}
}

func (l *LoginPage) WithProviders(providers ...string) *LoginPage {
	l.Providers = providers
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

func (l *LoginPage) WithForgotPasswordButtonText(forgotPasswordButtonText string) *LoginPage {
	l.ForgotPasswordButtonText = forgotPasswordButtonText
	return l
}

func (p *LoginPage) Render() app.UI {
	log.Printf("Rendering login page")
	return p.
		Content(
			app.Text("Hello World"),
			app.Div().
				Body(
					app.Div().
						Body(
							app.H1().Text(p.WelcomeMessage),

							app.Range(p.Providers).Slice(func(i int) app.UI {
								return sp.Button().
									Text(p.Providers[i]).
									OnClick(func(ctx app.Context, e app.Event) {
										app.Log(logs.New("clicked provider").WithTag("provider", p.Providers[i]))
									})
							}),

							app.P().Text(p.RegisterButtonText),
							app.P().Text(p.ForgotPasswordButtonText),
							app.Img().Src(p.Logo),
						),
				),
		)
}
