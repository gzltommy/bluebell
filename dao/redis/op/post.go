package op

import (
	redis2 "bluebell/dao/redis"
	"bluebell/model"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432 // 每一票的值 432 分
	PostPerAge               = 20
)

// CreatePost 使用 hash 存储帖子信息
func CreatePost(postID, userID uint64, title, summary string, communityID uint64) (err error) {
	now := float64(time.Now().Unix())
	votedKey := redis2.KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	communityKey := redis2.KeyCommunityPostSetPrefix + strconv.Itoa(int(communityID))
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	// 事务 pipeline
	pipeline := redis2.client.TxPipeline()
	pipeline.ZAdd(votedKey, redis.Z{ // 作者默认投赞成票
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds) // 一周时间

	pipeline.HMSet(redis2.KeyPostInfoHashPrefix+strconv.Itoa(int(postID)), postInfo)
	pipeline.ZAdd(redis2.KeyPostScoreZSet, redis.Z{ // 添加到分数的 ZSet
		Score:  now + VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(redis2.KeyPostTimeZSet, redis.Z{ // 添加到时间的 ZSet
		Score:  now,
		Member: postID,
	})
	pipeline.SAdd(communityKey, postID) // 添加到对应版块  把帖子添加到社区的 set
	_, err = pipeline.Exec()
	return
}

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3.ZREVRANGE 按照分数从大到小的顺序查询指定数量的元素
	return redis2.client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *model.ParamPostList2) ([]string, error) {
	// 从 redis 获取 id
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	key := redis2.KeyPostTimeZSet    // 默认是时间
	if p.Order == model.OrderScore { // 按照分数请求
		key = redis2.KeyPostScoreZSet
	}
	// 2.确定查询的索引起始点
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids{
	//	key := KeyPostVotedZSetPrefix + id
	//	// 查找key中分数是1的元素数量 -> 统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 使用 pipeline 一次发送多条命令减少 RTT
	pipeline := redis2.client.Pipeline()
	for _, id := range ids {
		key := redis2.KeyCommunityPostSetPrefix + id
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *model.ParamPostList2) ([]string, error) {
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	orderkey := redis2.KeyPostTimeZSet // 默认是时间
	if p.Order == model.OrderScore {   // 按照分数请求
		orderkey = redis2.KeyPostScoreZSet
	}

	// 使用zinterstore 把分区的帖子set与帖子分数的zset生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据

	// 社区的key
	cKey := redis2.KeyCommunityPostSetPrefix + strconv.Itoa(int(p.CommunityID))

	// 利用缓存key减少zinterstore执行的次数 缓存key
	key := orderkey + strconv.Itoa(int(p.CommunityID))
	if redis2.client.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipeline := redis2.client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX", // 将两个zset函数聚合的时候 求最大值
		}, cKey, orderkey) // zinterstore 计算
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在的就直接根据key查询ids
	return getIDsFormKey(key, p.Page, p.Size)
}
