package netprotocol

import (
	"context"

	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type ClientProvider interface {
	MakeClient(ctx context.Context, browserProfile BrowserProfile, proxy *proxyprotocol.Proxy) (HttpClient, error)
}
