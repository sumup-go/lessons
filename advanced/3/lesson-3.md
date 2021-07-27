# Lesson 3

## Modules

Modules are how Go manages dependencies. A module is a collection of packages that are released, versioned and distributed together. Modules may be downloaded directly from version control repositories or from proxy servers.

A module is identified by a module path, which is declared in a `go.mod` file. The *module root directory* is the directory that contains `go.mod` file. 

Start your project by creating a new directory and initializing a new module.

    mkdir demo
    go mod init demo

## Web Application

**Router or HTTP request multiplexer** -  matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL

We can either use a stdlib router:

```go
mux := http.NewServerMux()
```

Or we can use some third-party router, like Gorilla:

```go
import "github.com/gorilla/mux"

func somefunc() {
    router := mux.NewRouter()
}
```

**Handler** - is function that responds to an HTTP Request. Functions serving as handlers take `http.ResponseWriter` and `http.Request` as arguments. The response writer is used to fill in the HTTP response.**Handler Function** must have a following signature:

    func(http.ResponseWriter, *http.Request)

Golang provides a function that serves static HTML files:

    http.ServeFile(w, r, "./index.html")

A **web server** is an entity that accepts and serves requests via HTTP. The function `http.ListenAndServe` starts the web server and listens the incoming HTTP requests on the specified port. The first parameter specifies a port, and the second one specifies the handler.

    http.ListenAndServe(":8080", handler)

Now, let's create a simple web server that will serve a static HTML page

**Full working example**

```go
package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/static", static)
	
	log.Fatal(http.ListenAndServe(":9999", m))
}
```

Now, try to open your browser and type: `localost:9999/static`. You should see a simple web page which greets you as a Go leaner?

This page is static, i.e. it will always render the same information. How can we make it more alive, and show the name of the leaner? We need to collect an input from the user through the form. 

```go
func submission(w http.ResponseWriter, r *http.Request) {
    tmp := template.Must(template.ParseFiles("index.html"))
    
    if r.Method != http.MethodPost {
        tmp.Execute(w, nil)
        return
    }
    
    name := r.FormValue("name")
    log.Println("received name", name)
    if name == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    p := Page{Name: name}
    tmp.Execute(w, p)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/submission", submission).Methods(http.MethodGet, http.MethodPost)
    
    log.Fatal(http.ListenAndServe(":9098", router))
}
```

Now, try to open "localhost:9098/submission". The page shoud intitially ask your name and then greet you personally.


## Troubleshooting

### 404 Not Found

If you run a static web server, and receive 404 Not Found error when trying to access the page, make sure:

    * The router configuration is correct. I.e. if you have `m.HandleFunc("/static", static)` then the page should be displayed at `localhost:PORT/static`.

    * The template location is correct. I.e if you have `http.ServeFile(w, r, "index.html")`, then Go will try to find the template in the directory where `go run main.go` command were invoked. Make sure that you in the root directory of your module (the directory containing `go.mod` file) and that directory contains file `index.html`.


### This site can’t be reached. localhost refused to connect.

If you receive a following error from your browser when you are trying to access the webpage:

    This site can’t be reached
    localhost refused to connect.

Then make sure that your specified the correct port for server `http.ListenAndServe(":9999", m)`. In this example, we should access a page as `localhost:9999/page-name`.

### No Required Module Provides Package

If you try to run the program and receive something like:

    no required module provides package github.com/gorilla/mux: go.mod file not found in current directory or any parent directory; see 'go help modules'

Make sure that initialized the module:

    go mod init {MODULE_NAME}

where MODULE_NAME is some string, for example:

    go mod init static

will initialize a module with name static. In `go.mod` you will see:

Note, make sure that you have go 1.16 version by running `go version` command. The behavour of modules has been evolving with different versions of Go. 



    

