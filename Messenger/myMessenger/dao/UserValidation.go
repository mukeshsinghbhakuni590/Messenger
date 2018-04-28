package dao

import (
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Validation_post_struct struct {
	Email string
	Passwd string
}

type Validation_info struct {
	Session_id bson.ObjectId        `bson:"_id" json:"_id" ,omitempty`
	Email string           `json:"email"`
  	User_id bson.ObjectId   `json:"usrid"`
}

type Delete_session struct {
	Session_id string
	Email string 
  	User_id string 
}


func Create_Session(t Validation_post_struct) Validation_info {
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("user_sessions")
  user_info := Get_user_by_email(t.Email)
	session_id := bson.NewObjectId()
	validation_post_var := Validation_info{
		Session_id : session_id,
		Email : t.Email,
    User_id : user_info.User_id,
	}

	err_ := c.Insert(validation_post_var)
	if err_ != nil {
                panic(err_)
        }
	return validation_post_var
}

func Get_session_by_email(email string) Validation_info {
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("user_sessions")

	result := Validation_info{}
	c.Find(bson.M{"email": email}).One(&result)
	return result
}

func Get_session_by_id(id string) Validation_info {
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("user_sessions")

	result := Validation_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return result
}

func Delete_Session(t Delete_session){
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("user_sessions")
	c.Remove(bson.M{"_id": bson.ObjectIdHex(t.Session_id)})
}
