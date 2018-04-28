package validators

import (
	"myMessenger/dao"
	"gopkg.in/mgo.v2/bson"
)

func Ispresent(list []bson.ObjectId , ele string) bool {
	for _,v := range list {
		if v == bson.ObjectIdHex(ele) {
			return true
		}
	}
	return false
}

func Validate_for_makeAdmin(t dao.Group_put_struct) bool {
	grp_info := dao.Get_group_by_id(t.Group_id)
	if Ispresent(grp_info.Group_admins,t.Admin_id) && Ispresent(grp_info.Group_members,t.Member_id) && !Ispresent(grp_info.Group_admins,t.Member_id){
		return true
	}else{
		return false
	}
}

func Validate_for_addMember(t dao.Group_put_struct) bool {
	grp_info := dao.Get_group_by_id(t.Group_id)
	if Ispresent(grp_info.Group_admins,t.Admin_id) && !Ispresent(grp_info.Group_members,t.Member_id){
		return true
	}else{
		return false
	}
}

func Validate_for_remMember(t dao.Group_put_struct) bool {
	grp_info := dao.Get_group_by_id(t.Group_id)
	if Ispresent(grp_info.Group_admins,t.Admin_id) && Ispresent(grp_info.Group_members,t.Member_id){
		if (len(grp_info.Group_admins)) > 1{
			return true
		}else if t.Admin_id != t.Member_id {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}


func Validate_for_renameGroup(t dao.Group_rename) bool {
	grp_info := dao.Get_group_by_id(t.Group_id)
	if Ispresent(grp_info.Group_members,t.Member_id){
		return true
	}else{
		return false
	}
}
