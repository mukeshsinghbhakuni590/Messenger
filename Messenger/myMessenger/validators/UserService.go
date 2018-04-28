package validators

import (
	"myMessenger/dao"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)



func Validate_for_changename(t dao.User_change_name) bool {
	return true
}

func Validate_for_changestatus(t dao.User_change_status) bool {
	return true
}

func Validate_for_contacts(sid string,oid string) bool {
	 self_info := dao.Get_contact_by_id(sid)
	 other_info := dao.Get_contact_by_id(oid)
	 if (!reflect.DeepEqual(self_info, (dao.User_contect{})) && !reflect.DeepEqual(other_info, (dao.User_contect{}))){
		 return true
	 } else {
		 return false
	 }
}

func Validate_for_changepasswd(t dao.User_change_passwd) bool{
	user_info := dao.Get_user_by_id(t.User_id)
	if user_info.Passwd == t.Old_passwd {
		return true
	}else{
		return false
	}
}

func Validate_for_addcontact(t dao.Add_Block_contact) bool {
	cont_info := dao.Get_contact_by_id(t.Member_cont_id)
	if reflect.DeepEqual(cont_info, (dao.User_contect{})) {
		return false
	} else {
		return true
	}
}

func Validate_for_get_other_user(t dao.Other_user_get_struct) bool {
  user_info := dao.Get_user_by_id(t.User_id)
	if !(reflect.DeepEqual(user_info, (dao.User_info{}))) {
		cont_info := dao.Get_contact_by_id(string(user_info.Cont_id))
	  for _,v := range cont_info.Cont_list {
			  if v == bson.ObjectIdHex(t.Other_cont_id) {
				     return true
				}
		}
		return false
	} else {
	  return false
	}
}

func Validate_for_blockcontact(t dao.Add_Block_contact)bool {
	user_cont := dao.Get_contact_by_id(t.Self_cont_id)
	var a,b bool
	a = false
	b = false
	for _, other := range user_cont.Cont_list {
		if other == bson.ObjectIdHex(t.Member_cont_id) {
			a = true
			break
		}
	}
	for _, other := range user_cont.Block_list {
		if other == bson.ObjectIdHex(t.Member_cont_id) {
			b = true
			break
		}
	}
	if a && !b {
		return true
	}else{
		return false
	}
}

func Validate_for_leavegroup(t dao.Leave_group_struct) bool {
	user_cont := dao.Get_contact_by_id(t.Cont_id)
	group_info := dao.Get_group_by_id(t.Group_id)
	var a,b,c bool
	a = false
	b = false
	c = false
	for _, other := range user_cont.Group_list {
		if other == bson.ObjectIdHex(t.Group_id) {
			a = true
			break
		}
	}
	for _, other := range group_info.Group_members {
		if other == bson.ObjectIdHex(t.Cont_id) {
			b = true
			break
		}
	}
	for _, other := range group_info.Group_admins {
		if other == bson.ObjectIdHex(t.Cont_id) {
			c = true
			break
		}
	}
	if a && b && ((c && (len(group_info.Group_admins) > 1)) || !c) {
		return true
	}else{
		return false
	}
}
