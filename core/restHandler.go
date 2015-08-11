package core

import (
	//"io/ioutil"
	//	"strconv"
	//	"time"
	//"fmt"
	//	"bytes"
	//	"io"
	"encoding/json"
	"net/http"
)

type RestHandler struct {
	mottak *Mottak
}

func NyRestHandler() *RestHandler {
	return &RestHandler{mottak: NyttMottak()}
}

func (h RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.ReceiveGet(w, r)
	case "PUT":
		h.ReceivePut(w, r)
	case "DELETE":
		h.ReceiveDelete(w, r)
	case "OPTIONS":
		w.Write([]byte("The handler supports GET, PUT and DELETE!"))
	default:
		http.Error(w, "Method not supported.", 405)
		Error.Println("Metode er ikke støttet: " + r.Method)
	}
}

func (h RestHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eksempel på payload: "))

	fange := IsolatFange{FangeTilIsolat: Fange{Id: "1", Navn: "Arne"}, IsoleringsTid: 5}

	encoder := json.NewEncoder(w)
	encoder.Encode(&fange)
	Info.Println("Get mottat.")
}

func (h RestHandler) ReceivePut(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	fange := IsolatFange{}
	err := decoder.Decode(&fange)

	if err != nil {
		http.Error(w, "Invalid content.", 400)
		Error.Println("Ugyldig innhold.")

		return
	}

	h.mottak.Motta(fange)
	Info.Println("Mottat ny fange:." + fange.FangeTilIsolat.String())

}

func (h RestHandler) ReceiveDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE!"))
}
