package e2e

import (
	"testing"
	"time"
)

func TestCompleteUserWorkflow(t *testing.T) {
	// Test a complete user workflow from landing to interaction
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()

	// 1. User lands on home page
	h.AssertPageTitle("Home")
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Home")

	// 2. Check initial state - no user logged in
	// Top nav and side nav should be hidden for anonymous users
	h.AssertElementNotExists(".spectrum-TopNav")
	h.AssertElementNotExists(".spectrum-Sidenav")

	// 3. Navigate to login page
	h.Navigate(ServerURL + "/login")
	h.AssertPageTitle("Login")

	// 4. Check login form is present
	h.AssertElementExists("input") // Should have email and password fields
	
	// Find email and password fields
	labels := h.page.MustElements(".spectrum-FieldLabel")
	if len(labels) >= 2 {
		emailLabel := labels[0].MustText()
		passwordLabel := labels[1].MustText()
		
		if emailLabel != "Email" {
			t.Errorf("Expected first label to be 'Email', got '%s'", emailLabel)
		}
		if passwordLabel != "Password" {
			t.Errorf("Expected second label to be 'Password', got '%s'", passwordLabel)
		}
	}

	// 5. Test form interaction (fill but don't submit since auth isn't implemented)
	inputs := h.page.MustElements("input")
	if len(inputs) >= 2 {
		inputs[0].MustInput("test@example.com")
		inputs[1].MustInput("password123")
		
		// Verify input values
		emailValue := inputs[0].MustProperty("value").String()
		passwordValue := inputs[1].MustProperty("value").String()
		
		if emailValue != "test@example.com" {
			t.Errorf("Expected email value 'test@example.com', got '%s'", emailValue)
		}
		if passwordValue != "password123" {
			t.Errorf("Expected password value 'password123', got '%s'", passwordValue)
		}
	}

	// 6. Navigate back to home
	h.Navigate(ServerURL + "/")
	h.AssertPageTitle("Home")
}

func TestResponsiveNavigation(t *testing.T) {
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()

	// Test navigation on different screen sizes
	
	// Desktop view
	h.Desktop()
	h.AssertElementExists("body")
	
	// Check that content is properly sized
	body := h.page.MustElement("body")
	width := body.MustEval("() => this.offsetWidth").Int()
	if width < 1000 {
		t.Errorf("Desktop view should be at least 1000px wide, got %d", width)
	}

	// Mobile view
	h.Mobile()
	h.Reload() // Reload to trigger responsive layout
	
	// Content should still be accessible
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Home")

	// Tablet view
	h.Tablet()
	h.Reload()
	
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Home")
}

func TestPageLoadPerformance(t *testing.T) {
	// Test that pages load within reasonable time
	start := time.Now()
	
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()
	
	loadTime := time.Since(start)
	t.Logf("Home page load time: %v", loadTime)
	
	if loadTime > 10*time.Second {
		t.Errorf("Home page load time too slow: %v (expected < 10s)", loadTime)
	}

	// Test login page load time
	start = time.Now()
	h.Navigate(ServerURL + "/login")
	loginLoadTime := time.Since(start)
	
	t.Logf("Login page load time: %v", loginLoadTime)
	
	if loginLoadTime > 5*time.Second {
		t.Errorf("Login page load time too slow: %v (expected < 5s)", loginLoadTime)
	}
}

func TestErrorHandling(t *testing.T) {
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()

	// Test navigation to non-existent page
	h.Navigate(ServerURL + "/nonexistent")
	
	// Should either show 404 or redirect to home
	// For now, just verify the page loads without crashing
	h.WaitForNetworkIdle(2 * time.Second)
	
	// Check that no JavaScript errors occurred
	errorsResult := h.ExecuteJS("() => window.testErrors || []")
	errors := errorsResult.Arr()
	if len(errors) > 0 {
		t.Errorf("JavaScript errors found: %v", errors)
	}
}

func TestFormValidationStates(t *testing.T) {
	h := NewTestHelpers(t, ServerURL + "/login")
	defer h.Close()

	// Test empty form state
	inputs := h.page.MustElements("input")
	if len(inputs) >= 2 {
		// Clear any existing values
		inputs[0].MustSelectAllText().MustInput("")
		inputs[1].MustSelectAllText().MustInput("")
		
		// Verify fields are empty
		emailValue := inputs[0].MustProperty("value").String()
		passwordValue := inputs[1].MustProperty("value").String()
		
		if emailValue != "" {
			t.Errorf("Email field should be empty, got '%s'", emailValue)
		}
		if passwordValue != "" {
			t.Errorf("Password field should be empty, got '%s'", passwordValue)
		}
	}

	// Test partial form fill
	if len(inputs) >= 2 {
		inputs[0].MustInput("test@example.com")
		// Leave password empty
		
		emailValue := inputs[0].MustProperty("value").String()
		passwordValue := inputs[1].MustProperty("value").String()
		
		if emailValue != "test@example.com" {
			t.Errorf("Email field should be 'test@example.com', got '%s'", emailValue)
		}
		if passwordValue != "" {
			t.Errorf("Password field should be empty, got '%s'", passwordValue)
		}
	}
}

func TestAccessibilityCompliance(t *testing.T) {
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()

	// Check for basic accessibility features
	
	// 1. Page should have a proper title
	title := h.page.MustInfo().Title
	if title == "" {
		t.Error("Page should have a title for accessibility")
	}

	// 2. Should have proper heading structure
	h1Elements := h.page.MustElements("h1")
	if len(h1Elements) == 0 {
		t.Error("Page should have at least one h1 element")
	}
	if len(h1Elements) > 1 {
		t.Logf("WARNING: %s", "Page should typically have only one h1 element for accessibility")
	}

	// 3. Images should have alt attributes
	images := h.page.MustElements("img")
	for i, img := range images {
		alt := img.MustAttribute("alt")
		if alt == nil || *alt == "" {
			t.Errorf("Image %d missing alt attribute", i)
		}
	}

	// 4. Check language attribute
	html := h.page.MustElement("html")
	lang := html.MustAttribute("lang")
	if lang == nil || *lang == "" {
		t.Logf("WARNING: %s", "HTML element should have lang attribute for accessibility")
	}

	// Test login page accessibility
	h.Navigate(ServerURL + "/login")
	
	// Form inputs should have associated labels
	inputs := h.page.MustElements("input")
	labels := h.page.MustElements("label, .spectrum-FieldLabel")
	
	if len(inputs) > 0 && len(labels) < len(inputs) {
		t.Logf("WARNING: %s", "All form inputs should have associated labels")
	}
}

func TestStateConsistency(t *testing.T) {
	h := NewTestHelpers(t, ServerURL)
	defer h.Close()

	// Test that page state remains consistent across navigation
	
	// 1. Start on home page
	h.AssertElementText("h1", "Home")
	
	// 2. Navigate to login
	h.Navigate(ServerURL + "/login")
	h.AssertPageTitle("Login")
	
	// 3. Navigate back to home
	h.Navigate(ServerURL + "/")
	h.AssertPageTitle("Home")
	h.AssertElementText("h1", "Home")
	
	// 4. Refresh page
	h.Reload()
	h.AssertPageTitle("Home")
	h.AssertElementText("h1", "Home")
	
	// State should be consistent after refresh
	// (This tests that the app initializes properly)
}

