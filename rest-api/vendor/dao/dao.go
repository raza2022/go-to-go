package dao

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "log"

type MoviesDAO struct {
	Server   string
	Database string
}

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Establish a connection to database
//func (m *MoviesDAO) Connect() {
//	session, err := mgo.Dial(m.Server)
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = session.DB(m.Database)
//}

func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

//// Find list of movies
//func (m *MoviesDAO) FindAll() ([]Movie, error) {
//	var movies []Movie
//	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
//	return movies, err
//}
