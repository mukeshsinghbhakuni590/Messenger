package servicehandlers

import (
	"net/http"
)

type HttpServiceHandler interface {
	Get(*http.Request) (string,int)
	Put(*http.Request) (string,int)
	Post(*http.Request) (string,int)
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) {
	var v string
	var e int
	if r.Method == "GET" {
		v,e = p.Get(r)
		w.WriteHeader(e)
		w.Write([]byte(v))
	} else if r.Method == "PUT" {
		v,e = p.Put(r)
		w.WriteHeader(e)
		w.Write([]byte(v))
	} else if r.Method == "POST" {
		v,e = p.Post(r)
		w.WriteHeader(e)
		w.Write([]byte(v))
	}

}
