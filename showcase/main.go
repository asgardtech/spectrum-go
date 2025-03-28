package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/cli"
	"github.com/maxence-charriere/go-app/v10/pkg/errors"
	"github.com/maxence-charriere/go-app/v10/pkg/logs"
)

const (
	defaultTitle       = "Adobe Spectrum Components for Go-App"
	defaultDescription = "A package that allows you to use Adobe Spectrum components in your Go-App projects."
)

type localOptions struct {
	Port int `cli:"p"                 env:"GOAPP_DOCS_PORT"   help:"The port used by the server that serves the PWA."`
}

func main() {
	app.Route("/", app.NewZeroComponentFactory(newHomePage()))

	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	localOpts := localOptions{Port: 7777}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&localOpts)

	h := app.Handler{
		Name:        "Documentation for go-app",
		Title:       defaultTitle,
		Description: defaultDescription,
		Author:      "Vlad Iovanov",
		Image:       "https://go-app.dev/web/images/go-app.png",
		Keywords: []string{
			"go-app",
			"go",
			"golang",
			"app",
			"pwa",
			"progressive web app",
			"webassembly",
			"web assembly",
			"webapp",
			"web",
			"gui",
			"ui",
			"user interface",
			"graphical user interface",
			"frontend",
			"opensource",
			"open source",
			"github",
			"component library",
			"component",
			"components",
			"spectrum",
			"adobe spectrum",
			"adobe",
		},
		BackgroundColor:    "#d2d2d2",
		ThemeColor:         "#d2d2d2",
		LoadingLabel:       "go-app adobe spectrum components {progress}%",
		Styles:             []string{},
		Scripts:            []string{},
		RawHeaders:         []string{},
		CacheableResources: []string{},
	}

	runLocal(ctx, &h, localOpts)

}

func runLocal(ctx context.Context, h *app.Handler, opts localOptions) {
	app.Log(logs.New("starting go-app documentation service").
		WithTag("port", opts.Port).
		WithTag("version", h.Version),
	)

	// Serve the web directory
	http.Handle("/web/",
		// use a logging middleware
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.Log(logs.New("serving web directory").WithTag("path", r.URL.Path))
			http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))).ServeHTTP(w, r)
		}),
	)
	http.Handle("/", h)

	s := http.Server{
		Addr: fmt.Sprintf(":%v", opts.Port),
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Log("command failed:", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
