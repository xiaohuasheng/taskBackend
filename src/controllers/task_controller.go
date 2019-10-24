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
	if name == "" || taskType == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "name或type为空",
		})
		return
	}

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
	taskList := make([]Task, 0)
	taskList, _ = task.GetTasks()

	typeTaskMap := make(map[int][]Task)
	for _, task := range taskList {
		if _, ok := typeTaskMap[task.Type]; ok {
			//存在
			typeTaskMap[task.Type] = append(typeTaskMap[task.Type], task)
		} else {
			typeTaskMap[task.Type] = []Task{task}
		}
	}
	c.JSON(http.StatusOK, typeTaskMap)
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
