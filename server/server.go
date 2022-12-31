package main

import (
	"github.com/AntonioTrupac/socialHabitsTracker/database"
	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	resolvers "github.com/AntonioTrupac/socialHabitsTracker/graph/resolvers"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"github.com/AntonioTrupac/socialHabitsTracker/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	bookRepo := repository.NewBookService(db)
	userRepo := repository.NewUserService(db)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		BookRepository: bookRepo,
		UserRepository: userRepo,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	err := godotenv.Load(".env")

	db, err := database.InitDb()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Address{}, &models.Role{}, &models.UserRoles{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
