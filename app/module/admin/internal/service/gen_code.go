package service

import (
	"context"
	"payget/library/dbcolumn"
	"payget/library/tools"
	"strings"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/gview"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/frame/g"
)

var GenCode = genCodeService{}

type genCodeService struct {
}

// 代码生成
func (s *genCodeService) GenData(ctx context.Context, tableName string) error {
	// 获取表字段信息
	mapFields, err := g.DB("default").TableFields(ctx, tableName)
	if err != nil {
		return err
	}

	g.Dump(mapFields)

	// 去除表前缀
	for _, v := range []string{"app_", "sys_"} {
		if strings.Contains(tableName, v) {
			tableName = strings.Replace(tableName, v, "", -1)
			break
		}
	}

	// 过滤表字段
	columnHidden := []string{"created_at", "updated_at", "updated_by", "created_by", "id"}
	var fields []map[string]string
	for k, v := range mapFields {
		if !tools.InArray(k, columnHidden) {
			fields = append(fields, map[string]string{
				"name":    v.Name,
				"comment": v.Comment,
				"type":    dbcolumn.GetGolangType(v.Type),
			})
		}
	}

	// 模板变量
	data := g.Map{
		"name":   tableName,
		"fields": fields,
	}

	// 创建模板引擎
	view := gview.New()
	view.BindFuncMap(g.Map{
		"UcFirst": func(str string) string {
			return gstr.UcFirst(str)
		},
		"CaseCamel": func(str string) string {
			return gstr.CaseCamel(str)
		},
		//"add": func(a, b int) int {
		//	return a + b
		//},
	})
	_ = view.SetConfigWithMap(g.Map{
		"Paths":      []string{"template"}, // 模板文件搜索目录路径
		"Delimiters": []string{"${", "}"},  // 模板引擎变量分隔符号。默认为 ["{{", "}}"]
	})
	result, err := view.Parse(ctx, "code/define.tpl", g.Map{"table": data})
	if err != nil {
		return err
	}
	content, err := s.trimBreak(result)
	if err != nil {
		return err
	}
	// 模板保存到文件
	err = gfile.PutContents("./app/module/admin/internal/define/"+tableName+".go", content)
	return err
}

//剔除多余的换行
func (s *genCodeService) trimBreak(str string) (res string, err error) {
	var b []byte
	if b, err = gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str)); err == nil {
		res = gconv.String(b)
	}
	return
}
