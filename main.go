// isolat project isolat.go
package main

import (
	//"fmt"
	"github.com/goarne/web"
	//"io/ioutil"
	"mesan.no/fagark/isolat/core"
	"net/http"
	//"os"
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

func init() {
	core.LoadConfig("./config/logging.json")
	core.Info.Println("Starter isolat.")
}
