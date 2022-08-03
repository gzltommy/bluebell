package redis

import (
	"github.com/go-redis/redis"
	"math"
	"time"
)

/*
投票算法：http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
本项目使用简化版的投票分数
投一票就加432分 86400/200 -> 200张赞成票就可以给帖子在首页续天  -> 《redis实战》
*/

/* PostVote 为帖子投票
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票		 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过反对票，现在要改为赞成票	 --> 更新分数和投票记录	差值的绝对值：2	+432*2
v=0时，有两种情况
	1.之前投过反对票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过赞成票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	-432
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票		 --> 更新分数和投票记录	差值的绝对值：1	-432
	2.之前投过赞成票，现在要改为反对票	 --> 更新分数和投票记录	差值的绝对值：2	-432*2

投票的限制：
每个帖子子发表之日起一个星期之内允许用户投票，超过一个星期就不允许投票了
	1、到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2、到期之后删除那个 KeyPostVotedZSetPrefix
*/

func VoteForPost(userID string, postID string, v float64) (err error) {
	// 1.判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds { // Unix()时间戳
		// 不允许投票了
		return ErrorVoteTimeExpire
	}
	// 2、更新帖子的分数
	// 2和3 需要放到一个pipeline事务中操作
	// 判断是否已经投过票 查当前用户给当前帖子的投票记录
	key := KeyPostVotedZSetPrefix + postID
	ov := client.ZScore(key, userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if v == ov {
		return ErrVoteRepested
	}
	var op float64
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diffAbs := math.Abs(ov - v)                                                        // 计算两次投票的差值
	pipeline := client.TxPipeline()                                                    // 事务操作
	_, err = pipeline.ZIncrBy(KeyPostScoreZSet, VoteScore*diffAbs*op, postID).Result() // 更新分数
	if ErrorVoteTimeExpire != nil {
		return err
	}
	// 3、记录用户为该帖子投票的数据
	if v == 0 {
		_, err = client.ZRem(key, postID).Result()
	} else {
		pipeline.ZAdd(key, redis.Z{ // 记录已投票
			Score:  v, // 赞成票还是反对票
			Member: userID,
		})
	}

	//switch math.Abs(ov) - math.Abs(v) {
	//case 1:
	//	// 取消投票 ov=1/-1 v=0
	//	// 投票数-1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", -1)
	//case 0:
	//	// 反转投票 ov=-1/1 v=1/-1
	//	// 投票数不用更新
	//case -1:
	//	// 新增投票 ov=0 v=1/-1
	//	// 投票数+1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	//default:
	//	// 已经投过票了
	//	return ErrorVoted
	//}
	_, err = pipeline.Exec()
	return err
}
