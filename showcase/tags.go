package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// tagsPage represents the Tags component showcase
type tagsPage struct {
	app.Compo
}

// newTagsPage creates a new Tags component showcase
func newTagsPage() *tagsPage {
	return &tagsPage{}
}

// OnNav initializes the page when navigated to
func (p *tagsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *tagsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// Render renders the Tags component showcase
func (p *tagsPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Tags Component"),
		app.P().Text("Tags represent categories to which content belongs. They can be used for keywords, labels, or people, and are often grouped to describe an item or search request."),

		// Basic Tags
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Tags"),
				app.P().Text("Simple group of basic tags with default settings:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().Text("Default Tag"),
								sp.Tag().Text("Another Tag"),
								sp.Tag().Text("Third Tag"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().Text("Default Tag"),
	sp.Tag().Text("Another Tag"),
	sp.Tag().Text("Third Tag"),
)`),
			),

		// Tag Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tag Sizes"),
				app.P().Text("Tags come in three sizes: small (S), medium (M, default), and large (L):"),
				app.Div().
					Class("component-example").
					Body(
						app.H3().Text("Small (S)"),
						sp.Tags().
							Body(
								sp.Tag().SizeS().Text("Small Tag"),
								sp.Tag().SizeS().SetInvalid().Text("Invalid Tag"),
								sp.Tag().SizeS().SetDisabled().Text("Disabled Tag"),
							),
						app.H3().Text("Medium (M)"),
						sp.Tags().
							Body(
								sp.Tag().SizeM().Text("Medium Tag"),
								sp.Tag().SizeM().SetInvalid().Text("Invalid Tag"),
								sp.Tag().SizeM().SetDisabled().Text("Disabled Tag"),
							),
						app.H3().Text("Large (L)"),
						sp.Tags().
							Body(
								sp.Tag().SizeL().Text("Large Tag"),
								sp.Tag().SizeL().SetInvalid().Text("Invalid Tag"),
								sp.Tag().SizeL().SetDisabled().Text("Disabled Tag"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`// Small tags
sp.Tags().
	Body(
		sp.Tag().SizeS().Text("Small Tag"),
		sp.Tag().SizeS().SetInvalid().Text("Invalid Tag"),
		sp.Tag().SizeS().SetDisabled().Text("Disabled Tag"),
	)

// Medium tags
sp.Tags().
	Body(
		sp.Tag().SizeM().Text("Medium Tag"),
		sp.Tag().SizeM().SetInvalid().Text("Invalid Tag"),
		sp.Tag().SizeM().SetDisabled().Text("Disabled Tag"),
	)

// Large tags
sp.Tags().
	Body(
		sp.Tag().SizeL().Text("Large Tag"),
		sp.Tag().SizeL().SetInvalid().Text("Invalid Tag"),
		sp.Tag().SizeL().SetDisabled().Text("Disabled Tag"),
	)`),
			),

		// Deletable Tags
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Deletable Tags"),
				app.P().Text("Tags can be configured to be deletable. When clicked, they dispatch a 'delete' event:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().SetDeletable().Text("Deletable Tag"),
								sp.Tag().SetDeletable().SetInvalid().Text("Invalid Deletable"),
								sp.Tag().SetDeletable().SetDisabled().Text("Disabled Deletable"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().SetDeletable().Text("Deletable Tag"),
	sp.Tag().SetDeletable().SetInvalid().Text("Invalid Deletable"),
	sp.Tag().SetDeletable().SetDisabled().Text("Disabled Deletable"),
)`),
			),

		// Tags with Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tags with Icons"),
				app.P().Text("Tags can include icons to provide visual context:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().Text("User").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("user"),
											),
									),
								sp.Tag().Text("Calendar").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("calendar"),
											),
									),
								sp.Tag().Text("Tag").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("tag"),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().Text("User").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("user"),
				),
		),
	sp.Tag().Text("Calendar").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("calendar"),
				),
		),
	sp.Tag().Text("Tag").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("tag"),
				),
		),
)`),
			),

		// Tags with Avatars
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tags with Avatars"),
				app.P().Text("Tags can include avatars, which is useful for representing people:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().Text("John Doe").
									Avatar(
										app.Elem("div").
											Attr("slot", "avatar").
											Body(
												sp.Avatar().
													Label("John Doe"),
											),
									),
								sp.Tag().Text("Jane Smith").
									Avatar(
										app.Elem("div").
											Attr("slot", "avatar").
											Body(
												sp.Avatar().
													Label("Jane Smith"),
											),
									),
								sp.Tag().SetDeletable().Text("Alex Johnson").
									Avatar(
										app.Elem("div").
											Attr("slot", "avatar").
											Body(
												sp.Avatar().
													Label("Alex Johnson"),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().Text("John Doe").
		Avatar(
			app.Elem("div").
				Attr("slot", "avatar").
				Body(
					sp.Avatar().
						Label("John Doe"),
				),
		),
	sp.Tag().Text("Jane Smith").
		Avatar(
			app.Elem("div").
				Attr("slot", "avatar").
				Body(
					sp.Avatar().
						Label("Jane Smith"),
				),
		),
	sp.Tag().SetDeletable().Text("Alex Johnson").
		Avatar(
			app.Elem("div").
				Attr("slot", "avatar").
				Body(
					sp.Avatar().
						Label("Alex Johnson"),
				),
		),
)`),
			),

		// Tag States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tag States"),
				app.P().Text("Tags can be displayed in various states, including invalid, disabled, and readonly:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().Text("Default Tag"),
								sp.Tag().SetInvalid().Text("Invalid Tag"),
								sp.Tag().SetDisabled().Text("Disabled Tag"),
								sp.Tag().SetReadonly().Text("Readonly Tag"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().Text("Default Tag"),
	sp.Tag().SetInvalid().Text("Invalid Tag"),
	sp.Tag().SetDisabled().Text("Disabled Tag"),
	sp.Tag().SetReadonly().Text("Readonly Tag"),
)`),
			),

		// Complex Tag Combinations
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Complex Tag Combinations"),
				app.P().Text("Combined features with deletable tags, icons, and various states:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Tags().
							Body(
								sp.Tag().SetDeletable().Text("Deletable").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("tag"),
											),
									),
								sp.Tag().SetDeletable().SetInvalid().Text("Invalid").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("alert"),
											),
									),
								sp.Tag().SetDeletable().SetDisabled().Text("Disabled").
									Icon(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("lock"),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Tags().
Body(
	sp.Tag().SetDeletable().Text("Deletable").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("tag"),
				),
		),
	sp.Tag().SetDeletable().SetInvalid().Text("Invalid").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("alert"),
				),
		),
	sp.Tag().SetDeletable().SetDisabled().Text("Disabled").
		Icon(
			app.Elem("div").
				Attr("slot", "icon").
				Body(
					sp.Icon().Name("lock"),
				),
		),
)`),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.H3().Text("When to use"),
				app.Ul().
					Body(
						app.Li().Text("To categorize or label content"),
						app.Li().Text("To represent keywords in a search interface"),
						app.Li().Text("To show selected filter options"),
						app.Li().Text("To represent people in user interfaces"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Keep tag text concise and descriptive"),
						app.Li().Text("Use colors and icons consistently"),
						app.Li().Text("Ensure that tag states (invalid, disabled) are clear and meaningful"),
						app.Li().Text("For interactive tags, make sure that click and delete interactions are intuitive"),
					),
			),
	)
}
