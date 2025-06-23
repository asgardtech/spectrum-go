package prism

import (
	"fmt"
	"strings"
	"time"
)


// FormatDate formats a time.Time for display
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	return t.Format("Jan 2, 2006")
}

// FormatDateTime formats a time.Time with date and time
func FormatDateTime(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	return t.Format("Jan 2, 2006 at 3:04 PM")
}

// FormatTimeAgo formats a time.Time as "X ago" relative to now
func FormatTimeAgo(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	
	now := time.Now()
	diff := now.Sub(t)
	
	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if diff < 24*time.Hour {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	} else {
		return FormatDate(t)
	}
}

// FormatCurrency formats a float64 as currency
func FormatCurrency(amount float64, currency string) string {
	if currency == "" {
		currency = "USD"
	}
	
	switch currency {
	case "USD":
		return fmt.Sprintf("$%.2f", amount)
	case "EUR":
		return fmt.Sprintf("€%.2f", amount)
	case "GBP":
		return fmt.Sprintf("£%.2f", amount)
	default:
		return fmt.Sprintf("%.2f %s", amount, currency)
	}
}

// FormatNumber formats a number with commas
func FormatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}
	
	str := fmt.Sprintf("%d", n)
	result := ""
	
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}
	
	return result
}

// FormatFileSize formats bytes as human readable size
func FormatFileSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// TruncateText truncates text to maxLength and adds ellipsis
func TruncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength-3] + "..."
}

// FormatList formats a slice of strings as a comma-separated list
func FormatList(items []string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 2 {
		return items[0] + " and " + items[1]
	}
	
	return strings.Join(items[:len(items)-1], ", ") + ", and " + items[len(items)-1]
}

// Pluralize returns the singular or plural form based on count
func Pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

// FormatPercentage formats a float as a percentage
func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.1f%%", value*100)
}