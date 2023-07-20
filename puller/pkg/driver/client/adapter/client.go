package adapter

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	types "github.com/thnkrn/comet/puller/pkg/types"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	client := &http.Client{
		// REF: https://stackoverflow.com/questions/40338711/not-able-to-pass-bearer-token-in-headers-of-a-get-request-in-golang
		// NOTE: GO by default does not forward the headers, thus the bearer token was being lost in the middle. So, we had to override the client's CheckRedirect function and manually pass the headers to the new request.
		CheckRedirect: checkRedirectFunc,
		Timeout:       900 * time.Second,
	}
	return &HTTPClient{client}
}

func checkRedirectFunc(r *http.Request, via []*http.Request) error {
	r.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}

func WithAuth(key string) types.RequestModifier {
	return func(r *http.Request) {
		r.Header.Add("Authorization", key)
	}
}

func WithContentType(contentType string) types.RequestModifier {
	return func(r *http.Request) {
		r.Header.Add("Content-Type", contentType)
	}
}

func (hc *HTTPClient) Get(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	// NOTE: add headers from option functions
	for _, mod := range opts {
		mod(request)
	}

	resp, err := hc.client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result, responseError string
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		responseError = string(bodyBytes)
	} else {
		result = string(bodyBytes)
	}
	clientResponse := domain.ClientResponse{Result: result, StatusCode: resp.StatusCode, Error: responseError}
	copier.Copy(res, &clientResponse)

	return nil
}

func (hc *HTTPClient) Post(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	// NOTE: add headers from option functions
	for _, mod := range opts {
		mod(request)
	}

	resp, err := hc.client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return err
	}

	return nil
}

func (hc *HTTPClient) Put(url string, req interface{}, res interface{}, opts ...types.RequestModifier) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	// NOTE: add headers from option functions
	for _, mod := range opts {
		mod(request)
	}

	resp, err := hc.client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var responseError string
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		responseError = string(bodyBytes)
	}

	clientResponse := domain.ClientResponse{StatusCode: resp.StatusCode, Error: responseError}
	copier.Copy(res, &clientResponse)
	return nil
}
