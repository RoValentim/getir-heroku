package main

import (
	"log"
        "net/http"
	"rodrigo/restful/router"
	"rodrigo/restful/memory"
	"rodrigo/restful/database"
)

func main() {
log.Println("01")
	r := &router.Router{}
log.Println("02")

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Working!"))
	})
log.Println("03")

	r.Route(http.MethodPost, `/v1/database`, database.GetData )
log.Println("04")
	r.Route(http.MethodGet,  `/v1/memory`,   memory.GetData   )
log.Println("05")
	r.Route(http.MethodPost, `/v1/memory`,   memory.CreateData)
log.Println("06")

        http.ListenAndServe(":443", r)
log.Println("07")
}
