package goubi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewUbiClient(t *testing.T) {
	client := NewUbiClient(123, "ubic-123", "65as4df65as4f6sad4f")
	ubi_services := new(UbiServices)
	if reflect.TypeOf(client) != reflect.TypeOf(ubi_services) {
		t.Error("Error: Did not make the right type")
	}
}

func TestNewUbiClientWithUseServiceID(t *testing.T) {
	client := NewUbiClient(123, "ubic-123", "65as4df65as4f6sad4f", true)
	ubi_services := new(UbiServices)
	if reflect.TypeOf(client) != reflect.TypeOf(ubi_services) {
		t.Error("Error: Did not make the right type")
	}
}

func evalResponse(res interface{}, err error, exp interface{}, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("\nExpected: %+v\nGot: %+v", exp, res)
	}
}

func makeWellBehavedServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, `{"Backups":[{"id": "22"}, {"id": "88"}]}`)
	}))
}

func makeBadlyBehavedServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Error-Code", "33")
		w.Header().Set("X-Error-Message", "Failure")
		w.WriteHeader(404)
	}))
}

func TestMakeRequest(t *testing.T) {
	client := new(ubiApiClient)
	client.makeRequest("cloud.test", make(map[string]string, 0))
	client.makeRequest("cloud.test")
}

func TestCall(t *testing.T) {
	server := makeWellBehavedServer()
	defer server.Close()

	client := new(ubiApiClient)
	ubi_request := new(ubiRequest)
	ubi_request.URL = server.URL
	ubi_request.Params = map[string]string{"vm_id": "123"}
	client.call(ubi_request)
}

func TestCallWithUseServiceID(t *testing.T) {
	server := makeWellBehavedServer()
	defer server.Close()

	client := new(ubiApiClient)
	client.use_service_id = true
	ubi_request := new(ubiRequest)
	ubi_request.URL = server.URL
	ubi_request.Params = map[string]string{"vm_id": "123", "unknown": ""}
	client.call(ubi_request)
}

func TestCallErrors(t *testing.T) {
	server := makeBadlyBehavedServer()
	defer server.Close()

	client := new(ubiApiClient)
	ubi_request := new(ubiRequest)
	ubi_request.URL = server.URL
	ubi_request.Params = map[string]string{"vm_id": "123"}
	client.call(ubi_request)
}

func TestCallWithNoServer(t *testing.T) {
	client := new(ubiApiClient)
	ubi_request := new(ubiRequest)
	ubi_request.URL = "http://111.222.333.444/"
	ubi_request.Params = map[string]string{"vm_id": "123"}
	_, err := client.call(ubi_request)
	if err == nil {
		t.Error("Error: no error thrown on server connection error")
	}
}

func TestMalformedJSON(t *testing.T) {
	input := []byte(`{"Backups":[{"id": "22"}, {"id": "11"}}`)
	output := new(Backups)
	err := unmarshalToStruct(input, output)
	if err == nil {
		t.Error("Error: no error thrown on JSON error")
	}
}
