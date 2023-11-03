package main

import (
	"context"
	"fmt"
	"graphql/authentication"
	"graphql/database"
	"graphql/graph"
	"graphql/repository"
	"graphql/service"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8080"

func main() {
	svc, err := StartApp()
	if err != nil {
		log.Info().Err(err).Msg("could not startapp")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: svc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}
func StartApp() (service.UserService, error) {
	log.Info().Msg("Main: Started: Intilaizing authentication support")
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return &service.Service{}, fmt.Errorf("reading the auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return &service.Service{}, fmt.Errorf("parsing private key %w", err)
	}
	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return &service.Service{}, fmt.Errorf("reading the auth public key %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return &service.Service{}, fmt.Errorf("parsing public key %w", err)
	}
	a, err := authentication.NewAuth(privateKey, publicKey)
	if err != nil {
		return &service.Service{}, fmt.Errorf("constructing auth %w", err)
	}
	db, err := database.Open()
	if err != nil {
		return &service.Service{}, fmt.Errorf("connecting to database %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, fmt.Errorf("database is not connected: %w ", err)
	}
	repo, err := repository.NewRepository(db)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize repo layer: %w ", err)
	}
	svc, err := service.NewService(a, repo)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize service layer: %w ", err)
	}
	return svc, nil
}
