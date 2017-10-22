package models

import (
	"fmt"
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
		GetTaskByUserIdAndTaskId(bson.ObjectId, bson.ObjectId) (Task, error)
		CreateTask(Task) (Task, error)
		DeleteTaskByUserIdAndTaskId(bson.ObjectId, bson.ObjectId) (Task, error)
		UpdateTaskByUserIdAndTaskID(Task, bson.ObjectId, bson.ObjectId) (Task, error)
	}
)

/* task specific stuff */

// Returns all the tasks, a DataStore knows about
func (d *DataStore) GetAllTasksByUserId(userId bson.ObjectId) ([]Task, error) {
	tasks := []Task{}
	err := d.C(cNameTasks).Find(bson.M{"userid": userId}).Limit(100).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (d *DataStore) GetTaskByUserIdAndTaskId(userId bson.ObjectId, taskId bson.ObjectId) (Task, error) {
	task := Task{}
	err := d.C(cNameTasks).Find(bson.M{"userid": userId, "_id": taskId}).One(&task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (d *DataStore) CreateTask(task Task) (Task, error) {
	task.CreatedOn = time.Now()
	task.Id = bson.NewObjectId()
	err := d.C(cNameTasks).Insert(task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (d *DataStore) DeleteTaskByUserIdAndTaskId(userId bson.ObjectId, taskId bson.ObjectId) (Task, error) {
	task, err := d.GetTaskByUserIdAndTaskId(userId, taskId)
	if err != nil {
		return Task{}, err
	}
	err = d.C(cNameTasks).Remove(bson.M{"userid": userId, "_id": taskId})
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (d *DataStore) UpdateTaskByUserIdAndTaskID(newTask Task, userId bson.ObjectId, taskId bson.ObjectId) (Task, error) {

	fmt.Println(newTask)

	err := d.C(cNameTasks).Update(bson.M{"userid": userId, "_id": taskId},
		bson.M{"$set": bson.M{
			"name":        newTask.Name,
			"description": newTask.Description,
			"due":         newTask.Due,
			"status":      newTask.Status,
			"tags":        newTask.Tags,
		}})
	if err != nil {
		return Task{}, err
	}
	newTask.UserId = userId
	newTask.Id = taskId
	return newTask, nil
}
