package reqbuilder

import "net/http"

// Post generate a post request
func (r Request) Post() PostRequest {
	r.requestConfiguration.Method = http.MethodPost
	return PostRequest{
		Request:              &r,
		requestConfiguration: r.requestConfiguration,
	}
}

func (r Request) Delete() DeleteRequest {
	r.requestConfiguration.Method = http.MethodDelete
	return DeleteRequest{
		Request:              &r,
		requestConfiguration: r.requestConfiguration,
	}
}

func (r Request) Get() GetRequest {
	r.requestConfiguration.Method = http.MethodGet
	return GetRequest{
		Request:              &r,
		requestConfiguration: r.requestConfiguration,
	}
}

func (r Request) Put() PutRequest {
	r.requestConfiguration.Method = http.MethodPut
	return PutRequest{
		Request:              &r,
		requestConfiguration: r.requestConfiguration,
	}
}
