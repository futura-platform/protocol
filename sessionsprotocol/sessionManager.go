package sessionsprotocol

import (
	"context"
	"encoding/json"
	"fmt"
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

	ReplacementStatus  int
	ReplacementHeaders []MockResponseHeader
	ReplacementBody    string
}

type MockResponseHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExportSessionDetails struct {
	TargetLocation string

	CookieJar                    http.CookieJar
	SessionStorage, LocalStorage map[string]any

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

type ScopedProvider[K comparable, V any] struct{}
type providerWithContext interface {
	context.Context
	Provider
}

func keyToSessionID(key any) (string, error) {
	b, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}

	return string(b), nil
}

func (ScopedProvider[K, V]) SaveSession(p providerWithContext, key K, global bool, expireAt *time.Time, session V) error {
	sessionID, err := keyToSessionID(key)
	if err != nil {
		return fmt.Errorf("failed to convert key to session ID: %w", err)
	}
	return p.SaveSession(p, sessionID, reflect.TypeOf(session).String(), global, expireAt, session)
}

func (ScopedProvider[K, V]) LoadSession(p providerWithContext, key K, global bool) (*V, bool, error) {
	sessionID, err := keyToSessionID(key)
	if err != nil {
		return nil, false, fmt.Errorf("failed to convert key to session ID: %w", err)
	}
	session := new(V)
	ok, err := p.LoadSession(p, sessionID, reflect.TypeOf(*session).String(), global, session)

	return session, ok, err
}

func (ScopedProvider[K, V]) DeleteSession(p providerWithContext, key K, global bool) error {
	sessionID, err := keyToSessionID(key)
	if err != nil {
		return fmt.Errorf("failed to convert key to session ID: %w", err)
	}
	return p.DeleteSession(p, sessionID, reflect.TypeOf(*new(V)).String(), global)
}
