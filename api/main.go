package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct{
	r *httprouter.Router
}

func NewMiddleWareHandler (r *httprouter.Router) middleWareHandler{
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w,r)
}

func regesiterHandler()*httprouter.Router{
	router := httprouter.New()

	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

func main()  {
	r := regesiterHandler()
	m := NewMiddleWareHandler(r)
	http.ListenAndServe(":8888",m)
}

