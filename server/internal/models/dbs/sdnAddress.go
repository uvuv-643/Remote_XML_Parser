package dbs

type SDNAddress struct {
	UID             int64 `gorm:"primaryKey"`
	City            string
	Address1        string
	Address2        string
	Address3        string
	StateOrProvince string
	PostalCode      string
	Country         string
}
