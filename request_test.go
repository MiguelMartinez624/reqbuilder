package reqbuilder

import (
	"fmt"
	"reflect"
	"testing"
)

func compareRequest(t *testing.T, expected, got Request) {
	expectedEndpoint := expected.requestConfiguration.Endpoint
	gotEndpoint := got.requestConfiguration.Endpoint
	if !reflect.DeepEqual(expectedEndpoint, gotEndpoint) {
		t.Errorf("Endpoint() = %v, want %v", gotEndpoint, gotEndpoint)
	}

	expectedHeaders := expected.requestConfiguration.Headers
	gotdHeaders := got.requestConfiguration.Headers
	if !reflect.DeepEqual(expectedHeaders, gotdHeaders) {
		t.Errorf("Headers = %v, want %v", expectedHeaders, gotdHeaders)
	}
}

type newRequestTestCaseArgs struct {
	route string
}
type newRequestTestCase[T any] struct {
	name string
	args T
	want Request
}

func TestNewRequest(t *testing.T) {
	tests := []newRequestTestCase[newRequestTestCaseArgs]{
		{
			name: "should return a Request object with the same endpoint and init headers map",
			args: newRequestTestCaseArgs{"test"},
			want: Request{
				requestConfiguration: RequestConfiguration{
					Endpoint:   "test",
					Headers:    map[string]string{},
					ReqBuilder: defaultHTTPRequestBuilderFunc},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRequest(tt.args.route)
			compareRequest(t, tt.want, got)
		})
	}
}

type withHeaderstTestCaseArgs struct {
	key   string
	value string
}
type assignHeadersTestcase struct {
	name        string
	args        withHeaderstTestCaseArgs
	want        Request
	given       Request
	panicAssert func(t *testing.T, panicResult string)
}

func TestRequest_WithHeader(t *testing.T) {
	tests := []assignHeadersTestcase{
		{
			name: "should assign header to request if not exist already a key.",
			args: withHeaderstTestCaseArgs{"x-header", "test-value"},
			want: Request{
				requestConfiguration: RequestConfiguration{
					Headers:    HeadersMap{"x-header": "test-value"},
					ReqBuilder: defaultHTTPRequestBuilderFunc},
			},
			given: NewRequest(""),
		},
		{
			name: "should panic if you attempt to set a duplicated key.",
			args: withHeaderstTestCaseArgs{"x-header", "test-value"},
			want: Request{
				requestConfiguration: RequestConfiguration{
					Headers:    HeadersMap{"x-header": "test-value"},
					ReqBuilder: defaultHTTPRequestBuilderFunc,
				},
			},
			given: NewRequest("").WithHeader("x-header", "test-value"),
			panicAssert: func(t *testing.T, panicResult string) {
				expectedErr := "cannot use key x-header already exist on headers map"
				if panicResult != expectedErr {
					t.Errorf("WithHeader() =\n [%v] \n, want [%v]", fmt.Sprintf("%v", panicResult), expectedErr)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.panicAssert != nil {
						tt.panicAssert(t, r.(string))
					}
				}
			}()

			got := tt.given.WithHeader(tt.args.key, tt.args.value)
			compareRequest(t, tt.want, got)

		})
	}
}
