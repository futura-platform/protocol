package browserprotocol

import (
	"context"

	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
	"github.com/futura-platform/protocol/netprotocol"
)

type OnRequestHandler func(*fetch.EventRequestPaused) (chromedp.Action, chan error)
type SingleTabBrowser struct {
	CTX            context.Context
	OnRequest      func(OnRequestHandler)
	ClearOnRequest func()
}

type Spawner interface {
	SpawnSingleTabBrowser(proxy *netprotocol.Proxy) (*SingleTabBrowser, context.CancelFunc, error)
}
