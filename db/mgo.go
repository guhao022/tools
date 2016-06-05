package db

import (
	"gopkg.in/mgo.v2"
	"sync"
)

type Mongo struct {
	Session *mgo.Session
}

var (
	database	*mgo.Database
	once		sync.Once
)

func NewMgo(database,username,password string, addr ...string) *Mongo {
	m := new(Mongo)
	once.Do(func() { // 只执行一次

		info := &mgo.DialInfo{
			Database: database,
			Username: username,
			Password: password,
			Addrs: addr,
		}

		session, err := mgo.DialWithInfo(info)
		if err != nil {
			panic(err)
		}

		m.Session = session.Clone()
	})


	return m

}

func (m *Mongo) C(name string) (*mgo.Session, *mgo.Collection) {
	Collection := database.With(m.Session).C(name)

	return m.Session, Collection
}

func (m *Mongo) Close() {
	m.Session.Close()
}



