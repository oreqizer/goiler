//go:generate sqlboiler --wipe --no-hooks --no-tests psql
//go:generate go run github.com/99designs/gqlgen

package main

import (
	"context"
	"database/sql"
	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/handler"
	"github.com/getsentry/raven-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/oreqizer/goiler/config"
	"github.com/oreqizer/goiler/generated"
	"github.com/oreqizer/goiler/graphql"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/cache"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/schemas"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

const (
	complexity = 250
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// CORS
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.Cors},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler)

	r.Use(middleware.Logger)
	r.Use(raven.Recoverer)

	// DB
	dbi, err := sql.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer dbi.Close()

	if err := dbi.Ping(); err != nil {
		log.Fatal(err)
	}

	// Firebase
	opt := option.WithCredentialsFile(cfg.Firebase)
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(db.Middleware(dbi))
	r.Use(auth.Middleware(dbi, fb))
	r.Use(schemas.Middleware())

	resolver := generated.Config{Resolvers: &graphql.Resolver{}}

	graphql.Directives(&resolver)
	graphql.Complexity(&resolver)

	schema := generated.NewExecutableSchema(resolver)

	r.Get("/", handler.Playground("GraphQL playground", "/graphql"))
	r.Post("/graphql", handler.GraphQL(
		schema,
		handler.ComplexityLimit(complexity),
		handler.EnablePersistedQueryCache(cache.New(24*time.Hour)),
	))

	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
