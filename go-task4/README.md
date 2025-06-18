个人博客系统后端 - Go + Gin + GORM
https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go
https://img.shields.io/badge/Gin-1.9.0-00ADD8?logo=go
https://img.shields.io/badge/GORM-1.25.0-00ADD8
https://img.shields.io/badge/MySQL-8.0+-4479A1?logo=mysql

这是一个基于Go语言、Gin框架和GORM的个人博客系统后端，提供完整的用户认证、文章管理和评论功能。

功能特性
✅ 用户认证系统

用户注册与登录

JWT令牌认证

密码加密存储

📝 文章管理

创建、读取、更新、删除文章(CRUD)

文章列表分页

作者权限控制

💬 评论系统

文章评论功能

评论权限管理

🛡️ 安全特性

输入验证

错误处理

请求日志

技术栈
编程语言: Go 1.20+

Web框架: Gin

ORM: GORM

数据库: MySQL 8.0+

认证: JWT


快速开始
前提条件
Go 1.20+

MySQL 8.0+


配置数据库

创建MySQL数据库并修改配置文件 config/database.go:

go
dsn := "your_username:your_password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
安装依赖

bash
go mod tidy
运行迁移

bash
go run main.go migrate
启动服务器

bash
go run main.go
服务器将默认运行在 http://127.0.0.1:9000

API文档
完整的API文档可在 API文档 查看，或通过以下方式访问：



项目结构

ge-task4/
├── api/                # API路由和处理程序
│   ├── auth/           # 认证相关
│   ├── post/           # 文章相关
│   ├── comment/        # 评论相关
│   └── routes.go       # 路由定义
├── config/             # 配置
├── middleware/         # 中间件
├── models/             # 数据模型
├── pkg/                # 公共包
│   ├── utils/          # 工具函数
│   └── logger/         # 日志配置
├── go.mod
├── go.sum
└── main.go             # 主入口
开发指南
环境变量

bash
go run main.go
生产构建

bash
go build -o go-task4
./go-task4
测试
运行单元测试:

bash
go test ./...
贡献指南
欢迎贡献！请遵循以下步骤：

Fork项目

创建特性分支 (git checkout -b feature/AmazingFeature)

提交更改 (git commit -m 'Add some AmazingFeature')

推送到分支 (git push origin feature/AmazingFeature)

打开Pull Request

许可证
本项目采用 MIT 许可证 - 详情请见 LICENSE 文件

联系方式
如有任何问题，请联系：

作者: insthu

邮箱: insthu@163.com

项目链接: https://github.com/hclcxy/goTask/tree/main/go-task4