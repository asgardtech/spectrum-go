package prism

import (
	"reflect"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Pointer helper functions for safe nil handling

// StringPtr returns a pointer to the string value
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the int value
func IntPtr(i int) *int {
	return &i
}

// BoolPtr returns a pointer to the bool value
func BoolPtr(b bool) *bool {
	return &b
}

// SafeString safely dereferences a string pointer
func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SafeInt safely dereferences an int pointer
func SafeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// SafeBool safely dereferences a bool pointer
func SafeBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// Conditional rendering helpers

// If conditionally renders content
func If(condition bool, content app.UI) app.UI {
	if condition {
		return content
	}
	return app.Div()
}

// IfElse conditionally renders one of two content options
func IfElse(condition bool, ifContent, elseContent app.UI) app.UI {
	if condition {
		return ifContent
	}
	return elseContent
}

// Unless renders content if condition is false
func Unless(condition bool, content app.UI) app.UI {
	return If(!condition, content)
}

// IfElseString returns one string or another based on condition
func IfElseString(condition bool, ifValue, elseValue string) string {
	if condition {
		return ifValue
	}
	return elseValue
}

// ClassNames conditionally combines CSS class names
func ClassNames(classes ...string) string {
	var result []string
	for _, class := range classes {
		if strings.TrimSpace(class) != "" {
			result = append(result, strings.TrimSpace(class))
		}
	}
	return strings.Join(result, " ")
}

// ConditionalClass adds a class if condition is true
func ConditionalClass(condition bool, className string) string {
	if condition {
		return className
	}
	return ""
}

// Validation helpers

// IsEmpty checks if a value is considered empty
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// IsNotEmpty checks if a value is not empty
func IsNotEmpty(value interface{}) bool {
	return !IsEmpty(value)
}

// Contains checks if a slice contains a value
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Filter filters a slice based on a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map transforms a slice using a mapping function
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// Find finds the first item in a slice that matches the predicate
func Find[T any](slice []T, predicate func(T) bool) *T {
	for _, item := range slice {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

// FindIndex finds the index of the first item that matches the predicate
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, item := range slice {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// Unique removes duplicate items from a slice
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	
	return result
}

// Reverse reverses a slice
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, item := range slice {
		result[len(slice)-1-i] = item
	}
	return result
}

// Min returns the minimum of two ordered values
func Min[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two ordered values
func Max[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Coalesce returns the first non-empty value
func Coalesce[T comparable](values ...T) T {
	var zero T
	for _, value := range values {
		if value != zero {
			return value
		}
	}
	return zero
}