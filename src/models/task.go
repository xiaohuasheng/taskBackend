package models

import (
	. "db_utils"
	"log"
	"time"
)

type Task struct {
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Type       int    `json:"type" form:"type"`
	Status     int    `json:"status" form:"status"`
	CreateTime int    `json:"create_time" form:"create_time"`
	UpdateTime int    `json:"update_time" form:"update_time"`
}

func (p *Task) AddTask() (id int64, err error) {
	timeUnix := time.Now().Unix()
	rs, err := GetDB().Exec("INSERT INTO task(name , type, create_time) VALUES (?, ?, ?)", p.Name, p.Type, timeUnix)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Task) GetTasks() (persons []Task, err error) {
	persons = make([]Task, 0)
	rows, err := GetDB().Query("SELECT id,name,type,status,create_time,update_time FROM task where status=0")

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Name, &task.Type, &task.Status, &task.CreateTime, &task.UpdateTime)
		persons = append(persons, task)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func GetTask(id int) (person Task, err error) {
	err = GetDB().QueryRow("SELECT id, name, type, status, create_time, update_time FROM task WHERE id=?", id).Scan(
		&person.Id, &person.Name, &person.Type, &person.Status, &person.CreateTime, &person.UpdateTime)
	return
}

func ModTask(person Task) int64 {
	stmt, err := GetDB().Prepare("UPDATE task SET name=?, status=?, type=?, update_time=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	timeUnix := time.Now().Unix()
	rs, err := stmt.Exec(person.Name, person.Status, person.Type, timeUnix, person.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	return ra
}

func DelTask(id int) int64 {
	stmt, err := GetDB().Prepare("UPDATE task SET status=2, update_time=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	timeUnix := time.Now().Unix()
	rs, err := stmt.Exec(timeUnix, id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	return ra
}
