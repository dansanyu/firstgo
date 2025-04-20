package router

import (
	"firstgo/internal/dao"
	"firstgo/internal/handlers"
	"firstgo/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {

}

// SetupRouter 初始化路由
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 全局中间件（如日志、CORS）
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	userDao := dao.NewDao(db)
	UserService := service.NewService(userDao)
	handler := handlers.NewUserHandler(UserService)
	registerUserRoutes(r, handler)
	// 404 处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Not Found"})
	})

	return r
}

// 注册用户路由
func registerUserRoutes(r *gin.Engine, handler *handlers.UserHandler) {
	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/create", handler.CreateUser)
	}
}
