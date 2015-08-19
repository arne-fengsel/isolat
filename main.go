// isolat project isolat.go
package main

import (
	"flag"
	"github.com/goarne/web"
	"mesan.no/fagark/isolat/core"
	"net/http"
	"strconv"
	//"os"
)

func main() {
	configFile := flag.String("config", "./config/appconfig.json", "Fullt navn til applikasjonens konfigurasjonsfil (json)")
	flag.Parse()

	appConfig := core.AppConfig{}
	appConfig.ReadConfig(*configFile)

	logHandle := core.InitLogFile(appConfig.Logging)
	core.InitLoggers(logHandle, logHandle, logHandle, logHandle)

	router := web.NewWebRouter()
	router.AddRoute(opprettIsolatRessurs(appConfig.Server.Root))

	port := strconv.FormatInt(appConfig.Server.Port, 10)

	core.Info.Println("Laster appconfig:", *configFile)
	core.Info.Println("Starter isolat på port:", port)
	core.Info.Println("Isolat er tilgjengelig på ", appConfig.Server.Root)
	core.Info.Println("Skriver til loggfil: ", appConfig.Logging.Filename)

	http.ListenAndServe(":"+port, router)
}

func opprettIsolatRessurs(path string) *web.Route {
	r := web.NewRoute()

	r.Path(path)
	r.Method(web.HttpGet).Method(web.HttpPut)
	r.Handler(core.NyRestHandler())

	return r
}
