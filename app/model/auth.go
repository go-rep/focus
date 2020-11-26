// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
)

// Auth is the golang structure for table gf_auth.
type Auth internal.Auth

// 权限树形列表
type AuthTree struct {
	Id       uint        `json:"id"`        // 分类ID，自增主键
	ParentId uint        `json:"parent_id"` // 父级分类ID，用于层级管理
	Name     string      `json:"name"`      // 权限名称
	Key      string      `json:"key"`       // 权限键名(用于程序)
	Value    string      `json:"value"`     // 权限键值，部分自定义权限可能有键值存在
	Sort     int         `json:"sort"`      // 排序
	Icon     string      `json:"icon"`      // 展示图标
	Items    []*AuthTree `json:"items"`     // 子级数据项
}