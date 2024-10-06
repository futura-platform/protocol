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

	GetProxy() *proxyprotocol.Proxy
	SetProxy(proxy *proxyprotocol.Proxy) error

	GetCookieJar() http.CookieJar
	SetCookieJar(jar http.CookieJar)

	GetHeaderDefaults() map[string]string

	DoRequest(req *Request) (*Response, error)

	Get(url string, headers OrderedHeaders) (*Response, error)
	Post(url string, body any, headers OrderedHeaders) (*Response, error)
	Put(url string, body any, headers OrderedHeaders) (*Response, error)
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

	BodyReader io.ReadCloser
	readCache  []byte
}

func (r *Response) Body() []byte {
	if r.readCache != nil {
		return r.readCache
	}

	defer r.BodyReader.Close()
	var respBody []byte
	var err error
	if !(r.StatusCode >= 300 && r.StatusCode < 400) {
		respBody, err = io.ReadAll(r.BodyReader)
		if err != nil {
			fmt.Println("failed to read response body (2):", err)
		}
	}
	r.readCache = respBody
	return respBody
}

func (r *Response) JSON(v any) error {
	return json.Unmarshal(r.Body(), v)
}
