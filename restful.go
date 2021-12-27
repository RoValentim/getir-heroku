package main

import (
	"fmt"
        "net/http"
	"rodrigo/restful/router"
	"rodrigo/restful/memory"
	"rodrigo/restful/database"
)

func main() {
fmt.Println("01")
	r := &router.Router{}
fmt.Println("02")

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
fmt.Println("Found main route")
		w.Write([]byte("Working!"))
	})
fmt.Println("03")

	r.Route(http.MethodPost, `/v1/database`, database.GetData )
fmt.Println("04")
	r.Route(http.MethodGet,  `/v1/memory`,   memory.GetData   )
fmt.Println("05")
	r.Route(http.MethodPost, `/v1/memory`,   memory.CreateData)
fmt.Println("06")

        http.ListenAndServe(":443", r)
fmt.Println("07")
}
