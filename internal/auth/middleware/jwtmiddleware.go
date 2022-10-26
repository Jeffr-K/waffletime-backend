package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// https://sosedoff.com/2014/12/21/gin-middleware.html
// TODO: 미들웨어 사용법, 레디스에 토큰
func TokenValidationMiddleware(c *gin.Context) {
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed."})
		c.Abort()
		return
	}

	getToken := token.Value
	if getToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Token is None"})
		c.Abort()
		return
	}

	// TODO: 보내온 토큰이 현재 레디스에 저장되어 있는지 확인한 후, 이후 작업을 수행한다.
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(getToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "fail to validate token.",
		})
	}

	for key, value := range claims {
		if claims[key] ==
	}


	c.JSON(401, gin.H{
		"status":  200,
		"message": "success",
	})
}

func CustomTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
