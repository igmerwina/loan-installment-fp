package model

type MaxLoan struct {
	ID      int64 `json:"id" gorm:"primaryKey"`
	UserID  int64 `json:"user_id"`
	Limit   int64 `json:"limit"`
	Current int64 `json:"current"`
}
