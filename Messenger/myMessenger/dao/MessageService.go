package dao

import (
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Msg_put_sturct struct{
  Msg_id string
  Receiver_id string
  Status int
  Session_id string
}

type Receivers_struct struct{
	Receiver bson.ObjectId
	Status int
}

type Msg_post_struct struct {
	Sender string
	Receivers []string
	Msg_type string
	Msg_value string
	Views []string
	Session_id string
}

type Msg_info struct {
	Msg_id bson.ObjectId        `bson:"_id" json:"_id" ,omitempty`
	Sender bson.ObjectId       `json:"sender"`
	Receivers []Receivers_struct          `json:"receivers"`
	Msg_type string          `json:"msg_type"`
	Msg_value string          `json:"msg_value"`
	Views []bson.ObjectId			  `json:"views"`
}

func Get_msg_by_id(mid string) Msg_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("msgs")

	result := Msg_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(mid)}).One(&result)
	return result
}

func Create_msg(t Msg_post_struct) Msg_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("msgs")
	v := session.DB("myMessenger").C("views")
	msg_id := bson.NewObjectId()

	receivers := []Receivers_struct{}
	receiver:= Receivers_struct{}

	for _,r := range t.Receivers{
		receiver.Receiver = bson.ObjectIdHex(r)
		receiver.Status = 0
		receivers = append(receivers,receiver)
	}
	var views []bson.ObjectId
	result := View_info{}
  for _,o := range t.Views{
		views = append(views,bson.ObjectIdHex(o))
		v.Find(bson.M{"_id": bson.ObjectIdHex(o)}).One(&result)
		result.Msg_ids = append(result.Msg_ids,msg_id)
		v.Update(bson.M{"_id": bson.ObjectIdHex(o)}, bson.M{"$set": bson.M{"msgs":result.Msg_ids}})

  }

	msg_post_var := Msg_info{
		Msg_id : msg_id,
		Sender : bson.ObjectIdHex(t.Sender),
		Receivers : receivers,
		Msg_type : t.Msg_type,
		Msg_value : t.Msg_value,
		Views : views,
	}

	err_ := c.Insert(msg_post_var)
	if err_ != nil {
                panic(err_)
        }
	return msg_post_var
}

func Update_receiver_status(t Msg_put_sturct) Msg_info {
  session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("msgs")
  result := Msg_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.Msg_id)}).One(&result)
  for _, other := range result.Receivers {
    if other.Receiver == bson.ObjectIdHex(t.Receiver_id) {
      other.Status = t.Status
      //result.Receivers = append(result.Receivers[:i], result.Receivers[i+1:]...)
      break
    }
  }
  c.Update(bson.M{"_id": bson.ObjectIdHex(t.Msg_id)}, bson.M{"$set": bson.M{"receivers":result.Receivers}})
  //c.Find(bson.M{"_id": t.Msg_id}).One(&result)
	return result
}
