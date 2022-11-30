package midlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"tubes/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func Cors () gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Next()
	}
}