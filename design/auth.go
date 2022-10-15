package design

import (
	. "goa.design/goa/v3/dsl"
)

// Service describes a service
var _ = Service("authenticate", func() {
	Description("auth service")
	Error("badrequest", ErrorResponse, "Bad Request")
	Error("unauthorized", ErrorResponse, "Unauthorized")
	Error("internalservererror", ErrorResponse, "Internal Server Error")
	Method("authenticate", func() {
		Payload(InputAuth)
		Result(OutputAuth)
		HTTP(func() {
			POST("/auth/authentication/auth")
			Param("X-Provider", Int)
			Response(StatusOK)
			Response("badrequest", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("internalservererror", StatusInternalServerError)
		})
	})
	// Serve the file gen/http/openapi3.json for requests sent to
	// /openapi.json.
	Files("/openapi.json", "openapi3.json")
})

var InputAuth = Type("InputAuth", func() {
	Attribute("X-Provider", Int, "provider")
	Attribute("login", String, "A valid login name")
	Attribute("password", String, "A valid password")
	Attribute("privacyAccepted", Boolean, "privacyAccepted")
	Attribute("uuid", String, "uuid")
	Example("sample", func() {
		Description("")
		Value(Val{
			"X-Provider":      1,
			"login":           "login1",
			"password":        "password",
			"uuid":            "eFd1CEdC-93Fc-38db-dAae-e029817F045F",
			"privacyAccepted": true,
		})
	})
	Required("login", "password")
})

var OutputAuth = Type("OutputAuth", func() {
	Attribute("login", String, "login")
	Attribute("refreshToken", String, "refreshToken")
	Attribute("providerToken", String, "providerToken")
	Attribute("subscriber", Int, "subscriber")
	Attribute("people", String, "people")
	Attribute("locale", String, "locale")
	Example("sample", func() {
		Description("")
		Value(Val{
			"login":         "login",
			"refreshToken":  "b222d1c4-4877-11ec-81d3-0242ac130003",
			"providerToken": "",
			"subscriber":    1,
			"people":        "/people/peoples/1",
			"locale":        "fr_FR",
		})
	})
	Required("login", "refreshToken", "subscriber")
})
