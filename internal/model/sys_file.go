package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File       *ghttp.UploadFile `json:"file"`       // 文件对象
	Name       string            `json:"name"`       // 文件名
	RandomName bool              `json:"randomName"` // 是否随机文件名
}

type FileUploadOutput struct {
	Name string `json:"name"` // 文件名
	Path string `json:"path"` // 文件路径
	Url  string `json:"url"`  // 文件访问路径
}
