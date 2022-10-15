package design

import (
	. "goa.design/goa/v3/dsl"
)

var ErrorResponse = Type("TextMessage", func() {
	Description("returns details about the current error.")
	Field(1, "status", Int, "message status")
	Field(2, "message", String, "message")
	Field(3, "exception", String, "Name of the error", func() {
		Meta("struct:error:name")
	})
	Required("status", "message", "exception")
})
