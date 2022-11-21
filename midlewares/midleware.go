package midlewares

import (
	"tubes/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(paramerter gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - [%s] %s %s %s %s  ",
			paramerter.ClientIP,
			paramerter.TimeStamp.Format(time.RFC3339),
			paramerter.MethodColor(),
			paramerter.Path,
			paramerter.StatusCodeColor(),
			paramerter.Latency)
	})
}

func GenerateIdUnix() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := uuid.New()
		ctx.Set("uuid", uid)
		fmt.Println("uuid : ", uid)
		ctx.Next()
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")
		// bearer ..........
		if !strings.HasPrefix(headerToken, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request doesn't contain access token"})
			ctx.Abort()
			return
		}
		token := strings.Split(headerToken, " ")
		if len(token) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request doesn't contain access token"})
			ctx.Abort()
			return
		}
		tokenstr := token[1]
		if tokenstr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request doesn't contain access token"})
			ctx.Abort()
			return
		}
		err := auth.ValidateToken(tokenstr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
