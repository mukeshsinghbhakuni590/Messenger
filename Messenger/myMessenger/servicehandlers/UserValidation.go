package servicehandlers


import (
	"myMessenger/dao"
	//"myMessenger/validators"
	"encoding/json"
	"net/http"
	"reflect"
	"fmt"
)


type UserValidationHandler struct {
}

func (p UserValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p UserValidationHandler) Get(req *http.Request) (string,int) {
	return "Get UserValidationHandler called",408
}

func (p UserValidationHandler) Put(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.Delete_session
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	fmt.Print(t.Email ,"\n")
	user := dao.Get_user_by_email(t.Email)
	if (reflect.DeepEqual(user,(dao.User_info{}))) {
		fmt.Print(dao.User_info{},"\n")
		return ("user not exist!!"),404
	} else {
		user_session := dao.Get_session_by_email(t.Email)
		if (reflect.DeepEqual(user_session,(dao.Validation_info{}))) {
			return "You Have To Login Again",401
		}else{
			dao.Delete_Session(t)
			return "Session Deleted successfully",410
		}
	}
}

func (p UserValidationHandler) Post(req *http.Request) (string,int) {
	decoder := json.NewDecoder(req.Body)
    var t dao.Validation_post_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
	}
	defer req.Body.Close()
	fmt.Print(t.Email ,"\n")
	user := dao.Get_user_by_email(t.Email)

	if (reflect.DeepEqual(user,(dao.User_info{}))) {
		fmt.Print(dao.User_info{},"\n")
		return ("user not exist!!"),401
	} else {
		if t.Passwd != user.Passwd {
			return ("Incorrect Password !!"),401
		}else{
			user_session := dao.Get_session_by_email(t.Email)
			if !(reflect.DeepEqual(user_session,(dao.Validation_info{}))) {
				session,_ := json.Marshal(user_session)
				return (string(session)),200
			}else{
				session,_ := json.Marshal(dao.Create_Session(t))
				return (string(session)),200
			}
		}
	}
}
