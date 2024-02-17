package api

import (
	"fmt"
	"time"
	"net/http"
)

func Logger(f http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		f.ServeHTTP(w, r)
		fmt.Println("adding log", name, r.Method, r.RequestURI, time.Since(start))
	})
}
