package main

type Router struct{
	rules: map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router{
		rules: make(map[string]http.HandlerFunc)
	}
}