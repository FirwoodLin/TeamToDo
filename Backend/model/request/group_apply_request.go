package request

import "time"

// GroupCodeRequest 创建邀请码请求
type GroupCodeRequest struct {
	GroupID   uint      `json:"-"`
	InviterID uint      `json:"-"`
	Code      string    `json:"-"`
	ExpireAt  time.Time `json:"expireAt"`
}
