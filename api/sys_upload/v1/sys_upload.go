package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" auth:"open" tags:"File Upload" summary:"File Upload"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"Select Upload File"`
}

type FileUploadRes struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FileUploadForEditorReq struct {
	g.Meta `path:"/upload/forEditor" method:"post" mime:"multipart/form-data" auth:"open" tags:"File Upload" summary:"File Upload"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"Select Upload File"`
}

type FileUploadForEditorRes struct {
	Errno   int                      `json:"errno"`
	Message string                   `json:"message"`
	Data    *FileUploadForEditorData `json:"data"`
}

type FileUploadForEditorData struct {
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Href string `json:"href"`
}
