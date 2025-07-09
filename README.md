# 🌀 csdn-swipe-views-go

一个使用 **Go + Iris 框架** 开发的前后端一体化 Web 项目，支持用户输入 CSDN 用户 ID，获取其博客文章列表并自动模拟访问文章页面以刷取访问量。

---

## ✨ 项目功能

- 🔍 用户输入 CSDN ID，自动抓取其所有公开文章
- 📄 展示文章列表：标题 / 链接 / 发布时间 / 当前浏览量
- 🧠 模拟访问行为：自定义刷量次数、并发度、时间间隔
- 🧰 支持代理池和 User-Agent 池伪装请求
- 📊 实时展示刷量进度、成功数与失败数
- 📝 提供简单的访问日志记录

---

## 📁 项目结构
```bash
csdn-swipe-views-go/
├── main.go // 入口
├── routes/ // 路由注册
│ └── router.go
├── controller/ // 控制器
│ └── csdn.go
├── service/ // 核心业务逻辑
│ ├── csdn_parser.go
│ └── pageview.go
├── model/ // 数据结构定义
│ └── article.go
├── templates/ // 页面模板
│ ├── index.html
│ ├── result.html
│ └── logs.html
├── static/ // 静态资源
│ ├── css/
│ └── js/
├── utils/ // 工具库，如代理池等
│ └── proxy.go
└── go.mod / go.sum // 依赖管理
```


---

## 🚀 快速开始

### ✅ 环境要求

- Go 1.19+
- 已安装 [Iris](https://github.com/kataras/iris)（自动通过 `go get` 安装）
- 推荐使用 `proxy pool` 服务或免费代理接口（可选）

### 📦 安装依赖

```bash
git clone https://github.com/yourname/csdn-swipe-views-go.git
cd csdn-swipe-views-go

go mod tidy
```

### 🏃 运行项目

```bash
go run main.go
```
默认监听端口：http://localhost:8080

## 🧱 主要依赖
- [Iris](https://github.com/kataras/iris) - 高性能 Go Web 框架
- [Goquery](https://github.com/PuerkitoBio/goquery) - 类似 jQuery 的 Go HTML 解析库
- [Logrus](https://github.com/sirupsen/logrus) - 强大的日志库
- [Colly](https://github.com/gocolly/colly) - 高效的 Go 爬虫框架
- [goroutines + channels] - 并发控制与任务调度

## 🛠 TODO（开发中）

-[X] 项目结构搭建
-[ ] 首页输入 + 文章列表展示
-[ ] 爬取 CSDN 用户文章列表
-[ ] 模拟访问请求发送
-[ ] 日志页面实现
-[ ] 引入代理池 / User-Agent 池

## 📄 License
本项目仅供学习用途，请勿用于违反 CSDN 使用协议的行为。

