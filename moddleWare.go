package main
import (
	"time"
	"fmt"
	"net/http"

)
func Index(w http.ResponseWriter,r *http.Request){
	fmt.Println("调用方法", r.Method)
	fmt.Println("Bye Index")
}
func login(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		start:=time.Now()
		fmt.Println("调用中间件")
		next.ServeHTTP(w,r)
		fmt.Println("调用结束",time.Since(start))
	})
}
func main(){
	index:= http.HandlerFunc(Index)
	http.Handle("/",login(index))
	server:=&http.Server{
		Addr:":8080",
	}
	fmt.Println("Listening Port is ",server.Addr)
	server.ListenAndServe()
}