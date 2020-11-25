// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"database/sql"
	"focus/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

// RoleAuthDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type RoleAuthDao struct {
	gmvc.M
	Table   string
	Columns roleAuthColumns
}

// RoleAuthColumns defines and stores column names for table gf_role_auth.
type roleAuthColumns struct {
	Id         string // 自增ID    
    RoleId     string // 角色ID    
    AuthId     string // 权限ID    
    CreatedAt  string // 创建时间
}

var (
	// RoleAuth is globally public accessible object for table gf_role_auth operations.
	RoleAuth = &RoleAuthDao{
		M:     g.DB("default").Table("gf_role_auth").Safe(),
		Table: "gf_role_auth",
		Columns: roleAuthColumns{
			Id:        "id",          
            RoleId:    "role_id",     
            AuthId:    "auth_id",     
            CreatedAt: "created_at",
		},
	}
)

// As sets an alias name for current table.
func (d *RoleAuthDao) As(as string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *RoleAuthDao) TX(tx *gdb.TX) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *RoleAuthDao) Master() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *RoleAuthDao) Slave() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *RoleAuthDao) Args(args ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Args(args ...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *RoleAuthDao) LeftJoin(table ...string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *RoleAuthDao) RightJoin(table ...string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *RoleAuthDao) InnerJoin(table ...string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *RoleAuthDao) Fields(fieldNamesOrMapStruct ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *RoleAuthDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *RoleAuthDao) Option(option int) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *RoleAuthDao) OmitEmpty() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *RoleAuthDao) Filter() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *RoleAuthDao) Where(where interface{}, args ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *RoleAuthDao) WherePri(where interface{}, args ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *RoleAuthDao) And(where interface{}, args ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *RoleAuthDao) Or(where interface{}, args ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *RoleAuthDao) Group(groupBy string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *RoleAuthDao) Order(orderBy ...string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *RoleAuthDao) Limit(limit ...int) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *RoleAuthDao) Offset(offset int) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *RoleAuthDao) Page(page, limit int) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *RoleAuthDao) Batch(batch int) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *RoleAuthDao) Cache(duration time.Duration, name ...string) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *RoleAuthDao) Data(data ...interface{}) *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.RoleAuth.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *RoleAuthDao) All(where ...interface{}) ([]*model.RoleAuth, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.RoleAuth
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.RoleAuth.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *RoleAuthDao) One(where ...interface{}) (*model.RoleAuth, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.RoleAuth
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *RoleAuthDao) FindOne(where ...interface{}) (*model.RoleAuth, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.RoleAuth
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *RoleAuthDao) FindAll(where ...interface{}) ([]*model.RoleAuth, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.RoleAuth
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (d *RoleAuthDao) Chunk(limit int, callback func(entities []*model.RoleAuth, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.RoleAuth
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *RoleAuthDao) LockUpdate() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *RoleAuthDao) LockShared() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *RoleAuthDao) Unscoped() *RoleAuthDao {
	return &RoleAuthDao{M:d.M.Unscoped()}
}