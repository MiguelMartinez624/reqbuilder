package reqbuilder

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

type HeadersMap map[string]string
type QueryParamsMap map[string]string

type RequestConfiguration struct {
	Method      string
	Endpoint    string
	Headers     HeadersMap
	QueryParams QueryParamsMap
	ReqBuilder  func(method string, target string, body io.Reader) (*http.Request, error)
}

type Request struct {
	requestConfiguration RequestConfiguration
}

var defaultHTTPRequestBuilderFunc = func(method string, target string, body io.Reader) (*http.Request, error) {
	return httptest.NewRequest(method, target, body), nil
}

// WithHeader add a header to the request configuration, will panic if
// the key already exist on the headers
func (r Request) WithHeader(key, value string) Request {
	if _, ok := r.requestConfiguration.Headers[key]; ok {
		panic(fmt.Sprintf("cannot use key %v already exist on headers map", key))
	}

	r.requestConfiguration.Headers[key] = value
	return Request{requestConfiguration: r.requestConfiguration}
}

func (r Request) QueryParam(key, value string) Request {
	if _, ok := r.requestConfiguration.QueryParams[key]; ok {
		panic(fmt.Sprintf("cannot use key %v already exist on headers map", key))
	}

	r.requestConfiguration.QueryParams[key] = value
	return Request{requestConfiguration: r.requestConfiguration}
}

func (r Request) generateEndpoint() string {
	rawPath := r.requestConfiguration.Endpoint
	var endpoint string
	for paramKey, paramValue := range r.requestConfiguration.QueryParams {
		endpoint = strings.Replace(rawPath, fmt.Sprintf(":%v", paramKey), paramValue, 1)
	}

	return endpoint

}

func (r Request) appendRequestHeaders(req *http.Request) {
	headers := r.requestConfiguration.Headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func NewRequest(route string) Request {
	return Request{
		RequestConfiguration{
			Endpoint:    route,
			Headers:     map[string]string{},
			QueryParams: map[string]string{},
			// by default will build a valid http request
			ReqBuilder: defaultHTTPRequestBuilderFunc,
		},
	}
}
