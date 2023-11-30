package model

type MaxLoan struct {
	ID      int64 `json:"id" gorm:"primaryKey" swagger:"description(ID)"`
	UserID  int64 `json:"user_id" swagger:"description(UserID)"`
	Limit   int64 `json:"limit" swagger:"description(Limit)"`
	Current int64 `json:"current" swagger:"description(Current)"`
}
