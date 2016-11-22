package mongo

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"github.com/nairufan/yh-share/util"
)

const (
	DB_Name = "youhui"
)

type Option struct {
	Sort   []string
	Limit  *int
	Offset *int
	Select bson.M
}

var globalSession *mgo.Session
var warningMongoQueryDuration = time.Millisecond * 500

func init() {
	mongodbUrl := beego.AppConfig.String("mongodb_url")
	session, err := mgo.Dial(mongodbUrl)
	if err != nil {
		util.Panic(err)
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
		util.Panic(err)
	}
}


func (s *Session) Find(collection string, query interface{}, result interface{}) error {
	return s.C(collection).Find(query).All(result)
}

func (s *Session) MustFind(collection string, query interface{}, result interface{}) {
	if err := s.Find(collection, query, result); err != nil {
		util.Panic(err)
	}
}

func (s *Session) FindId(collectionName string, id interface{}, result interface{}) error {
	return s.C(collectionName).FindId(id).One(result)
}

func (s *Session) MustFindId(collection string, id interface{}, result interface{}) {
	if err := s.FindId(collection, id, result); err != nil {
		util.Panic(err)
	}
}

func (s *Session) FindWithOptions(collection string, query interface{}, options Option, result interface{}) error {
	q := s.C(collection).Find(query)
	if len(options.Sort) > 0 {
		q = q.Sort(options.Sort...)
	}
	if options.Offset != nil {
		q = q.Skip(*options.Offset)
	}
	if options.Limit != nil {
		q = q.Limit(*options.Limit)
	}
	if len(options.Select) != 0 {
		q = q.Select(options.Select)
	}
	return q.All(result)
}

func (s *Session) MustFindWithOptions(collection string, query interface{}, options Option, result interface{}) {
	if err := s.FindWithOptions(collection, query, options, result); err != nil {
		util.Panic(err)
	}
}