package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func icoHandler(w http.ResponseWriter, r *http.Request) {

}
func index(w http.ResponseWriter, r *http.Request) {

	log.Println("Execte index handler")
	fmt.Fprintf(w, "welcome!")
}
func about(w http.ResponseWriter, r *http.Request) {

	log.Println("Execute about handler")
	fmt.Fprintf(w, "About GoLang!")
}
func logginHandler(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		star := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) 
		log.Printf("COmpleted %s in %v", r.URL.Path, time.Since(star))
	})

}

func main() {
	
	http.HandleFunc("favicon.icon", icoHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", logginHandler(indexHandler))
	http.Handle("/about", logginHandler(aboutHandler))

	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening....")
	server.ListenAndServe()
	// http.ListenAndServe("8080",nil)
}
