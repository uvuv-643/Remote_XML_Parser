package dbs

type SDNNationality struct {
	UID       int64 `gorm:"primaryKey"`
	Country   string
	MainEntry bool
}
