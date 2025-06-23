package prism

import (
	"time"

	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ComprehensiveTestPage showcases all Spectrum components for E2E testing
type ComprehensiveTestPage struct {
	app.Compo

	toastManager *ToastManager

	// State for various components
	checkboxValue  bool
	radioValue     string
	toggleValue    bool
	sliderValue    int
	textAreaValue  string
	selectValue    string
	switchValue    bool
	isModalOpen    bool
	toastMessage   string
	toastType      string
	tabActiveIndex int
	accordionOpen  bool
	tableData      []*TestUser
	selectedRows   map[int]bool
	formData       *FormData

	// Test data
	testUsers []*TestUser
}

type TestUser struct {
	ID     int
	Name   string
	Email  string
	Role   string
	Active bool
	Score  int
}

func NewComprehensiveTestPage() *ComprehensiveTestPage {
	testUsers := []*TestUser{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", Role: "Admin", Active: true, Score: 95},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com", Role: "User", Active: true, Score: 82},
		{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com", Role: "Moderator", Active: false, Score: 78},
		{ID: 4, Name: "Diana Prince", Email: "diana@example.com", Role: "User", Active: true, Score: 91},
		{ID: 5, Name: "Eve Wilson", Email: "eve@example.com", Role: "Admin", Active: false, Score: 87},
	}

	formData := NewFormData().
		AddField(&FormField{
			Name:       "testName",
			Label:      "Full Name",
			Type:       "text",
			Required:   true,
			Validators: []Validator{MinLengthValidator(2)},
		}).
		AddField(&FormField{
			Name:       "testEmail",
			Label:      "Email Address",
			Type:       "email",
			Required:   true,
			Validators: []Validator{EmailValidator()},
		}).
		AddField(&FormField{
			Name:     "testPhone",
			Label:    "Phone Number",
			Type:     "tel",
			Required: false,
		})

	return &ComprehensiveTestPage{
		radioValue:   "option1",
		sliderValue:  50,
		selectValue:  "option1",
		selectedRows: make(map[int]bool),
		testUsers:    testUsers,
		tableData:    testUsers,
		formData:     formData,
		toastManager: NewToastManager(),
	}
}

func (c *ComprehensiveTestPage) Render() app.UI {
	return NewPage().
		WithName("Comprehensive Component Test").
		WithUser(User{LoggedIn: true, Name: "Test User"}).
		Content(
			app.Div().
				Class("comprehensive-test-page").
				Style("padding", "24px").
				Style("max-width", "1200px").
				Style("margin", "0 auto").
				Body(
					c.toastManager,
					c.renderHeader(),
					c.renderButtonSection(),
					c.renderFormSection(),
					c.renderDataDisplaySection(),
					c.renderLayoutSection(),
					c.renderInteractionSection(),
					c.renderNavigationSection(),
					c.renderAdvancedSection(),
				),
		)
}

func (c *ComprehensiveTestPage) renderHeader() app.UI {
	return app.Div().
		Class("test-header").
		Style("margin-bottom", "32px").
		Body(
			app.H1().
				Style("margin", "0 0 16px 0").
				Style("color", "#333").
				Text("Comprehensive Component Test Suite"),
			app.P().
				Style("color", "#666").
				Style("font-size", "16px").
				Text("This page demonstrates all Spectrum components for comprehensive E2E testing."),
		)
}

func (c *ComprehensiveTestPage) renderButtonSection() app.UI {
	return app.Div().
		Class("test-section buttons-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("buttons-section").
				Text("Buttons & Actions"),
			app.Div().
				Class("component-grid").
				Style("display", "grid").
				Style("grid-template-columns", "repeat(auto-fit, minmax(300px, 1fr))").
				Style("gap", "24px").
				Body(
					c.renderBasicButtons(),
					c.renderButtonVariants(),
					c.renderButtonStates(),
					c.renderIconButtons(),
				),
		)
}

func (c *ComprehensiveTestPage) renderBasicButtons() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Basic Buttons"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.Button().
						Id("primary-button").
						Text("Primary Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Primary button clicked!", "info")
						}),
					sp.Button().
						Id("secondary-button").
						Text("Secondary Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Secondary button clicked!", "info")
						}),
					sp.Button().
						Id("action-button").
						Text("Action Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Action performed!", "success")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderButtonVariants() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Button Variants"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.Button().
						Id("cta-button").
						Text("Call to Action").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("CTA clicked!", "success")
						}),
					sp.Button().
						Id("outline-button").
						Text("Outline Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Outline clicked!", "info")
						}),
					sp.Button().
						Id("quiet-button").
						Text("Quiet Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Quiet clicked!", "info")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderButtonStates() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Button States"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.Button().
						Id("disabled-button").
						Text("Disabled Button").
						Disabled(true),
					sp.Button().
						Id("loading-button").
						Text("Loading Button").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Loading simulation!", "info")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderIconButtons() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Icon Buttons"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.Button().
						Id("edit-icon-button").
						Text("✏️ Edit").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Edit action!", "info")
						}),
					sp.Button().
						Id("delete-icon-button").
						Text("🗑️ Delete").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Delete action!", "warning")
						}),
					sp.Button().
						Id("save-icon-button").
						Text("💾 Save").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Save action!", "success")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderFormSection() app.UI {
	return app.Div().
		Class("test-section forms-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("forms-section").
				Text("Form Components"),
			app.Div().
				Class("component-grid").
				Style("display", "grid").
				Style("grid-template-columns", "repeat(auto-fit, minmax(300px, 1fr))").
				Style("gap", "24px").
				Body(
					c.renderInputFields(),
					c.renderSelectionControls(),
					c.renderFormValidation(),
				),
		)
}

