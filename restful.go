package main

import (
        "net/http"
	"rodrigo/restful/router"
	"rodrigo/restful/memory"
	"rodrigo/restful/database"
)

func main() {
	r := &router.Router{}

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Working!"))
	})

	r.Route(http.MethodPost, `/v1/database`, database.GetData )
	r.Route(http.MethodGet,  `/v1/memory`,   memory.GetData   )
	r.Route(http.MethodPost, `/v1/memory`,   memory.CreateData)

        http.ListenAndServe(":443", r)
}
