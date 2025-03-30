package sp

import (
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type system struct {
	app.Compo
	navCtx         app.Context
	removeListener app.Func
}

type ISystem interface {
	app.UI
}

func System() ISystem {
	return &system{}
}

func (s *system) OnMount(ctx app.Context) {
	s.navCtx = ctx

	if app.IsClient {
		s.removeListener = app.FuncOf(s.interceptLinkClicks)
		app.Window().Get("document").Call("addEventListener", "click", s.removeListener)
	}
}

func (s *system) OnDismount() {
	if app.IsClient && s.removeListener != nil {
		app.Window().Get("document").Call("removeEventListener", "click", s.removeListener)
		s.removeListener.Release()
		s.removeListener = nil
	}
}

func (s *system) Render() app.UI {
	return app.Div() // Invisible; exists only to bootstrap context
}

func (s *system) interceptLinkClicks(this app.Value, args []app.Value) interface{} {
	e := args[0]

	if e.Get("defaultPrevented").Bool() {
		return nil
	}

	if e.Get("button").Int() != 0 || e.Get("ctrlKey").Bool() || e.Get("metaKey").Bool() || e.Get("shiftKey").Bool() {
		return nil
	}

	target := findAnchor(e.Get("target"))
	if target.IsNull() {
		return nil
	}

	if target.Get("target").String() == "_blank" {
		return nil
	}

	href := target.Get("href").String()
	if href == "" || strings.HasPrefix(href, "#") {
		return nil
	}

	origin := app.Window().Get("location").Get("origin").String()
	if !strings.HasPrefix(href, origin) {
		return nil
	}

	path := strings.TrimPrefix(href, origin)

	e.Call("preventDefault")

	s.navCtx.Navigate(path)

	return nil
}

func findAnchor(target app.Value) app.Value {
	for !target.IsNull() && !target.IsUndefined() {
		// If it's an anchor, done
		if strings.ToUpper(target.Get("tagName").String()) == "A" {
			return target
		}

		// Check shadow root
		shadowRoot := target.Get("shadowRoot")
		if !shadowRoot.IsUndefined() && !shadowRoot.IsNull() {
			anchor := shadowRoot.Call("querySelector", "a")
			if anchor.Truthy() {
				return anchor
			}
		}

		// Climb
		target = target.Get("parentElement")
	}
	return app.Null()
}
