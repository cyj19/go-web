package util

import "github.com/cyj19/go-web/internal/pkg/model"

// GenWhereOrderByStruct 根据结构体生成sql条件
func GenWhereOrderByStruct(value interface{}) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)
	switch val := value.(type) {
	case model.SysUser:
		if val.Id > 0 {
			v := val.Id
			whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{v}})
		}
		if val.Username != "" {
			v := "%" + val.Username + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "username like ?", Value: []interface{}{v}})
		}
		if val.Status != nil {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*val.Status}})
		}
	case model.SysRole:
		if val.Name != "" {
			v := "%" + val.Name + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
		}
		if val.NameZh != "" {
			v := "%" + val.NameZh + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "name_zh like ?", Value: []interface{}{v}})
		}
		if val.Status != nil {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*val.Status}})
		}
		if val.Sort != nil {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "sort = ?", Value: []interface{}{*val.Sort}})
		}
	case model.SysMenu:
		if val.Name != "" {
			v := "%" + val.Name + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
		}
		if val.Title != "" {
			v := "%" + val.Title + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "title = ?", Value: []interface{}{v}})
		}
		if val.Status != nil {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*val.Status}})
		}
	case model.SysApi:
		if val.Id > 0 {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{val.Id}})
		}
		if val.Method != "" {
			v := "%" + val.Method + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "method like ?", Value: []interface{}{v}})
		}
		if val.Path != "" {
			v := "%" + val.Path + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "path like ?", Value: []interface{}{v}})
		}
		if val.Category != "" {
			v := "%" + val.Category + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "category like ?", Value: []interface{}{v}})
		}
		if val.Creator != "" {
			v := "%" + val.Creator + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "creator like ?", Value: []interface{}{v}})
		}
	}

	return whereOrders
}
