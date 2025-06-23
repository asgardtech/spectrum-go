package prism

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// DebugDump displays a JSON representation of any value - useful for development
func DebugDump(v any) {
	DebugDumpWithLabel("Debug", v)
}

// DebugDumpWithLabel displays a JSON representation with a custom label
func DebugDumpWithLabel(label string, v any) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("[%s] Error marshaling to JSON: %v", label, err)
		return
	}
	log.Printf("[%s] %s", label, string(jsonBytes))
}

// DebugLogf logs formatted debug messages (only in development)
func DebugLogf(format string, args ...any) {
	if IsDebugMode() {
		log.Printf("[DEBUG] "+format, args...)
	}
}

// IsDebugMode checks if we're in debug/development mode
func IsDebugMode() bool {
	// In go-app, we can check if we're running in the browser vs server
	// or use build tags for more sophisticated debug detection
	return app.IsClient || app.IsServer
}

// DumpState logs the current state of a component for debugging
func DumpState(component app.Compo, label string) {
	if IsDebugMode() {
		DebugDumpWithLabel(fmt.Sprintf("State[%s]", label), component)
	}
}

// LogError logs errors with context
func LogError(err error, context string) {
	if err != nil {
		log.Printf("[ERROR] %s: %v", context, err)
	}
}

// LogWarning logs warnings
func LogWarning(message string, context string) {
	log.Printf("[WARNING] %s: %s", context, message)
}

// LogInfo logs informational messages
func LogInfo(message string, context string) {
	log.Printf("[INFO] %s: %s", context, message)
}