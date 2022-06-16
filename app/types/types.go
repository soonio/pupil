package types

type AdminDictSaveRequest struct {
	K string `json:"k" form:"k" validate:"required"` // 键
	V string `json:"v" form:"v" validate:"required"` // 值
}
