package models

import "strconv"

type LogReqAPI struct {
	Method     string
	Path       string
	StatusCode int
}

func (l LogReqAPI) String() string {
	return "Method: " + l.Method + ", Path: " + l.Path + ", StatusCode: " + strconv.Itoa(l.StatusCode)
}