func (c *ComprehensiveTestPage) renderInputFields() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Input Fields"),
			app.Div().
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "16px").
				Body(
					sp.Textfield().
						Id("text-input").
						Placeholder("Enter text here").
						OnInput(func(ctx app.Context, e app.Event) {
							c.showToast("Text input changed!", "info")
						}),
					sp.Textfield().
						Id("password-input").
						Type("password").
						Placeholder("Enter password").
						OnInput(func(ctx app.Context, e app.Event) {
							c.showToast("Password input changed!", "info")
						}),
					sp.Textfield().
						Id("email-input").
						Type("email").
						Placeholder("Enter email").
						OnInput(func(ctx app.Context, e app.Event) {
							c.showToast("Email input changed!", "info")
						}),
					sp.Textfield().
						Id("number-input").
						Type("number").
						Placeholder("Enter number").
						OnInput(func(ctx app.Context, e app.Event) {
							c.showToast("Number input changed!", "info")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderSelectionControls() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Selection Controls"),
			app.Div().
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "16px").
				Body(
					// Checkbox
					app.Label().
						Style("display", "flex").
						Style("align-items", "center").
						Style("gap", "8px").
						Body(
							sp.Checkbox().
								Id("test-checkbox").
								Checked(c.checkboxValue).
								OnChange(func(ctx app.Context, e app.Event) {
									c.checkboxValue = !c.checkboxValue
									c.showToast("Checkbox toggled!", "info")
								}),
							app.Text("Test Checkbox"),
						),

					// Radio Group
					app.Div().
						Body(
							app.Text("Radio Group:"),
							sp.RadioGroup().
								Id("test-radio-group").
								Body(
									sp.Radio().
										Id("radio-option1").
										Value("option1").
										Checked(c.radioValue == "option1").
										OnChange(func(ctx app.Context, e app.Event) {
											c.radioValue = "option1"
											c.showToast("Radio option 1 selected!", "info")
										}),
									app.Text("Option 1"),
									sp.Radio().
										Id("radio-option2").
										Value("option2").
										Checked(c.radioValue == "option2").
										OnChange(func(ctx app.Context, e app.Event) {
											c.radioValue = "option2"
											c.showToast("Radio option 2 selected!", "info")
										}),
									app.Text("Option 2"),
								),
						),

					// Switch
					app.Label().
						Style("display", "flex").
						Style("align-items", "center").
						Style("gap", "8px").
						Body(
							sp.Switch().
								Id("test-switch").
								Checked(c.switchValue).
								OnChange(func(ctx app.Context, e app.Event) {
									c.switchValue = !c.switchValue
									c.showToast("Switch toggled!", "info")
								}),
							app.Text("Test Switch"),
						),
				),
		)
}

