package utils

import (
	"gopkg.in/mgo.v2"
	"time"
)

type Mongodb struct {
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
	query      *mgo.Query
}

//New Mongodb client
func NewMongodb(host string, port string, source string, username string, password string) *Mongodb {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{host + ":" + port},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "",
		Source:    source,
		Username:  username,
		Password:  password,
		PoolLimit: 4096,
	}
	c, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	m := &Mongodb{session: c}
	return m
}

//Select database
func (m *Mongodb) DB(database string) *Mongodb {
	m.database = m.session.DB(database)
	return m
}

//Select collection
func (m *Mongodb) Collection(collection string) *Mongodb {
	m.collection = m.database.C(collection)
	return m
}

//Get database all collections
func (m *Mongodb) AllCollections() (names []string) {
	names, _ = m.database.CollectionNames()
	return names
}

//Where condition
func (m *Mongodb) Where(condition interface{}) *Mongodb {
	m.query = m.collection.Find(condition)
	return m
}

//Limit condition
func (m *Mongodb) Limit(size int) *Mongodb {
	m.query = m.query.Limit(size)
	return m
}

//Execute query and get all result
func (m *Mongodb) All(result interface{}) {
	m.query.All(result)
}

//Execute query and get one result
func (m *Mongodb) One(result interface{}) {
	m.query.One(result)
}

//Close Mongodb
func (m *Mongodb) Close() {
	m.session.Close()
}
