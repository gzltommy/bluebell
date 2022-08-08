package controller

import "bluebell/model"

// _ResponsePostList 因为我们的接口文档返回的数据格式是一致的，但是具体的 data 类型不一致
type _ResponsePostList struct {
	Code    int                    `json:"code"`    // 业务响应状态码
	Message string                 `json:"message"` // 提示信息
	Data    []*model.ApiPostDetail `json:"data"`    // 数据
}

type _RespComment struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"` // omitempty 当 data 为空时,不展示这个字段
}
