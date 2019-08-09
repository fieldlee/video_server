package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWare struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWare(r *httprouter.Router,conn int) middleWare {
	m := middleWare{}
	m.r = r
	m.l = NewConnLimiter(conn)
	return m
}

func (m middleWare)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if m.l.GetConn() == false {
		sendErrorResponse(w,http.StatusTooManyRequests,"too many request")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers()*httprouter.Router{
	router := httprouter.New()

	router.GET("/videos/:vid_id",streamHandler)
	router.POST("/upload/:vid_id",uploadHandler)

	return router
}

func main()  {
	r := RegisterHandlers()
	m := NewMiddleWare(r,2)
	http.ListenAndServe(":9000",m)
}