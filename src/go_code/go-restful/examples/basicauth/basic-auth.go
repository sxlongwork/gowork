package main

import (
	restful "github.com/emicklei/go-restful"
	_ "github.com/emicklei/go-restful-swagger12"
	"log"
	"net/http"
	"io"
	"fmt"
)

func main(){
	ws := new(restful.WebService)
	ws.Route(ws.GET("/secret").Filter(basicAuthenticate).To(secret))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func basicAuthenticate(req *restful.Request, res *restful.Response, chain *restful.FilterChain){
	u, p, ok := req.Request.BasicAuth()
	fmt.Printf("u=%v, p=%v, ok=%v\n", u, p, ok)
	if !ok || u != "admin" || p != "admin" {
		res.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		res.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, res)
}

func secret(req *restful.Request, res *restful.Response){
	io.WriteString(res,"acess success")
}