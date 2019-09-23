package main

import (
	"net/http"
	// "html/template"
	// "io/ioutil"
	// "os"
	"blog/model"
	"blog/controller"
	
	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()
	
	// http.ListenAndServe(":8888", nil)
	http.ListenAndServe(":8880", context.ClearHandler(http.DefaultServeMux))
}