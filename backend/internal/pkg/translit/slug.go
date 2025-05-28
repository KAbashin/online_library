package translit

import (
	"regexp"
	"strings"
)

// Удаление неразрешённых символов для slug
var nonAlnumRegex = regexp.MustCompile(`[^a-z0-9\-]+`)
var hyphenRegex = regexp.MustCompile(`-{2,}`)

// ToSlug — транслитерация и очистка строки для использования в URL-slug
func ToSlug(input string) string {
	input = ToLatin(input)
	input = strings.ToLower(input)

	var result strings.Builder
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result.WriteRune(r)
		} else if r == ' ' || r == '-' || r == '_' {
			result.WriteRune('-')
		}
	}

	slug := result.String()
	slug = nonAlnumRegex.ReplaceAllString(slug, "-")
	slug = hyphenRegex.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}
