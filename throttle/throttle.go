package throttle

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/XoliqberdiyevBehruz/wtc_backend/services/common"
// )

// func (h *common.Handler) isThrottled(ctx context.Context, key string, limit int, duration time.Duration) bool {
// 	count, err := h.redis.Incr(ctx, key).Result()
// 	if err != nil {
// 		fmt.Println("Redis error:", err)
// 		return true
// 	}

// 	if count == 1 {
// 		h.redis.Expire(ctx, key, duration)
// 	}

// 	return count > int64(limit)
// }