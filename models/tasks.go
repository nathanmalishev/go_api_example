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
		GetAllTasks(*DataStore) ([]Task, error)
	}
)

func (*DataStore) GetAllTasks() ([]Task, error) {
	return nil, nil
}

//[> tasks specific models <]

//func (d *DataStore) FindAllNotes() []models.Task {
//tasks := []models.Task{}
//itr := d.DB(common.Config.DbName).C(collectionName).Find(nil).Limit(100).Iter()
//err := itr.All(&tasks)
//if err != nil {
//log.Fatal(err)
//}
//return tasks
//}
