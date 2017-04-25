package controllers

import "github.com/revel/revel"

func init(){
    revel.InterceptMethod((*ModelController).Begin, revel.BEFORE)
    revel.InterceptMethod((*ModelController).CreateTable, revel.BEFORE)

}