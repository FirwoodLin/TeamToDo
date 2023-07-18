package utils

import (
	"TeamToDo/global"
	"TeamToDo/model/response"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(id string) string {
	// 这里jwtKey不会为空
	//jwtKey := viper.GetString("JWT.SecretKey")
	jwtKey := global.Server.JWT.SecretKey
	//expireTime := global.Server.JWT.ExpireTime // TODO：待完善：添加过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": id,
		"exp":    time.Now().Add(time.Second * time.Duration(global.Server.JWT.ExpireTime)).Unix(),
	})
	// 这里不会有error被返回
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Println("[error]:", err)
		return ""
	}
	return tokenString
}

func MiddlewareJWTAuthorize() gin.HandlerFunc {
	jwtKey := global.Server.JWT.SecretKey

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) <= len("Bearer ") {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		tokenString := auth[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		id, err := strconv.ParseUint((*claims)["userID"].(string), 10, 32)
		if err != nil {
			log.Println("[error]:", err)
			c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("校验令牌失败"))
			c.Abort()
			return
		}
		c.Set("userID", uint(id))
	}
}
