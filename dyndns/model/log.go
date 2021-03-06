package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	Status    bool
	Message   string
	Host      Host
	HostID    uint
	SentIP    string
	CallerIP  string
	TimeStamp time.Time
	UserAgent string
}
