package sessionsprotocol

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type ChromeCookie struct {
	Url    string `json:"url"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	Secure bool   `json:"secure"`
}

type MockResponse struct {
	Url    string
	Method string

	ReplacementStatus int
	ReplacementBody   string
}

type ExportSessionDetails struct {
	TargetLocation string

	CookieJar      http.CookieJar
	SessionStorage map[string]any

	Proxy *proxyprotocol.Proxy

	MockedResponses []MockResponse
}

type Provider interface {
	// Deprecated: use the exported "SaveSession" func instead for type safety
	SaveSession(ctx context.Context, sessionID, typeKey string, global bool, expireAt *time.Time, session any) error
	// Deprecated: use the exported "LoadSession" func instead for type safety
	LoadSession(ctx context.Context, sessionID, typeKey string, global bool, target any) (bool, error)
	// Deprecated: use the exported "DeleteSession" func instead for type safety
	DeleteSession(ctx context.Context, sessionID, typeKey string, global bool) error

	ExportSession(ctx context.Context, details ExportSessionDetails) (*url.URL, error)
}

type providerWithContext interface {
	context.Context
	Provider
}

func SaveSession[T any](p providerWithContext, sessionID string, global bool, expireAt *time.Time, session T) error {
	return p.SaveSession(p, sessionID, reflect.TypeOf(session).String(), global, expireAt, session)
}

func LoadSession[T any](p providerWithContext, sessionID string, global bool) (*T, bool, error) {
	session := new(T)
	ok, err := p.LoadSession(p, sessionID, reflect.TypeOf(*session).String(), global, session)

	return session, ok, err
}

func DeleteSession[T any](p providerWithContext, sessionID string, global bool) error {
	return p.DeleteSession(p, sessionID, reflect.TypeOf(*new(T)).String(), global)
}
