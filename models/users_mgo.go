package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"

	"github.com/globalsign/mgo"
)

type userServiceMGO struct {
	s  *mgo.Session
	db *mgo.Database
	c  *mgo.Collection
}

type userMGO struct {
	User `bson:",inline"`
	ID   bson.ObjectId `bson:"_id"`
}

// NewUserServiceMGO creates a UserService which accesses users
// in a MongoDB database using the GlobalSign mgo driver.
func NewUserServiceMGO(host, db, username, password string) (UserService, error) {
	usm := &userServiceMGO{}
	var err error
	usm.s, err = mgo.Dial("mongodb://" + host)
	if err != nil {
		return nil, err
	}

	err = usm.s.Login(&mgo.Credential{
		Username: username,
		Password: password,
		Source:   db,
	})
	if err != nil {
		return nil, err
	}

	usm.db = usm.s.DB(db)
	usm.c = usm.db.C("users")
	return &userService{UserDB: usm}, nil
}

func (us *userServiceMGO) ByUUID(uuid uuid.UUID) (*User, error) {
	u := &User{}
	err := us.c.Find(bson.M{"uuid": uuid}).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *userServiceMGO) ByEmail(email string) (*User, error) {
	u := &User{}
	err := us.c.Find(bson.M{"email": email}).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *userServiceMGO) Create(user *User) error {
	ctime := time.Now()
	user.UUID = uuid.New()
	user.CreatedAt = ctime
	user.ModifiedAt = ctime
	user.DeletedAt = time.Time{}
	// The _id field will be automatically added with a valid ObjectID.
	return us.c.Insert(user)
}

func (us *userServiceMGO) Update(user *User) error {
	// Only update the mutable user data.
	return us.c.Update(
		bson.M{"uuid": user.UUID},
		bson.M{"$set": &user.UserMutableData, "$currentDate": bson.M{"modifiedat": true}})
}

func (us *userServiceMGO) Delete(uuid uuid.UUID) error {
	return us.c.Remove(bson.M{"uuid": uuid})
}

func (us *userServiceMGO) Close() error {
	us.s.Close()
	return nil
}
