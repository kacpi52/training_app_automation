package data

import "mime/multipart"

type Params struct {
	AppLanguage string
	Header   string
	Query    string
	Param    string
	Json     map[string]interface{}
	FormData map[string][]*multipart.FileHeader
	FormDataParams map[string]interface{}
}