package main

import (
	"fmt"
	"net/http"

	"github.com/yakushou730/go-web-development/controllers"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Ops, something went wrong.</h1>")
}

func main() {
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods(http.MethodGet)
	r.Handle("/contact", staticC.Contact).Methods(http.MethodGet)
	r.Handle("/faq", staticC.Faq).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.New).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.Create).Methods(http.MethodPost)
	r.HandleFunc("/galleries/new", galleriesC.New).Methods(http.MethodGet)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
