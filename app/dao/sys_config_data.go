// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"glow-admin/app/dao/internal"
)

// sysConfigDataDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysConfigDataDao struct {
	*internal.SysConfigDataDao
}

var (
	// SysConfigData is globally public accessible object for table sys_config_data operations.
	SysConfigData sysConfigDataDao
)

func init() {
	SysConfigData = sysConfigDataDao{
		internal.NewSysConfigDataDao(),
	}
}

// Fill with you ideas below.
