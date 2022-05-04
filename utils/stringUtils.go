package utils

import "strings"

func StringJoin(args ...string) string {
	var builder strings.Builder
	for _, s := range args {
		builder.WriteString(s)
	}
	return builder.String()
}
