package main

import (
	"github.com/Scorpio69t/csdn-swipe-views-go/router"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./template", ".html").Layout("layout.html"))
	router.Register(app)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
		return
	}
}
