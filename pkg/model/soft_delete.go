package model

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

//由于gorm提供的DeletedAt没有json tag， 因此自定义为yyyy-MM-dd hh:mm:ss格式
type DeletedAt struct {
	time.Time
}

//实现driver包中的Valuer接口，将自定义的类型值转换为驱动支持的Value类型值，ps: 接收器不能是指针类型，否则会报invalid memory address or nil pointer dereference
func (t DeletedAt) Value() (driver.Value, error) {
	var zeroTime time.Time
	//零值判断
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return t.Time, nil
}

// 实现sql包中的Scanner接口，被Rows和Row的Scan方法使用
func (t *DeletedAt) Scan(v interface{}) error {
	//把数据库中的时间赋值给t
	if v, ok := v.(time.Time); ok {
		*t = DeletedAt{Time: v}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 重写time.Time的MarshalJSON方法
func (t *DeletedAt) MarshalJSON() ([]byte, error) {
	timeStr := t.Format("2006-01-02 15:04:05")
	//零值判断
	if t.IsZero() {
		timeStr = ""
	}
	formatted := fmt.Sprintf("\"%s\"", timeStr)
	return []byte(formatted), nil
}

// 重写time.Time的UnmarshalJSON方法
func (t *DeletedAt) UnmarshalJSON(data []byte) error {
	//去除json格式中的双引号
	timeStr := strings.Trim(string(data), "\"")
	//空值判断，因为在MarshalJSON中时间是零值，json为空字符串
	if timeStr == "null" || strings.TrimSpace(timeStr) == "" {
		*t = DeletedAt{Time: time.Time{}}
		return nil
	}
	t.SetString(timeStr)
	return nil
}

// 设置字符串
func (t *DeletedAt) SetString(str string) *DeletedAt {
	if t != nil {
		// 指定解析的格式(设置转为本地格式)
		now, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
		if err == nil {
			*t = DeletedAt{Time: now}
		}
	}
	return t
}

/*
	以下直接从gorm源码的DeletedAt中copy的条款，实现gorm的DeletedAt一样的功能，主要是软删除能力
*/

func (DeletedAt) QueryClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{SoftDeleteQueryClause{Field: f}}
}

type SoftDeleteQueryClause struct {
	Field *schema.Field
}

func (sd SoftDeleteQueryClause) Name() string {
	return ""
}

func (sd SoftDeleteQueryClause) Build(clause.Builder) {
}

func (sd SoftDeleteQueryClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteQueryClause) ModifyStatement(stmt *gorm.Statement) {
	if _, ok := stmt.Clauses["soft_delete_enabled"]; !ok {
		if c, ok := stmt.Clauses["WHERE"]; ok {
			if where, ok := c.Expression.(clause.Where); ok && len(where.Exprs) > 1 {
				for _, expr := range where.Exprs {
					if orCond, ok := expr.(clause.OrConditions); ok && len(orCond.Exprs) == 1 {
						where.Exprs = []clause.Expression{clause.And(where.Exprs...)}
						c.Expression = where
						stmt.Clauses["WHERE"] = c
						break
					}
				}
			}
		}

		stmt.AddClause(clause.Where{Exprs: []clause.Expression{
			clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: sd.Field.DBName}, Value: nil},
		}})
		stmt.Clauses["soft_delete_enabled"] = clause.Clause{}
	}
}

func (DeletedAt) UpdateClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{SoftDeleteUpdateClause{Field: f}}
}

type SoftDeleteUpdateClause struct {
	Field *schema.Field
}

func (sd SoftDeleteUpdateClause) Name() string {
	return ""
}

func (sd SoftDeleteUpdateClause) Build(clause.Builder) {
}

func (sd SoftDeleteUpdateClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteUpdateClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.String() == "" {
		if _, ok := stmt.Clauses["WHERE"]; stmt.DB.AllowGlobalUpdate || ok {
			SoftDeleteQueryClause(sd).ModifyStatement(stmt)
		}
	}
}

func (DeletedAt) DeleteClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{SoftDeleteDeleteClause{Field: f}}
}

type SoftDeleteDeleteClause struct {
	Field *schema.Field
}

func (sd SoftDeleteDeleteClause) Name() string {
	return ""
}

func (sd SoftDeleteDeleteClause) Build(clause.Builder) {
}

func (sd SoftDeleteDeleteClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteDeleteClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.String() == "" {
		curTime := stmt.DB.NowFunc()
		stmt.AddClause(clause.Set{{Column: clause.Column{Name: sd.Field.DBName}, Value: curTime}})
		stmt.SetColumn(sd.Field.DBName, curTime, true)

		if stmt.Schema != nil {
			_, queryValues := schema.GetIdentityFieldValuesMap(stmt.ReflectValue, stmt.Schema.PrimaryFields)
			column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

			if len(values) > 0 {
				stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}

			if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
				_, queryValues = schema.GetIdentityFieldValuesMap(reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
				column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

				if len(values) > 0 {
					stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
				}
			}
		}

		if _, ok := stmt.Clauses["WHERE"]; !stmt.DB.AllowGlobalUpdate && !ok {
			stmt.DB.AddError(gorm.ErrMissingWhereClause)
		} else {
			SoftDeleteQueryClause(sd).ModifyStatement(stmt)
		}

		stmt.AddClauseIfNotExists(clause.Update{})
		stmt.Build("UPDATE", "SET", "WHERE")
	}
}
