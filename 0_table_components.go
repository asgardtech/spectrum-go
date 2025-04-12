package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// --- Interfaces ---

// ITableBody represents the <sp-table-body> element.
type ITableBody interface {
	app.UI
	Styler[ITableBody]
	Classer[ITableBody]
	Ider[ITableBody]
	Rows(...ITableRow) ITableBody
	AddRow(ITableRow) ITableBody
}

// ITableCell represents the <sp-table-cell> element.
type ITableCell interface {
	app.UI
	Styler[ITableCell]
	Classer[ITableCell]
	Ider[ITableCell]
	Body(...app.UI) ITableCell
	AddToBody(app.UI) ITableCell
	Text(string) ITableCell
}

// ITableCheckboxCell represents the <sp-table-checkbox-cell> element.
type ITableCheckboxCell interface {
	app.UI
	Styler[ITableCheckboxCell]
	Classer[ITableCheckboxCell]
	Ider[ITableCheckboxCell]
	Body(...app.UI) ITableCheckboxCell // Usually contains sp-checkbox
	AddToBody(app.UI) ITableCheckboxCell
	Text(string) ITableCheckboxCell
}

// ITableHead represents the <sp-table-head> element.
type ITableHead interface {
	app.UI
	Styler[ITableHead]
	Classer[ITableHead]
	Ider[ITableHead]
	Rows(...ITableRow) ITableHead
	AddRow(ITableRow) ITableHead
}

// ITableHeadCell represents the <sp-table-head-cell> element.
type ITableHeadCell interface {
	app.UI
	Styler[ITableHeadCell]
	Classer[ITableHeadCell]
	Ider[ITableHeadCell]
	Body(...app.UI) ITableHeadCell // Can contain text, icons, etc.
	AddToBody(app.UI) ITableHeadCell
	Text(string) ITableHeadCell
	// TODO: Add properties like SortDirection, SortKey?
}

// ITableRow represents the <sp-table-row> element.
type ITableRow interface {
	app.UI
	Styler[ITableRow]
	Classer[ITableRow]
	Ider[ITableRow]
	Cells(...app.UI) ITableRow                    // Accepts ITableCell, ITableHeadCell, ITableCheckboxCell
	AddCell(ITableCell) ITableRow                 // Convenience method
	AddHeadCell(ITableHeadCell) ITableRow         // Convenience method
	AddCheckboxCell(ITableCheckboxCell) ITableRow // Convenience method
}

// --- Implementations ---

// <sp-table-body>
type tableBody struct {
	app.Compo
	*styler[*tableBody]
	*classer[*tableBody]
	*ider[*tableBody]

	PropRows []ITableRow
}

