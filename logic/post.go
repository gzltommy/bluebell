package logic

import (
	model2 "bluebell/dao/mysql/op"
	"bluebell/dao/redis/op"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func GetPostList(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {
	postList, err := model2.GetPostList(p.Page, p.Size)
	if err != nil {
		zap.L().Error("mysql.GetPostList(p.Page, p.Size) failed", zap.Error(err))
		return
	}
	data = make([]*model.ApiPostDetail, 0, len(postList)) // data 初始化
	for _, post := range postList {
		// 根据作者 id 查询作者信息
		user, err := model2.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区 id 查询社区详细信息
		community, err := model2.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postDetail := &model.ApiPostDetail{
			Post:            post,
			CommunityDetail: community,
			AuthorName:      user.UserName,
		}
		data = append(data, postDetail)
	}
	return
}

func CreatePost(p *model.ParamCreatePost) error {
	// 1、 生成 post_id (生成帖子ID)
	post := &model.Post{
		PostID:      uint64(snowflake.GenID()), //生成 post_id (生成帖子ID)
		AuthorId:    p.AuthorId,
		CommunityID: p.CommunityID,
		Status:      0,
		Title:       p.Title,
		Content:     p.Content,
		//CreateTime:  time.Time{},
	}

	// 2、创建帖子保存到数据库
	if err := model2.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	community, err := model2.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}

	// redis 存储帖子信息
	if err := op.CreatePost(
		post.PostID,
		post.AuthorId,
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityID); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return nil
}

func GetPostById(postID int64) (data *model.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	// 查询帖子信息
	post, err := model2.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed",
			zap.Int64("postID", postID),
			zap.Error(err))
		return nil, err
	}
	// 根据作者 id 查询作者信息
	user, err := model2.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed",
			zap.Uint64("postID", post.AuthorId),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := model2.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &model.ApiPostDetail{
		Post:            post,
		CommunityDetail: community,
		AuthorName:      user.UserName,
	}
	return
}

func GetPostListNew(p *model.ParamPostList2) (data []*model.ApiPostDetail, err error) {
	// 根据请求参数的不同,执行不同的业务逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区 id 查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return
}

func GetPostList2(p *model.ParamPostList2) (data []*model.ApiPostDetail, err error) {
	// 2、去 redis 查询 id 列表
	ids, err := op.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 提前查询好每篇帖子的投票数
	voteData, err := op.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := model2.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := model2.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := model2.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postDetail := &model.ApiPostDetail{
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
			AuthorName:      user.UserName,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *model.ParamPostList2) (data []*model.ApiPostDetail, err error) {
	// 2、去 redis 查询 id 列表
	ids, err := op.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostList(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 提前查询好每篇帖子的投票数
	voteData, err := op.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := model2.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := model2.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := model2.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postdetail := &model.ApiPostDetail{
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
			AuthorName:      user.UserName,
		}
		data = append(data, postdetail)
	}
	return
}
