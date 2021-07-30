package service

import (
	"context"
	"fmt"
	"payget/library/dbcolumn"
	"payget/library/tools"
	"strings"

	"github.com/gogf/gf/os/gfile"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/gview"

	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/frame/g"
)

var GenCode = genCodeService{
	tablePrefix:  []string{"app_", "sys_"},                                               // 表前缀
	columnHidden: []string{"created_at", "updated_at", "updated_by", "created_by", "id"}, // 过滤表字段
	modules:      []string{"define", "service", "api"},                                   // 自动生成模块
}

type genCodeService struct {
	tablePrefix, columnHidden, modules []string
}

func (s *genCodeService) GenData(ctx context.Context, tableName string) error {
	// 获取表备注
	tableComment, err := s.getTableComment(tableName)
	if err != nil {
		return err
	}

	// 获取表字段信息
	mapFields, err := g.DB().TableFields(ctx, tableName)
	if err != nil {
		return err
	}

	// 去除表前缀
	shortName := ""
	for _, v := range s.tablePrefix {
		if strings.Contains(tableName, v) {
			shortName = strings.Replace(tableName, v, "", -1)
			break
		}
	}

	// 过滤表字段
	var fields []map[string]string
	for k, v := range mapFields {
		if !tools.InArray(k, s.columnHidden) {
			fields = append(fields, map[string]string{
				"name":    v.Name,
				"comment": v.Comment,
				"type":    dbcolumn.GetGolangType(v.Type),
			})
		}
	}

	// 模板变量
	fillData := g.Map{
		"short_name": shortName,
		"name":       tableName,
		"fields":     fields,
		"comment":    tableComment,
	}

	for _, v := range s.modules {
		// 填充模板
		content, err := s.fillAndGetContent(fmt.Sprintf("code/%s.tpl", v), fillData)
		if err != nil {
			return err
		}
		// 模板保存到文件
		if err = gfile.PutContents(fmt.Sprintf("./app/module/admin/internal/%s/%s.go", v, shortName), content); err != nil {
			return err
		}
	}

	// 打印路由
	fmt.Println(fmt.Sprintf("// %s", tableComment))
	fmt.Println(fmt.Sprintf(`group.GET("/%s/index", respond.Convert(api.%s.GetList))`, shortName, gstr.UcFirst(shortName)))
	fmt.Println(fmt.Sprintf(`group.POST("/%s/add", respond.Convert(api.%s.Create))`, shortName, gstr.UcFirst(shortName)))
	fmt.Println(fmt.Sprintf(`group.POST("/%s/edit", respond.Convert(api.%s.Update))`, shortName, gstr.UcFirst(shortName)))
	fmt.Println(fmt.Sprintf(`group.POST("/%s/delete/:id", respond.Convert(api.%s.Delete))`, shortName, gstr.UcFirst(shortName)))
	fmt.Println(fmt.Sprintf(`group.GET("/%s/list", respond.Convert(api.%s.GetAll))`, shortName, gstr.UcFirst(shortName)))
	fmt.Println(fmt.Sprintf(`group.GET("/%s/info/:id", respond.Convert(api.%s.GetDetail))`, shortName, gstr.UcFirst(shortName)))

	return nil
}

// 填充模板并且获取填充后的模板
func (s *genCodeService) fillAndGetContent(tplName string, fillData g.Map) (string, error) {
	// 创建模板引擎
	view := gview.New()
	view.BindFuncMap(g.Map{
		"UcFirst": func(str string) string {
			return gstr.UcFirst(str)
		},
		"CaseCamel": func(str string) string {
			return gstr.CaseCamel(str)
		},
	})
	// 配置
	_ = view.SetConfigWithMap(g.Map{
		"Paths":      []string{"template"}, // 模板文件搜索目录路径
		"Delimiters": []string{"${", "}"},  // 模板引擎变量分隔符号。默认为 ["{{", "}}"]
	})
	// 填充模板
	result, err := view.Parse(context.Background(), tplName, g.Map{"table": fillData})
	if err != nil {
		return "", err
	}
	// 剔除多余的换行
	//content, err := s.trimBreak(result)
	//if err != nil {
	//	return "", err
	//}

	return result, nil
}

// 获取table备注
func (s *genCodeService) getTableComment(tableName string) (string, error) {
	link := g.Cfg().Get("database.link")
	if link == "" {
		return "", gerror.New("数据库配置读取错误")
	}
	linkArr := strings.Split(link.(string), "/")
	data, err := g.DB().GetArray(`SELECT table_comment FROM information_schema.TABLES WHERE TABLE_SCHEMA = ? and table_name = ?`, linkArr[1], tableName)
	if err != nil {
		return "", err
	}
	tableComment := ""
	if len(data) > 0 {
		tableComment = data[0].String()
	}
	tableComment = strings.TrimRight(tableComment, "表")
	return tableComment, nil
}

// 剔除多余的换行
func (s *genCodeService) trimBreak(str string) (string, error) {
	b, err := gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str))
	if err != nil {
		return "", err
	}
	return gconv.String(b), nil
}
