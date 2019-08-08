package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	ID string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	AuthorName string
	Content string
}

type SimpeSession struct {
	UserName string
	TTL int64
}

type SignedUP struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}