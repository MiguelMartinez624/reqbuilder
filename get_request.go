package reqbuilder

import (
	"fmt"
	"net/http"
	"strings"
)

type GetRequest struct {
	Request
	requestConfiguration RequestConfiguration
}

func (pr GetRequest) Build() (*http.Request, error) {
	reqBuilder := pr.requestConfiguration.ReqBuilder
	method := pr.requestConfiguration.Method
	rawPath := pr.requestConfiguration.Endpoint
	var endpoint string
	for paramKey, paramValue := range pr.requestConfiguration.QueryParams {
		endpoint = strings.Replace(rawPath, fmt.Sprintf(":%v", paramKey), paramValue, 1)
	}

	req, err := reqBuilder(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	headers := pr.requestConfiguration.Headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}
