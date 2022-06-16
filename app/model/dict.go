package model

// Dict 字典表
type Dict struct {
	K string `gorm:"column:k;primary_key"` // 键
	V string `gorm:"column:v;NOT NULL"`    // 值
}

func (m *Dict) TableName() string {
	return "dict"
}
