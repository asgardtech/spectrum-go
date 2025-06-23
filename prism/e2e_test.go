package prism

import (
	"context"
	"testing"
	"time"
)

// Skip E2E tests for now - focus on unit tests
func TestComprehensiveDemo(t *testing.T) {
	t.Skip("E2E tests require browser automation setup")
}

func TestFormComponents(t *testing.T) {
	t.Run("FormData validation", func(t *testing.T) {
		formData := NewFormData().
			AddField(&FormField{
				Name:       "email",
				Label:      "Email",
				Type:       "email",
				Required:   true,
				Validators: []Validator{EmailValidator()},
			})

		// Test required validation
		formData.SetFieldValue("email", "")
		if formData.ValidateAll() {
			t.Error("Empty required field should fail validation")
		}

		// Test email validation
		formData.SetFieldValue("email", "invalid-email")
		if formData.ValidateAll() {
			t.Error("Invalid email should fail validation")
		}

		// Test valid email
		formData.SetFieldValue("email", "test@example.com")
		if !formData.ValidateAll() {
			t.Error("Valid email should pass validation")
		}
	})

	t.Run("Validator functions", func(t *testing.T) {
		// Test EmailValidator
		emailValidator := EmailValidator()
		
		result := emailValidator("invalid")
		if result.Valid {
			t.Error("EmailValidator should reject invalid email")
		}

		result = emailValidator("test@example.com")
		if !result.Valid {
			t.Error("EmailValidator should accept valid email")
		}

		// Test MinLengthValidator
		minLengthValidator := MinLengthValidator(5)
		
		result = minLengthValidator("abc")
		if result.Valid {
			t.Error("MinLengthValidator should reject short string")
		}

		result = minLengthValidator("abcdef")
		if !result.Valid {
			t.Error("MinLengthValidator should accept long enough string")
		}
	})
}

func TestDataTableComponents(t *testing.T) {
	t.Run("DataTable creation and configuration", func(t *testing.T) {
		users := []*DemoUser{
			{ID: 1, Name: "John", Email: "john@example.com", Role: "Admin"},
			{ID: 2, Name: "Jane", Email: "jane@example.com", Role: "User"},
		}

		table := NewDataTable[*DemoUser]().
			WithData(users).
			WithSelectable(true).
			WithShowFilters(true).
			WithPageSize(10)

		if !table.Selectable {
			t.Error("Table should be selectable")
		}

		if !table.ShowFilters {
			t.Error("Table should show filters")
		}

		if table.State.PageSize != 10 {
			t.Error("Table page size should be 10")
		}

		if len(table.State.Data) != 2 {
			t.Error("Table should have 2 users")
		}
	})

	t.Run("DataTable filtering", func(t *testing.T) {
		users := []*DemoUser{
			{ID: 1, Name: "John", Email: "john@example.com", Role: "Admin"},
			{ID: 2, Name: "Jane", Email: "jane@example.com", Role: "User"},
			{ID: 3, Name: "Bob", Email: "bob@example.com", Role: "User"},
		}

		table := NewDataTable[*DemoUser]().WithData(users)
		
		// Add a filter
		table.AddFilter("name", "John", "text")
		
		// Check filtered data
		if len(table.State.FilteredData) == 0 {
			t.Error("Filter should return at least one result for 'John'")
		}

		// Clear filter
		table.AddFilter("name", "", "text")
		
		if len(table.State.FilteredData) != 3 {
			t.Error("Clearing filter should return all data")
		}
	})

	t.Run("DataTable selection", func(t *testing.T) {
		users := []*DemoUser{
			{ID: 1, Name: "John", Email: "john@example.com", Role: "Admin"},
			{ID: 2, Name: "Jane", Email: "jane@example.com", Role: "User"},
		}

		table := NewDataTable[*DemoUser]().WithData(users)

		// Test row selection
		table.ToggleRowSelection(0)
		if !table.State.SelectedRows[0] {
			t.Error("Row 0 should be selected")
		}

		// Test select all
		table.SelectAll()
		if len(table.State.SelectedRows) != 2 {
			t.Error("All rows should be selected")
		}

		// Test clear selection
		table.ClearSelection()
		if len(table.State.SelectedRows) != 0 {
			t.Error("No rows should be selected after clear")
		}
	})
}

// Helper function to run tests with timeout
func runWithTimeout(t *testing.T, timeout time.Duration, testFunc func(*testing.T)) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan bool)
	go func() {
		testFunc(t)
		done <- true
	}()

	select {
	case <-ctx.Done():
		t.Fatal("Test timed out")
	case <-done:
		// Test completed
	}
}