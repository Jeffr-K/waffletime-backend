package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://sosedoff.com/2014/12/21/gin-middleware.html
// TODO: 미들웨어 사용법
// TODO: 만약 Cookie 에 엑세스 토큰이 없다면? -> 요청 거부
// TODO: 만료된 경우, Refresh Token 을 사용해서 Access Token 을 재발급 받아야한다.

func TokenValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if len(accessToken) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed."})
			c.Abort()
			return
		}

		if permit := IsPermitTokenExpireTime(accessToken); permit == false {
			// permit 됐으면,
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Time is expired"})
		}

		// TODO:
		// 토큰 정보 파싱
		// redis 토큰 정보 파싱
		// 서로 같은 클레임이냐 여부 true => next() false => return not same user

		if valid := isValidToken(accessToken); valid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Time is expired"})
		}

		c.Next()
	}
}

// IsPermitTokenExpireTime 만료 시간
func IsPermitTokenExpireTime(accessToken string) bool {
	return true
}

func isValidToken(accessToken string) bool {
	return true
}

//func TokenValidationMiddleware(c *gin.Context) {
//	token, err := c.Request.Cookie("access-token")
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed."})
//		c.Abort()
//		return
//	}
//
//	getToken := token.Value
//	if getToken == "" {
//		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Token is None"})
//		c.Abort()
//		return
//	}
//
//	// TODO: 보내온 토큰이 현재 레디스에 저장되어 있는지 확인한 후, 이후 작업을 수행한다.
//	claims := jwt.MapClaims{}
//	_, err = jwt.ParseWithClaims(getToken, &claims, func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
//	})
//	if err != nil {
//		c.JSON(401, gin.H{
//			"status":  401,
//			"message": "fail to validate token.",
//		})
//	}
//
//	for key, value := range claims {
//		if claims[key] ==
//	}
//
//
//	c.JSON(401, gin.H{
//		"status":  200,
//		"message": "success",
//	})
//}
//
//func CustomTokenMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//func responseWithError(c *gin.Context, code int, message interface{}) {
//	c.AbortWithStatusJSON(code, gin.H{"error": message})
//}
