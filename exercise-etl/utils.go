package main

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

func Map[T, U any](slc []T, fnc func(T) U) []U {
	result := make([]U, len(slc))
	for i, v := range slc {
		result[i] = fnc(v)
	}
	return result
}

func NormalizeString(str string) string {
	return caser.String(strings.ToLower(strings.TrimSpace(str)))
}

func NormalizeStringPtr(str *string) *string {
	if str == nil {
		return nil
	}
	title := NormalizeString(*str)
	return &title
}
