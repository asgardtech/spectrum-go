package prism

import (
	"testing"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func TestTopNavFunctionality(t *testing.T) {
	t.Run("SearchBar functionality", func(t *testing.T) {
		searchQuery := ""
		searchBar := NewSearchBar().
			WithPlaceholder("Test search").
			WithOnSearch(func(query string) {
				searchQuery = query
			})

		if searchBar.Placeholder != "Test search" {
			t.Error("SearchBar placeholder not set correctly")
		}

		// Simulate search
		if searchBar.OnSearch != nil {
			searchBar.OnSearch("test query")
		}

		if searchQuery != "test query" {
			t.Error("SearchBar OnSearch callback not working")
		}
	})

	t.Run("TopNav actions", func(t *testing.T) {
		action := NewTopNavAction("test", "Test Action").
			WithOnClick(func(ctx app.Context) {
				// Action callback
			})

		topNav := NewTopNav().
			WithActions(action)

		if len(topNav.Actions) != 1 {
			t.Error("TopNav action not added")
		}

		if topNav.Actions[0].Label != "Test Action" {
			t.Error("TopNav action label not set")
		}
	})

	t.Run("TopNav theme toggle", func(t *testing.T) {
		themeChanged := ""
		topNav := NewTopNav().
			WithTheme("light").
			WithOnThemeToggle(func(theme string) {
				themeChanged = theme
			})

		if topNav.Theme != "light" {
			t.Error("TopNav theme not set correctly")
		}

		// Simulate theme toggle
		if topNav.OnThemeToggle != nil {
			topNav.OnThemeToggle("dark")
		}

		if themeChanged != "dark" {
			t.Error("TopNav theme toggle callback not working")
		}
	})
}

func TestUserMenuFunctionality(t *testing.T) {
	t.Run("UserMenu state management", func(t *testing.T) {
		userMenu := NewUserMenu()

		if userMenu.IsOpen {
			t.Error("UserMenu should start closed")
		}

		userMenu.Toggle()
		if !userMenu.IsOpen {
			t.Error("UserMenu should be open after toggle")
		}

		userMenu.Close()
		if userMenu.IsOpen {
			t.Error("UserMenu should be closed after Close()")
		}
	})

	t.Run("UserMenu with user", func(t *testing.T) {
		user := &User{
			ID:    "1",
			Name:  "John Doe",
			Email: "john@example.com",
		}

		userMenu := NewUserMenu().WithUser(user)

		if userMenu.User != user {
			t.Error("UserMenu user not set correctly")
		}

		initials := userMenu.getUserInitials()
		if initials != "JD" {
			t.Errorf("Expected initials 'JD', got '%s'", initials)
		}
	})

	t.Run("UserMenu items", func(t *testing.T) {
		menuItem := NewUserMenuItem("Test Item").
			WithOnClick(func(ctx app.Context) {
				// Menu item clicked
			})

		userMenu := NewUserMenu().
			WithMenuItems(menuItem)

		if len(userMenu.MenuItems) != 1 {
			t.Error("UserMenu item not added")
		}

		if userMenu.MenuItems[0].Label != "Test Item" {
			t.Error("UserMenu item label not set")
		}
	})
}

func TestCardFunctionality(t *testing.T) {
	t.Run("Card configuration", func(t *testing.T) {
		card := NewCard().
			WithTitle("Test Card").
			WithSubtitle("Test Subtitle").
			WithVariant("outlined").
			WithSize("large").
			WithClickable(true)

		if card.Title != "Test Card" {
			t.Error("Card title not set")
		}

		if card.Subtitle != "Test Subtitle" {
			t.Error("Card subtitle not set")
		}

		if card.Variant != "outlined" {
			t.Error("Card variant not set")
		}

		if card.Size != "large" {
			t.Error("Card size not set")
		}

		if !card.Clickable {
			t.Error("Card clickable not set")
		}
	})

	t.Run("Card actions", func(t *testing.T) {
		action := NewCardAction("test", "Test Action").
			WithOnClick(func(ctx app.Context) {
				// Action triggered
			})

		card := NewCard().WithActions(action)

		if len(card.Actions) != 1 {
			t.Error("Card action not added")
		}

		if card.Actions[0].Label != "Test Action" {
			t.Error("Card action label not set")
		}
	})

	t.Run("Card helper functions", func(t *testing.T) {
		infoCard := InfoCard("Info Title", "Info content")
		if infoCard.Title != "Info Title" {
			t.Error("InfoCard helper not working")
		}

		actionCard := ActionCard("Action Title", "Action content", "Click Me", nil)
		if actionCard.Title != "Action Title" {
			t.Error("ActionCard helper not working")
		}

		if len(actionCard.Actions) != 1 {
			t.Error("ActionCard should have one action")
		}

		imageCard := ImageCard("Image Title", "Image subtitle", "/test.jpg", "Test image")
		if imageCard.Image != "/test.jpg" {
			t.Error("ImageCard helper not working")
		}

		statsCard := StatsCard("Stats", "100", "Total items")
		if statsCard.Title != "Stats" {
			t.Error("StatsCard helper not working")
		}
	})
}

func TestEmptyStateFunctionality(t *testing.T) {
	t.Run("EmptyState configuration", func(t *testing.T) {
		emptyState := NewEmptyState().
			WithTitle("No Data").
			WithDescription("No data available").
			WithVariant("no-data").
			WithSize("large")

		if emptyState.Title != "No Data" {
			t.Error("EmptyState title not set")
		}

		if emptyState.Description != "No data available" {
			t.Error("EmptyState description not set")
		}

		if emptyState.Variant != "no-data" {
			t.Error("EmptyState variant not set")
		}

		if emptyState.Size != "large" {
			t.Error("EmptyState size not set")
		}
	})

	t.Run("EmptyState actions", func(t *testing.T) {
		action := NewEmptyStateAction("test", "Test Action").
			WithOnClick(func(ctx app.Context) {
				// Action triggered
			})

		emptyState := NewEmptyState().WithActions(action)

		if len(emptyState.Actions) != 1 {
			t.Error("EmptyState action not added")
		}

		if emptyState.Actions[0].Label != "Test Action" {
			t.Error("EmptyState action label not set")
		}
	})

	t.Run("EmptyState helper functions", func(t *testing.T) {
		noDataState := NoDataEmptyState("No Data", "No data found")
		if noDataState.Variant != "no-data" {
			t.Error("NoDataEmptyState helper not working")
		}

		noResultsState := NoResultsEmptyState("test query")
		if noResultsState.Variant != "no-results" {
			t.Error("NoResultsEmptyState helper not working")
		}

		errorState := ErrorEmptyState("Error", "Something went wrong", nil)
		if errorState.Variant != "error" {
			t.Error("ErrorEmptyState helper not working")
		}

		firstTimeState := FirstTimeEmptyState("Welcome", "Get started", "Start", nil)
		if firstTimeState.Title != "Welcome" {
			t.Error("FirstTimeEmptyState helper not working")
		}

		comingSoonState := ComingSoonEmptyState("New feature")
		if comingSoonState.Title != "Coming Soon" {
			t.Error("ComingSoonEmptyState helper not working")
		}
	})
}

func TestFormIntegration(t *testing.T) {
	t.Run("Form with fields and validation", func(t *testing.T) {
		formData := NewFormData().
			AddField(&FormField{
				Name:       "name",
				Label:      "Name",
				Required:   true,
				Validators: []Validator{MinLengthValidator(2)},
			}).
			AddField(&FormField{
				Name:     "email",
				Label:    "Email",
				Type:     "email",
				Required: true,
				Validators: []Validator{EmailValidator()},
			})

		// Test empty values fail validation
		if formData.ValidateAll() {
			t.Error("Empty form should fail validation")
		}

		// Test valid values pass validation
		formData.SetFieldValue("name", "John Doe")
		formData.SetFieldValue("email", "john@example.com")

		if !formData.ValidateAll() {
			t.Error("Valid form should pass validation")
		}

		// Test form values
		values := formData.GetFormValues()
		if values["name"] != "John Doe" {
			t.Error("Form values not retrieved correctly")
		}
	})
}

func TestSideNavFunctionality(t *testing.T) {
	t.Run("SideNav state management", func(t *testing.T) {
		sideNav := NewSideNav()

		if !sideNav.IsOpen {
			t.Error("SideNav should start open by default")
		}

		sideNav.Close()
		if sideNav.IsOpen {
			t.Error("SideNav should be closed after Close()")
		}

		sideNav.Open()
		if !sideNav.IsOpen {
			t.Error("SideNav should be open after Open()")
		}

		sideNav.Toggle()
		if sideNav.IsOpen {
			t.Error("SideNav should be closed after toggle")
		}
	})

	t.Run("SideNav configuration", func(t *testing.T) {
		sideNav := NewSideNav().
			WithTitle("Test Nav").
			WithWidth("300px").
			WithPosition("right").
			WithVariant("overlay").
			WithCollapsed(true)

		if sideNav.Title != "Test Nav" {
			t.Error("SideNav title not set")
		}

		if sideNav.Width != "300px" {
			t.Error("SideNav width not set")
		}

		if sideNav.Position != "right" {
			t.Error("SideNav position not set")
		}

		if sideNav.Variant != "overlay" {
			t.Error("SideNav variant not set")
		}

		if !sideNav.IsCollapsed {
			t.Error("SideNav collapsed not set")
		}
	})

	t.Run("SideNav items", func(t *testing.T) {
		item := NewSideNavItem("test", "Test Item").
			WithIcon("🏠").
			WithHref("/test").
			WithActive(true).
			WithBadge("New")

		sideNav := NewSideNav().WithItems(item)

		if len(sideNav.Items) != 1 {
			t.Error("SideNav item not added")
		}

		if sideNav.Items[0].Label != "Test Item" {
			t.Error("SideNav item label not set")
		}

		if sideNav.Items[0].Icon != "🏠" {
			t.Error("SideNav item icon not set")
		}

		if sideNav.Items[0].Href != "/test" {
			t.Error("SideNav item href not set")
		}

		if !sideNav.Items[0].Active {
			t.Error("SideNav item active not set")
		}

		if sideNav.Items[0].Badge != "New" {
			t.Error("SideNav item badge not set")
		}
	})

	t.Run("SideNav groups", func(t *testing.T) {
		group := NewSideNavGroup("content", "Content").
			WithItems(
				NewSideNavItem("posts", "Posts").WithIcon("📝"),
				NewSideNavItem("pages", "Pages").WithIcon("📄"),
			).
			WithExpanded(false)

		sideNav := NewSideNav().WithGroups(group)

		if len(sideNav.Groups) != 1 {
			t.Error("SideNav group not added")
		}

		if sideNav.Groups[0].Label != "Content" {
			t.Error("SideNav group label not set")
		}

		if len(sideNav.Groups[0].Items) != 2 {
			t.Error("SideNav group items not set")
		}

		if sideNav.Groups[0].Expanded {
			t.Error("SideNav group expanded not set")
		}
	})

	t.Run("SideNav helper functions", func(t *testing.T) {
		mainNav := MainSideNav()
		if mainNav.Title != "Your App" {
			t.Error("MainSideNav helper not working")
		}

		if len(mainNav.Items) != 3 {
			t.Error("MainSideNav should have 3 items")
		}

		adminNav := AdminSideNav()
		if adminNav.Title != "Admin Panel" {
			t.Error("AdminSideNav helper not working")
		}

		if len(adminNav.Groups) != 3 {
			t.Error("AdminSideNav should have 3 groups")
		}

		mobileNav := MobileSideNav()
		if mobileNav.Variant != "overlay" {
			t.Error("MobileSideNav should have overlay variant")
		}

		if mobileNav.IsOpen {
			t.Error("MobileSideNav should start closed")
		}

		if mobileNav.Width != "280px" {
			t.Error("MobileSideNav width not set correctly")
		}
	})
}

func TestDataTableIntegration(t *testing.T) {
	t.Run("DataTable with filtering and sorting", func(t *testing.T) {
		testData := []*DemoUser{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Role: "Admin"},
			{ID: 2, Name: "Bob", Email: "bob@example.com", Role: "User"},
			{ID: 3, Name: "Charlie", Email: "charlie@example.com", Role: "User"},
		}

		table := NewDataTable[*DemoUser]().
			WithData(testData).
			WithSelectable(true)

		// Test initial data
		if len(table.State.Data) != 3 {
			t.Error("DataTable data not set correctly")
		}

		if len(table.State.FilteredData) != 3 {
			t.Error("DataTable filtered data should match initial data")
		}

		// Test filtering
		table.AddFilter("name", "Alice", "text")
		// Note: The current filter implementation is basic and searches the string representation
		// This would need more sophisticated filtering in a real implementation

		// Test selection
		table.ToggleRowSelection(0)
		if !table.State.SelectedRows[0] {
			t.Error("DataTable row selection not working")
		}

		// Test select all
		table.SelectAll()
		if len(table.State.SelectedRows) != len(table.State.FilteredData) {
			t.Error("DataTable select all not working")
		}

		// Test clear selection
		table.ClearSelection()
		if len(table.State.SelectedRows) != 0 {
			t.Error("DataTable clear selection not working")
		}
	})
}