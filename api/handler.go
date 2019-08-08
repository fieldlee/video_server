package main
import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	io.WriteString(w,"create user successful")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}