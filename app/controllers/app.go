package controllers

import (
	"perpus/app/models"
	"github.com/revel/revel"
	"fmt"
)

type App struct {
	ModelController
}

func (c App) Index() revel.Result {
	var bukus []models.Buku
	buku := c.Orm.Find(&bukus)
	if buku.Error != nil{
		panic(buku.Error)
	}
	fmt.Println(bukus)
	return c.Render(bukus)
}

func (c App) New() revel.Result{
	return c.Render()
}

func (c App) Simpan(nama string, deskripsi string) revel.Result{
	buku := models.Buku{Nama:nama, Deskripsi:deskripsi}
	c.Orm.Save(&buku)
	return c.Redirect(App.Index)
}

func (c App) Show(id int) revel.Result{
	var bukus models.Buku
	buku := c.Orm.First(&bukus, id)
	if buku.Error != nil{
		panic(buku.Error)
	}
	return c.Render(bukus)
}

func (c App) Edit(id int) revel.Result{
	var bukus models.Buku
	buku := c.Orm.First(&bukus, id)
	if buku.Error != nil{
		panic(buku.Error)
	}
	return c.Render(bukus)
}


func (c App) Update(id int, nama string, deskripsi string) revel.Result{
	bukus := models.Buku{}
	c.Orm.Model(&bukus).Where("ID=?", id).Updates(map[string]interface{}{"Nama":nama, "Deskripsi":deskripsi})
	return c.Redirect(App.Index)
}

func (c App) Delete(id int) revel.Result{
	buku := models.Buku{}
	c.Orm.Where("ID=?", id).Delete(&buku)
	return c.Redirect(App.Index)
}