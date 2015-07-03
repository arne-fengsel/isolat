// isolat project isolat.go
package main

import (
	"github.com/goarne/web"
	"mesan.no/fagark/isolat/core"
	"net/http"
)

func main() {
	router := web.NewWebRouter()
	r := web.NewRoute()
	r.Path("/isolat/{id}")
	r.Method("POST").Method("GET").Method("PUT")
	r.Handler(core.NyRestHandler())

	router.AddRoute(r)
	http.ListenAndServe(":9998", router)
}
