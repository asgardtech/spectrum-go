package prism

import (
	"fmt"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ComprehensiveDemo demonstrates all the implemented components working together
type ComprehensiveDemo struct {
	app.Compo

	// Form state
	formData *FormData
	
	// Table state  
	tableData    []*DemoUser
	userTable    *DataTable[*DemoUser]
	
	// UI state
	showModal    bool
	selectedUser *DemoUser
	notification string
}

// DemoUser represents a user for our demo table
type DemoUser struct {
	ID       int
	Name     string
	Email    string
	Role     string
	Status   string
	Created  time.Time
	LastSeen time.Time
}

// Generate demo data
func generateDemoUsers() []*DemoUser {
	users := []*DemoUser{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Role: "Admin", Status: "Active", Created: time.Now().AddDate(0, -2, 0), LastSeen: time.Now().Add(-1 * time.Hour)},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Role: "User", Status: "Active", Created: time.Now().AddDate(0, -1, -15), LastSeen: time.Now().Add(-30 * time.Minute)},
		{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", Role: "Manager", Status: "Inactive", Created: time.Now().AddDate(0, -3, -5), LastSeen: time.Now().AddDate(0, 0, -7)},
		{ID: 4, Name: "Alice Brown", Email: "alice@example.com", Role: "User", Status: "Active", Created: time.Now().AddDate(0, -1, 0), LastSeen: time.Now().Add(-10 * time.Minute)},
		{ID: 5, Name: "Charlie Wilson", Email: "charlie@example.com", Role: "User", Status: "Pending", Created: time.Now().AddDate(0, 0, -1), LastSeen: time.Time{}},
	}
	return users
}

