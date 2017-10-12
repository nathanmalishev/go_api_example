package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const collectionName = "tasks"

// move each model into its respective file
// write an interface for each function we need
// Datastore will implement each of these functions

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
	err := d.C(collectionName).Find(nil).Limit(100).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (d *DataStore) InsertTask(t Task) error {
	return d.C(collectionName).Insert(t)
}
