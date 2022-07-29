package redis

//import (
//	"github.com/go-redis/redis"
//	"strconv"
//	"time"
//)
//
//// CreatePost 使用 hash 存储帖子信息
//func CreatePost(postID, userID uint64, title, summary string, CommunityID uint64) (err error) {
//	now := float64(time.Now().Unix())
//	votedKey := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
//	communityKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(CommunityID))
//	postInfo := map[string]interface{}{
//		"title":    title,
//		"summary":  summary,
//		"post:id":  postID,
//		"user:id":  userID,
//		"time":     now,
//		"votes":    1,
//		"comments": 0,
//	}
//
//	// 事务操作
//	pipeline := client.TxPipeline()
//	pipeline.ZAdd(votedKey, redis.Z{ // 作者默认投赞成票
//		Score:  1,
//		Member: userID,
//	})
//	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds) // 一周时间
//
//	pipeline.HMSet(KeyPostInfoHashPrefix+strconv.Itoa(int(postID)), postInfo)
//	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{ // 添加到分数的ZSet
//		Score:  now + VoteScore,
//		Member: postID,
//	})
//	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{ // 添加到时间的ZSet
//		Score:  now,
//		Member: postID,
//	})
//	pipeline.SAdd(communityKey, postID) // 添加到对应版块  把帖子添加到社区的set
//	_, err = pipeline.Exec()
//	return
//}
