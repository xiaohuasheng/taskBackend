package models

import . "db_utils"

type Task struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Type   int    `json:"type" form:"type"`
	Status int    `json:"status" form:"status"`
}

func (p *Task) AddTask() (id int64, err error) {
	rs, err := GetDB().Exec("INSERT INTO person(name , type) VALUES (?, ?)", p.Name, p.Type)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Task) GetTasks() (persons []Task, err error) {
	persons = make([]Task, 0)
	rows, err := GetDB().Query("SELECT id,name,type,status FROM task")

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Name, &task.Type, &task.Status)
		persons = append(persons, task)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
