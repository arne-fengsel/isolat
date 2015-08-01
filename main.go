// isolat project isolat.go
package main

import (
	"github.com/goarne/web"
	"mesan.no/fagark/isolat/core"
	"net/http"
)

func main() {
	router := web.NewWebRouter()

	router.AddRoute(opprettIsolatRessurs())
	http.ListenAndServe(":9998", router)
}

func opprettIsolatRessurs() *web.Route {
	r := web.NewRoute()

	r.Path("/isolat(/)?")
	r.Method(web.HttpGet).Method(web.HttpPut)
	r.Handler(core.NyRestHandler())

	return r
}
