package main

import (
    //"fmt"
    "net/http"
    //"context"
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

    /* using Context instead, refactor, 2021-07-29
    r.GET("/", indexHandler)
    r.GET("/hello", helloHandler)
    */
    //end of refactor
    r.GET("/", func(c *ego.Context) {
        c.HTML(http.StatusOK, "<h1>HELLO EGO</h1>")
        })
    r.GET("/hello", func(c *ego.Context) {
        // expect /hello?name=ego
        c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"),c.Path)
    })

    r.POST("/login", func(c *ego.Context) {
        c.JSON(http.StatusOK, ego.NewJSON {
            "username": c.PostForm("username"),
            "password": c.PostForm("password"),
        })
    })
 
    r.Run(":9999")
}

// handler echoes r.URL.Path
/* using Context instead, refactor, 2021-07-29
func indexHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}


// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
    for k, v := range req.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
}
*/

