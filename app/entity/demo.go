package entity

type Demo struct {
	ID int64 `json:"id"`
}

func (d Demo) TableName() string {
	return "demo"
}
