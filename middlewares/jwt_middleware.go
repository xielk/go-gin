package middlewares

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/models"
	
	"gorm.io/gorm"
	"reflect"
	// "github.com/swaggo/gin-swagger"

)


func AuthMiddleware(db *gorm.DB, config config.Config) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(config.Jwt.Key),
		Timeout:     time.Hour * 720,
		MaxRefresh:  time.Hour * 720,
		IdentityKey: "username",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("PayloadFunc: Data ", data, reflect.TypeOf(data))
			if v, ok := data.(models.User); ok {
				fmt.Println("PayloadFunc: Type of data:", reflect.TypeOf(data))
				fmt.Println("PayloadFunc: Username:", v.Username)
				return jwt.MapClaims{
					"username": v.Username,
					"id":       v.ID,
				}
			}
			fmt.Println("PayloadFunc: Data is not *models.User")
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if username, ok := claims["username"].(string); ok {
				// 在这里获取ID值，如果claims中有ID字段的值的话
				var userID uint // 假设ID是uint类型的
				if id, ok := claims["id"].(float64); ok {
					userID = uint(id)
				}

				return &models.User{
					ID:       userID, // 设置ID值
					Username: username,
				}
			}
			return nil // 或者返回适当的错误值
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user models.User
			if err := c.BindJSON(&user); err != nil {
				fmt.Println("Authenticator: Error binding JSON:", err)
				return nil, jwt.ErrMissingLoginValues
			}

			// 在这里执行数据库查询，验证用户名和密码是否匹配
			// 假设您的数据库操作是通过一个名为DB的全局变量来执行的
			// 您需要根据您的项目结构和数据库逻辑来实现这部分逻辑
			var dbUser models.User
			if err := db.Where("user_name = ? AND pass = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
				fmt.Println("Authenticator: Error querying database:", err)
				return nil, jwt.ErrFailedAuthentication
			}

			fmt.Println("Authenticator: Database user found:", dbUser)
			return dbUser, nil
		},

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if token != "" {
				fmt.Println("LoginResponse: Token generated successfully")
				c.JSON(http.StatusOK, gin.H{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				})
			} else {
				fmt.Println("LoginResponse: Token is nil")
				c.JSON(http.StatusOK, gin.H{
					"token": "",
				})
			}
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println("Unauthorized: Code:", code, "Message:", message)
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	return authMiddleware, err
}
