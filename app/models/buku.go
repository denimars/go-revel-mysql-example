package models

import(
    "github.com/jinzhu/gorm"
)

type Buku struct{
    gorm.Model
    Nama string `gorm:"type:varchar(100)"`
    Deskripsi string `gorm:"type:varchar(255)"`
}