package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter,sc int,errMsg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,errMsg)
}

func sendNormalResponse(w http.ResponseWriter,sc int,info string)  {
	w.WriteHeader(sc)
	io.WriteString(w,info)
}