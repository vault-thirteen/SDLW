package c

import "strings"

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

func DetachPointerMarks(sIn string) (sOut string) {
	return strings.ReplaceAll(sIn, "*", " * ")
}

func GlueParts(parts []string) (sOut string) {
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part)
	}
	return sb.String()
}
