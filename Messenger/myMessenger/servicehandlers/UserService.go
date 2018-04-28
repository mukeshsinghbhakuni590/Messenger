package servicehandlers

import (
	"encoding/json"
	"fmt"
	"myMessenger/dao"
	"myMessenger/validators"
	"net/http"
	"reflect"
)

type UserHandler struct {
}

func (p UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p UserHandler) Get(req *http.Request) (string, int) {
	userid, ok1 := req.URL.Query()["userid"]
	sessionid, ok2 := req.URL.Query()["sessionid"]

	if !ok1 || !ok2 {
		return ("Somthing!!"), 404
	} else if userid[0] != "" && sessionid[0] != "" {
		fmt.Println(userid[0])
		fmt.Println(sessionid[0])
		user_info := dao.Get_user_by_id(userid[0])
		if reflect.DeepEqual(user_info, (dao.User_info{})) {
			return ("user not exist!! "), 404
		} else {
			session_info := dao.Get_session_by_email(user_info.Email)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again"), 401
			} else {
				usr, _ := json.Marshal(user_info)
				return (string(usr)), 200
			}
		}
	} else {
		return ("Parameters Not correctly defined !!"), 404
	}
}

//----------------------------------------------------------------------------
func (p UserHandler) Put(req *http.Request) (string, int) {
	operation, ok := req.URL.Query()["operation"]
	if !ok {
		return ("Something Worng !!"), 404
	} else if len(operation) == 1 {
		decoder := json.NewDecoder(req.Body)
		if operation[0] == "changeName" {
			var t dao.User_change_name
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_changename(t) {
					grp, _ := json.Marshal(dao.Change_user_name(t))
					return (string(grp)), 200
				} else {
					return ("Change Username request denied !!"), 400
				}
			}
		} else if operation[0] == "changeStatus" {
			var t dao.User_change_status
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_changestatus(t) {
					grp, _ := json.Marshal(dao.Change_user_status(t))
					return (string(grp)), 200
				} else {
					return ("Change Status request denied !!"), 400
				}
			}
		} else if operation[0] == "changePasswd" {
			var t dao.User_change_passwd
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_changepasswd(t) {
					grp, _ := json.Marshal(dao.Change_user_passwd(t))
					return (string(grp)), 200
				} else {
					return ("Change Password request denied !!"), 400
				}
			}
		} else if operation[0] == "addCont" {
			var t dao.Add_Block_contact
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_addcontact(t) {
					grp, _ := json.Marshal(dao.Update_add_block_list(t, "ADD"))
					return (string(grp)), 200
				} else {
					return ("Add Contact request denied !!"), 400
				}
			}
		} else if operation[0] == "blockCont" {
			var t dao.Add_Block_contact
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_blockcontact(t) {
					grp, _ := json.Marshal(dao.Update_add_block_list(t, "BLOCK"))
					return (string(grp)), 200
				} else {
					return ("Block Contact request denied !!"), 400
				}
			}
		} else if operation[0] == "leaveGroup" {
			var t dao.Leave_group_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
				if validators.Validate_for_leavegroup(t) {
					dao.Remove_self_from_group(t)
					grp, _ := json.Marshal(dao.Leave_group(t))
					return (string(grp)), 200
				} else {
					return ("Leave Group request denied !!"), 400
				}
			}
		} else if operation[0] == "continfo" {
			var t dao.Get_cont_info
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer req.Body.Close()
			session_info := dao.Get_session_by_id(t.Session_id)
			if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
				return (" You have to login Again "), 401
			} else {
					cont_info, _ := json.Marshal(dao.Get_contact_by_id(t.Cont_id))
					return (string(cont_info)), 200
				}
			} else if operation[0] == "otherinfo" {
				var t dao.Other_user_get_struct
				err := decoder.Decode(&t)
				if err != nil {
					panic(err)
				}
				defer req.Body.Close()
				session_info := dao.Get_session_by_id(t.Session_id)
				if reflect.DeepEqual(session_info, (dao.Validation_info{})) {
					return (" You have to login Again "), 401
				} else {
						other_info, _ := json.Marshal(dao.Get_other_user_info(t))
						return (string(other_info)), 200
					}
				}  else {
				return ("Invalid Operation"), 405
		}
	} else {
		return ("multiple operations  or none !!"), 406
	}
}

//-----------------------------------------------------------------------------
func (p UserHandler) Post(req *http.Request) (string, int) {

	decoder := json.NewDecoder(req.Body)
	var t dao.User_post_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Print(t.Email, "\n")
	user := dao.Get_user_by_email(t.Email)

	if !(reflect.DeepEqual(user, (dao.User_info{}))) {
		fmt.Print(dao.User_info{}, "\n")
		return ("user already exist!!"), 409
	} else {
		usr, _ := json.Marshal(dao.Create_user(t))
		return (string(usr)), 200
	}
}
