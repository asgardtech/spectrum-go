package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/cli"
	"github.com/maxence-charriere/go-app/v10/pkg/logs"

	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

const (
	defaultTitle       = "Adobe Spectrum Components for Go-App"
	defaultDescription = "A package that allows you to use Adobe Spectrum components in your Go-App projects."
)

type localOptions struct {
	Port int `cli:"p"                 env:"GOAPP_DOCS_PORT"   help:"The port used by the server that serves the PWA."`
}

func main() {
	app.Route("/", app.NewZeroComponentFactory(newPage()))
	app.Route("/checkbox", app.NewZeroComponentFactory(newCheckboxPage()))
	app.Route("/button", app.NewZeroComponentFactory(newButtonPage()))
	app.Route("/badge", app.NewZeroComponentFactory(newBadgePage()))
	app.Route("/switch", app.NewZeroComponentFactory(newSwitchPage()))
	app.Route("/radio", app.NewZeroComponentFactory(newRadioPage()))
	app.Route("/progress-bar", app.NewZeroComponentFactory(newProgressBarPage()))
	app.Route("/progress-circle", app.NewZeroComponentFactory(newProgressCirclePage()))
	app.Route("/link", app.NewZeroComponentFactory(newLinkPage()))
	app.Route("/icon", app.NewZeroComponentFactory(newIconPage()))
	app.Route("/avatar", app.NewZeroComponentFactory(newAvatarPage()))
	app.Route("/status-light", app.NewZeroComponentFactory(newStatusLightPage()))
	app.Route("/accordion", app.NewZeroComponentFactory(newAccordionPage()))
	app.Route("/coachmark", app.NewZeroComponentFactory(newCoachmarkPage()))
	app.Route("/alert-banner", app.NewZeroComponentFactory(newAlertBannerPage()))
	app.Route("/alert-dialog", app.NewZeroComponentFactory(newAlertDialogPage()))
	app.Route("/asset", app.NewZeroComponentFactory(newAssetPage()))
	app.Route("/button-group", app.NewZeroComponentFactory(newButtonGroupPage()))
	app.Route("/card", app.NewZeroComponentFactory(newCardPage()))
	app.Route("/toast", app.NewZeroComponentFactory(newToastPage()))
	app.Route("/tooltip", app.NewZeroComponentFactory(newTooltipPage()))
	app.Route("/breadcrumbs", app.NewZeroComponentFactory(newBreadcrumbsPage()))
	app.Route("/tabs", app.NewZeroComponentFactory(newTabsPage()))
	app.Route("/tags", app.NewZeroComponentFactory(newTagsPage()))
	app.Route("/action-button", app.NewZeroComponentFactory(newActionButtonPage()))
	app.Route("/divider", app.NewZeroComponentFactory(newDividerPage()))
	app.Route("/textfield", app.NewZeroComponentFactory(newTextfieldPage()))
	app.Route("/textarea", app.NewZeroComponentFactory(newTextareaPage()))
	app.Route("/popover", app.NewZeroComponentFactory(newPopoverPage()))
	app.Route("/menu", app.NewZeroComponentFactory(newMenuPage()))
	app.Route("/meter", app.NewZeroComponentFactory(newMeterPage()))
	app.Route("/search", app.NewZeroComponentFactory(newSearchPage()))
	app.Route("/field-label", app.NewZeroComponentFactory(newFieldLabelPage()))
	app.Route("/help-text", app.NewZeroComponentFactory(newHelpTextPage()))
	app.Route("/slider", app.NewZeroComponentFactory(newSliderPage()))
	app.Route("/sidenav", app.NewZeroComponentFactory(newSidenavPage()))
	app.Route("/number-field", app.NewZeroComponentFactory(newNumberFieldPage()))
	app.Route("/combobox", app.NewZeroComponentFactory(newComboboxPage()))
	app.Route("/split-view", app.NewZeroComponentFactory(newSplitViewPage()))
	app.Route("/dialog", app.NewZeroComponentFactory(newDialogPage()))
	app.Route("/contextual-help", app.NewZeroComponentFactory(newContextualHelpPage()))
	app.Route("/field-group", app.NewZeroComponentFactory(newFieldGroupPage()))
	app.Route("/illustrated-message", app.NewZeroComponentFactory(newIllustratedMessagePage()))
	app.Route("/thumbnail", app.NewZeroComponentFactory(newThumbnailPage()))
	app.Route("/picker", app.NewZeroComponentFactory(newPickerPage()))
	app.Route("/picker-button", app.NewZeroComponentFactory(newPickerButtonPage()))
	app.Route("/swatch", app.NewZeroComponentFactory(newSwatchPage()))
	app.Route("/swatch-group", app.NewZeroComponentFactory(newSwatchGroupPage()))
	app.Route("/action-bar", app.NewZeroComponentFactory(newActionBarPage()))
	app.Route("/top-nav", app.NewZeroComponentFactory(newTopNavPage()))
	app.Route("/overlay", app.NewZeroComponentFactory(newOverlayPage()))
	app.Route("/overlay-trigger", app.NewZeroComponentFactory(newOverlayTriggerPage()))

	app.RunWhenOnBrowser()

	ctx := signals.SetupSignalHandler()

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
		BackgroundColor: "#d2d2d2",
		ThemeColor:      "#d2d2d2",
		LoadingLabel:    "go-app adobe spectrum components {progress}%",
		Styles: []string{
			"/web/app.css",
		},
		Scripts: []string{},
		RawHeaders: []string{
			// Load Spectrum Web Components as recommended in their docs
			`<script src="https://jspm.dev/@spectrum-web-components/bundle/elements.js" type="module" async></script>`,
		},
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
			http.StripPrefix("/web/", http.FileServer(http.Dir("./showcase/web"))).ServeHTTP(w, r)
		}),
	)
	http.Handle("/", h)

	s := http.Server{
		Addr: fmt.Sprintf(":%v", opts.Port),
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
		log.Println("Starting server on http://localhost:8080")
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for a signal
	<-sigChan

	// Cleanup
	log.Println("Shutting down...")
}
