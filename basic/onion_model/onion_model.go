package main

import "net/http"

type HTTPMiddlerware func(handler http.Handler) http.Handler

func HTTPChain(opt ...HTTPMiddlerware) HTTPMiddlerware {
	return nil
}

