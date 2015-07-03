package core

import (
	"io/ioutil"
	//	"strconv"
	//	"time"
	//"fmt"
	//	"bytes"
	//	"io"
	"net/http"
)

type RestHandler struct {
	mottak *Mottak
}

func NyRestHandler() *RestHandler {
	return &RestHandler{mottak: NyttMottak()}
}

func (h RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
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
	}
}

func (h RestHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET!"))
}

func (h RestHandler) ReceivePut(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("PUT!"))
	b, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte(string(b)))
	h.mottak.Motta(string(b))
}

func (h RestHandler) ReceiveDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE!"))
}
