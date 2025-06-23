package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ExamplePageState demonstrates the recommended state management patterns
type ExamplePageState struct {
	// Page-level state
	Title       string
	Description string
	
	// Data management using ViewModel
	PayoutsViewModel *ViewModel[Payout]
	
	// Individual modals and UI state
	ConfirmDeleteDialog *ConfirmDialog
	ToastManager        *ToastManager
	
	// Form state
	FilterForm      *FilterFormState
	CreatePayoutForm *CreatePayoutFormState
	
	// User context
	CurrentUser *User
}

// FilterFormState represents filters for the page
type FilterFormState struct {
	Status      string
	DateFrom    string
	DateTo      string
	SearchQuery string
}

// CreatePayoutFormState represents form data for creating payouts
type CreatePayoutFormState struct {
	Amount      float64
	Description string
	Recipient   string
	LoaderState *LoaderStateMachine
}

func NewExamplePageState() *ExamplePageState {
	return &ExamplePageState{
		Title:               "Example Page",
		Description:         "Demonstrates prism state management patterns",
		PayoutsViewModel:    NewViewModel[Payout](),
		ConfirmDeleteDialog: NewConfirmDialog(),
		ToastManager:        NewToastManager(),
		FilterForm:          NewFilterFormState(),
		CreatePayoutForm:    NewCreatePayoutFormState(),
	}
}

func NewFilterFormState() *FilterFormState {
	return &FilterFormState{
		Status: "all",
	}
}

func NewCreatePayoutFormState() *CreatePayoutFormState {
	return &CreatePayoutFormState{
		LoaderState: NewLoaderState(),
	}
}

// Derived state helpers
func (eps *ExamplePageState) ShouldShowEmpty() bool {
	return eps.PayoutsViewModel.ShouldShowEmpty()
}

func (eps *ExamplePageState) ShouldShowLoading() bool {
	return eps.PayoutsViewModel.ShouldShowLoading()
}

func (eps *ExamplePageState) ShouldShowError() bool {
	return eps.PayoutsViewModel.ShouldShowError()
}

func (eps *ExamplePageState) HasFiltersApplied() bool {
	return eps.FilterForm.Status != "all" ||
		eps.FilterForm.DateFrom != "" ||
		eps.FilterForm.DateTo != "" ||
		eps.FilterForm.SearchQuery != ""
}

func (eps *ExamplePageState) GetSelectedPayoutCount() int {
	count := 0
	for _, payout := range eps.PayoutsViewModel.Items {
		if payout.Selected {
			count++
		}
	}
	return count
}

func (eps *ExamplePageState) HasSelectedPayouts() bool {
	return eps.GetSelectedPayoutCount() > 0
}

func (eps *ExamplePageState) CanCreatePayout() bool {
	return eps.CurrentUser != nil && eps.CurrentUser.LoggedIn
}

func (eps *ExamplePageState) IsFormValid() bool {
	return eps.CreatePayoutForm.Amount > 0 &&
		eps.CreatePayoutForm.Description != "" &&
		eps.CreatePayoutForm.Recipient != ""
}

// WithX pattern methods for state mutations
func (eps *ExamplePageState) WithCurrentUser(user *User) *ExamplePageState {
	eps.CurrentUser = user
	return eps
}

func (eps *ExamplePageState) WithPayouts(payouts []Payout) *ExamplePageState {
	eps.PayoutsViewModel.WithItems(payouts)
	return eps
}

func (eps *ExamplePageState) WithFilter(key, value string) *ExamplePageState {
	eps.PayoutsViewModel.WithFilter(key, value)
	return eps
}

func (eps *ExamplePageState) ShowCreateModal() *ExamplePageState {
	eps.PayoutsViewModel.ShowCreateModal()
	return eps
}

func (eps *ExamplePageState) ShowEditModal(payout *Payout) *ExamplePageState {
	eps.PayoutsViewModel.ShowEditModal(payout)
	return eps
}

func (eps *ExamplePageState) HideModal() *ExamplePageState {
	eps.PayoutsViewModel.HideModal()
	return eps
}

func (eps *ExamplePageState) ShowDeleteConfirm(payout *Payout) *ExamplePageState {
	eps.ConfirmDeleteDialog = ShowDeleteConfirm("payout", func(ctx app.Context) {
		// Handle delete logic here
		eps.ToastManager.ShowSuccess("Payout deleted successfully")
	})
	return eps
}

func (eps *ExamplePageState) ShowToast(message string, variant ToastVariant) *ExamplePageState {
	eps.ToastManager.Show(message, variant, 5000)
	return eps
}