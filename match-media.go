package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// MatchMediaUpdatedSymbol is a sentinel value to indicate when a media query match state has changed
var MatchMediaUpdatedSymbol = "match-media-updated"

// spectrumMatchMediaController binds a CSS media query to a host element and
// allows for reactivity based on whether that query is currently matching or not
type spectrumMatchMediaController struct {
	PropHost       app.Composer
	PropMediaQuery string
	PropMatches    bool
	mql            app.Value
	ctx            *app.Context
}

// MatchMediaController creates a new spectrumMatchMediaController
func MatchMediaController(host app.Composer, mediaQuery string) *spectrumMatchMediaController {
	controller := &spectrumMatchMediaController{
		PropHost:       host,
		PropMediaQuery: mediaQuery,
	}

	// Initialize the controller and start listening for changes
	controller.initialize()

	return controller
}

// SetMediaQuery updates the media query for this controller
func (c *spectrumMatchMediaController) SetMediaQuery(mediaQuery string) *spectrumMatchMediaController {
	if c.PropMediaQuery == mediaQuery {
		return c
	}

	c.PropMediaQuery = mediaQuery
	c.initialize()

	return c
}

// Connect starts listening for media query changes
func (c *spectrumMatchMediaController) Connect() {
	if c.PropMediaQuery != "" && c.mql.IsNull() {
		c.initialize()
	}
}

// Disconnect stops listening for media query changes
func (c *spectrumMatchMediaController) Disconnect() {
	if !c.mql.IsNull() {
		// Remove event listener
		c.mql.Call("removeEventListener", "change", app.FuncOf(c.handleChange))
		c.mql = nil
	}
}

// initialize sets up the MediaQueryList and starts listening for changes
func (c *spectrumMatchMediaController) initialize() {
	if c.PropMediaQuery == "" {
		return
	}

	// Clean up any existing listener
	if !c.mql.IsNull() {
		c.mql.Call("removeEventListener", "change", app.FuncOf(c.handleChange))
	}

	// Create a new MediaQueryList
	c.mql = app.Window().Call("matchMedia", c.PropMediaQuery)

	// Set initial matches state
	c.PropMatches = c.mql.Get("matches").Bool()

	// Add change listener
	c.mql.Call("addEventListener", "change", app.FuncOf(c.handleChange))
}

// handleChange is called when the media query match state changes
func (c *spectrumMatchMediaController) handleChange(this app.Value, args []app.Value) interface{} {
	// Update matches state
	oldMatches := c.PropMatches
	c.PropMatches = args[0].Get("matches").Bool()

	// Only trigger update if the match state has changed
	if oldMatches != c.PropMatches && c.ctx != nil {
		c.ctx.Update()
	}

	return nil
}

func (c *spectrumMatchMediaController) OnMount(ctx app.Context) {
	c.ctx = &ctx
}

func (c *spectrumMatchMediaController) OnUnmount() {
	c.ctx = nil
}
