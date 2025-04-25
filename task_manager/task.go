package task_manager

import (
	"fmt"
	"time"
)

type Task struct {
	Id           int
	Title        string
	Priority     Priority
	Description  string
	CreationDate time.Time
	Author       string
}

type Priority string

const (
	Low    Priority = "Low"
	Medium Priority = "Medium"
	High   Priority = "High"
)

func (t *Task) String() string {
	return fmt.Sprintf(
		"Task: {\n  Id: %d, \n  Title: %s, \n  Priority: %s, \n  Description: %s, \n  CreationDate: %s, \n  Author: %s\n}",
		t.Id,
		t.Title,
		t.Priority,
		t.Description,
		t.CreationDate.Format("2006-01-02"),
		t.Author,
	)
}
