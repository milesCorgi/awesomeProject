# awesomeProject
实际上是用来复健写go的

# 项目构建命令

`go build -tags netgo -ldflags '-s -w' -o ./target/app`  
(发现直接戳goland的构建也能出包)

执行 ./target/app就能运行

# 选型

gin ： web框架，比beego精简。是有中文文档的：https://gin-gonic.com/zh-cn/docs/  

当然也有很多这样的中文教程：

https://eddycjy.gitbook.io/golang/di-3-ke-gin/api-01  

https://zhuanlan.zhihu.com/p/614838910

Gorm ： gorm是一个使用Go语言编写的ORM框架。它文档齐全，对开发者友好，支持主流数据库。  
有中文文档：https://gorm.io/zh_CN/docs/  

Go-validator ： 代码层面的一个校验脚手架

日志打印：zap

https://segmentfault.com/a/1190000022461706

定时任务：https://github.com/xxl-job/xxl-job-executor-go  我们熟悉的xxl-job

单元测试：
还不熟，考虑以下两个

gomock  https://github.com/golang/mock

Mockey https://github.com/bytedance/mockey

↑这个还有点像mockito

# TodoList

## 第一期

- 重写核心功能如下
  - 添加片段记录，后续成为Note（某个线上视频，或者其他线上内容值得注意的点）
  - 检索现有已经添加的片段记录
  - youtube的爬虫（因为需要网络梯子支持，部署方法待定）（研究用用ssr？）
  - b站的爬虫（直接用定时任务写在后端）
  - youtube的爬虫视频信息
  - 查询B站的爬虫视频信息


```	
加入片段记录

beego.Router("/api/add_video_note_info", &controllers.AddController{}, "post:AddTwoSetVideoInfo")

查询一个已登录的片段信息

beego.Router("/api/query_video_info", &controllers.QueryController{}, "post:ShowVideoInfo")

查询youtube的爬虫视频信息

beego.Router("/api/query_youtube_video_list", &controllers.QueryController{}, "post:ShowYoutubeVideoList")

查询B站的爬虫视频信息

beego.Router("/api/query_bilibili_video_list", &controllers.QueryController{}, "post:ShowBiliBiliVideoList")
```