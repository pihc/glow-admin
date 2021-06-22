// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"payget/app/dao/internal"
)

// sysLevelDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysLevelDao struct {
	*internal.SysLevelDao
}

var (
	// SysLevel is globally public accessible object for table sys_level operations.
	SysLevel sysLevelDao
)

func init() {
	SysLevel = sysLevelDao{
		internal.NewSysLevelDao(),
	}
}

// Fill with you ideas below.