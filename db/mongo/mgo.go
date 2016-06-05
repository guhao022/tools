package mongo

import (
	"gopkg.in/mgo.v2"
	"sync"
	"fmt"
)

var (
	session *mgo.Session
	once sync.Once
)

func Connect(host string) error {
	var err error
	s, err := mgo.Dial(host)
	if err != nil {
		return err
	}

	session = s

	return nil
}

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func getSession() *mgo.Session {
	//最大连接池默认为4096
	return session.Clone()
}

func DB(database string) (*mgo.Session, *mgo.Database) {
	session := getSession()

	db := session.DB(database)
	if db == nil {
		session.Close()
		fmt.Println("数据库连接失败!")
	}

	return session, db
}

func AuthDB(database, username, password string) (*mgo.Session, *mgo.Database) {

	session := getSession()

	db := session.DB(database)
	if db == nil {
		session.Close()
		fmt.Println("数据库连接失败!")
	}
	err := db.Login(username, password)

	if err != nil {
		fmt.Printf("mongodb登陆失败: %s\n", err)
	}

	return session, db
}
