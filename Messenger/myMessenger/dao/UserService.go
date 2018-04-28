package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Single_veiw_struct struct {
	Cont_id bson.ObjectId
	View_id bson.ObjectId
}

type Group_view_struct struct {
	Group_id bson.ObjectId
	View_id  bson.ObjectId
}

type User_post_struct struct {
	Uname  string
	Email  string
	Passwd string
}

type User_info struct {
	User_id     bson.ObjectId        `bson:"_id" json:"_id" ,omitempty`
	User_name   string               `json:"user_name"`
	Email       string               `json:"email"`
	Passwd      string               `json:"passwd"`
	User_status string               `json:"user_status"`
	Cont_id     bson.ObjectId        `json:"cont_id"`
}

type User_contect struct {
	Cont_id    bson.ObjectId   `bson:"_id" json:"_id" ,omitempty`
  	User_id     bson.ObjectId        `json:"user_id"`
	Cont_list  []bson.ObjectId `json:"cont_list"`
	Block_list []bson.ObjectId `json:"block_list"`
	Group_list []bson.ObjectId `json:"group_list"`
	Sviews      []Single_veiw_struct `json:"sviews"`
	Gviews      []Group_view_struct  `json:"gviews"`
}

type User_change_name struct {
	User_id       string
	New_User_name string
	Session_id    string
}

type User_change_status struct {
	User_id         string
	New_User_status string
	Session_id      string
}

type User_change_passwd struct {
	User_id    string
	Old_passwd string
	New_passwd string
	Session_id string
}

type Add_Block_contact struct {
	Self_cont_id   string
	Member_cont_id string
	Session_id     string
}

type Leave_group_struct struct {
	Cont_id    string
	Group_id   string
	Session_id string
}

type Get_cont_info struct {
	Session_id string
	Cont_id    string
}

type Other_user_get_struct struct {
	User_id string
	Session_id string
	Other_cont_id string
}

type Other_user_info struct {
	Cont_id     bson.ObjectId
	User_name   string
	User_status string
}

func Get_other_user_info(t Other_user_get_struct) Other_user_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")
	result := User_info{}
	c.Find(bson.M{"cont_id": t.Other_cont_id}).One(&result)
  other_info_var := Other_user_info{
		Cont_id : result.Cont_id,
		User_name : result.User_name,
		User_status : result.User_status,
	}
	return other_info_var
}

func Get_user_by_email(email string) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")

	result := User_info{}
	c.Find(bson.M{"email": email}).One(&result)
	return result
}

func Get_contact_by_id(cid string) User_contect {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("contacts")

	result := User_contect{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(cid)}).One(&result)
	return result
}

func Get_user_by_id(id string) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")

	fmt.Println(bson.ObjectIdHex(id))
	fmt.Println(bson.ObjectIdHex(id))
	result := User_info{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	fmt.Println(result)
	return result
}



func Get_all_users() []User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")

	result := []User_info{}
	c.Find(nil).All(&result)
	return result
}
func Create_user(u User_post_struct) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	user_contect_c := session.DB("myMessenger").C("contacts")
	user_info_c := session.DB("myMessenger").C("users")

	cont_id := bson.NewObjectId()
	user_id := bson.NewObjectId()
	contects_post_var := User_contect{
		Cont_id:    cont_id,
		User_id:    user_id,
		Cont_list:  nil,
		Block_list: nil,
		Group_list: nil,
		Sviews:      nil,
		Gviews:      nil,
	}

	err_user_cont := user_contect_c.Insert(contects_post_var)
	if err_user_cont != nil {
		panic(err_user_cont)
	}


	user_post_var := User_info{
		User_id:     user_id,
		User_name:   u.Uname,
		Email:       u.Email,
		Passwd:      u.Passwd,
		User_status: "",
		Cont_id:     cont_id,
	}

	err_user_info := user_info_c.Insert(user_post_var)
	if err_user_info != nil {
		panic(err_user_info)
	}
	return user_post_var
}

func Update_member_groupList_gviews(t Group_put_struct, ope string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	user_contect_c := session.DB("myMessenger").C("contacts")

	result := User_contect{}
	user_contect_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Member_id)}).One(&result)
	if ope == "ADD" {
		admin_cont := User_contect{}
		user_contect_c.Find(bson.M{"_id": bson.ObjectIdHex(t.Admin_id)}).One(&admin_cont)
		result.Group_list = append(result.Group_list, bson.ObjectIdHex(t.Group_id))
		for _,v := range admin_cont.Gviews {
			if v.Group_id == bson.ObjectIdHex(t.Group_id) {
				result.Gviews = append(result.Gviews,v)
			  	break
			}
		}
	} else if ope == "REM" {
		for j,v := range result.Group_list {
			if v == bson.ObjectIdHex(t.Group_id){
				result.Group_list = append(result.Group_list[:j],result.Group_list[j+1:]...)
				for k,g := range result.Gviews{
					if v == g.Group_id {
						result.Gviews = append(result.Gviews[:k],result.Gviews[k+1:]...)
						break
					}
				}
				break
			}
		}
	}
	user_contect_c.Update(bson.M{"_id": bson.ObjectIdHex(t.Member_id)}, bson.M{"$set": bson.M{"group_list": result.Group_list , "gviews":result.Gviews}})
}

func Change_user_name(t User_change_name) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")
	result := User_info{}
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.User_id)}, bson.M{"$set": bson.M{"user_name": t.New_User_name}})
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.User_id)}).One(&result)
	return result
}

func Change_user_status(t User_change_status) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")
	result := User_info{}
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.User_id)}, bson.M{"$set": bson.M{"user_name": t.New_User_status}})
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.User_id)}).One(&result)
	return result
}

func Change_user_passwd(t User_change_passwd) User_info {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("users")
	result := User_info{}
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.User_id)}, bson.M{"$set": bson.M{"user_name": t.New_passwd}})
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.User_id)}).One(&result)
	return result
}

func Update_add_block_list(t Add_Block_contact, opt string) User_contect {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("contacts")
	result := User_contect{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.Self_cont_id)}).One(&result)
	if opt == "ADD" {
		result.Cont_list = append(result.Cont_list, bson.ObjectIdHex(t.Member_cont_id))
		c.Update(bson.M{"_id": bson.ObjectIdHex(t.Self_cont_id)}, bson.M{"$set": bson.M{"cont_list": result.Cont_list}})
	} else if opt == "BLOCK" {
		result.Block_list = append(result.Block_list, bson.ObjectIdHex(t.Member_cont_id))
		c.Update(bson.M{"_id": bson.ObjectIdHex(t.Self_cont_id)}, bson.M{"$set": bson.M{"block_list": result.Block_list}})
	}
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.Self_cont_id)}).One(&result)
	return result
}

func Leave_group(t Leave_group_struct) User_contect {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("myMessenger").C("contacts")
	result := User_contect{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.Cont_id)}).One(&result)
	for i, other := range result.Group_list {
		if other == bson.ObjectIdHex(t.Group_id) {
			result.Group_list = append(result.Group_list[:i], result.Group_list[i+1:]...)
			for j,v := range result.Gviews{
				if v.Group_id == other{
					result.Gviews = append(result.Gviews[:j], result.Gviews[j+1:]...)
				}
			}
			break
		}
	}
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.Cont_id)}, bson.M{"$set": bson.M{"group_list": result.Group_list}})
	c.Update(bson.M{"_id": bson.ObjectIdHex(t.Cont_id)}, bson.M{"$set": bson.M{"gviews": result.Gviews}})
	c.Find(bson.M{"_id": bson.ObjectIdHex(t.Cont_id)}).One(&result)
	return result
}
