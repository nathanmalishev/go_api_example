package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const cNameTasks = "tasks"

type (
	Task struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId      bson.ObjectId `bson:"userid" json:"userid"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	TaskStore interface {
		GetAllTasksByUserId(bson.ObjectId) ([]Task, error)
	}
)

/* task specific stuff */

// Returns all the tasks, a DataStore knows about
func (d *DataStore) GetAllTasksByUserId(userId bson.ObjectId) ([]Task, error) {
	tasks := []Task{}
	err := d.C(cNameTasks).Find(bson.M{"UserId": userId}).Limit(100).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (d *DataStore) InsertTask(t Task) error {
	return d.C(cNameTasks).Insert(t)
}
