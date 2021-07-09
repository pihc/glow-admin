package define

import (
    "payget/library/query"
    "strings"

    "xorm.io/builder"
)

// ==========================================================================================
// API
// ==========================================================================================

// API删除
type ${.table.short_name|UcFirst}ApiDeleteReq struct {
    Id uint `v:"min:1#请选择需要删除的选项"`
}

// API创建/修改基类
type ${.table.short_name|UcFirst}ApiCreateUpdateBase struct {
${- range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type} `v:"required#请输入${$elem.comment}"` //${$elem.comment}
${- end}
}

// API创建
type ${.table.short_name|UcFirst}ApiCreateReq struct {
    ${.table.short_name|UcFirst}ApiCreateUpdateBase
}

// API修改
type ${.table.short_name|UcFirst}ApiUpdateReq struct {
    ${.table.short_name|UcFirst}ApiCreateUpdateBase
    Id uint `v:"min:1#请选择需要修改的选项"`
}

// ==========================================================================================
// Service
// ==========================================================================================

// Service查询
type ${.table.short_name|UcFirst}ServiceGetListReq struct {
    query.Params
    Name string `json:"short_name"`
}
func (q *${.table.short_name|UcFirst}ServiceGetListReq) Build() builder.Cond {
    cond := builder.NewCond()
    if q.Name != "" {
        cond = cond.And(builder.Like{"role.short_name", strings.TrimSpace(q.Name)})
    }
    return cond
}

// Service创建/修改基类
type ${.table.short_name|UcFirst}ServiceCreateUpdateBase struct {
${- range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type} `json:"${$elem.short_name}"` //${$elem.comment}
${- end}
}

// Service创建
type ${.table.short_name|UcFirst}ServiceCreateReq struct {
    ${.table.short_name|UcFirst}ServiceCreateUpdateBase
    CreatedBy uint `json:"created_by"`
}

// Service修改
type ${.table.short_name|UcFirst}ServiceUpdateReq struct {
    ${.table.short_name|UcFirst}ServiceCreateUpdateBase
    Id        uint `json:"id"`
    UpdatedBy uint `json:"updated_by"`
}

// Service创建返回结果
type ${.table.short_name|UcFirst}ServiceCreateRes struct {
    ${.table.short_name|UcFirst}Id uint `json:"${.table.short_name}_id"`
}