func (c *ComprehensiveTestPage) renderFormValidation() app.UI {
	form := NewForm().
		WithFormData(c.formData).
		WithTitle("Validation Form").
		WithOnSubmit(func(ctx app.Context, values map[string]string) error {
			if c.formData.ValidateAll() {
				c.showToast("Form validation passed!", "success")
			} else {
				c.showToast("Form validation failed!", "error")
			}
			return nil
		})

	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Form Validation"),
			form,
		)
}

func (c *ComprehensiveTestPage) renderDataDisplaySection() app.UI {
	columns := []TableColumn[*TestUser]{
		{
			Key:   "id",
			Label: "ID",
			Render: func(user *TestUser) app.UI {
				return app.Text(FormatNumber(user.ID))
			},
		},
		{
			Key:   "name",
			Label: "Name",
			Render: func(user *TestUser) app.UI {
				return app.Text(user.Name)
			},
		},
		{
			Key:   "email",
			Label: "Email",
			Render: func(user *TestUser) app.UI {
				return app.Text(user.Email)
			},
		},
		{
			Key:   "role",
			Label: "Role",
			Render: func(user *TestUser) app.UI {
				return app.Text(user.Role)
			},
		},
		{
			Key:   "active",
			Label: "Status",
			Render: func(user *TestUser) app.UI {
				return app.Text(IfElseString(user.Active, "Active", "Inactive"))
			},
		},
		{
			Key:   "actions",
			Label: "Actions",
			Render: func(user *TestUser) app.UI {
				return app.Div().
					Style("display", "flex").
					Style("gap", "8px").
					Body(
						sp.Button().
							Text("View").
							OnClick(func(ctx app.Context, e app.Event) {
								c.showToast("Viewing user: "+user.Name, "info")
							}),
						sp.Button().
							Text("Edit").
							OnClick(func(ctx app.Context, e app.Event) {
								c.showToast("Editing user: "+user.Name, "info")
							}),
					)
			},
		},
	}

	table := NewDataTable[*TestUser]().
		WithData(c.tableData).
		WithColumns(columns...).
		WithSelectable(true).
		WithShowFilters(true).
		WithPageSize(10)

	return app.Div().
		Class("test-section data-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("data-section").
				Text("Data Display"),
			app.Div().
				Class("component-demo").
				Body(
					app.H3().Text("Data Table"),
					app.Div().
						Style("margin-bottom", "16px").
						Body(
							sp.Button().
								Id("select-all-button").
								Text("Select All").
								OnClick(func(ctx app.Context, e app.Event) {
									table.SelectAll()
									c.showToast("All rows selected!", "info")
								}),
							sp.Button().
								Id("clear-selection-button").
								Text("Clear Selection").
								Style("margin-left", "8px").
								OnClick(func(ctx app.Context, e app.Event) {
									table.ClearSelection()
									c.showToast("Selection cleared!", "info")
								}),
						),
					table,
				),
			c.renderCards(),
		)
}

func (c *ComprehensiveTestPage) renderCards() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Cards"),
			app.Div().
				Style("display", "grid").
				Style("grid-template-columns", "repeat(auto-fit, minmax(250px, 1fr))").
				Style("gap", "16px").
				Body(
					NewCard().
						WithTitle("Basic Card").
						WithSubtitle("Simple card example").
						WithContent(
							app.P().Text("This is a basic card with title, subtitle, and content."),
						),

					NewCard().
						WithTitle("Action Card").
						WithSubtitle("Card with actions").
						WithContent(
							app.P().Text("This card has action buttons."),
						).
						WithActions(
							NewCardAction("view", "View").
								WithOnClick(func(ctx app.Context) {
									c.showToast("Card action: View", "info")
								}),
							NewCardAction("edit", "Edit").
								WithOnClick(func(ctx app.Context) {
									c.showToast("Card action: Edit", "info")
								}),
						),

					NewCard().
						WithTitle("Image Card").
						WithSubtitle("Card with image").
						WithContent(
							app.P().Text("This card could have an image."),
						).
						WithVariant("outlined"),
				),
		)
}

