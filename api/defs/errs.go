package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}
type ErrorResponse struct {
	HttpSC int
	Error Err

}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,Error:Err{Error:"request body parse failed",ErrorCode:"201"}}
	ErrorNOTAuthUser = ErrorResponse{HttpSC:401,Error:Err{Error:"user auth error",ErrorCode:"202"}}
)