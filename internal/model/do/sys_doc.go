// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDoc is the golang structure of table sys_doc for DAO operations like Where/Data.
type SysDoc struct {
	g.Meta      `orm:"table:sys_doc, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Key         interface{} // 标识
	Title       interface{} // 标题
	Desc        interface{} // 简介
	Content     interface{} // 内容，当内容为文件时，存储文件路径及文件类型的json
	ContentType interface{} // 1 html 2 markdowm 3 json
	Status      interface{} // 状态 1 正常 2 禁用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
