package plugin

import (
	"html/template"
	"strings"
)

func ToUpper() template.FuncMap {
	f := make(template.FuncMap)
	f["Upper"] = func(s string) string {
		return strings.ToUpper(s)
	}
	return f
}
