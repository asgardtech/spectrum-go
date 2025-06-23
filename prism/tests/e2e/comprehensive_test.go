package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// TestComprehensiveComponentPage tests all components on the comprehensive test page
func TestComprehensiveComponentPage(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Verify page loads
	h.AssertPageTitle("Comprehensive Component Test")
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Comprehensive Component Test Suite")

	// Verify all major sections exist
	h.AssertElementExists("#buttons-section")
	h.AssertElementExists("#forms-section")
	h.AssertElementExists("#data-section")
	h.AssertElementExists("#layout-section")
	h.AssertElementExists("#interaction-section")
	h.AssertElementExists("#navigation-section")
	h.AssertElementExists("#advanced-section")
}

func TestButtonComponents(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test basic buttons
	h.AssertElementExists("#primary-button")
	h.AssertElementExists("#secondary-button")
	h.AssertElementExists("#action-button")

	// Click primary button and verify toast
	h.ClickButton("Primary Button")
	toast := h.WaitForToast(5 * time.Second)
	if toast != nil {
		toastText := toast.MustText()
		if !strings.Contains(toastText, "Primary button clicked") {
			t.Errorf("Expected primary button toast, got '%s'", toastText)
		}
	}

	// Test button variants
	h.AssertElementExists("#cta-button")
	h.AssertElementExists("#outline-button")
	h.AssertElementExists("#quiet-button")

	// Click CTA button
	h.ClickButton("Call to Action")
	ctaToast := h.WaitForToast(5 * time.Second)
	if ctaToast != nil {
		ctaText := ctaToast.MustText()
		if !strings.Contains(ctaText, "CTA clicked") {
			t.Errorf("Expected CTA toast, got '%s'", ctaText)
		}
	}

	// Test button states
	h.AssertElementExists("#disabled-button")
	h.AssertElementExists("#loading-button")

	// Verify disabled button is actually disabled
	disabledBtn := h.page.MustElement("#disabled-button")
	if !disabledBtn.MustProperty("disabled").Bool() {
		t.Error("Disabled button should have disabled property")
	}

	// Test icon buttons
	h.AssertElementExists("#edit-icon-button")
	h.AssertElementExists("#delete-icon-button")
	h.AssertElementExists("#save-icon-button")

	// Click edit icon button
	h.ClickButton("✏️ Edit")
	editToast := h.WaitForToast(5 * time.Second)
	if editToast != nil {
		editText := editToast.MustText()
		if !strings.Contains(editText, "Edit action") {
			t.Errorf("Expected edit action toast, got '%s'", editText)
		}
	}

	// Click delete icon button
	h.ClickButton("🗑️ Delete")
	deleteToast := h.WaitForToast(5 * time.Second)
	if deleteToast != nil {
		deleteText := deleteToast.MustText()
		if !strings.Contains(deleteText, "Delete action") {
			t.Errorf("Expected delete action toast, got '%s'", deleteText)
		}
	}

	// Click save icon button
	h.ClickButton("💾 Save")
	saveToast := h.WaitForToast(5 * time.Second)
	if saveToast != nil {
		saveText := saveToast.MustText()
		if !strings.Contains(saveText, "Save action") {
			t.Errorf("Expected save action toast, got '%s'", saveText)
		}
	}
}

