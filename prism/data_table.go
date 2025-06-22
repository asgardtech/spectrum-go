package prism

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// DataTable is a generic table component. Only a stub – it currently renders the
// number of rows and column template but wires the public API so user code
// compiles.

type DataTable[T any] struct {
	app.Compo

	dense   bool
	columns []int
	rows    []T
	rowAct  func(int) []TableAction
}

func DataTable[T any]() *DataTable[T] { return &DataTable[T]{dense: false} }

func (t *DataTable[T]) Dense(d bool) *DataTable[T]                          { t.dense = d; return t }
func (t *DataTable[T]) Columns(cs ...int) *DataTable[T]                     { t.columns = cs; return t }
func (t *DataTable[T]) Rows(r []T) *DataTable[T]                            { t.rows = r; return t }
func (t *DataTable[T]) RowActions(fn func(int) []TableAction) *DataTable[T] { t.rowAct = fn; return t }

func (t *DataTable[T]) Render() app.UI {
	cols := make([]string, len(t.columns))
	for i, c := range t.columns {
		cols[i] = fmt.Sprintf("%dfr", c)
	}
	summary := fmt.Sprintf("DataTable stub – %d rows, columns: %s", len(t.rows), strings.Join(cols, " "))
	return app.Div().Text(summary)
}
