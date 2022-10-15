package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("api", func() {
	Title("Staffea Authentication API")
	Description("This API permits to authenticate or to refresh a token")
	Version("0.1.0")

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("api", func() {

		// List the Hosts and their transport URLs.
		Host("localhost", func() {
			Description("localhost hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:8080")
		})
	})
})
