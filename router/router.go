package router

import (
	"fmt"
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
	fmt.Println("ServerHTTP - 001")
	defer func() {
	fmt.Println("ServerHTTP - 002")
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
			http.Error(w, "Error: ", http.StatusInternalServerError)
		}
	}()

	fmt.Println("ServerHTTP - 003")
	for _, e := range rtr.routes {
	fmt.Println("ServerHTTP - 004")
	    params := e.Match(r)
	fmt.Println("ServerHTTP - 005")
	    if params == nil {
	fmt.Println("ServerHTTP - 006")
		    continue
	    }

	fmt.Println("ServerHTTP - 007")
	    // Create new request with params stored in context
	    ctx := context.WithValue(r.Context(), "params", params)
	fmt.Println("ServerHTTP - 008")
	    e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))
	fmt.Println("ServerHTTP - 009")

	    return
	}

	fmt.Println("ServerHTTP - 010")
        http.NotFound(w, r)
}
