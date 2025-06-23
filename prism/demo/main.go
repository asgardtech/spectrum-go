package main

import (
	"context"
	"log"

	"github.com/asgardtech/spectrum-go/prism"
)

func main() {

	app := prism.NewApp().
		WithBaseUrl("http://localhost:9090").
		WithDescription("Prism Demo").
		WithTitle("Prism Demo").
		WithPage(
			prism.NewLoginPage().
				WithUrlPath("/login").
				WithName("Login").
				WithDescription("Login to your account").
				WithTopNav(prism.SectionVisibilityHidden).
				WithSidenav(prism.SectionVisibilityHidden),
		)

	err := app.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Configure routes
	// 	app.Route("/", func() app.Composer { return &prism.HomePage{} })
	// app.Route("/login", func() app.Composer { return prism.NewLoginPage() })
	// app.Route("/demo", func() app.Composer { return prism.NewDemoPage() })
	// app.Route("/test", func() app.Composer { return prism.NewComprehensiveTestPage() })

	// // Setup the app
	// app.RunWhenOnBrowser()

	// // Start server when on server
	// http.Handle("/", &app.Handler{
	// 	Name:         "Prism Demo",
	// 	Description:  "A demo of the Prism Go+WASM UI framework",
	// 	Author:       "Prism Team",
	// 	ThemeColor:   "#2196f3",
	// 	LoadingLabel: "Loading Prism Demo...",
	// 	Styles: []string{
	// 		"https://spectrum-css.adobe.com/spectrum-css/dist/index-vars.css",
	// 	},
	// })

	// if err := http.ListenAndServe(":9090", nil); err != nil {
	// 	log.Fatal(err)
	// }
}
