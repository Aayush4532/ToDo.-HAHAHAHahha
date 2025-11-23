package controllers

import "github.com/gin-gonic/gin"

func PostToDo(c *gin.Context, tasks *[]string, count *int) {
	type ToDo struct {
		Task string `json:"task"`
		Id   int    `json:"id"`
	}

	var newToDo ToDo

	if err := c.BindJSON(&newToDo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	*count++
	newToDo.Id = *count

	if newToDo.Task == "" {
		c.JSON(400, gin.H{"error": "Task cannot be empty"})
		return
	}

	*tasks = append(*tasks, newToDo.Task)
	c.JSON(200, gin.H{
		"message": "To-Do item created successfully",
		"task":    newToDo.Task,
	})
}

func GetToDos(c *gin.Context, tasks *[]string) {
	c.JSON(200, gin.H{
		"tasks": *tasks,
	})
}

func UpdateToDo(c *gin.Context, tasks *[]string) {
	type ToDo struct {
		Task string `json:"task"`
		Id   int    `json:"id"`
	}
	var updatedToDo ToDo

	if err := c.BindJSON(&updatedToDo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updatedToDo.Id <= 0 || updatedToDo.Id > len(*tasks) {
		c.JSON(404, gin.H{"error": "To-Do item not found"})
		return
	}

	(*tasks)[updatedToDo.Id-1] = updatedToDo.Task

	c.JSON(200, gin.H{
		"message": "To-Do item updated successfully",
		"task":    updatedToDo.Task,
	})
}

func DeleteToDo(c *gin.Context, tasks *[]string) {
	type ToDo struct {
		Id int `json:"id"`
	}
	var toDoToDelete ToDo

	if err := c.BindJSON(&toDoToDelete); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if toDoToDelete.Id <= 0 || toDoToDelete.Id > len(*tasks) {
		c.JSON(404, gin.H{"error": "To-Do item not found"})
		return
	}

	*tasks = append((*tasks)[:toDoToDelete.Id-1], (*tasks)[toDoToDelete.Id:]...)

	c.JSON(200, gin.H{
		"message": "To-Do item deleted successfully",
		"task":    (*tasks),
	})
}
