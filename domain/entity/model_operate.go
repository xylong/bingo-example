package entity

import (
	"gorm.io/gorm"
)

const (
	_            = iota
	Equal        // =
	NotEqual     // <>
	GreaterThan  // >
	GreaterEqual // >=
	LessThan     // <
	LessEqual    // <=
	In           // in
	NotIn        // not in
	Like         // like
	NotLike      // not like
)

// Scope 查询范围
type Scope func(db *gorm.DB) *gorm.DB

// Paginate 分页
func Paginate(page, pageSize int) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((getPage(page) - 1) * getLimit(pageSize)).Limit(getLimit(pageSize))
	}
}

func getPage(page int) int {
	if page <= 0 {
		return 1
	}

	return page
}

func getLimit(pageSize int) int {
	switch {
	case pageSize <= 0:
		return 10
	case pageSize >= 100:
		return 100
	default:
		return pageSize
	}
}

func Select(query interface{}, arg ...interface{}) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, arg...)
	}
}

func Order(order interface{}) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// With 关联预加载
func With(relations map[string][]string) func(*gorm.DB) *gorm.DB {
	if relations == nil {
		return nil
	}

	length := len(relations)
	if length == 0 {
		return nil
	}

	return func(db *gorm.DB) *gorm.DB {
		for key, value := range relations {
			if len(value) == 0 {
				continue
			}

			if key == "" {
				db.Select(value)
			} else {
				func(column []string) *gorm.DB {
					return db.Preload(key, func(db2 *gorm.DB) *gorm.DB {
						return db2.Select(column)
					})
				}(value)
			}
		}

		return db
	}
}
