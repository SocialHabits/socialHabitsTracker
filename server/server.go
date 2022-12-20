package main

import (
	"github.com/AntonioTrupac/socialHabitsTracker/database"
	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	resolvers "github.com/AntonioTrupac/socialHabitsTracker/graph/resolvers"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"github.com/AntonioTrupac/socialHabitsTracker/repository"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")

	db, err := database.InitDb()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Address{}, &models.Role{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bookRepo := repository.NewBookService(db)
	userRepo := repository.NewUserService(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		BookRepository: bookRepo,
		UserRepository: userRepo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
