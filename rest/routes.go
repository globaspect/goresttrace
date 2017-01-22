package rest

import (
	"net/http"
	"time"
)

// define struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// array of struct Route
type Routes []Route

var routes = Routes{
	Route{
		"test " + time.Now().String(),
		"GET",
		"/health",
		healthHandler,
	},
	Route{
		"catch trace message ...",
		"POST",
		"/trace",
		traceInputHandler,
	},
	Route{
		"catch trace message ...",
		"POST",
		"/trace/string",
		traceInputHandlerString,
	},

}
