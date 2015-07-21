package goubi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*****************************************************************************/

// API URI constants
const (
	host = "http://api.ubiquityhosting.com"
	path = "/v25/api.php"
)

// Structure for available API services
type UbiServices struct {
	Cloud cloudService
}

type postClient interface {
	makeRequest(method string, params ...map[string]string) *ubiRequest
	call(ubireq *ubiRequest) ([]byte, error)
}

type ubiApiClient struct {
	use_service_id bool
}

type ubiRequest struct {
	URL    string
	Params map[string]string
}

/*****************************************************************************/

// Instantiate and return an interface to API client services
func NewUbiClient(clientid int, username string, token string, use_service_id ...bool) *UbiServices {
	setCredentials(clientid, username, token)
	ubic := &UbiServices{}
	api_client := new(ubiApiClient)

	// Flag for using service_id as vm_id
	api_client.use_service_id = false
	if len(use_service_id) > 0 {
		api_client.use_service_id = use_service_id[0]
	}

	ubic.Cloud = &Cloud{apiClient: api_client}
	return ubic
}

func (c *ubiApiClient) makeRequest(method string, params ...map[string]string) *ubiRequest {
	url := getAPIURI(method)
	if len(params) > 0 {
		ubireq := &ubiRequest{url, params[0]}
		return ubireq
	}
	ubireq := &ubiRequest{url, make(map[string]string, 0)}
	return ubireq
}

func (req *ubiApiClient) call(ubireq *ubiRequest) ([]byte, error) {

	username, password := getCredentials()

	data := url.Values{}
	for k, v := range ubireq.Params {
		if req.use_service_id == true && k == "vm_id" {
			k = "service_id"
		}
		if v == "" {
			continue
		}
		data.Add(k, v)
	}

	r, _ := http.NewRequest("POST", ubireq.URL, bytes.NewBufferString(data.Encode()))

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Ubiquity API Client Go/1.1")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	r.SetBasicAuth(username, password)

	return getResponse(r)
}

func getAPIURI(method string) string {
	query := fmt.Sprintf("?method=%s", method)
	return fmt.Sprintf("%s%s%s", host, path, query)
}

// Returns a response for a request
func getResponse(r *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if _, ok := resp.Header["X-Error-Code"]; ok {
		return nil, errors.New(resp.Header.Get("X-Error-Message"))
	}

	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil
}

// Converts a JSON string into a golang structure
func unmarshalToStruct(jstring []byte, t interface{}) error {
	if err := json.Unmarshal(jstring, &t); err != nil {
		return err
	}
	return nil
}
