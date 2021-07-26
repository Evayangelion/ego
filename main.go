package main

import (
    "fmt"
    "net/http"

    "ego"
)

func main() {
    r := ego.New()

    //lht, 2021-07-26, refactor
    /*
    r.GET("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    })

    r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
        for k, v := range req.Header {
            fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
    })
    */
    r.GET("/", indexHandler)
    r.GET("/hello", helloHandler)
    //end of refactor
    r.Run(":9999")
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
    for k, v := range req.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
}