func NewComprehensiveDemo() *ComprehensiveDemo {
	demo := &ComprehensiveDemo{
		tableData: generateDemoUsers(),
	}
	
	// Initialize form data
	demo.formData = NewFormData().
		AddField(&FormField{
			Name:        "name",
			Label:       "Full Name",
			Type:        "text",
			Required:    true,
			Placeholder: "Enter full name",
			Validators:  []Validator{MinLengthValidator(2)},
		}).
		AddField(&FormField{
			Name:        "email",
			Label:       "Email Address",
			Type:        "email",
			Required:    true,
			Placeholder: "Enter email address",
			Validators:  []Validator{EmailValidator()},
		}).
		AddField(&FormField{
			Name:        "role",
			Label:       "Role",
			Type:        "text",
			Required:    true,
			Placeholder: "Enter role",
		}).
		AddField(&FormField{
			Name:     "bio",
			Label:    "Biography",
			Type:     "textarea",
			Required: false,
			HelpText: "Optional biography or notes",
		})
	
	// Initialize data table
	demo.userTable = NewDataTable[*DemoUser]().
		WithColumns(
			TableColumn[*DemoUser]{
				Key:      "id",
				Label:    "ID",
				Width:    "60px",
				Sortable: true,
				Render: func(user *DemoUser) app.UI {
					return app.Text(fmt.Sprintf("#%d", user.ID))
				},
				Compare: func(a, b *DemoUser) int {
					return a.ID - b.ID
				},
			},
			TableColumn[*DemoUser]{
				Key:        "name",
				Label:      "Name",
				Width:      "200px",
				Sortable:   true,
				Filterable: true,
				Render: func(user *DemoUser) app.UI {
					return app.Div().
						Style("font-weight", "600").
						Text(user.Name)
				},
				Compare: func(a, b *DemoUser) int {
					if a.Name < b.Name {
						return -1
					} else if a.Name > b.Name {
						return 1
					}
					return 0
				},
			},
			TableColumn[*DemoUser]{
				Key:        "email",
				Label:      "Email",
				Width:      "250px",
				Sortable:   true,
				Filterable: true,
				Render: func(user *DemoUser) app.UI {
					return app.A().
						Href("mailto:"+user.Email).
						Style("color", "#0066cc").
						Text(user.Email)
				},
			},
			TableColumn[*DemoUser]{
				Key:        "role",
				Label:      "Role",
				Width:      "120px",
				Sortable:   true,
				Filterable: true,
				Render: func(user *DemoUser) app.UI {
					color := "#666"
					if user.Role == "Admin" {
						color = "#d73a49"
					} else if user.Role == "Manager" {
						color = "#0066cc"
					}
					return app.Span().
						Style("color", color).
						Style("font-weight", "600").
						Text(user.Role)
				},
			},
			TableColumn[*DemoUser]{
				Key:      "status",
				Label:    "Status",
				Width:    "100px",
				Sortable: true,
				Render: func(user *DemoUser) app.UI {
					var color, bgColor string
					switch user.Status {
					case "Active":
						color = "#22863a"
						bgColor = "#dcffe4"
					case "Inactive":
						color = "#d73a49"
						bgColor = "#ffeaea"
					case "Pending":
						color = "#b08800"
						bgColor = "#fff5b4"
					default:
						color = "#666"
						bgColor = "#f6f8fa"
					}
					return app.Span().
						Style("color", color).
						Style("background-color", bgColor).
						Style("padding", "2px 8px").
						Style("border-radius", "12px").
						Style("font-size", "12px").
						Style("font-weight", "600").
						Text(user.Status)
				},
			},
			TableColumn[*DemoUser]{
				Key:      "created",
				Label:    "Created",
				Width:    "150px",
				Sortable: true,
				Render: func(user *DemoUser) app.UI {
					return app.Div().
						Style("font-size", "14px").
						Body(
							app.Div().Text(FormatDate(user.Created)),
							app.Div().
								Style("color", "#666").
								Style("font-size", "12px").
								Text(FormatTimeAgo(user.Created)),
						)
				},
				Compare: func(a, b *DemoUser) int {
					if a.Created.Before(b.Created) {
						return -1
					} else if a.Created.After(b.Created) {
						return 1
					}
					return 0
				},
			},
			TableColumn[*DemoUser]{
				Key:   "actions",
				Label: "Actions",
				Width: "120px",
				Render: func(user *DemoUser) app.UI {
					return app.Div().
						Style("display", "flex").
						Style("gap", "8px").
						Body(
							app.Button().
								Style("padding", "4px 8px").
								Style("font-size", "12px").
								Style("border", "1px solid #0066cc").
								Style("background", "white").
								Style("color", "#0066cc").
								Style("cursor", "pointer").
								Text("Edit").
								OnClick(func(ctx app.Context, e app.Event) {
									demo.editUser(ctx, user)
								}),
							app.Button().
								Style("padding", "4px 8px").
								Style("font-size", "12px").
								Style("border", "1px solid #d73a49").
								Style("background", "white").
								Style("color", "#d73a49").
								Style("cursor", "pointer").
								Text("Delete").
								OnClick(func(ctx app.Context, e app.Event) {
									demo.deleteUser(ctx, user)
								}),
						)
				},
			},
		).
		WithData(demo.tableData).
		WithSelectable(true).
		WithOnRowClick(func(user *DemoUser, index int) {
			demo.selectedUser = user
			demo.showModal = true
		}).
		WithOnSelectionChange(func(indices []int) {
			demo.notification = fmt.Sprintf("Selected %d users", len(indices))
		})
	
	return demo
}

func (dp *ComprehensiveDemo) editUser(ctx app.Context, user *DemoUser) {
	dp.selectedUser = user
	
	// Pre-populate form with user data
	dp.formData.SetFieldValue("name", user.Name).
		SetFieldValue("email", user.Email).
		SetFieldValue("role", user.Role)
	
	dp.showModal = true
	ctx.Reload()
}

func (dp *ComprehensiveDemo) deleteUser(ctx app.Context, user *DemoUser) {
	// Remove user from table data
	for i, u := range dp.tableData {
		if u.ID == user.ID {
			dp.tableData = append(dp.tableData[:i], dp.tableData[i+1:]...)
			break
		}
	}
	
	dp.userTable.WithData(dp.tableData)
	dp.notification = fmt.Sprintf("Deleted user: %s", user.Name)
	ctx.Reload()
}

