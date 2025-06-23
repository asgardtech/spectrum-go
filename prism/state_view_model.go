package prism

// ViewModel is a generic view model for managing collections with common patterns
type ViewModel[T any] struct {
	Items        []T
	Selected     *T
	Filters      map[string]string
	ShowModal    bool
	ModalData    *T
	ModalMode    string // "create", "edit", "view"
	LoaderState  *LoaderStateMachine
	SearchQuery  string
	SortBy       string
	SortDesc     bool
	Page         int
	PageSize     int
	TotalItems   int
}

func NewViewModel[T any]() *ViewModel[T] {
	return &ViewModel[T]{
		Items:       make([]T, 0),
		Filters:     make(map[string]string),
		ShowModal:   false,
		LoaderState: NewLoaderState(),
		Page:        1,
		PageSize:    10,
	}
}

// Selection methods
func (vm *ViewModel[T]) WithSelected(item *T) *ViewModel[T] {
	vm.Selected = item
	return vm
}

func (vm *ViewModel[T]) ClearSelected() *ViewModel[T] {
	vm.Selected = nil
	return vm
}

func (vm *ViewModel[T]) HasSelected() bool {
	return vm.Selected != nil
}

// Modal methods
func (vm *ViewModel[T]) ShowCreateModal() *ViewModel[T] {
	vm.ShowModal = true
	vm.ModalMode = "create"
	vm.ModalData = nil
	return vm
}

func (vm *ViewModel[T]) ShowEditModal(item *T) *ViewModel[T] {
	vm.ShowModal = true
	vm.ModalMode = "edit"
	vm.ModalData = item
	return vm
}

func (vm *ViewModel[T]) ShowViewModal(item *T) *ViewModel[T] {
	vm.ShowModal = true
	vm.ModalMode = "view"
	vm.ModalData = item
	return vm
}

func (vm *ViewModel[T]) HideModal() *ViewModel[T] {
	vm.ShowModal = false
	vm.ModalMode = ""
	vm.ModalData = nil
	return vm
}

func (vm *ViewModel[T]) IsModalCreate() bool { return vm.ModalMode == "create" }
func (vm *ViewModel[T]) IsModalEdit() bool   { return vm.ModalMode == "edit" }
func (vm *ViewModel[T]) IsModalView() bool   { return vm.ModalMode == "view" }

// Filter methods
func (vm *ViewModel[T]) WithFilter(key, value string) *ViewModel[T] {
	vm.Filters[key] = value
	return vm
}

func (vm *ViewModel[T]) ClearFilter(key string) *ViewModel[T] {
	delete(vm.Filters, key)
	return vm
}

func (vm *ViewModel[T]) ClearAllFilters() *ViewModel[T] {
	vm.Filters = make(map[string]string)
	return vm
}

func (vm *ViewModel[T]) HasFilter(key string) bool {
	_, exists := vm.Filters[key]
	return exists
}

func (vm *ViewModel[T]) GetFilter(key string) string {
	return vm.Filters[key]
}

func (vm *ViewModel[T]) HasAnyFilters() bool {
	return len(vm.Filters) > 0
}

// Search methods
func (vm *ViewModel[T]) WithSearch(query string) *ViewModel[T] {
	vm.SearchQuery = query
	return vm
}

func (vm *ViewModel[T]) ClearSearch() *ViewModel[T] {
	vm.SearchQuery = ""
	return vm
}

func (vm *ViewModel[T]) HasSearch() bool {
	return vm.SearchQuery != ""
}

// Sort methods
func (vm *ViewModel[T]) WithSort(field string, desc bool) *ViewModel[T] {
	vm.SortBy = field
	vm.SortDesc = desc
	return vm
}

func (vm *ViewModel[T]) ToggleSort(field string) *ViewModel[T] {
	if vm.SortBy == field {
		vm.SortDesc = !vm.SortDesc
	} else {
		vm.SortBy = field
		vm.SortDesc = false
	}
	return vm
}

func (vm *ViewModel[T]) IsSortedBy(field string) bool {
	return vm.SortBy == field
}

func (vm *ViewModel[T]) IsSortedDesc() bool {
	return vm.SortDesc
}

// Pagination methods
func (vm *ViewModel[T]) WithPage(page int) *ViewModel[T] {
	vm.Page = page
	return vm
}

func (vm *ViewModel[T]) NextPage() *ViewModel[T] {
	if vm.HasNextPage() {
		vm.Page++
	}
	return vm
}

func (vm *ViewModel[T]) PrevPage() *ViewModel[T] {
	if vm.HasPrevPage() {
		vm.Page--
	}
	return vm
}

func (vm *ViewModel[T]) HasNextPage() bool {
	return vm.Page*vm.PageSize < vm.TotalItems
}

func (vm *ViewModel[T]) HasPrevPage() bool {
	return vm.Page > 1
}

func (vm *ViewModel[T]) GetTotalPages() int {
	if vm.PageSize == 0 {
		return 1
	}
	return (vm.TotalItems + vm.PageSize - 1) / vm.PageSize
}

// Data methods
func (vm *ViewModel[T]) WithItems(items []T) *ViewModel[T] {
	vm.Items = items
	return vm
}

func (vm *ViewModel[T]) AddItem(item T) *ViewModel[T] {
	vm.Items = append(vm.Items, item)
	return vm
}

func (vm *ViewModel[T]) RemoveItem(index int) *ViewModel[T] {
	if index >= 0 && index < len(vm.Items) {
		vm.Items = append(vm.Items[:index], vm.Items[index+1:]...)
	}
	return vm
}

func (vm *ViewModel[T]) UpdateItem(index int, item T) *ViewModel[T] {
	if index >= 0 && index < len(vm.Items) {
		vm.Items[index] = item
	}
	return vm
}

func (vm *ViewModel[T]) GetItemsCount() int {
	return len(vm.Items)
}

func (vm *ViewModel[T]) HasItems() bool {
	return len(vm.Items) > 0
}

func (vm *ViewModel[T]) IsEmpty() bool {
	return len(vm.Items) == 0
}

// Derived state helpers
func (vm *ViewModel[T]) ShouldShowEmpty() bool {
	return vm.IsEmpty() && vm.LoaderState.IsReady()
}

func (vm *ViewModel[T]) ShouldShowLoading() bool {
	return vm.LoaderState.IsLoading()
}

func (vm *ViewModel[T]) ShouldShowError() bool {
	return vm.LoaderState.IsError()
}

func (vm *ViewModel[T]) ShouldShowContent() bool {
	return vm.LoaderState.IsReady() && vm.HasItems()
}