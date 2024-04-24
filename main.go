package main

import (
	"fmt"
	"go-gin/config"
	"go-gin/controllers"
	"log"
	"net/http"
	"os"
	// "time"
	// "reflect"
	// "go-gin/models"

	// jwt "github.com/dgrijalva/jwt-go"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-gin/middlewares"
	"go-gin/docs"

	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

// var jwtKey []byte

type login struct {
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	

	// 设置环境变量
	env := os.Getenv("ENV")

	// 读取配置文件
	var configFile string
	if env == "prod" {
		configFile = "configs/prod.yaml"
	} else {
		configFile = "configs/debug.yaml"
	}

	var config config.Config
	if err := readConfig(configFile, &config); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + fmt.Sprintf(":%d", config.Database.Port) + ")/" + config.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Println(dsn)
	log.Println(config)
	// jwtKey = []byte(config.Jwt.Key)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建控制器实例
	articleController := &controllers.ArticleController{
		DB: db,
	}

	// 初始化 Gin 引擎
	router := gin.Default()

	// 定义路由
	router.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server":   config.Server,
			"database": config.Database,
		})
	})

	userController := &controllers.UserController{}
	testController := &controllers.TestController{}

	// 用户登录路由
	router.GET("/user", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	router.GET("/article", articleController.GetArticle)
	router.GET("/articles", articleController.GetArticles)
	router.POST("/article", articleController.CreateArticle)
	router.PUT("/article/:id", articleController.UpdateArticle)
	router.DELETE("/article/:id", articleController.DeleteArticle)
	router.GET("/test", testController.Test)
	
	authMiddleware, err := middlewares.AuthMiddleware(db, config)
    if err != nil {
        log.Fatalf("Error creating AuthMiddleware: %v", err)
    }

	// 添加JWT认证中间件到路由
	router.POST("/login", authMiddleware.LoginHandler)

	// 要求身份验证的路由组
	authGroup := router.Group("/auth")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		authGroup.GET("/profile", func(c *gin.Context) {
			// 从上下文中获取用户信息
			 claims := jwt.ExtractClaims(c)
        // 将claims打印出来
        	c.JSON(http.StatusOK, gin.H{"claims": claims})
			// c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "username": username})
		})
	}
	docs.SwaggerInfo.Title = "go-gin"
    docs.SwaggerInfo.Description = "by xielk"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.BasePath = "/"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 启动服务器
	port := fmt.Sprintf(":%d", config.Server.Port)
	router.Run(port)
}


// readConfig 从配置文件中读取配置信息
func readConfig(configFile string, config interface{}) error {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}

