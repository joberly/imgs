package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joberly/imgs/ctlrs"
)

func main() {
	staticCtlr := ctlrs.NewStatic()
	usersCtlr := ctlrs.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticCtlr.Home).Methods("GET")
	r.Handle("/contact", staticCtlr.Contact).Methods("GET")
	r.HandleFunc("/signup", usersCtlr.New).Methods("GET")
	r.HandleFunc("/signup", usersCtlr.Create).Methods("POST")
	r.NotFoundHandler = staticCtlr.NotFound
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
