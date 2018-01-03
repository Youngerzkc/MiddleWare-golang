package main
import (
	"fmt"
	"log"
	"net/http"	
)
func index(w http.ResponseWriter,r *http.Request){
	log.Println("Execute index Handler")
	fmt.Fprintf(w,"Welcome!")
}
func message(w http.ResponseWriter,r *http.Request){
	log.Println("Execute message Handler")
	fmt.Fprintln(w,"Message Go!")
}
func middlewareFirst(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		log.Println("MiddlewareFirst-Before Handler")
		next.ServeHTTP(w,r)
		log.Println("MiddlerwareFirst-After Handler")
	})
}
func middlewareSecond(next http.Handler) http.Handler{

	   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		   log.Println("MiddlerwareSecond-Before Handler")
		   if r.URL.Path=="/message"{
			   if r.URL.Query().Get("password")=="123456"{
				   log.Println("Authorized to system")
				   next.ServeHTTP(w,r)
			   }else{
				   log.Println("failed to Authorized to system")
			   }
		   }else{
			   next.ServeHTTP(w,r)
		   }
		   log.Println("MiddlewareSecond-After Handler")
	   })
}
func main(){
	http.Handle("/",middlewareFirst(middlewareSecond(http.HandlerFunc(index))))
	http.Handle("/message",middlewareFirst(middlewareSecond(http.HandlerFunc(message))))
	server:=http.Server{
		Addr:":8080",
	}
	log.Println("Listening....")
	server.ListenAndServe()
}
