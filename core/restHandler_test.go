package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSkalSetteIIsolat(t *testing.T) {
	fange := &IsolatFange{FangeTilIsolat: Fange{Id: "1", Navn: "Ola Norman"}, IsoleringsTid: 5}
	fangeBytes, _ := json.Marshal(&fange)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "test", bytes.NewReader(fangeBytes))

	r := NyRestHandler()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}

}
