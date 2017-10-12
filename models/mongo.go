package models

import (
	"log"

	"github.com/nathanmalishev/taskmanager/common"
	mgo "gopkg.in/mgo.v2"
)

type DataStore struct {
	Session *mgo.Session
}

/* getters / setters */
func CreateStore(dialInfo *mgo.DialInfo) *DataStore {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &DataStore{session}
}

func (d *DataStore) GetStore() *DataStore {
	if d.Session == nil {
		log.Fatal("There is no session active")
	}
	// important, open new session for concurreny
	return &DataStore{d.Session.Copy()}

}

func (d *DataStore) Close() {
	if d.Session == nil {
		log.Fatal("You cannot close an empty session")
	}
	d.Session.Close()
}

/* general database functions */
func (d *DataStore) C(collection string) *mgo.Collection {
	return d.Session.DB(common.AppConfig.DbName).C(collection)
}
