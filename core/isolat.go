package core

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"net/http"
	"time"
)

type Isolat struct {
	fange         Fange
	isoleringsTid int
	callbackUrl   string
	method        string
}

func (i *Isolat) StartSoning() {
	Trace.Println(i.fange.String(), "Fange starter isolat.")

	time.Sleep(time.Duration(i.isoleringsTid) * time.Second)
	fangeBytes, _ := json.Marshal(i.fange)

	client := &http.Client{}
	req, _ := http.NewRequest(i.method, i.callbackUrl, bytes.NewReader(fangeBytes))
	resp, e := client.Do(req)

	if e != nil {
		Error.Println(i.fange.String(), "Feil ved request: ", e.Error())
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		Error.Println(i.fange.String(), "Fange ble ikke løslat.")
		return
	}

	Trace.Println(i.fange.String(), "Fange løslates.")

}
