package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

var browser *rod.Browser

// TestHelpers provides common testing utilities
type TestHelpers struct {
	page *rod.Page
	t    *testing.T
}

// NewTestHelpers creates a new test helper instance
func NewTestHelpers(t *testing.T, url string) *TestHelpers {
	page := browser.MustPage(url)
	page.MustWaitLoad()

	return &TestHelpers{
		page: page,
		t:    t,
	}
}

// Close cleans up the test page
func (h *TestHelpers) Close() {
	if h.page != nil {
		h.page.Close()
	}
}

// WaitForElement waits for an element to appear with timeout
func (h *TestHelpers) WaitForElement(selector string, timeout time.Duration) *rod.Element {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	element, err := h.page.Context(ctx).Element(selector)
	if err != nil {
		h.t.Fatalf("Element %s not found within %v: %v", selector, timeout, err)
	}
	return element
}

// WaitForElementText waits for an element with specific text
func (h *TestHelpers) WaitForElementText(selector, expectedText string, timeout time.Duration) *rod.Element {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			h.t.Fatalf("Element %s with text '%s' not found within %v", selector, expectedText, timeout)
		default:
			elements := h.page.MustElements(selector)
			for _, el := range elements {
				if el.MustText() == expectedText {
					return el
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// AssertElementExists checks if an element exists
func (h *TestHelpers) AssertElementExists(selector string) {
	elements := h.page.MustElements(selector)
	if len(elements) == 0 {
		h.t.Errorf("Expected element %s to exist", selector)
	}
}

// AssertElementNotExists checks if an element does not exist
func (h *TestHelpers) AssertElementNotExists(selector string) {
	elements := h.page.MustElements(selector)
	if len(elements) > 0 {
		h.t.Errorf("Expected element %s to not exist", selector)
	}
}

// AssertElementText checks if an element has expected text
func (h *TestHelpers) AssertElementText(selector, expectedText string) {
	element := h.page.MustElement(selector)
	actualText := element.MustText()
	if actualText != expectedText {
		h.t.Errorf("Expected element %s to have text '%s', got '%s'", selector, expectedText, actualText)
	}
}

// AssertElementVisible checks if an element is visible
func (h *TestHelpers) AssertElementVisible(selector string) {
	element := h.page.MustElement(selector)
	if !element.MustVisible() {
		h.t.Errorf("Expected element %s to be visible", selector)
	}
}

// AssertElementHidden checks if an element is hidden
func (h *TestHelpers) AssertElementHidden(selector string) {
	elements := h.page.MustElements(selector)
	if len(elements) > 0 && elements[0].MustVisible() {
		h.t.Errorf("Expected element %s to be hidden", selector)
	}
}

// ClickButton clicks a button by text content
func (h *TestHelpers) ClickButton(buttonText string) {
	button := h.WaitForElementText("button, .spectrum-Button, sp-button", buttonText, 5*time.Second)
	button.MustEval("() => this.click()")
}

// FillInput fills an input field
func (h *TestHelpers) FillInput(selector, value string) {
	input := h.page.MustElement(selector)
	input.MustSelectAllText()
	input.MustInput(value)
}

// FillInputByLabel fills an input field by its label
func (h *TestHelpers) FillInputByLabel(labelText, value string) {
	label := h.WaitForElementText("label, .spectrum-FieldLabel", labelText, 5*time.Second)

	// Find associated input
	var input *rod.Element

	// Try to find input by for attribute
	if forAttr := label.MustAttribute("for"); forAttr != nil {
		input = h.page.MustElement("#" + *forAttr)
	} else {
		// Find input within or after the label
		inputs := h.page.MustElements("input")
		for _, inp := range inputs {
			// Simple heuristic: find input near the label
			if inp.MustVisible() {
				input = inp
				break
			}
		}
	}

	if input == nil {
		h.t.Fatalf("Could not find input for label '%s'", labelText)
	}

	input.MustSelectAllText()
	input.MustInput(value)
}

// WaitForModal waits for a modal to appear
func (h *TestHelpers) WaitForModal(timeout time.Duration) *rod.Element {
	return h.WaitForElement(".spectrum-Modal, .spectrum-Dialog", timeout)
}

// WaitForToast waits for a toast notification to appear
func (h *TestHelpers) WaitForToast(timeout time.Duration) *rod.Element {
	return h.WaitForElement(".spectrum-Toast, .toast-container", timeout)
}

// AssertPageTitle checks the page title
func (h *TestHelpers) AssertPageTitle(expectedTitle string) {
	actualTitle := h.page.MustInfo().Title
	if actualTitle != expectedTitle {
		h.t.Errorf("Expected page title '%s', got '%s'", expectedTitle, actualTitle)
	}
}

// AssertURL checks the current URL
func (h *TestHelpers) AssertURL(expectedURL string) {
	actualURL := h.page.MustInfo().URL
	if actualURL != expectedURL {
		h.t.Errorf("Expected URL '%s', got '%s'", expectedURL, actualURL)
	}
}

// Navigate navigates to a new URL
func (h *TestHelpers) Navigate(url string) {
	h.page.MustNavigate(url)
	h.page.MustWaitLoad()
}

// Reload reloads the current page
func (h *TestHelpers) Reload() {
	h.page.MustReload()
	h.page.MustWaitLoad()
}

// Screenshot takes a screenshot for debugging
func (h *TestHelpers) Screenshot(filename string) {
	data := h.page.MustScreenshot()
	if err := utils.OutputFile(filename, data); err != nil {
		h.t.Logf("Failed to save screenshot %s: %v", filename, err)
	} else {
		h.t.Logf("Screenshot saved: %s", filename)
	}
}

// ExecuteJS executes JavaScript and returns the result
func (h *TestHelpers) ExecuteJS(js string) gson.JSON {
	return h.page.MustEval(js)
}

// WaitForNetworkIdle waits for network requests to finish
func (h *TestHelpers) WaitForNetworkIdle(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	h.page.Context(ctx).MustWaitRequestIdle()
}

// SetViewport sets the viewport size
func (h *TestHelpers) SetViewport(width, height int) {
	h.page.MustSetViewport(width, height, 1, false)
}

// Mobile sets mobile viewport
func (h *TestHelpers) Mobile() {
	h.SetViewport(375, 667)
}

// Desktop sets desktop viewport
func (h *TestHelpers) Desktop() {
	h.SetViewport(1200, 800)
}

// Tablet sets tablet viewport
func (h *TestHelpers) Tablet() {
	h.SetViewport(768, 1024)
}
