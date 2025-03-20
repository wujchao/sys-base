package model

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type PaginationInput struct {
	Param    map[string]interface{} `json:"param" dc:"Search field parameters, written as param[field name], value supports string and []string, array [0] is the comparison operator, [1] is the value"`
	Keyword  string                 `json:"keyword" dc:"Search Keywords"` //搜索关键字
	Between  string                 `json:"between" p:"between"`          //日期范围
	OrderBy  string                 `json:"orderBy"`                      //排序方式
	Page     int                    `json:"page" in:"query" d:"1"  v:"min:-1#Wrong page number"     dc:"Paging number, default 1"`
	PageSize int                    `json:"pageSize" in:"query" d:"10" v:"max:500#Maximum number of pages: 500" dc:"Number of pages, maximum 500"`
}

type PaginationOutput struct {
	CurrentPage int `json:"currentPage" dc:"Current Page"`
	Total       int `dc:"total"`
}

type PaginationConfig struct {
	*PaginationInput
	Fields *map[string]interface{} `json:"fields"`
}

// Pagination PaginationModel 通用查询条件处理
func Pagination(cfg *PaginationConfig) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if cfg.Keyword != "" {
			keywordField := defVal(cfg.Fields, "keyword", "name")
			m = m.WhereLike(keywordField, cfg.Keyword)
		}
		if between := betweenToTimes(cfg.Between); between != nil {
			betweenFiled := defVal(cfg.Fields, "between", "created_at")
			m = m.WhereBetween(betweenFiled, between[0], between[1])
		}

		// 默认倒序
		if cfg.OrderBy != "" {
			if gstr.ContainsI(cfg.OrderBy, "desc") || gstr.ContainsI(cfg.OrderBy, "asc") {
				m = m.Order(cfg.OrderBy)
			} else {
				m = m.OrderDesc(cfg.OrderBy)
			}
		} else {
			if ok, _ := m.HasField("id"); ok {
				m = m.OrderDesc("id")
			}
		}

		if cfg.Param != nil {
			// 字段 array
			if f, ok := (*cfg.Fields)["Columns"]; ok {
				searchKeys := f.([]string)
				operator := []string{"like", "LIKE", "between", "BETWEEN", ">", ">=", "=", "<", "<=", "<>"}
				// 参数 等于判断
				for k, v := range cfg.Param {
					if gstr.InArray(searchKeys, k) {
						switch value := v.(type) {
						case []string:
							if len(value) >= 2 && gstr.InArray(operator, value[0]) {
								// like
								if gstr.Equal(value[0], "like") {
									m = m.Where(k+" like ?", value[1])
								} else if gstr.Equal(value[0], "between") && len(value) >= 3 {
									// between
									m = m.WhereBetween(k, value[1], value[2])
								} else {
									// >,>=,=,<,<=
									m = m.Where(k+" "+value[0]+" ?", value[1])
								}
							}
							break
						default:
							m = m.Where(k, v)
						}
					}
				}
			}
		}

		// page && pageSize = -1时 返回所有数据
		if cfg.Page != -1 && cfg.PageSize != -1 {
			if cfg.Page == 0 {
				cfg.Page = 1
			}
			if cfg.PageSize == 0 {
				cfg.PageSize = 10
			}
			m = m.Page(cfg.Page, cfg.PageSize)
		}

		return m
	}
}

func betweenToTimes(between string) []string {
	if !gstr.Contains(between, ",") {
		return nil
	}
	betweenArray := gstr.Split(between, ",")
	if len(betweenArray) != 2 {
		return nil
	}
	if gstr.IsNumeric(betweenArray[0]) {
		// 时间戳
		return []string{
			gconv.String(gtime.NewFromTimeStamp(gconv.Int64(betweenArray[0])).String()),
			gconv.String(gtime.NewFromTimeStamp(gconv.Int64(betweenArray[1])).String()),
		}
	} else {
		// 字符串
		return []string{betweenArray[0], betweenArray[1]}
	}
}

func defVal(fields *map[string]interface{}, field, def string) string {
	if fields == nil {
		return def
	}
	if f, ok := (*fields)[field]; ok {
		if tField, sOk := f.(string); sOk {
			return tField
		}
	}
	return def
}