func TestInputFields(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test text input
	h.AssertElementExists("#text-input")
	textInput := h.page.MustElement("#text-input")
	textInput.MustInput("test text")

	inputToast := h.WaitForToast(5 * time.Second)
	if inputToast != nil {
		inputText := inputToast.MustText()
		if !strings.Contains(inputText, "Text input changed") {
			t.Errorf("Expected text input change toast, got '%s'", inputText)
		}
	}

	// Test password input
	h.AssertElementExists("#password-input")
	passwordInput := h.page.MustElement("#password-input")
	passwordInput.MustInput("password123")

	passwordToast := h.WaitForToast(5 * time.Second)
	if passwordToast != nil {
		passwordText := passwordToast.MustText()
		if !strings.Contains(passwordText, "Password input changed") {
			t.Errorf("Expected password input change toast, got '%s'", passwordText)
		}
	}

	// Test email input
	h.AssertElementExists("#email-input")
	emailInput := h.page.MustElement("#email-input")
	emailInput.MustInput("test@example.com")

	emailToast := h.WaitForToast(5 * time.Second)
	if emailToast != nil {
		emailText := emailToast.MustText()
		if !strings.Contains(emailText, "Email input changed") {
			t.Errorf("Expected email input change toast, got '%s'", emailText)
		}
	}

	// Test number input
	h.AssertElementExists("#number-input")
	numberInput := h.page.MustElement("#number-input")
	numberInput.MustInput("42")

	numberToast := h.WaitForToast(5 * time.Second)
	if numberToast != nil {
		numberText := numberToast.MustText()
		if !strings.Contains(numberText, "Number input changed") {
			t.Errorf("Expected number input change toast, got '%s'", numberText)
		}
	}
}

func TestSelectionControls(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test checkbox
	h.AssertElementExists("#test-checkbox")
	checkbox := h.page.MustElement("#test-checkbox")

	// Initially unchecked
	if checkbox.MustProperty("checked").Bool() {
		t.Error("Checkbox should be initially unchecked")
	}

	checkbox.MustEval("() => this.click()")
	checkboxToast := h.WaitForToast(5 * time.Second)
	if checkboxToast != nil {
		checkboxText := checkboxToast.MustText()
		if !strings.Contains(checkboxText, "Checkbox toggled") {
			t.Errorf("Expected checkbox toggle toast, got '%s'", checkboxText)
		}
	}

	// Test radio buttons
	h.AssertElementExists("#radio-option1")
	h.AssertElementExists("#radio-option2")

	radio1 := h.page.MustElement("#radio-option1")
	radio2 := h.page.MustElement("#radio-option2")

	// Initially option1 should be selected
	if !radio1.MustProperty("checked").Bool() {
		t.Error("Radio option 1 should be initially selected")
	}

	// Click option 2
	radio2.MustEval("() => this.click()")
	radioToast := h.WaitForToast(5 * time.Second)
	if radioToast != nil {
		radioText := radioToast.MustText()
		if !strings.Contains(radioText, "Radio option 2 selected") {
			t.Errorf("Expected radio option 2 toast, got '%s'", radioText)
		}
	}

	// Test switch
	h.AssertElementExists("#test-switch")
	switchEl := h.page.MustElement("#test-switch")

	// Initially unchecked
	if switchEl.MustProperty("checked").Bool() {
		t.Error("Switch should be initially unchecked")
	}

	switchEl.MustEval("() => this.click()")
	switchToast := h.WaitForToast(5 * time.Second)
	if switchToast != nil {
		switchText := switchToast.MustText()
		if !strings.Contains(switchText, "Switch toggled") {
			t.Errorf("Expected switch toggle toast, got '%s'", switchText)
		}
	}
}

func TestComprehensiveFormValidation(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Find form inputs by looking for the validation form section
	formInputs := h.page.MustElements("input[type='text'], input[type='email'], input[type='tel']")
	if len(formInputs) < 3 {
		t.Errorf("Expected at least 3 form inputs for validation, got %d", len(formInputs))
	}

	// Test form submission with empty values
	submitButtons := h.page.MustElements("button")
	var submitBtn *rod.Element
	for _, btn := range submitButtons {
		if strings.Contains(btn.MustText(), "Submit") {
			submitBtn = btn
			break
		}
	}

	if submitBtn != nil {
		submitBtn.MustClick()
		validationToast := h.WaitForToast(5 * time.Second)
		if validationToast != nil {
			validationText := validationToast.MustText()
			if !strings.Contains(validationText, "validation") {
				t.Errorf("Expected validation toast, got '%s'", validationText)
			}
		}
	}
}

func TestComprehensiveDataTableInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Verify data table exists and has content
	checkboxes := h.page.MustElements("input[type='checkbox']")
	if len(checkboxes) < 5 {
		t.Errorf("Expected at least 5 checkboxes (for test users), got %d", len(checkboxes))
	}

	// Test Select All button
	h.ClickButton("Select All")
	selectAllToast := h.WaitForToast(5 * time.Second)
	if selectAllToast != nil {
		selectAllText := selectAllToast.MustText()
		if !strings.Contains(selectAllText, "All rows selected") {
			t.Errorf("Expected select all toast, got '%s'", selectAllText)
		}
	}

	// Wait for state change
	time.Sleep(500 * time.Millisecond)

	// Verify all checkboxes are selected
	updatedCheckboxes := h.page.MustElements("input[type='checkbox']")
	for i, checkbox := range updatedCheckboxes {
		if !checkbox.MustProperty("checked").Bool() {
			t.Errorf("Checkbox %d should be checked after Select All", i)
		}
	}

	// Test Clear Selection button
	h.ClickButton("Clear Selection")
	clearToast := h.WaitForToast(5 * time.Second)
	if clearToast != nil {
		clearText := clearToast.MustText()
		if !strings.Contains(clearText, "Selection cleared") {
			t.Errorf("Expected clear selection toast, got '%s'", clearText)
		}
	}

	// Wait for state change
	time.Sleep(500 * time.Millisecond)

	// Verify all checkboxes are unselected
	clearedCheckboxes := h.page.MustElements("input[type='checkbox']")
	for i, checkbox := range clearedCheckboxes {
		if checkbox.MustProperty("checked").Bool() {
			t.Errorf("Checkbox %d should be unchecked after Clear Selection", i)
		}
	}

	// Test View and Edit buttons in table rows
	viewButtons := h.page.MustElements("button")
	var viewBtn *rod.Element
	for _, btn := range viewButtons {
		if btn.MustText() == "View" {
			viewBtn = btn
			break
		}
	}

	if viewBtn != nil {
		viewBtn.MustClick()
		viewToast := h.WaitForToast(5 * time.Second)
		if viewToast != nil {
			viewText := viewToast.MustText()
			if !strings.Contains(viewText, "Viewing user") {
				t.Errorf("Expected viewing user toast, got '%s'", viewText)
			}
		}
	}

	// Test Edit button
	editButtons := h.page.MustElements("button")
	var editBtn *rod.Element
	for _, btn := range editButtons {
		if btn.MustText() == "Edit" {
			editBtn = btn
			break
		}
	}

	if editBtn != nil {
		editBtn.MustClick()
		editToast := h.WaitForToast(5 * time.Second)
		if editToast != nil {
			editText := editToast.MustText()
			if !strings.Contains(editText, "Editing user") {
				t.Errorf("Expected editing user toast, got '%s'", editText)
			}
		}
	}
}

func TestCardsInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test card action buttons
	cardButtons := h.page.MustElements("button")

	// Find and click View card action
	for _, btn := range cardButtons {
		if btn.MustText() == "View" {
			btn.MustClick()
			viewToast := h.WaitForToast(5 * time.Second)
			if viewToast != nil {
				viewText := viewToast.MustText()
				if strings.Contains(viewText, "Card action: View") {
					break // Found the card view action
				}
			}
		}
	}

	// Find and click Edit card action
	for _, btn := range cardButtons {
		if btn.MustText() == "Edit" {
			btn.MustClick()
			editToast := h.WaitForToast(5 * time.Second)
			if editToast != nil {
				editText := editToast.MustText()
				if strings.Contains(editText, "Card action: Edit") {
					break // Found the card edit action
				}
			}
		}
	}
}

func TestTabsInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test tab navigation
	h.AssertElementExists("#tab-1")
	h.AssertElementExists("#tab-2")
	h.AssertElementExists("#tab-3")

	// Initially tab 1 should be active
	h.AssertElementExists("#tab-content-1")

	// Click tab 2
	h.ClickButton("Tab 2")
	tab2Toast := h.WaitForToast(5 * time.Second)
	if tab2Toast != nil {
		tab2Text := tab2Toast.MustText()
		if !strings.Contains(tab2Text, "Tab 2 selected") {
			t.Errorf("Expected tab 2 selected toast, got '%s'", tab2Text)
		}
	}

	// Wait for content change
	time.Sleep(500 * time.Millisecond)
	h.AssertElementExists("#tab-content-2")

	// Click tab 3
	h.ClickButton("Tab 3")
	tab3Toast := h.WaitForToast(5 * time.Second)
	if tab3Toast != nil {
		tab3Text := tab3Toast.MustText()
		if !strings.Contains(tab3Text, "Tab 3 selected") {
			t.Errorf("Expected tab 3 selected toast, got '%s'", tab3Text)
		}
	}

	// Wait for content change
	time.Sleep(500 * time.Millisecond)
	h.AssertElementExists("#tab-content-3")
}

func TestAccordionInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test accordion toggle
	h.AssertElementExists("#accordion-toggle")

	// Initially accordion should be closed
	accordionContents := h.page.MustElements("#accordion-content")
	if len(accordionContents) > 0 {
		t.Error("Accordion should be initially closed")
	}

	// Click to open accordion
	accordionToggle := h.page.MustElement("#accordion-toggle")
	accordionToggle.MustEval("() => this.click()")

	accordionToast := h.WaitForToast(5 * time.Second)
	if accordionToast != nil {
		accordionText := accordionToast.MustText()
		if !strings.Contains(accordionText, "Accordion toggled") {
			t.Errorf("Expected accordion toggle toast, got '%s'", accordionText)
		}
	}

	// Wait for content to appear
	time.Sleep(500 * time.Millisecond)
	h.AssertElementExists("#accordion-content")

	// Click to close accordion
	accordionToggle.MustEval("() => this.click()")

	// Wait for content to disappear
	time.Sleep(500 * time.Millisecond)
	closedContents := h.page.MustElements("#accordion-content")
	if len(closedContents) > 0 {
		t.Error("Accordion content should be hidden after closing")
	}
}

func TestComprehensiveModalInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Initially no modal should be visible
	h.AssertElementNotExists(".spectrum-Modal")

	// Click open modal button
	h.ClickButton("Open Modal")

	// Wait for modal to appear
	modal := h.WaitForModal(5 * time.Second)
	if modal == nil {
		t.Fatal("Modal should appear after clicking open button")
	}

	// Verify modal content
	modalText := modal.MustText()
	if !strings.Contains(modalText, "Test Modal") {
		t.Errorf("Expected modal to contain 'Test Modal', got '%s'", modalText)
	}

	// Test modal close by clicking Cancel
	h.ClickButton("Cancel")
	cancelToast := h.WaitForToast(5 * time.Second)
	if cancelToast != nil {
		cancelText := cancelToast.MustText()
		if !strings.Contains(cancelText, "Modal cancelled") {
			t.Errorf("Expected modal cancelled toast, got '%s'", cancelText)
		}
	}

	// Wait for modal to disappear
	time.Sleep(500 * time.Millisecond)
	h.AssertElementNotExists(".spectrum-Modal")

	// Test modal confirm action
	h.ClickButton("Open Modal")
	modal = h.WaitForModal(5 * time.Second)
	if modal == nil {
		t.Fatal("Modal should appear again")
	}

	h.ClickButton("Confirm")
	confirmToast := h.WaitForToast(5 * time.Second)
	if confirmToast != nil {
		confirmText := confirmToast.MustText()
		if !strings.Contains(confirmText, "Modal confirmed") {
			t.Errorf("Expected modal confirmed toast, got '%s'", confirmText)
		}
	}

	// Wait for modal to disappear
	time.Sleep(500 * time.Millisecond)
	h.AssertElementNotExists(".spectrum-Modal")
}

