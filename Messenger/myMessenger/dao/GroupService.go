package dao

import (
  "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Group_post_struct struct {
	Gname string
	Admins []string
	Gmembers []string
	Session_id string
}

type Group_info struct {
	Group_id bson.ObjectId   `bson:"_id" json:"_id" ,omitempty`
	Group_name string       `json:"group_name"`
	Group_admins []bson.ObjectId         `json:"group_admins"`
	Group_members []bson.ObjectId         `json:"group_members"`
}

type Group_put_struct struct {
	Group_id string
	Admin_id string
	Member_id string
	Session_id string
}

type Group_rename struct {
	Group_id string
	Member_id string
	New_group_name string
	Session_id string
}

func Make_admin(t Group_put_struct) Group_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")

	result := Group_info{}
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	result.Group_admins = append(result.Group_admins , bson.ObjectIdHex(t.Member_id))
	group_info_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}, bson.M{"$set": bson.M{"group_admins":result.Group_admins}})
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	return result
}

func Add_member(t Group_put_struct) Group_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")

	result := Group_info{}
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	result.Group_members = append(result.Group_members , bson.ObjectIdHex(t.Member_id))
	group_info_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}, bson.M{"$set": bson.M{"group_members":result.Group_members}})
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	return result
}

func Remove_member(t Group_put_struct) Group_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    group_info_c := session.DB("myMessenger").C("groups")
    
    result := Group_info{}
    group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
    for i, other := range result.Group_members {
	    if other == bson.ObjectIdHex(t.Member_id) {
		    result.Group_members = append(result.Group_members[:i], result.Group_members[i+1:]...)
		    break
		}
	}
	for i, other := range result.Group_admins {
        if other == bson.ObjectIdHex(t.Member_id) {
			result.Group_admins = append(result.Group_admins[:i], result.Group_admins[i+1:]...)
			break
        }
    }
    group_info_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}, bson.M{"$set": bson.M{"group_members":result.Group_members,"group_admins":result.Group_admins}})
    group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
    return result
}

func Rename_group(t Group_rename) Group_info {
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")

	result := Group_info{}
	group_info_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}, bson.M{"$set": bson.M{"group_name":t.New_group_name}})
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	return result
}


func Get_group_by_id(gid string) Group_info {
    session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")

	result := Group_info{}
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(gid)}).One(&result)
	return result
}


func Create_group(g Group_post_struct) Group_info {
    session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")
  view_post_var := View_post_struct{
		Msg_ids : nil,
    Session_id : g.Session_id,
	}
  view_info := Create_View(view_post_var)
	group_id := bson.NewObjectId()
  var gadmins []bson.ObjectId
  var gmembers []bson.ObjectId
  grp_view_var := Group_view_struct{
    Group_id : group_id,
    View_id  : view_info.View_id,
  }
  for _,v := range g.Admins {
    gadmins = append(gadmins,bson.ObjectIdHex(v))
  }
  for _,v := range g.Gmembers {
    gmembers = append(gmembers,bson.ObjectIdHex(v))
    cont_info := Get_contact_by_id(v)
    cont_info.Gviews = append(cont_info.Gviews,grp_view_var)
  }

	group_post_var := Group_info{
		Group_id : group_id,
		Group_name : g.Gname,
		Group_admins : gadmins,
		Group_members : gmembers,
	}

	err_group_info := group_info_c.Insert(group_post_var)
	if err_group_info != nil {
                panic(err_group_info)
        }
	return group_post_var
}

func Remove_self_from_group(t Leave_group_struct){
	session, err := mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	group_info_c := session.DB("myMessenger").C("groups")
	result := Group_info{}
	group_info_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}).One(&result)
	for i, other := range result.Group_members {
        if other == bson.ObjectIdHex(t.Cont_id) {
			result.Group_members = append(result.Group_members[:i], result.Group_members[i+1:]...)
			break
        }
	}
	for i, other := range result.Group_admins {
        if other == bson.ObjectIdHex(t.Cont_id) {
			result.Group_admins = append(result.Group_admins[:i], result.Group_admins[i+1:]...)
			break
        }
	}
	group_info_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Group_id)}, bson.M{"$set": bson.M{"group_admins":result.Group_admins,"group_members":result.Group_members}})
}
