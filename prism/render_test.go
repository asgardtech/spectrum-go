package prism

import (
	"testing"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func TestComponentRendering(t *testing.T) {
	t.Run("FormComponent renders without error", func(t *testing.T) {
		formData := NewFormData().
			AddField(&FormField{
				Name:     "test",
				Label:    "Test Field",
				Type:     "text",
				Required: true,
			})

		form := NewForm().
			WithFormData(formData).
			WithTitle("Test Form")

		// Test that Render returns a UI component
		ui := form.Render()
		if ui == nil {
			t.Error("Form.Render() should return a UI component")
		}
	})

	t.Run("DataTable renders without error", func(t *testing.T) {
		users := []*DemoUser{
			{ID: 1, Name: "John", Email: "john@example.com", Role: "Admin"},
		}

		columns := []TableColumn[*DemoUser]{
			{
				Key:   "name",
				Label: "Name",
				Render: func(user *DemoUser) app.UI {
					return app.Text(user.Name)
				},
			},
		}

		table := NewDataTable[*DemoUser]().
			WithData(users).
			WithColumns(columns...)

		// Test that Render returns a UI component
		ui := table.Render()
		if ui == nil {
			t.Error("DataTable.Render() should return a UI component")
		}
	})

	t.Run("EditableField renders without error", func(t *testing.T) {
		field := NewEditableField("test", "Test Field").
			WithType("text").
			WithValue("test value").
			WithRequired(true)

		// Test that Render returns a UI component
		ui := field.Render()
		if ui == nil {
			t.Error("EditableField.Render() should return a UI component")
		}
	})

	t.Run("FieldGroup renders without error", func(t *testing.T) {
		field1 := NewEditableField("field1", "Field 1")
		field2 := NewEditableField("field2", "Field 2")

		group := NewFieldGroup().
			WithTitle("Test Group").
			WithFields(field1, field2).
			WithLayout("vertical")

		// Test that Render returns a UI component
		ui := group.Render()
		if ui == nil {
			t.Error("FieldGroup.Render() should return a UI component")
		}
	})

	t.Run("ComprehensiveDemo renders without error", func(t *testing.T) {
		demo := NewComprehensiveDemo()

		// Test that Render returns a UI component
		ui := demo.Render()
		if ui == nil {
			t.Error("ComprehensiveDemo.Render() should return a UI component")
		}
	})

	t.Run("TopNav renders without error", func(t *testing.T) {
		topNav := NewTopNav().
			WithTitle("Test App").
			WithSearchBar(NewSearchBar().WithPlaceholder("Search..."))

		ui := topNav.Render()
		if ui == nil {
			t.Error("TopNav.Render() should return a UI component")
		}
	})

	t.Run("UserMenu renders without error", func(t *testing.T) {
		user := &User{
			ID:    "1",
			Name:  "Test User",
			Email: "test@example.com",
		}

		userMenu := NewUserMenu().
			WithUser(user).
			WithMenuItems(DefaultUserMenuItems()...)

		ui := userMenu.Render()
		if ui == nil {
			t.Error("UserMenu.Render() should return a UI component")
		}
	})

	t.Run("Card renders without error", func(t *testing.T) {
		card := NewCard().
			WithTitle("Test Card").
			WithContent(app.P().Text("This is test content")).
			WithActions(
				NewCardAction("test", "Test Action").WithVariant("primary"),
			)

		ui := card.Render()
		if ui == nil {
			t.Error("Card.Render() should return a UI component")
		}
	})

	t.Run("EmptyState renders without error", func(t *testing.T) {
		emptyState := NewEmptyState().
			WithTitle("No Data").
			WithDescription("There is no data to display").
			WithActions(
				NewEmptyStateAction("reload", "Reload").WithVariant("primary"),
			)

		ui := emptyState.Render()
		if ui == nil {
			t.Error("EmptyState.Render() should return a UI component")
		}
	})

	t.Run("SideNav renders without error", func(t *testing.T) {
		sideNav := NewSideNav().
			WithTitle("Test Navigation").
			WithItems(
				NewSideNavItem("home", "Home").WithIcon("🏠").WithActive(true),
				NewSideNavItem("about", "About").WithIcon("ℹ️"),
			)

		ui := sideNav.Render()
		if ui == nil {
			t.Error("SideNav.Render() should return a UI component")
		}
	})

	t.Run("SideNav with groups renders without error", func(t *testing.T) {
		sideNav := NewSideNav().
			WithTitle("Admin Panel").
			WithGroups(
				NewSideNavGroup("content", "Content").
					WithItems(
						NewSideNavItem("posts", "Posts").WithIcon("📝"),
						NewSideNavItem("pages", "Pages").WithIcon("📄"),
					),
			)

		ui := sideNav.Render()
		if ui == nil {
			t.Error("SideNav with groups should render")
		}
	})
}

func TestUtilityFunctions(t *testing.T) {
	t.Run("IfElseString works correctly", func(t *testing.T) {
		result := IfElseString(true, "yes", "no")
		if result != "yes" {
			t.Errorf("Expected 'yes', got '%s'", result)
		}

		result = IfElseString(false, "yes", "no")
		if result != "no" {
			t.Errorf("Expected 'no', got '%s'", result)
		}
	})

	t.Run("Min works correctly", func(t *testing.T) {
		result := Min(5, 3)
		if result != 3 {
			t.Errorf("Min(5, 3) should return 3, got %d", result)
		}

		result = Min(2, 8)
		if result != 2 {
			t.Errorf("Min(2, 8) should return 2, got %d", result)
		}
	})

	t.Run("Format functions work", func(t *testing.T) {
		// Test FormatNumber
		result := FormatNumber(1234)
		if result != "1,234" {
			t.Errorf("FormatNumber(1234) should return '1,234', got '%s'", result)
		}

		// Test FormatCurrency
		currency := FormatCurrency(99.99, "USD")
		if currency != "$99.99" {
			t.Errorf("FormatCurrency(99.99, 'USD') should return '$99.99', got '%s'", currency)
		}

		// Test TruncateText
		truncated := TruncateText("This is a long text", 10)
		if truncated != "This is..." {
			t.Errorf("TruncateText should truncate text properly, got '%s'", truncated)
		}
	})
}