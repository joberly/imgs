package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joberly/imgs/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "we couldn't find what you're looking for. get support at <a href=\"mailto:joberly@gmail.com\">joberly@gmail.com</a>.")
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", handlerHome)
	r.HandleFunc("/contact", handlerContact)
	r.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
