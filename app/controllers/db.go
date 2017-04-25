package controllers

import (
    "perpus/app/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    r "github.com/revel/revel"
    "fmt"
)

var DB *gorm.DB

func init(){
    var err error
    DB, err = gorm.Open("mysql", "root:root@/perpus?charset=utf8&parseTime=True&loc=Local")
    if err != nil{
        panic(fmt.Sprintf("gagal terkoneksi kedatabase, %v", err))
    }else{
        fmt.Println("Koneksi berhasil")
    }
}

type ModelController struct{
    *r.Controller
    Orm *gorm.DB
}

func (c *ModelController) Begin() r.Result{
    c.Orm = DB
    return nil
}

func (c *ModelController) CreateTable() r.Result{
    c.Orm.CreateTable(models.Buku{})
    return nil
}