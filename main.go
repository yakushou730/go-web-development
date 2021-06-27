package main

import (
	"fmt"
	"net/http"

	"github.com/yakushou730/go-web-development/middleware"

	"github.com/yakushou730/go-web-development/models"

	"github.com/yakushou730/go-web-development/controllers"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "TsengYaoShang"
	password = ""
	dbname   = "go_web_dev"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Ops, something went wrong.</h1>")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	services, err := models.NewServices(psqlInfo)

	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, r)

	requireUserMw := middleware.RequireUser{
		UserService: services.User,
	}
	newGallery := requireUserMw.Apply(galleriesC.New)
	createGallery := requireUserMw.ApplyFn(galleriesC.Create)

	r.Handle("/", staticC.Home).Methods(http.MethodGet)
	r.Handle("/contact", staticC.Contact).Methods(http.MethodGet)
	r.Handle("/faq", staticC.Faq).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.New).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.Create).Methods(http.MethodPost)
	r.Handle("/login", usersC.LoginView).Methods(http.MethodGet)
	r.HandleFunc("/login", usersC.Login).Methods(http.MethodPost)
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods(http.MethodGet)
	r.Handle("/galleries/new", newGallery).Methods(http.MethodGet)
	r.HandleFunc("/galleries", createGallery).Methods(http.MethodPost)
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods(http.MethodGet).
		Name(controllers.ShowGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", requireUserMw.ApplyFn(galleriesC.Edit)).
		Methods(http.MethodGet)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
