package main

import (
        "net/http"
	"rodrigo/restful/router"
	"rodrigo/restful/memory"
	"rodrigo/restful/database"
)

func main() {
	r := &router.Router{}

	r.Route(http.MethodPost, `/v1/database`, database.GetData )
	r.Route(http.MethodGet,  `/v1/memory`,   memory.GetData   )
	r.Route(http.MethodPost, `/v1/memory`,   memory.CreateData)

        http.ListenAndServe(":8080", r)
}
