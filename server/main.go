package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"share_movie/server/config"
	"share_movie/server/controllers"
	"share_movie/server/middleware"
)

func main() {
	// 初始化数据库连接
	if err := config.InitDB(); err != nil {
		fmt.Printf("数据库初始化失败: %v\n", err)
		return
	}

	// 创建gin引擎
	r := gin.Default()

	// 配置CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API路由组
	api := r.Group("/api")
	{
		// 认证相关路由
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// 需要认证的路由组
		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			// 电影相关路由
			auth.GET("/movies", controllers.ListMovies)
			auth.GET("/movies/:id", controllers.GetMovie)
			auth.POST("/movies", controllers.CreateMovie)
			auth.PUT("/movies/:id", controllers.UpdateMovie)
			auth.DELETE("/movies/:id", controllers.DeleteMovie)

			// 管理员路由组
			admin := auth.Group("/admin")
			admin.Use(middleware.AdminAuthMiddleware())
			{
				// 管理员特权操作路由
				admin.GET("/users", controllers.GetUsers)
				admin.PUT("/users/:id/status", controllers.UpdateUserStatus)
			}
		}
	}

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}