package prism

import (
	"fmt"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type DemoPage struct {
	app.Compo
	Page

	state *DemoPageState
}

type DemoPageState struct {
	ShowModal      bool
	ModalTitle     string
	ModalContent   string
	ToastManager   *ToastManager
	ConfirmDialog  *ConfirmDialog
	SamplePayouts  []Payout
	SelectedPayout *Payout
	LoaderState    *LoaderStateMachine
}

func NewDemoPage() *DemoPage {
	return &DemoPage{
		state: NewDemoPageState(),
	}
}

func NewDemoPageState() *DemoPageState {
	return &DemoPageState{
		ShowModal:     false,
		ToastManager:  NewToastManager(),
		ConfirmDialog: NewConfirmDialog(),
		LoaderState:   NewLoaderState().SetReady(),
		SamplePayouts: []Payout{
			{
				ID:          "1",
				Amount:      150.00,
				Currency:    "USD",
				Description: "Freelance work payment",
				Recipient:   "john@example.com",
				Status:      "completed",
				Selected:    false,
			},
			{
				ID:          "2",
				Amount:      75.50,
				Currency:    "USD",
				Description: "Refund for cancelled order",
				Recipient:   "jane@example.com",
				Status:      "pending",
				Selected:    false,
			},
			{
				ID:          "3",
				Amount:      300.00,
				Currency:    "USD",
				Description: "Consulting fee",
				Recipient:   "bob@example.com",
				Status:      "completed",
				Selected:    false,
			},
		},
	}
}

func (p *DemoPage) OnMount(ctx app.Context) {
	// Initialize component
}

func (p *DemoPage) Render() app.UI {
	return NewPage(PageOptions{
		Name:        "Prism Demo",
		Description: "Interactive demo of the Prism framework components",
		User:        User{LoggedIn: true, Name: "Demo User"},
	}).
		Content(
			p.renderHeader(),
			p.renderComponentShowcase(),
			p.renderDataTable(),
			p.renderModals(),
			p.state.ToastManager,
		)
}

func (p *DemoPage) renderHeader() app.UI {
	return app.Div().
		Style("padding", "24px").
		Style("border-bottom", "1px solid #e0e0e0").
		Style("margin-bottom", "24px").
		Body(
			app.H1().
				Style("margin", "0 0 8px 0").
				Text("Prism Framework Demo"),
			app.P().
				Style("margin", "0").
				Style("color", "#666").
				Text("Interactive showcase of UI components, state management, and user interactions"),
		)
}

func (p *DemoPage) renderComponentShowcase() app.UI {
	return app.Div().
		Style("padding", "0 24px 24px").
		Body(
			app.H2().Text("Component Showcase"),
			app.Div().
				Style("display", "flex").
				Style("gap", "16px").
				Style("margin-bottom", "24px").
				Body(
					sp.Button().
						Text("Show Info Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ToastManager.ShowInfo("This is an info toast message!")
							ctx.Reload()
						}),
					sp.Button().
						Text("Show Success Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ToastManager.ShowSuccess("Operation completed successfully!")
							ctx.Reload()
						}),
					sp.Button().
						Text("Show Warning Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ToastManager.ShowWarning("Please review your settings")
							ctx.Reload()
						}),
					sp.Button().
						Text("Show Error Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ToastManager.ShowError("An error occurred while processing")
							ctx.Reload()
						}),
				),
			app.Div().
				Style("display", "flex").
				Style("gap", "16px").
				Style("margin-bottom", "24px").
				Body(
					sp.Button().
						Text("Open Modal").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ShowModal = true
							p.state.ModalTitle = "Demo Modal"
							p.state.ModalContent = "This is a sample modal dialog with some content."
							ctx.Reload()
						}),
					sp.Button().
						Text("Show Confirm Dialog").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ConfirmDialog = NewConfirmDialog().
								Title("Confirm Action").
								Message("Are you sure you want to perform this action?").
								OnConfirm(func(ctx app.Context) {
									p.state.ToastManager.ShowSuccess("Action confirmed!")
									ctx.Reload()
								}).
								OnCancel(func(ctx app.Context) {
									p.state.ToastManager.ShowInfo("Action cancelled")
									ctx.Reload()
								}).
								Show()
							ctx.Reload()
						}),
					sp.Button().
						Text("Show Delete Confirm").
						OnClick(func(ctx app.Context, e app.Event) {
							p.state.ConfirmDialog = ShowDeleteConfirm("item", func(ctx app.Context) {
								p.state.ToastManager.ShowSuccess("Item deleted successfully")
								ctx.Reload()
							})
							ctx.Reload()
						}),
				),
		)
}

