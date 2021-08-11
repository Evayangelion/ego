package ego

import (
    //"fmt"
    "net/http"
    //"context"
)

// using Context instead, refactor 2021-07-29
// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

type Engine struct {
    // define a map which key-value pair is "string-HandleFunc" called "router"
    // using Context instead, refactor 2021-07-29
    // router map[string]HandlerFunc
    router *router
}

func New() *Engine {
    // using Context instead, refactor 2021-07-29
    //return &Engine{router: make(map[string]HandlerFunc)}
    return &Engine{router: newRouter()}
}

// define a mathod named "addRoute" for "Engine" class(or you can call is struct)
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
    /* using Context instead, refactor 2021-07-29
    key := method + "-" + pattern
    engine.router[key] = handler
    */
    engine.router.addRoute(method, pattern, handler)
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
    /*  using Context instead, refactor 2021-07-29
    key := req.Method + "-" + req.URL.Path
    if handler, ok := engine.router[key]; ok {
        handler(w, req)
    } else {
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
    */
    c := newContext(w, req)
    engine.router.handle(c)
}
