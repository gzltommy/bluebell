package logic

import (
	"bluebell/dao/redis/op"
	"bluebell/model"
	"go.uber.org/zap"
	"strconv"
)

func VoteForPost(userId uint64, p *model.ParamVote) error {
	zap.L().Debug("VoteForPost",
		zap.Uint64("userId", userId),
		zap.String("postId", p.PostID),
		zap.Int8("Direction", p.Direction))
	return op.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))
}
