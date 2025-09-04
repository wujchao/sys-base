// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDoc is the golang structure for table sys_doc.
type SysDoc struct {
	Id          string      `json:"id"          orm:"id"           description:""`                            //
	Name        string      `json:"name"        orm:"name"         description:"名称"`                          // 名称
	Key         string      `json:"key"         orm:"key"          description:"标识"`                          // 标识
	Title       string      `json:"title"       orm:"title"        description:"标题"`                          // 标题
	Desc        string      `json:"desc"        orm:"desc"         description:"简介"`                          // 简介
	Content     string      `json:"content"     orm:"content"      description:"内容，当内容为文件时，存储文件路径及文件类型的json"` // 内容，当内容为文件时，存储文件路径及文件类型的json
	ContentType uint        `json:"contentType" orm:"content_type" description:"1 html 2 markdowm 3 json"`    // 1 html 2 markdowm 3 json
	Status      uint        `json:"status"      orm:"status"       description:"状态 1 正常 2 禁用"`                // 状态 1 正常 2 禁用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                        // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`                        // 更新时间
}
