package api

import (
	mux "github.com/gorilla/mux"
	"time"
	"log"
	"net/http"
	"fmt"
)


type controller struct {}

func (c *controller) staticHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "from static handler %d\n", http.StatusOK)
}

func (c *controller) Routes() Routes {
	return Routes{
		{
			"product",
			"POST",
			"/product",
			handle,

		},
		{
			"static",
			"GET",
			"/static/",
			c.staticHandler,
		},
		{
			"hello",
			"GET",
			"/hello",
			handle,
		},
	}
}

func InitServer(host string){

	c := &controller{}
	r := mux.NewRouter().StrictSlash(false)
	routes := c.Routes()

	for _, route := range routes{

		name := route.Name
		path := route.Pattern
		handler := route.HandlerFunc
		method := route.Method
		log.Println("reg :", name, path, handler, method)

		r.HandleFunc(path, handler).Methods(method).Schemes("http")
	}

	server := &http.Server{
		Addr:		host,
		Handler:	r,
		WriteTimeout:	time.Second * 15,
		ReadTimeout:	time.Second * 15,
	}

	log.Fatal(server.ListenAndServe())
}
