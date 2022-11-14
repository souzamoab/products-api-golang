package migrations

import (
	"github.com/souzamoab/products-api-golang/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Product{})
}
