package dbs

type SDNPlaceOfBirth struct {
	UID       int64 `gorm:"primaryKey"`
	Place     string
	MainEntry bool
}
