package dbops

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

//create table users (id int(4) PRIMARY KEY AUTO_INCREMENT, login_name varchar(64),pwd TEXT);
//create table video_info(id varchar(64) primary key not null,author_id int(4),name text,display_ctime text,create_time datetime);
//create table comments(id varchar(64) primary key not null , video_id varchar(64),author_id int(4),content TEXT , time datetime);
//create table sessions (session_id varchar(100) primary key not null,ttl tinytext,login_name varchar(64));

func AddUserCredential(loginName string,pwd string) error  {
	ins,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_,err = ins.Exec(loginName,pwd)
	if err != nil {
		return err
	}
	defer  ins.Close()
	return nil
}

func GetUserCredential(loginName string)(string,error){
	 outs,err := dbConn.Prepare("SELECT pwd from users where  login_name = ?")
	if err != nil {
		log.Printf("GetUserCredential :%s",err.Error())
		return "",err
	}

	 var pwd string
	 err = outs.QueryRow(loginName).Scan(&pwd)
	 if err != nil && err != sql.ErrNoRows {
	 	return "",err
	 }
	 defer outs.Close()
	 return pwd,nil
}

func DeleteUserCredential(loginName string,pwd string) error  {
	del,err := dbConn.Prepare("delete FROM users WHERE  login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("DeleteUserCredential : %s",err.Error())
		return err
	}
	_,err = del.Exec(loginName,pwd)
	if err != nil {
		return err
	}
	defer del.Close()
	return nil
}

func AddNewVideo(aid int,name string) (*defs.VideoInfo,error) {
	//create uid
	uid := utils.NewUUID()
	if uid == ""{
		return nil,errors.New("uid err")
	}
	// createtime
	createtime := time.Now()
	ctime := createtime.Format("Jan 02 2006,15:04:05")
	ins,err := dbConn.Prepare("INSERT into video_info (id,author_id,name,display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil,err
	}
	_,err = ins.Exec(uid,aid,name,ctime)
	if err != nil {
		return nil,err
	}
	res := defs.VideoInfo{
		ID:uid,
		AuthorId:aid,
		Name:name,
		DisplayCtime:ctime,
	}
	defer ins.Close()
	return &res,nil
}
func GetVideo(vid string) (*defs.VideoInfo,error) {

	stmt,err := dbConn.Prepare("SELECT author_id ,name,display_ctime FROM video_info where id = ?")
	if err != nil {
		return nil,err
	}
	var aid int
	var name string
	var ctime string
	err = stmt.QueryRow(vid).Scan(&aid,&name,&ctime)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}
	res := defs.VideoInfo{
		ID:vid,
		AuthorId:aid,
		Name:name,
		DisplayCtime:ctime,
	}
	defer stmt.Close()
	return &res,nil
}

func DeleteVideo(vid string)error{
	delstmt,err := dbConn.Prepare("delete FROM video_info WHERE  id = ?")
	if err != nil {
		return err
	}
	_,err = delstmt.Exec(vid)
	if err != nil {
		return err
	}
	defer delstmt.Close()
	return nil
}

func AddNewComment(vid string,aid int, content string)error{
	id := utils.NewUUID()
	if id == "" {
		return errors.New("id error")
	}
	instmt,err := dbConn.Prepare("INSERT into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}

	_,err = instmt.Exec(id,vid,aid,content)

	if err != nil {
		return err
	}
	defer instmt.Close()
	return nil
}

func ListComments(vid string)([]*defs.Comment,error){
	stm , err := dbConn.Prepare("SELECT comments.id,users.login_name,comments.content FROM comments inner join users on comments.author_id = users.id where comments.video_id = ? ")
	if err != nil {
		return nil , err
	}
	rows,err := stm.Query(vid)
	if err != nil {
		return nil , err
	}

	var res []*defs.Comment

	for rows.Next() {
		var id , name , content string
		if err := rows.Scan(&id,&name,&content); err != nil{
			return res,err
		}

		c := &defs.Comment{
			id,
			vid,
			name,
			content,
		}

		res = append(res,c)
	}
	return res,nil
}