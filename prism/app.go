package prism

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/logs"
	"github.com/mohae/deepcopy"
	"github.com/rotisserie/eris"
)

type App struct {
	pages       []IPage
	baseURL     string
	title       string
	description string
}

func NewApp() *App {
	return &App{}
}

func (a *App) WithPage(page IPage) *App {
	a.pages = append(a.pages, page)
	return a
}

func (a *App) WithBaseUrl(baseUrl string) *App {
	a.baseURL = baseUrl
	return a
}

func (a *App) WithTitle(title string) *App {
	a.title = title
	return a
}

func (a *App) WithDescription(description string) *App {
	a.description = description
	return a
}

func (a *App) Mount() {
	// TODO: implement routing with go-app once pages are fully defined.
}

func (a *App) Run(ctx context.Context) error {
	// Handle all pages routes
	for _, page := range a.pages {
		pageName := page.GetName()
		pagePath := page.GetUrlPath()

		log.Printf("Registering page [%s] on path [%s]", pageName, pagePath)

		app.Route(pagePath, func() app.Composer {
			return deepcopy.Copy(page).(app.Composer)
		})
	}

	app.RunWhenOnBrowser()

	// Start server when on server
	appHandler := &app.Handler{
		Name:            a.title,
		Description:     a.description,
		Author:          "Prism Team",
		ThemeColor:      "#2196f3",
		LoadingLabel:    "Loading Prism Demo...",
		BackgroundColor: "#d2d2d2",
		Styles: []string{
			"/web/app.css",
		},
		RawHeaders: []string{
			// Load Spectrum Web Components as recommended in their docs
			`<script src="https://jspm.dev/@spectrum-web-components/bundle/elements.js" type="module" async></script>`,
		},
	}

	parsedURL, err := url.Parse(a.baseURL)
	if err != nil {
		return eris.Wrap(err, "failed to parse base URL")
	}

	log.Printf("Starting server on [%s]", parsedURL.Host)

	runLocal(ctx, appHandler, parsedURL.Host)

	return nil
}

func runLocal(ctx context.Context, h *app.Handler, host string) {
	// Serve the web directory
	http.Handle("/web/",
		// use a logging middleware
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/web/app.css" {
				http.ServeFile(w, r, sp.CssContent)
			} else {
				app.Log(logs.New("serving web directory").WithTag("path", r.URL.Path))
				http.StripPrefix("/web/", http.FileServer(http.Dir("./showcase/web"))).ServeHTTP(w, r)
			}
		}),
	)
	http.Handle("/", h)

	s := http.Server{
		Addr: host,
	}

	// Create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	// Start the server
	go func() {
		log.Printf("Starting server on %s", host)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for a signal
	<-sigChan

	// Cleanup
	log.Println("Shutting down...")
}
