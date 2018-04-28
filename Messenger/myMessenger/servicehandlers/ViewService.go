package servicehandlers

import (
	"myMessenger/dao"
	"encoding/json"
	"net/http"
	"reflect"
	"myMessenger/validators"
)

type ViewHandler struct {
}

func (p ViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ViewHandler) Get(req *http.Request) (string,int) {
	viewid,ok1 := req.URL.Query()["viewid"]
	sessionid,ok2 := req.URL.Query()["sessionid"]
	if !ok1 || !ok2 {
		return (" Somthing Wrong !!"),404
	}else if viewid[0] != "" && sessionid[0] != "" {
		session_info := dao.Get_session_by_id(sessionid[0])
		if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
			return (" You have to login Again "),401
		}else{
			view,_ := json.Marshal(dao.Get_view_by_id(viewid[0]))
			return string(view),200
		}
	}else {
		return ("Parameters Not correctly defined !!"),400
	}
}

func (p ViewHandler) Put(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.View_update_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	session_info := dao.Get_session_by_id(t.Session_id)
	if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
		return (" You have to login Again "),401
	} else {
		if validators.Validate_for_delete_msg_from_view(t){
			view,_ := json.Marshal(dao.Update_view(t))
			return (string(view)),200
		}else {
			return ("Error : You cant delete msg"),401
		}
	}
}

func (p ViewHandler) Post(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.View_post_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	session_info := dao.Get_session_by_id(t.Session_id)
	if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
		return (" You have to login Again "),401
	} else {
		view_info := dao.Create_View(t)
		if validators.Validate_for_contacts(t.Self_cont_id,t.Other_cont_id) {
	  			dao.Add_single_view_id(t.Self_cont_id,t.Other_cont_id,string(view_info.View_id))
	  }
		view,_ := json.Marshal(view_info)
		return (string(view)),200
	}
}
