package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type PrefixedRedisClient struct {
	*redis.Client
	prefix string
}

func (p *PrefixedRedisClient) doPrefix(key string) string {
	return p.prefix + key
}

func (p *PrefixedRedisClient) doPrefixKeys(keys []string) []string {
	k := make([]string, len(keys))
	for i, key := range keys {
		k[i] = p.doPrefix(key)
	}
	return k
}

func (p *PrefixedRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return p.Client.Del(ctx, p.doPrefixKeys(keys)...)
}

func (p *PrefixedRedisClient) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	return p.Client.Unlink(ctx, p.doPrefixKeys(keys)...)
}

func (p *PrefixedRedisClient) Dump(ctx context.Context, key string) *redis.StringCmd {
	return p.Client.Dump(ctx, p.doPrefix(key))
}

func (p *PrefixedRedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return p.Client.Exists(ctx, p.doPrefixKeys(keys)...)
}

func (p *PrefixedRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return p.Client.Expire(ctx, p.doPrefix(key), expiration)
}

func (p *PrefixedRedisClient) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	return p.Client.ExpireAt(ctx, p.doPrefix(key), tm)
}

func (p *PrefixedRedisClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return p.Client.Keys(ctx, p.doPrefix(pattern))
}

func (p *PrefixedRedisClient) Migrate(ctx context.Context, host string, port string, key string, db int, timeout time.Duration) *redis.StatusCmd {
	return p.Client.Migrate(ctx, host, port, p.doPrefix(key), db, timeout)
}

func (p *PrefixedRedisClient) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	return p.Client.Move(ctx, p.doPrefix(key), db)
}

func (p *PrefixedRedisClient) WithTimeout(timeout time.Duration) Redis {
	return &PrefixedRedisClient{p.Client.WithTimeout(timeout), p.prefix}
}

func (p *PrefixedRedisClient) WithContext(ctx context.Context) Redis {
	return &PrefixedRedisClient{p.Client.WithContext(ctx), p.prefix}
}
