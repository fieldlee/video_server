package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)


var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB(){
	m,err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	m.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpeSession)
		sessionMap.Store(k,ss)
		return true
	})
}

func GenerateNewSessionId(un string)string{
	id := utils.NewUUID()
	if id == ""{
		return ""
	}
	ct := time.Now().UnixNano()/100000
	ttl := ct + 30 * 60 * 1000 /// 超时时间
	ss := &defs.SimpeSession{
		UserName:un,
		TTL:ttl,
	}
	sessionMap.Store(id,ss)

	err := dbops.InsertSession(id,ttl,un)
	if err != nil {
		return ""
	}
	return id
}

func IsSessionExpired(sid string)(string,bool){
	ss,ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano()/100000
		if ct > ss.(*defs.SimpeSession).TTL{
			//delete
			sessionMap.Delete(sid)
			err := dbops.DeleteSession(sid)
			if err != nil {

				return "",true
			}
			return "",true
		}
		return ss.(*defs.SimpeSession).UserName,false
	}
	return "",true
}