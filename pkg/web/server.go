package web

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

//go:embed ui/public
var staticFiles embed.FS

func NewUIServer(port string) *Server {
  var staticFS = fs.FS(staticFiles)
	staticContent, err := fs.Sub(staticFS, "ui/public")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(staticContent))

  r := chi.NewRouter()
	r.Use(middleware.Logger)
  r.Mount("/", fs)

  return &Server{
    router: r,
    port: port,
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
    port: port,
  }
}

type Server struct {
  router *chi.Mux
  port string
}

func (s *Server) ListenAndServe() error {
  log.Printf("serving on port 0.0.0.0:%s", s.port)
  return http.ListenAndServe(":"+s.port, s.router)
}
