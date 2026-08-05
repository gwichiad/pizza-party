package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type api struct {
	config config
}
type config struct {
	addr string
	store *store
}

func (api *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/data", api.ingestHandler)

	/*r.Route("/data/", func(r chi.Router) {
		r.Get("/{field}", api.dataHandler)
		r.Get("/{field}/{subfield}", api.dataHandler)
	})*/

	return r
}

func (api *api) run(mux http.Handler) error {
	srv := &http.Server{
		Addr: api.config.addr,
		Handler: mux,
	}

	log.Printf("Server has started at %s", api.config.addr)

	return srv.ListenAndServe()
}
