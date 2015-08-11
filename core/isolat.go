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
	replyUrl      string
	method        string
}

func (i *Isolat) StartSoning() {
	Trace.Println("Fange " + i.fange.String() + " starter isolat.")

	time.Sleep(time.Duration(i.isoleringsTid) * time.Second)
	fangeBytes, _ := json.Marshal(i.fange)

	client := &http.Client{}
	req, _ := http.NewRequest(i.method, i.replyUrl, bytes.NewReader(fangeBytes))
	_, e := client.Do(req)
	Trace.Println("Fange " + i.fange.String() + " løslates." + e.Error())

}
