package main

import (
	"log"
	"net/http"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type User struct {
	Id int
	Name string
	Email string
}

var (
	us = make(map[string]User)
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	
	log.Println("User requested at ", time.Now())
	w.WriteHeader(http.StatusOK)

	uresp, err := json.Marshal(us[vars["id"]])
	if err != nil {
		log.Println("error: ", err.Error())
	}

	w.Write([]byte(uresp))
}

func TriggerLogs(w http.ResponseWriter, r *http.Request) {

	count, err := strconv.Atoi(mux.Vars(r)["count"])
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	for i := 1; i<= count; i++ {
		go logValue(i)
	}
}

func logValue(i int) {
	log.Println("Logging number ", i)
}

func main() {
	
	rtr := mux.NewRouter()

	createUsers()

	rtr.Handle("/api/user/{id}", http.HandlerFunc(GetUser)).Methods("GET")
	
	rtr.Handle("/api/log/{count}", http.HandlerFunc(TriggerLogs)).Methods("POST")

	server :=
		&http.Server{
			Addr:    ":8182",
			Handler: rtr,
		}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("An error has occurred: %s", err)
	}

}

func createUsers() {
	us["1"] = User{
		Id: 1,
		Name: "Jansson",
		Email: "bamine_chatx@z-mail.cf",
	}

	us["2"] = User{
		Id: 2,
		Name: "Eriksson",
		Email: "ghoussam_fi@babyfriendly.app",
	}
}