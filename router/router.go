package router

import "github.com/kataras/iris/v12"

func Register(app *iris.Application) {
	// 首页：输入 CSDN ID
	//app.Get("/", controller.Index)

	// 查询用户文章列表
	//app.Post("/articles", controller.FetchArticles)

	// 启动刷访问量任务
	//app.Post("/swipe", controller.StartSwipe)

	// 查询访问日志
	//app.Get("/logs", controller.ShowLogs)

	// 静态资源映射
	app.HandleDir("/static", "./static")
}
