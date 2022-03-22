package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"random.chars.jp/git/misskey/config"
)

var Instance Redis

// New returns pointer to a new Client with the parameters
// specified in the config file.
func New() Redis {
	c := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.External.Redis.Host,
			config.External.Redis.Port),
		DB:       config.External.Redis.DB,
		Password: config.External.Redis.Password,
	})

	if config.External.Redis.Prefix == "" {
		return &Client{c}
	} else {
		return &PrefixedRedisClient{c, config.External.Redis.Prefix}
	}
}

type Client struct {
	*redis.Client
}

func (r *Client) WithTimeout(timeout time.Duration) Redis {
	return &Client{r.Client.WithTimeout(timeout)}
}

func (r *Client) WithContext(ctx context.Context) Redis {
	return &Client{r.Client.WithContext(ctx)}
}

type Redis interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Unlink(ctx context.Context, keys ...string) *redis.IntCmd
	Dump(ctx context.Context, key string) *redis.StringCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Migrate(ctx context.Context, host string, port string, key string, db int, timeout time.Duration) *redis.StatusCmd
	Move(ctx context.Context, key string, db int) *redis.BoolCmd
	//ObjectRefCount(ctx context.Context, key string) *redis.IntCmd
	//ObjectEncoding(ctx context.Context, key string) *redis.StringCmd
	//ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd
	//Persist(ctx context.Context, key string) *redis.BoolCmd
	//PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	//PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	//PTTL(ctx context.Context, key string) *redis.DurationCmd
	//RandomKey(ctx context.Context) *redis.StringCmd
	//Rename(ctx context.Context, key string, newkey string) *redis.StatusCmd
	//RenameNX(ctx context.Context, key string, newkey string) *redis.BoolCmd
	//Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	//RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	//Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd
	//SortStore(ctx context.Context, key string, store string, sort *redis.Sort) *redis.IntCmd
	//SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd
	//Touch(ctx context.Context, keys ...string) *redis.IntCmd
	//TTL(ctx context.Context, key string) *redis.DurationCmd
	//Type(ctx context.Context, key string) *redis.StatusCmd
	//Append(ctx context.Context, key string, value string) *redis.IntCmd
	//Decr(ctx context.Context, key string) *redis.IntCmd
	//DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd
	//Get(ctx context.Context, key string) *redis.StringCmd
	//GetRange(ctx context.Context, key string, start int64, end int64) *redis.StringCmd
	//GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd
	//GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd
	//GetDel(ctx context.Context, key string) *redis.StringCmd
	//Incr(ctx context.Context, key string) *redis.IntCmd
	//IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd
	//IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd
	//MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	//MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	//MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd
	//Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	//SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd
	//SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	//SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	//SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	//SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd
	//StrLen(ctx context.Context, key string) *redis.IntCmd
	//GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd
	//SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd
	//BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd
	//BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	//BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	//BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	//BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd
	//BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd
	//BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd
	//Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	//ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd
	//SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	//HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	//ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	//HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	//HExists(ctx context.Context, key string, field string) *redis.BoolCmd
	//HGet(ctx context.Context, key string, field string) *redis.StringCmd
	//HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	//HIncrBy(ctx context.Context, key string, field string, incr int64) *redis.IntCmd
	//HIncrByFloat(ctx context.Context, key string, field string, incr float64) *redis.FloatCmd
	//HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	//HLen(ctx context.Context, key string) *redis.IntCmd
	//HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	//HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	//HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	//HSetNX(ctx context.Context, key string, field string, value interface{}) *redis.BoolCmd
	//HVals(ctx context.Context, key string) *redis.StringSliceCmd
	//HRandField(ctx context.Context, key string, count int, withValues bool) *redis.StringSliceCmd
	//BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	//BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	//BRPopLPush(ctx context.Context, source string, destination string, timeout time.Duration) *redis.StringCmd
	//LIndex(ctx context.Context, key string, index int64) *redis.StringCmd
	//LInsert(ctx context.Context, key string, op string, pivot interface{}, value interface{}) *redis.IntCmd
	//LInsertBefore(ctx context.Context, key string, pivot interface{}, value interface{}) *redis.IntCmd
	//LInsertAfter(ctx context.Context, key string, pivot interface{}, value interface{}) *redis.IntCmd
	//LLen(ctx context.Context, key string) *redis.IntCmd
	//LPop(ctx context.Context, key string) *redis.StringCmd
	//LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd
	//LPos(ctx context.Context, key string, value string, a redis.LPosArgs) *redis.IntCmd
	//LPosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *redis.IntSliceCmd
	//LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	//LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	//LRange(ctx context.Context, key string, start int64, stop int64) *redis.StringSliceCmd
	//LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	//LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd
	//LTrim(ctx context.Context, key string, start int64, stop int64) *redis.StatusCmd
	//RPop(ctx context.Context, key string) *redis.StringCmd
	//RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd
	//RPopLPush(ctx context.Context, source string, destination string) *redis.StringCmd
	//RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	//RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	//LMove(ctx context.Context, source string, destination string, srcpos string, destpos string) *redis.StringCmd
	//BLMove(ctx context.Context, source string, destination string, srcpos string, destpos string, timeout time.Duration) *redis.StringCmd
	//SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	//SCard(ctx context.Context, key string) *redis.IntCmd
	//SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd
	//SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	//SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd
	//SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	//SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd
	//SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd
	//SMembers(ctx context.Context, key string) *redis.StringSliceCmd
	//SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd
	//SMove(ctx context.Context, source string, destination string, member interface{}) *redis.BoolCmd
	//SPop(ctx context.Context, key string) *redis.StringCmd
	//SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	//SRandMember(ctx context.Context, key string) *redis.StringCmd
	//SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	//SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	//SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd
	//SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	//XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
	//XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd
	//XLen(ctx context.Context, stream string) *redis.IntCmd
	//XRange(ctx context.Context, stream string, start string, stop string) *redis.XMessageSliceCmd
	//XRangeN(ctx context.Context, stream string, start string, stop string, count int64) *redis.XMessageSliceCmd
	//XRevRange(ctx context.Context, stream string, start string, stop string) *redis.XMessageSliceCmd
	//XRevRangeN(ctx context.Context, stream string, start string, stop string, count int64) *redis.XMessageSliceCmd
	//XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd
	//XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd
	//XGroupCreate(ctx context.Context, stream string, group string, start string) *redis.StatusCmd
	//XGroupCreateMkStream(ctx context.Context, stream string, group string, start string) *redis.StatusCmd
	//XGroupSetID(ctx context.Context, stream string, group string, start string) *redis.StatusCmd
	//XGroupDestroy(ctx context.Context, stream string, group string) *redis.IntCmd
	//XGroupCreateConsumer(ctx context.Context, stream string, group string, consumer string) *redis.IntCmd
	//XGroupDelConsumer(ctx context.Context, stream string, group string, consumer string) *redis.IntCmd
	//XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd
	//XAck(ctx context.Context, stream string, group string, ids ...string) *redis.IntCmd
	//XPending(ctx context.Context, stream string, group string) *redis.XPendingCmd
	//XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd
	//XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd
	//XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd
	//XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd
	//XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd
	//XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	//XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	//XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	//XTrimMaxLenApprox(ctx context.Context, key string, maxLen int64, limit int64) *redis.IntCmd
	//XTrimMinID(ctx context.Context, key string, minID string) *redis.IntCmd
	//XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *redis.IntCmd
	//XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd
	//XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd
	//XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd
	//XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd
	//BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	//BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	//ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd
	//ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd
	//ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	//ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	//ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	//ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	//ZCard(ctx context.Context, key string) *redis.IntCmd
	//ZCount(ctx context.Context, key string, min string, max string) *redis.IntCmd
	//ZLexCount(ctx context.Context, key string, min string, max string) *redis.IntCmd
	//ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd
	//ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd
	//ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd
	//ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd
	//ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd
	//ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	//ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	//ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd
	//ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd
	//ZRange(ctx context.Context, key string, start int64, stop int64) *redis.StringSliceCmd
	//ZRangeWithScores(ctx context.Context, key string, start int64, stop int64) *redis.ZSliceCmd
	//ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	//ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	//ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	//ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd
	//ZRank(ctx context.Context, key string, member string) *redis.IntCmd
	//ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	//ZRemRangeByRank(ctx context.Context, key string, start int64, stop int64) *redis.IntCmd
	//ZRemRangeByScore(ctx context.Context, key string, min string, max string) *redis.IntCmd
	//ZRemRangeByLex(ctx context.Context, key string, min string, max string) *redis.IntCmd
	//ZRevRange(ctx context.Context, key string, start int64, stop int64) *redis.StringSliceCmd
	//ZRevRangeWithScores(ctx context.Context, key string, start int64, stop int64) *redis.ZSliceCmd
	//ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	//ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	//ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	//ZRevRank(ctx context.Context, key string, member string) *redis.IntCmd
	//ZScore(ctx context.Context, key string, member string) *redis.FloatCmd
	//ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd
	//ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd
	//ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd
	//ZRandMember(ctx context.Context, key string, count int, withScores bool) *redis.StringSliceCmd
	//ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd
	//ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd
	//ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	//PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd
	//PFCount(ctx context.Context, keys ...string) *redis.IntCmd
	//PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd
	//MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd

	Info(ctx context.Context, section ...string) *redis.StringCmd
	Save(ctx context.Context) *redis.StatusCmd
	Shutdown(ctx context.Context) *redis.StatusCmd
	Close() error
}
