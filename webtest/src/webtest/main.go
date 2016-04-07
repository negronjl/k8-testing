package main

import (
	"fmt"
	"html"
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func MyRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Home",
		"GET",
		"/",
		HomeHandler,
	},
	Route{
		"Login",
		"POST",
		"/login",
		LoginHandler,
	},
	Route{
		"Metrics",
		"POST",
		"/metrics",
		MetricsHandler,
	},
}

func main() {
	router := MyRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	deviceid := mux.Vars(r)["deviceid"]
	fmt.Fprintf(w, "/login - device: %s", deviceid) 
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	deviceid := mux.Vars(r)["deviceid"]
	timestamp := time.Now()
	fmt.Fprintf(w, "/metrics - device: %s, timestamp: %s", deviceid, timestamp.Format("2016-04-01 12:06:46.153985"))
}

