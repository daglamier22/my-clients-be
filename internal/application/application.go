package application

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/daglamier22/my-clients-be/internal/services"
	"github.com/daglamier22/my-clients-be/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Application struct {
	Config  Config
	Db      *sql.DB
	Store   store.Storage
	Service services.Service
	Server  *http.Server
}

type Config struct {
	Addr string
	Db   DbConfig
}

type DbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

func (app *Application) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", app.signupHandler)
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/", app.GetAllUsers)
		})
	})

	return r
}

func (app *Application) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Store = store.NewStorage(app.Db)
	app.Service = services.NewServices(app.Store)

	log.Printf("Server is running at %s\n\r", app.Config.Addr)

	app.Server = srv

	return srv.ListenAndServe()
}
