package dbops

import (
	"testing"
)

// login db -> truncate talbes -> run tests -> clear tables truncatetable
func clearTable(){
	dbConn.Exec("Truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTable()
	m.Run()
	clearTable()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("add",TestAddUser)
	t.Run("get",TestGetUser)
	t.Run("del",TestDeleteUser)
	t.Run("reget",TestReGet)
	t.Run("test view",TestAddNewVideo)
	t.Run("add new comment",TestAddNewComment)
	t.Run("list comments",testListComments)
}

func TestAddNewVideo(t *testing.T) {
	v , err := AddNewVideo(1,"testview")
	if err != nil {
		t.Errorf("TestAddNewVideo :%s",err.Error())
	}
	t.Logf("TestAddNewVideo :%v",v)


	gv , err := GetVideo(v.ID)

	if err != nil {
		t.Errorf("get view :%s",err.Error())
	}
	t.Logf("GetVideo :%v",gv)


	err = DeleteVideo(v.ID)
	if err != nil {
		t.Errorf("delete video :%s",err.Error())
	}

	gv2 , err := GetVideo(v.ID)
	if err != nil {
		t.Errorf("get view :%s",err.Error())
	}
	t.Logf("reGetVideo :%v",gv2)
}



func TestAddUser(t *testing.T){
	err := AddUserCredential("lidepeng","12345")
	if err != nil {
		t.Errorf("err add user :%s",err.Error())
	}
}

func TestGetUser(t *testing.T){
	u, err :=  GetUserCredential("lidepeng")
	if err != nil {
		t.Errorf("err get user :%s",err.Error())
	}
	t.Logf("GET USER :%v",u)
}

func TestDeleteUser(t *testing.T){
	err := DeleteUserCredential("lidepeng","12345")
	if err != nil {
		t.Errorf("err del user :%s",err.Error())
	}

}

func TestReGet(t *testing.T){
	u, err :=  GetUserCredential("lidepeng")
	if err != nil {
		t.Errorf("err reget user :%s",err.Error())
	}
	t.Logf("reGET USER :%v",u)
}

func TestAddNewComment(t *testing.T) {
	vid := "123456"
	aid := 1
	content := "sdfadsfasdf"
	err := AddNewComment(vid,aid,content)
	if err != nil {
		t.Errorf("TestAddNewComment :%s",err.Error())
	}
}

func testListComments(t *testing.T) {
	comlist, err := ListComments("123456")
	if err != nil {
		t.Errorf("TestAddNewComment :%s",err.Error())
	}
	t.Logf("TestListComments:%v",comlist)
}