func (c *ComprehensiveTestPage) renderLayoutSection() app.UI {
	return app.Div().
		Class("test-section layout-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("layout-section").
				Text("Layout Components"),
			c.renderTabs(),
			c.renderAccordion(),
		)
}

func (c *ComprehensiveTestPage) renderTabs() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Tabs"),
			app.Div().
				ID("test-tabs").
				Body(
					// Tab Headers
					app.Div().
						Class("tab-headers").
						Style("display", "flex").
						Style("border-bottom", "1px solid #e1e1e1").
						Style("margin-bottom", "16px").
						Body(
							sp.Button().
								Id("tab-1").
								Text("Tab 1").
								Style("border", "none").
								Style("background", IfElseString(c.tabActiveIndex == 0, "#f0f0f0", "transparent")).
								OnClick(func(ctx app.Context, e app.Event) {
									c.tabActiveIndex = 0
									c.showToast("Tab 1 selected!", "info")
								}),
							sp.Button().
								Id("tab-2").
								Text("Tab 2").
								Style("border", "none").
								Style("background", IfElseString(c.tabActiveIndex == 1, "#f0f0f0", "transparent")).
								OnClick(func(ctx app.Context, e app.Event) {
									c.tabActiveIndex = 1
									c.showToast("Tab 2 selected!", "info")
								}),
							sp.Button().
								Id("tab-3").
								Text("Tab 3").
								Style("border", "none").
								Style("background", IfElseString(c.tabActiveIndex == 2, "#f0f0f0", "transparent")).
								OnClick(func(ctx app.Context, e app.Event) {
									c.tabActiveIndex = 2
									c.showToast("Tab 3 selected!", "info")
								}),
						),
					// Tab Content
					app.Div().
						Class("tab-content").
						Body(
							app.If(c.tabActiveIndex == 0, func() app.UI {
								return app.Div().ID("tab-content-1").Text("Content for Tab 1")
							}),
							app.If(c.tabActiveIndex == 1, func() app.UI {
								return app.Div().ID("tab-content-2").Text("Content for Tab 2")
							}),
							app.If(c.tabActiveIndex == 2, func() app.UI {
								return app.Div().ID("tab-content-3").Text("Content for Tab 3")
							}),
						),
				),
		)
}

func (c *ComprehensiveTestPage) renderAccordion() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Accordion"),
			app.Div().
				ID("test-accordion").
				Body(
					app.Div().
						Class("accordion-item").
						Style("border", "1px solid #e1e1e1").
						Style("margin-bottom", "8px").
						Body(
							sp.Button().
								Id("accordion-toggle").
								Text(IfElseString(c.accordionOpen, "▼ Accordion Item (Click to close)", "▶ Accordion Item (Click to open)")).
								Style("width", "100%").
								Style("text-align", "left").
								Style("padding", "12px").
								OnClick(func(ctx app.Context, e app.Event) {
									c.accordionOpen = !c.accordionOpen
									c.showToast("Accordion toggled!", "info")
								}),
							app.If(c.accordionOpen, func() app.UI {
								return app.Div().
									ID("accordion-content").
									Style("padding", "16px").
									Style("background", "#f9f9f9").
									Text("This is the accordion content that shows when expanded.")
							}),
						),
				),
		)
}

func (c *ComprehensiveTestPage) renderInteractionSection() app.UI {
	return app.Div().
		Class("test-section interaction-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("interaction-section").
				Text("Interaction Components"),
			c.renderModalExample(),
			c.renderToastExample(),
			c.renderTooltipExample(),
		)
}

