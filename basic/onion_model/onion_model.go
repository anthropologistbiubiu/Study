package main

import (
	"log"
	"net/http"
)

type HTTPMiddlerware func(handler http.Handler) http.Handler

func HTTPChain(opts ...HTTPMiddlerware) HTTPMiddlerware {
	return func(next http.Handler) http.Handler {
		for i := len(opts) - 1; i >= 0; i-- {
			next = opts[i](next)
		}
		return next
	}
}

func LoggerMiddler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(r http.ResponseWriter, w *http.Request) {
		log.Println("befor handler")
		next.ServeHTTP(r, w)
		log.Println("afeter handler")
	})
}

type Router struct {
	routers map[string]string
}

func (router *Router) ServeHTTP(r http.ResponseWriter, w *http.Request) {

}

