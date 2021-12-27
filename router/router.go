package router

import (
	"context"
	"log"
	"regexp"
        "net/http"
)

type Router struct {
        routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")

        e := RouteEntry{
                Method:      method,
                Path:        exactPath,
                HandlerFunc: handlerFunc,
        }

        rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
			http.Error(w, "Error: ", http.StatusInternalServerError)
		}
	}()

	for _, e := range rtr.routes {
	    params := e.Match(r)
	    if params == nil {
		    continue
	    }

	    // Create new request with params stored in context
	    ctx := context.WithValue(r.Context(), "params", params)
	    e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))

	    return
	}

        http.NotFound(w, r)
}
