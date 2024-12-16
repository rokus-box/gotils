package gotils

import "regexp"

// ExtractEmails extracts all email addresses from the provided biography string.
func ExtractEmails(biography string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(biography, -1)

	return emails
}
