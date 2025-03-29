package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// MatchMediaUpdatedSymbol is a sentinel value to indicate when a media query match state has changed
var MatchMediaUpdatedSymbol = "match-media-updated"

// SpectrumMatchMediaController binds a CSS media query to a host element and
// allows for reactivity based on whether that query is currently matching or not
type SpectrumMatchMediaController struct {
	host       app.Composer
	MediaQuery string
	Matches    bool
	mql        app.Value
	ctx        *app.Context
}

// MatchMediaController creates a new SpectrumMatchMediaController
func MatchMediaController(host app.Composer, mediaQuery string) *SpectrumMatchMediaController {
	controller := &SpectrumMatchMediaController{
		host:       host,
		MediaQuery: mediaQuery,
	}

	// Initialize the controller and start listening for changes
	controller.initialize()

	return controller
}

// SetMediaQuery updates the media query for this controller
func (c *SpectrumMatchMediaController) SetMediaQuery(mediaQuery string) *SpectrumMatchMediaController {
	if c.MediaQuery == mediaQuery {
		return c
	}

	c.MediaQuery = mediaQuery
	c.initialize()

	return c
}

// Connect starts listening for media query changes
func (c *SpectrumMatchMediaController) Connect() {
	if c.MediaQuery != "" && c.mql.IsNull() {
		c.initialize()
	}
}

// Disconnect stops listening for media query changes
func (c *SpectrumMatchMediaController) Disconnect() {
	if !c.mql.IsNull() {
		// Remove event listener
		c.mql.Call("removeEventListener", "change", app.FuncOf(c.handleChange))
		c.mql = nil
	}
}

// initialize sets up the MediaQueryList and starts listening for changes
func (c *SpectrumMatchMediaController) initialize() {
	if c.MediaQuery == "" {
		return
	}

	// Clean up any existing listener
	if !c.mql.IsNull() {
		c.mql.Call("removeEventListener", "change", app.FuncOf(c.handleChange))
	}

	// Create a new MediaQueryList
	c.mql = app.Window().Call("matchMedia", c.MediaQuery)

	// Set initial matches state
	c.Matches = c.mql.Get("matches").Bool()

	// Add change listener
	c.mql.Call("addEventListener", "change", app.FuncOf(c.handleChange))
}

// handleChange is called when the media query match state changes
func (c *SpectrumMatchMediaController) handleChange(this app.Value, args []app.Value) interface{} {
	// Update matches state
	oldMatches := c.Matches
	c.Matches = args[0].Get("matches").Bool()

	// Only trigger update if the match state has changed
	if oldMatches != c.Matches && c.ctx != nil {
		c.ctx.Update()
	}

	return nil
}

func (c *SpectrumMatchMediaController) OnMount(ctx app.Context) {
	c.ctx = &ctx
}

func (c *SpectrumMatchMediaController) OnUnmount() {
	c.ctx = nil
}
