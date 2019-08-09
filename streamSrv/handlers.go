package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


func streamHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid_id")
	vl := VIDEOPATH + vid
	video , err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w,http.StatusInternalServerError,"internal error")
		return
	}
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,vid,time.Now(),video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w,http.StatusBadRequest,"too big for this video")
		return
	}
	file ,_, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w,http.StatusBadRequest,err.Error())
		return
	}
	//accept := "video/*"
	data , err := ioutil.ReadAll(file)
	if err != nil {
		sendErrorResponse(w,http.StatusBadRequest,err.Error())
		return
	}
	fn := p.ByName("vid_id")
	path := VIDEOPATH + fn
	err = ioutil.WriteFile(path,data,0666)
	if err != nil {
		sendErrorResponse(w,http.StatusBadRequest,err.Error())
		return
	}
	sendNormalResponse(w,200,"video had uploaded")
	return
}