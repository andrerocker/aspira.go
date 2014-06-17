package util

import (
    "regexp"
    "strings"
)

func Clean(message string) string {
   return strings.TrimSpace(strings.Split(message, "aspira:")[1])
}

func RemoveHtml(message string) string {
    pattern := regexp.MustCompile("<[^<>]*>")
    return pattern.ReplaceAllString(message, "")
}
