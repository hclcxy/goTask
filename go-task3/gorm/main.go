package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Print(getMostCommentedPost())

}

type User struct {
	ID       uint   `gorm:"primaryKey"`        // 用户ID
	Username string `gorm:"unique;not null"`   // 用户名
	Email    string `gorm:"unique;not null"`   // 邮箱
	Posts    []Post `gorm:"foreignKey:UserID"` // 用户发布的文章
}
type Post struct {
	ID       uint      `gorm:"primaryKey"`        // 文章ID
	Title    string    `gorm:"not null"`          // 文章标题
	Content  string    `gorm:"type:text"`         // 文章内容
	UserID   uint      `gorm:"not null"`          // 外键，关联到 User
	Comments []Comment `gorm:"foreignKey:PostID"` // 文章的评论
}
type Comment struct {
	ID      uint   `gorm:"primaryKey"`        // 评论ID
	Content string `gorm:"type:text"`         // 评论内容
	PostID  uint   `gorm:"not null"`          // 外键，关联到 Post
	UserID  uint   `gorm:"not null"`          // 外键，关联到 User
	User    User   `gorm:"foreignKey:UserID"` // 评论的用户
	Post    Post   `gorm:"foreignKey:PostID"` // 评论的文章
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func getMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&post).Error
	if err != nil {
		return Post{}, err
	}
	return post, nil
}
