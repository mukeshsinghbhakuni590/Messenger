package dao

import (
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type View_update_struct struct {
	View_id string
	Session_id string
	Msg_id string
	Cont_id string
}


type View_post_struct struct {
	Msg_ids []string
	Session_id string
  Self_cont_id string
  Other_cont_id string
}

type View_info struct {
	View_id bson.ObjectId        `bson:"_id" json:"_id" ,omitempty`
	Msg_ids []bson.ObjectId			  `json:"msgs"`
}

func Add_single_view_id(sid string,oid string,vid string) {
  session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("myMessenger").C("contacts")
  self_info := Get_contact_by_id(sid)
  other_info := Get_contact_by_id(oid)
  self_sview := Single_veiw_struct{
     Cont_id : bson.ObjectIdHex(oid),
     View_id : bson.ObjectIdHex(vid),
  }
  other_sview := Single_veiw_struct{
     Cont_id : bson.ObjectIdHex(sid),
     View_id : bson.ObjectIdHex(vid),
  }
  self_info.Sviews = append(self_info.Sviews,self_sview)
  c.Update(bson.M{"_id": bson.ObjectIdHex(sid)}, bson.M{"$set": bson.M{"sviews":self_info.Sviews}})
  other_info.Sviews = append(other_info.Sviews,other_sview)
  c.Update(bson.M{"_id": bson.ObjectIdHex(oid)}, bson.M{"$set": bson.M{"sviews":other_info.Sviews}})
}

func Get_view_by_id(vid string) View_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("views")
	result := View_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(vid)}).One(&result)
	return result
}


func Create_View(t View_post_struct) View_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("views")

	view_id := bson.NewObjectId()
  var msgids []bson.ObjectId
  for _,v := range t.Msg_ids {
    msgids = append(msgids,bson.ObjectIdHex(v))
  }

	view_post_var := View_info{
		View_id : view_id,
		Msg_ids : msgids,
	}

	err_ := c.Insert(view_post_var)
	if err_ != nil {
                panic(err_)
        }
	return view_post_var
}

func Update_view(t View_update_struct) View_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("views")
	result := View_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.View_id)}).One(&result)
	for i, other := range result.Msg_ids {
		if other == bson.ObjectIdHex(t.Msg_id) {
			result.Msg_ids = append(result.Msg_ids[:i], result.Msg_ids[i+1:]...)
			break
		}
	}
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.View_id)}, bson.M{"$set": bson.M{"msgs":result.Msg_ids}})
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.View_id)}).One(&result)
	return result
}
