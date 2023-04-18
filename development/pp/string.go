package pp

import "strings"

// RemoveDoubleSpaces removes all double spaces from string.
func RemoveDoubleSpaces(sIn string) (sOut string) {
	sOut = sIn
	for {
		if strings.Contains(sOut, "  ") {
			sOut = strings.ReplaceAll(sOut, "  ", " ")
		} else {
			break
		}
	}
	return sOut
}

// DetachPointerMarks makes pointer symbols detached from other words.
// Beware that this function can create some redundant spaces.
func DetachPointerMarks(sIn string) (sOut string) {
	return strings.ReplaceAll(sIn, "*", " * ")
}

// GlueParts concatenates string parts into a single string.
func GlueParts(parts []string) (sOut string) {
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part)
	}
	return sb.String()
}
