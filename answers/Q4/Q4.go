package main

import (
	"fmt"  
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"encoding/json"
)



type movies struct {
	Title string  `json : "title"`
	Year  int  	  `json : "year"`
	Director string `json : "director"`
	Caste string    `json : "caste"`
	Genre  string   `json : "genre"`
	Notes string    `json : "notes"`  
}




func create_entry(j int,k int,f []movies) {
	fmt.Println("hello")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("Movies").C("movies")
	for i:=j;i<k;i++ {
		post_var := movies {
			Title : f[i].Title,
	   		Year : f[i].Year,
	   		Director : f[i].Director,
	   		Caste : f[i].Caste,
	   		Genre : f[i].Genre,
	   		Notes : f[i].Notes,
		}
		
			err = c.Insert(post_var)
			if err != nil {
				panic(err)
			}	
	  }
	
}



func main() {
	raw, err := ioutil.ReadFile("./movies.json")
    if err != nil {
		fmt.Println(err)
	}
	var f []movies
	json.Unmarshal([]byte(raw),&f)
	
	n := len(f)
	k := (n/4)
	i := 0 
	go create_entry(i,k,f)
	go create_entry(k,k+k,f)
	go create_entry(k+k,k+k+k,f)
	go create_entry(k+k+k,n,f)
}