// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"focus/app/dao/internal"
)

// userLoginDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userLoginDao struct {
	*internal.UserLoginDao
}

var (
	// UserLogin is globally public accessible object for table {TplTableName} operations.
	UserLogin = &userLoginDao{
		internal.UserLogin,
	}
)

// Fill with you ideas below.