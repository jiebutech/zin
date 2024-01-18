package dto

type IdReq struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}
