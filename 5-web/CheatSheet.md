# Cheat Sheet

## Handler Example

```go
package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %v.", r.URL.Path[1:])
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

```

## Routes
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello")
})

http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Welcome")
})

http.ListenAndServe(":8080", nil)
```

## JSON

```go
import "encoding/json"

type Person struct {
	Name string `json:"name"`
	age  int    `json:"age"`// private! won't serialize
}

// GET
p := Person{"Gopher", 4}
data, _ := json.Marshal(p)
w.Write(data)

// POST
body, err := ioutil.ReadAll(r.Body)
var p Person
err := json.Unmarshal(body, &p)

```

## Templates
```go
type Data struct { Path string }

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Template error")
		return
	}

	t.Execute(w, &Data{r.URL.Path[1:]})
})

```

## Static Files
```go
http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
})

```


## Gorilla Mux

https://github.com/gorilla/mux

```go
r := mux.NewRouter()
r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Welcome %v!", name)
}).Methods("GET")

http.Handle("/", r)
```