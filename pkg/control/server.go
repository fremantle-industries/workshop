package control

import (
	"embed"
	"fmt"
	"html"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	//go:embed ui/dist/assets/*
	assets embed.FS

	//go:embed ui/dist/index.html
	index []byte

	//go:embed ui/dist/favicon.svg
	favicon []byte
)

func NewUIServer(port string) *Server {
	assetsContent, err := fs.Sub(assets, "ui/dist/assets")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle(
		"/assets/*",
		http.StripPrefix("/assets/", http.FileServer(http.FS(assetsContent))),
	)
	r.Get("/favicon.svg", func(w http.ResponseWriter, r *http.Request) {
		// TODO: can this come from the file system?
		w.WriteHeader(200)
		w.Write(favicon)
	})
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// TODO: only match when there is no file extension
		// TODO: can this come from the file system?
		w.WriteHeader(200)
		w.Write(index)
	})

	return &Server{
		router: r,
		port:   port,
	}
}

func NewAPIServer(port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return &Server{
		router: r,
		port:   port,
	}
}

type Server struct {
	router *chi.Mux
	port   string
}

func (s *Server) ListenAndServe() error {
	log.Printf("serving on port 0.0.0.0:%s", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}
