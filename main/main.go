package main

import (
	"fmt"
	"learning-go/task_manager"
	"time"
)

func main() {

	store := task_manager.NewTaskStore()
	manager := task_manager.NewTaskManager(store)

	task := task_manager.Task{
		Id:           1,
		Title:        "Срочная задача",
		Priority:     task_manager.High,
		Description:  "Надо сделать эту задачу как можно скорее",
		CreationDate: time.Now(),
		Author:       "Артемова Анастасия",
	}

	manager.AddTask(&task)
	tasks := manager.GetAllData()

	for _, v := range tasks {
		fmt.Println(v)
	}
}
