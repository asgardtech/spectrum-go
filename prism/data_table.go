package prism

import (
	"fmt"
	"sort"
	"strings"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TableColumn defines a column in the data table
type TableColumn[T any] struct {
	Key         string
	Label       string
	Width       string
	Sortable    bool
	Filterable  bool
	Render      func(T) app.UI
	Compare     func(T, T) int
	FilterType  string // "text", "select", "date", "number"
	FilterOptions []string
}

// TableFilter represents a filter for a column
type TableFilter struct {
	Column string
	Value  string
	Type   string
}

// SortConfig represents sorting configuration
type SortConfig struct {
	Column string
	Desc   bool
}

// TableState holds the current state of the table
type TableState[T any] struct {
	Data           []T
	FilteredData   []T
	Filters        []TableFilter
	Sort           *SortConfig
	SelectedRows   map[int]bool
	CurrentPage    int
	PageSize       int
	Loading        bool
	Error          string
}

// DataTable is a comprehensive table component with filtering, sorting, and pagination
type DataTable[T any] struct {
	app.Compo

	Columns       []TableColumn[T]
	State         *TableState[T]
	Dense         bool
	Selectable    bool
	ShowFilters   bool
	ShowPagination bool
	OnRowClick    func(T, int)
	OnSelectionChange func([]int)
	EmptyMessage  string
	LoadingMessage string
}

// NewDataTable creates a new data table
func NewDataTable[T any]() *DataTable[T] {
	return &DataTable[T]{
		State: &TableState[T]{
			Data:         make([]T, 0),
			FilteredData: make([]T, 0),
			Filters:      make([]TableFilter, 0),
			SelectedRows: make(map[int]bool),
			CurrentPage:  1,
			PageSize:     10,
			Loading:      false,
		},
		Dense:          false,
		Selectable:     false,
		ShowFilters:    true,
		ShowPagination: true,
		EmptyMessage:   "No data available",
		LoadingMessage: "Loading...",
	}
}

func (dt *DataTable[T]) WithColumns(columns ...TableColumn[T]) *DataTable[T] {
	dt.Columns = columns
	return dt
}

func (dt *DataTable[T]) WithData(data []T) *DataTable[T] {
	dt.State.Data = data
	dt.applyFiltersAndSort()
	return dt
}

func (dt *DataTable[T]) WithDense(dense bool) *DataTable[T] {
	dt.Dense = dense
	return dt
}

func (dt *DataTable[T]) WithSelectable(selectable bool) *DataTable[T] {
	dt.Selectable = selectable
	return dt
}

func (dt *DataTable[T]) WithShowFilters(show bool) *DataTable[T] {
	dt.ShowFilters = show
	return dt
}

func (dt *DataTable[T]) WithShowPagination(show bool) *DataTable[T] {
	dt.ShowPagination = show
	return dt
}

func (dt *DataTable[T]) WithPageSize(size int) *DataTable[T] {
	dt.State.PageSize = size
	return dt
}

func (dt *DataTable[T]) WithOnRowClick(fn func(T, int)) *DataTable[T] {
	dt.OnRowClick = fn
	return dt
}

func (dt *DataTable[T]) WithOnSelectionChange(fn func([]int)) *DataTable[T] {
	dt.OnSelectionChange = fn
	return dt
}

func (dt *DataTable[T]) WithEmptyMessage(message string) *DataTable[T] {
	dt.EmptyMessage = message
	return dt
}

func (dt *DataTable[T]) WithLoadingMessage(message string) *DataTable[T] {
	dt.LoadingMessage = message
	return dt
}

// AddFilter adds a filter to the table
func (dt *DataTable[T]) AddFilter(column, value, filterType string) {
	// Remove existing filter for this column
	for i, filter := range dt.State.Filters {
		if filter.Column == column {
			dt.State.Filters = append(dt.State.Filters[:i], dt.State.Filters[i+1:]...)
			break
		}
	}

	// Add new filter if value is not empty
	if value != "" {
		dt.State.Filters = append(dt.State.Filters, TableFilter{
			Column: column,
			Value:  value,
			Type:   filterType,
		})
	}

	dt.applyFiltersAndSort()
}

// SetSort sets the sort configuration
func (dt *DataTable[T]) SetSort(column string, desc bool) {
	if dt.State.Sort != nil && dt.State.Sort.Column == column {
		dt.State.Sort.Desc = !dt.State.Sort.Desc
	} else {
		dt.State.Sort = &SortConfig{
			Column: column,
			Desc:   desc,
		}
	}
	dt.applyFiltersAndSort()
}

// ToggleRowSelection toggles the selection of a row
func (dt *DataTable[T]) ToggleRowSelection(index int) {
	if dt.State.SelectedRows[index] {
		delete(dt.State.SelectedRows, index)
	} else {
		dt.State.SelectedRows[index] = true
	}

	if dt.OnSelectionChange != nil {
		selectedIndices := make([]int, 0)
		for idx := range dt.State.SelectedRows {
			selectedIndices = append(selectedIndices, idx)
		}
		dt.OnSelectionChange(selectedIndices)
	}
}

// SelectAll selects all rows
func (dt *DataTable[T]) SelectAll() {
	dt.State.SelectedRows = make(map[int]bool)
	for i := range dt.State.FilteredData {
		dt.State.SelectedRows[i] = true
	}

	if dt.OnSelectionChange != nil {
		selectedIndices := make([]int, 0)
		for idx := range dt.State.SelectedRows {
			selectedIndices = append(selectedIndices, idx)
		}
		dt.OnSelectionChange(selectedIndices)
	}
}

// ClearSelection clears all row selections
func (dt *DataTable[T]) ClearSelection() {
	dt.State.SelectedRows = make(map[int]bool)

	if dt.OnSelectionChange != nil {
		dt.OnSelectionChange([]int{})
	}
}

// applyFiltersAndSort applies filters and sorting to the data
func (dt *DataTable[T]) applyFiltersAndSort() {
	// Start with all data
	filtered := make([]T, len(dt.State.Data))
	copy(filtered, dt.State.Data)

	// Apply filters
	for _, filter := range dt.State.Filters {
		filtered = dt.applyFilter(filtered, filter)
	}

	// Apply sorting
	if dt.State.Sort != nil {
		filtered = dt.applySort(filtered, dt.State.Sort)
	}

	dt.State.FilteredData = filtered
	
	// Reset to first page when data changes
	dt.State.CurrentPage = 1
}

// applyFilter applies a single filter to the data
func (dt *DataTable[T]) applyFilter(data []T, filter TableFilter) []T {
	column := dt.getColumn(filter.Column)
	if column == nil {
		return data
	}

	var filtered []T
	for _, item := range data {
		if dt.matchesFilter(item, filter, column) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// matchesFilter checks if an item matches a filter
func (dt *DataTable[T]) matchesFilter(item T, filter TableFilter, column *TableColumn[T]) bool {
	// This is a simple text-based filter
	// In a real implementation, you'd extract the field value and compare properly
	
	// For now, we'll assume the render function produces text we can search
	if column.Render != nil {
		// This is a simplified approach - in practice you'd need a way to extract field values
		return strings.Contains(strings.ToLower(fmt.Sprintf("%v", item)), strings.ToLower(filter.Value))
	}
	
	return true
}

// applySort applies sorting to the data
func (dt *DataTable[T]) applySort(data []T, sortConfig *SortConfig) []T {
	column := dt.getColumn(sortConfig.Column)
	if column == nil || column.Compare == nil {
		return data
	}

	sorted := make([]T, len(data))
	copy(sorted, data)

	sort.Slice(sorted, func(i, j int) bool {
		result := column.Compare(sorted[i], sorted[j])
		if sortConfig.Desc {
			return result > 0
		}
		return result < 0
	})

	return sorted
}

// getColumn gets a column by key
func (dt *DataTable[T]) getColumn(key string) *TableColumn[T] {
	for i := range dt.Columns {
		if dt.Columns[i].Key == key {
			return &dt.Columns[i]
		}
	}
	return nil
}

// getPaginatedData gets the data for the current page
func (dt *DataTable[T]) getPaginatedData() []T {
	if !dt.ShowPagination {
		return dt.State.FilteredData
	}

	start := (dt.State.CurrentPage - 1) * dt.State.PageSize
	end := start + dt.State.PageSize

	if start >= len(dt.State.FilteredData) {
		return []T{}
	}

	if end > len(dt.State.FilteredData) {
		end = len(dt.State.FilteredData)
	}

	return dt.State.FilteredData[start:end]
}

// getTotalPages gets the total number of pages
func (dt *DataTable[T]) getTotalPages() int {
	if dt.State.PageSize == 0 {
		return 1
	}
	return (len(dt.State.FilteredData) + dt.State.PageSize - 1) / dt.State.PageSize
}

func (dt *DataTable[T]) Render() app.UI {
	return app.Div().
		Class("prism-data-table").
		Body(
			dt.renderFilters(),
			dt.renderTable(),
			dt.renderPagination(),
		)
}

func (dt *DataTable[T]) renderFilters() app.UI {
	if !dt.ShowFilters {
		return app.Div()
	}

	var filterElements []app.UI
	for _, column := range dt.Columns {
		if column.Filterable {
			filterElements = append(filterElements, dt.renderColumnFilter(column))
		}
	}

	if len(filterElements) == 0 {
		return app.Div()
	}

	return app.Div().
		Class("prism-table-filters").
		Style("display", "flex").
		Style("gap", "16px").
		Style("margin-bottom", "16px").
		Style("padding", "16px").
		Style("background", "#f5f5f5").
		Style("border-radius", "4px").
		Body(filterElements...)
}

func (dt *DataTable[T]) renderColumnFilter(column TableColumn[T]) app.UI {
	currentFilter := ""
	for _, filter := range dt.State.Filters {
		if filter.Column == column.Key {
			currentFilter = filter.Value
			break
		}
	}

	return app.Div().
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("gap", "4px").
		Body(
			app.Label().
				Style("font-size", "12px").
				Style("font-weight", "600").
				Text(column.Label),
			sp.Textfield().
				Placeholder("Filter "+column.Label).
				Value(currentFilter).
				OnInput(func(ctx app.Context, e app.Event) {
					value := ctx.JSSrc().Get("value").String()
					dt.AddFilter(column.Key, value, "text")
					ctx.Reload()
				}),
		)
}

func (dt *DataTable[T]) renderTable() app.UI {
	if dt.State.Loading {
		return dt.renderLoading()
	}

	if len(dt.State.FilteredData) == 0 {
		return dt.renderEmpty()
	}

	return app.Div().
		Class("prism-table-container").
		Style("border", "1px solid #e0e0e0").
		Style("border-radius", "4px").
		Style("overflow", "hidden").
		Body(
			dt.renderTableHeader(),
			dt.renderTableBody(),
		)
}

func (dt *DataTable[T]) renderTableHeader() app.UI {
	headerCells := []app.UI{}

	// Selection column
	if dt.Selectable {
		allSelected := len(dt.State.SelectedRows) == len(dt.State.FilteredData) && len(dt.State.FilteredData) > 0
		headerCells = append(headerCells,
			app.Th().
				Style("width", "40px").
				Style("padding", "12px").
				Style("background", "#f5f5f5").
				Style("border-bottom", "1px solid #e0e0e0").
				Body(
					sp.Checkbox().
						Checked(allSelected).
						OnChange(func(ctx app.Context, e app.Event) {
							if allSelected {
								dt.ClearSelection()
							} else {
								dt.SelectAll()
							}
							ctx.Reload()
						}),
				),
		)
	}

	// Data columns
	for _, column := range dt.Columns {
		headerCells = append(headerCells, dt.renderHeaderCell(column))
	}

	return app.THead().
		Body(
			app.Tr().Body(headerCells...),
		)
}

func (dt *DataTable[T]) renderHeaderCell(column TableColumn[T]) app.UI {
	cell := app.Th().
		Style("padding", "12px").
		Style("background", "#f5f5f5").
		Style("border-bottom", "1px solid #e0e0e0").
		Style("text-align", "left").
		Style("font-weight", "600")

	if column.Width != "" {
		cell = cell.Style("width", column.Width)
	}

	content := []app.UI{app.Text(column.Label)}

	if column.Sortable {
		isSorted := dt.State.Sort != nil && dt.State.Sort.Column == column.Key
		sortIcon := "↕"
		if isSorted {
			if dt.State.Sort.Desc {
				sortIcon = "↓"
			} else {
				sortIcon = "↑"
			}
		}

		content = append(content,
			app.Span().
				Style("margin-left", "8px").
				Style("color", IfElseString(isSorted, "#0066cc", "#ccc")).
				Text(sortIcon),
		)

		cell = cell.
			Style("cursor", "pointer").
			OnClick(func(ctx app.Context, e app.Event) {
				dt.SetSort(column.Key, false)
				ctx.Reload()
			})
	}

	return cell.Body(content...)
}

func (dt *DataTable[T]) renderTableBody() app.UI {
	data := dt.getPaginatedData()
	rows := make([]app.UI, len(data))

	for i, item := range data {
		rows[i] = dt.renderTableRow(item, i)
	}

	return app.TBody().Body(rows...)
}

func (dt *DataTable[T]) renderTableRow(item T, index int) app.UI {
	cells := []app.UI{}
	isSelected := dt.State.SelectedRows[index]

	// Selection cell
	if dt.Selectable {
		cells = append(cells,
			app.Td().
				Style("padding", "12px").
				Style("border-bottom", "1px solid #e0e0e0").
				Body(
					sp.Checkbox().
						Checked(isSelected).
						OnChange(func(ctx app.Context, e app.Event) {
							dt.ToggleRowSelection(index)
							ctx.Reload()
						}),
				),
		)
	}

	// Data cells
	for _, column := range dt.Columns {
		var cellContent app.UI
		if column.Render != nil {
			cellContent = column.Render(item)
		} else {
			cellContent = app.Text(fmt.Sprintf("%v", item))
		}

		cells = append(cells,
			app.Td().
				Style("padding", "12px").
				Style("border-bottom", "1px solid #e0e0e0").
				Body(cellContent),
		)
	}

	row := app.Tr().
		Style("background", IfElseString(isSelected, "#f0f8ff", "white")).
		Body(cells...)

	if dt.OnRowClick != nil {
		row = row.
			Style("cursor", "pointer").
			OnClick(func(ctx app.Context, e app.Event) {
				dt.OnRowClick(item, index)
			})
	}

	return row
}

func (dt *DataTable[T]) renderLoading() app.UI {
	return app.Div().
		Style("text-align", "center").
		Style("padding", "48px").
		Style("color", "#666").
		Body(
			sp.ProgressCircle().
				Indeterminate(true),
			app.Div().
				Style("margin-top", "16px").
				Text(dt.LoadingMessage),
		)
}

func (dt *DataTable[T]) renderEmpty() app.UI {
	return app.Div().
		Style("text-align", "center").
		Style("padding", "48px").
		Style("color", "#666").
		Body(
			app.Div().
				Style("font-size", "18px").
				Style("margin-bottom", "8px").
				Text("📄"),
			app.Div().Text(dt.EmptyMessage),
		)
}

func (dt *DataTable[T]) renderPagination() app.UI {
	if !dt.ShowPagination || len(dt.State.FilteredData) <= dt.State.PageSize {
		return app.Div()
	}

	totalPages := dt.getTotalPages()
	currentPage := dt.State.CurrentPage

	return app.Div().
		Class("prism-table-pagination").
		Style("display", "flex").
		Style("justify-content", "space-between").
		Style("align-items", "center").
		Style("margin-top", "16px").
		Style("padding", "16px").
		Body(
			app.Div().
				Text(fmt.Sprintf("Showing %d-%d of %d items", 
					(currentPage-1)*dt.State.PageSize+1,
					Min(currentPage*dt.State.PageSize, len(dt.State.FilteredData)),
					len(dt.State.FilteredData))),
			app.Div().
				Style("display", "flex").
				Style("gap", "8px").
				Body(
					sp.Button().
						Text("Previous").
						Disabled(currentPage <= 1).
						OnClick(func(ctx app.Context, e app.Event) {
							if currentPage > 1 {
								dt.State.CurrentPage--
								ctx.Reload()
							}
						}),
					app.Span().
						Style("padding", "8px 16px").
						Text(fmt.Sprintf("Page %d of %d", currentPage, totalPages)),
					sp.Button().
						Text("Next").
						Disabled(currentPage >= totalPages).
						OnClick(func(ctx app.Context, e app.Event) {
							if currentPage < totalPages {
								dt.State.CurrentPage++
								ctx.Reload()
							}
						}),
				),
		)
}
