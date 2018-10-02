package service

import "net/http"

// Defines a single route, name, HTT method and the pattern the function will
// execute when the route is called
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just a slice of Route structs
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetAccount", // Name
		"GET",        // HTTP method
		"/accounts/{accountId}", // Route pattern
		GetAccount,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
