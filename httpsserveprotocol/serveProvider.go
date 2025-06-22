package httpsserveprotocol

import "net/http"

type Provider interface {
	// Serve creates an https endpoint which serves the given http.ServeMux.
	// It returns the URL prefix that the server is listening on,
	// and a function to close the server.
	Serve(*http.ServeMux) (urlPrefix string, close func(), err error)
}
