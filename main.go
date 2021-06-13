package main

import (
	"fmt"
	"net/http"

	"github.com/yakushou730/go-web-development/views"

	"github.com/gorilla/mux"
)

var homeTemplate *views.View
var contactTemplate *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is faq page</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Ops, something went wrong.</h1>")
}

func main() {
	homeTemplate = views.NewView("views/home.gohtml")
	contactTemplate = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
