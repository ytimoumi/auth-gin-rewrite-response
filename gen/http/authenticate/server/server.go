// Code generated by goa v3.10.1, DO NOT EDIT.
//
// authenticate HTTP server
//
// Command:
// $ goa gen auth/design

package server

import (
	authenticate "auth/gen/authenticate"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the authenticate service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	Authenticate http.Handler
	Openapi3JSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the authenticate service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *authenticate.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemOpenapi3JSON http.FileSystem,
) *Server {
	if fileSystemOpenapi3JSON == nil {
		fileSystemOpenapi3JSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Authenticate", "POST", "/auth/authentication/auth"},
			{"openapi3.json", "GET", "/openapi.json"},
		},
		Authenticate: NewAuthenticateHandler(e.Authenticate, mux, decoder, encoder, errhandler, formatter),
		Openapi3JSON: http.FileServer(fileSystemOpenapi3JSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "authenticate" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Authenticate = m(s.Authenticate)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return authenticate.MethodNames[:] }

// Mount configures the mux to serve the authenticate endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAuthenticateHandler(mux, h.Authenticate)
	MountOpenapi3JSON(mux, goahttp.Replace("", "/openapi3.json", h.Openapi3JSON))
}

// Mount configures the mux to serve the authenticate endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountAuthenticateHandler configures the mux to serve the "authenticate"
// service "authenticate" endpoint.
func MountAuthenticateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/auth/authentication/auth", f)
}

// NewAuthenticateHandler creates a HTTP handler which loads the HTTP request
// and calls the "authenticate" service "authenticate" endpoint.
func NewAuthenticateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAuthenticateRequest(mux, decoder)
		encodeResponse = EncodeAuthenticateResponse(encoder)
		encodeError    = EncodeAuthenticateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "authenticate")
		ctx = context.WithValue(ctx, goa.ServiceKey, "authenticate")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountOpenapi3JSON configures the mux to serve GET request made to
// "/openapi.json".
func MountOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}