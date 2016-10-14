package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/franela/play-with-docker/handlers"
	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
)

func main() {
	mux := bone.New()

	mux.Get("/ping", http.HandlerFunc(handlers.Ping))
	mux.Get("/", http.HandlerFunc(handlers.NewSession))
	mux.Get("/sessions/:sessionId", http.HandlerFunc(handlers.GetSession))
	mux.Post("/sessions/:sessionId/instances", http.HandlerFunc(handlers.NewInstance))
	mux.Delete("/sessions/:sessionId/instances/:instanceName", http.HandlerFunc(handlers.DeleteInstance))

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./www/index.html")
	})
	mux.Get("/p/:sessionId", h)
	mux.Get("/assets/*", http.FileServer(http.Dir("./www")))

	mux.Get("/sessions/:sessionId/instances/:instanceName/attach", websocket.Handler(handlers.Exec))

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", n))

}