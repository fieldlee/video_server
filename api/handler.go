package main
import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res,ubody); err != nil {
		errResp := defs.ErrorResponse{
			HttpSC:400,
			Error:defs.Err{
				Error:err.Error(),
				ErrorCode:"201",
			},
		}
		sendErrorResponse(w,errResp)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd);err != nil {
		sendErrorResponse(w,defs.ErrorDBERROR)
		return
	}

	id := utils.NewUUID()
	s_id := session.GenerateNewSessionId(id)

	su := &defs.SignedUP{Success:true,SessionId:s_id}
	subyte , err := json.Marshal(su)
	if err != nil {
		sendErrorResponse(w,defs.ErrorMarsha)
		return
	}
	sendNormalResponse(w,string(subyte),200)
	return
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}