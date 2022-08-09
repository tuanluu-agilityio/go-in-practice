package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

// Adds paths to internal lookup
type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

// Creates new initialized pathResolver
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func main() {
	// Gets an instance of a path-based router
	pr := newPathResolver()
	// Maps functions to paths
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	// Sets the HTTP server to use your router
	http.ListenAndServe(":8080", pr)
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Constructs our method + path to check
	check := req.Method + " " + req.URL.Path
	// Iterates over registerd paths
	for pattern, handlerFunc := range p.handlers {
		// Checks whether current path matches a registered one
		if ok, err := path.Match(pattern, check); ok && err == nil {
			// Executes the handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	// If no path matches, the page wasn't found.
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Tuan Luu"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Tuan Luu"
	}
	fmt.Fprint(res, "Goodbye ", name)
}