func (c *ComprehensiveTestPage) renderModalExample() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Modal"),
			sp.Button().
				Id("open-modal-button").
				Text("Open Modal").
				OnClick(func(ctx app.Context, e app.Event) {
					c.isModalOpen = true
				}),
			NewActionModal().
				Title("Test Modal").
				Visible(c.isModalOpen).
				Body(
					app.P().Text("This is a test modal with content."),
				).
				OnCancel(func(ctx app.Context) {
					c.isModalOpen = false
					c.showToast("Modal cancelled!", "info")
				}).
				OnConfirm(func(ctx app.Context) {
					c.isModalOpen = false
					c.showToast("Modal confirmed!", "success")
				}).
				OnDismiss(func(ctx app.Context) {
					c.isModalOpen = false
				}),
		)
}

func (c *ComprehensiveTestPage) renderToastExample() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Toast Notifications"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.Button().
						Id("toast-info-button").
						Text("Show Info Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("This is an info toast!", "info")
						}),
					sp.Button().
						Id("toast-success-button").
						Text("Show Success Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("This is a success toast!", "success")
						}),
					sp.Button().
						Id("toast-warning-button").
						Text("Show Warning Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("This is a warning toast!", "warning")
						}),
					sp.Button().
						Id("toast-error-button").
						Text("Show Error Toast").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("This is an error toast!", "error")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderTooltipExample() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Tooltips"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					sp.OverlayTrigger().
						Trigger(
							sp.Button().
								Id("tooltip-button-1").
								Text("Hover for Tooltip").
								OnClick(func(ctx app.Context, e app.Event) {
									c.showToast("Button with tooltip clicked!", "info")
								}),
						).
						HoverContent(
							sp.Tooltip().Body(app.Text("This is a tooltip message")),
						),
					sp.OverlayTrigger().
						Trigger(
							sp.Button().
								Id("tooltip-button-2").
								Text("Another Tooltip").
								OnClick(func(ctx app.Context, e app.Event) {
									c.showToast("Another tooltip button clicked!", "info")
								}),
						).
						HoverContent(
							sp.Tooltip().Body(app.Text("Another tooltip message")),
						),
				),
		)
}

func (c *ComprehensiveTestPage) renderNavigationSection() app.UI {
	return app.Div().
		Class("test-section navigation-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("navigation-section").
				Text("Navigation Components"),
			c.renderBreadcrumbs(),
			c.renderPagination(),
			c.renderMenu(),
			c.renderSideNavExample(),
			c.renderTopNavExample(),
		)
}

func (c *ComprehensiveTestPage) renderSideNavExample() app.UI {
	sideNav := NewSideNav().
		WithTitle("Test Navigation").
		WithItems(
			NewSideNavItem("nav-home", "Home").
				WithIcon("🏠").
				WithActive(true).
				WithOnClick(func(ctx app.Context) {
					c.showToast("Home navigation clicked!", "info")
				}),
			NewSideNavItem("nav-users", "Users").
				WithIcon("👥").
				WithBadge("5").
				WithOnClick(func(ctx app.Context) {
					c.showToast("Users navigation clicked!", "info")
				}),
			NewSideNavItem("nav-settings", "Settings").
				WithIcon("⚙️").
				WithOnClick(func(ctx app.Context) {
					c.showToast("Settings navigation clicked!", "info")
				}),
		)

	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Side Navigation"),
			app.Div().
				ID("test-sidenav").
				Style("height", "300px").
				Style("border", "1px solid #e1e1e1").
				Style("border-radius", "4px").
				Body(sideNav),
		)
}

