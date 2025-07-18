package netprotocol

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type OrderedHeaders [][]string
type BaseHttpClient interface {
	GetContext() context.Context

	GetBrowserProfile() BrowserProfile
	SetBrowserProfile(profile BrowserProfile) error

	GetProxy() *proxyprotocol.Proxy
	SetProxy(proxy *proxyprotocol.Proxy) error

	GetCookieJar() http.CookieJar
	SetCookieJar(jar http.CookieJar)

	GetHeaderDefaults() map[string]string
	SetHeaderDefaults(h map[string]string)

	DoRequest(req *Request) (*Response, error)
}

type HttpClient struct{ BaseHttpClient }

func (h *HttpClient) Get(url string, headers OrderedHeaders) (*Response, error) {
	return h.DoRequest(&Request{
		Method:  "GET",
		Url:     url,
		Headers: headers,
	})
}

func (h *HttpClient) Post(url string, headers OrderedHeaders, body any) (*Response, error) {
	return h.DoRequest(&Request{
		Method:  "POST",
		Url:     url,
		Headers: headers,
		Body:    body,
	})
}

func (h *HttpClient) Put(url string, headers OrderedHeaders, body any) (*Response, error) {
	return h.DoRequest(&Request{
		Method:  "PUT",
		Url:     url,
		Headers: headers,
		Body:    body,
	})
}

func (h *HttpClient) Delete(url string, headers OrderedHeaders) (*Response, error) {
	return h.DoRequest(&Request{
		Method:  "DELETE",
		Url:     url,
		Headers: headers,
	})
}

func (h *HttpClient) Patch(url string, headers OrderedHeaders, body any) (*Response, error) {
	return h.DoRequest(&Request{
		Method:  "PATCH",
		Url:     url,
		Headers: headers,
		Body:    body,
	})
}

type RequestPriority int

const (
	/*h2 weight: 73*/ THROTTLED RequestPriority = 0 // Used to signal that resources
	// should be reserved for following
	// requests (i.e. that higher priority
	// following requests are expected).
	MINIMUM_PRIORITY RequestPriority = THROTTLED
	/*h2 weight: 109*/ IDLE RequestPriority = 1 // Default "as resources available" level.
	/*h2 weight: 146*/ LOWEST RequestPriority = 2
	DEFAULT_PRIORITY          RequestPriority = LOWEST
	/*h2 weight: 182*/ LOW RequestPriority = 3
	/*h2 weight: 219*/ MEDIUM RequestPriority = 4
	/*h2 weight: 255*/ HIGHEST RequestPriority = 5
	MAXIMUM_PRIORITY           RequestPriority = HIGHEST
)

func PriorityPtr(p RequestPriority) *RequestPriority {
	return &p
}

type Request struct {
	Method  string
	Url     string
	Headers OrderedHeaders

	Body any

	RequestPriority *RequestPriority

	PreloadAndSendAt time.Time
	PreloadSleepFunc func(time.Duration)

	ForceIp net.IP

	DontReadResponseBody bool
}

type Response struct {
	*http.Response

	readCache []byte
}

func (r *Response) GetBody() []byte {
	if r.readCache != nil {
		return r.readCache
	}

	defer r.Response.Body.Close()
	var respBody []byte
	var err error
	if !(r.StatusCode >= 300 && r.StatusCode < 400) {
		respBody, err = io.ReadAll(r.Response.Body)
		if err != nil {
			fmt.Println("failed to read response body (2):", err)
		}
	}
	r.readCache = respBody
	return respBody
}

func (r *Response) JSON(v any) error {
	return json.Unmarshal(r.GetBody(), v)
}
