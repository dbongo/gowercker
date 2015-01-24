package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = flag.String("p", "3000", "server port")
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	return router
}

func routerHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

// IndexHandler ...
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'hello':'wercker!'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	flag.Parse()
	handler := routerHandler(router())
	if err := http.ListenAndServe(":"+*port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
