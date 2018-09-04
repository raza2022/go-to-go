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

func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a movie by its id
func (m *MoviesDAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
