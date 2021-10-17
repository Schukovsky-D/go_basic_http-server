// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Links []Link
}

type Link struct {
	Title string
	Url   string
}

func materialsHandler(w http.ResponseWriter, req *http.Request) {
	data := ViewData{
		Title: "Materials",
		Links: []Link{
			{Title: "Go by Example: HTTP Servers", Url: "https://gobyexample.com/http-servers"},
			{Title: "http package", Url: "https://golang.org/pkg/net/http/"},
			{Title: "Writing Web Applications", Url: "https://golang.org/doc/articles/wiki/"},
			{Title: "Hello world HTTP server example", Url: "https://yourbasic.org/golang/http-server-example/"},
			{Title: "How I write Go HTTP services after seven years", Url: "https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831"},
			{Title: "Gorilla web toolkit", Url: "https://www.gorillatoolkit.org/pkg/mux"},
		},
	}

	tmpl, _ := template.ParseFiles("templates/materials.html")
	tmpl.Execute(w, data)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	//literally returns an empty page with a 200 code
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/materials", materialsHandler)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8090", nil)
}
