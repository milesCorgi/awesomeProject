# awesomeProject
实际上是用来复健写go的

# 项目构建命令

go build -tags netgo -ldflags '-s -w' -o ./target/app(发现直接戳goland的构建也能出包)

执行 app就能运行

# 选型

gin ： web框架，比beego精简。是有中文文档的：https://gin-gonic.com/zh-cn/docs/   当然也有很多这样的中文教程：https://eddycjy.gitbook.io/golang/di-3-ke-gin/api-01
Gorm ： gorm是一个使用Go语言编写的ORM框架。它文档齐全，对开发者友好，支持主流数据库。有中文文档：https://gorm.io/zh_CN/docs/
Go-validator ： 代码层面的一个校验脚手架