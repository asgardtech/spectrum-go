package prism

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

const (
	userStateKey = "current-user"
)

type App struct {
	pageConstructors []PageConstructor
	baseURL          string
	title            string
	description      string
	user             User

	ctxMutex sync.RWMutex
	ctx      app.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) WithPage(pageConstructor PageConstructor) *App {
	a.pageConstructors = append(a.pageConstructors, pageConstructor)
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

func (a *App) SetUser(user User) {
	log.Printf("Setting user: %+v", user)

	a.user = user

	a.ctx.SetState(userStateKey, user).
		Persist().
		Broadcast()
}

func (a *App) setContext(ctx app.Context) {
	a.ctxMutex.Lock()
	defer a.ctxMutex.Unlock()
	a.ctx = ctx
}

func (a *App) GetUser() User {
	return a.user
}

func (a *App) Run() error {
	// Handle all pages routes
	for _, pageConstructor := range a.pageConstructors {
		page := pageConstructor()
		pageName := page.GetName()
		pagePath := page.GetUrlPath()

		// TODO: do a bunch of validation on the page constructor

		log.Printf("Registering page [%s] on path [%s]", pageName, pagePath)

		app.Route(pagePath, func() app.Composer {
			concretePage := pageConstructor()
			concretePage.setApp(a)
			return concretePage
		})
	}

	app.RunWhenOnBrowser()

	ctx := signals.SetupSignalHandler()

	// Start server when on server
	appHandler := &app.Handler{
		Name:            a.title,
		Description:     a.description,
		Author:          "Prism Team",
		LoadingLabel:    "Loading Prism Demo...",
		BackgroundColor: "#d2d2d2",
		ThemeColor:      "#d2d2d2",
		Image:           "https://go-app.dev/web/images/go-app.png",
		Styles: []string{
			"/web/app.css",
		},
		RawHeaders: []string{
			// Load Spectrum Web Components as recommended in their docs
			`
    <script type="importmap">
      {
        "imports": {
          "@spectrum-web-components/bundle/elements.js": "https://ga.jspm.io/npm:@spectrum-web-components/bundle@0.40.0/elements.js"
        },
        "scopes": {
          "https://ga.jspm.io/": {
            "@ctrl/tinycolor": "https://ga.jspm.io/npm:@ctrl/tinycolor@4.1.0/dist/module/public_api.js",
            "@floating-ui/core": "https://ga.jspm.io/npm:@floating-ui/core@1.6.9/dist/floating-ui.core.mjs",
            "@floating-ui/dom": "https://ga.jspm.io/npm:@floating-ui/dom@1.6.13/dist/floating-ui.dom.mjs",
            "@floating-ui/utils": "https://ga.jspm.io/npm:@floating-ui/utils@0.2.9/dist/floating-ui.utils.mjs",
            "@floating-ui/utils/dom": "https://ga.jspm.io/npm:@floating-ui/utils@0.2.9/dist/floating-ui.utils.dom.mjs",
            "@internationalized/number": "https://ga.jspm.io/npm:@internationalized/number@3.6.1/dist/import.mjs",
            "@lit-labs/observers/": "https://ga.jspm.io/npm:@lit-labs/observers@2.0.5/",
            "@lit-labs/virtualizer/": "https://ga.jspm.io/npm:@lit-labs/virtualizer@2.1.0/",
            "@lit/reactive-element": "https://ga.jspm.io/npm:@lit/reactive-element@1.6.3/reactive-element.js",
            "@lit/reactive-element/decorators/": "https://ga.jspm.io/npm:@lit/reactive-element@1.6.3/decorators/",
            "@spectrum-web-components/accordion/": "https://ga.jspm.io/npm:@spectrum-web-components/accordion@0.40.5/",
            "@spectrum-web-components/action-bar/sp-action-bar.js": "https://ga.jspm.io/npm:@spectrum-web-components/action-bar@0.40.5/sp-action-bar.js",
            "@spectrum-web-components/action-button/sp-action-button.js": "https://ga.jspm.io/npm:@spectrum-web-components/action-button@0.40.5/sp-action-button.js",
            "@spectrum-web-components/action-group/sp-action-group.js": "https://ga.jspm.io/npm:@spectrum-web-components/action-group@0.40.5/sp-action-group.js",
            "@spectrum-web-components/action-menu/sp-action-menu.js": "https://ga.jspm.io/npm:@spectrum-web-components/action-menu@0.40.5/sp-action-menu.js",
            "@spectrum-web-components/alert-dialog/": "https://ga.jspm.io/npm:@spectrum-web-components/alert-dialog@0.40.5/",
            "@spectrum-web-components/asset/sp-asset.js": "https://ga.jspm.io/npm:@spectrum-web-components/asset@0.40.5/sp-asset.js",
            "@spectrum-web-components/avatar/sp-avatar.js": "https://ga.jspm.io/npm:@spectrum-web-components/avatar@0.40.5/sp-avatar.js",
            "@spectrum-web-components/badge/sp-badge.js": "https://ga.jspm.io/npm:@spectrum-web-components/badge@0.40.5/sp-badge.js",
            "@spectrum-web-components/banner/sp-banner.js": "https://ga.jspm.io/npm:@spectrum-web-components/banner@0.40.5/sp-banner.js",
            "@spectrum-web-components/base": "https://ga.jspm.io/npm:@spectrum-web-components/base@0.40.5/src/index.js",
            "@spectrum-web-components/base/src/": "https://ga.jspm.io/npm:@spectrum-web-components/base@0.40.5/src/",
            "@spectrum-web-components/button": "https://ga.jspm.io/npm:@spectrum-web-components/button@0.40.5/src/index.js",
            "@spectrum-web-components/button-group/sp-button-group.js": "https://ga.jspm.io/npm:@spectrum-web-components/button-group@0.40.5/sp-button-group.js",
            "@spectrum-web-components/button/": "https://ga.jspm.io/npm:@spectrum-web-components/button@0.40.5/",
            "@spectrum-web-components/card/sp-card.js": "https://ga.jspm.io/npm:@spectrum-web-components/card@0.40.5/sp-card.js",
            "@spectrum-web-components/checkbox/": "https://ga.jspm.io/npm:@spectrum-web-components/checkbox@0.40.5/",
            "@spectrum-web-components/clear-button/src/clear-button.css.js": "https://ga.jspm.io/npm:@spectrum-web-components/clear-button@0.40.5/src/clear-button.css.js",
            "@spectrum-web-components/close-button/src/close-button.css.js": "https://ga.jspm.io/npm:@spectrum-web-components/close-button@0.40.5/src/close-button.css.js",
            "@spectrum-web-components/coachmark/sp-coachmark.js": "https://ga.jspm.io/npm:@spectrum-web-components/coachmark@0.40.5/sp-coachmark.js",
            "@spectrum-web-components/color-area/sp-color-area.js": "https://ga.jspm.io/npm:@spectrum-web-components/color-area@0.40.5/sp-color-area.js",
            "@spectrum-web-components/color-handle/sp-color-handle.js": "https://ga.jspm.io/npm:@spectrum-web-components/color-handle@0.40.5/sp-color-handle.js",
            "@spectrum-web-components/color-loupe/sp-color-loupe.js": "https://ga.jspm.io/npm:@spectrum-web-components/color-loupe@0.40.5/sp-color-loupe.js",
            "@spectrum-web-components/color-slider/sp-color-slider.js": "https://ga.jspm.io/npm:@spectrum-web-components/color-slider@0.40.5/sp-color-slider.js",
            "@spectrum-web-components/color-wheel/sp-color-wheel.js": "https://ga.jspm.io/npm:@spectrum-web-components/color-wheel@0.40.5/sp-color-wheel.js",
            "@spectrum-web-components/dialog/": "https://ga.jspm.io/npm:@spectrum-web-components/dialog@0.40.5/",
            "@spectrum-web-components/divider/": "https://ga.jspm.io/npm:@spectrum-web-components/divider@0.40.5/",
            "@spectrum-web-components/dropzone/sp-dropzone.js": "https://ga.jspm.io/npm:@spectrum-web-components/dropzone@0.40.5/sp-dropzone.js",
            "@spectrum-web-components/field-group": "https://ga.jspm.io/npm:@spectrum-web-components/field-group@0.40.5/src/index.js",
            "@spectrum-web-components/field-group/sp-field-group.js": "https://ga.jspm.io/npm:@spectrum-web-components/field-group@0.40.5/sp-field-group.js",
            "@spectrum-web-components/field-label/sp-field-label.js": "https://ga.jspm.io/npm:@spectrum-web-components/field-label@0.40.5/sp-field-label.js",
            "@spectrum-web-components/help-text/": "https://ga.jspm.io/npm:@spectrum-web-components/help-text@0.40.5/",
            "@spectrum-web-components/icon": "https://ga.jspm.io/npm:@spectrum-web-components/icon@0.40.5/src/index.js",
            "@spectrum-web-components/icon/": "https://ga.jspm.io/npm:@spectrum-web-components/icon@0.40.5/",
            "@spectrum-web-components/icons-ui/icons/": "https://ga.jspm.io/npm:@spectrum-web-components/icons-ui@0.40.5/icons/",
            "@spectrum-web-components/icons-workflow/icons/": "https://ga.jspm.io/npm:@spectrum-web-components/icons-workflow@0.40.5/icons/",
            "@spectrum-web-components/icons/": "https://ga.jspm.io/npm:@spectrum-web-components/icons@0.40.5/",
            "@spectrum-web-components/iconset/src/": "https://ga.jspm.io/npm:@spectrum-web-components/iconset@0.40.5/src/",
            "@spectrum-web-components/illustrated-message/sp-illustrated-message.js": "https://ga.jspm.io/npm:@spectrum-web-components/illustrated-message@0.40.5/sp-illustrated-message.js",
            "@spectrum-web-components/infield-button/sp-infield-button.js": "https://ga.jspm.io/npm:@spectrum-web-components/infield-button@0.40.5/sp-infield-button.js",
            "@spectrum-web-components/link/sp-link.js": "https://ga.jspm.io/npm:@spectrum-web-components/link@0.40.5/sp-link.js",
            "@spectrum-web-components/menu/": "https://ga.jspm.io/npm:@spectrum-web-components/menu@0.40.5/",
            "@spectrum-web-components/meter/sp-meter.js": "https://ga.jspm.io/npm:@spectrum-web-components/meter@0.40.5/sp-meter.js",
            "@spectrum-web-components/modal/src/": "https://ga.jspm.io/npm:@spectrum-web-components/modal@0.40.5/src/",
            "@spectrum-web-components/number-field/sp-number-field.js": "https://ga.jspm.io/npm:@spectrum-web-components/number-field@0.40.5/sp-number-field.js",
            "@spectrum-web-components/opacity-checkerboard/src/": "https://ga.jspm.io/npm:@spectrum-web-components/opacity-checkerboard@0.40.5/src/",
            "@spectrum-web-components/overlay/": "https://ga.jspm.io/npm:@spectrum-web-components/overlay@0.40.5/",
            "@spectrum-web-components/picker": "https://ga.jspm.io/npm:@spectrum-web-components/picker@0.40.5/src/index.js",
            "@spectrum-web-components/picker-button/sp-picker-button.js": "https://ga.jspm.io/npm:@spectrum-web-components/picker-button@0.40.5/sp-picker-button.js",
            "@spectrum-web-components/picker/sp-picker.js": "https://ga.jspm.io/npm:@spectrum-web-components/picker@0.40.5/sp-picker.js",
            "@spectrum-web-components/popover/sp-popover.js": "https://ga.jspm.io/npm:@spectrum-web-components/popover@0.40.5/sp-popover.js",
            "@spectrum-web-components/progress-bar/sp-progress-bar.js": "https://ga.jspm.io/npm:@spectrum-web-components/progress-bar@0.40.5/sp-progress-bar.js",
            "@spectrum-web-components/progress-circle/sp-progress-circle.js": "https://ga.jspm.io/npm:@spectrum-web-components/progress-circle@0.40.5/sp-progress-circle.js",
            "@spectrum-web-components/quick-actions/sp-quick-actions.js": "https://ga.jspm.io/npm:@spectrum-web-components/quick-actions@0.40.5/sp-quick-actions.js",
            "@spectrum-web-components/radio/": "https://ga.jspm.io/npm:@spectrum-web-components/radio@0.40.5/",
            "@spectrum-web-components/reactive-controllers/src/": "https://ga.jspm.io/npm:@spectrum-web-components/reactive-controllers@0.40.5/src/",
            "@spectrum-web-components/search/sp-search.js": "https://ga.jspm.io/npm:@spectrum-web-components/search@0.40.5/sp-search.js",
            "@spectrum-web-components/shared": "https://ga.jspm.io/npm:@spectrum-web-components/shared@0.40.5/src/index.js",
            "@spectrum-web-components/shared/src/": "https://ga.jspm.io/npm:@spectrum-web-components/shared@0.40.5/src/",
            "@spectrum-web-components/sidenav/": "https://ga.jspm.io/npm:@spectrum-web-components/sidenav@0.40.5/",
            "@spectrum-web-components/slider/": "https://ga.jspm.io/npm:@spectrum-web-components/slider@0.40.5/",
            "@spectrum-web-components/split-button/sp-split-button.js": "https://ga.jspm.io/npm:@spectrum-web-components/split-button@0.40.5/sp-split-button.js",
            "@spectrum-web-components/split-view/sp-split-view.js": "https://ga.jspm.io/npm:@spectrum-web-components/split-view@0.40.5/sp-split-view.js",
            "@spectrum-web-components/status-light/sp-status-light.js": "https://ga.jspm.io/npm:@spectrum-web-components/status-light@0.40.5/sp-status-light.js",
            "@spectrum-web-components/styles/": "https://ga.jspm.io/npm:@spectrum-web-components/styles@0.40.5/",
            "@spectrum-web-components/swatch/": "https://ga.jspm.io/npm:@spectrum-web-components/swatch@0.40.5/",
            "@spectrum-web-components/switch/sp-switch.js": "https://ga.jspm.io/npm:@spectrum-web-components/switch@0.40.5/sp-switch.js",
            "@spectrum-web-components/table/": "https://ga.jspm.io/npm:@spectrum-web-components/table@0.40.5/",
            "@spectrum-web-components/tabs/": "https://ga.jspm.io/npm:@spectrum-web-components/tabs@0.40.5/",
            "@spectrum-web-components/tags/": "https://ga.jspm.io/npm:@spectrum-web-components/tags@0.40.5/",
            "@spectrum-web-components/textfield": "https://ga.jspm.io/npm:@spectrum-web-components/textfield@0.40.5/src/index.js",
            "@spectrum-web-components/textfield/sp-textfield.js": "https://ga.jspm.io/npm:@spectrum-web-components/textfield@0.40.5/sp-textfield.js",
            "@spectrum-web-components/theme/": "https://ga.jspm.io/npm:@spectrum-web-components/theme@0.40.5/",
            "@spectrum-web-components/thumbnail/sp-thumbnail.js": "https://ga.jspm.io/npm:@spectrum-web-components/thumbnail@0.40.5/sp-thumbnail.js",
            "@spectrum-web-components/toast/sp-toast.js": "https://ga.jspm.io/npm:@spectrum-web-components/toast@0.40.5/sp-toast.js",
            "@spectrum-web-components/tooltip/sp-tooltip.js": "https://ga.jspm.io/npm:@spectrum-web-components/tooltip@0.40.5/sp-tooltip.js",
            "@spectrum-web-components/top-nav/": "https://ga.jspm.io/npm:@spectrum-web-components/top-nav@0.40.5/",
            "@spectrum-web-components/tray/sp-tray.js": "https://ga.jspm.io/npm:@spectrum-web-components/tray@0.40.5/sp-tray.js",
            "@spectrum-web-components/underlay/sp-underlay.js": "https://ga.jspm.io/npm:@spectrum-web-components/underlay@0.40.5/sp-underlay.js",
            "focus-visible": "https://ga.jspm.io/npm:focus-visible@5.2.1/dist/focus-visible.js",
            "lit": "https://ga.jspm.io/npm:lit@2.8.0/index.js",
            "lit-element/lit-element.js": "https://ga.jspm.io/npm:lit-element@3.3.3/lit-element.js",
            "lit-html": "https://ga.jspm.io/npm:lit-html@2.8.0/lit-html.js",
            "lit-html/": "https://ga.jspm.io/npm:lit-html@3.3.0/",
            "lit/": "https://ga.jspm.io/npm:lit@2.8.0/"
          },
          "https://ga.jspm.io/npm:@lit-labs/virtualizer@2.1.0/": {
            "lit": "https://ga.jspm.io/npm:lit@3.3.0/index.js",
            "lit/": "https://ga.jspm.io/npm:lit@3.3.0/"
          },
          "https://ga.jspm.io/npm:@spectrum-web-components/base@1.5.0/": {
            "@spectrum-web-components/base/src/version.js": "https://ga.jspm.io/npm:@spectrum-web-components/base@1.5.0/src/version.js"
          },
          "https://ga.jspm.io/npm:@spectrum-web-components/button@0.40.5/": {
            "@spectrum-web-components/progress-circle/sp-progress-circle.js": "https://ga.jspm.io/npm:@spectrum-web-components/progress-circle@1.5.0/sp-progress-circle.js"
          },
          "https://ga.jspm.io/npm:@spectrum-web-components/progress-circle@1.5.0/": {
            "@spectrum-web-components/base": "https://ga.jspm.io/npm:@spectrum-web-components/base@1.5.0/src/index.js",
            "@spectrum-web-components/base/src/": "https://ga.jspm.io/npm:@spectrum-web-components/base@1.5.0/src/",
            "@spectrum-web-components/shared/src/get-label-from-slot.js": "https://ga.jspm.io/npm:@spectrum-web-components/shared@1.5.0/src/get-label-from-slot.js"
          },
          "https://ga.jspm.io/npm:lit-element@4.2.0/": {
            "@lit/reactive-element": "https://ga.jspm.io/npm:@lit/reactive-element@2.1.0/reactive-element.js",
            "lit-html": "https://ga.jspm.io/npm:lit-html@3.3.0/lit-html.js"
          },
          "https://ga.jspm.io/npm:lit@2.8.0/": {
            "lit-html/is-server.js": "https://ga.jspm.io/npm:lit-html@2.8.0/is-server.js"
          },
          "https://ga.jspm.io/npm:lit@3.3.0/": {
            "@lit/reactive-element": "https://ga.jspm.io/npm:@lit/reactive-element@2.1.0/reactive-element.js",
            "lit-element/lit-element.js": "https://ga.jspm.io/npm:lit-element@4.2.0/lit-element.js",
            "lit-html": "https://ga.jspm.io/npm:lit-html@3.3.0/lit-html.js",
            "lit-html/": "https://ga.jspm.io/npm:lit-html@3.3.0/"
          }
        }
      }
    </script>
		<script
      async
      src="https://ga.jspm.io/npm:es-module-shims@2.0.10/dist/es-module-shims.js"
      crossorigin="anonymous"
    ></script>

    <script type="module">
      import * as WebComponentsBundleElements from '@spectrum-web-components/bundle/elements.js';
      // Use any Spectrum component
    </script>
			`,
		},
		CacheableResources: []string{},
	}

	parsedURL, err := url.Parse(a.baseURL)
	if err != nil {
		return eris.Wrap(err, "failed to parse base URL")
	}

	log.Printf("Starting server on [%s]", parsedURL.Host)

	runServer(ctx, appHandler, parsedURL.Host)

	return nil
}

func binaryDir() string {
	dir, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(dir)
}

func webDir() string {
	dir := binaryDir()
	if dir == "" {
		return ""
	}
	return filepath.Join(dir, "web")
}

func runServer(ctx context.Context, h *app.Handler, host string) {
	// Serve the web directory
	http.Handle("/web/",
		// use a logging middleware
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/web/app.css" {
				content := sp.CssContent
				w.Header().Set("Content-Type", "text/css; charset=utf-8")
				w.Write([]byte(content))
			} else {
				dir := webDir()
				log.Printf("Serving web directory: %s", dir)
				http.StripPrefix("/web/", http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
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
