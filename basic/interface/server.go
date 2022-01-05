package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type skHttpServer struct {
	Name string
}

func (s *skHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	panic("implement me")
}

func (s *skHttpServer) Start(address string) error {
	panic("implement me")
}
