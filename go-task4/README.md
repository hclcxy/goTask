# 个人博客系统后端

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-1.9.0-00ADD8)
![GORM](https://img.shields.io/badge/GORM-1.25.0-00ADD8)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?logo=mysql)
![License](https://img.shields.io/badge/License-MIT-blue)

基于Go语言和Gin框架开发的个人博客系统后端，提供完整的用户认证、文章管理和评论功能。

## 功能特性

### 用户认证
- ✅ 用户注册/登录
- 🔐 JWT认证
- 🔒 Bcrypt密码加密

### 文章管理
- ✏️ 文章CRUD操作
- 📄 文章列表分页
- 🔐 作者权限控制

### 评论系统
- 💬 文章评论功能
- 🚫 评论删除权限控制

### 其他
- 📊 请求日志记录
- 🛡️ 输入验证
- ⚠️ 统一错误处理

## 技术栈

- **后端框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **日志**: Zerolog

## 快速开始

### 安装要求

- Go 1.20+
- MySQL 8.0+

### 安装步骤

1. 克隆仓库
```bash
git clone   https://github.com/hclcxy/goTask.git
cd blog-system
配置数据库

bash
cp .env.example .env
# 编辑.env文件配置数据库连接
安装依赖

bash
go mod tidy
数据库迁移

bash
go run main.go migrate
启动服务

bash
go run main.go
服务将运行在 http://127.0.0.1:9000

API文档
查看完整API文档



项目结构
text
blog-system/
├── api/                # API路由和处理
│   ├── auth/           # 认证相关
│   ├── post/           # 文章管理
│   ├── comment/        # 评论系统
│   └── routes.go       # 路由定义
├── config/             # 配置
├── middleware/         # 中间件
├── models/             # 数据模型
├── pkg/                # 公共包
│   ├── utils/          # 工具函数
│   └── logger/         # 日志配置
├── go.mod
├── go.sum
└── main.go             # 程序入口
开发指南
环境变量
复制.env.example为.env并配置：

ini
DB_DSN=username:password@tcp(localhost:3306)/blog_system
JWT_SECRET=your_jwt_secret_key
常用命令
bash
# 开发模式
go run main.go

# 生产构建
go build -o blog-system

# 运行测试
go test ./...
贡献指南
欢迎贡献！请遵循以下流程：

联系方式
如有任何问题，请联系：

作者: insthu

邮箱: insthu@163.com

项目链接: https://github.com/hclcxy/goTask/tree/main/go-task4




