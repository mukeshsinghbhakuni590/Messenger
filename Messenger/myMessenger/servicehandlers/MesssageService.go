package servicehandlers

import (
	"myMessenger/dao"
	"encoding/json"
	"net/http"
	"reflect"
)

type MsgHandler struct {
}

func (p MsgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p MsgHandler) Get(req *http.Request) (string,int) {
	msgid,ok1 := req.URL.Query()["msgid"]
	sessionid,ok2 := req.URL.Query()["sessionid"]
	if !ok1 || !ok2 {
		return ("Something Wrong !!"),404
	}else if msgid[0] != "" && sessionid[0] != "" {
		session_info := dao.Get_session_by_id(sessionid[0])
		if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
			return (" You have to login Again "),401
		}else{
			msg,_ := json.Marshal(dao.Get_msg_by_id(msgid[0]))
			return string(msg),200
		}
	}else {
		return ("Parameters Not correctly defined !!"),404
	}
}

func (p MsgHandler) Put(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.Msg_put_sturct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	session_info := dao.Get_session_by_id(t.Session_id)
	if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
		return (" You have to login Again "),401
	} else {
		view,_ := json.Marshal(dao.Update_receiver_status(t))
		return (string(view)),200
	}
}

func (p MsgHandler) Post(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.Msg_post_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	session_info := dao.Get_session_by_id(t.Session_id)
	if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
		return (" You have to login Again "),401
	} else {
		msg,_ := json.Marshal(dao.Create_msg(t))
		return (string(msg)),200
	}
}
