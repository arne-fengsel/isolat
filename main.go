// isolat project isolat.go
package main

import (
	//"fmt"
	"github.com/goarne/web"
	//"io/ioutil"
	"mesan.no/fagark/isolat/core"
	"net/http"
	"strconv"
	//"os"
)

func main() {
	appConfig := core.AppConfig{}
	appConfig.ReadConfig("./config/appconfig.json")
	core.InitLogFile(appConfig.Logging)

	router := web.NewWebRouter()

	router.AddRoute(opprettIsolatRessurs())
	http.ListenAndServe(":"+strconv.FormatInt(appConfig.Server.Port, 10), router)
}

func opprettIsolatRessurs() *web.Route {
	r := web.NewRoute()

	r.Path("/isolat(/)?")
	r.Method(web.HttpGet).Method(web.HttpPut)
	r.Handler(core.NyRestHandler())

	return r
}
