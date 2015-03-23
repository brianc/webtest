package test

import (
	"github.com/go-martini/martini"
	"net/http"
)

func MartiniServer() http.Handler {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	return m
}
