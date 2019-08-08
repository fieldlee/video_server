package dbops

import (
	"database/sql"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string,ttl int64,uname string)error{
	ttlstr := strconv.FormatInt(ttl,10)
	stmins , err := dbConn.Prepare("INSERT INTO sessions (session_id,ttl,login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_,err = stmins.Exec(sid,ttlstr,uname)
	if err != nil {
		return err
	}
	defer stmins.Close()
	return nil
}

func RetrieveSession(sid string)(*defs.SimpeSession,error)  {
	smtout , err := dbConn.Prepare("select ttl,login_name from sessions where session_id = ?")
	if err != nil {
		return nil , err
	}
	var ttl string
	var uname string
	err = smtout.QueryRow(sid).Scan(&ttl,&uname)

	if err != nil && err != sql.ErrNoRows{
		return nil , err
	}
	ttlint ,err := strconv.ParseInt(ttl,10,64)
	if err != nil {
		return nil , err
	}
	ss := &defs.SimpeSession{
		UserName:uname,
		TTL:ttlint,
	}
	return ss,nil
}

func RetrieveAllSession()(*sync.Map,error){
	m := &sync.Map{}
	smtout , err :=dbConn.Prepare("select session_id,ttl,login_name from sessions")
	if err != nil {
		return nil , err
	}
	rows,err := smtout.Query()
	if err != nil {
		return nil , err
	}
	for rows.Next(){
		var id string
		var ttlstr string
		var uname string
		if err := rows.Scan(&id,&ttlstr,&uname);err != nil {
			break
		}
		if ttl,err := strconv.ParseInt(ttlstr,10,64);err != nil {
			ss := &defs.SimpeSession{UserName:uname,TTL:ttl,}
			m.Store(id,ss)
		}
	}
	return m,nil
}

func DeleteSession(sid string)error{
	delstmt,err := dbConn.Prepare("delete FROM sessions WHERE  session_id = ?")
	if err != nil {
		return err
	}
	_,err = delstmt.Query(sid)
	if err != nil {
		return err
	}
	defer delstmt.Close()

	return nil
}