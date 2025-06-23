package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IfAuthorized conditionally renders content based on user permissions
func IfAuthorized(user *User, permission string, content app.UI) app.UI {
	if user == nil || !user.LoggedIn {
		return app.Div() // Return empty div if not logged in
	}

	if HasPermission(user, permission) {
		return content
	}

	return app.Div() // Return empty div if not authorized
}

// IfLoggedIn conditionally renders content for logged-in users
func IfLoggedIn(user *User, content app.UI) app.UI {
	if user != nil && user.LoggedIn {
		return content
	}
	return app.Div()
}

// IfLoggedOut conditionally renders content for logged-out users
func IfLoggedOut(user *User, content app.UI) app.UI {
	if user == nil || !user.LoggedIn {
		return content
	}
	return app.Div()
}

// HasPermission checks if user has a specific permission
func HasPermission(user *User, permission string) bool {
	if user == nil || !user.LoggedIn {
		return false
	}

	// For now, basic role-based permissions
	// In real implementation, this would check against user roles/permissions
	switch permission {
	case "admin":
		return user.IsAdmin()
	case "create":
		return user.CanCreate()
	case "edit":
		return user.CanEdit()
	case "delete":
		return user.CanDelete()
	case "view":
		return true // All logged-in users can view
	default:
		return false
	}
}

// RequireAuth redirects to login if not authenticated
func RequireAuth(user *User, ctx app.Context) bool {
	if user == nil || !user.LoggedIn {
		ctx.Navigate("/login")
		return false
	}
	return true
}

// RequirePermission redirects to unauthorized page if permission denied
func RequirePermission(user *User, permission string, ctx app.Context) bool {
	if !HasPermission(user, permission) {
		ctx.Navigate("/unauthorized")
		return false
	}
	return true
}

// GetCurrentUser retrieves user from context/state
func GetCurrentUser(ctx app.Context) *User {
	var user User
	ctx.GetState("current-user", &user)
	if user.ID != "" {
		return &user
	}
	return nil
}

// SetCurrentUser stores user in context/state
func SetCurrentUser(ctx app.Context, user *User) {
	ctx.SetState("current-user", user).Persist().Broadcast()
}

// ClearCurrentUser removes user from context/state
func ClearCurrentUser(ctx app.Context) {
	var emptyUser *User
	ctx.SetState("current-user", emptyUser).Persist().Broadcast()
}