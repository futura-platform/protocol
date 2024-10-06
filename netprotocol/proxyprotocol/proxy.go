package proxyprotocol

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
)

type Proxy url.URL

var _ basicgroupsprotocol.Parsable[*Proxy] = (*Proxy)(nil)

func (*Proxy) ParseEntry(u string) (*Proxy, error) {
	p, err := url.Parse(u)
	if err != nil {
		spl := strings.Split(u, ":")
		if len(spl) == 2 {
			return &Proxy{
				Scheme: "http",
				Host:   fmt.Sprintf("%s:%s", spl[0], spl[1]),
			}, nil
		} else if len(spl) == 4 {
			return &Proxy{
				Scheme: "http",
				Host:   fmt.Sprintf("%s:%s", spl[0], spl[1]),
				User:   url.UserPassword(spl[2], spl[3]),
			}, nil
		}

		return nil, fmt.Errorf("failed to parse proxy: %w", err)
	}
	return (*Proxy)(p), nil
}

func (p *Proxy) SerializeEntry() string {
	return p.String()
}

func (p *Proxy) Equals(p2 *Proxy) bool {
	return p.String() == p2.String()
}

func (p *Proxy) GetGroupConfig() basicgroupsprotocol.GroupConfig {
	return basicgroupsprotocol.GroupConfig{
		EntryTypeSingular: "Proxy",
		EntryTypePlural:   "Proxies",
		EntryPlaceholder:  "http://user:password@host:port or host:port:user:password",
		Icon:              "mdi:proxy",
	}
}

func (p *Proxy) String() string {
	return (*url.URL)(p).String()
}

func (p *Proxy) HostName() string {
	return (*url.URL)(p).Hostname()
}

func (p *Proxy) Port() int {
	portString := (*url.URL)(p).Port()
	n, err := strconv.Atoi(portString)
	if err != nil {
		switch p.Scheme {
		case "http":
			return 80
		case "https":
			return 443
		case "socks5":
			return 1080
		}
	}

	return n
}

// define marhsal and unmarshal methods for Proxy as url.Parse and url.String
func (p Proxy) MarshalJSON() ([]byte, error) {
	return json.Marshal((*url.URL)(&p).String())
}

func (p *Proxy) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	u, err := url.Parse(str)
	if err != nil {
		return err
	}
	*p = Proxy(*u)
	return nil
}
