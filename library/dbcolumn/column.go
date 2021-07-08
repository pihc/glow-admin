package dbcolumn

import (
	"payget/library/tools"
	"strings"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
)

//数据库字符串类型
var ColumnTypeStr = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

//数据库时间类型
var ColumnTypeTime = []string{"datetime", "time", "date", "timestamp"}

//数据库数字类型
var ColumnTypeNumber = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

//页面不需要显示的列表字段
var COLUMNNAME_NOT_LIST = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

//页面不需要显示的列表字段
func IsNotList(value string) bool {
	return !tools.InArray(value, COLUMNNAME_NOT_LIST)
}

//获取数据库类型字段
func getDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

func GetGolangType(columnType string) string {
	goType := ""
	t := getDbType(columnType)
	if tools.InArray(t, ColumnTypeStr) {
		goType = "string"
	} else if tools.InArray(t, ColumnTypeTime) {
		goType = "Time"
	} else if tools.InArray(t, ColumnTypeNumber) {
		t, _ := gregex.ReplaceString(`\(.+\)`, "", columnType)
		t = gstr.Split(gstr.Trim(t), " ")[0]
		t = gstr.ToLower(t)
		// 如果是浮点型
		switch t {
		case "float", "double", "decimal":
			goType = "float64"
		case "bit", "int", "tinyint", "small_int", "smallint", "medium_int", "mediumint":
			if gstr.ContainsI(columnType, "unsigned") {
				goType = "uint"
			}
			goType = "int"
		case "big_int", "bigint":
			if gstr.ContainsI(columnType, "unsigned") {
				goType = "uint64"
			}
			goType = "int64"
		}
	}

	return goType
}
