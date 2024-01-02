package dbs

type SDNAka struct {
	UID       int64 `gorm:"primaryKey"`
	Type      string
	Category  string
	FirstName string
	LastName  string
}
