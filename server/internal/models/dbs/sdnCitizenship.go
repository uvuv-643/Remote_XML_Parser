package dbs

type SDNCitizenship struct {
	UID       int64 `gorm:"primaryKey"`
	Country   string
	MainEntry bool
}
