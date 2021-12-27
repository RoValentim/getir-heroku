package router

import (
	"regexp"
	"strings"
        "net/http"
)

type RouteEntry struct {
        Path        *regexp.Regexp
        Method      string
        HandlerFunc http.HandlerFunc
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil || r.Method != ent.Method {
		return nil
	}

	params  := make(map[string]string)
	filters := r.URL.Query()

	for i, group := range filters {
		params[strings.ToLower(i)] = strings.Join(group,",")
	}

	// For this assignment, post method can't have filter
	if r.Method == http.MethodPost && len(params) > 0 {
		return nil
	}

	return params
}