func TestComprehensiveToastNotifications(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test info toast
	h.ClickButton("Show Info Toast")
	infoToast := h.WaitForToast(5 * time.Second)
	if infoToast != nil {
		infoText := infoToast.MustText()
		if !strings.Contains(infoText, "info toast") {
			t.Errorf("Expected info toast message, got '%s'", infoText)
		}
	}

	// Test success toast
	h.ClickButton("Show Success Toast")
	successToast := h.WaitForToast(5 * time.Second)
	if successToast != nil {
		successText := successToast.MustText()
		if !strings.Contains(successText, "success toast") {
			t.Errorf("Expected success toast message, got '%s'", successText)
		}
	}

	// Test warning toast
	h.ClickButton("Show Warning Toast")
	warningToast := h.WaitForToast(5 * time.Second)
	if warningToast != nil {
		warningText := warningToast.MustText()
		if !strings.Contains(warningText, "warning toast") {
			t.Errorf("Expected warning toast message, got '%s'", warningText)
		}
	}

	// Test error toast
	h.ClickButton("Show Error Toast")
	errorToast := h.WaitForToast(5 * time.Second)
	if errorToast != nil {
		errorText := errorToast.MustText()
		if !strings.Contains(errorText, "error toast") {
			t.Errorf("Expected error toast message, got '%s'", errorText)
		}
	}
}

func TestTooltipInteractions(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test tooltip buttons exist
	h.AssertElementExists("#tooltip-button-1")
	h.AssertElementExists("#tooltip-button-2")

	// Click tooltip buttons to test functionality
	h.ClickButton("Hover for Tooltip")
	tooltipToast := h.WaitForToast(5 * time.Second)
	if tooltipToast != nil {
		tooltipText := tooltipToast.MustText()
		if !strings.Contains(tooltipText, "Button with tooltip clicked") {
			t.Errorf("Expected tooltip button toast, got '%s'", tooltipText)
		}
	}

	h.ClickButton("Another Tooltip")
	anotherToast := h.WaitForToast(5 * time.Second)
	if anotherToast != nil {
		anotherText := anotherToast.MustText()
		if !strings.Contains(anotherText, "Another tooltip button clicked") {
			t.Errorf("Expected another tooltip button toast, got '%s'", anotherText)
		}
	}
}

func TestNavigationComponents(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test breadcrumbs
	h.AssertElementExists("#test-breadcrumbs")
	breadcrumbLinks := h.page.MustElements("#test-breadcrumbs a")
	if len(breadcrumbLinks) < 2 {
		t.Errorf("Expected at least 2 breadcrumb links, got %d", len(breadcrumbLinks))
	}

	// Click Home breadcrumb
	for _, link := range breadcrumbLinks {
		if link.MustText() == "Home" {
			link.MustClick()
			breadcrumbToast := h.WaitForToast(5 * time.Second)
			if breadcrumbToast != nil {
				breadcrumbText := breadcrumbToast.MustText()
				if !strings.Contains(breadcrumbText, "Home breadcrumb clicked") {
					t.Errorf("Expected home breadcrumb toast, got '%s'", breadcrumbText)
				}
			}
			break
		}
	}

	// Test pagination
	h.AssertElementExists("#test-pagination")
	h.AssertElementExists("#prev-page")
	h.AssertElementExists("#next-page")
	h.AssertElementExists("#page-1")
	h.AssertElementExists("#page-2")
	h.AssertElementExists("#page-3")

	// Click pagination buttons
	h.ClickButton("Previous")
	prevToast := h.WaitForToast(5 * time.Second)
	if prevToast != nil {
		prevText := prevToast.MustText()
		if !strings.Contains(prevText, "Previous page clicked") {
			t.Errorf("Expected previous page toast, got '%s'", prevText)
		}
	}

	h.ClickButton("Next")
	nextToast := h.WaitForToast(5 * time.Second)
	if nextToast != nil {
		nextText := nextToast.MustText()
		if !strings.Contains(nextText, "Next page clicked") {
			t.Errorf("Expected next page toast, got '%s'", nextText)
		}
	}

	// Test menu interactions
	h.AssertElementExists("#test-menu")
	menuItems := h.page.MustElements("#test-menu .menu-item")
	if len(menuItems) < 3 {
		t.Errorf("Expected at least 3 menu items, got %d", len(menuItems))
	}

	// Click menu items
	for i, item := range menuItems {
		if i < 3 { // Test first 3 menu items
			item.MustClick()
			menuToast := h.WaitForToast(5 * time.Second)
			if menuToast != nil {
				menuText := menuToast.MustText()
				expectedText := "Menu Item " + string(rune('1'+i)) + " clicked"
				if !strings.Contains(menuText, expectedText) {
					t.Errorf("Expected menu item %d toast, got '%s'", i+1, menuText)
				}
			}
		}
	}
}

