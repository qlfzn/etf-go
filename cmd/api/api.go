package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	// define services and dependencies
	Addr string
}

func (a *Application) NewServer() http.Handler {
	r := chi.NewRouter()

	// define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello it's finlio here!"))
	})

	return r

}

func (a *Application) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr: a.Addr,
		Handler: mux,
	}

	log.Println("app is listening at http://localhost:8080")

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
