package shared

import (
	"github.com/gin-gonic/gin"
)

func WithUserRequestContext(ctx *gin.Context, userId int64) {
	ctx.Set("userId", userId)
}

func GetContextStreamingPlatformId(ctx *gin.Context) int64 {
	return ctx.GetInt64("userId")
}
