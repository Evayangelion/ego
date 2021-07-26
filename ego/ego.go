package ego

import (
    "fmt"
    "net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
    // define a map which key-value pair is "string-HandleFunc" called "router"
    router map[string]HandlerFunc
}

func New() *Engine {
    return &Engine{router: make(map[string]HandlerFunc)}
}

// define a mathod named "addRoute" for "Engine" class(or you can call is struct)
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
    key := method + "-" + pattern
    engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
    engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
    engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
    return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    key := req.Method + "-" + req.URL.Path
    if handler, ok := engine.router[key]; ok {
        handler(w, req)
    } else {
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
}
