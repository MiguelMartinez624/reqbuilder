package reqbuilder

import (
	"net/http"
)

type GetRequest struct {
	*Request
	requestConfiguration RequestConfiguration
}

func (pr GetRequest) Build() (*http.Request, error) {
	reqBuilder := pr.requestConfiguration.ReqBuilder
	method := pr.requestConfiguration.Method

	endpoint := pr.generateEndpoint()
	req, err := reqBuilder(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	pr.appendRequestHeaders(req)

	return req, nil
}
