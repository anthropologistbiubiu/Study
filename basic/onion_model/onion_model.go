package main

import (
	"fmt"
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
	routers map[string]http.Handler
}

func (router *Router) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	handler, ok := router.routers[w.URL.Path]
	if ok {
		handler.ServeHTTP(r, w)
	} else {
		http.NotFound(r, w)
	}
}

func (router *Router) Handler(path string, handler http.Handler) {
	router.routers[path] = handler
}

func main() {
	router := &Router{
		routers: make(map[string]http.Handler),
	}
	router.Handler("/hello", http.HandlerFunc(func(r http.ResponseWriter, w *http.Request) {
		r.Write([]byte("hello world!"))
	}))
	composer := HTTPChain(LoggerMiddler)(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: composer,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("run failed", err)
	} else {
		fmt.Println("server is run on 8081")
	}

}
