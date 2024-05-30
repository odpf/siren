package template

import (
	"regexp"
	"strconv"
	"strings"
	texttemplate "text/template"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var defaultFuncMap = func() texttemplate.FuncMap {
	f := sprig.TxtFuncMap()

	extra := texttemplate.FuncMap{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
		"title":   cases.Title(language.AmericanEnglish).String,
		// join is equal to strings.Join but inverts the argument order
		// for easier pipelining in templates.
		"join": func(sep string, s []string) string {
			return strings.Join(s, sep)
		},
		"joinStringValues": func(sep string, ms map[string]string) string {
			var joinedString []string
			for _, v := range ms {
				joinedString = append(joinedString, v)
			}
			return strings.Join(joinedString, sep)
		},
		"match": regexp.MatchString,
		"reReplaceAll": func(pattern, repl, text string) string {
			re := regexp.MustCompile(pattern)
			return re.ReplaceAllString(text, repl)
		},
		"stringSlice": func(s ...string) []string {
			return s
		},
		"sub": func(a, b string) string {
			aint, err := strconv.ParseInt(a, 10, 64)
			if err != nil {
				return "unknown" // fallback
			}
			bint, err := strconv.ParseInt(b, 10, 64)
			if err != nil {
				return "unknown" // fallback
			}

			return strconv.FormatInt(aint-bint, 10)
		},
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}()
