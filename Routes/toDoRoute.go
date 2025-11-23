package routes

import (
	controllers "todo/Controllers"
	"github.com/gin-gonic/gin"
)

func ToDoRoutes (r *gin.RouterGroup, tasks *[]string, count *int) {
	r.POST("/task", func(c *gin.Context) {
		controllers.PostToDo(c, tasks, count);
	})
	r.GET("/todos", func(c *gin.Context) {
		controllers.GetToDos(c, tasks);
	})
	r.PUT("/task", func(c *gin.Context) {
		controllers.UpdateToDo(c, tasks);
	})
	r.DELETE("/task", func(c *gin.Context) {
		controllers.DeleteToDo(c, tasks);
	})
}