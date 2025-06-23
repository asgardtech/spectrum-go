package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

const (
	ServerPort  = "9090"
	ServerURL   = "http://localhost:" + ServerPort
	TestTimeout = 30 * time.Second
)

var (
	serverCmd *exec.Cmd
)

func TestMain(m *testing.M) {
	// Setup
	if err := setupServer(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}

	if err := setupBrowser(); err != nil {
		fmt.Printf("Failed to setup browser: %v\n", err)
		cleanup()
		os.Exit(1)
	}

	// Wait for server to be ready
	if err := waitForServer(); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		cleanup()
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Cleanup
	cleanup()
	os.Exit(code)
}

func setupServer() error {
	// Build the demo app first
	buildCmd := exec.Command("go", "build", "-o", "demo-server", ".")
	buildCmd.Dir = "../../demo"
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build demo app: %v", err)
	}

	// Start the server
	serverCmd = exec.Command("./demo-server")
	serverCmd.Dir = "../../demo"
	serverCmd.Env = append(os.Environ(), "PORT="+ServerPort)

	if err := serverCmd.Start(); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}

func setupBrowser() error {
	// Launch browser in headless mode for CI
	l := launcher.New().Headless(true).NoSandbox(true)
	url, err := l.Launch()
	if err != nil {
		return fmt.Errorf("failed to launch browser: %v", err)
	}

	browser = rod.New().ControlURL(url)
	if err := browser.Connect(); err != nil {
		return fmt.Errorf("failed to connect to browser: %v", err)
	}

	return nil
}

func waitForServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for server to start")
		default:
			page := browser.MustPage()

			err := page.Navigate(ServerURL)
			if err == nil {
				page.Close()
				return nil
			}

			page.Close()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func cleanup() {
	if browser != nil {
		browser.Close()
	}

	if serverCmd != nil && serverCmd.Process != nil {
		serverCmd.Process.Signal(syscall.SIGTERM)
		serverCmd.Wait()
	}
}

func TestHomePage(t *testing.T) {
	page := browser.MustPage(ServerURL)
	defer page.Close()

	// Wait for page to load
	page.MustWaitLoad()

	// Check if the page title is set correctly
	title := page.MustInfo().Title
	if title != "Home" {
		t.Errorf("Expected page title 'Home', got '%s'", title)
	}

	// Check if the main heading exists
	h1 := page.MustElement("h1")
	text := h1.MustText()
	if text != "Home" {
		t.Errorf("Expected h1 text 'Home', got '%s'", text)
	}

	// Check if the top navigation is visible (when logged in)
	// For now, we expect it to be hidden since no user is logged in
	topNavs := page.MustElements(".spectrum-TopNav")
	if len(topNavs) > 0 {
		// Top nav exists, check if it's visible
		visible := topNavs[0].MustVisible()
		t.Logf("Top navigation visible: %v", visible)
	}
}

func TestLoginPage(t *testing.T) {
	page := browser.MustPage(ServerURL + "/login")
	defer page.Close()

	// Wait for page to load
	page.MustWaitLoad()

	// Check if the login page loaded
	title := page.MustInfo().Title
	if title != "Login" {
		t.Errorf("Expected page title 'Login', got '%s'", title)
	}

	// Check for login form elements
	emailFields := page.MustElements("sp-textfield")
	if len(emailFields) < 2 {
		t.Errorf("Expected at least 2 input fields (email, password), got %d", len(emailFields))
	}

	// Check for field labels
	labels := page.MustElements("sp-field-label")
	if len(labels) < 2 {
		t.Errorf("Expected at least 2 field labels, got %d", len(labels))
	}

	// Verify email label
	emailLabel := labels[0].MustText()
	if emailLabel != "Email" {
		t.Errorf("Expected first label to be 'Email', got '%s'", emailLabel)
	}

	// Verify password label
	passwordLabel := labels[1].MustText()
	if passwordLabel != "Password" {
		t.Errorf("Expected second label to be 'Password', got '%s'", passwordLabel)
	}
}

func TestPageNavigation(t *testing.T) {
	page := browser.MustPage(ServerURL)
	defer page.Close()

	// Start on home page
	page.MustWaitLoad()
	currentURL := page.MustInfo().URL
	if currentURL != ServerURL+"/" {
		t.Errorf("Expected to be on home page, got '%s'", currentURL)
	}

	// Navigate to login page
	page.MustNavigate(ServerURL + "/login")
	page.MustWaitLoad()

	currentURL = page.MustInfo().URL
	if currentURL != ServerURL+"/login" {
		t.Errorf("Expected to be on login page, got '%s'", currentURL)
	}

	// Navigate back to home
	page.MustNavigate(ServerURL + "/")
	page.MustWaitLoad()

	currentURL = page.MustInfo().URL
	if currentURL != ServerURL+"/" {
		t.Errorf("Expected to be back on home page, got '%s'", currentURL)
	}
}

func TestBasicResponsiveLayout(t *testing.T) {
	page := browser.MustPage(ServerURL)
	defer page.Close()

	page.MustWaitLoad()

	// Test desktop viewport
	page.MustSetViewport(1200, 800, 1, false)
	page.MustWaitLoad()

	// Check if content is properly laid out
	body := page.MustElement("body")
	bodyWidth := body.MustEval("() => this.offsetWidth").Int()
	if bodyWidth < 1000 {
		t.Errorf("Expected body width to be at least 1000px on desktop, got %d", bodyWidth)
	}

	// Test mobile viewport
	page.MustSetViewport(375, 667, 1, false)
	page.MustWaitLoad()

	// Content should still be accessible
	h1 := page.MustElement("h1")
	text := h1.MustText()
	if text != "Home" {
		t.Errorf("Expected h1 text 'Home' on mobile, got '%s'", text)
	}
}

// Helper function to create a test page
func createTestPage(t *testing.T, url string) *rod.Page {
	ctx, cancel := context.WithTimeout(context.Background(), TestTimeout)
	defer cancel()

	page := browser.Context(ctx).MustPage()

	if err := page.Navigate(url); err != nil {
		page.Close()
		t.Fatalf("Failed to navigate to %s: %v", url, err)
	}

	return page
}

// Helper function to wait for element with timeout
func waitForElement(page *rod.Page, selector string, timeout time.Duration) (*rod.Element, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return page.Context(ctx).Element(selector)
}
