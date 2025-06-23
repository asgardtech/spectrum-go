package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/go-rod/rod"
)

func TestDemoPageComponents(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// 1. Verify demo page loads
	h.AssertPageTitle("Prism Demo")
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Prism Framework Demo")

	// 2. Check header and description
	h.AssertElementExists("p")
	description := h.page.MustElement("p").MustText()
	if !strings.Contains(description, "Interactive showcase") {
		t.Errorf("Expected description to contain 'Interactive showcase', got '%s'", description)
	}

	// 3. Verify component showcase section
	h.AssertElementExists("h2")
	h2Elements := h.page.MustElements("h2")
	if len(h2Elements) > 0 && h2Elements[0].MustText() != "Component Showcase" {
		t.Errorf("Expected first h2 to be 'Component Showcase', got '%s'", h2Elements[0].MustText())
	}

	// 4. Check that toast buttons exist
	buttons := h.page.MustElements("button, .spectrum-Button")
	buttonTexts := make([]string, 0)
	for _, btn := range buttons {
		buttonTexts = append(buttonTexts, btn.MustText())
	}

	expectedButtons := []string{
		"Show Info Toast",
		"Show Success Toast",
		"Show Warning Toast",
		"Show Error Toast",
		"Open Modal",
		"Show Confirm Dialog",
		"Show Delete Confirm",
	}

	for _, expected := range expectedButtons {
		found := false
		for _, actual := range buttonTexts {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected button '%s' not found. Available buttons: %v", expected, buttonTexts)
		}
	}
}

func TestToastInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Test info toast
	h.ClickButton("Show Info Toast")

	// Wait for toast to appear
	toast := h.WaitForToast(5 * time.Second)
	if toast == nil {
		t.Fatal("Toast should appear after clicking info button")
	}

	// Verify toast content
	toastText := toast.MustText()
	if !strings.Contains(toastText, "info toast message") {
		t.Errorf("Expected toast to contain 'info toast message', got '%s'", toastText)
	}

	// Test success toast
	h.ClickButton("Show Success Toast")

	// Wait for success toast
	successToast := h.WaitForToast(5 * time.Second)
	successText := successToast.MustText()
	if !strings.Contains(successText, "successfully") {
		t.Errorf("Expected success toast to contain 'successfully', got '%s'", successText)
	}

	// Test warning toast
	h.ClickButton("Show Warning Toast")

	warningToast := h.WaitForToast(5 * time.Second)
	warningText := warningToast.MustText()
	if !strings.Contains(warningText, "review") {
		t.Errorf("Expected warning toast to contain 'review', got '%s'", warningText)
	}

	// Test error toast
	h.ClickButton("Show Error Toast")

	errorToast := h.WaitForToast(5 * time.Second)
	errorText := errorToast.MustText()
	if !strings.Contains(errorText, "error") {
		t.Errorf("Expected error toast to contain 'error', got '%s'", errorText)
	}
}

func TestDemoModalInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Initially no modal should be visible
	h.AssertElementNotExists(".spectrum-Modal")

	// Click to open modal
	h.ClickButton("Open Modal")

	// Wait for modal to appear
	modal := h.WaitForModal(5 * time.Second)
	if modal == nil {
		t.Fatal("Modal should appear after clicking open button")
	}

	// Verify modal content
	modalText := modal.MustText()
	if !strings.Contains(modalText, "Demo Modal") {
		t.Errorf("Expected modal to contain 'Demo Modal', got '%s'", modalText)
	}

	// Modal should be visible
	h.AssertElementVisible(".spectrum-Modal")

	// Test modal close functionality
	// Look for cancel/close button in modal
	cancelButtons := h.page.MustElements(".spectrum-Modal button")
	if len(cancelButtons) > 0 {
		cancelButtons[0].MustClick()

		// Wait for modal to disappear
		time.Sleep(500 * time.Millisecond)
		h.AssertElementNotExists(".spectrum-Modal")
	}
}

func TestConfirmDialogInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Test confirm dialog
	h.ClickButton("Show Confirm Dialog")

	// Wait for confirm dialog to appear
	dialog := h.WaitForModal(5 * time.Second)
	if dialog == nil {
		t.Fatal("Confirm dialog should appear")
	}

	// Verify dialog content
	dialogText := dialog.MustText()
	if !strings.Contains(dialogText, "Confirm Action") {
		t.Errorf("Expected dialog to contain 'Confirm Action', got '%s'", dialogText)
	}

	// Test cancel functionality
	cancelBtn := h.WaitForElementText("button", "Cancel", 3*time.Second)
	cancelBtn.MustClick()

	// Wait for dialog to disappear
	time.Sleep(500 * time.Millisecond)
	h.AssertElementNotExists(".spectrum-Modal")

	// Test delete confirm dialog
	h.ClickButton("Show Delete Confirm")

	deleteDialog := h.WaitForModal(5 * time.Second)
	deleteText := deleteDialog.MustText()
	if !strings.Contains(deleteText, "Delete") {
		t.Errorf("Expected delete dialog to contain 'Delete', got '%s'", deleteText)
	}

	// Cancel the delete dialog too
	cancelBtn = h.WaitForElementText("button", "Cancel", 3*time.Second)
	cancelBtn.MustClick()

	time.Sleep(500 * time.Millisecond)
	h.AssertElementNotExists(".spectrum-Modal")
}

func TestDataTableInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Verify data table section exists
	h2Elements := h.page.MustElements("h2")
	found := false
	for _, h2 := range h2Elements {
		if h2.MustText() == "Sample Data Table" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'Sample Data Table' section")
	}

	// Check that table rows exist
	checkboxes := h.page.MustElements("input[type='checkbox']")
	if len(checkboxes) < 3 {
		t.Errorf("Expected at least 3 checkboxes (for sample payouts), got %d", len(checkboxes))
	}

	// Test row selection
	if len(checkboxes) > 0 {
		// Click first checkbox
		checkboxes[0].MustClick()

		// Verify it's checked
		if !checkboxes[0].MustProperty("checked").Bool() {
			t.Error("Checkbox should be checked after clicking")
		}

		// Click again to uncheck
		checkboxes[0].MustClick()

		// Verify it's unchecked
		if checkboxes[0].MustProperty("checked").Bool() {
			t.Error("Checkbox should be unchecked after clicking again")
		}
	}

	// Test Select All button
	h.ClickButton("Select All")

	// Wait a moment for state change
	time.Sleep(500 * time.Millisecond)

	// Check that all checkboxes are now selected
	updatedCheckboxes := h.page.MustElements("input[type='checkbox']")
	for i, checkbox := range updatedCheckboxes {
		if !checkbox.MustProperty("checked").Bool() {
			t.Errorf("Checkbox %d should be checked after Select All", i)
		}
	}

	// Test Clear Selection button
	h.ClickButton("Clear Selection")

	time.Sleep(500 * time.Millisecond)

	// Check that all checkboxes are now unselected
	clearedCheckboxes := h.page.MustElements("input[type='checkbox']")
	for i, checkbox := range clearedCheckboxes {
		if checkbox.MustProperty("checked").Bool() {
			t.Errorf("Checkbox %d should be unchecked after Clear Selection", i)
		}
	}
}

func TestPayoutRowInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Find View buttons
	viewButtons := h.page.MustElements("button")
	var viewBtn *rod.Element
	for _, btn := range viewButtons {
		if btn.MustText() == "View" {
			viewBtn = btn
			break
		}
	}

	if viewBtn == nil {
		t.Fatal("Expected at least one View button in payout rows")
	}

	// Click View button to open payout details modal
	viewBtn.MustClick()

	// Wait for modal with payout details
	modal := h.WaitForModal(5 * time.Second)
	modalText := modal.MustText()

	if !strings.Contains(modalText, "Payout Details") {
		t.Errorf("Expected modal to contain 'Payout Details', got '%s'", modalText)
	}

	// Modal should contain payout information
	if !strings.Contains(modalText, "Payout ID") {
		t.Error("Expected modal to contain payout ID information")
	}

	// Close the modal
	cancelButtons := h.page.MustElements(".spectrum-Modal button")
	if len(cancelButtons) > 0 {
		cancelButtons[0].MustClick()
		time.Sleep(500 * time.Millisecond)
		h.AssertElementNotExists(".spectrum-Modal")
	}
}

func TestProcessSelectedWorkflow(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// 1. Test processing with no selection
	h.ClickButton("Process Selected")

	// Should show warning toast
	toast := h.WaitForToast(5 * time.Second)
	toastText := toast.MustText()
	if !strings.Contains(toastText, "No payouts selected") {
		t.Errorf("Expected warning about no selection, got '%s'", toastText)
	}

	// 2. Select some payouts and process
	h.ClickButton("Select All")
	time.Sleep(500 * time.Millisecond)

	h.ClickButton("Process Selected")

	// Should show success toast
	successToast := h.WaitForToast(5 * time.Second)
	successText := successToast.MustText()
	if !strings.Contains(successText, "Processing") {
		t.Errorf("Expected success message about processing, got '%s'", successText)
	}
}

func TestNavigationAndLayoutConsistency(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Demo page should have logged-in user layout
	// Should show top nav and side nav since user is logged in
	h.AssertElementExists(".spectrum-TopNav")
	h.AssertElementExists(".spectrum-SplitView")

	// Navigate to other pages and back
	h.Navigate(ServerURL + "/")
	h.AssertPageTitle("Home")

	h.Navigate(ServerURL + "/demo")
	h.AssertPageTitle("Prism Demo")

	// Content should be consistent after navigation
	h.AssertElementText("h1", "Prism Framework Demo")
}

func TestResponsiveDemo(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/demo")
	defer h.Close()

	// Test desktop layout
	h.Desktop()
	h.Reload()

	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Prism Framework Demo")

	// Test mobile layout
	h.Mobile()
	h.Reload()

	// Content should still be accessible
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Prism Framework Demo")

	// Buttons should still be clickable
	buttons := h.page.MustElements("button")
	if len(buttons) == 0 {
		t.Error("Expected buttons to be available on mobile")
	}

	// Test tablet layout
	h.Tablet()
	h.Reload()

	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Prism Framework Demo")
}
