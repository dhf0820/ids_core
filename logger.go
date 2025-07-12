package main

// import (
// 	"fmt"
// 	vsLog "github.com/dhf0820/vslog"
// 	"log"
// 	"net/http"
// 	"time"
// )

import (
	"fmt"
	vsLog "github.com/dhf0820/vslog"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		vsLog.Debug1(fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		inner.ServeHTTP(w, r)

// 		vsLog.Info(fmt.Sprintf(
// 			"%s\t%s\t%s\t%s",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			time.Since(start),
// 		))
// 	})
// }

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("\n\n\n")
// 		vsLog.Info(fmt.Sprintf(
// 			"%s\t%s\t%s\t",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			//time.Since(start),
// 		))
// 		// vsLog.Info("Handle Request: " + name)
// 		fmt.Printf("\n\n\n")
// 		start := time.Now()
// 		inner.ServeHTTP(w, r)

// 		vsLog.Info(fmt.Sprintf(
// 			"%s\t%s\t%s\t%s",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			time.Since(start),
// 		))
// 	})
// }
