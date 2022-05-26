package http

import (
	"errors"
	"fmt"
	"net/http"

	"visitor/internal/http/handler"

	"github.com/go-chi/chi"
)

type Handlers struct {
	Auth *handler.Auth

	APITheme   *handler.APITheme
	APIVisitor *handler.APIVisitor

	Page *handler.Page
}

type Server struct {
	http.Server

	router *chi.Mux
	hs     *Handlers
}

func NewServer(addr string, hs *Handlers) *Server {
	r := chi.NewRouter()

	s := &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: r,
		},
		router: r,
		hs:     hs,
	}

	s.routeStatus()
	s.routeAPI()
	s.routeWeb()

	return s
}

func (s *Server) routeStatus() {
	s.router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})
}

func (s *Server) routeAPI() {
	s.router.Route("/api", func(r chi.Router) {
		r.Get("/theme", s.hs.Auth.Do(s.hs.APITheme.Get))
		r.Put("/theme", s.hs.Auth.Do(s.hs.APITheme.Update))

		r.Get("/visitors", s.hs.Auth.Do(s.hs.APIVisitor.GetAll))
		r.Post("/visitor", s.hs.Auth.Do(s.hs.APIVisitor.Save))
	})
}

func (s *Server) routeWeb() {
	s.router.Get("/", s.hs.Page.Load)
	s.router.Post("/", s.hs.APIVisitor.Save)
}

func (s *Server) Run() error {
	err := s.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
