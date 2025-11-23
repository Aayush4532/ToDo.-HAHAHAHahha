package main

import (
	"os"
	routes "todo/Routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err);
	}
	r := gin.Default();
	// cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"content-type"},
	}))

	todoGroup := r.Group("api/");
	tasks := []string{}
	count := 0

	routes.ToDoRoutes(todoGroup, &tasks, &count);
	r.Run(":" + os.Getenv("PORT"));
}