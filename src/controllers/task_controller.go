package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	. "models"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddTaskApi(c *gin.Context) {
	name := c.Request.FormValue("name")
	taskType, _ := strconv.Atoi(c.Request.FormValue("type"))

	p := Task{Name: name, Type: taskType}

	ra, err := p.AddTask()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetTasksApi(c *gin.Context) {
	task := Task{}
	persons := make([]Task, 0)
	persons, _ = task.GetTasks()
	c.JSON(http.StatusOK, gin.H{
		"msg": persons,
	})
}
