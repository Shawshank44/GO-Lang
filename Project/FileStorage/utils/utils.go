package utils

import "strings"

func EscapeField(s string, delim rune) string {
	var b strings.Builder

	for _, r := range s {
		if r == '\\' {
			b.WriteRune('\\')
			b.WriteRune('\\')
		} else if r == delim {
			b.WriteRune('\\')
			b.WriteRune(delim)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func UnescapeField(s string, delim rune) (string, error) {
	var b strings.Builder
	esc := false
	for _, r := range s {
		if esc {
			if r == delim || r == '\\' {
				b.WriteRune(r)
				esc = false
				continue
			}
			b.WriteRune('\\')
			b.WriteRune(r)
			esc = false
			continue
		}
		if r == '\\' {
			esc = true
			continue
		}
		b.WriteRune(r)
	}
	if esc {
		b.WriteRune('\\')
	}
	return b.String(), nil
}

func SplitEscaped(line string, delim rune) ([]string, error) {
	var parts []string
	var b strings.Builder

	esc := false

	for _, r := range line {
		if esc {
			if r == delim || r == '\\' {
				b.WriteRune(r)
			} else {
				b.WriteRune('\\')
				b.WriteRune(r)
			}
			esc = false
			continue
		}
		if r == '\\' {
			esc = true
			continue
		}
		if r == delim {
			parts = append(parts, b.String())
			b.Reset()
			continue
		}
		b.WriteRune(r)
	}
	parts = append(parts, b.String())
	return parts, nil
}
