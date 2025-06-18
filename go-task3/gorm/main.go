package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func main() {

	db, err := ConnectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{}) // 自动迁移模型
	// AddUser(db, "Alice")                                                   // 添加用户
	// AddUser(db, "Bob")                                                     // 添加用户
	// AddPost(db, "First Post", "This is the content of the first post.", 1) // 添加文章
	// AddComment(db, "This is a comment.", 1, 1)                             // 添加评论
	fmt.Println(GetUserPostsWithComments(db, 1)) // 查询用户发布的所有文章及其对应的评论信息
	fmt.Println(GetMostCommentedPost(db))        // 查询评论数量最多的文章信息
	//添加评论
	//AddComment(db, "This is another comment.", 1, 1) // 添加评论
	//添加评论
	//AddComment(db, "This is a comment by Bob.", 1, 2) // 添加评论

	// 删除评论
	DeleteComment(db, 2) // 假设删除ID为1的评论

}
func ConnectDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}

// 添加用户
func AddUser(db *gorm.DB, username string) error {
	user := User{Username: username, PostNum: 0}
	return db.Create(&user).Error
}

// 添加文章
func AddPost(db *gorm.DB, title, content string, userID uint) error {
	post := Post{Title: title, Content: content, Status: "草稿", UserID: userID}
	return db.Create(&post).Error
}

// 添加评论
func AddComment(db *gorm.DB, content string, postID, userID uint) error {
	comment := Comment{Content: content, PostID: postID, UserID: userID}
	return db.Create(&comment).Error
}

// 删除评论
func DeleteComment(db *gorm.DB, commentID uint) error {
	var comment Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		return err // 如果评论不存在，返回错误
	}
	return db.Delete(&comment).Error // 删除评论
}

// 题目1：模型定义
type User struct {
	ID       uint   `gorm:"primaryKey"` // 用户ID
	Username string // 用户名
	//Email    string // 邮箱
	PostNum int    // 文章数量统计字段
	Posts   []Post `gorm:"foreignKey:UserID"` // 用户发布的文章
}
type Post struct {
	ID            uint      `gorm:"primaryKey"`        // 文章ID
	Title         string    `gorm:"not null"`          // 文章标题
	Content       string    `gorm:"type:text"`         // 文章内容
	Status        string    `gorm:"not null"`          // 文章状态，默认为草稿
	UserID        uint      `gorm:"not null"`          // 外键，关联到 User
	CommentsCount int64     `gorm:"default:0"`         // 评论数量S
	Comments      []Comment `gorm:"foreignKey:PostID"` // 文章的评论
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
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func GetMostCommentedPost(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Preload("Comments").Order("comments_count DESC").Limit(1).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).First(&user).Error; err != nil {
		return err
	}
	user.PostNum++
	return tx.Save(&user).Error
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).First(&post).Error; err != nil {
		return err
	}
	post.CommentsCount--
	if post.CommentsCount == 0 {
		post.Status = "无评论"
	}
	return tx.Save(&post).Error
}
