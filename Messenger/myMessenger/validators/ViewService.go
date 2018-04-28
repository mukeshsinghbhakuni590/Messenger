package validators

import (
	"myMessenger/dao"
	"gopkg.in/mgo.v2/bson"
)


func Validate_for_delete_msg_from_view(t dao.View_update_struct) bool {
	msg := dao.Get_msg_by_id(t.Msg_id)
	if msg.Sender == bson.ObjectIdHex(t.Cont_id) {
		return true
	}else{
		return false
	}
}
