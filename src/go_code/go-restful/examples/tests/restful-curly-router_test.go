package main

import (
	restful "github.com/emicklei/go-restful"
	"net/http"
	"log"
	"time"
	"errors"
	"testing"
	"bytes"
)

type User struct{
	Id, Name string
}

type UserResource struct{
	users map[string]User
}

func (user *UserResource) regiister(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("users").
	Consumes(restful.MIME_XML, restful.MIME_JSON).
	Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(user.findUser))
	ws.Route(ws.POST("").To(user.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(user.createUser))
	ws.Route(ws.DELETE("/{user-id}").To(user.removeUser))

	container.Add(ws)
}

func (user UserResource) findUser(req *restful.Request, res *restful.Response){
	uid := req.PathParameter("user-id")
	u := user.users[uid]
	if len(u.Id) == 0{
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusNotFound, "User could not be founf.")
	} else {
		res.WriteEntity(u)
	}
}

func (user *UserResource) updateUser(req *restful.Request, res *restful.Response){
	usr := new(User)
	err := req.ReadEntity(&usr)
	if err == nil{
		user.users[usr.Id] = *usr
		res.WriteEntity(usr)
	} else {
		res.AddHeader("Content-type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (user *UserResource) createUser(req *restful.Request, res *restful.Response){
	usr := User{Id: req.PathParameter("user-id")}
	err := req.ReadEntity(&usr)
	if err == nil{
		user.users[usr.Id] = usr
		res.WriteHeader(http.StatusCreated)
		res.WriteEntity(usr)
	} else {
		res.AddHeader("Content-type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (user *UserResource) removeUser(req *restful.Request, res *restful.Response){
	uid := req.PathParameter("user-id")
	delete(user.users, uid)
}

func RunRestfulCurlyRouterServer()  {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	u := UserResource{users: make(map[string]User)}
	u.regiister(wsContainer)

	log.Print("strat listening on localhost: 8090")
	server := &http.Server{Addr: ":8090", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func waitServerUp(serverurl string) error {
	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(5 *time.Second){
		_, err := http.Get(serverurl +"/")
		if err == nil {
			return nil
		}
	}
	return errors.New("waiting for server timed out")
}

func TestServer(t *testing.T){
	serverurl := "http://localhost:8090"
	go func() {
		RunRestfulCurlyRouterServer()
	}()
	if err := waitServerUp(serverurl); err != nil{
		t.Errorf("%v", err)
	}

	// GET request ,give a 405
	res, err := http.Get(serverurl + "/users/")
	if err != nil{
		t.Errorf("unexpected error in GET /users/: %v", err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexpected response: %v, expected: %v", res.StatusCode, http.StatusOK)
	}
	// send a post request
	var jsonstr =[]byte(`{"id":"1", "name":"user1"}`)
	req, err := http.NewRequest("POST", serverurl + "/users/", bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-type", restful.MIME_JSON)

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if res.StatusCode != http.StatusOK{
		t.Errorf("unexpected response: %v, expected: %v", res.StatusCode, http.StatusOK)
	}
	// Test that Get works
	res, err = http.Get(serverurl + "/users/1")
	if err != nil {
		t.Errorf("unexpected error in GET /users/1: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected response: %v, expected: %v", res.StatusCode, http.StatusOK)
	}
}