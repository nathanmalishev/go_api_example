package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const cNameTasks = "tasks"

type (
	Task struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	TaskStore interface {
		GetAllTasks() ([]Task, error)
	}
)

/* task specific stuff */

// Returns all the tasks, a DataStore knows about
func (d *DataStore) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	err := d.C(cNameTasks).Find(nil).Limit(100).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (d *DataStore) InsertTask(t Task) error {
	return d.C(cNameTasks).Insert(t)
}
