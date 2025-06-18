package api

import (
	"go-task4/api/auth"
	"go-task4/api/comment"
	"go-task4/api/post"
	"go-task4/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// 初始化依赖
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	postRepo := post.NewPostRepository(db)
	postService := post.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postService)

	commentRepo := comment.NewCommentRepository(db)
	commentService := comment.NewCommentService(commentRepo, postRepo)
	commentHandler := comment.NewCommentHandler(commentService)

	// 认证路由
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// 文章路由
	postGroup := router.Group("/posts")
	{
		postGroup.GET("", postHandler.GetAllPosts)
		postGroup.GET("/:id", postHandler.GetPost)

		// 需要认证的路由
		authPostGroup := postGroup.Group("").Use(middleware.AuthMiddleware())
		{
			authPostGroup.POST("", postHandler.CreatePost)
			authPostGroup.PUT("/:id", postHandler.UpdatePost)
			authPostGroup.DELETE("/:id", postHandler.DeletePost)
		}
	}

	// 评论路由
	commentGroup := router.Group("/comments")
	{
		commentGroup.GET("/post/:postId", commentHandler.GetCommentsByPostID)

		// 需要认证的路由
		authCommentGroup := commentGroup.Group("").Use(middleware.AuthMiddleware())
		{
			authCommentGroup.POST("", commentHandler.CreateComment)
			authCommentGroup.DELETE("/:id", commentHandler.DeleteComment)
		}
	}
}
