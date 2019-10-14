// Code generated by go-swagger; DO NOT EDIT.

package todoclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/IAD/go-swagger-template/examples/todo/gen/client/todoclient/todos"
)

// Default todo HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new todo HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Todo {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new todo HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Todo {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new todo client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Todo {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Todo)
	cli.Transport = transport

	cli.Todos = todos.New(transport, formats)

	return cli
}

func NewDateTime(t time.Time) strfmt.DateTime {
	return strfmt.DateTime(t)
}

func NewClientWithBasePath(url string, endpoint string) *Todo {
	schemes := []string{"http"}
	if strings.HasSuffix(url, ":443") {
		schemes = []string{"https"}
	}

	transport := httptransport.New(url, endpoint, schemes)
	return New(transport, strfmt.Default)
}

//nolint:check-panic
func NewExternalClientWithEnv() *Todo {
	host := "todo"
	v := os.Getenv("TODO_EXTERNAL_HOST")
	if v != "" {
		host = v
	}

	port := 443
	v = os.Getenv("TODO_EXTERNAL_PORT")
	if v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid port provided for TODO_EXTERNAL_PORT=%s %s", v, err))
		}
		port = i
	}

	schemes := []string{"http"}
	if port == 443 {
		schemes = []string{"https"}
	}

	url := fmt.Sprintf("%s:%d", host, port)

	transport := httptransport.New(url, "", schemes)
	transport.Consumers["text/html"] = runtime.TextConsumer()

	return New(transport, strfmt.Default)
}

//nolint:check-panic
func NewClientWithEnv() *Todo {
	port := 80
	v := os.Getenv("TODO_LOCAL_PORT")
	if v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid port provided for TODO_LOCAL_PORT=%s %s", v, err))
		}
		port = i
	}

	schemes := []string{"http"}
	if port == 443 {
		schemes = []string{"https"}
	}

	url := fmt.Sprintf("todo:%d", port)

	transport := httptransport.New(url, "", schemes)
	transport.Consumers["text/html"] = runtime.TextConsumer()

	return New(transport, strfmt.Default)
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Todo is a client for todo
type Todo struct {
	Todos *todos.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Todo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Todos.SetTransport(transport)

}
