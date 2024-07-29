package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from %s article with ID %s", vars["category"], vars["id"])

}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is settings of an article with ID %s", mux.Vars(r)["id"])
}

func detailsHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is details of an article with ID %s", mux.Vars(r)["id"])
}

func QueriesHandler(w http.ResponseWriter, r *http.Request){

	queryParam:=r.URL.Query()
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Got parameter id:%s!\n",queryParam["id"])
	fmt.Fprintf(w,"Got parameter category:%s!\n",queryParam["category"])


}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	s := r.PathPrefix("/articles").Subrouter() //common to both  settings and details routes

	s.HandleFunc(":{id}/settings", settingsHandler)
	s.HandleFunc(":{id}/details", detailsHandler)

	r.HandleFunc("/articles",QueriesHandler)
	r.Queries("id", "category") // query based path



	// host based 
	r.Host("aaa.bbb.ccc")
	// r.HandleFunc("/id1/id2/id3",MyHandler)

	
	server := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
