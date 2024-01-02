package dto

type UserName struct {
	RecordKey string `json:"-"`
	UID       int64  `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
