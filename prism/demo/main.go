package main

import (
	"log"

	"github.com/asgardtech/spectrum-go/prism"
)

func main() {
	app := prism.NewApp().
		WithBaseUrl("http://localhost:9090").
		WithDescription("Prism Demo").
		WithTitle("Prism Demo").
		WithPage(
			prism.LoginPageConstructor(
				prism.LoginPageOptions{
					PageOptions: prism.PageOptions{
						Name:    "Login",
						UrlPath: "/login",
					},
				},
			),
		).
		WithPage(
			prism.HomePageConstructor(
				prism.PageOptions{
					Name:         "Home",
					UrlPath:      "/",
					SidenavGroup: "Dashboard",
				},
			),
		)

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