func (c *ComprehensiveTestPage) renderTopNavExample() app.UI {
	userMenu := NewUserMenu().
		WithUser(&User{
			ID:    "test-user",
			Name:  "Test User",
			Email: "test@example.com",
		}).
		WithMenuItems(
			NewUserMenuItem("Profile").
				WithOnClick(func(ctx app.Context) {
					c.showToast("Profile clicked!", "info")
				}),
			NewUserMenuItem("Logout").
				WithOnClick(func(ctx app.Context) {
					c.showToast("Logout clicked!", "info")
				}),
		)

	topNav := NewTopNav().
		WithTitle("Test App").
		WithUserMenu(userMenu).
		WithSearchBar(
			NewSearchBar().
				WithPlaceholder("Search...").
				WithOnSearch(func(query string) {
					c.showToast("Searching for: "+query, "info")
				}),
		)

	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Top Navigation"),
			app.Div().
				ID("test-topnav").
				Style("border", "1px solid #e1e1e1").
				Style("border-radius", "4px").
				Body(topNav),
		)
}

func (c *ComprehensiveTestPage) renderBreadcrumbs() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Breadcrumbs"),
			app.Nav().
				ID("test-breadcrumbs").
				Style("display", "flex").
				Style("align-items", "center").
				Style("gap", "8px").
				Body(
					app.A().
						Href("/").
						Text("Home").
						OnClick(func(ctx app.Context, e app.Event) {
							e.PreventDefault()
							c.showToast("Home breadcrumb clicked!", "info")
						}),
					app.Text(" > "),
					app.A().
						Href("/test").
						Text("Test").
						OnClick(func(ctx app.Context, e app.Event) {
							e.PreventDefault()
							c.showToast("Test breadcrumb clicked!", "info")
						}),
					app.Text(" > "),
					app.Text("Current Page"),
				),
		)
}

func (c *ComprehensiveTestPage) renderPagination() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Pagination"),
			app.Div().
				ID("test-pagination").
				Style("display", "flex").
				Style("align-items", "center").
				Style("gap", "8px").
				Body(
					sp.Button().
						Id("prev-page").
						Text("Previous").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Previous page clicked!", "info")
						}),
					sp.Button().
						Id("page-1").
						Text("1").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Page 1 clicked!", "info")
						}),
					sp.Button().
						Id("page-2").
						Text("2").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Page 2 clicked!", "info")
						}),
					sp.Button().
						Id("page-3").
						Text("3").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Page 3 clicked!", "info")
						}),
					sp.Button().
						Id("next-page").
						Text("Next").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Next page clicked!", "info")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderMenu() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Menu"),
			app.Div().
				ID("test-menu").
				Style("border", "1px solid #e1e1e1").
				Style("border-radius", "4px").
				Style("padding", "8px 0").
				Style("max-width", "200px").
				Body(
					app.Div().
						Class("menu-item").
						Style("padding", "8px 16px").
						Style("cursor", "pointer").
						Text("Menu Item 1").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Menu Item 1 clicked!", "info")
						}),
					app.Div().
						Class("menu-item").
						Style("padding", "8px 16px").
						Style("cursor", "pointer").
						Text("Menu Item 2").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Menu Item 2 clicked!", "info")
						}),
					app.Div().
						Class("menu-item").
						Style("padding", "8px 16px").
						Style("cursor", "pointer").
						Text("Menu Item 3").
						OnClick(func(ctx app.Context, e app.Event) {
							c.showToast("Menu Item 3 clicked!", "info")
						}),
				),
		)
}

func (c *ComprehensiveTestPage) renderAdvancedSection() app.UI {
	return app.Div().
		Class("test-section advanced-section").
		Style("margin-bottom", "48px").
		Body(
			app.H2().
				ID("advanced-section").
				Text("Advanced Components"),
			c.renderProgressBars(),
			c.renderBadges(),
			c.renderChips(),
		)
}