func (p *DemoPage) renderDataTable() app.UI {
	return app.Div().
		Style("padding", "0 24px 24px").
		Body(
			app.H2().Text("Sample Data Table"),
			app.P().
				Style("color", "#666").
				Style("margin-bottom", "16px").
				Text("Click on rows to select them, then use action buttons."),
			app.Div().
				Style("border", "1px solid #e0e0e0").
				Style("border-radius", "4px").
				Style("overflow", "hidden").
				Body(
					// Table header
					app.Div().
						Style("display", "grid").
						Style("grid-template-columns", "40px 1fr 100px 150px 120px 100px").
						Style("background", "#f5f5f5").
						Style("border-bottom", "1px solid #e0e0e0").
						Style("padding", "12px").
						Style("font-weight", "bold").
						Body(
							app.Div().Text(""),
							app.Div().Text("Description"),
							app.Div().Text("Amount"),
							app.Div().Text("Recipient"),
							app.Div().Text("Status"),
							app.Div().Text("Actions"),
						),
					// Table rows
					app.Range(p.state.SamplePayouts).Slice(func(i int) app.UI {
						payout := p.state.SamplePayouts[i]
						return p.renderPayoutRow(payout, i)
					}),
				),
			app.Div().
				Style("margin-top", "16px").
				Style("display", "flex").
				Style("gap", "16px").
				Body(
					sp.Button().
						Text("Select All").
						OnClick(func(ctx app.Context, e app.Event) {
							for i := range p.state.SamplePayouts {
								p.state.SamplePayouts[i].Selected = true
							}
							ctx.Reload()
						}),
					sp.Button().
						Text("Clear Selection").
						OnClick(func(ctx app.Context, e app.Event) {
							for i := range p.state.SamplePayouts {
								p.state.SamplePayouts[i].Selected = false
							}
							ctx.Reload()
						}),
					sp.Button().
						Text("Process Selected").
						OnClick(func(ctx app.Context, e app.Event) {
							selectedCount := 0
							for _, payout := range p.state.SamplePayouts {
								if payout.Selected {
									selectedCount++
								}
							}
							if selectedCount > 0 {
								p.state.ToastManager.ShowSuccess(fmt.Sprintf("Processing %d payouts", selectedCount))
							} else {
								p.state.ToastManager.ShowWarning("No payouts selected")
							}
							ctx.Reload()
						}),
				),
		)
}

func (p *DemoPage) renderPayoutRow(payout Payout, index int) app.UI {
	return app.Div().
		Style("display", "grid").
		Style("grid-template-columns", "40px 1fr 100px 150px 120px 100px").
		Style("padding", "12px").
		Style("border-bottom", "1px solid #e0e0e0").
		Style("cursor", "pointer").
		Style("background", IfElseString(payout.Selected, "#f0f8ff", "white")).
		OnClick(func(ctx app.Context, e app.Event) {
			p.state.SamplePayouts[index].Selected = !p.state.SamplePayouts[index].Selected
			ctx.Reload()
		}).
		Body(
			app.Input().
				Type("checkbox").
				Checked(payout.Selected).
				OnChange(func(ctx app.Context, e app.Event) {
					p.state.SamplePayouts[index].Selected = !p.state.SamplePayouts[index].Selected
					ctx.Reload()
				}),
			app.Div().Text(payout.Description),
			app.Div().Text(FormatCurrency(payout.Amount, payout.Currency)),
			app.Div().Text(payout.Recipient),
			app.Div().
				Style("padding", "4px 8px").
				Style("border-radius", "12px").
				Style("font-size", "12px").
				Style("background", IfElseString(payout.Status == "completed", "#d4edda", "#fff3cd")).
				Style("color", IfElseString(payout.Status == "completed", "#155724", "#856404")).
				Text(payout.Status),
			app.Div().
				Body(
					sp.Button().
						Text("View").
						OnClick(func(ctx app.Context, e app.Event) {
							e.PreventDefault()
							p.state.SelectedPayout = &payout
							p.state.ShowModal = true
							p.state.ModalTitle = "Payout Details"
							p.state.ModalContent = fmt.Sprintf("Payout ID: %s\nAmount: %s\nRecipient: %s\nStatus: %s",
								payout.ID, FormatCurrency(payout.Amount, payout.Currency), payout.Recipient, payout.Status)
							ctx.Reload()
						}),
				),
		)
}

func (p *DemoPage) renderModals() app.UI {
	return app.Div().Body(
		NewActionModal().
			Title(p.state.ModalTitle).
			Body(app.P().Text(p.state.ModalContent)).
			Visible(p.state.ShowModal).
			OnDismiss(func(ctx app.Context) {
				p.state.ShowModal = false
				ctx.Reload()
			}).
			OnCancel(func(ctx app.Context) {
				p.state.ShowModal = false
				ctx.Reload()
			}),
		p.state.ConfirmDialog,
	)
}
