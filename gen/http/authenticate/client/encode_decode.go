// Code generated by goa v3.10.1, DO NOT EDIT.
//
// authenticate HTTP client encoders and decoders
//
// Command:
// $ goa gen auth/design

package client

import (
	authenticate "auth/gen/authenticate"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildAuthenticateRequest instantiates a HTTP request object with method and
// path set to call the "authenticate" service "authenticate" endpoint
func (c *Client) BuildAuthenticateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthenticateAuthenticatePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("authenticate", "authenticate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthenticateRequest returns an encoder for requests sent to the
// authenticate authenticate server.
func EncodeAuthenticateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*authenticate.InputAuth)
		if !ok {
			return goahttp.ErrInvalidType("authenticate", "authenticate", "*authenticate.InputAuth", v)
		}
		values := req.URL.Query()
		if p.XProvider != nil {
			values.Add("X-Provider", fmt.Sprintf("%v", *p.XProvider))
		}
		req.URL.RawQuery = values.Encode()
		body := NewAuthenticateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("authenticate", "authenticate", err)
		}
		return nil
	}
}

// DecodeAuthenticateResponse returns a decoder for responses returned by the
// authenticate authenticate endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeAuthenticateResponse may return the following errors:
//	- "badrequest" (type *authenticate.TextMessage): http.StatusBadRequest
//	- "unauthorized" (type *authenticate.TextMessage): http.StatusUnauthorized
//	- "internalservererror" (type *authenticate.TextMessage): http.StatusInternalServerError
//	- error: internal error
func DecodeAuthenticateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AuthenticateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("authenticate", "authenticate", err)
			}
			err = ValidateAuthenticateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("authenticate", "authenticate", err)
			}
			res := NewAuthenticateOutputAuthOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body AuthenticateBadrequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("authenticate", "authenticate", err)
			}
			err = ValidateAuthenticateBadrequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("authenticate", "authenticate", err)
			}
			return nil, NewAuthenticateBadrequest(&body)
		case http.StatusUnauthorized:
			var (
				body AuthenticateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("authenticate", "authenticate", err)
			}
			err = ValidateAuthenticateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("authenticate", "authenticate", err)
			}
			return nil, NewAuthenticateUnauthorized(&body)
		case http.StatusInternalServerError:
			var (
				body AuthenticateInternalservererrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("authenticate", "authenticate", err)
			}
			err = ValidateAuthenticateInternalservererrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("authenticate", "authenticate", err)
			}
			return nil, NewAuthenticateInternalservererror(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("authenticate", "authenticate", resp.StatusCode, string(body))
		}
	}
}