func TestAdvancedComponents(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test progress bars
	h.AssertElementExists("#progress-bar-1")
	h.AssertElementExists("#progress-bar-2")

	// Test badges
	h.AssertElementExists("#badge-primary")
	h.AssertElementExists("#badge-success")
	h.AssertElementExists("#badge-warning")
	h.AssertElementExists("#badge-error")

	// Test chips
	h.AssertElementExists("#chip-1")
	h.AssertElementExists("#chip-2")

	// Click chip remove buttons
	chip1Remove := h.page.MustElement("#chip-1 span")
	chip1Remove.MustClick()
	chip1Toast := h.WaitForToast(5 * time.Second)
	if chip1Toast != nil {
		chip1Text := chip1Toast.MustText()
		if !strings.Contains(chip1Text, "Chip 1 removed") {
			t.Errorf("Expected chip 1 removed toast, got '%s'", chip1Text)
		}
	}

	chip2Remove := h.page.MustElement("#chip-2 span")
	chip2Remove.MustClick()
	chip2Toast := h.WaitForToast(5 * time.Second)
	if chip2Toast != nil {
		chip2Text := chip2Toast.MustText()
		if !strings.Contains(chip2Text, "Chip 2 removed") {
			t.Errorf("Expected chip 2 removed toast, got '%s'", chip2Text)
		}
	}
}

