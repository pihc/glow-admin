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
type ${.table.name|UcFirst}ApiDeleteReq struct {
    Id uint `v:"min:1#请选择需要删除的选项"`
}

// API创建/修改基类
type ${.table.name|UcFirst}ApiCreateUpdateBase struct {
${range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type} `v:"required#请输入${$elem.comment}"` //${$elem.comment}
${end}
}

// API创建
type ${.table.name|UcFirst}ApiCreateReq struct {
    ${.table.name|UcFirst}ApiCreateUpdateBase
}

// API修改
type ${.table.name|UcFirst}ApiUpdateReq struct {
    ${.table.name|UcFirst}ApiCreateUpdateBase
    Id uint `v:"min:1#请选择需要修改的选项"`
}

// ==========================================================================================
// Service
// ==========================================================================================

// Service查询
type ${.table.name|UcFirst}ServiceGetListReq struct {
    query.Params
    Name string `json:"name"`
}
func (q *${.table.name|UcFirst}ServiceGetListReq) Build() builder.Cond {
    cond := builder.NewCond()
    if q.Name != "" {
        cond = cond.And(builder.Like{"role.name", strings.TrimSpace(q.Name)})
    }
    return cond
}

// Service创建/修改基类
type ${.table.name|UcFirst}ServiceCreateUpdateBase struct {
${range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type} `json:"${$elem.name}"` //${$elem.comment}
${end}
}

// Service创建
type ${.table.name|UcFirst}ServiceCreateReq struct {
    ${.table.name|UcFirst}ServiceCreateUpdateBase
    CreatedBy uint `json:"created_by"`
}

// Service修改
type ${.table.name|UcFirst}ServiceUpdateReq struct {
    ${.table.name|UcFirst}ServiceCreateUpdateBase
    Id        uint `json:"id"`
    UpdatedBy uint `json:"updated_by"`
}

// Service创建返回结果
type ${.table.name|UcFirst}ServiceCreateRes struct {
    ${.table.name|UcFirst}Id uint `json:"${.table.name}_id"`
}
