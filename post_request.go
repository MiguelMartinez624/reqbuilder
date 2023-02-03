package reqbuilder

import (
	"bytes"
	"net/http"
)

type PostRequest struct {
	*Request
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

	endpoint := pr.generateEndpoint()

	req, err := reqBuilder(method, endpoint, pr.body)
	if err != nil {
		return nil, err
	}
	pr.appendRequestHeaders(req)

	return req, nil
}
