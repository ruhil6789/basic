package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type City struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {

	// if method is post
	if r.Method == "POST" {
		var tempCity City

		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		fmt.Fprintf(w, "City: %s, Area: %d", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

// middleware to checkcontent
func filterContentType(handler http.Handler) http.Handler {

	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("currently in the ceck content type middleware")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - please send json Unsupported MediaType"))
			return
		}
		handler.ServeHTTP(w, r)
	})

}

// middleware to set server time cookie
func setServerTimeCookie(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("set server time ")

	})

}

func main() {

	// http.HandleFunc("/city",mainLogic)
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))

	// chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	// http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)

}
