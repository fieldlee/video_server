package dbops

import "log"

func ReadVideoDeletionRecord(count int)([]string,error){
	var ids []string
	stmt,err := dbConn.Prepare("SELECT video_id from video_del_rec LIMIT ?")
	if err != nil {
		return ids , err
	}
	rows,err := stmt.Query(count)
	if err != nil {
		log.Printf("log : %v",err)
		return ids,err
	}
	for rows.Next(){
		var id string
		if err := rows.Scan(&id);err != nil {
			return ids,err
		}
		ids = append(ids,id)
	}
	defer stmt.Close()
	return ids,nil
}

func DeleteVideoDeletionRecord(vid string)error{
	stmt,err := dbConn.Prepare("DELETE FROM video_del_rec where video_id = ?")
	if err != nil {
		return   err
	}
	_,err = stmt.Exec(vid)
	if err != nil {
		log.Printf("DeleteVideoDeletionRecord : %v",err)
		return   err
	}
	return nil
}