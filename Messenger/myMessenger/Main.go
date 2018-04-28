package main

import (
	"myMessenger/servicehandlers"
	"log"
	"net/http"
)

func main() {

	u := servicehandlers.UserHandler{}
	g := servicehandlers.GroupHandler{}
	s := servicehandlers.UserValidationHandler{}
	m := servicehandlers.MsgHandler{}
	v := servicehandlers.ViewHandler{}
	//p := servicehandlers.PingHandler{}

	http.Handle("/myMessenger/user", u)
	http.Handle("/myMessenger/group",g)
	http.Handle("/myMessenger/userValidation",s)
	http.Handle("/myMessenger/message",m)
	http.Handle("/myMessenger/view",v)
	//http.Handle("/ping", p)

	x := http.ListenAndServe(":8080", nil)
	log.Fatal(x)
}
