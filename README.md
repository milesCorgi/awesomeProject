# awesomeProject
实际上是用来复健写go的,根目录底下的sql是用于调试的数据脚本（因为是重构所以现有表结构暂时不动了）

# 项目构建命令

`go build -tags netgo -ldflags '-s -w' -o ./target/app`  
(发现直接戳goland的构建也能出包)

执行 ./target/app就能运行

# 选型

*web框架* ： gin，比beego精简。是有中文文档的：https://gin-gonic.com/zh-cn/docs/  

当然也有很多这样的中文教程：

https://eddycjy.gitbook.io/golang/di-3-ke-gin/api-01  

https://zhuanlan.zhihu.com/p/614838910

gin 路由获取get post请求参数 : https://zhuanlan.zhihu.com/p/474758711

*ORM框架* ： gorm，这是一个使用Go语言编写的ORM框架。它文档齐全，对开发者友好，支持主流数据库。  
有中文文档：https://gorm.io/zh_CN/docs/  

*数据库分表*： Sharding
中文说明https://gorm.io/zh_CN/docs/sharding.html

Go-validator ： 代码层面的一个校验脚手架

*日志打印*：zap

https://segmentfault.com/a/1190000022461706
https://www.cnblogs.com/you-men/p/14694928.html#_label4

*定时任务*：

我们熟悉的xxl-job

https://github.com/xxl-job/xxl-job-executor-go  

cron ： 服务自己的一些定时事务，用这个也行（搭xxl-job也是要钱的）

https://darjun.github.io/2020/06/25/godailylib/cron/

*单元测试*：
还不熟，考虑以下两个

gomock  https://github.com/golang/mock

Mockey https://github.com/bytedance/mockey

↑这个还有点像mockito

*本地缓存*：用这个方案造了个本地缓存（因为买redis要钱）（实际改用了读写锁）

https://www.jianshu.com/p/f95fb8f0f737

读写锁的科普

https://www.jianshu.com/p/679041bdaa39

选锁总结：
- Mutex（互斥锁）适用于读写不确定，并且只有一个读或者写的场景
- RWMutex（读写锁）适用于读多写少的场景（业务代码应该用这个会更多）

# TodoList

## 第一期

- 重写核心功能如下
  - 添加片段记录，后续成为Note（某个线上视频，或者其他线上内容值得注意的点）--- done
  - 检索现有已经添加的片段记录  --- done，测试中
  - youtube的爬虫（因为需要网络梯子支持，部署方法待定）（研究用用ssr？）
  - b站的爬虫（直接用定时任务写在后端 --- > 考虑用上xxl-job）
  - youtube的爬虫视频信息 --- done
  - 查询B站的爬虫视频信息 --- done


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