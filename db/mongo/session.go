package mongo

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/astaxie/beego"
)

const (
	DB_Name = "youhui"
)

var globalSession *mgo.Session
var warningMongoQueryDuration = time.Millisecond * 500

func init() {
	mongodbUrl := beego.AppConfig.String("mongodb_url")
	session, err := mgo.Dial(mongodbUrl)
	if err != nil {
		panic(err)
	}
	globalSession = session
}

func NewSession() *mgo.Session {
	return globalSession.Copy()
}

type Session struct {
	*mgo.Session
}

func Get() *Session {
	return &Session{
		globalSession.Copy(),
	}
}

func (s *Session) C(name string) *mgo.Collection {
	return s.DB(DB_Name).C(name)
}

func (s *Session) Insert(collectionName string, docs ...interface{}) error {
	return s.C(collectionName).Insert(docs...)
}

func (s *Session) MustInsert(collectionName string, docs ...interface{}) {
	if err := s.Insert(collectionName, docs...); err != nil {
		panic(err)
	}
}


func (s *Session) Find(collection string, query interface{}, result interface{}) error {
	return s.C(collection).Find(query).All(result)
}

func (s *Session) MustFind(collection string, query interface{}, result interface{}) {
	if err := s.Find(collection, query, result); err != nil {
		panic(err)
	}
}