func (dp *ComprehensiveDemo) handleFormSubmit(ctx app.Context, values map[string]string) error {
	// Simulate form processing
	if dp.selectedUser != nil {
		// Update existing user
		dp.selectedUser.Name = values["name"]
		dp.selectedUser.Email = values["email"]
		dp.selectedUser.Role = values["role"]
		dp.notification = fmt.Sprintf("Updated user: %s", dp.selectedUser.Name)
	} else {
		// Create new user
		newUser := &DemoUser{
			ID:      len(dp.tableData) + 1,
			Name:    values["name"],
			Email:   values["email"],
			Role:    values["role"],
			Status:  "Active",
			Created: time.Now(),
			LastSeen: time.Now(),
		}
		dp.tableData = append(dp.tableData, newUser)
		dp.notification = fmt.Sprintf("Created user: %s", newUser.Name)
	}
	
	dp.userTable.WithData(dp.tableData)
	dp.showModal = false
	dp.selectedUser = nil
	
	// Clear form
	for name := range dp.formData.Fields {
		dp.formData.SetFieldValue(name, "")
	}
	
	ctx.Reload()
	return nil
}

func (dp *ComprehensiveDemo) Render() app.UI {
	return app.Div().
		Class("demo-page").
		Style("padding", "24px").
		Style("max-width", "1200px").
		Style("margin", "0 auto").
		Body(
			app.H1().
				Style("margin-bottom", "32px").
				Text("Prism UI Framework Demo"),
			
			app.If(dp.notification != "", func() app.UI {
				return app.Div().
					Style("background", "#dcffe4").
					Style("color", "#22863a").
					Style("padding", "12px").
					Style("border-radius", "4px").
					Style("margin-bottom", "24px").
					Style("border", "1px solid #34d058").
					Text(dp.notification).
					OnClick(func(ctx app.Context, e app.Event) {
						dp.notification = ""
						ctx.Reload()
					})
			}),
			
			app.Div().
				Style("display", "flex").
				Style("justify-content", "space-between").
				Style("align-items", "center").
				Style("margin-bottom", "24px").
				Body(
					app.H2().
						Style("margin", "0").
						Text("User Management"),
					app.Button().
						Style("padding", "8px 16px").
						Style("background", "#0066cc").
						Style("color", "white").
						Style("border", "none").
						Style("border-radius", "4px").
						Style("cursor", "pointer").
						Text("Add User").
						OnClick(func(ctx app.Context, e app.Event) {
							dp.selectedUser = nil
							dp.showModal = true
							// Clear form
							for name := range dp.formData.Fields {
								dp.formData.SetFieldValue(name, "")
							}
							ctx.Reload()
						}),
				),
			
			dp.userTable,
			
			app.If(dp.showModal, func() app.UI {
				return dp.renderModal()
			}),
		)
}

func (dp *ComprehensiveDemo) renderModal() app.UI {
	title := "Add User"
	if dp.selectedUser != nil {
		title = "Edit User"
	}
	
	return app.Div().
		Style("position", "fixed").
		Style("top", "0").
		Style("left", "0").
		Style("right", "0").
		Style("bottom", "0").
		Style("background", "rgba(0,0,0,0.5)").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("z-index", "1000").
		OnClick(func(ctx app.Context, e app.Event) {
			// Close modal when clicking backdrop
			if ctx.JSSrc().Get("target") == ctx.JSSrc().Get("currentTarget") {
				dp.showModal = false
				ctx.Reload()
			}
		}).
		Body(
			app.Div().
				Style("background", "white").
				Style("border-radius", "8px").
				Style("padding", "24px").
				Style("width", "500px").
				Style("max-width", "90vw").
				Style("max-height", "90vh").
				Style("overflow", "auto").
				Body(
					NewForm().
						WithTitle(title).
						WithFormData(dp.formData).
						WithShowCancel(true).
						WithOnSubmit(dp.handleFormSubmit).
						WithOnCancel(func(ctx app.Context) {
							dp.showModal = false
							dp.selectedUser = nil
							ctx.Reload()
						}),
				),
		)
}