package e2e

import (
	"fmt"
	"testing"
	"time"
)

func TestModalFramework(t *testing.T) {
	// This test verifies the modal framework components exist
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Verify that we can access the page without errors
	title := page.MustInfo().Title
	if title == "" {
		t.Error("Page title should not be empty")
	}

	t.Log("Modal framework test - basic structure verified")
}

func TestModalToastNotifications(t *testing.T) {
	// This test will verify toast notifications once we add demo triggers
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Future implementation would:
	// 1. Trigger an action that shows a toast
	// 2. Verify toast appears in the correct position
	// 3. Verify toast auto-dismisses after timeout
	// 4. Test manual toast dismissal
	// 5. Test toast actions if present

	t.Log("Toast test placeholder - needs demo page with toast triggers")

	// Check that toast container area is properly set up
	toastContainers := page.MustElements(".toast-container")
	t.Logf("Found %d toast containers", len(toastContainers))
}

func TestConfirmDialogs(t *testing.T) {
	// This test will verify confirm dialog functionality
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Future implementation would:
	// 1. Trigger a destructive action
	// 2. Verify confirm dialog appears
	// 3. Test cancel functionality
	// 4. Test confirm functionality
	// 5. Verify appropriate actions are taken

	t.Log("Confirm dialog test placeholder - needs demo page with confirm triggers")
}

func TestFormValidation(t *testing.T) {
	// Test form validation on the login page
	page := browser.MustPage(ServerURL + "/login")
	defer page.Close()

	page.MustWaitLoad()

	// Get form elements
	textFields := page.MustElements(".spectrum-Textfield")
	if len(textFields) < 2 {
		t.Skip("Login form not fully implemented yet")
		return
	}

	// Test empty form submission (if submit button exists)
	submitButtons := page.MustElements("button[type='submit'], .spectrum-Button")
	if len(submitButtons) > 0 {
		// Try to interact with the form
		t.Log("Form elements found, validation tests would go here")
	}

	t.Log("Form validation test placeholder - needs interactive form elements")
}

func TestStateManagement(t *testing.T) {
	// Test state persistence across page interactions
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Future tests would verify:
	// 1. State changes are reflected in UI
	// 2. State persists across page navigation
	// 3. State resets appropriately
	// 4. Loading states are shown correctly
	// 5. Error states are handled properly

	t.Log("State management test placeholder - needs stateful demo components")

	// Verify basic page state
	body := page.MustElement("body")
	if body == nil {
		t.Error("Body element should exist")
	}
}

func TestModalAccessibility(t *testing.T) {
	// Basic accessibility tests
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Check for basic accessibility requirements

	// Verify page has a title
	title := page.MustInfo().Title
	if title == "" {
		t.Error("Page should have a title for screen readers")
	}

	// Check for proper heading structure
	h1Elements := page.MustElements("h1")
	if len(h1Elements) == 0 {
		t.Error("Page should have at least one h1 element")
	}
	if len(h1Elements) > 1 {
		t.Logf("WARNING: %s", "Page should typically have only one h1 element")
	}

	// Check for alt text on images (if any)
	images := page.MustElements("img")
	for i, img := range images {
		alt := img.MustAttribute("alt")
		if alt == nil {
			t.Errorf("Image %d missing alt attribute", i)
		}
	}

	// Check for form labels
	inputs := page.MustElements("input")
	for i, input := range inputs {
		id := input.MustAttribute("id")
		if id != nil {
			labels := page.MustElements(fmt.Sprintf("label[for='%s']", *id))
			if len(labels) == 0 {
				t.Errorf("Input %d missing associated label", i)
			}
		}
	}
}

func TestPerformance(t *testing.T) {
	// Basic performance tests
	page := browser.MustPage()
	defer page.Close()

	// Measure page load time
	start := time.Now()
	page.MustNavigate(ServerURL)
	page.MustWaitLoad()
	loadTime := time.Since(start)

	t.Logf("Page load time: %v", loadTime)

	// Page should load reasonably quickly
	if loadTime > 5*time.Second {
		t.Errorf("Page load time too slow: %v", loadTime)
	}

	// Check for WASM loading
	wasmElements := page.MustElements("script[src$='.wasm']")
	t.Logf("Found %d WASM script elements", len(wasmElements))

	// Verify no JavaScript errors
	page.MustEval(`() => {
		if (window.console && console.error) {
			window.testErrors = [];
			const originalError = console.error;
			console.error = function(...args) {
				window.testErrors.push(args.join(' '));
				originalError.apply(console, args);
			};
		}
	}`)

	// Wait a bit for any async errors
	time.Sleep(2 * time.Second)

	errors := page.MustEval("() => window.testErrors || []").Arr()
	if len(errors) > 0 {
		t.Errorf("Found JavaScript errors: %v", errors)
	}
}
