package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Gorm telah menyediakan hooks yang dapat kita gunakan,
// salah satunya adalah method BeforeCreate.
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// Method BeforeCreate akan tereksekusi sebelum melakukan create data.
	// Disini penulis akan mengimplementasikan BeforeCreate untuk Product,
	// ketika ingin membuat data product maka
	// method BeforeCreate tersebut akan terkesekusi
	fmt.Println("Product before Create()")

	// Ketika function createProduct dijalankan, maka akan menghasilkan error
	// karena panjang karakter dari property Name
	// dari struct Product kurang dari 4
	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}
