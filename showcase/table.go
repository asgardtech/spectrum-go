package main

import (
	"fmt"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type tablePage struct {
	app.Compo
	ctx      app.Context
	selected map[int]bool
}

type tableRowData struct {
	ID   int
	Col1 string
	Col2 string
}

func newTablePage() *tablePage {
	return &tablePage{
		selected: make(map[int]bool),
	}
}

func (p *tablePage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *tablePage) onSelect(id int) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		if p.selected[id] {
			delete(p.selected, id)
		} else {
			p.selected[id] = true
		}
		p.ctx.Update()
	}
}

func (p *tablePage) Render() app.UI {
	dataRows := []tableRowData{
		{ID: 1, Col1: "Row 1, Cell 1", Col2: "Row 1, Cell 2"},
		{ID: 2, Col1: "Row 2, Cell 1", Col2: "Row 2, Cell 2"},
		{ID: 3, Col1: "Row 3, Cell 1", Col2: "Row 3, Cell 2"},
	}

	selectedItems := []app.UI{}
	for id := range p.selected {
		selectedItems = append(selectedItems, app.Li().Text(fmt.Sprintf("Row %d", id)))
	}

	tableRows := make([]sp.ITableRow, len(dataRows))
	for i, data := range dataRows {
		tableRows[i] = sp.TableRow().
			AddCell(sp.TableCell().Body(
				sp.Checkbox().OnChange(p.onSelect(data.ID)),
			)).
			AddCell(sp.TableCell().Text(data.Col1)).
			AddCell(sp.TableCell().Text(data.Col2))
	}

	return newPage().
		Content(
			app.H1().Text("Table"),
			app.P().Text("A component for displaying tabular data."),
			app.If(len(selectedItems) > 0, func() app.UI {
				return app.Div().Body(
					app.H3().Text("Selected Rows:"),
					app.Ul().Body(selectedItems...),
				)
			}),
			app.H2().Text("Standard"),
			sp.Table().Body(
				sp.TableHead().AddRow(
					sp.TableRow().
						AddHeadCell(sp.TableHeadCell().Body(sp.Checkbox())).
						AddHeadCell(sp.TableHeadCell().Text("Column 1")).
						AddHeadCell(sp.TableHeadCell().Text("Column 2")),
				),
				sp.TableBody().Rows(tableRows...),
			),
		)
}
