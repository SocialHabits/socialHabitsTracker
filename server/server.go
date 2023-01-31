package main

import (
	"fmt"

	"github.com/AntonioTrupac/socialHabitsTracker/database"
	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	resolvers "github.com/AntonioTrupac/socialHabitsTracker/graph/resolvers"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"github.com/AntonioTrupac/socialHabitsTracker/repository"
	"github.com/gin-contrib/cors"
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
	moodRepo := repository.NewMoodService(db)
	habitRepo := repository.NewHabitService(db)

	c := generated.Config{Resolvers: &resolvers.Resolver{
		BookRepository:  bookRepo,
		UserRepository:  userRepo,
		MoodRepository:  moodRepo,
		HabitRepository: habitRepo,
	}}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

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

	if err != nil {
		panic(err)
	}

	db, err := database.InitDb()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Book{}, &models.User{}, &models.Address{}, &models.Mood{}, &models.Habit{})

	if err != nil {
		fmt.Printf("Error while migrating: %v", err)
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(middleware.AuthMiddleware())
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	err = r.Run(":" + port)

	if err != nil {
		fmt.Printf("Error while running server: %v", err)
		panic(err)
	}
}
