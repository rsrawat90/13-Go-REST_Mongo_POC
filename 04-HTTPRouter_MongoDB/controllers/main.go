package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/RSR2019/GO-POC/04-RESTService_MongoDB/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)
// UserController represents the controller for operating on the User resource
type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("test-rest-service").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
    uc.session.DB("test-rest-service").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
	w.WriteHeader(http.StatusNotFound) // 404
	return
}
    oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
	w.WriteHeader(404)
	return
}
}