package prism

import (
	"strings"
	"time"
)

type User struct {
	ID       string
	Name     string
	Email    string
	LoggedIn bool
	Avatar   string
	Role     string
	Permissions []string
	CreatedAt   time.Time
	LastLoginAt time.Time
}

// Permission check methods
func (u *User) IsAdmin() bool {
	return u.Role == "admin" || u.HasPermission("admin")
}

func (u *User) CanCreate() bool {
	return u.IsAdmin() || u.HasPermission("create")
}

func (u *User) CanEdit() bool {
	return u.IsAdmin() || u.HasPermission("edit")
}

func (u *User) CanDelete() bool {
	return u.IsAdmin() || u.HasPermission("delete")
}

func (u *User) CanView() bool {
	return u.LoggedIn // All logged-in users can view
}

func (u *User) HasPermission(permission string) bool {
	for _, p := range u.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func (u *User) HasRole(role string) bool {
	return u.Role == role
}

// Display methods
func (u *User) GetDisplayName() string {
	if u.Name != "" {
		return u.Name
	}
	return u.Email
}

func (u *User) GetInitials() string {
	if u.Name == "" {
		return "U"
	}
	
	parts := strings.Split(u.Name, " ")
	if len(parts) >= 2 {
		return string(parts[0][0]) + string(parts[1][0])
	}
	return string(u.Name[0])
}

func (u *User) IsActive() bool {
	return u.LoggedIn
}
