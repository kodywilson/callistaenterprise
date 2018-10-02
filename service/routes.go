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
		"Get/Account", // name
		"GET",         // HTTP Method
		"/accounts/{accountId}", // Route Pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
