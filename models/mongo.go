package models

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type (
	DataStorer interface {
		GetStore() *DataStore
		Close()
		InitIndexs() error
		UserStore
		TaskStore
	}
	// Implementation of the DataStorer
	DataStore struct {
		Session *mgo.Session
		DbName  string
	}
)

/* getters / setters */
func CreateStore(dialInfo *mgo.DialInfo, dbName string) *DataStore {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &DataStore{session, dbName}
}

// Returns a new data store with a copy of the mgo session
func (d *DataStore) GetStore() *DataStore {
	if d.Session == nil {
		log.Fatal("There is no session active")
	}
	// important, open new session for concurreny
	return &DataStore{d.Session.Copy(), d.DbName}
}

// Closes the mgo session, within the DataStore
func (d *DataStore) Close() {
	if d.Session == nil {
		log.Fatal("You cannot close an empty session")
	}
	d.Session.Close()
}

/* indexs */
func (d *DataStore) InitIndexs() error {
	return d.UserIndexs()
}

/* general database functions */
func (d *DataStore) C(collection string) *mgo.Collection {
	return d.Session.DB(d.DbName).C(collection)
}
