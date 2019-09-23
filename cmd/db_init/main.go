package main

import (
	"log"
	// "fmt"

	"blog/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB init...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	// users := []model.User {
	// 	{
	// 		Username: "ray",
	// 		PasswordHash: model.GeneratePasswordHash("ray123"),
	// 		Email: "ray@126.com",
	// 		Avatar: fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("ray@126.com")),
	// 		Posts: []model.Post{
	// 			{Body: "Nice try, ray!"},
	// 			{Body: "This is my blog, written by ME!"},
	// 		},
	// 	},
	// 	{
	// 		Username: "jack",
	// 		PasswordHash: model.GeneratePasswordHash("jack123"),
	// 		Email: "jack@126.com",
	// 		Avatar: fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("jack@126.com")),
	// 		Posts: []model.Post{
	// 			{Body: "Wonderful!"},
	// 			{Body: "Sun shine is beautiful!"},
	// 		},
	// 	},
	// }

	// for _, u := range users {
	// 	db.Debug().Create(&u)
	// }

	model.AddUser("ray", "ray123", "ray@168.me")
	model.AddUser("jack", "jack123", "jack@168.me")

	u1, _ := model.GetUserByUsername("ray")
	u1.CreatePost("Nice,ray!")
	model.UpdateAboutMe(u1.Username, `Hellow there.`)
	model.UpdateAboutMe(u1.Username, `Can't wait to see you.`)

	u2, _ := model.GetUserByUsername("jack")
	u2.CreatePost("Great job, jack")
	model.UpdateAboutMe(u2.Username, `That movie was so cool.`)

	u1.Follow(u2.Username)
}