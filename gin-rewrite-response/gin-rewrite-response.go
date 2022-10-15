package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/**
 * The purpose of this test is to solve a problem during a request on a gin server. The gin server receives requests
 * and acts as a reverse proxy for the requested url.
 * The entry point is a json mock server http://dummyjson.com.
 * We want to add in the output payload a new signature field. We use a middleware that modifies the writing context
 * to retrieve and modify it.
 * However, an error 500 on the call of http://localhost:8080/products/1
 *
 * How to solve it?
 * Launch the test: go run test.go

 * Note: this system only works for responses with a single json object. However, how to add the signature in
 * each element if the response is an array of objects.
 */

// BufferedResponseWriter is a writing context to replace that of gin by providing additional methods for
// managing the write buffer response. The default context does not provide methods for editing and accessing content.
type BufferedResponseWriter struct {
	Buf bytes.Buffer
	gin.ResponseWriter
}

func (b *BufferedResponseWriter) WriteString(s string) (int, error) {
	b.Header().Del("Content-Length")
	return b.Buf.WriteString(s)
}

func (b *BufferedResponseWriter) Write(buf []byte) (int, error) {
	b.Header().Del("Content-Length")
	return b.Buf.Write(buf)
}

// main application start, creates a gin server listening and :8080.
func main() {
	r := gin.Default()

	// Signature middleware is applied for each request
	r.Use(signature())

	// Requests are proxified using httputil.ReverseProxy
	r.Any("/*proxyPath", proxy())

	r.Run()
}

// proxy middleware will redirect request to json placeholder server https://dummyjson.com
func proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse("https://dummyjson.com")
		if err != nil {
			panic(err)
		}

		c.Request.Header.Del("If-None-Match")

		pxy := httputil.NewSingleHostReverseProxy(remote)
		pxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		pxy.ServeHTTP(c.Writer, c.Request)
		c.Next()
	}
}

// signature middleware add a custom signature field in the output payload with auto-generated uuid
func signature() gin.HandlerFunc {
	return func(c *gin.Context) {

		// buf replace the standard writer buffer to add new functions to manipulate him
		buf := &BufferedResponseWriter{
			ResponseWriter: c.Writer,
		}
		c.Writer = buf

		c.Next()

		// Restore the original response writer
		c.Writer = buf.ResponseWriter

		// Retrieve the response body to manipulate him
		var m map[string]interface{}
		if err := json.Unmarshal(buf.Buf.Bytes(), &m); err != nil {
			log.Panic(err)
			return
		}

		// Add the signature
		m["uuid"] = uuid.New()
		b, _ := json.Marshal(m)

		// Wrote the new content with the signature
		_, err := c.Writer.Write(b)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}
