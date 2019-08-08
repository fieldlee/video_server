package utils

import uuid "github.com/satori/go.uuid"

func NewUUID()string{
	uid,err := uuid.NewV1()
	if err != nil {
		return ""
	}
	return uid.String()
}
