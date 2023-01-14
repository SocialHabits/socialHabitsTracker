package main

import (
	"github.com/AntonioTrupac/socialHabitsTracker/database"
	"github.com/AntonioTrupac/socialHabitsTracker/directives"
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

func init() {
	defaultTranslation()
}

func defaultTranslation() {
	directives.ValidateAddTranslation("email", "Email is not valid")
}

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	bookRepo := repository.NewBookService(db)
	userRepo := repository.NewUserService(db)

	c := generated.Config{Resolvers: &resolvers.Resolver{
		BookRepository: bookRepo,
		UserRepository: userRepo,
	}}

	c.Directives.Binding = directives.Binding

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

	db, err := database.InitDb()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Address{})

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
	r.Run(":" + port)
}