func TableBody() ITableBody {
	c := &tableBody{
		PropRows: []ITableRow{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableBody) Rows(rows ...ITableRow) ITableBody {
	c.PropRows = rows
	return c
}

func (c *tableBody) AddRow(row ITableRow) ITableBody {
	c.PropRows = append(c.PropRows, row)
	return c
}

func (c *tableBody) Style(key, format string, values ...any) ITableBody {
	return c.styler.Style(key, format, values...)
}
func (c *tableBody) Styles(styles map[string]string) ITableBody { return c.styler.Styles(styles) }
func (c *tableBody) Class(class string) ITableBody              { return c.classer.Class(class) }
func (c *tableBody) Classes(classes ...string) ITableBody       { return c.classer.Classes(classes...) }
func (c *tableBody) Id(id string) ITableBody                    { return c.ider.Id(id) }

func (c *tableBody) Render() app.UI {
	element := app.Elem("sp-table-body")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropRows) > 0 {
		rows := make([]app.UI, len(c.PropRows))
		for i, r := range c.PropRows {
			rows[i] = r
		}
		element = element.Body(rows...)
	}
	return element
}

// <sp-table-cell>
type tableCell struct {
	app.Compo
	*styler[*tableCell]
	*classer[*tableCell]
	*ider[*tableCell]

	PropBody []app.UI
}

func TableCell() ITableCell {
	c := &tableCell{
		PropBody: []app.UI{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableCell) Body(elements ...app.UI) ITableCell {
	c.PropBody = elements
	return c
}

func (c *tableCell) AddToBody(element app.UI) ITableCell {
	c.PropBody = append(c.PropBody, element)
	return c
}

func (c *tableCell) Text(text string) ITableCell {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

func (c *tableCell) Style(key, format string, values ...any) ITableCell {
	return c.styler.Style(key, format, values...)
}
func (c *tableCell) Styles(styles map[string]string) ITableCell { return c.styler.Styles(styles) }
func (c *tableCell) Class(class string) ITableCell              { return c.classer.Class(class) }
func (c *tableCell) Classes(classes ...string) ITableCell       { return c.classer.Classes(classes...) }
func (c *tableCell) Id(id string) ITableCell                    { return c.ider.Id(id) }

func (c *tableCell) Render() app.UI {
	element := app.Elem("sp-table-cell")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropBody) > 0 {
		element = element.Body(c.PropBody...)
	}
	return element
}

// <sp-table-checkbox-cell>
type tableCheckboxCell struct {
	app.Compo
	*styler[*tableCheckboxCell]
	*classer[*tableCheckboxCell]
	*ider[*tableCheckboxCell]

	PropBody []app.UI
}

func TableCheckboxCell() ITableCheckboxCell {
	c := &tableCheckboxCell{
		PropBody: []app.UI{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableCheckboxCell) Body(elements ...app.UI) ITableCheckboxCell {
	c.PropBody = elements
	return c
}

func (c *tableCheckboxCell) AddToBody(element app.UI) ITableCheckboxCell {
	c.PropBody = append(c.PropBody, element)
	return c
}

func (c *tableCheckboxCell) Text(text string) ITableCheckboxCell {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

func (c *tableCheckboxCell) Style(key, format string, values ...any) ITableCheckboxCell {
	return c.styler.Style(key, format, values...)
}
func (c *tableCheckboxCell) Styles(styles map[string]string) ITableCheckboxCell {
	return c.styler.Styles(styles)
}
func (c *tableCheckboxCell) Class(class string) ITableCheckboxCell { return c.classer.Class(class) }
func (c *tableCheckboxCell) Classes(classes ...string) ITableCheckboxCell {
	return c.classer.Classes(classes...)
}
func (c *tableCheckboxCell) Id(id string) ITableCheckboxCell { return c.ider.Id(id) }

func (c *tableCheckboxCell) Render() app.UI {
	element := app.Elem("sp-table-checkbox-cell")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropBody) > 0 {
		element = element.Body(c.PropBody...)
	}
	return element
}

// <sp-table-head>
type tableHead struct {
	app.Compo
	*styler[*tableHead]
	*classer[*tableHead]
	*ider[*tableHead]

	PropRows []ITableRow
}

func TableHead() ITableHead {
	c := &tableHead{
		PropRows: []ITableRow{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableHead) Rows(rows ...ITableRow) ITableHead {
	c.PropRows = rows
	return c
}

func (c *tableHead) AddRow(row ITableRow) ITableHead {
	c.PropRows = append(c.PropRows, row)
	return c
}

func (c *tableHead) Style(key, format string, values ...any) ITableHead {
	return c.styler.Style(key, format, values...)
}
func (c *tableHead) Styles(styles map[string]string) ITableHead { return c.styler.Styles(styles) }
func (c *tableHead) Class(class string) ITableHead              { return c.classer.Class(class) }
func (c *tableHead) Classes(classes ...string) ITableHead       { return c.classer.Classes(classes...) }
func (c *tableHead) Id(id string) ITableHead                    { return c.ider.Id(id) }

func (c *tableHead) Render() app.UI {
	element := app.Elem("sp-table-head")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropRows) > 0 {
		rows := make([]app.UI, len(c.PropRows))
		for i, r := range c.PropRows {
			rows[i] = r
		}
		element = element.Body(rows...)
	}
	return element
}

// <sp-table-head-cell>
type tableHeadCell struct {
	app.Compo
	*styler[*tableHeadCell]
	*classer[*tableHeadCell]
	*ider[*tableHeadCell]

	PropBody []app.UI
}

func TableHeadCell() ITableHeadCell {
	c := &tableHeadCell{
		PropBody: []app.UI{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableHeadCell) Body(elements ...app.UI) ITableHeadCell {
	c.PropBody = elements
	return c
}

func (c *tableHeadCell) AddToBody(element app.UI) ITableHeadCell {
	c.PropBody = append(c.PropBody, element)
	return c
}

func (c *tableHeadCell) Text(text string) ITableHeadCell {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

func (c *tableHeadCell) Style(key, format string, values ...any) ITableHeadCell {
	return c.styler.Style(key, format, values...)
}
func (c *tableHeadCell) Styles(styles map[string]string) ITableHeadCell {
	return c.styler.Styles(styles)
}
func (c *tableHeadCell) Class(class string) ITableHeadCell { return c.classer.Class(class) }
func (c *tableHeadCell) Classes(classes ...string) ITableHeadCell {
	return c.classer.Classes(classes...)
}
func (c *tableHeadCell) Id(id string) ITableHeadCell { return c.ider.Id(id) }

func (c *tableHeadCell) Render() app.UI {
	element := app.Elem("sp-table-head-cell")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropBody) > 0 {
		element = element.Body(c.PropBody...)
	}
	return element
}

// <sp-table-row>
type tableRow struct {
	app.Compo
	*styler[*tableRow]
	*classer[*tableRow]
	*ider[*tableRow]

	PropCells []app.UI
}

func TableRow() ITableRow {
	c := &tableRow{
		PropCells: []app.UI{},
	}
	c.styler = newStyler(c)
	c.classer = newClasser(c)
	c.ider = newIder(c)
	return c
}

func (c *tableRow) Cells(cells ...app.UI) ITableRow {
	c.PropCells = cells
	return c
}

func (c *tableRow) AddCell(cell ITableCell) ITableRow {
	c.PropCells = append(c.PropCells, cell)
	return c
}

func (c *tableRow) AddHeadCell(cell ITableHeadCell) ITableRow {
	c.PropCells = append(c.PropCells, cell)
	return c
}

func (c *tableRow) AddCheckboxCell(cell ITableCheckboxCell) ITableRow {
	c.PropCells = append(c.PropCells, cell)
	return c
}

func (c *tableRow) Style(key, format string, values ...any) ITableRow {
	return c.styler.Style(key, format, values...)
}
func (c *tableRow) Styles(styles map[string]string) ITableRow { return c.styler.Styles(styles) }
func (c *tableRow) Class(class string) ITableRow              { return c.classer.Class(class) }
func (c *tableRow) Classes(classes ...string) ITableRow       { return c.classer.Classes(classes...) }
func (c *tableRow) Id(id string) ITableRow                    { return c.ider.Id(id) }

func (c *tableRow) Render() app.UI {
	element := app.Elem("sp-table-row")
	element = element.Styles(c.styler.styles)
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}
	if len(c.PropCells) > 0 {
		element = element.Body(c.PropCells...)
	}
	return element
}
