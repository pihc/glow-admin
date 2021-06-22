// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"payget/app/dao/internal"
)

// sysExampleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysExampleDao struct {
	*internal.SysExampleDao
}

var (
	// SysExample is globally public accessible object for table sys_example operations.
	SysExample sysExampleDao
)

func init() {
	SysExample = sysExampleDao{
		internal.NewSysExampleDao(),
	}
}

// Fill with you ideas below.