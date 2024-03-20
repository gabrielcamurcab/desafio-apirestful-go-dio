package utils

import "strings"

func GetURLParam(path string) string {
	segments := strings.Split(path, "/")
	if len(segments) > 2 {
		return segments[len(segments)-1]
	}
	return ""
}
