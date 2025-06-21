package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type dropzonePage struct {
	app.Compo
	ctx          app.Context
	droppedFiles []string
}

func newDropzonePage() *dropzonePage {
	return &dropzonePage{}
}

func (p *dropzonePage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *dropzonePage) onDrop(ctx app.Context, e app.Event) {
	e.PreventDefault()
	files := e.Get("dataTransfer").Get("files")
	length := files.Length()
	dropped := make([]string, length)
	for i := 0; i < length; i++ {
		dropped[i] = files.Index(i).Get("name").String()
	}
	p.droppedFiles = dropped
	p.ctx.Update()
}

func (p *dropzonePage) onDragOver(ctx app.Context, e app.Event) {
	e.PreventDefault()
}

func (p *dropzonePage) Render() app.UI {
	return prism.NewLayout().
		Content(
			app.H1().Text("Dropzone"),
			app.P().Text("A dropzone is an area into which objects can be dragged and dropped."),
			app.H2().Text("Standard"),
			app.Div().On("dragover", p.onDragOver).Body(
				sp.Dropzone().Id("dropzone").Body(
					sp.IllustratedMessage().Heading("Drag and drop your file"),
				).OnSpDropzoneDrop(p.onDrop),
			),
			app.If(len(p.droppedFiles) > 0, func() app.UI {
				return app.Div().Body(
					app.H3().Text("Dropped Files:"),
					app.Ul().Body(
						app.Range(p.droppedFiles).Slice(func(i int) app.UI {
							return app.Li().Text(p.droppedFiles[i])
						}),
					),
				)
			}),
		)
}
