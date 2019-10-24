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
	var task Task
	persons := make([]Task, 0)
	persons, _ = task.GetTasks()
	c.JSON(http.StatusOK, persons)
}

func GetTaskApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task Task
	var err error
	task, err = GetTask(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, task)
}

func ModTaskApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	task := Task{Id: id}
	err = c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}
	ra := ModTask(task)
	msg := fmt.Sprintf("Update task %d successful %d", task.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelTaskApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	ra := DelTask(id)
	msg := fmt.Sprintf("Delete task %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
