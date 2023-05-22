package repo

import (
	"context"
	"time"
)

type Cache interface {
	// Put ctx防止超时，expire为过期时间
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
