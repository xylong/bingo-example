package dto

// CountList 带统计的列表
type CountList struct {
	Total int64         `json:"total"`
	List  []interface{} `json:"list"`
}
