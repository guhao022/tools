package db

import "gopkg.in/mgo.v2"

type Mongo struct {
	session *mgo.Session
	database *mgo.Database
}

func NewMgo(database,username,password string, addr ...string) *Mongo {
	m := new(Mongo)

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

	m.session = session.Clone()

	return m

}

func (m *Mongo) C(name string) *mgo.Collection {
	collection := m.database.With(m.session).C(name)

	return collection

}

func (m *Mongo) Close() {
	m.session.Close()
}
