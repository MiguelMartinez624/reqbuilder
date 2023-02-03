package reqbuilder

import (
	"bytes"
	"net/http"
)

type PutRequest struct {
	*Request
	requestConfiguration RequestConfiguration
	body                 *bytes.Reader
}

func (pr PutRequest) Body(body string) PutRequest {
	pr.body = bytes.NewReader([]byte(body))
	return pr
}

func (pr PutRequest) Build() (*http.Request, error) {
	reqBuilder := pr.requestConfiguration.ReqBuilder
	method := pr.requestConfiguration.Method

	endpoint := pr.generateEndpoint()

	req, err := reqBuilder(method, endpoint, pr.body)
	if err != nil {
		return nil, err
	}
	pr.appendRequestHeaders(req)

	return req, nil
}
