package api

import (
	"fmt"
	"net/http"
)

//endpoint
type Route struct{
	Name	string
	Method	string
	Pattern string
	HandlerFunc http.HandlerFunc
}

//collection of endpoints
type Routes []Route


func handle(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Method, " URI ", r.RequestURI)
	fmt.Fprintf(w, "hello %d\n", http.StatusOK)
}