func (c *ComprehensiveTestPage) renderProgressBars() app.UI {
	return app.Div().
		Class("component-demo").
		Body(
			app.H3().Text("Progress Indicators"),
			app.Div().
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "16px").
				Body(
					app.Div().
						Body(
							app.Text("Progress Bar (50%):"),
							app.Div().
								ID("progress-bar-1").
								Style("width", "100%").
								Style("height", "8px").
								Style("background", "#e1e1e1").
								Style("border-radius", "4px").
								Style("overflow", "hidden").
								Body(
									app.Div().
										Style("width", "50%").
										Style("height", "100%").
										Style("background", "#0066cc").
										Style("transition", "width 0.3s ease"),
								),
						),
					app.Div().
						Body(
							app.Text("Progress Bar (75%):"),
							app.Div().
								ID("progress-bar-2").
								Style("width", "100%").
								Style("height", "8px").
								Style("background", "#e1e1e1").
								Style("border-radius", "4px").
								Style("overflow", "hidden").
								Body(
									app.Div().
										Style("width", "75%").
										Style("height", "100%").
										Style("background", "#28a745").
										Style("transition", "width 0.3s ease"),
								),
						),
				),
		)
}

func (c *ComprehensiveTestPage) renderBadges() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Badges"),
			app.Div().
				Style("display", "flex").
				Style("gap", "12px").
				Style("flex-wrap", "wrap").
				Body(
					app.Span().
						ID("badge-primary").
						Class("badge badge-primary").
						Style("background", "#0066cc").
						Style("color", "white").
						Style("padding", "4px 8px").
						Style("border-radius", "12px").
						Style("font-size", "12px").
						Text("Primary"),
					app.Span().
						ID("badge-success").
						Class("badge badge-success").
						Style("background", "#28a745").
						Style("color", "white").
						Style("padding", "4px 8px").
						Style("border-radius", "12px").
						Style("font-size", "12px").
						Text("Success"),
					app.Span().
						ID("badge-warning").
						Class("badge badge-warning").
						Style("background", "#ffc107").
						Style("color", "black").
						Style("padding", "4px 8px").
						Style("border-radius", "12px").
						Style("font-size", "12px").
						Text("Warning"),
					app.Span().
						ID("badge-error").
						Class("badge badge-error").
						Style("background", "#dc3545").
						Style("color", "white").
						Style("padding", "4px 8px").
						Style("border-radius", "12px").
						Style("font-size", "12px").
						Text("Error"),
				),
		)
}

func (c *ComprehensiveTestPage) renderChips() app.UI {
	return app.Div().
		Class("component-demo").
		Style("margin-top", "32px").
		Body(
			app.H3().Text("Chips"),
			app.Div().
				Style("display", "flex").
				Style("gap", "8px").
				Style("flex-wrap", "wrap").
				Body(
					app.Div().
						ID("chip-1").
						Class("chip").
						Style("display", "flex").
						Style("align-items", "center").
						Style("gap", "4px").
						Style("background", "#f0f0f0").
						Style("border", "1px solid #ccc").
						Style("border-radius", "16px").
						Style("padding", "4px 12px").
						Style("font-size", "14px").
						Body(
							app.Text("Tag 1"),
							app.Span().
								Style("cursor", "pointer").
								Style("margin-left", "4px").
								Text("×").
								OnClick(func(ctx app.Context, e app.Event) {
									c.showToast("Chip 1 removed!", "info")
								}),
						),
					app.Div().
						ID("chip-2").
						Class("chip").
						Style("display", "flex").
						Style("align-items", "center").
						Style("gap", "4px").
						Style("background", "#f0f0f0").
						Style("border", "1px solid #ccc").
						Style("border-radius", "16px").
						Style("padding", "4px 12px").
						Style("font-size", "14px").
						Body(
							app.Text("Tag 2"),
							app.Span().
								Style("cursor", "pointer").
								Style("margin-left", "4px").
								Text("×").
								OnClick(func(ctx app.Context, e app.Event) {
									c.showToast("Chip 2 removed!", "info")
								}),
						),
				),
		)
}

func (c *ComprehensiveTestPage) showToast(message, toastType string) {
	c.toastMessage = message
	c.toastType = toastType

	var variant ToastVariant
	switch toastType {
	case "success":
		variant = ToastVariantSuccess
	case "warning":
		variant = ToastVariantWarning
	case "error":
		variant = ToastVariantError
	default:
		variant = ToastVariantInfo
	}

	c.toastManager.Show(message, variant, 5*time.Second)
}
