package global

import "time"

type UpdateStatusName string

const (
	Empty    UpdateStatusName = "empty"
	Updating UpdateStatusName = "updating"
	Ok       UpdateStatusName = "ok"
)

type UpdateStatus struct {
	ID        int64 `gorm:"primary_key"`
	Status    UpdateStatusName
	CreatedAt time.Time
	UpdatedAt time.Time
}
