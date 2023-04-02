package dependency

import "gorm.io/gorm"

type Databases struct {
	Main *gorm.DB
}