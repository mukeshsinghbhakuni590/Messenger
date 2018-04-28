package servicehandlers

import (
	"myMessenger/dao"
	"myMessenger/validators"
	"encoding/json"
	"net/http"
	"reflect"
)

type GroupHandler struct {
}

func (p GroupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}


//------------------------------------------------------------------------
func (p GroupHandler) Get(req *http.Request) (string,int) {
	grpid,ok1 := req.URL.Query()["grpid"]
	session_id,ok2 := req.URL.Query()["sessionid"]
	if !ok1 || !ok2 {
		return ("Something Wrong !!"),404
	}else if grpid[0] != "" && session_id[0] != "" {
		session_info := dao.Get_session_by_id(session_id[0])
		if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
			return (" You have to login Again "),401
		}else {
			grp,_ := json.Marshal(dao.Get_group_by_id(grpid[0]))
			return (string(grp)),200
		}
	}else {
		return ("Group Id or Session Id is Empty !!"),404
	}
}

//--------------------------------------------------------------------------
func (p GroupHandler) Put(req *http.Request) (string,int) {
	operation,ok := req.URL.Query()["operation"]
	if !ok {
		return ("Something Wrong !!"),404
	}else if len(operation) == 1 {
		decoder := json.NewDecoder(req.Body)
		if operation[0] == "makeAdmin" {
			var t dao.Group_put_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
				return (" You have to login Again "),401
			}else {
				if validators.Validate_for_makeAdmin(t) {
					grp,_ := json.Marshal(dao.Make_admin(t))
					return (string(grp)),200
				}else{
					return ("Make Admin request denied !!"),400
				}
			}
		} else if operation[0] == "addMember" {
			var t dao.Group_put_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
				return (" You have to login Again "),401
			}else{
				if validators.Validate_for_addMember(t){
					grp,_ := json.Marshal(dao.Add_member(t))
					dao.Update_member_groupList_gviews(t,"ADD")
					return (string(grp)),200
				}else{
					return ("Add Member request denied !!"),400
				}
			}
		} else if operation[0] == "remMember" {
			var t dao.Group_put_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
				return (" You have to login Again "),401
			}else{
				if validators.Validate_for_remMember(t){
					grp,_ := json.Marshal(dao.Remove_member(t))
					dao.Update_member_groupList_gviews(t,"REM")
					return (string(grp)),200
				}else{
					return ("Remove Member request denied !!"),400
				}
			}
		} else if operation[0] == "renameGroup"{
			var t dao.Group_rename
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
				return (" You have to login Again "),401
			}else{
				if validators.Validate_for_renameGroup(t){
					grp,_ := json.Marshal(dao.Rename_group(t))
					return (string(grp)),200
				}else{
					return ("Rename Group request denied !!"),400
				}
			}
		} else {
			return ("Invalid Operation"),405
		}
	}else {
		return ("multiple operations  or none !!"),406
	}
}

//--------------------------------------------------------------------------
func (p GroupHandler) Post(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.Group_post_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	session_info := dao.Get_session_by_id(t.Session_id)
	if (reflect.DeepEqual(session_info,(dao.Validation_info{}))) {
		return (" You have to login Again "),401
	}
	grp,_ := json.Marshal(dao.Create_group(t))

	return (string(grp)),200
}
