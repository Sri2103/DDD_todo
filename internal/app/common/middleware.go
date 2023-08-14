package common

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w,r)
		elapsed:= time.Since(start)
		log.Printf(" %s %s %s %s ", r.Method, r.RemoteAddr,r.RequestURI,elapsed)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if err := recover(); err != nil {
				log.Printf("Panic occurred:%v",err)
				http.Error(w, "Internal server error",http.StatusInternalServerError)

			}
		}()
		next.ServeHTTP(w,r)			
	})
}