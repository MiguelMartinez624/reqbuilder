package reqbuilder

import (
	"bytes"
	"net/http"
	"strings"
)

type PostRequest struct {
	requestConfiguration RequestConfiguration
	body                 *bytes.Reader
}

func (pr PostRequest) Body(body string) PostRequest {
	pr.body = bytes.NewReader([]byte(body))
	return pr
}

func (pr PostRequest) Build() (*http.Request, error) {
	reqBuilder := pr.requestConfiguration.ReqBuilder
	method := pr.requestConfiguration.Method

	rawPath := pr.requestConfiguration.Endpoint
	var endpoint string
	for paramKey, paramValue := range pr.requestConfiguration.QueryParams {
		endpoint = strings.Replace(rawPath, paramKey, paramValue, 1)
	}

	req, err := reqBuilder(method, endpoint, pr.body)
	if err != nil {
		return nil, err
	}
	headers := pr.requestConfiguration.Headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}