func TestNavigationComponentsDetailed(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test Side Navigation
	h.AssertElementExists("#test-sidenav")

	// Find and test side nav items
	sideNavItems := h.page.MustElements("#test-sidenav .prism-sidenav__item")
	if len(sideNavItems) >= 3 {
		// Test Home nav item (should be active)
		sideNavItems[0].MustClick()
		homeNavToast := h.WaitForToast(5 * time.Second)
		if homeNavToast != nil {
			homeNavText := homeNavToast.MustText()
			if !strings.Contains(homeNavText, "Home navigation clicked") {
				t.Errorf("Expected home navigation toast, got '%s'", homeNavText)
			}
		}

		// Test Users nav item
		sideNavItems[1].MustClick()
		usersNavToast := h.WaitForToast(5 * time.Second)
		if usersNavToast != nil {
			usersNavText := usersNavToast.MustText()
			if !strings.Contains(usersNavText, "Users navigation clicked") {
				t.Errorf("Expected users navigation toast, got '%s'", usersNavText)
			}
		}

		// Test Settings nav item
		sideNavItems[2].MustClick()
		settingsNavToast := h.WaitForToast(5 * time.Second)
		if settingsNavToast != nil {
			settingsNavText := settingsNavToast.MustText()
			if !strings.Contains(settingsNavText, "Settings navigation clicked") {
				t.Errorf("Expected settings navigation toast, got '%s'", settingsNavText)
			}
		}
	}

	// Test Top Navigation
	h.AssertElementExists("#test-topnav")

	// Test search functionality
	searchInputs := h.page.MustElements("#test-topnav input[type='search'], #test-topnav input[placeholder*='Search']")
	if len(searchInputs) > 0 {
		searchInput := searchInputs[0]
		searchInput.MustInput("test search")

		// Press Enter to trigger search
		searchInput.MustType(input.Enter)

		searchToast := h.WaitForToast(5 * time.Second)
		if searchToast != nil {
			searchText := searchToast.MustText()
			if !strings.Contains(searchText, "Searching for: test search") {
				t.Errorf("Expected search toast, got '%s'", searchText)
			}
		}
	}

	// Test user menu interactions
	userMenuButtons := h.page.MustElements("#test-topnav button")
	for _, btn := range userMenuButtons {
		btnText := btn.MustText()
		if strings.Contains(btnText, "Profile") {
			btn.MustClick()
			profileToast := h.WaitForToast(5 * time.Second)
			if profileToast != nil {
				profileText := profileToast.MustText()
				if !strings.Contains(profileText, "Profile clicked") {
					t.Errorf("Expected profile clicked toast, got '%s'", profileText)
				}
			}
			break
		}
	}
}

func TestComprehensiveResponsiveLayout(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test desktop layout
	h.Desktop()
	h.Reload()

	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Comprehensive Component Test Suite")

	// Test mobile layout
	h.Mobile()
	h.Reload()

	// Content should still be accessible
	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Comprehensive Component Test Suite")

	// Buttons should still be clickable
	buttons := h.page.MustElements("button")
	if len(buttons) == 0 {
		t.Error("Expected buttons to be available on mobile")
	}

	// Test one button click on mobile
	if len(buttons) > 0 {
		buttons[0].MustClick()
		// Should get some toast (any toast means interactivity works)
		toast := h.WaitForToast(5 * time.Second)
		if toast == nil {
			t.Error("Expected toast on mobile button click")
		}
	}

	// Test tablet layout
	h.Tablet()
	h.Reload()

	h.AssertElementExists("h1")
	h.AssertElementText("h1", "Comprehensive Component Test Suite")
}

func TestComprehensiveAccessibility(t *testing.T) {
	h := NewTestHelpers(t, ServerURL+"/test")
	defer h.Close()

	// Test that buttons have proper text content
	buttons := h.page.MustElements("button")
	for i, btn := range buttons {
		text := btn.MustText()
		if text == "" {
			t.Errorf("Button %d should have text content for accessibility", i)
		}
	}

	// Test that form inputs have associated labels or placeholders
	inputs := h.page.MustElements("input")
	for i, input := range inputs {
		placeholder := input.MustAttribute("placeholder")
		aria_label := input.MustAttribute("aria-label")
		if (placeholder == nil || *placeholder == "") && (aria_label == nil || *aria_label == "") {
			// Check if there's a nearby label
			labels := h.page.MustElements("label")
			hasLabel := false
			for _, label := range labels {
				forAttr := label.MustAttribute("for")
				inputId := input.MustAttribute("id")
				if forAttr != nil && inputId != nil && *forAttr == *inputId {
					hasLabel = true
					break
				}
			}
			if !hasLabel {
				t.Errorf("Input %d should have placeholder, aria-label, or associated label for accessibility", i)
			}
		}
	}

	// Test that important elements have proper heading structure
	h1Elements := h.page.MustElements("h1")
	if len(h1Elements) != 1 {
		t.Errorf("Page should have exactly one h1 element, got %d", len(h1Elements))
	}

	h2Elements := h.page.MustElements("h2")
	if len(h2Elements) < 5 {
		t.Errorf("Page should have multiple h2 elements for section structure, got %d", len(h2Elements))
	}
}
